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

// checks if the SubscriptionsAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsAnyOf{}

// SubscriptionsAnyOf struct for SubscriptionsAnyOf
type SubscriptionsAnyOf struct {
	Current              Current     `json:"current"`
	NextProduct          interface{} `json:"nextProduct"`
	Next                 interface{} `json:"next"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionsAnyOf SubscriptionsAnyOf

// NewSubscriptionsAnyOf instantiates a new SubscriptionsAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsAnyOf(current Current, nextProduct interface{}, next interface{}) *SubscriptionsAnyOf {
	this := SubscriptionsAnyOf{}
	this.Current = current
	this.NextProduct = nextProduct
	this.Next = next
	return &this
}

// NewSubscriptionsAnyOfWithDefaults instantiates a new SubscriptionsAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsAnyOfWithDefaults() *SubscriptionsAnyOf {
	this := SubscriptionsAnyOf{}
	return &this
}

// GetCurrent returns the Current field value
func (o *SubscriptionsAnyOf) GetCurrent() Current {
	if o == nil {
		var ret Current
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf) GetCurrentOk() (*Current, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *SubscriptionsAnyOf) SetCurrent(v Current) {
	o.Current = v
}

// GetNextProduct returns the NextProduct field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SubscriptionsAnyOf) GetNextProduct() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.NextProduct
}

// GetNextProductOk returns a tuple with the NextProduct field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionsAnyOf) GetNextProductOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NextProduct) {
		return nil, false
	}
	return &o.NextProduct, true
}

// SetNextProduct sets field value
func (o *SubscriptionsAnyOf) SetNextProduct(v interface{}) {
	o.NextProduct = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SubscriptionsAnyOf) GetNext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionsAnyOf) GetNextOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *SubscriptionsAnyOf) SetNext(v interface{}) {
	o.Next = v
}

func (o SubscriptionsAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	if o.NextProduct != nil {
		toSerialize["nextProduct"] = o.NextProduct
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionsAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current",
		"nextProduct",
		"next",
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

	varSubscriptionsAnyOf := _SubscriptionsAnyOf{}

	err = json.Unmarshal(data, &varSubscriptionsAnyOf)

	if err != nil {
		return err
	}

	*o = SubscriptionsAnyOf(varSubscriptionsAnyOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current")
		delete(additionalProperties, "nextProduct")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionsAnyOf struct {
	value *SubscriptionsAnyOf
	isSet bool
}

func (v NullableSubscriptionsAnyOf) Get() *SubscriptionsAnyOf {
	return v.value
}

func (v *NullableSubscriptionsAnyOf) Set(val *SubscriptionsAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsAnyOf(val *SubscriptionsAnyOf) *NullableSubscriptionsAnyOf {
	return &NullableSubscriptionsAnyOf{value: val, isSet: true}
}

func (v NullableSubscriptionsAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
