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

// checks if the ItemsInner1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemsInner1{}

// ItemsInner1 struct for ItemsInner1
type ItemsInner1 struct {
	Id                       float32                                    `json:"id"`
	OrgId                    float32                                    `json:"orgId"`
	UserId                   float32                                    `json:"userId"`
	Status                   float32                                    `json:"status"`
	CreatedAt                string                                     `json:"createdAt"`
	UpdatedAt                NullableString                             `json:"updatedAt"`
	DefaultOrg               float32                                    `json:"defaultOrg"`
	Role                     string                                     `json:"role"`
	Privacy                  float32                                    `json:"privacy"`
	Billing                  float32                                    `json:"billing"`
	CreatedBy                string                                     `json:"createdBy"`
	UpdatedBy                string                                     `json:"updatedBy"`
	OrgName                  string                                     `json:"orgName"`
	OrgSlug                  string                                     `json:"orgSlug"`
	OrgUrl                   string                                     `json:"orgUrl"`
	GrafanaCloud             float32                                    `json:"grafanaCloud"`
	ResellerId               NullableFloat32                            `json:"resellerId"`
	ContractTypeId           float32                                    `json:"contractTypeId"`
	AllowGCloudTrial         bool                                       `json:"allowGCloudTrial"`
	HlUsage                  float32                                    `json:"hlUsage"`
	HmCurrentGraphiteUsage   float32                                    `json:"hmCurrentGraphiteUsage"`
	HmCurrentPrometheusUsage float32                                    `json:"hmCurrentPrometheusUsage"`
	HgDatasourceCnts         string                                     `json:"hgDatasourceCnts"`
	UserFirstName            string                                     `json:"userFirstName"`
	UserLastName             string                                     `json:"userLastName"`
	UserUsername             string                                     `json:"userUsername"`
	UserStatus               float32                                    `json:"userStatus"`
	UserEmail                string                                     `json:"userEmail"`
	UserName                 string                                     `json:"userName"`
	Subscriptions            Subscriptions                              `json:"subscriptions"`
	MarketplaceSubscription  NullableItemsInner1MarketplaceSubscription `json:"marketplaceSubscription"`
	AdditionalProperties     map[string]interface{}
}

type _ItemsInner1 ItemsInner1

// NewItemsInner1 instantiates a new ItemsInner1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemsInner1(id float32, orgId float32, userId float32, status float32, createdAt string, updatedAt NullableString, defaultOrg float32, role string, privacy float32, billing float32, createdBy string, updatedBy string, orgName string, orgSlug string, orgUrl string, grafanaCloud float32, resellerId NullableFloat32, contractTypeId float32, allowGCloudTrial bool, hlUsage float32, hmCurrentGraphiteUsage float32, hmCurrentPrometheusUsage float32, hgDatasourceCnts string, userFirstName string, userLastName string, userUsername string, userStatus float32, userEmail string, userName string, subscriptions Subscriptions, marketplaceSubscription NullableItemsInner1MarketplaceSubscription) *ItemsInner1 {
	this := ItemsInner1{}
	this.Id = id
	this.OrgId = orgId
	this.UserId = userId
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DefaultOrg = defaultOrg
	this.Role = role
	this.Privacy = privacy
	this.Billing = billing
	this.CreatedBy = createdBy
	this.UpdatedBy = updatedBy
	this.OrgName = orgName
	this.OrgSlug = orgSlug
	this.OrgUrl = orgUrl
	this.GrafanaCloud = grafanaCloud
	this.ResellerId = resellerId
	this.ContractTypeId = contractTypeId
	this.AllowGCloudTrial = allowGCloudTrial
	this.HlUsage = hlUsage
	this.HmCurrentGraphiteUsage = hmCurrentGraphiteUsage
	this.HmCurrentPrometheusUsage = hmCurrentPrometheusUsage
	this.HgDatasourceCnts = hgDatasourceCnts
	this.UserFirstName = userFirstName
	this.UserLastName = userLastName
	this.UserUsername = userUsername
	this.UserStatus = userStatus
	this.UserEmail = userEmail
	this.UserName = userName
	this.Subscriptions = subscriptions
	this.MarketplaceSubscription = marketplaceSubscription
	return &this
}

// NewItemsInner1WithDefaults instantiates a new ItemsInner1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsInner1WithDefaults() *ItemsInner1 {
	this := ItemsInner1{}
	return &this
}

