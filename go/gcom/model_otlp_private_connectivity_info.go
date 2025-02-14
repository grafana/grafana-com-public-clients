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

// OtlpPrivateConnectivityInfo struct for OtlpPrivateConnectivityInfo
type OtlpPrivateConnectivityInfo struct {
	PdcPrivateConnectivityInfoAnyOf  *PdcPrivateConnectivityInfoAnyOf
	PdcPrivateConnectivityInfoAnyOf1 *PdcPrivateConnectivityInfoAnyOf1
	PdcPrivateConnectivityInfoAnyOf2 *PdcPrivateConnectivityInfoAnyOf2
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OtlpPrivateConnectivityInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PdcPrivateConnectivityInfoAnyOf
	err = json.Unmarshal(data, &dst.PdcPrivateConnectivityInfoAnyOf)
	if err == nil {
		jsonPdcPrivateConnectivityInfoAnyOf, _ := json.Marshal(dst.PdcPrivateConnectivityInfoAnyOf)
		if string(jsonPdcPrivateConnectivityInfoAnyOf) == "{}" { // empty struct
			dst.PdcPrivateConnectivityInfoAnyOf = nil
		} else {
			return nil // data stored in dst.PdcPrivateConnectivityInfoAnyOf, return on the first match
		}
	} else {
		dst.PdcPrivateConnectivityInfoAnyOf = nil
	}

	// try to unmarshal JSON data into PdcPrivateConnectivityInfoAnyOf1
	err = json.Unmarshal(data, &dst.PdcPrivateConnectivityInfoAnyOf1)
	if err == nil {
		jsonPdcPrivateConnectivityInfoAnyOf1, _ := json.Marshal(dst.PdcPrivateConnectivityInfoAnyOf1)
		if string(jsonPdcPrivateConnectivityInfoAnyOf1) == "{}" { // empty struct
			dst.PdcPrivateConnectivityInfoAnyOf1 = nil
		} else {
			return nil // data stored in dst.PdcPrivateConnectivityInfoAnyOf1, return on the first match
		}
	} else {
		dst.PdcPrivateConnectivityInfoAnyOf1 = nil
	}

	// try to unmarshal JSON data into PdcPrivateConnectivityInfoAnyOf2
	err = json.Unmarshal(data, &dst.PdcPrivateConnectivityInfoAnyOf2)
	if err == nil {
		jsonPdcPrivateConnectivityInfoAnyOf2, _ := json.Marshal(dst.PdcPrivateConnectivityInfoAnyOf2)
		if string(jsonPdcPrivateConnectivityInfoAnyOf2) == "{}" { // empty struct
			dst.PdcPrivateConnectivityInfoAnyOf2 = nil
		} else {
			return nil // data stored in dst.PdcPrivateConnectivityInfoAnyOf2, return on the first match
		}
	} else {
		dst.PdcPrivateConnectivityInfoAnyOf2 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(OtlpPrivateConnectivityInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OtlpPrivateConnectivityInfo) MarshalJSON() ([]byte, error) {
	if src.PdcPrivateConnectivityInfoAnyOf != nil {
		return json.Marshal(&src.PdcPrivateConnectivityInfoAnyOf)
	}

	if src.PdcPrivateConnectivityInfoAnyOf1 != nil {
		return json.Marshal(&src.PdcPrivateConnectivityInfoAnyOf1)
	}

	if src.PdcPrivateConnectivityInfoAnyOf2 != nil {
		return json.Marshal(&src.PdcPrivateConnectivityInfoAnyOf2)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOtlpPrivateConnectivityInfo struct {
	value *OtlpPrivateConnectivityInfo
	isSet bool
}

func (v NullableOtlpPrivateConnectivityInfo) Get() *OtlpPrivateConnectivityInfo {
	return v.value
}

func (v *NullableOtlpPrivateConnectivityInfo) Set(val *OtlpPrivateConnectivityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOtlpPrivateConnectivityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOtlpPrivateConnectivityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtlpPrivateConnectivityInfo(val *OtlpPrivateConnectivityInfo) *NullableOtlpPrivateConnectivityInfo {
	return &NullableOtlpPrivateConnectivityInfo{value: val, isSet: true}
}

func (v NullableOtlpPrivateConnectivityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtlpPrivateConnectivityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
