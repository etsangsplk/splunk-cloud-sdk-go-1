/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package idp

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Supported authentication flows
//
//   Authorization code with PKCE (pkce) -- known/trusted app
//      client_id + code challenge + redirect_uri + username + password =>
//          access, id_token, refresh_token[*]
//
//   Client credentials (client) -- private service
//      client_id + client_secret + custom scope =>
//          access
//
//   * refresh_token is available if enabled and if the offline_access scope
//     is requested.
//
// Note: pkce is normally browser based and involves redirection. Support is
// also provided for IdPs with extensions to the standard OIDC flows that allow
// client code to first authenticate with user credentials against a "primary"
// endpoint (/authn) and retrieve a one time session token, which when used with
// these flows, will result in the requested grants being returned directly in
// the redirect url.

// OIDCScope defines scopes that are OpenID Connect compatible, see: https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims
type OIDCScope string

const (
	// ScopeOpenID - The basic (and required) scope for OpenID Connect
	ScopeOpenID OIDCScope = "openid"
	// ScopeEmail - This scope value requests access to the email and email_verified Claims
	ScopeEmail OIDCScope = "email"
	// ScopeProfile - This scope value requests access to the End-User's default profile Claims, which are: name, family_name,
	// given_name, middle_name, nickname, preferred_username, profile, picture, website, gender, birthdate, zoneinfo,
	// locale, and updated_at
	ScopeProfile OIDCScope = "profile"
	// ScopeAddress - This scope value requests access to the address Claim
	ScopeAddress OIDCScope = "address"
	// ScopePhone - This scope value requests access to the phone_number and phone_number_verified Claims
	ScopePhone OIDCScope = "phone"
	// ScopeOffline - This scope value requests that an OAuth 2.0 Refresh Token be issued that can be used to obtain an Access Token that grants access to the End-User's UserInfo Endpoint even when the End-User is not present (not logged in).
	ScopeOffline OIDCScope = "offline_access"
)

var (
	// DefaultOIDCScopes defines the default OpenID Connect scopes to use in authn requests - "openid email profile"
	DefaultOIDCScopes   = fmt.Sprintf("%s %s %s", ScopeOpenID, ScopeEmail, ScopeProfile)
	DefaultRefreshScope = fmt.Sprintf("%s %s %s %s", ScopeOpenID, ScopeEmail, ScopeOffline, ScopeProfile)
)

// Read and decode JSON data from given reader and return as a map.
func load(r io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Returns the requested string value from the given map.
func gets(m map[string]interface{}, key string) (string, error) {
	value, ok := m[key]
	if !ok {
		return "", fmt.Errorf("key error: '%s'", key)
	}
	result, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("type error: %T", value)
	}
	return result, nil
}

// HTTPError Represents an error response
type HTTPError struct {
	StatusCode int                    `json:"status,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty"`
}

// Error handles marshalling of the HttpError to error type
func (e *HTTPError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Returns a golang error corresponding to the given http response.
func httpError(response *http.Response) error {
	var result = &HTTPError{StatusCode: response.StatusCode}

	// ignore if we cant read body details
	_ = json.NewDecoder(response.Body).Decode(&result.Body)
	return result
}

// Context Represents an authentication "context", which is the result of a
// successful OAuth authentication flow.
type Context struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	StartTime    int64
}

const (
	defaultAuthnPath     = "authn"
	defaultAuthorizePath = "authorize"
	defaultTokenPath     = "token"
)

// Client captures url and route information for the IdP endpoints
type Client struct {
	ProviderHost  string
	AuthnPath     string
	AuthorizePath string
	TokenPath     string
}

// NewClient Returns a new IdP client object.
//   providerURL: should be of the form https://example.com or optionally https://example.com:port
func NewClient(providerURL string, authnPath string, authorizePath string, tokenPath string) *Client {
	// Add a trailing slash if none
	if providerURL[len(providerURL)-1:] != "/" {
		providerURL = providerURL + "/"
	}
	if authnPath == "" {
		authnPath = defaultAuthnPath
	}
	if authorizePath == "" {
		authorizePath = defaultAuthorizePath
	}
	if tokenPath == "" {
		tokenPath = defaultTokenPath
	}
	return &Client{
		ProviderHost:  providerURL,
		AuthnPath:     authnPath,
		AuthorizePath: authorizePath,
		TokenPath:     tokenPath,
	}
}

// Returns a new HTTP client object with redirects disabled.
func newHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}}
}

// Decode and return the given response body.
func decode(response *http.Response) (*Context, error) {
	context := &Context{}
	if err := json.NewDecoder(response.Body).Decode(&context); err != nil {
		return nil, err
	}
	now := time.Now()
	context.StartTime = now.Unix()
	return context, nil
}

