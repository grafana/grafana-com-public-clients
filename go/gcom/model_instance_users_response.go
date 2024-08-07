/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
)

// checks if the InstanceUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceUsersResponse{}

// InstanceUsersResponse struct for InstanceUsersResponse
type InstanceUsersResponse struct {
	Items                []ItemsInner1 `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _InstanceUsersResponse InstanceUsersResponse

// NewInstanceUsersResponse instantiates a new InstanceUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceUsersResponse(items []ItemsInner1) *InstanceUsersResponse {
	this := InstanceUsersResponse{}
	this.Items = items
	return &this
}

// NewInstanceUsersResponseWithDefaults instantiates a new InstanceUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceUsersResponseWithDefaults() *InstanceUsersResponse {
	this := InstanceUsersResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *InstanceUsersResponse) GetItems() []ItemsInner1 {
	if o == nil {
		var ret []ItemsInner1
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InstanceUsersResponse) GetItemsOk() ([]ItemsInner1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InstanceUsersResponse) SetItems(v []ItemsInner1) {
	o.Items = v
}

func (o InstanceUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstanceUsersResponse) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varInstanceUsersResponse := _InstanceUsersResponse{}

	err = json.Unmarshal(data, &varInstanceUsersResponse)

	if err != nil {
		return err
	}

	*o = InstanceUsersResponse(varInstanceUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstanceUsersResponse struct {
	value *InstanceUsersResponse
	isSet bool
}

func (v NullableInstanceUsersResponse) Get() *InstanceUsersResponse {
	return v.value
}

func (v *NullableInstanceUsersResponse) Set(val *InstanceUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceUsersResponse(val *InstanceUsersResponse) *NullableInstanceUsersResponse {
	return &NullableInstanceUsersResponse{value: val, isSet: true}
}

func (v NullableInstanceUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
