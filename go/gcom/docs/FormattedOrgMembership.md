# FormattedOrgMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**DefaultOrg** | Pointer to **float32** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to **float32** |  | [optional] 
**Billing** | Pointer to **float32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**OrgSlug** | Pointer to **string** |  | [optional] 
**OrgUrl** | Pointer to **string** |  | [optional] 
**GrafanaCloud** | Pointer to **float32** |  | [optional] 
**ResellerId** | Pointer to **NullableFloat32** |  | [optional] 
**ContractTypeId** | Pointer to **float32** |  | [optional] 
**AllowGCloudTrial** | Pointer to **bool** |  | [optional] 
**HlUsage** | Pointer to **float32** |  | [optional] 
**HmCurrentGraphiteUsage** | Pointer to **float32** |  | [optional] 
**HmCurrentPrometheusUsage** | Pointer to **float32** |  | [optional] 
**HgDatasourceCnts** | Pointer to **string** |  | [optional] 
**UserFirstName** | Pointer to **string** |  | [optional] 
**UserLastName** | Pointer to **string** |  | [optional] 
**UserUsername** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **float32** |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**Subscriptions**](Subscriptions.md) |  | [optional] 
**MarketplaceSubscription** | Pointer to [**NullableItemsInner1MarketplaceSubscription**](ItemsInner1MarketplaceSubscription.md) |  | [optional] 

## Methods

### NewFormattedOrgMembership

`func NewFormattedOrgMembership() *FormattedOrgMembership`

NewFormattedOrgMembership instantiates a new FormattedOrgMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedOrgMembershipWithDefaults

`func NewFormattedOrgMembershipWithDefaults() *FormattedOrgMembership`

NewFormattedOrgMembershipWithDefaults instantiates a new FormattedOrgMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormattedOrgMembership) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedOrgMembership) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedOrgMembership) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *FormattedOrgMembership) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *FormattedOrgMembership) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FormattedOrgMembership) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FormattedOrgMembership) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FormattedOrgMembership) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUserId

`func (o *FormattedOrgMembership) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FormattedOrgMembership) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FormattedOrgMembership) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FormattedOrgMembership) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *FormattedOrgMembership) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormattedOrgMembership) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormattedOrgMembership) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FormattedOrgMembership) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FormattedOrgMembership) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedOrgMembership) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedOrgMembership) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormattedOrgMembership) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormattedOrgMembership) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedOrgMembership) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedOrgMembership) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormattedOrgMembership) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *FormattedOrgMembership) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedOrgMembership) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDefaultOrg

`func (o *FormattedOrgMembership) GetDefaultOrg() float32`

GetDefaultOrg returns the DefaultOrg field if non-nil, zero value otherwise.

### GetDefaultOrgOk

`func (o *FormattedOrgMembership) GetDefaultOrgOk() (*float32, bool)`

GetDefaultOrgOk returns a tuple with the DefaultOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrg

`func (o *FormattedOrgMembership) SetDefaultOrg(v float32)`

SetDefaultOrg sets DefaultOrg field to given value.

### HasDefaultOrg

`func (o *FormattedOrgMembership) HasDefaultOrg() bool`

HasDefaultOrg returns a boolean if a field has been set.

### GetRole

`func (o *FormattedOrgMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FormattedOrgMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FormattedOrgMembership) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FormattedOrgMembership) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPrivacy

`func (o *FormattedOrgMembership) GetPrivacy() float32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *FormattedOrgMembership) GetPrivacyOk() (*float32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *FormattedOrgMembership) SetPrivacy(v float32)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *FormattedOrgMembership) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetBilling

`func (o *FormattedOrgMembership) GetBilling() float32`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *FormattedOrgMembership) GetBillingOk() (*float32, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *FormattedOrgMembership) SetBilling(v float32)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *FormattedOrgMembership) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FormattedOrgMembership) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FormattedOrgMembership) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FormattedOrgMembership) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FormattedOrgMembership) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FormattedOrgMembership) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FormattedOrgMembership) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FormattedOrgMembership) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FormattedOrgMembership) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetOrgName

`func (o *FormattedOrgMembership) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FormattedOrgMembership) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FormattedOrgMembership) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *FormattedOrgMembership) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgSlug

`func (o *FormattedOrgMembership) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *FormattedOrgMembership) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *FormattedOrgMembership) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.

### HasOrgSlug

`func (o *FormattedOrgMembership) HasOrgSlug() bool`

HasOrgSlug returns a boolean if a field has been set.

### GetOrgUrl

`func (o *FormattedOrgMembership) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *FormattedOrgMembership) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *FormattedOrgMembership) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.

### HasOrgUrl

`func (o *FormattedOrgMembership) HasOrgUrl() bool`

HasOrgUrl returns a boolean if a field has been set.

### GetGrafanaCloud