func newGet(reqURL string, params url.Values) (*http.Request, error) {
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.URL.RawQuery = params.Encode()
	return request, nil
}

func newFormPost(reqURL string, data url.Values) (*http.Request, error) {
	reader := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", reqURL, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

func get(reqURL string, params url.Values) (*http.Response, error) {
	request, err := newGet(reqURL, params)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

// Encode the given value and return its reader.
func encode(value interface{}) (*strings.Reader, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(data)), nil
}

func newPost(reqURL string, body interface{}) (*http.Request, error) {
	reader, err := encode(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", reqURL, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func post(reqURL string, body interface{}) (*http.Response, error) {
	request, err := newPost(reqURL, body)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

func formPost(reqURL string, data url.Values) (*http.Response, error) {
	request, err := newFormPost(reqURL, data)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

// Returns a synthetic state value.
func state() string {
	result, _ := time.Now().MarshalText()
	return string(result)
}

// Return a full URL based on the given host and path template.
func (c *Client) makeURL(path string) string {
	return fmt.Sprintf("%s%s", c.ProviderHost, path)
}

// ClientFlow will authenticate using the "client credentials" flow.
func (c *Client) ClientFlow(clientID, clientSecret, scope string) (*Context, error) {
	form := url.Values{
		"grant_type": {"client_credentials"},
		"scope":      {scope}}
	request, err := newFormPost(c.makeURL(c.TokenPath), form)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(clientID, clientSecret)
	response, err := newHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, httpError(response)
	}
	return decode(response)
}

// GetSessionToken Returns a one-time session token by authenticating using a
// "primary" endpoint (/authn).
func (c *Client) GetSessionToken(username, password string) (string, error) {
	body := map[string]string{
		"username": username,
		"password": password}

	response, err := post(c.makeURL(c.AuthnPath), body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	data, err := load(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		msg := response.Status
		if errMsg, found := data["description"]; found {
			msg = fmt.Sprintf("%s %s", response.Status, errMsg)
		}
		return "", errors.New(msg)
	}
	status, err := gets(data, "status")
	if err != nil {
		return "", err
	}
	if status != "SUCCESS" { // eg: LOCKED_OUT
		return "", errors.New(status)
	}
	sessionToken, err := gets(data, "sessionToken")
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}

// Returns a codeVerfier and codeChallenge for use in a PKCE flow.
func createCodeChallenge(n int) (string, string, error) {
	const safe = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._~"
	if n < 43 || n > 128 {
		return "", "", errors.New("invalid argument")
	}
	buff := make([]byte, n)
	if _, err := rand.Read(buff); err != nil {
		return "", "", err
	}
	nsafe := byte(len(safe))
	for i, b := range buff {
		b = b % nsafe
		buff[i] = safe[b]
	}
	cv := base64.RawURLEncoding.EncodeToString(buff)
	s256 := sha256.Sum256([]byte(cv))
	cc := base64.RawURLEncoding.EncodeToString(s256[:])
	return cv, cc, nil
}

// PKCEFlow will authenticate using the "proof key for code exchange" flow.
func (c *Client) PKCEFlow(clientID, redirectURI, scope, username, password string) (*Context, error) {
	// retrieve one-time session token
	sessionToken, err := c.GetSessionToken(username, password)
	if err != nil {
		return nil, err
	}

	cv, cc, err := createCodeChallenge(50)

	if err != nil {
		return nil, err
	}

	// request authorization code
	params := url.Values{
		"client_id":             {clientID},
		"code_challenge":        {cc},
		"code_challenge_method": {"S256"},
		"nonce":                 {"none"},
		"redirect_uri":          {redirectURI},
		"response_type":         {"code"},
		"scope":                 {scope},
		"session_token":         {sessionToken},
		"state":                 {state()}}
	response, err := get(c.makeURL(c.AuthorizePath), params)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusFound {
		return nil, httpError(response)
	}

	// retrieve the authorization code from the redirect url query string
	location := response.Header.Get("Location")
	locationURL, err := url.Parse(location)
	if err != nil {
		return nil, err
	}
	code := locationURL.Query().Get("code")

	// exchange authorization code for token(s)
	form := url.Values{
		"client_id":     {clientID},
		"code":          {code},
		"code_verifier": {cv},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirectURI}}
	response, err = formPost(c.makeURL(c.TokenPath), form)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, httpError(response)
	}

	return decode(response)
}

// Refresh will authenticate using a refresh token.
func (c *Client) Refresh(clientID, scope, refreshToken string) (*Context, error) {
	form := url.Values{
		"client_id":     {clientID},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"scope":         {scope}}
	response, err := formPost(c.makeURL(c.TokenPath), form)
	if err != nil {
		return nil, err
	}
	return decode(response)
}
