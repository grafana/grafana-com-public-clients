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

// checks if the Payload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payload{}

// Payload struct for Payload
type Payload struct {
	WithAddons           []string `json:"withAddons,omitempty"`
	LicenseAllPlugins    *bool    `json:"licenseAllPlugins,omitempty"`
	CaseId               *string  `json:"caseId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Payload Payload

// NewPayload instantiates a new Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayload() *Payload {
	this := Payload{}
	return &this
}

// NewPayloadWithDefaults instantiates a new Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadWithDefaults() *Payload {
	this := Payload{}
	return &this
}

// GetWithAddons returns the WithAddons field value if set, zero value otherwise.
func (o *Payload) GetWithAddons() []string {
	if o == nil || IsNil(o.WithAddons) {
		var ret []string
		return ret
	}
	return o.WithAddons
}

// GetWithAddonsOk returns a tuple with the WithAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetWithAddonsOk() ([]string, bool) {
	if o == nil || IsNil(o.WithAddons) {
		return nil, false
	}
	return o.WithAddons, true
}

// HasWithAddons returns a boolean if a field has been set.
func (o *Payload) HasWithAddons() bool {
	if o != nil && !IsNil(o.WithAddons) {
		return true
	}

	return false
}

// SetWithAddons gets a reference to the given []string and assigns it to the WithAddons field.
func (o *Payload) SetWithAddons(v []string) {
	o.WithAddons = v
}

// GetLicenseAllPlugins returns the LicenseAllPlugins field value if set, zero value otherwise.
func (o *Payload) GetLicenseAllPlugins() bool {
	if o == nil || IsNil(o.LicenseAllPlugins) {
		var ret bool
		return ret
	}
	return *o.LicenseAllPlugins
}

// GetLicenseAllPluginsOk returns a tuple with the LicenseAllPlugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetLicenseAllPluginsOk() (*bool, bool) {
	if o == nil || IsNil(o.LicenseAllPlugins) {
		return nil, false
	}
	return o.LicenseAllPlugins, true
}

// HasLicenseAllPlugins returns a boolean if a field has been set.
func (o *Payload) HasLicenseAllPlugins() bool {
	if o != nil && !IsNil(o.LicenseAllPlugins) {
		return true
	}

	return false
}

// SetLicenseAllPlugins gets a reference to the given bool and assigns it to the LicenseAllPlugins field.
func (o *Payload) SetLicenseAllPlugins(v bool) {
	o.LicenseAllPlugins = &v
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *Payload) GetCaseId() string {
	if o == nil || IsNil(o.CaseId) {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetCaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *Payload) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *Payload) SetCaseId(v string) {
	o.CaseId = &v
}

func (o Payload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WithAddons) {
		toSerialize["withAddons"] = o.WithAddons
	}
	if !IsNil(o.LicenseAllPlugins) {
		toSerialize["licenseAllPlugins"] = o.LicenseAllPlugins
	}
	if !IsNil(o.CaseId) {
		toSerialize["caseId"] = o.CaseId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Payload) UnmarshalJSON(data []byte) (err error) {
	varPayload := _Payload{}

	err = json.Unmarshal(data, &varPayload)

	if err != nil {
		return err
	}

	*o = Payload(varPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "withAddons")
		delete(additionalProperties, "licenseAllPlugins")
		delete(additionalProperties, "caseId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayload struct {
	value *Payload
	isSet bool
}

func (v NullablePayload) Get() *Payload {
	return v.value
}

func (v *NullablePayload) Set(val *Payload) {
	v.value = val
	v.isSet = true
}

func (v NullablePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayload(val *Payload) *NullablePayload {
	return &NullablePayload{value: val, isSet: true}
}

func (v NullablePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