// GetId returns the Id field value
func (o *ItemsInner1) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ItemsInner1) SetId(v float32) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *ItemsInner1) GetOrgId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetOrgIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ItemsInner1) SetOrgId(v float32) {
	o.OrgId = v
}

// GetUserId returns the UserId field value
func (o *ItemsInner1) GetUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ItemsInner1) SetUserId(v float32) {
	o.UserId = v
}

// GetStatus returns the Status field value
func (o *ItemsInner1) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ItemsInner1) SetStatus(v float32) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ItemsInner1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ItemsInner1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemsInner1) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner1) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *ItemsInner1) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// GetDefaultOrg returns the DefaultOrg field value
func (o *ItemsInner1) GetDefaultOrg() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DefaultOrg
}

// GetDefaultOrgOk returns a tuple with the DefaultOrg field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetDefaultOrgOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultOrg, true
}

// SetDefaultOrg sets field value
func (o *ItemsInner1) SetDefaultOrg(v float32) {
	o.DefaultOrg = v
}

// GetRole returns the Role field value
func (o *ItemsInner1) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ItemsInner1) SetRole(v string) {
	o.Role = v
}

// GetPrivacy returns the Privacy field value
func (o *ItemsInner1) GetPrivacy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetPrivacyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *ItemsInner1) SetPrivacy(v float32) {
	o.Privacy = v
}

// GetBilling returns the Billing field value
func (o *ItemsInner1) GetBilling() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Billing
}

// GetBillingOk returns a tuple with the Billing field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetBillingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Billing, true
}

// SetBilling sets field value
func (o *ItemsInner1) SetBilling(v float32) {
	o.Billing = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ItemsInner1) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ItemsInner1) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ItemsInner1) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ItemsInner1) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetOrgName returns the OrgName field value
func (o *ItemsInner1) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *ItemsInner1) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgSlug returns the OrgSlug field value
func (o *ItemsInner1) GetOrgSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgSlug
}

// GetOrgSlugOk returns a tuple with the OrgSlug field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetOrgSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSlug, true
}

// SetOrgSlug sets field value
func (o *ItemsInner1) SetOrgSlug(v string) {
	o.OrgSlug = v
}

// GetOrgUrl returns the OrgUrl field value
func (o *ItemsInner1) GetOrgUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgUrl
}

// GetOrgUrlOk returns a tuple with the OrgUrl field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetOrgUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUrl, true
}

// SetOrgUrl sets field value
func (o *ItemsInner1) SetOrgUrl(v string) {
	o.OrgUrl = v
}

// GetGrafanaCloud returns the GrafanaCloud field value
func (o *ItemsInner1) GetGrafanaCloud() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GrafanaCloud
}

// GetGrafanaCloudOk returns a tuple with the GrafanaCloud field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetGrafanaCloudOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrafanaCloud, true
}

// SetGrafanaCloud sets field value
func (o *ItemsInner1) SetGrafanaCloud(v float32) {
	o.GrafanaCloud = v
}

// GetResellerId returns the ResellerId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ItemsInner1) GetResellerId() float32 {
	if o == nil || o.ResellerId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ResellerId.Get()
}

// GetResellerIdOk returns a tuple with the ResellerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner1) GetResellerIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResellerId.Get(), o.ResellerId.IsSet()
}

// SetResellerId sets field value
func (o *ItemsInner1) SetResellerId(v float32) {
	o.ResellerId.Set(&v)
}

// GetContractTypeId returns the ContractTypeId field value
func (o *ItemsInner1) GetContractTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContractTypeId
}

// GetContractTypeIdOk returns a tuple with the ContractTypeId field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetContractTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractTypeId, true
}

// SetContractTypeId sets field value
func (o *ItemsInner1) SetContractTypeId(v float32) {
	o.ContractTypeId = v
}

// GetAllowGCloudTrial returns the AllowGCloudTrial field value
func (o *ItemsInner1) GetAllowGCloudTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowGCloudTrial
}

// GetAllowGCloudTrialOk returns a tuple with the AllowGCloudTrial field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetAllowGCloudTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowGCloudTrial, true
}

// SetAllowGCloudTrial sets field value
func (o *ItemsInner1) SetAllowGCloudTrial(v bool) {
	o.AllowGCloudTrial = v
}

