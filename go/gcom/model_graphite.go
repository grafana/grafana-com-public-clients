/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"fmt"
)

// checks if the Graphite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Graphite{}

// Graphite struct for Graphite
type Graphite struct {
	PrivateDNS           string `json:"privateDNS"`
	ServiceName          string `json:"serviceName"`
	AdditionalProperties map[string]interface{}
}

type _Graphite Graphite

// NewGraphite instantiates a new Graphite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphite(privateDNS string, serviceName string) *Graphite {
	this := Graphite{}
	this.PrivateDNS = privateDNS
	this.ServiceName = serviceName
	return &this
}

// NewGraphiteWithDefaults instantiates a new Graphite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphiteWithDefaults() *Graphite {
	this := Graphite{}
	return &this
}

// GetPrivateDNS returns the PrivateDNS field value
func (o *Graphite) GetPrivateDNS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateDNS
}

// GetPrivateDNSOk returns a tuple with the PrivateDNS field value
// and a boolean to check if the value has been set.
func (o *Graphite) GetPrivateDNSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateDNS, true
}

// SetPrivateDNS sets field value
func (o *Graphite) SetPrivateDNS(v string) {
	o.PrivateDNS = v
}

// GetServiceName returns the ServiceName field value
func (o *Graphite) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *Graphite) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *Graphite) SetServiceName(v string) {
	o.ServiceName = v
}

func (o Graphite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Graphite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privateDNS"] = o.PrivateDNS
	toSerialize["serviceName"] = o.ServiceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Graphite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privateDNS",
		"serviceName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGraphite := _Graphite{}

	err = json.Unmarshal(data, &varGraphite)

	if err != nil {
		return err
	}

	*o = Graphite(varGraphite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "privateDNS")
		delete(additionalProperties, "serviceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphite struct {
	value *Graphite
	isSet bool
}

func (v NullableGraphite) Get() *Graphite {
	return v.value
}

func (v *NullableGraphite) Set(val *Graphite) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphite) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphite(val *Graphite) *NullableGraphite {
	return &NullableGraphite{value: val, isSet: true}
}

func (v NullableGraphite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
