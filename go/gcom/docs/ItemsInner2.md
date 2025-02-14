# ItemsInner2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**OrgId** | **float32** |  | 
**UserId** | **float32** |  | 
**Status** | **float32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 
**DefaultOrg** | **float32** |  | 
**Role** | **string** |  | 
**Privacy** | **float32** |  | 
**Billing** | **float32** |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**OrgUrl** | **string** |  | 
**GrafanaCloud** | **float32** |  | 
**ResellerId** | **NullableFloat32** |  | 
**ContractTypeId** | **float32** |  | 
**AllowGCloudTrial** | **bool** |  | 
**HlUsage** | **float32** |  | 
**HmCurrentGraphiteUsage** | **float32** |  | 
**HmCurrentPrometheusUsage** | **float32** |  | 
**HgDatasourceCnts** | **string** |  | 
**UserFirstName** | **string** |  | 
**UserLastName** | **string** |  | 
**UserUsername** | **string** |  | 
**UserStatus** | **float32** |  | 
**UserEmail** | **string** |  | 
**UserName** | **string** |  | 
**Subscriptions** | [**Subscriptions**](Subscriptions.md) |  | 
**MarketplaceSubscription** | [**NullableItemsInner2MarketplaceSubscription**](ItemsInner2MarketplaceSubscription.md) |  | 
**ExtraPermissions** | Pointer to **[]string** |  | [optional] 
**GrafanaStaffAccess** | Pointer to [**NullableItemsInner2GrafanaStaffAccess**](ItemsInner2GrafanaStaffAccess.md) |  | [optional] 

## Methods

### NewItemsInner2

`func NewItemsInner2(id float32, orgId float32, userId float32, status float32, createdAt string, updatedAt NullableString, defaultOrg float32, role string, privacy float32, billing float32, createdBy string, updatedBy string, orgName string, orgSlug string, orgUrl string, grafanaCloud float32, resellerId NullableFloat32, contractTypeId float32, allowGCloudTrial bool, hlUsage float32, hmCurrentGraphiteUsage float32, hmCurrentPrometheusUsage float32, hgDatasourceCnts string, userFirstName string, userLastName string, userUsername string, userStatus float32, userEmail string, userName string, subscriptions Subscriptions, marketplaceSubscription NullableItemsInner2MarketplaceSubscription, ) *ItemsInner2`

NewItemsInner2 instantiates a new ItemsInner2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsInner2WithDefaults

`func NewItemsInner2WithDefaults() *ItemsInner2`

NewItemsInner2WithDefaults instantiates a new ItemsInner2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemsInner2) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemsInner2) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemsInner2) SetId(v float32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ItemsInner2) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ItemsInner2) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ItemsInner2) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *ItemsInner2) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ItemsInner2) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ItemsInner2) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *ItemsInner2) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemsInner2) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemsInner2) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ItemsInner2) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemsInner2) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemsInner2) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ItemsInner2) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemsInner2) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemsInner2) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ItemsInner2) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ItemsInner2) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDefaultOrg

`func (o *ItemsInner2) GetDefaultOrg() float32`

GetDefaultOrg returns the DefaultOrg field if non-nil, zero value otherwise.

### GetDefaultOrgOk

`func (o *ItemsInner2) GetDefaultOrgOk() (*float32, bool)`

GetDefaultOrgOk returns a tuple with the DefaultOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrg

`func (o *ItemsInner2) SetDefaultOrg(v float32)`

SetDefaultOrg sets DefaultOrg field to given value.


### GetRole

`func (o *ItemsInner2) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ItemsInner2) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ItemsInner2) SetRole(v string)`

SetRole sets Role field to given value.


### GetPrivacy

`func (o *ItemsInner2) GetPrivacy() float32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ItemsInner2) GetPrivacyOk() (*float32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ItemsInner2) SetPrivacy(v float32)`

SetPrivacy sets Privacy field to given value.


### GetBilling

`func (o *ItemsInner2) GetBilling() float32`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ItemsInner2) GetBillingOk() (*float32, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ItemsInner2) SetBilling(v float32)`

SetBilling sets Billing field to given value.


### GetCreatedBy

`func (o *ItemsInner2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ItemsInner2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ItemsInner2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *ItemsInner2) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ItemsInner2) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ItemsInner2) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetOrgName

`func (o *ItemsInner2) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ItemsInner2) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ItemsInner2) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *ItemsInner2) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *ItemsInner2) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *ItemsInner2) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetOrgUrl

`func (o *ItemsInner2) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *ItemsInner2) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *ItemsInner2) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.


### GetGrafanaCloud

`func (o *ItemsInner2) GetGrafanaCloud() float32`

GetGrafanaCloud returns the GrafanaCloud field if non-nil, zero value otherwise.

