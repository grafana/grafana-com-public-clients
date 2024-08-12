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

// checks if the GrafanaNewApiKeyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrafanaNewApiKeyResult{}

// GrafanaNewApiKeyResult struct for GrafanaNewApiKeyResult
type GrafanaNewApiKeyResult struct {
	Id                   *int64  `json:"id,omitempty"`
	Key                  *string `json:"key,omitempty"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrafanaNewApiKeyResult GrafanaNewApiKeyResult

// NewGrafanaNewApiKeyResult instantiates a new GrafanaNewApiKeyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrafanaNewApiKeyResult() *GrafanaNewApiKeyResult {
	this := GrafanaNewApiKeyResult{}
	return &this
}

// NewGrafanaNewApiKeyResultWithDefaults instantiates a new GrafanaNewApiKeyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaNewApiKeyResultWithDefaults() *GrafanaNewApiKeyResult {
	this := GrafanaNewApiKeyResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GrafanaNewApiKeyResult) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaNewApiKeyResult) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrafanaNewApiKeyResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GrafanaNewApiKeyResult) SetId(v int64) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GrafanaNewApiKeyResult) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaNewApiKeyResult) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GrafanaNewApiKeyResult) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GrafanaNewApiKeyResult) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GrafanaNewApiKeyResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaNewApiKeyResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GrafanaNewApiKeyResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GrafanaNewApiKeyResult) SetName(v string) {
	o.Name = &v
}

func (o GrafanaNewApiKeyResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrafanaNewApiKeyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrafanaNewApiKeyResult) UnmarshalJSON(data []byte) (err error) {
	varGrafanaNewApiKeyResult := _GrafanaNewApiKeyResult{}

	err = json.Unmarshal(data, &varGrafanaNewApiKeyResult)

	if err != nil {
		return err
	}

	*o = GrafanaNewApiKeyResult(varGrafanaNewApiKeyResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrafanaNewApiKeyResult struct {
	value *GrafanaNewApiKeyResult
	isSet bool
}

func (v NullableGrafanaNewApiKeyResult) Get() *GrafanaNewApiKeyResult {
	return v.value
}

func (v *NullableGrafanaNewApiKeyResult) Set(val *GrafanaNewApiKeyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGrafanaNewApiKeyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaNewApiKeyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaNewApiKeyResult(val *GrafanaNewApiKeyResult) *NullableGrafanaNewApiKeyResult {
	return &NullableGrafanaNewApiKeyResult{value: val, isSet: true}
}

func (v NullableGrafanaNewApiKeyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaNewApiKeyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}