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

// checks if the PostInstanceOAuthGithubRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstanceOAuthGithubRequest{}

// PostInstanceOAuthGithubRequest struct for PostInstanceOAuthGithubRequest
type PostInstanceOAuthGithubRequest struct {
	AllowedOrganizations []string `json:"allowedOrganizations,omitempty"`
	ClientId             string   `json:"clientId"`
	ClientSecret         *string  `json:"clientSecret,omitempty"`
	TeamIds              []string `json:"teamIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostInstanceOAuthGithubRequest PostInstanceOAuthGithubRequest

// NewPostInstanceOAuthGithubRequest instantiates a new PostInstanceOAuthGithubRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstanceOAuthGithubRequest(clientId string) *PostInstanceOAuthGithubRequest {
	this := PostInstanceOAuthGithubRequest{}
	this.ClientId = clientId
	return &this
}

// NewPostInstanceOAuthGithubRequestWithDefaults instantiates a new PostInstanceOAuthGithubRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstanceOAuthGithubRequestWithDefaults() *PostInstanceOAuthGithubRequest {
	this := PostInstanceOAuthGithubRequest{}
	return &this
}

// GetAllowedOrganizations returns the AllowedOrganizations field value if set, zero value otherwise.
func (o *PostInstanceOAuthGithubRequest) GetAllowedOrganizations() []string {
	if o == nil || IsNil(o.AllowedOrganizations) {
		var ret []string
		return ret
	}
	return o.AllowedOrganizations
}

// GetAllowedOrganizationsOk returns a tuple with the AllowedOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGithubRequest) GetAllowedOrganizationsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedOrganizations) {
		return nil, false
	}
	return o.AllowedOrganizations, true
}

// HasAllowedOrganizations returns a boolean if a field has been set.
func (o *PostInstanceOAuthGithubRequest) HasAllowedOrganizations() bool {
	if o != nil && !IsNil(o.AllowedOrganizations) {
		return true
	}

	return false
}

// SetAllowedOrganizations gets a reference to the given []string and assigns it to the AllowedOrganizations field.
func (o *PostInstanceOAuthGithubRequest) SetAllowedOrganizations(v []string) {
	o.AllowedOrganizations = v
}

// GetClientId returns the ClientId field value
func (o *PostInstanceOAuthGithubRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGithubRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PostInstanceOAuthGithubRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PostInstanceOAuthGithubRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGithubRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PostInstanceOAuthGithubRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PostInstanceOAuthGithubRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *PostInstanceOAuthGithubRequest) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthGithubRequest) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *PostInstanceOAuthGithubRequest) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *PostInstanceOAuthGithubRequest) SetTeamIds(v []string) {
	o.TeamIds = v
}

func (o PostInstanceOAuthGithubRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstanceOAuthGithubRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedOrganizations) {
		toSerialize["allowedOrganizations"] = o.AllowedOrganizations
	}
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.TeamIds) {
		toSerialize["teamIds"] = o.TeamIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostInstanceOAuthGithubRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPostInstanceOAuthGithubRequest := _PostInstanceOAuthGithubRequest{}

	err = json.Unmarshal(data, &varPostInstanceOAuthGithubRequest)

	if err != nil {
		return err
	}

	*o = PostInstanceOAuthGithubRequest(varPostInstanceOAuthGithubRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowedOrganizations")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "teamIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostInstanceOAuthGithubRequest struct {
	value *PostInstanceOAuthGithubRequest
	isSet bool
}

func (v NullablePostInstanceOAuthGithubRequest) Get() *PostInstanceOAuthGithubRequest {
	return v.value
}

func (v *NullablePostInstanceOAuthGithubRequest) Set(val *PostInstanceOAuthGithubRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstanceOAuthGithubRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstanceOAuthGithubRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstanceOAuthGithubRequest(val *PostInstanceOAuthGithubRequest) *NullablePostInstanceOAuthGithubRequest {
	return &NullablePostInstanceOAuthGithubRequest{value: val, isSet: true}
}

func (v NullablePostInstanceOAuthGithubRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstanceOAuthGithubRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
