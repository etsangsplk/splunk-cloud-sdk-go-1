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

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

// ParseResponse parses json-formatted http response and decodes it into pre-defined models
func ParseResponse(model interface{}, response *http.Response) error {
	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		return errors.New("model parameter to ParseResponse() must be a pointer")
	}

	if response == nil {
		return errors.New("nil response provided")
	}
	if response.StatusCode == 204 {
		return nil
	}
	err := json.NewDecoder(response.Body).Decode(model)
	return err
}

// ParseURLParams parses a struct into url params based on its "key" tag
// It parses basic values and slices, and will parse structs recursively
func ParseURLParams(model interface{}) url.Values {
	values := url.Values{}
	if model == nil {
		return nil
	}
	indirect := reflect.Indirect(reflect.ValueOf(model))
	// must be a struct, this should always be a *QueryParams model
	if indirect.Kind() != reflect.Struct {
		return values
	}
	return toURLValues("", reflect.ValueOf(model), true)
}

func toURLValues(keyName string, value reflect.Value, explode bool) url.Values {
	values := url.Values{}
	// If pointer, follow the pointer
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return values
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		// explode=true: repeat k=v pairs e.g. ?k1=v1&k1=v2&k1=v3&...
		// explode=false: comma separate values e.g. ?k1=v1,v2,v3,...
		vals := []string{}
		for i := 0; i < value.Len(); i++ {
			uvals := toURLValues(keyName, value.Index(i), true)
			for uk, vslice := range uvals {
				// don't support a slice of maps
				if uk != keyName {
					continue
				}
				for _, v := range vslice {
					if explode {
						values.Add(uk, v)
					} else {
						vals = append(vals, v)
					}
				}
			}
		}
		if !explode && len(vals) != 0 {
			values.Set(keyName, strings.Join(vals, ","))
		}
	case reflect.Struct:
		// explode=true: ignore keyName and add k=v pairs for fields e.g. ?k1=v1&k2=v2&k3=v3&...
		// explode=false: comma separate keys and values e.g. ?keyName=k1,v1,k2,v2,...
		keyvals := []string{}
		for f := 0; f < value.Type().NumField(); f++ {
			field := value.Type().Field(f)
			fval := value.FieldByName(field.Name)
			// If explode is not specified, maps and arrays should be comma-separated
			fx := false
			if v, ok := field.Tag.Lookup("explode"); ok && v == "true" {
				fx = true
			}
			// If `key:` tag is specified, use it otherwise default to field name
			fkey := field.Name
			if v, ok := field.Tag.Lookup("key"); ok {
				fkey = v
			}
			uvals := toURLValues(fkey, fval, fx)
			for uk, vslice := range uvals {
				for _, v := range vslice {
					if explode {
						values.Add(uk, v)
					} else {
						keyvals = append(keyvals, uk)
						keyvals = append(keyvals, v)
					}
				}
			}
		}
		if !explode && len(keyvals) != 0 {
			values.Set(keyName, strings.Join(keyvals, ","))
		}
	case reflect.Map:
		// explode=true: ignore keyName and add k=v pairs for fields e.g. ?k1=v1&k2=v2&k3=v3&...
		// explode=false: comma separate keys and values e.g. ?keyName=k1,v1,k2,v2,...
		keyvals := []string{}
		for _, k := range value.MapKeys() {
			key := fmt.Sprintf("%v", k.Interface())
			if len(key) == 0 {
				continue
			}
			mv := value.MapIndex(k)
			uvals := toURLValues(key, mv, true)
			for uk, vslice := range uvals {
				for _, v := range vslice {
					if explode {
						values.Set(uk, v)
					} else {
						keyvals = append(keyvals, uk)
						keyvals = append(keyvals, v)
					}
				}
			}
		}
		if !explode && len(keyvals) != 0 {
			values.Set(keyName, strings.Join(keyvals, ","))
		}
	default:
		if val := fmt.Sprintf("%v", value.Interface()); len(val) > 0 {
			values.Set(keyName, val)
		}
	}
	return values
}

// ParseTemplatedPath parses a url-like path using an input template and outputs a map of matched values.
// template should place the named params to be extracted in single braces, e.g.
// ParseTemplatedPath("{my_param1}/literal/path/{my_param2}/parts", "foo/literal/path/bar/parts")
// returns {"my_param1": "foo", "my_param2": "bar"}
func ParseTemplatedPath(template string, path string) (map[string]string, error) {
	// escape any needed strings for regex
	template = regexp.QuoteMeta(template)
	// convert braces to capture groups
	template = strings.Replace(template, `\{`, `(?P<`, -1)
	template = "^" + strings.Replace(template, `\}`, `>[a-zA-z0-9_\-]+)`, -1)
	rex, err := regexp.Compile(template)
	if err != nil {
		return nil, err
	}
	match := rex.FindStringSubmatch(path)
	names := rex.SubexpNames()
	params := map[string]string{}
	for i, name := range names {
		if i != 0 && len(match) > i {
			params[name] = match[i]
		}
	}
	return params, nil
}
