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
 * Identity and Access Control
 *
 * With the Splunk Cloud Identity and Access Control (IAC) Service, you can authenticate and authorize Splunk API users.
 *
 * API version: v2beta1.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package identity

type AddGroupMemberBody struct {
	Name string `json:"name"`
}

type AddGroupRoleBody struct {
	Name string `json:"name"`
}

type AddMemberBody struct {
	Name string `json:"name"`
}

type AddRolePermissionBody string

type CreateGroupBody struct {
	Name string `json:"name"`
}

type CreateRoleBody struct {
	Name string `json:"name"`
}

type Group struct {
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
	MemberCount int32  `json:"memberCount"`
	Name        string `json:"name"`
	RoleCount   int32  `json:"roleCount"`
	Tenant      string `json:"tenant"`
}

// Represents a member that belongs to a group
type GroupMember struct {
	AddedAt   string `json:"addedAt"`
	AddedBy   string `json:"addedBy"`
	Group     string `json:"group"`
	Principal string `json:"principal"`
	Tenant    string `json:"tenant"`
}

// Represents a role that is assigned to a group
type GroupRole struct {
	AddedAt string `json:"addedAt"`
	AddedBy string `json:"addedBy"`
	Group   string `json:"group"`
	Role    string `json:"role"`
	Tenant  string `json:"tenant"`
}

// Represents a member that belongs to a tenant.
type Member struct {
	// When the principal was added to the tenant.
	AddedAt string            `json:"addedAt"`
	AddedBy string            `json:"addedBy"`
	Name    string            `json:"name"`
	Tenant  string            `json:"tenant"`
	Profile *PrincipalProfile `json:"profile,omitempty"`
}

type Principal struct {
	CreatedAt string            `json:"createdAt"`
	CreatedBy string            `json:"createdBy"`
	Kind      PrincipalKind     `json:"kind"`
	Name      string            `json:"name"`
	Tenants   []string          `json:"tenants"`
	UpdatedAt string            `json:"updatedAt"`
	UpdatedBy string            `json:"updatedBy"`
	Profile   *PrincipalProfile `json:"profile,omitempty"`
}

type PrincipalKind string

// List of PrincipalKind
const (
	PrincipalKindUser           PrincipalKind = "user"
	PrincipalKindServiceAccount PrincipalKind = "service_account"
)

// Profile information for a principal
type PrincipalProfile struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
}

type Role struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	Name      string `json:"name"`
	Tenant    string `json:"tenant"`
}

type RolePermission struct {
	AddedAt    string `json:"addedAt"`
	AddedBy    string `json:"addedBy"`
	Permission string `json:"permission"`
	Role       string `json:"role"`
	Tenant     string `json:"tenant"`
}

type Tenant struct {
	CreatedAt string       `json:"createdAt"`
	CreatedBy string       `json:"createdBy"`
	Name      string       `json:"name"`
	Status    TenantStatus `json:"status"`
}

type TenantName string

type TenantStatus string

// List of TenantStatus
const (
	TenantStatusProvisioning TenantStatus = "provisioning"
	TenantStatusFailed       TenantStatus = "failed"
	TenantStatusReady        TenantStatus = "ready"
	TenantStatusDeleting     TenantStatus = "deleting"
	TenantStatusDeleted      TenantStatus = "deleted"
	TenantStatusSuspended    TenantStatus = "suspended"
)

type ValidateInfo struct {
	ClientId  string     `json:"clientId"`
	Name      string     `json:"name"`
	Principal *Principal `json:"principal,omitempty"`
	Tenant    *Tenant    `json:"tenant,omitempty"`
}