### GetGrafanaCloudOk

`func (o *ItemsInner2) GetGrafanaCloudOk() (*float32, bool)`

GetGrafanaCloudOk returns a tuple with the GrafanaCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaCloud

`func (o *ItemsInner2) SetGrafanaCloud(v float32)`

SetGrafanaCloud sets GrafanaCloud field to given value.


### GetResellerId

`func (o *ItemsInner2) GetResellerId() float32`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *ItemsInner2) GetResellerIdOk() (*float32, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *ItemsInner2) SetResellerId(v float32)`

SetResellerId sets ResellerId field to given value.


### SetResellerIdNil

`func (o *ItemsInner2) SetResellerIdNil(b bool)`

 SetResellerIdNil sets the value for ResellerId to be an explicit nil

### UnsetResellerId
`func (o *ItemsInner2) UnsetResellerId()`

UnsetResellerId ensures that no value is present for ResellerId, not even an explicit nil
### GetContractTypeId

`func (o *ItemsInner2) GetContractTypeId() float32`

GetContractTypeId returns the ContractTypeId field if non-nil, zero value otherwise.

### GetContractTypeIdOk

`func (o *ItemsInner2) GetContractTypeIdOk() (*float32, bool)`

GetContractTypeIdOk returns a tuple with the ContractTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypeId

`func (o *ItemsInner2) SetContractTypeId(v float32)`

SetContractTypeId sets ContractTypeId field to given value.


### GetAllowGCloudTrial

`func (o *ItemsInner2) GetAllowGCloudTrial() bool`

GetAllowGCloudTrial returns the AllowGCloudTrial field if non-nil, zero value otherwise.

### GetAllowGCloudTrialOk

`func (o *ItemsInner2) GetAllowGCloudTrialOk() (*bool, bool)`

GetAllowGCloudTrialOk returns a tuple with the AllowGCloudTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGCloudTrial

`func (o *ItemsInner2) SetAllowGCloudTrial(v bool)`

SetAllowGCloudTrial sets AllowGCloudTrial field to given value.


### GetHlUsage

`func (o *ItemsInner2) GetHlUsage() float32`

GetHlUsage returns the HlUsage field if non-nil, zero value otherwise.

### GetHlUsageOk

`func (o *ItemsInner2) GetHlUsageOk() (*float32, bool)`

GetHlUsageOk returns a tuple with the HlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlUsage

`func (o *ItemsInner2) SetHlUsage(v float32)`

SetHlUsage sets HlUsage field to given value.


### GetHmCurrentGraphiteUsage

`func (o *ItemsInner2) GetHmCurrentGraphiteUsage() float32`

GetHmCurrentGraphiteUsage returns the HmCurrentGraphiteUsage field if non-nil, zero value otherwise.

### GetHmCurrentGraphiteUsageOk

`func (o *ItemsInner2) GetHmCurrentGraphiteUsageOk() (*float32, bool)`

GetHmCurrentGraphiteUsageOk returns a tuple with the HmCurrentGraphiteUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentGraphiteUsage

`func (o *ItemsInner2) SetHmCurrentGraphiteUsage(v float32)`

SetHmCurrentGraphiteUsage sets HmCurrentGraphiteUsage field to given value.


### GetHmCurrentPrometheusUsage

`func (o *ItemsInner2) GetHmCurrentPrometheusUsage() float32`

GetHmCurrentPrometheusUsage returns the HmCurrentPrometheusUsage field if non-nil, zero value otherwise.

### GetHmCurrentPrometheusUsageOk

`func (o *ItemsInner2) GetHmCurrentPrometheusUsageOk() (*float32, bool)`

GetHmCurrentPrometheusUsageOk returns a tuple with the HmCurrentPrometheusUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentPrometheusUsage

`func (o *ItemsInner2) SetHmCurrentPrometheusUsage(v float32)`

SetHmCurrentPrometheusUsage sets HmCurrentPrometheusUsage field to given value.


### GetHgDatasourceCnts

`func (o *ItemsInner2) GetHgDatasourceCnts() string`

GetHgDatasourceCnts returns the HgDatasourceCnts field if non-nil, zero value otherwise.

### GetHgDatasourceCntsOk

`func (o *ItemsInner2) GetHgDatasourceCntsOk() (*string, bool)`

GetHgDatasourceCntsOk returns a tuple with the HgDatasourceCnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgDatasourceCnts

`func (o *ItemsInner2) SetHgDatasourceCnts(v string)`

SetHgDatasourceCnts sets HgDatasourceCnts field to given value.


### GetUserFirstName

`func (o *ItemsInner2) GetUserFirstName() string`

GetUserFirstName returns the UserFirstName field if non-nil, zero value otherwise.

### GetUserFirstNameOk

`func (o *ItemsInner2) GetUserFirstNameOk() (*string, bool)`

GetUserFirstNameOk returns a tuple with the UserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFirstName

`func (o *ItemsInner2) SetUserFirstName(v string)`

SetUserFirstName sets UserFirstName field to given value.


### GetUserLastName

`func (o *ItemsInner2) GetUserLastName() string`

GetUserLastName returns the UserLastName field if non-nil, zero value otherwise.

### GetUserLastNameOk

`func (o *ItemsInner2) GetUserLastNameOk() (*string, bool)`

GetUserLastNameOk returns a tuple with the UserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLastName

`func (o *ItemsInner2) SetUserLastName(v string)`

SetUserLastName sets UserLastName field to given value.


### GetUserUsername

`func (o *ItemsInner2) GetUserUsername() string`

GetUserUsername returns the UserUsername field if non-nil, zero value otherwise.

### GetUserUsernameOk

`func (o *ItemsInner2) GetUserUsernameOk() (*string, bool)`

GetUserUsernameOk returns a tuple with the UserUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUsername

`func (o *ItemsInner2) SetUserUsername(v string)`

SetUserUsername sets UserUsername field to given value.


### GetUserStatus

`func (o *ItemsInner2) GetUserStatus() float32`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *ItemsInner2) GetUserStatusOk() (*float32, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *ItemsInner2) SetUserStatus(v float32)`

