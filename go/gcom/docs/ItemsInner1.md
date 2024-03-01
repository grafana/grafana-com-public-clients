# ItemsInner1

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

### NewItemsInner1

`func NewItemsInner1() *ItemsInner1`

NewItemsInner1 instantiates a new ItemsInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsInner1WithDefaults

`func NewItemsInner1WithDefaults() *ItemsInner1`

NewItemsInner1WithDefaults instantiates a new ItemsInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemsInner1) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemsInner1) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemsInner1) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ItemsInner1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *ItemsInner1) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ItemsInner1) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ItemsInner1) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ItemsInner1) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUserId

`func (o *ItemsInner1) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ItemsInner1) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ItemsInner1) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ItemsInner1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *ItemsInner1) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemsInner1) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemsInner1) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemsInner1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ItemsInner1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemsInner1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemsInner1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ItemsInner1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ItemsInner1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemsInner1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemsInner1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ItemsInner1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ItemsInner1) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ItemsInner1) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDefaultOrg

`func (o *ItemsInner1) GetDefaultOrg() float32`

GetDefaultOrg returns the DefaultOrg field if non-nil, zero value otherwise.

### GetDefaultOrgOk

`func (o *ItemsInner1) GetDefaultOrgOk() (*float32, bool)`

GetDefaultOrgOk returns a tuple with the DefaultOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrg

`func (o *ItemsInner1) SetDefaultOrg(v float32)`

SetDefaultOrg sets DefaultOrg field to given value.

### HasDefaultOrg

`func (o *ItemsInner1) HasDefaultOrg() bool`

HasDefaultOrg returns a boolean if a field has been set.

### GetRole

`func (o *ItemsInner1) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ItemsInner1) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ItemsInner1) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ItemsInner1) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPrivacy

`func (o *ItemsInner1) GetPrivacy() float32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ItemsInner1) GetPrivacyOk() (*float32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ItemsInner1) SetPrivacy(v float32)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *ItemsInner1) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetBilling

`func (o *ItemsInner1) GetBilling() float32`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ItemsInner1) GetBillingOk() (*float32, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ItemsInner1) SetBilling(v float32)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *ItemsInner1) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ItemsInner1) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ItemsInner1) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ItemsInner1) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ItemsInner1) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ItemsInner1) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ItemsInner1) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ItemsInner1) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ItemsInner1) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetOrgName

`func (o *ItemsInner1) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ItemsInner1) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ItemsInner1) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ItemsInner1) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgSlug

`func (o *ItemsInner1) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *ItemsInner1) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *ItemsInner1) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.

### HasOrgSlug

`func (o *ItemsInner1) HasOrgSlug() bool`

HasOrgSlug returns a boolean if a field has been set.

### GetOrgUrl

`func (o *ItemsInner1) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *ItemsInner1) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *ItemsInner1) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.

### HasOrgUrl

`func (o *ItemsInner1) HasOrgUrl() bool`

HasOrgUrl returns a boolean if a field has been set.

### GetGrafanaCloud

`func (o *ItemsInner1) GetGrafanaCloud() float32`

GetGrafanaCloud returns the GrafanaCloud field if non-nil, zero value otherwise.

### GetGrafanaCloudOk

`func (o *ItemsInner1) GetGrafanaCloudOk() (*float32, bool)`

GetGrafanaCloudOk returns a tuple with the GrafanaCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaCloud

`func (o *ItemsInner1) SetGrafanaCloud(v float32)`

SetGrafanaCloud sets GrafanaCloud field to given value.

### HasGrafanaCloud

`func (o *ItemsInner1) HasGrafanaCloud() bool`

HasGrafanaCloud returns a boolean if a field has been set.

### GetResellerId

`func (o *ItemsInner1) GetResellerId() float32`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *ItemsInner1) GetResellerIdOk() (*float32, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *ItemsInner1) SetResellerId(v float32)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *ItemsInner1) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### SetResellerIdNil

`func (o *ItemsInner1) SetResellerIdNil(b bool)`

 SetResellerIdNil sets the value for ResellerId to be an explicit nil

