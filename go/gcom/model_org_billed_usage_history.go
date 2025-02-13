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

// checks if the OrgBilledUsageHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgBilledUsageHistory{}

// OrgBilledUsageHistory struct for OrgBilledUsageHistory
type OrgBilledUsageHistory struct {
	Items                []ItemsInner4 `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _OrgBilledUsageHistory OrgBilledUsageHistory

// NewOrgBilledUsageHistory instantiates a new OrgBilledUsageHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgBilledUsageHistory(items []ItemsInner4) *OrgBilledUsageHistory {
	this := OrgBilledUsageHistory{}
	this.Items = items
	return &this
}

// NewOrgBilledUsageHistoryWithDefaults instantiates a new OrgBilledUsageHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgBilledUsageHistoryWithDefaults() *OrgBilledUsageHistory {
	this := OrgBilledUsageHistory{}
	return &this
}

// GetItems returns the Items field value
func (o *OrgBilledUsageHistory) GetItems() []ItemsInner4 {
	if o == nil {
		var ret []ItemsInner4
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrgBilledUsageHistory) GetItemsOk() ([]ItemsInner4, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrgBilledUsageHistory) SetItems(v []ItemsInner4) {
	o.Items = v
}

func (o OrgBilledUsageHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgBilledUsageHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgBilledUsageHistory) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varOrgBilledUsageHistory := _OrgBilledUsageHistory{}

	err = json.Unmarshal(data, &varOrgBilledUsageHistory)

	if err != nil {
		return err
	}

	*o = OrgBilledUsageHistory(varOrgBilledUsageHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgBilledUsageHistory struct {
	value *OrgBilledUsageHistory
	isSet bool
}

func (v NullableOrgBilledUsageHistory) Get() *OrgBilledUsageHistory {
	return v.value
}

func (v *NullableOrgBilledUsageHistory) Set(val *OrgBilledUsageHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgBilledUsageHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgBilledUsageHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgBilledUsageHistory(val *OrgBilledUsageHistory) *NullableOrgBilledUsageHistory {
	return &NullableOrgBilledUsageHistory{value: val, isSet: true}
}

func (v NullableOrgBilledUsageHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgBilledUsageHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
