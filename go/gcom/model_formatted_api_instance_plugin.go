/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"time"
)

// checks if the FormattedApiInstancePlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormattedApiInstancePlugin{}

// FormattedApiInstancePlugin struct for FormattedApiInstancePlugin
type FormattedApiInstancePlugin struct {
	Id                   float32       `json:"id"`
	InstanceId           float32       `json:"instanceId"`
	InstanceUrl          string        `json:"instanceUrl"`
	InstanceSlug         string        `json:"instanceSlug"`
	PluginId             float32       `json:"pluginId"`
	PluginSlug           string        `json:"pluginSlug"`
	PluginName           string        `json:"pluginName"`
	Version              string        `json:"version"`
	LatestVersion        string        `json:"latestVersion"`
	CreatedAt            string        `json:"createdAt"`
	UpdatedAt            *time.Time    `json:"updatedAt,omitempty"`
	Links                []LinksInner1 `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _FormattedApiInstancePlugin FormattedApiInstancePlugin

// NewFormattedApiInstancePlugin instantiates a new FormattedApiInstancePlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormattedApiInstancePlugin(id float32, instanceId float32, instanceUrl string, instanceSlug string, pluginId float32, pluginSlug string, pluginName string, version string, latestVersion string, createdAt string, links []LinksInner1) *FormattedApiInstancePlugin {
	this := FormattedApiInstancePlugin{}
	this.Id = id
	this.InstanceId = instanceId
	this.InstanceUrl = instanceUrl
	this.InstanceSlug = instanceSlug
	this.PluginId = pluginId
	this.PluginSlug = pluginSlug
	this.PluginName = pluginName
	this.Version = version
	this.LatestVersion = latestVersion
	this.CreatedAt = createdAt
	this.Links = links
	return &this
}

// NewFormattedApiInstancePluginWithDefaults instantiates a new FormattedApiInstancePlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormattedApiInstancePluginWithDefaults() *FormattedApiInstancePlugin {
	this := FormattedApiInstancePlugin{}
	return &this
}

// GetId returns the Id field value
func (o *FormattedApiInstancePlugin) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FormattedApiInstancePlugin) SetId(v float32) {
	o.Id = v
}

// GetInstanceId returns the InstanceId field value
func (o *FormattedApiInstancePlugin) GetInstanceId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetInstanceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *FormattedApiInstancePlugin) SetInstanceId(v float32) {
	o.InstanceId = v
}

// GetInstanceUrl returns the InstanceUrl field value
func (o *FormattedApiInstancePlugin) GetInstanceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetInstanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceUrl, true
}

// SetInstanceUrl sets field value
func (o *FormattedApiInstancePlugin) SetInstanceUrl(v string) {
	o.InstanceUrl = v
}

// GetInstanceSlug returns the InstanceSlug field value
func (o *FormattedApiInstancePlugin) GetInstanceSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceSlug
}

// GetInstanceSlugOk returns a tuple with the InstanceSlug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetInstanceSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceSlug, true
}

// SetInstanceSlug sets field value
func (o *FormattedApiInstancePlugin) SetInstanceSlug(v string) {
	o.InstanceSlug = v
}

// GetPluginId returns the PluginId field value
func (o *FormattedApiInstancePlugin) GetPluginId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetPluginIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginId, true
}

// SetPluginId sets field value
func (o *FormattedApiInstancePlugin) SetPluginId(v float32) {
	o.PluginId = v
}

// GetPluginSlug returns the PluginSlug field value
func (o *FormattedApiInstancePlugin) GetPluginSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginSlug
}

// GetPluginSlugOk returns a tuple with the PluginSlug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetPluginSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginSlug, true
}

// SetPluginSlug sets field value
func (o *FormattedApiInstancePlugin) SetPluginSlug(v string) {
	o.PluginSlug = v
}

// GetPluginName returns the PluginName field value
func (o *FormattedApiInstancePlugin) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *FormattedApiInstancePlugin) SetPluginName(v string) {
	o.PluginName = v
}

// GetVersion returns the Version field value
func (o *FormattedApiInstancePlugin) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FormattedApiInstancePlugin) SetVersion(v string) {
	o.Version = v
}

// GetLatestVersion returns the LatestVersion field value
func (o *FormattedApiInstancePlugin) GetLatestVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetLatestVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestVersion, true
}

// SetLatestVersion sets field value
func (o *FormattedApiInstancePlugin) SetLatestVersion(v string) {
	o.LatestVersion = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FormattedApiInstancePlugin) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FormattedApiInstancePlugin) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FormattedApiInstancePlugin) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FormattedApiInstancePlugin) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FormattedApiInstancePlugin) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value
func (o *FormattedApiInstancePlugin) GetLinks() []LinksInner1 {
	if o == nil {
		var ret []LinksInner1
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FormattedApiInstancePlugin) GetLinksOk() ([]LinksInner1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *FormattedApiInstancePlugin) SetLinks(v []LinksInner1) {
	o.Links = v
}

func (o FormattedApiInstancePlugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormattedApiInstancePlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["instanceUrl"] = o.InstanceUrl
	toSerialize["instanceSlug"] = o.InstanceSlug
	toSerialize["pluginId"] = o.PluginId
	toSerialize["pluginSlug"] = o.PluginSlug
	toSerialize["pluginName"] = o.PluginName
	toSerialize["version"] = o.Version
	toSerialize["latestVersion"] = o.LatestVersion
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormattedApiInstancePlugin) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varFormattedApiInstancePlugin := _FormattedApiInstancePlugin{}

	err = json.Unmarshal(data, &varFormattedApiInstancePlugin)

	if err != nil {
		return err
	}

	*o = FormattedApiInstancePlugin(varFormattedApiInstancePlugin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "instanceUrl")
		delete(additionalProperties, "instanceSlug")
		delete(additionalProperties, "pluginId")
		delete(additionalProperties, "pluginSlug")
		delete(additionalProperties, "pluginName")
		delete(additionalProperties, "version")
		delete(additionalProperties, "latestVersion")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormattedApiInstancePlugin struct {
	value *FormattedApiInstancePlugin
	isSet bool
}

func (v NullableFormattedApiInstancePlugin) Get() *FormattedApiInstancePlugin {
	return v.value
}

func (v *NullableFormattedApiInstancePlugin) Set(val *FormattedApiInstancePlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableFormattedApiInstancePlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableFormattedApiInstancePlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormattedApiInstancePlugin(val *FormattedApiInstancePlugin) *NullableFormattedApiInstancePlugin {
	return &NullableFormattedApiInstancePlugin{value: val, isSet: true}
}

func (v NullableFormattedApiInstancePlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormattedApiInstancePlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