// GetHlUsage returns the HlUsage field value
func (o *ItemsInner1) GetHlUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HlUsage
}

// GetHlUsageOk returns a tuple with the HlUsage field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetHlUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HlUsage, true
}

// SetHlUsage sets field value
func (o *ItemsInner1) SetHlUsage(v float32) {
	o.HlUsage = v
}

// GetHmCurrentGraphiteUsage returns the HmCurrentGraphiteUsage field value
func (o *ItemsInner1) GetHmCurrentGraphiteUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HmCurrentGraphiteUsage
}

// GetHmCurrentGraphiteUsageOk returns a tuple with the HmCurrentGraphiteUsage field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetHmCurrentGraphiteUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmCurrentGraphiteUsage, true
}

// SetHmCurrentGraphiteUsage sets field value
func (o *ItemsInner1) SetHmCurrentGraphiteUsage(v float32) {
	o.HmCurrentGraphiteUsage = v
}

// GetHmCurrentPrometheusUsage returns the HmCurrentPrometheusUsage field value
func (o *ItemsInner1) GetHmCurrentPrometheusUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HmCurrentPrometheusUsage
}

// GetHmCurrentPrometheusUsageOk returns a tuple with the HmCurrentPrometheusUsage field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetHmCurrentPrometheusUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmCurrentPrometheusUsage, true
}

// SetHmCurrentPrometheusUsage sets field value
func (o *ItemsInner1) SetHmCurrentPrometheusUsage(v float32) {
	o.HmCurrentPrometheusUsage = v
}

// GetHgDatasourceCnts returns the HgDatasourceCnts field value
func (o *ItemsInner1) GetHgDatasourceCnts() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HgDatasourceCnts
}

// GetHgDatasourceCntsOk returns a tuple with the HgDatasourceCnts field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetHgDatasourceCntsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HgDatasourceCnts, true
}

// SetHgDatasourceCnts sets field value
func (o *ItemsInner1) SetHgDatasourceCnts(v string) {
	o.HgDatasourceCnts = v
}

// GetUserFirstName returns the UserFirstName field value
func (o *ItemsInner1) GetUserFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserFirstName
}

// GetUserFirstNameOk returns a tuple with the UserFirstName field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserFirstName, true
}

// SetUserFirstName sets field value
func (o *ItemsInner1) SetUserFirstName(v string) {
	o.UserFirstName = v
}

// GetUserLastName returns the UserLastName field value
func (o *ItemsInner1) GetUserLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserLastName
}

// GetUserLastNameOk returns a tuple with the UserLastName field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLastName, true
}

// SetUserLastName sets field value
func (o *ItemsInner1) SetUserLastName(v string) {
	o.UserLastName = v
}

// GetUserUsername returns the UserUsername field value
func (o *ItemsInner1) GetUserUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserUsername
}

// GetUserUsernameOk returns a tuple with the UserUsername field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUsername, true
}

// SetUserUsername sets field value
func (o *ItemsInner1) SetUserUsername(v string) {
	o.UserUsername = v
}

// GetUserStatus returns the UserStatus field value
func (o *ItemsInner1) GetUserStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserStatus, true
}

// SetUserStatus sets field value
func (o *ItemsInner1) SetUserStatus(v float32) {
	o.UserStatus = v
}

// GetUserEmail returns the UserEmail field value
func (o *ItemsInner1) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *ItemsInner1) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetUserName returns the UserName field value
func (o *ItemsInner1) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ItemsInner1) SetUserName(v string) {
	o.UserName = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *ItemsInner1) GetSubscriptions() Subscriptions {
	if o == nil {
		var ret Subscriptions
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ItemsInner1) GetSubscriptionsOk() (*Subscriptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ItemsInner1) SetSubscriptions(v Subscriptions) {
	o.Subscriptions = v
}

// GetMarketplaceSubscription returns the MarketplaceSubscription field value
// If the value is explicit nil, the zero value for ItemsInner1MarketplaceSubscription will be returned
func (o *ItemsInner1) GetMarketplaceSubscription() ItemsInner1MarketplaceSubscription {
	if o == nil || o.MarketplaceSubscription.Get() == nil {
		var ret ItemsInner1MarketplaceSubscription
		return ret
	}

	return *o.MarketplaceSubscription.Get()
}

// GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner1) GetMarketplaceSubscriptionOk() (*ItemsInner1MarketplaceSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceSubscription.Get(), o.MarketplaceSubscription.IsSet()
}