### UnsetResellerId
`func (o *ItemsInner1) UnsetResellerId()`

UnsetResellerId ensures that no value is present for ResellerId, not even an explicit nil
### GetContractTypeId

`func (o *ItemsInner1) GetContractTypeId() float32`

GetContractTypeId returns the ContractTypeId field if non-nil, zero value otherwise.

### GetContractTypeIdOk

`func (o *ItemsInner1) GetContractTypeIdOk() (*float32, bool)`

GetContractTypeIdOk returns a tuple with the ContractTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypeId

`func (o *ItemsInner1) SetContractTypeId(v float32)`

SetContractTypeId sets ContractTypeId field to given value.

### HasContractTypeId

`func (o *ItemsInner1) HasContractTypeId() bool`

HasContractTypeId returns a boolean if a field has been set.

### GetAllowGCloudTrial

`func (o *ItemsInner1) GetAllowGCloudTrial() bool`

GetAllowGCloudTrial returns the AllowGCloudTrial field if non-nil, zero value otherwise.

### GetAllowGCloudTrialOk

`func (o *ItemsInner1) GetAllowGCloudTrialOk() (*bool, bool)`

GetAllowGCloudTrialOk returns a tuple with the AllowGCloudTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGCloudTrial

`func (o *ItemsInner1) SetAllowGCloudTrial(v bool)`

SetAllowGCloudTrial sets AllowGCloudTrial field to given value.

### HasAllowGCloudTrial

`func (o *ItemsInner1) HasAllowGCloudTrial() bool`

HasAllowGCloudTrial returns a boolean if a field has been set.

### GetHlUsage

`func (o *ItemsInner1) GetHlUsage() float32`

GetHlUsage returns the HlUsage field if non-nil, zero value otherwise.

### GetHlUsageOk

`func (o *ItemsInner1) GetHlUsageOk() (*float32, bool)`

GetHlUsageOk returns a tuple with the HlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlUsage

`func (o *ItemsInner1) SetHlUsage(v float32)`

SetHlUsage sets HlUsage field to given value.

### HasHlUsage

`func (o *ItemsInner1) HasHlUsage() bool`

HasHlUsage returns a boolean if a field has been set.

### GetHmCurrentGraphiteUsage

`func (o *ItemsInner1) GetHmCurrentGraphiteUsage() float32`

GetHmCurrentGraphiteUsage returns the HmCurrentGraphiteUsage field if non-nil, zero value otherwise.

### GetHmCurrentGraphiteUsageOk

`func (o *ItemsInner1) GetHmCurrentGraphiteUsageOk() (*float32, bool)`

GetHmCurrentGraphiteUsageOk returns a tuple with the HmCurrentGraphiteUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentGraphiteUsage

`func (o *ItemsInner1) SetHmCurrentGraphiteUsage(v float32)`

SetHmCurrentGraphiteUsage sets HmCurrentGraphiteUsage field to given value.

### HasHmCurrentGraphiteUsage

`func (o *ItemsInner1) HasHmCurrentGraphiteUsage() bool`

HasHmCurrentGraphiteUsage returns a boolean if a field has been set.

### GetHmCurrentPrometheusUsage

`func (o *ItemsInner1) GetHmCurrentPrometheusUsage() float32`

GetHmCurrentPrometheusUsage returns the HmCurrentPrometheusUsage field if non-nil, zero value otherwise.

### GetHmCurrentPrometheusUsageOk

`func (o *ItemsInner1) GetHmCurrentPrometheusUsageOk() (*float32, bool)`

GetHmCurrentPrometheusUsageOk returns a tuple with the HmCurrentPrometheusUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentPrometheusUsage

`func (o *ItemsInner1) SetHmCurrentPrometheusUsage(v float32)`

SetHmCurrentPrometheusUsage sets HmCurrentPrometheusUsage field to given value.

### HasHmCurrentPrometheusUsage

`func (o *ItemsInner1) HasHmCurrentPrometheusUsage() bool`

