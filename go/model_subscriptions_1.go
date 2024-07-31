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

// Subscriptions1 struct for Subscriptions1
type Subscriptions1 struct {
	Subscriptions1AnyOf  *Subscriptions1AnyOf
	Subscriptions1AnyOf1 *Subscriptions1AnyOf1
	Subscriptions1AnyOf2 *Subscriptions1AnyOf2
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Subscriptions1) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Subscriptions1AnyOf
	err = json.Unmarshal(data, &dst.Subscriptions1AnyOf)
	if err == nil {
		jsonSubscriptions1AnyOf, _ := json.Marshal(dst.Subscriptions1AnyOf)
		if string(jsonSubscriptions1AnyOf) == "{}" { // empty struct
			dst.Subscriptions1AnyOf = nil
		} else {
			return nil // data stored in dst.Subscriptions1AnyOf, return on the first match
		}
	} else {
		dst.Subscriptions1AnyOf = nil
	}

	// try to unmarshal JSON data into Subscriptions1AnyOf1
	err = json.Unmarshal(data, &dst.Subscriptions1AnyOf1)
	if err == nil {
		jsonSubscriptions1AnyOf1, _ := json.Marshal(dst.Subscriptions1AnyOf1)
		if string(jsonSubscriptions1AnyOf1) == "{}" { // empty struct
			dst.Subscriptions1AnyOf1 = nil
		} else {
			return nil // data stored in dst.Subscriptions1AnyOf1, return on the first match
		}
	} else {
		dst.Subscriptions1AnyOf1 = nil
	}

	// try to unmarshal JSON data into Subscriptions1AnyOf2
	err = json.Unmarshal(data, &dst.Subscriptions1AnyOf2)
	if err == nil {
		jsonSubscriptions1AnyOf2, _ := json.Marshal(dst.Subscriptions1AnyOf2)
		if string(jsonSubscriptions1AnyOf2) == "{}" { // empty struct
			dst.Subscriptions1AnyOf2 = nil
		} else {
			return nil // data stored in dst.Subscriptions1AnyOf2, return on the first match
		}
	} else {
		dst.Subscriptions1AnyOf2 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Subscriptions1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Subscriptions1) MarshalJSON() ([]byte, error) {
	if src.Subscriptions1AnyOf != nil {
		return json.Marshal(&src.Subscriptions1AnyOf)
	}

	if src.Subscriptions1AnyOf1 != nil {
		return json.Marshal(&src.Subscriptions1AnyOf1)
	}

	if src.Subscriptions1AnyOf2 != nil {
		return json.Marshal(&src.Subscriptions1AnyOf2)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubscriptions1 struct {
	value *Subscriptions1
	isSet bool
}

func (v NullableSubscriptions1) Get() *Subscriptions1 {
	return v.value
}

func (v *NullableSubscriptions1) Set(val *Subscriptions1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptions1(val *Subscriptions1) *NullableSubscriptions1 {
	return &NullableSubscriptions1{value: val, isSet: true}
}

func (v NullableSubscriptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}