/*
GCOM API

 Grafana.com API (or GCOM). This documentation includes all endpoints of GCOM API including the staff ones.  Looking for GCOM API client packages? You can find them at [grafana-com-clients](https://github.com/grafana/grafana-com-clients) repository.  If you have any questions, please contact us at #grafana_com on Slack or open an issue at [Grafana-com repository](https://github.com/grafana/grafana-com/issues/new).  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: internal
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"fmt"
)

// checks if the SubscriptionsAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsAnyOf1{}

// SubscriptionsAnyOf1 struct for SubscriptionsAnyOf1
type SubscriptionsAnyOf1 struct {
	Current              Current1 `json:"current"`
	NextProduct          string   `json:"nextProduct"`
	Next                 Next     `json:"next"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionsAnyOf1 SubscriptionsAnyOf1

// NewSubscriptionsAnyOf1 instantiates a new SubscriptionsAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsAnyOf1(current Current1, nextProduct string, next Next) *SubscriptionsAnyOf1 {
	this := SubscriptionsAnyOf1{}
	this.Current = current
	this.NextProduct = nextProduct
	this.Next = next
	return &this
}

// NewSubscriptionsAnyOf1WithDefaults instantiates a new SubscriptionsAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsAnyOf1WithDefaults() *SubscriptionsAnyOf1 {
	this := SubscriptionsAnyOf1{}
	return &this
}

// GetCurrent returns the Current field value
func (o *SubscriptionsAnyOf1) GetCurrent() Current1 {
	if o == nil {
		var ret Current1
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetCurrentOk() (*Current1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *SubscriptionsAnyOf1) SetCurrent(v Current1) {
	o.Current = v
}

// GetNextProduct returns the NextProduct field value
func (o *SubscriptionsAnyOf1) GetNextProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextProduct
}

// GetNextProductOk returns a tuple with the NextProduct field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetNextProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextProduct, true
}

// SetNextProduct sets field value
func (o *SubscriptionsAnyOf1) SetNextProduct(v string) {
	o.NextProduct = v
}

// GetNext returns the Next field value
func (o *SubscriptionsAnyOf1) GetNext() Next {
	if o == nil {
		var ret Next
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetNextOk() (*Next, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *SubscriptionsAnyOf1) SetNext(v Next) {
	o.Next = v
}

func (o SubscriptionsAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	toSerialize["nextProduct"] = o.NextProduct
	toSerialize["next"] = o.Next

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionsAnyOf1) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionsAnyOf1 := _SubscriptionsAnyOf1{}

	err = json.Unmarshal(data, &varSubscriptionsAnyOf1)

	if err != nil {
		return err
	}

	*o = SubscriptionsAnyOf1(varSubscriptionsAnyOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current")
		delete(additionalProperties, "nextProduct")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionsAnyOf1 struct {
	value *SubscriptionsAnyOf1
	isSet bool
}

func (v NullableSubscriptionsAnyOf1) Get() *SubscriptionsAnyOf1 {
	return v.value
}

func (v *NullableSubscriptionsAnyOf1) Set(val *SubscriptionsAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsAnyOf1(val *SubscriptionsAnyOf1) *NullableSubscriptionsAnyOf1 {
	return &NullableSubscriptionsAnyOf1{value: val, isSet: true}
}

func (v NullableSubscriptionsAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
