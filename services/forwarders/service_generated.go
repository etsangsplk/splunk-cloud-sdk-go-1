/*
 * Copyright © 2019 Splunk, Inc.
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
 *
 * Splunk Forwarder Service
 *
 * Send data from a heavy forwarder to the Splunk Forwarder Service.
 *
 * API version: v2beta1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package forwarders

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new forwarders service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
	AddCertificate - Adds a certificate to a vacant slot on a tenant.
	Each tenant can have up to five certificates.
	Parameters:
		certificate
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) AddCertificate(certificate Certificate, resp ...*http.Response) (*CertificateInfo, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/forwarders/v2beta1/certificates`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: certificate})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb CertificateInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteCertificate - Removes a certificate on a particular slot on a tenant.
	Parameters:
		slot
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteCertificate(slot string, resp ...*http.Response) error {
	pp := struct {
		Slot string
	}{
		Slot: slot,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/forwarders/v2beta1/certificates/{{.Slot}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	DeleteCertificates - Removes all certificates on a tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteCertificates(resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/forwarders/v2beta1/certificates`, nil)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	ListCertificates - Returns a list of all certificates for a tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListCertificates(resp ...*http.Response) ([]CertificateInfo, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/forwarders/v2beta1/certificates`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []CertificateInfo
	err = util.ParseResponse(&rb, response)
	return rb, err
}