// SetMarketplaceSubscription sets field value
func (o *ItemsInner1) SetMarketplaceSubscription(v ItemsInner1MarketplaceSubscription) {
	o.MarketplaceSubscription.Set(&v)
}

func (o ItemsInner1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemsInner1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["orgId"] = o.OrgId
	toSerialize["userId"] = o.UserId
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt.Get()
	toSerialize["defaultOrg"] = o.DefaultOrg
	toSerialize["role"] = o.Role
	toSerialize["privacy"] = o.Privacy
	toSerialize["billing"] = o.Billing
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["updatedBy"] = o.UpdatedBy
	toSerialize["orgName"] = o.OrgName
	toSerialize["orgSlug"] = o.OrgSlug
	toSerialize["orgUrl"] = o.OrgUrl
	toSerialize["grafanaCloud"] = o.GrafanaCloud
	toSerialize["resellerId"] = o.ResellerId.Get()
	toSerialize["contractTypeId"] = o.ContractTypeId
	toSerialize["allowGCloudTrial"] = o.AllowGCloudTrial
	toSerialize["hlUsage"] = o.HlUsage
	toSerialize["hmCurrentGraphiteUsage"] = o.HmCurrentGraphiteUsage
	toSerialize["hmCurrentPrometheusUsage"] = o.HmCurrentPrometheusUsage
	toSerialize["hgDatasourceCnts"] = o.HgDatasourceCnts
	toSerialize["userFirstName"] = o.UserFirstName
	toSerialize["userLastName"] = o.UserLastName
	toSerialize["userUsername"] = o.UserUsername
	toSerialize["userStatus"] = o.UserStatus
	toSerialize["userEmail"] = o.UserEmail
	toSerialize["userName"] = o.UserName
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["marketplaceSubscription"] = o.MarketplaceSubscription.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ItemsInner1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"orgId",
		"userId",
		"status",
		"createdAt",
		"updatedAt",
		"defaultOrg",
		"role",
		"privacy",
		"billing",
		"createdBy",
		"updatedBy",
		"orgName",
		"orgSlug",
		"orgUrl",
		"grafanaCloud",
		"resellerId",
		"contractTypeId",
		"allowGCloudTrial",
		"hlUsage",
		"hmCurrentGraphiteUsage",
		"hmCurrentPrometheusUsage",
		"hgDatasourceCnts",
		"userFirstName",
		"userLastName",
		"userUsername",
		"userStatus",
		"userEmail",
		"userName",
		"subscriptions",
		"marketplaceSubscription",
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

	varItemsInner1 := _ItemsInner1{}

	err = json.Unmarshal(data, &varItemsInner1)

	if err != nil {
		return err
	}

	*o = ItemsInner1(varItemsInner1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "defaultOrg")
		delete(additionalProperties, "role")
		delete(additionalProperties, "privacy")
		delete(additionalProperties, "billing")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "updatedBy")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgSlug")
		delete(additionalProperties, "orgUrl")
		delete(additionalProperties, "grafanaCloud")
		delete(additionalProperties, "resellerId")
		delete(additionalProperties, "contractTypeId")
		delete(additionalProperties, "allowGCloudTrial")
		delete(additionalProperties, "hlUsage")
		delete(additionalProperties, "hmCurrentGraphiteUsage")
		delete(additionalProperties, "hmCurrentPrometheusUsage")
		delete(additionalProperties, "hgDatasourceCnts")
		delete(additionalProperties, "userFirstName")
		delete(additionalProperties, "userLastName")
		delete(additionalProperties, "userUsername")
		delete(additionalProperties, "userStatus")
		delete(additionalProperties, "userEmail")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "subscriptions")
		delete(additionalProperties, "marketplaceSubscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemsInner1 struct {
	value *ItemsInner1
	isSet bool
}

func (v NullableItemsInner1) Get() *ItemsInner1 {
	return v.value
}

func (v *NullableItemsInner1) Set(val *ItemsInner1) {
	v.value = val
	v.isSet = true
}

func (v NullableItemsInner1) IsSet() bool {
	return v.isSet
}

func (v *NullableItemsInner1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemsInner1(val *ItemsInner1) *NullableItemsInner1 {
	return &NullableItemsInner1{value: val, isSet: true}
}

func (v NullableItemsInner1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemsInner1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
