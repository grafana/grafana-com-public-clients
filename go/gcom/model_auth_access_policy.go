/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"time"
)

// checks if the AuthAccessPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAccessPolicy{}

// AuthAccessPolicy struct for AuthAccessPolicy
type AuthAccessPolicy struct {
	Id    *string `json:"id,omitempty"`
	OrgId *string `json:"orgId,omitempty"`
	Name  string  `json:"name"`
	// Will be set to `name` if not provided.
	DisplayName *string `json:"displayName,omitempty"`
	// Source of the Access Policy (requires system token).
	Source     *string                       `json:"source,omitempty"`
	Scopes     []string                      `json:"scopes"`
	Realms     []AuthAccessPolicyRealmsInner `json:"realms"`
	CreatedAt  *time.Time                    `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time                    `json:"updatedAt,omitempty"`
	Conditions *AuthAccessPolicyConditions   `json:"conditions,omitempty"`
	Attributes *AuthAccessPolicyAttributes   `json:"attributes,omitempty"`
	// The status of the access policy.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthAccessPolicy AuthAccessPolicy

// NewAuthAccessPolicy instantiates a new AuthAccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAccessPolicy(name string, scopes []string, realms []AuthAccessPolicyRealmsInner) *AuthAccessPolicy {
	this := AuthAccessPolicy{}
	this.Name = name
	this.Scopes = scopes
	this.Realms = realms
	return &this
}

// NewAuthAccessPolicyWithDefaults instantiates a new AuthAccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAccessPolicyWithDefaults() *AuthAccessPolicy {
	this := AuthAccessPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthAccessPolicy) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *AuthAccessPolicy) SetOrgId(v string) {
	o.OrgId = &v
}

// GetName returns the Name field value
func (o *AuthAccessPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthAccessPolicy) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthAccessPolicy) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AuthAccessPolicy) SetSource(v string) {
	o.Source = &v
}

// GetScopes returns the Scopes field value
func (o *AuthAccessPolicy) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *AuthAccessPolicy) SetScopes(v []string) {
	o.Scopes = v
}

// GetRealms returns the Realms field value
func (o *AuthAccessPolicy) GetRealms() []AuthAccessPolicyRealmsInner {
	if o == nil {
		var ret []AuthAccessPolicyRealmsInner
		return ret
	}

	return o.Realms
}

// GetRealmsOk returns a tuple with the Realms field value
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetRealmsOk() ([]AuthAccessPolicyRealmsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Realms, true
}

// SetRealms sets field value
func (o *AuthAccessPolicy) SetRealms(v []AuthAccessPolicyRealmsInner) {
	o.Realms = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuthAccessPolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AuthAccessPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetConditions() AuthAccessPolicyConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret AuthAccessPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetConditionsOk() (*AuthAccessPolicyConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthAccessPolicyConditions and assigns it to the Conditions field.
func (o *AuthAccessPolicy) SetConditions(v AuthAccessPolicyConditions) {
	o.Conditions = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetAttributes() AuthAccessPolicyAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AuthAccessPolicyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetAttributesOk() (*AuthAccessPolicyAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AuthAccessPolicyAttributes and assigns it to the Attributes field.
func (o *AuthAccessPolicy) SetAttributes(v AuthAccessPolicyAttributes) {
	o.Attributes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthAccessPolicy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthAccessPolicy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthAccessPolicy) SetStatus(v string) {
	o.Status = &v
}

func (o AuthAccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAccessPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["scopes"] = o.Scopes
	toSerialize["realms"] = o.Realms
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthAccessPolicy) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varAuthAccessPolicy := _AuthAccessPolicy{}

	err = json.Unmarshal(data, &varAuthAccessPolicy)

	if err != nil {
		return err
	}

	*o = AuthAccessPolicy(varAuthAccessPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "source")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "realms")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthAccessPolicy struct {
	value *AuthAccessPolicy
	isSet bool
}

func (v NullableAuthAccessPolicy) Get() *AuthAccessPolicy {
	return v.value
}

func (v *NullableAuthAccessPolicy) Set(val *AuthAccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAccessPolicy(val *AuthAccessPolicy) *NullableAuthAccessPolicy {
	return &NullableAuthAccessPolicy{value: val, isSet: true}
}

func (v NullableAuthAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
