/*
GCOM API

 Grafana.com API (or GCOM). This documentation includes all endpoints of GCOM API including the staff ones.  Looking for GCOM API client packages? You can find them at [grafana-com-clients](https://github.com/grafana/grafana-com-clients) repository.  If you have any questions, please contact us at #grafana_com on Slack or open an issue at [Grafana-com repository](https://github.com/grafana/grafana-com/issues/new).  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: internal
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
)

// checks if the PostInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstanceRequest{}

// PostInstanceRequest struct for PostInstanceRequest
type PostInstanceRequest struct {
	Alerts               *bool              `json:"alerts,omitempty"`
	Description          *string            `json:"description,omitempty"`
	Graphite             *bool              `json:"graphite,omitempty"`
	HlInstanceId         *int32             `json:"hlInstanceId,omitempty"`
	K6OrgId              *int32             `json:"k6OrgId,omitempty"`
	Labels               *map[string]string `json:"labels,omitempty"`
	Logs                 *bool              `json:"logs,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	Plan                 *string            `json:"plan,omitempty"`
	Prometheus           *bool              `json:"prometheus,omitempty"`
	Slug                 *string            `json:"slug,omitempty"`
	Url                  *string            `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostInstanceRequest PostInstanceRequest

// NewPostInstanceRequest instantiates a new PostInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstanceRequest() *PostInstanceRequest {
	this := PostInstanceRequest{}
	return &this
}

// NewPostInstanceRequestWithDefaults instantiates a new PostInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstanceRequestWithDefaults() *PostInstanceRequest {
	this := PostInstanceRequest{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetAlerts() bool {
	if o == nil || IsNil(o.Alerts) {
		var ret bool
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetAlertsOk() (*bool, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given bool and assigns it to the Alerts field.
func (o *PostInstanceRequest) SetAlerts(v bool) {
	o.Alerts = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostInstanceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetGraphite() bool {
	if o == nil || IsNil(o.Graphite) {
		var ret bool
		return ret
	}
	return *o.Graphite
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetGraphiteOk() (*bool, bool) {
	if o == nil || IsNil(o.Graphite) {
		return nil, false
	}
	return o.Graphite, true
}

// HasGraphite returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasGraphite() bool {
	if o != nil && !IsNil(o.Graphite) {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given bool and assigns it to the Graphite field.
func (o *PostInstanceRequest) SetGraphite(v bool) {
	o.Graphite = &v
}

// GetHlInstanceId returns the HlInstanceId field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetHlInstanceId() int32 {
	if o == nil || IsNil(o.HlInstanceId) {
		var ret int32
		return ret
	}
	return *o.HlInstanceId
}

// GetHlInstanceIdOk returns a tuple with the HlInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetHlInstanceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HlInstanceId) {
		return nil, false
	}
	return o.HlInstanceId, true
}

// HasHlInstanceId returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasHlInstanceId() bool {
	if o != nil && !IsNil(o.HlInstanceId) {
		return true
	}

	return false
}

// SetHlInstanceId gets a reference to the given int32 and assigns it to the HlInstanceId field.
func (o *PostInstanceRequest) SetHlInstanceId(v int32) {
	o.HlInstanceId = &v
}

// GetK6OrgId returns the K6OrgId field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetK6OrgId() int32 {
	if o == nil || IsNil(o.K6OrgId) {
		var ret int32
		return ret
	}
	return *o.K6OrgId
}

// GetK6OrgIdOk returns a tuple with the K6OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetK6OrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.K6OrgId) {
		return nil, false
	}
	return o.K6OrgId, true
}

// HasK6OrgId returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasK6OrgId() bool {
	if o != nil && !IsNil(o.K6OrgId) {
		return true
	}

	return false
}

// SetK6OrgId gets a reference to the given int32 and assigns it to the K6OrgId field.
func (o *PostInstanceRequest) SetK6OrgId(v int32) {
	o.K6OrgId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *PostInstanceRequest) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetLogs() bool {
	if o == nil || IsNil(o.Logs) {
		var ret bool
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given bool and assigns it to the Logs field.
func (o *PostInstanceRequest) SetLogs(v bool) {
	o.Logs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostInstanceRequest) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *PostInstanceRequest) SetPlan(v string) {
	o.Plan = &v
}

// GetPrometheus returns the Prometheus field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetPrometheus() bool {
	if o == nil || IsNil(o.Prometheus) {
		var ret bool
		return ret
	}
	return *o.Prometheus
}

// GetPrometheusOk returns a tuple with the Prometheus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetPrometheusOk() (*bool, bool) {
	if o == nil || IsNil(o.Prometheus) {
		return nil, false
	}
	return o.Prometheus, true
}

// HasPrometheus returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasPrometheus() bool {
	if o != nil && !IsNil(o.Prometheus) {
		return true
	}

	return false
}

// SetPrometheus gets a reference to the given bool and assigns it to the Prometheus field.
func (o *PostInstanceRequest) SetPrometheus(v bool) {
	o.Prometheus = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PostInstanceRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PostInstanceRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstanceRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PostInstanceRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PostInstanceRequest) SetUrl(v string) {
	o.Url = &v
}

func (o PostInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Graphite) {
		toSerialize["graphite"] = o.Graphite
	}
	if !IsNil(o.HlInstanceId) {
		toSerialize["hlInstanceId"] = o.HlInstanceId
	}
	if !IsNil(o.K6OrgId) {
		toSerialize["k6OrgId"] = o.K6OrgId
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Prometheus) {
		toSerialize["prometheus"] = o.Prometheus
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	varPostInstanceRequest := _PostInstanceRequest{}

	err = json.Unmarshal(data, &varPostInstanceRequest)

	if err != nil {
		return err
	}

	*o = PostInstanceRequest(varPostInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alerts")
		delete(additionalProperties, "description")
		delete(additionalProperties, "graphite")
		delete(additionalProperties, "hlInstanceId")
		delete(additionalProperties, "k6OrgId")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "logs")
		delete(additionalProperties, "name")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "prometheus")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostInstanceRequest struct {
	value *PostInstanceRequest
	isSet bool
}

func (v NullablePostInstanceRequest) Get() *PostInstanceRequest {
	return v.value
}

func (v *NullablePostInstanceRequest) Set(val *PostInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstanceRequest(val *PostInstanceRequest) *NullablePostInstanceRequest {
	return &NullablePostInstanceRequest{value: val, isSet: true}
}

func (v NullablePostInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