`func (o *FormattedOrgMembership) GetGrafanaCloud() float32`

GetGrafanaCloud returns the GrafanaCloud field if non-nil, zero value otherwise.

### GetGrafanaCloudOk

`func (o *FormattedOrgMembership) GetGrafanaCloudOk() (*float32, bool)`

GetGrafanaCloudOk returns a tuple with the GrafanaCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaCloud

`func (o *FormattedOrgMembership) SetGrafanaCloud(v float32)`

SetGrafanaCloud sets GrafanaCloud field to given value.

### HasGrafanaCloud

`func (o *FormattedOrgMembership) HasGrafanaCloud() bool`

HasGrafanaCloud returns a boolean if a field has been set.

### GetResellerId

`func (o *FormattedOrgMembership) GetResellerId() float32`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *FormattedOrgMembership) GetResellerIdOk() (*float32, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *FormattedOrgMembership) SetResellerId(v float32)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *FormattedOrgMembership) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### SetResellerIdNil

`func (o *FormattedOrgMembership) SetResellerIdNil(b bool)`

 SetResellerIdNil sets the value for ResellerId to be an explicit nil

### UnsetResellerId
`func (o *FormattedOrgMembership) UnsetResellerId()`

UnsetResellerId ensures that no value is present for ResellerId, not even an explicit nil
### GetContractTypeId

`func (o *FormattedOrgMembership) GetContractTypeId() float32`

GetContractTypeId returns the ContractTypeId field if non-nil, zero value otherwise.

### GetContractTypeIdOk

`func (o *FormattedOrgMembership) GetContractTypeIdOk() (*float32, bool)`

GetContractTypeIdOk returns a tuple with the ContractTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypeId

`func (o *FormattedOrgMembership) SetContractTypeId(v float32)`

SetContractTypeId sets ContractTypeId field to given value.

### HasContractTypeId

`func (o *FormattedOrgMembership) HasContractTypeId() bool`

HasContractTypeId returns a boolean if a field has been set.

### GetAllowGCloudTrial

`func (o *FormattedOrgMembership) GetAllowGCloudTrial() bool`

GetAllowGCloudTrial returns the AllowGCloudTrial field if non-nil, zero value otherwise.

### GetAllowGCloudTrialOk

`func (o *FormattedOrgMembership) GetAllowGCloudTrialOk() (*bool, bool)`

GetAllowGCloudTrialOk returns a tuple with the AllowGCloudTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGCloudTrial

`func (o *FormattedOrgMembership) SetAllowGCloudTrial(v bool)`

SetAllowGCloudTrial sets AllowGCloudTrial field to given value.

### HasAllowGCloudTrial

`func (o *FormattedOrgMembership) HasAllowGCloudTrial() bool`

HasAllowGCloudTrial returns a boolean if a field has been set.

### GetHlUsage

`func (o *FormattedOrgMembership) GetHlUsage() float32`

GetHlUsage returns the HlUsage field if non-nil, zero value otherwise.

### GetHlUsageOk

`func (o *FormattedOrgMembership) GetHlUsageOk() (*float32, bool)`

GetHlUsageOk returns a tuple with the HlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlUsage

`func (o *FormattedOrgMembership) SetHlUsage(v float32)`

SetHlUsage sets HlUsage field to given value.

### HasHlUsage

`func (o *FormattedOrgMembership) HasHlUsage() bool`

HasHlUsage returns a boolean if a field has been set.

### GetHmCurrentGraphiteUsage

`func (o *FormattedOrgMembership) GetHmCurrentGraphiteUsage() float32`

GetHmCurrentGraphiteUsage returns the HmCurrentGraphiteUsage field if non-nil, zero value otherwise.

### GetHmCurrentGraphiteUsageOk

`func (o *FormattedOrgMembership) GetHmCurrentGraphiteUsageOk() (*float32, bool)`

GetHmCurrentGraphiteUsageOk returns a tuple with the HmCurrentGraphiteUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentGraphiteUsage

`func (o *FormattedOrgMembership) SetHmCurrentGraphiteUsage(v float32)`

SetHmCurrentGraphiteUsage sets HmCurrentGraphiteUsage field to given value.

### HasHmCurrentGraphiteUsage

`func (o *FormattedOrgMembership) HasHmCurrentGraphiteUsage() bool`

HasHmCurrentGraphiteUsage returns a boolean if a field has been set.

### GetHmCurrentPrometheusUsage

`func (o *FormattedOrgMembership) GetHmCurrentPrometheusUsage() float32`

GetHmCurrentPrometheusUsage returns the HmCurrentPrometheusUsage field if non-nil, zero value otherwise.

### GetHmCurrentPrometheusUsageOk

`func (o *FormattedOrgMembership) GetHmCurrentPrometheusUsageOk() (*float32, bool)`

