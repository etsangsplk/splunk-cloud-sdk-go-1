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
 * App Registry
 *
 * With the Splunk Cloud App Registry service, you can create, update, and manage apps built with Splunk Developer Cloud.
 *
 * API version: v1beta2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package appregistry

// ListSubscriptionsQueryParams represents valid query parameters for the ListSubscriptions operation
// For convenience ListSubscriptionsQueryParams can be formed in a single statement, for example:
//     `v := ListSubscriptionsQueryParams{}.SetKind(...)`
type ListSubscriptionsQueryParams struct {
	// Kind : The kind of application.
	Kind *AppResourceKind `key:"kind"`
}

func (q ListSubscriptionsQueryParams) SetKind(v AppResourceKind) ListSubscriptionsQueryParams {
	q.Kind = &v
	return q
}
