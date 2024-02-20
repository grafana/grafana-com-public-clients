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

// checks if the PostInstanceOAuthGoogleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstanceOAuthGoogleRequest{}

// PostInstanceOAuthGoogleRequest struct for PostInstanceOAuthGoogleRequest
type PostInstanceOAuthGoogleRequest struct {
	AllowedDomains       []string `json:"allowedDomains"`
	ClientId             string   `json:"clientId"`
	ClientSecret         *string  `json:"clientSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostInstanceOAuthGoogleRequest PostInstanceOAuthGoogleRequest

// NewPostInstanceOAuthGoogleRequest instantiates a new PostInstanceOAuthGoogleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstanceOAuthGoogleRequest(allowedDomains []string, clientId string) *PostInstanceOAuthGoogleRequest {
	this := PostInstanceOAuthGoogleRequest{}
	this.AllowedDomains = allowedDomains
	this.ClientId = clientId
	return &this
}

// NewPostInstanceOAuthGoogleRequestWithDefaults instantiates a new PostInstanceOAuthGoogleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstanceOAuthGoogleRequestWithDefaults() *PostInstanceOAuthGoogleRequest {
	this := PostInstanceOAuthGoogleRequest{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value
func (o *PostInstanceOAuthGoogleRequest) GetAllowedDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGoogleRequest) GetAllowedDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// SetAllowedDomains sets field value
func (o *PostInstanceOAuthGoogleRequest) SetAllowedDomains(v []string) {
	o.AllowedDomains = v
}

// GetClientId returns the ClientId field value
func (o *PostInstanceOAuthGoogleRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGoogleRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PostInstanceOAuthGoogleRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PostInstanceOAuthGoogleRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGoogleRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PostInstanceOAuthGoogleRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PostInstanceOAuthGoogleRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o PostInstanceOAuthGoogleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstanceOAuthGoogleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowedDomains"] = o.AllowedDomains
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostInstanceOAuthGoogleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowedDomains",
		"clientId",
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

	varPostInstanceOAuthGoogleRequest := _PostInstanceOAuthGoogleRequest{}

	err = json.Unmarshal(data, &varPostInstanceOAuthGoogleRequest)

	if err != nil {
		return err
	}

	*o = PostInstanceOAuthGoogleRequest(varPostInstanceOAuthGoogleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowedDomains")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostInstanceOAuthGoogleRequest struct {
	value *PostInstanceOAuthGoogleRequest
	isSet bool
}

func (v NullablePostInstanceOAuthGoogleRequest) Get() *PostInstanceOAuthGoogleRequest {
	return v.value
}

func (v *NullablePostInstanceOAuthGoogleRequest) Set(val *PostInstanceOAuthGoogleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstanceOAuthGoogleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstanceOAuthGoogleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstanceOAuthGoogleRequest(val *PostInstanceOAuthGoogleRequest) *NullablePostInstanceOAuthGoogleRequest {
	return &NullablePostInstanceOAuthGoogleRequest{value: val, isSet: true}
}

func (v NullablePostInstanceOAuthGoogleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstanceOAuthGoogleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
