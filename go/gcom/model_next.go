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

// checks if the Next type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Next{}

// Next struct for Next
type Next struct {
	Product              string                 `json:"product"`
	Payload              map[string]interface{} `json:"payload"`
	Plan                 NullableString         `json:"plan"`
	PublicName           NullableString         `json:"publicName"`
	AdditionalProperties map[string]interface{}
}

type _Next Next

// NewNext instantiates a new Next object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNext(product string, payload map[string]interface{}, plan NullableString, publicName NullableString) *Next {
	this := Next{}
	this.Product = product
	this.Payload = payload
	this.Plan = plan
	this.PublicName = publicName
	return &this
}

// NewNextWithDefaults instantiates a new Next object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextWithDefaults() *Next {
	this := Next{}
	return &this
}

// GetProduct returns the Product field value
func (o *Next) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *Next) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *Next) SetProduct(v string) {
	o.Product = v
}

// GetPayload returns the Payload field value
func (o *Next) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Next) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *Next) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Next) GetPlan() string {
	if o == nil || o.Plan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Next) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *Next) SetPlan(v string) {
	o.Plan.Set(&v)
}

// GetPublicName returns the PublicName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Next) GetPublicName() string {
	if o == nil || o.PublicName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicName.Get()
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Next) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicName.Get(), o.PublicName.IsSet()
}

// SetPublicName sets field value
func (o *Next) SetPublicName(v string) {
	o.PublicName.Set(&v)
}

func (o Next) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Next) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["payload"] = o.Payload
	toSerialize["plan"] = o.Plan.Get()
	toSerialize["publicName"] = o.PublicName.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Next) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"payload",
		"plan",
		"publicName",
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

	varNext := _Next{}

	err = json.Unmarshal(data, &varNext)

	if err != nil {
		return err
	}

	*o = Next(varNext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "publicName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNext struct {
	value *Next
	isSet bool
}

func (v NullableNext) Get() *Next {
	return v.value
}

func (v *NullableNext) Set(val *Next) {
	v.value = val
	v.isSet = true
}

func (v NullableNext) IsSet() bool {
	return v.isSet
}

func (v *NullableNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNext(val *Next) *NullableNext {
	return &NullableNext{value: val, isSet: true}
}

func (v NullableNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
