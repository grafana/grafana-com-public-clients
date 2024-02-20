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

// Subscriptions struct for Subscriptions
type Subscriptions struct {
	SubscriptionsAnyOf  *SubscriptionsAnyOf
	SubscriptionsAnyOf1 *SubscriptionsAnyOf1
	SubscriptionsAnyOf2 *SubscriptionsAnyOf2
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Subscriptions) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubscriptionsAnyOf
	err = json.Unmarshal(data, &dst.SubscriptionsAnyOf)
	if err == nil {
		jsonSubscriptionsAnyOf, _ := json.Marshal(dst.SubscriptionsAnyOf)
		if string(jsonSubscriptionsAnyOf) == "{}" { // empty struct
			dst.SubscriptionsAnyOf = nil
		} else {
			return nil // data stored in dst.SubscriptionsAnyOf, return on the first match
		}
	} else {
		dst.SubscriptionsAnyOf = nil
	}

	// try to unmarshal JSON data into SubscriptionsAnyOf1
	err = json.Unmarshal(data, &dst.SubscriptionsAnyOf1)
	if err == nil {
		jsonSubscriptionsAnyOf1, _ := json.Marshal(dst.SubscriptionsAnyOf1)
		if string(jsonSubscriptionsAnyOf1) == "{}" { // empty struct
			dst.SubscriptionsAnyOf1 = nil
		} else {
			return nil // data stored in dst.SubscriptionsAnyOf1, return on the first match
		}
	} else {
		dst.SubscriptionsAnyOf1 = nil
	}

	// try to unmarshal JSON data into SubscriptionsAnyOf2
	err = json.Unmarshal(data, &dst.SubscriptionsAnyOf2)
	if err == nil {
		jsonSubscriptionsAnyOf2, _ := json.Marshal(dst.SubscriptionsAnyOf2)
		if string(jsonSubscriptionsAnyOf2) == "{}" { // empty struct
			dst.SubscriptionsAnyOf2 = nil
		} else {
			return nil // data stored in dst.SubscriptionsAnyOf2, return on the first match
		}
	} else {
		dst.SubscriptionsAnyOf2 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Subscriptions)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Subscriptions) MarshalJSON() ([]byte, error) {
	if src.SubscriptionsAnyOf != nil {
		return json.Marshal(&src.SubscriptionsAnyOf)
	}

	if src.SubscriptionsAnyOf1 != nil {
		return json.Marshal(&src.SubscriptionsAnyOf1)
	}

	if src.SubscriptionsAnyOf2 != nil {
		return json.Marshal(&src.SubscriptionsAnyOf2)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubscriptions struct {
	value *Subscriptions
	isSet bool
}

func (v NullableSubscriptions) Get() *Subscriptions {
	return v.value
}

func (v *NullableSubscriptions) Set(val *Subscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptions(val *Subscriptions) *NullableSubscriptions {
	return &NullableSubscriptions{value: val, isSet: true}
}

func (v NullableSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