HasHmCurrentPrometheusUsage returns a boolean if a field has been set.

### GetHgDatasourceCnts

`func (o *ItemsInner1) GetHgDatasourceCnts() string`

GetHgDatasourceCnts returns the HgDatasourceCnts field if non-nil, zero value otherwise.

### GetHgDatasourceCntsOk

`func (o *ItemsInner1) GetHgDatasourceCntsOk() (*string, bool)`

GetHgDatasourceCntsOk returns a tuple with the HgDatasourceCnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgDatasourceCnts

`func (o *ItemsInner1) SetHgDatasourceCnts(v string)`

SetHgDatasourceCnts sets HgDatasourceCnts field to given value.

### HasHgDatasourceCnts

`func (o *ItemsInner1) HasHgDatasourceCnts() bool`

HasHgDatasourceCnts returns a boolean if a field has been set.

### GetUserFirstName

`func (o *ItemsInner1) GetUserFirstName() string`

GetUserFirstName returns the UserFirstName field if non-nil, zero value otherwise.

### GetUserFirstNameOk

`func (o *ItemsInner1) GetUserFirstNameOk() (*string, bool)`

GetUserFirstNameOk returns a tuple with the UserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFirstName

`func (o *ItemsInner1) SetUserFirstName(v string)`

SetUserFirstName sets UserFirstName field to given value.

### HasUserFirstName

`func (o *ItemsInner1) HasUserFirstName() bool`

HasUserFirstName returns a boolean if a field has been set.

### GetUserLastName

`func (o *ItemsInner1) GetUserLastName() string`

GetUserLastName returns the UserLastName field if non-nil, zero value otherwise.

### GetUserLastNameOk

`func (o *ItemsInner1) GetUserLastNameOk() (*string, bool)`

GetUserLastNameOk returns a tuple with the UserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLastName

`func (o *ItemsInner1) SetUserLastName(v string)`

SetUserLastName sets UserLastName field to given value.

### HasUserLastName

`func (o *ItemsInner1) HasUserLastName() bool`

HasUserLastName returns a boolean if a field has been set.

### GetUserUsername

`func (o *ItemsInner1) GetUserUsername() string`

GetUserUsername returns the UserUsername field if non-nil, zero value otherwise.

### GetUserUsernameOk

`func (o *ItemsInner1) GetUserUsernameOk() (*string, bool)`

GetUserUsernameOk returns a tuple with the UserUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUsername

`func (o *ItemsInner1) SetUserUsername(v string)`

SetUserUsername sets UserUsername field to given value.

### HasUserUsername

`func (o *ItemsInner1) HasUserUsername() bool`

HasUserUsername returns a boolean if a field has been set.

### GetUserStatus

`func (o *ItemsInner1) GetUserStatus() float32`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *ItemsInner1) GetUserStatusOk() (*float32, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *ItemsInner1) SetUserStatus(v float32)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *ItemsInner1) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetUserEmail

`func (o *ItemsInner1) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ItemsInner1) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ItemsInner1) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *ItemsInner1) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserName

`func (o *ItemsInner1) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ItemsInner1) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ItemsInner1) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ItemsInner1) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ItemsInner1) GetSubscriptions() Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ItemsInner1) GetSubscriptionsOk() (*Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ItemsInner1) SetSubscriptions(v Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ItemsInner1) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *ItemsInner1) GetMarketplaceSubscription() ItemsInner1MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *ItemsInner1) GetMarketplaceSubscriptionOk() (*ItemsInner1MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *ItemsInner1) SetMarketplaceSubscription(v ItemsInner1MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *ItemsInner1) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### SetMarketplaceSubscriptionNil

`func (o *ItemsInner1) SetMarketplaceSubscriptionNil(b bool)`

 SetMarketplaceSubscriptionNil sets the value for MarketplaceSubscription to be an explicit nil

### UnsetMarketplaceSubscription
`func (o *ItemsInner1) UnsetMarketplaceSubscription()`

UnsetMarketplaceSubscription ensures that no value is present for MarketplaceSubscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


