/*
 * Copyright © 2020 Splunk, Inc.
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
 * Metadata Catalog
 *
 * With the Metadata Catalog in Splunk Cloud Services you can create and manage knowledge objects such as datasets, fields, rules, actions, dashboards, and workflows.
 *
 * API version: v2beta1.4 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package catalog

// GetDatasetQueryParams represents valid query parameters for the GetDataset operation
// For convenience GetDatasetQueryParams can be formed in a single statement, for example:
//     `v := GetDatasetQueryParams{}.SetMaxstale(...)`
type GetDatasetQueryParams struct {
	// Maxstale : The number of seconds beyond which we will refresh index metadata.
	Maxstale *int32 `key:"maxstale"`
}

func (q GetDatasetQueryParams) SetMaxstale(v int32) GetDatasetQueryParams {
	q.Maxstale = &v
	return q
}

// GetDatasetByIdQueryParams represents valid query parameters for the GetDatasetById operation
// For convenience GetDatasetByIdQueryParams can be formed in a single statement, for example:
//     `v := GetDatasetByIdQueryParams{}.SetMaxstale(...)`
type GetDatasetByIdQueryParams struct {
	// Maxstale : The number of seconds beyond which we will refresh index metadata.
	Maxstale *int32 `key:"maxstale"`
}

func (q GetDatasetByIdQueryParams) SetMaxstale(v int32) GetDatasetByIdQueryParams {
	q.Maxstale = &v
	return q
}

// ListActionsForRuleQueryParams represents valid query parameters for the ListActionsForRule operation
// For convenience ListActionsForRuleQueryParams can be formed in a single statement, for example:
//     `v := ListActionsForRuleQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListActionsForRuleQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListActionsForRuleQueryParams) SetCount(v int32) ListActionsForRuleQueryParams {
	q.Count = &v
	return q
}

func (q ListActionsForRuleQueryParams) SetFilter(v string) ListActionsForRuleQueryParams {
	q.Filter = v
	return q
}

func (q ListActionsForRuleQueryParams) SetOffset(v int32) ListActionsForRuleQueryParams {
	q.Offset = &v
	return q
}

func (q ListActionsForRuleQueryParams) SetOrderby(v []string) ListActionsForRuleQueryParams {
	q.Orderby = v
	return q
}

// ListActionsForRuleByIdQueryParams represents valid query parameters for the ListActionsForRuleById operation
// For convenience ListActionsForRuleByIdQueryParams can be formed in a single statement, for example:
//     `v := ListActionsForRuleByIdQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListActionsForRuleByIdQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListActionsForRuleByIdQueryParams) SetCount(v int32) ListActionsForRuleByIdQueryParams {
	q.Count = &v
	return q
}

func (q ListActionsForRuleByIdQueryParams) SetFilter(v string) ListActionsForRuleByIdQueryParams {
	q.Filter = v
	return q
}

func (q ListActionsForRuleByIdQueryParams) SetOffset(v int32) ListActionsForRuleByIdQueryParams {
	q.Offset = &v
	return q
}

func (q ListActionsForRuleByIdQueryParams) SetOrderby(v []string) ListActionsForRuleByIdQueryParams {
	q.Orderby = v
	return q
}

// ListAnnotationsQueryParams represents valid query parameters for the ListAnnotations operation
// For convenience ListAnnotationsQueryParams can be formed in a single statement, for example:
//     `v := ListAnnotationsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListAnnotationsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListAnnotationsQueryParams) SetCount(v int32) ListAnnotationsQueryParams {
	q.Count = &v
	return q
}

func (q ListAnnotationsQueryParams) SetFilter(v string) ListAnnotationsQueryParams {
	q.Filter = v
	return q
}

func (q ListAnnotationsQueryParams) SetOffset(v int32) ListAnnotationsQueryParams {
	q.Offset = &v
	return q
}

func (q ListAnnotationsQueryParams) SetOrderby(v []string) ListAnnotationsQueryParams {
	q.Orderby = v
	return q
}

// ListAnnotationsForDashboardByIdQueryParams represents valid query parameters for the ListAnnotationsForDashboardById operation
// For convenience ListAnnotationsForDashboardByIdQueryParams can be formed in a single statement, for example:
//     `v := ListAnnotationsForDashboardByIdQueryParams{}.SetFilter(...)`
type ListAnnotationsForDashboardByIdQueryParams struct {
	// Filter : A filter query to apply to the annotations.
	Filter string `key:"filter"`
}

func (q ListAnnotationsForDashboardByIdQueryParams) SetFilter(v string) ListAnnotationsForDashboardByIdQueryParams {
	q.Filter = v
	return q
}

// ListAnnotationsForDashboardByResourceNameQueryParams represents valid query parameters for the ListAnnotationsForDashboardByResourceName operation
// For convenience ListAnnotationsForDashboardByResourceNameQueryParams can be formed in a single statement, for example:
//     `v := ListAnnotationsForDashboardByResourceNameQueryParams{}.SetFilter(...)`
type ListAnnotationsForDashboardByResourceNameQueryParams struct {
	// Filter : A filter query to apply to the annotations.
	Filter string `key:"filter"`
}

func (q ListAnnotationsForDashboardByResourceNameQueryParams) SetFilter(v string) ListAnnotationsForDashboardByResourceNameQueryParams {
	q.Filter = v
	return q
}

// ListAnnotationsForDatasetByIdQueryParams represents valid query parameters for the ListAnnotationsForDatasetById operation
// For convenience ListAnnotationsForDatasetByIdQueryParams can be formed in a single statement, for example:
//     `v := ListAnnotationsForDatasetByIdQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListAnnotationsForDatasetByIdQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListAnnotationsForDatasetByIdQueryParams) SetCount(v int32) ListAnnotationsForDatasetByIdQueryParams {
	q.Count = &v
	return q
}

func (q ListAnnotationsForDatasetByIdQueryParams) SetFilter(v string) ListAnnotationsForDatasetByIdQueryParams {
	q.Filter = v
	return q
}

func (q ListAnnotationsForDatasetByIdQueryParams) SetOffset(v int32) ListAnnotationsForDatasetByIdQueryParams {
	q.Offset = &v
	return q
}

func (q ListAnnotationsForDatasetByIdQueryParams) SetOrderby(v []string) ListAnnotationsForDatasetByIdQueryParams {
	q.Orderby = v
	return q
}

// ListAnnotationsForDatasetByResourceNameQueryParams represents valid query parameters for the ListAnnotationsForDatasetByResourceName operation
// For convenience ListAnnotationsForDatasetByResourceNameQueryParams can be formed in a single statement, for example:
//     `v := ListAnnotationsForDatasetByResourceNameQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListAnnotationsForDatasetByResourceNameQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListAnnotationsForDatasetByResourceNameQueryParams) SetCount(v int32) ListAnnotationsForDatasetByResourceNameQueryParams {
	q.Count = &v
	return q
}

func (q ListAnnotationsForDatasetByResourceNameQueryParams) SetFilter(v string) ListAnnotationsForDatasetByResourceNameQueryParams {
	q.Filter = v
	return q
}

func (q ListAnnotationsForDatasetByResourceNameQueryParams) SetOffset(v int32) ListAnnotationsForDatasetByResourceNameQueryParams {
	q.Offset = &v
	return q
}

func (q ListAnnotationsForDatasetByResourceNameQueryParams) SetOrderby(v []string) ListAnnotationsForDatasetByResourceNameQueryParams {
	q.Orderby = v
	return q
}

// ListDashboardsQueryParams represents valid query parameters for the ListDashboards operation
// For convenience ListDashboardsQueryParams can be formed in a single statement, for example:
//     `v := ListDashboardsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListDashboardsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListDashboardsQueryParams) SetCount(v int32) ListDashboardsQueryParams {
	q.Count = &v
	return q
}

func (q ListDashboardsQueryParams) SetFilter(v string) ListDashboardsQueryParams {
	q.Filter = v
	return q
}

func (q ListDashboardsQueryParams) SetOffset(v int32) ListDashboardsQueryParams {
	q.Offset = &v
	return q
}

func (q ListDashboardsQueryParams) SetOrderby(v []string) ListDashboardsQueryParams {
	q.Orderby = v
	return q
}

// ListDatasetsQueryParams represents valid query parameters for the ListDatasets operation
// For convenience ListDatasetsQueryParams can be formed in a single statement, for example:
//     `v := ListDatasetsQueryParams{}.SetCount(...).SetFilter(...).SetMaxstale(...).SetOffset(...).SetOrderby(...)`
type ListDatasetsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the dataset list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Maxstale : The number of seconds beyond which we will refresh index metadata.
	Maxstale *int32 `key:"maxstale"`
	// Offset : The number of results to skip before the first result is returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc\&quot;.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListDatasetsQueryParams) SetCount(v int32) ListDatasetsQueryParams {
	q.Count = &v
	return q
}

func (q ListDatasetsQueryParams) SetFilter(v string) ListDatasetsQueryParams {
	q.Filter = v
	return q
}

func (q ListDatasetsQueryParams) SetMaxstale(v int32) ListDatasetsQueryParams {
	q.Maxstale = &v
	return q
}

func (q ListDatasetsQueryParams) SetOffset(v int32) ListDatasetsQueryParams {
	q.Offset = &v
	return q
}

func (q ListDatasetsQueryParams) SetOrderby(v []string) ListDatasetsQueryParams {
	q.Orderby = v
	return q
}

// ListFieldsQueryParams represents valid query parameters for the ListFields operation
// For convenience ListFieldsQueryParams can be formed in a single statement, for example:
//     `v := ListFieldsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListFieldsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the dataset list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListFieldsQueryParams) SetCount(v int32) ListFieldsQueryParams {
	q.Count = &v
	return q
}

func (q ListFieldsQueryParams) SetFilter(v string) ListFieldsQueryParams {
	q.Filter = v
	return q
}

func (q ListFieldsQueryParams) SetOffset(v int32) ListFieldsQueryParams {
	q.Offset = &v
	return q
}

func (q ListFieldsQueryParams) SetOrderby(v []string) ListFieldsQueryParams {
	q.Orderby = v
	return q
}

// ListFieldsForDatasetQueryParams represents valid query parameters for the ListFieldsForDataset operation
// For convenience ListFieldsForDatasetQueryParams can be formed in a single statement, for example:
//     `v := ListFieldsForDatasetQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListFieldsForDatasetQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the dataset list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by. You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListFieldsForDatasetQueryParams) SetCount(v int32) ListFieldsForDatasetQueryParams {
	q.Count = &v
	return q
}

func (q ListFieldsForDatasetQueryParams) SetFilter(v string) ListFieldsForDatasetQueryParams {
	q.Filter = v
	return q
}

func (q ListFieldsForDatasetQueryParams) SetOffset(v int32) ListFieldsForDatasetQueryParams {
	q.Offset = &v
	return q
}

func (q ListFieldsForDatasetQueryParams) SetOrderby(v []string) ListFieldsForDatasetQueryParams {
	q.Orderby = v
	return q
}

// ListFieldsForDatasetByIdQueryParams represents valid query parameters for the ListFieldsForDatasetById operation
// For convenience ListFieldsForDatasetByIdQueryParams can be formed in a single statement, for example:
//     `v := ListFieldsForDatasetByIdQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListFieldsForDatasetByIdQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the dataset list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by. You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListFieldsForDatasetByIdQueryParams) SetCount(v int32) ListFieldsForDatasetByIdQueryParams {
	q.Count = &v
	return q
}

func (q ListFieldsForDatasetByIdQueryParams) SetFilter(v string) ListFieldsForDatasetByIdQueryParams {
	q.Filter = v
	return q
}

func (q ListFieldsForDatasetByIdQueryParams) SetOffset(v int32) ListFieldsForDatasetByIdQueryParams {
	q.Offset = &v
	return q
}

func (q ListFieldsForDatasetByIdQueryParams) SetOrderby(v []string) ListFieldsForDatasetByIdQueryParams {
	q.Orderby = v
	return q
}

// ListModulesQueryParams represents valid query parameters for the ListModules operation
// For convenience ListModulesQueryParams can be formed in a single statement, for example:
//     `v := ListModulesQueryParams{}.SetFilter(...)`
type ListModulesQueryParams struct {
	// Filter : A filter to apply to the modules.
	Filter string `key:"filter"`
}

func (q ListModulesQueryParams) SetFilter(v string) ListModulesQueryParams {
	q.Filter = v
	return q
}

// ListRelationshipsQueryParams represents valid query parameters for the ListRelationships operation
// For convenience ListRelationshipsQueryParams can be formed in a single statement, for example:
//     `v := ListRelationshipsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListRelationshipsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListRelationshipsQueryParams) SetCount(v int32) ListRelationshipsQueryParams {
	q.Count = &v
	return q
}

func (q ListRelationshipsQueryParams) SetFilter(v string) ListRelationshipsQueryParams {
	q.Filter = v
	return q
}

func (q ListRelationshipsQueryParams) SetOffset(v int32) ListRelationshipsQueryParams {
	q.Offset = &v
	return q
}

func (q ListRelationshipsQueryParams) SetOrderby(v []string) ListRelationshipsQueryParams {
	q.Orderby = v
	return q
}

// ListRulesQueryParams represents valid query parameters for the ListRules operation
// For convenience ListRulesQueryParams can be formed in a single statement, for example:
//     `v := ListRulesQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListRulesQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListRulesQueryParams) SetCount(v int32) ListRulesQueryParams {
	q.Count = &v
	return q
}

func (q ListRulesQueryParams) SetFilter(v string) ListRulesQueryParams {
	q.Filter = v
	return q
}

func (q ListRulesQueryParams) SetOffset(v int32) ListRulesQueryParams {
	q.Offset = &v
	return q
}

func (q ListRulesQueryParams) SetOrderby(v []string) ListRulesQueryParams {
	q.Orderby = v
	return q
}

// ListWorkflowBuildsQueryParams represents valid query parameters for the ListWorkflowBuilds operation
// For convenience ListWorkflowBuildsQueryParams can be formed in a single statement, for example:
//     `v := ListWorkflowBuildsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListWorkflowBuildsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListWorkflowBuildsQueryParams) SetCount(v int32) ListWorkflowBuildsQueryParams {
	q.Count = &v
	return q
}

func (q ListWorkflowBuildsQueryParams) SetFilter(v string) ListWorkflowBuildsQueryParams {
	q.Filter = v
	return q
}

func (q ListWorkflowBuildsQueryParams) SetOffset(v int32) ListWorkflowBuildsQueryParams {
	q.Offset = &v
	return q
}

func (q ListWorkflowBuildsQueryParams) SetOrderby(v []string) ListWorkflowBuildsQueryParams {
	q.Orderby = v
	return q
}

// ListWorkflowRunsQueryParams represents valid query parameters for the ListWorkflowRuns operation
// For convenience ListWorkflowRunsQueryParams can be formed in a single statement, for example:
//     `v := ListWorkflowRunsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListWorkflowRunsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListWorkflowRunsQueryParams) SetCount(v int32) ListWorkflowRunsQueryParams {
	q.Count = &v
	return q
}

func (q ListWorkflowRunsQueryParams) SetFilter(v string) ListWorkflowRunsQueryParams {
	q.Filter = v
	return q
}

func (q ListWorkflowRunsQueryParams) SetOffset(v int32) ListWorkflowRunsQueryParams {
	q.Offset = &v
	return q
}

func (q ListWorkflowRunsQueryParams) SetOrderby(v []string) ListWorkflowRunsQueryParams {
	q.Orderby = v
	return q
}

// ListWorkflowsQueryParams represents valid query parameters for the ListWorkflows operation
// For convenience ListWorkflowsQueryParams can be formed in a single statement, for example:
//     `v := ListWorkflowsQueryParams{}.SetCount(...).SetFilter(...).SetOffset(...).SetOrderby(...)`
type ListWorkflowsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the results list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Offset : The number of results to skip before the first one returned.
	Offset *int32 `key:"offset"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}

func (q ListWorkflowsQueryParams) SetCount(v int32) ListWorkflowsQueryParams {
	q.Count = &v
	return q
}

func (q ListWorkflowsQueryParams) SetFilter(v string) ListWorkflowsQueryParams {
	q.Filter = v
	return q
}

func (q ListWorkflowsQueryParams) SetOffset(v int32) ListWorkflowsQueryParams {
	q.Offset = &v
	return q
}

func (q ListWorkflowsQueryParams) SetOrderby(v []string) ListWorkflowsQueryParams {
	q.Orderby = v
	return q
}
