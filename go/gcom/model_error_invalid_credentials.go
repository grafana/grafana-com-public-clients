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

// checks if the ErrorInvalidCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorInvalidCredentials{}

// ErrorInvalidCredentials struct for ErrorInvalidCredentials
type ErrorInvalidCredentials struct {
	Message              *string `json:"message,omitempty"`
	Code                 *string `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorInvalidCredentials ErrorInvalidCredentials

// NewErrorInvalidCredentials instantiates a new ErrorInvalidCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorInvalidCredentials() *ErrorInvalidCredentials {
	this := ErrorInvalidCredentials{}
	return &this
}

// NewErrorInvalidCredentialsWithDefaults instantiates a new ErrorInvalidCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorInvalidCredentialsWithDefaults() *ErrorInvalidCredentials {
	this := ErrorInvalidCredentials{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorInvalidCredentials) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInvalidCredentials) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorInvalidCredentials) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorInvalidCredentials) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorInvalidCredentials) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInvalidCredentials) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorInvalidCredentials) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorInvalidCredentials) SetCode(v string) {
	o.Code = &v
}

func (o ErrorInvalidCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorInvalidCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorInvalidCredentials) UnmarshalJSON(data []byte) (err error) {
	varErrorInvalidCredentials := _ErrorInvalidCredentials{}

	err = json.Unmarshal(data, &varErrorInvalidCredentials)

	if err != nil {
		return err
	}

	*o = ErrorInvalidCredentials(varErrorInvalidCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorInvalidCredentials struct {
	value *ErrorInvalidCredentials
	isSet bool
}

func (v NullableErrorInvalidCredentials) Get() *ErrorInvalidCredentials {
	return v.value
}

func (v *NullableErrorInvalidCredentials) Set(val *ErrorInvalidCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorInvalidCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorInvalidCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorInvalidCredentials(val *ErrorInvalidCredentials) *NullableErrorInvalidCredentials {
	return &NullableErrorInvalidCredentials{value: val, isSet: true}
}

func (v NullableErrorInvalidCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorInvalidCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
