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

// checks if the PostInstancesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstancesRequest{}

// PostInstancesRequest struct for PostInstancesRequest
type PostInstancesRequest struct {
	AdminUserInstance               *bool    `json:"adminUserInstance,omitempty"`
	Alerts                          *bool    `json:"alerts,omitempty"`
	Cluster                         *string  `json:"cluster,omitempty"`
	CreateTemporaryLicenseIfMissing *bool    `json:"createTemporaryLicenseIfMissing,omitempty"`
	Description                     *string  `json:"description,omitempty"`
	Graphite                        *bool    `json:"graphite,omitempty"`
	HlInstanceId                    *int32   `json:"hlInstanceId,omitempty"`
	Hosted                          *bool    `json:"hosted,omitempty"`
	K6OrgId                         *int32   `json:"k6OrgId,omitempty"`
	Logs                            *bool    `json:"logs,omitempty"`
	Name                            string   `json:"name"`
	Org                             *string  `json:"org,omitempty"`
	Plan                            *string  `json:"plan,omitempty"`
	Plugins                         []string `json:"plugins,omitempty"`
	Prometheus                      *bool    `json:"prometheus,omitempty"`
	PublicInstance                  *bool    `json:"publicInstance,omitempty"`
	Region                          *string  `json:"region,omitempty"`
	Slug                            *string  `json:"slug,omitempty"`
	Url                             *string  `json:"url,omitempty"`
	UsernameOrEmail                 *string  `json:"usernameOrEmail,omitempty"`
	Version                         *string  `json:"version,omitempty"`
	WaitForReadiness                *bool    `json:"waitForReadiness,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _PostInstancesRequest PostInstancesRequest

// NewPostInstancesRequest instantiates a new PostInstancesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstancesRequest(name string) *PostInstancesRequest {
	this := PostInstancesRequest{}
	this.Name = name
	var waitForReadiness bool = true
	this.WaitForReadiness = &waitForReadiness
	return &this
}

// NewPostInstancesRequestWithDefaults instantiates a new PostInstancesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstancesRequestWithDefaults() *PostInstancesRequest {
	this := PostInstancesRequest{}
	var waitForReadiness bool = true
	this.WaitForReadiness = &waitForReadiness
	return &this
}

// GetAdminUserInstance returns the AdminUserInstance field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetAdminUserInstance() bool {
	if o == nil || IsNil(o.AdminUserInstance) {
		var ret bool
		return ret
	}
	return *o.AdminUserInstance
}

// GetAdminUserInstanceOk returns a tuple with the AdminUserInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetAdminUserInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminUserInstance) {
		return nil, false
	}
	return o.AdminUserInstance, true
}

// HasAdminUserInstance returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasAdminUserInstance() bool {
	if o != nil && !IsNil(o.AdminUserInstance) {
		return true
	}

	return false
}

// SetAdminUserInstance gets a reference to the given bool and assigns it to the AdminUserInstance field.
func (o *PostInstancesRequest) SetAdminUserInstance(v bool) {
	o.AdminUserInstance = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetAlerts() bool {
	if o == nil || IsNil(o.Alerts) {
		var ret bool
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetAlertsOk() (*bool, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given bool and assigns it to the Alerts field.
func (o *PostInstancesRequest) SetAlerts(v bool) {
	o.Alerts = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *PostInstancesRequest) SetCluster(v string) {
	o.Cluster = &v
}

// GetCreateTemporaryLicenseIfMissing returns the CreateTemporaryLicenseIfMissing field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetCreateTemporaryLicenseIfMissing() bool {
	if o == nil || IsNil(o.CreateTemporaryLicenseIfMissing) {
		var ret bool
		return ret
	}
	return *o.CreateTemporaryLicenseIfMissing
}

// GetCreateTemporaryLicenseIfMissingOk returns a tuple with the CreateTemporaryLicenseIfMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetCreateTemporaryLicenseIfMissingOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateTemporaryLicenseIfMissing) {
		return nil, false
	}
	return o.CreateTemporaryLicenseIfMissing, true
}

// HasCreateTemporaryLicenseIfMissing returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasCreateTemporaryLicenseIfMissing() bool {
	if o != nil && !IsNil(o.CreateTemporaryLicenseIfMissing) {
		return true
	}

	return false
}

// SetCreateTemporaryLicenseIfMissing gets a reference to the given bool and assigns it to the CreateTemporaryLicenseIfMissing field.
func (o *PostInstancesRequest) SetCreateTemporaryLicenseIfMissing(v bool) {
	o.CreateTemporaryLicenseIfMissing = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostInstancesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetGraphite() bool {
	if o == nil || IsNil(o.Graphite) {
		var ret bool
		return ret
	}
	return *o.Graphite
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetGraphiteOk() (*bool, bool) {
	if o == nil || IsNil(o.Graphite) {
		return nil, false
	}
	return o.Graphite, true
}

// HasGraphite returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasGraphite() bool {
	if o != nil && !IsNil(o.Graphite) {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given bool and assigns it to the Graphite field.
func (o *PostInstancesRequest) SetGraphite(v bool) {
	o.Graphite = &v
}

// GetHlInstanceId returns the HlInstanceId field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetHlInstanceId() int32 {
	if o == nil || IsNil(o.HlInstanceId) {
		var ret int32
		return ret
	}
	return *o.HlInstanceId
}

// GetHlInstanceIdOk returns a tuple with the HlInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetHlInstanceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HlInstanceId) {
		return nil, false
	}
	return o.HlInstanceId, true
}

// HasHlInstanceId returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasHlInstanceId() bool {
	if o != nil && !IsNil(o.HlInstanceId) {
		return true
	}

	return false
}

// SetHlInstanceId gets a reference to the given int32 and assigns it to the HlInstanceId field.
func (o *PostInstancesRequest) SetHlInstanceId(v int32) {
	o.HlInstanceId = &v
}

// GetHosted returns the Hosted field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetHosted() bool {
	if o == nil || IsNil(o.Hosted) {
		var ret bool
		return ret
	}
	return *o.Hosted
}

// GetHostedOk returns a tuple with the Hosted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetHostedOk() (*bool, bool) {
	if o == nil || IsNil(o.Hosted) {
		return nil, false
	}
	return o.Hosted, true
}

// HasHosted returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasHosted() bool {
	if o != nil && !IsNil(o.Hosted) {
		return true
	}

	return false
}

// SetHosted gets a reference to the given bool and assigns it to the Hosted field.
func (o *PostInstancesRequest) SetHosted(v bool) {
	o.Hosted = &v
}

// GetK6OrgId returns the K6OrgId field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetK6OrgId() int32 {
	if o == nil || IsNil(o.K6OrgId) {
		var ret int32
		return ret
	}
	return *o.K6OrgId
}

// GetK6OrgIdOk returns a tuple with the K6OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetK6OrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.K6OrgId) {
		return nil, false
	}
	return o.K6OrgId, true
}

// HasK6OrgId returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasK6OrgId() bool {
	if o != nil && !IsNil(o.K6OrgId) {
		return true
	}

	return false
}

// SetK6OrgId gets a reference to the given int32 and assigns it to the K6OrgId field.
func (o *PostInstancesRequest) SetK6OrgId(v int32) {
	o.K6OrgId = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetLogs() bool {
	if o == nil || IsNil(o.Logs) {
		var ret bool
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given bool and assigns it to the Logs field.
func (o *PostInstancesRequest) SetLogs(v bool) {
	o.Logs = &v
}

// GetName returns the Name field value
func (o *PostInstancesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostInstancesRequest) SetName(v string) {
	o.Name = v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetOrg() string {
	if o == nil || IsNil(o.Org) {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetOrgOk() (*string, bool) {
	if o == nil || IsNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasOrg() bool {
	if o != nil && !IsNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *PostInstancesRequest) SetOrg(v string) {
	o.Org = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *PostInstancesRequest) SetPlan(v string) {
	o.Plan = &v
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetPlugins() []string {
	if o == nil || IsNil(o.Plugins) {
		var ret []string
		return ret
	}
	return o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetPluginsOk() ([]string, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasPlugins() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given []string and assigns it to the Plugins field.
func (o *PostInstancesRequest) SetPlugins(v []string) {
	o.Plugins = v
}

// GetPrometheus returns the Prometheus field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetPrometheus() bool {
	if o == nil || IsNil(o.Prometheus) {
		var ret bool
		return ret
	}
	return *o.Prometheus
}

// GetPrometheusOk returns a tuple with the Prometheus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetPrometheusOk() (*bool, bool) {
	if o == nil || IsNil(o.Prometheus) {
		return nil, false
	}
	return o.Prometheus, true
}

// HasPrometheus returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasPrometheus() bool {
	if o != nil && !IsNil(o.Prometheus) {
		return true
	}

	return false
}

// SetPrometheus gets a reference to the given bool and assigns it to the Prometheus field.
func (o *PostInstancesRequest) SetPrometheus(v bool) {
	o.Prometheus = &v
}

// GetPublicInstance returns the PublicInstance field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetPublicInstance() bool {
	if o == nil || IsNil(o.PublicInstance) {
		var ret bool
		return ret
	}
	return *o.PublicInstance
}

// GetPublicInstanceOk returns a tuple with the PublicInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetPublicInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicInstance) {
		return nil, false
	}
	return o.PublicInstance, true
}

// HasPublicInstance returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasPublicInstance() bool {
	if o != nil && !IsNil(o.PublicInstance) {
		return true
	}

	return false
}

// SetPublicInstance gets a reference to the given bool and assigns it to the PublicInstance field.
func (o *PostInstancesRequest) SetPublicInstance(v bool) {
	o.PublicInstance = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PostInstancesRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PostInstancesRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PostInstancesRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsernameOrEmail returns the UsernameOrEmail field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetUsernameOrEmail() string {
	if o == nil || IsNil(o.UsernameOrEmail) {
		var ret string
		return ret
	}
	return *o.UsernameOrEmail
}

// GetUsernameOrEmailOk returns a tuple with the UsernameOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetUsernameOrEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UsernameOrEmail) {
		return nil, false
	}
	return o.UsernameOrEmail, true
}

// HasUsernameOrEmail returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasUsernameOrEmail() bool {
	if o != nil && !IsNil(o.UsernameOrEmail) {
		return true
	}

	return false
}

// SetUsernameOrEmail gets a reference to the given string and assigns it to the UsernameOrEmail field.
func (o *PostInstancesRequest) SetUsernameOrEmail(v string) {
	o.UsernameOrEmail = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PostInstancesRequest) SetVersion(v string) {
	o.Version = &v
}

// GetWaitForReadiness returns the WaitForReadiness field value if set, zero value otherwise.
func (o *PostInstancesRequest) GetWaitForReadiness() bool {
	if o == nil || IsNil(o.WaitForReadiness) {
		var ret bool
		return ret
	}
	return *o.WaitForReadiness
}

// GetWaitForReadinessOk returns a tuple with the WaitForReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancesRequest) GetWaitForReadinessOk() (*bool, bool) {
	if o == nil || IsNil(o.WaitForReadiness) {
		return nil, false
	}
	return o.WaitForReadiness, true
}

// HasWaitForReadiness returns a boolean if a field has been set.
func (o *PostInstancesRequest) HasWaitForReadiness() bool {
	if o != nil && !IsNil(o.WaitForReadiness) {
		return true
	}

	return false
}

// SetWaitForReadiness gets a reference to the given bool and assigns it to the WaitForReadiness field.
func (o *PostInstancesRequest) SetWaitForReadiness(v bool) {
	o.WaitForReadiness = &v
}

func (o PostInstancesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstancesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminUserInstance) {
		toSerialize["adminUserInstance"] = o.AdminUserInstance
	}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.CreateTemporaryLicenseIfMissing) {
		toSerialize["createTemporaryLicenseIfMissing"] = o.CreateTemporaryLicenseIfMissing
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
	if !IsNil(o.Hosted) {
		toSerialize["hosted"] = o.Hosted
	}
	if !IsNil(o.K6OrgId) {
		toSerialize["k6OrgId"] = o.K6OrgId
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	if !IsNil(o.Prometheus) {
		toSerialize["prometheus"] = o.Prometheus
	}
	if !IsNil(o.PublicInstance) {
		toSerialize["publicInstance"] = o.PublicInstance
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UsernameOrEmail) {
		toSerialize["usernameOrEmail"] = o.UsernameOrEmail
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.WaitForReadiness) {
		toSerialize["waitForReadiness"] = o.WaitForReadiness
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostInstancesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varPostInstancesRequest := _PostInstancesRequest{}

	err = json.Unmarshal(data, &varPostInstancesRequest)

	if err != nil {
		return err
	}

	*o = PostInstancesRequest(varPostInstancesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "adminUserInstance")
		delete(additionalProperties, "alerts")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "createTemporaryLicenseIfMissing")
		delete(additionalProperties, "description")
		delete(additionalProperties, "graphite")
		delete(additionalProperties, "hlInstanceId")
		delete(additionalProperties, "hosted")
		delete(additionalProperties, "k6OrgId")
		delete(additionalProperties, "logs")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "plugins")
		delete(additionalProperties, "prometheus")
		delete(additionalProperties, "publicInstance")
		delete(additionalProperties, "region")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "url")
		delete(additionalProperties, "usernameOrEmail")
		delete(additionalProperties, "version")
		delete(additionalProperties, "waitForReadiness")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostInstancesRequest struct {
	value *PostInstancesRequest
	isSet bool
}

func (v NullablePostInstancesRequest) Get() *PostInstancesRequest {
	return v.value
}

func (v *NullablePostInstancesRequest) Set(val *PostInstancesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstancesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstancesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstancesRequest(val *PostInstancesRequest) *NullablePostInstancesRequest {
	return &NullablePostInstancesRequest{value: val, isSet: true}
}

func (v NullablePostInstancesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstancesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
