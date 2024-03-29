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

// checks if the PostTokensRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTokensRequest{}

// PostTokensRequest struct for PostTokensRequest
type PostTokensRequest struct {
	AccessPolicyId       string     `json:"accessPolicyId"`
	DisplayName          *string    `json:"displayName,omitempty"`
	ExpiresAt            *time.Time `json:"expiresAt,omitempty"`
	Name                 string     `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _PostTokensRequest PostTokensRequest

// NewPostTokensRequest instantiates a new PostTokensRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTokensRequest(accessPolicyId string, name string) *PostTokensRequest {
	this := PostTokensRequest{}
	this.AccessPolicyId = accessPolicyId
	this.Name = name
	return &this
}

// NewPostTokensRequestWithDefaults instantiates a new PostTokensRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTokensRequestWithDefaults() *PostTokensRequest {
	this := PostTokensRequest{}
	return &this
}

// GetAccessPolicyId returns the AccessPolicyId field value
func (o *PostTokensRequest) GetAccessPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessPolicyId
}

// GetAccessPolicyIdOk returns a tuple with the AccessPolicyId field value
// and a boolean to check if the value has been set.
func (o *PostTokensRequest) GetAccessPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessPolicyId, true
}

// SetAccessPolicyId sets field value
func (o *PostTokensRequest) SetAccessPolicyId(v string) {
	o.AccessPolicyId = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PostTokensRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokensRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PostTokensRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PostTokensRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PostTokensRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokensRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PostTokensRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *PostTokensRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetName returns the Name field value
func (o *PostTokensRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostTokensRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostTokensRequest) SetName(v string) {
	o.Name = v
}

func (o PostTokensRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTokensRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessPolicyId"] = o.AccessPolicyId
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostTokensRequest) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varPostTokensRequest := _PostTokensRequest{}

	err = json.Unmarshal(data, &varPostTokensRequest)

	if err != nil {
		return err
	}

	*o = PostTokensRequest(varPostTokensRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPolicyId")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostTokensRequest struct {
	value *PostTokensRequest
	isSet bool
}

func (v NullablePostTokensRequest) Get() *PostTokensRequest {
	return v.value
}

func (v *NullablePostTokensRequest) Set(val *PostTokensRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTokensRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTokensRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTokensRequest(val *PostTokensRequest) *NullablePostTokensRequest {
	return &NullablePostTokensRequest{value: val, isSet: true}
}

func (v NullablePostTokensRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTokensRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
