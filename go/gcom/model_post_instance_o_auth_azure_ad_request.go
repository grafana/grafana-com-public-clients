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

// checks if the PostInstanceOAuthAzureADRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstanceOAuthAzureADRequest{}

// PostInstanceOAuthAzureADRequest struct for PostInstanceOAuthAzureADRequest
type PostInstanceOAuthAzureADRequest struct {
	AllowedDomains       []string `json:"allowedDomains,omitempty"`
	AllowedGroups        []string `json:"allowedGroups,omitempty"`
	AuthUrl              string   `json:"authUrl"`
	ClientId             string   `json:"clientId"`
	ClientSecret         *string  `json:"clientSecret,omitempty"`
	Scopes               []string `json:"scopes,omitempty"`
	TokenUrl             string   `json:"tokenUrl"`
	AdditionalProperties map[string]interface{}
}

type _PostInstanceOAuthAzureADRequest PostInstanceOAuthAzureADRequest

// NewPostInstanceOAuthAzureADRequest instantiates a new PostInstanceOAuthAzureADRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstanceOAuthAzureADRequest(authUrl string, clientId string, tokenUrl string) *PostInstanceOAuthAzureADRequest {
	this := PostInstanceOAuthAzureADRequest{}
	this.AuthUrl = authUrl
	this.ClientId = clientId
	this.TokenUrl = tokenUrl
	return &this
}

// NewPostInstanceOAuthAzureADRequestWithDefaults instantiates a new PostInstanceOAuthAzureADRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstanceOAuthAzureADRequestWithDefaults() *PostInstanceOAuthAzureADRequest {
	this := PostInstanceOAuthAzureADRequest{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *PostInstanceOAuthAzureADRequest) GetAllowedDomains() []string {
	if o == nil || IsNil(o.AllowedDomains) {
		var ret []string
		return ret
	}
	return o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetAllowedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedDomains) {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *PostInstanceOAuthAzureADRequest) HasAllowedDomains() bool {
	if o != nil && !IsNil(o.AllowedDomains) {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *PostInstanceOAuthAzureADRequest) SetAllowedDomains(v []string) {
	o.AllowedDomains = v
}

// GetAllowedGroups returns the AllowedGroups field value if set, zero value otherwise.
func (o *PostInstanceOAuthAzureADRequest) GetAllowedGroups() []string {
	if o == nil || IsNil(o.AllowedGroups) {
		var ret []string
		return ret
	}
	return o.AllowedGroups
}

// GetAllowedGroupsOk returns a tuple with the AllowedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetAllowedGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedGroups) {
		return nil, false
	}
	return o.AllowedGroups, true
}

// HasAllowedGroups returns a boolean if a field has been set.
func (o *PostInstanceOAuthAzureADRequest) HasAllowedGroups() bool {
	if o != nil && !IsNil(o.AllowedGroups) {
		return true
	}

	return false
}

// SetAllowedGroups gets a reference to the given []string and assigns it to the AllowedGroups field.
func (o *PostInstanceOAuthAzureADRequest) SetAllowedGroups(v []string) {
	o.AllowedGroups = v
}

// GetAuthUrl returns the AuthUrl field value
func (o *PostInstanceOAuthAzureADRequest) GetAuthUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthUrl, true
}

// SetAuthUrl sets field value
func (o *PostInstanceOAuthAzureADRequest) SetAuthUrl(v string) {
	o.AuthUrl = v
}

// GetClientId returns the ClientId field value
func (o *PostInstanceOAuthAzureADRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PostInstanceOAuthAzureADRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PostInstanceOAuthAzureADRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PostInstanceOAuthAzureADRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PostInstanceOAuthAzureADRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PostInstanceOAuthAzureADRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PostInstanceOAuthAzureADRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *PostInstanceOAuthAzureADRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenUrl returns the TokenUrl field value
func (o *PostInstanceOAuthAzureADRequest) GetTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *PostInstanceOAuthAzureADRequest) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *PostInstanceOAuthAzureADRequest) SetTokenUrl(v string) {
	o.TokenUrl = v
}

func (o PostInstanceOAuthAzureADRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstanceOAuthAzureADRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedDomains) {
		toSerialize["allowedDomains"] = o.AllowedDomains
	}
	if !IsNil(o.AllowedGroups) {
		toSerialize["allowedGroups"] = o.AllowedGroups
	}
	toSerialize["authUrl"] = o.AuthUrl
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["tokenUrl"] = o.TokenUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostInstanceOAuthAzureADRequest) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varPostInstanceOAuthAzureADRequest := _PostInstanceOAuthAzureADRequest{}

	err = json.Unmarshal(data, &varPostInstanceOAuthAzureADRequest)

	if err != nil {
		return err
	}

	*o = PostInstanceOAuthAzureADRequest(varPostInstanceOAuthAzureADRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowedDomains")
		delete(additionalProperties, "allowedGroups")
		delete(additionalProperties, "authUrl")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "tokenUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostInstanceOAuthAzureADRequest struct {
	value *PostInstanceOAuthAzureADRequest
	isSet bool
}

func (v NullablePostInstanceOAuthAzureADRequest) Get() *PostInstanceOAuthAzureADRequest {
	return v.value
}

func (v *NullablePostInstanceOAuthAzureADRequest) Set(val *PostInstanceOAuthAzureADRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstanceOAuthAzureADRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstanceOAuthAzureADRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstanceOAuthAzureADRequest(val *PostInstanceOAuthAzureADRequest) *NullablePostInstanceOAuthAzureADRequest {
	return &NullablePostInstanceOAuthAzureADRequest{value: val, isSet: true}
}

func (v NullablePostInstanceOAuthAzureADRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstanceOAuthAzureADRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