GetHmCurrentPrometheusUsageOk returns a tuple with the HmCurrentPrometheusUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentPrometheusUsage

`func (o *FormattedOrgMembership) SetHmCurrentPrometheusUsage(v float32)`

SetHmCurrentPrometheusUsage sets HmCurrentPrometheusUsage field to given value.

### HasHmCurrentPrometheusUsage

`func (o *FormattedOrgMembership) HasHmCurrentPrometheusUsage() bool`

HasHmCurrentPrometheusUsage returns a boolean if a field has been set.

### GetHgDatasourceCnts

`func (o *FormattedOrgMembership) GetHgDatasourceCnts() string`

GetHgDatasourceCnts returns the HgDatasourceCnts field if non-nil, zero value otherwise.

### GetHgDatasourceCntsOk

`func (o *FormattedOrgMembership) GetHgDatasourceCntsOk() (*string, bool)`

GetHgDatasourceCntsOk returns a tuple with the HgDatasourceCnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgDatasourceCnts

`func (o *FormattedOrgMembership) SetHgDatasourceCnts(v string)`

SetHgDatasourceCnts sets HgDatasourceCnts field to given value.

### HasHgDatasourceCnts

`func (o *FormattedOrgMembership) HasHgDatasourceCnts() bool`

HasHgDatasourceCnts returns a boolean if a field has been set.

### GetUserFirstName

`func (o *FormattedOrgMembership) GetUserFirstName() string`

GetUserFirstName returns the UserFirstName field if non-nil, zero value otherwise.

### GetUserFirstNameOk

`func (o *FormattedOrgMembership) GetUserFirstNameOk() (*string, bool)`

GetUserFirstNameOk returns a tuple with the UserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFirstName

`func (o *FormattedOrgMembership) SetUserFirstName(v string)`

SetUserFirstName sets UserFirstName field to given value.

### HasUserFirstName

`func (o *FormattedOrgMembership) HasUserFirstName() bool`

HasUserFirstName returns a boolean if a field has been set.

### GetUserLastName

`func (o *FormattedOrgMembership) GetUserLastName() string`

GetUserLastName returns the UserLastName field if non-nil, zero value otherwise.

### GetUserLastNameOk

`func (o *FormattedOrgMembership) GetUserLastNameOk() (*string, bool)`

GetUserLastNameOk returns a tuple with the UserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLastName

`func (o *FormattedOrgMembership) SetUserLastName(v string)`

SetUserLastName sets UserLastName field to given value.

### HasUserLastName

`func (o *FormattedOrgMembership) HasUserLastName() bool`

HasUserLastName returns a boolean if a field has been set.

### GetUserUsername

`func (o *FormattedOrgMembership) GetUserUsername() string`

GetUserUsername returns the UserUsername field if non-nil, zero value otherwise.

### GetUserUsernameOk

`func (o *FormattedOrgMembership) GetUserUsernameOk() (*string, bool)`

GetUserUsernameOk returns a tuple with the UserUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUsername

`func (o *FormattedOrgMembership) SetUserUsername(v string)`

SetUserUsername sets UserUsername field to given value.

### HasUserUsername

`func (o *FormattedOrgMembership) HasUserUsername() bool`

HasUserUsername returns a boolean if a field has been set.

### GetUserStatus

`func (o *FormattedOrgMembership) GetUserStatus() float32`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *FormattedOrgMembership) GetUserStatusOk() (*float32, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *FormattedOrgMembership) SetUserStatus(v float32)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *FormattedOrgMembership) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetUserEmail

`func (o *FormattedOrgMembership) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *FormattedOrgMembership) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *FormattedOrgMembership) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *FormattedOrgMembership) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserName

`func (o *FormattedOrgMembership) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FormattedOrgMembership) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FormattedOrgMembership) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FormattedOrgMembership) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetSubscriptions

`func (o *FormattedOrgMembership) GetSubscriptions() Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FormattedOrgMembership) GetSubscriptionsOk() (*Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FormattedOrgMembership) SetSubscriptions(v Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FormattedOrgMembership) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *FormattedOrgMembership) GetMarketplaceSubscription() ItemsInner1MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *FormattedOrgMembership) GetMarketplaceSubscriptionOk() (*ItemsInner1MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *FormattedOrgMembership) SetMarketplaceSubscription(v ItemsInner1MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *FormattedOrgMembership) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### SetMarketplaceSubscriptionNil

`func (o *FormattedOrgMembership) SetMarketplaceSubscriptionNil(b bool)`

 SetMarketplaceSubscriptionNil sets the value for MarketplaceSubscription to be an explicit nil

### UnsetMarketplaceSubscription
`func (o *FormattedOrgMembership) UnsetMarketplaceSubscription()`

UnsetMarketplaceSubscription ensures that no value is present for MarketplaceSubscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