SetUserStatus sets UserStatus field to given value.


### GetUserEmail

`func (o *ItemsInner2) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ItemsInner2) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ItemsInner2) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserName

`func (o *ItemsInner2) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ItemsInner2) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ItemsInner2) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSubscriptions

`func (o *ItemsInner2) GetSubscriptions() Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ItemsInner2) GetSubscriptionsOk() (*Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ItemsInner2) SetSubscriptions(v Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.


### GetMarketplaceSubscription

`func (o *ItemsInner2) GetMarketplaceSubscription() ItemsInner2MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *ItemsInner2) GetMarketplaceSubscriptionOk() (*ItemsInner2MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *ItemsInner2) SetMarketplaceSubscription(v ItemsInner2MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.


### SetMarketplaceSubscriptionNil

`func (o *ItemsInner2) SetMarketplaceSubscriptionNil(b bool)`

 SetMarketplaceSubscriptionNil sets the value for MarketplaceSubscription to be an explicit nil

### UnsetMarketplaceSubscription
`func (o *ItemsInner2) UnsetMarketplaceSubscription()`

UnsetMarketplaceSubscription ensures that no value is present for MarketplaceSubscription, not even an explicit nil
### GetExtraPermissions

`func (o *ItemsInner2) GetExtraPermissions() []string`

GetExtraPermissions returns the ExtraPermissions field if non-nil, zero value otherwise.

### GetExtraPermissionsOk

`func (o *ItemsInner2) GetExtraPermissionsOk() (*[]string, bool)`

GetExtraPermissionsOk returns a tuple with the ExtraPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPermissions

`func (o *ItemsInner2) SetExtraPermissions(v []string)`

SetExtraPermissions sets ExtraPermissions field to given value.

### HasExtraPermissions

`func (o *ItemsInner2) HasExtraPermissions() bool`

HasExtraPermissions returns a boolean if a field has been set.

### GetGrafanaStaffAccess

`func (o *ItemsInner2) GetGrafanaStaffAccess() ItemsInner2GrafanaStaffAccess`

GetGrafanaStaffAccess returns the GrafanaStaffAccess field if non-nil, zero value otherwise.

### GetGrafanaStaffAccessOk

`func (o *ItemsInner2) GetGrafanaStaffAccessOk() (*ItemsInner2GrafanaStaffAccess, bool)`

GetGrafanaStaffAccessOk returns a tuple with the GrafanaStaffAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaStaffAccess

`func (o *ItemsInner2) SetGrafanaStaffAccess(v ItemsInner2GrafanaStaffAccess)`

SetGrafanaStaffAccess sets GrafanaStaffAccess field to given value.

### HasGrafanaStaffAccess

`func (o *ItemsInner2) HasGrafanaStaffAccess() bool`

HasGrafanaStaffAccess returns a boolean if a field has been set.

### SetGrafanaStaffAccessNil

`func (o *ItemsInner2) SetGrafanaStaffAccessNil(b bool)`

 SetGrafanaStaffAccessNil sets the value for GrafanaStaffAccess to be an explicit nil

### UnsetGrafanaStaffAccess
`func (o *ItemsInner2) UnsetGrafanaStaffAccess()`

UnsetGrafanaStaffAccess ensures that no value is present for GrafanaStaffAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


