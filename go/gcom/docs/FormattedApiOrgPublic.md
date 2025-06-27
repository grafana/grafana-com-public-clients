# FormattedApiOrgPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Url** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**UpdatedBy** | **NullableString** |  | 
**Avatar** | **NullableString** |  | 
**ChecksPerMonth** | **float32** |  | 
**WpPlan** | **string** |  | 
**HgInstanceLimit** | **float32** |  | 
**HmInstanceLimit** | **float32** |  | 
**HlInstanceLimit** | **float32** |  | 
**UserQuota** | **float32** |  | 
**SupportPlan** | **string** |  | 
**CreditApproved** | **float32** |  | 
**MsaSignedAt** | **NullableString** |  | 
**MsaSignedBy** | **NullableString** |  | 
**EnterprisePlugins** | **float32** |  | 
**LicenseProducts** | **[]string** |  | 
**GrafanaCloud** | **float32** |  | 
**Privacy** | **string** |  | 
**Reseller** | **string** |  | 
**ResellerId** | **NullableFloat32** |  | 
**ResellerName** | **NullableString** |  | 
**EmergencySupport** | **bool** |  | 
**GcloudMonthlyCost** | **float32** |  | 
**HgUsage** | **float32** |  | 
**HgCurrentActiveUsers** | **float32** |  | 
**HgGrafanaUsage** | **float32** |  | 
**HgOnCallUsage** | **float32** |  | 
**HmUsage** | **float32** |  | 
**HmCurrentUsage** | **float32** |  | 
**HmGraphiteUsage** | **float32** |  | 
**HlUsage** | **float32** |  | 
**HlRetentionUsage** | **float32** |  | 
**HtUsage** | **float32** |  | 
**HpUsage** | **float32** |  | 
**IrmUsage** | **float32** |  | 
**K6VuhUsage** | **float32** |  | 
**K6IPUsage** | **float32** |  | 
**FeO11YUsage** | **float32** |  | 
**AppO11YUsage** | **float32** |  | 
**SmUsage** | **float32** |  | 
**InfraO11YHostsUsage** | **float32** |  | 
**InfraO11YContainersUsage** | **float32** |  | 
**GeUsersUsage** | **float32** |  | 
**GeInstancesUsage** | **float32** |  | 
**SmBrowserUsage** | **float32** |  | 
**AwsMarketplaceSupport** | **float32** |  | 
**TrialStartDate** | **NullableString** |  | 
**TrialEndDate** | **NullableString** |  | 
**TrialLengthDays** | **NullableFloat32** |  | 
**TrialNoticeDate** | **NullableString** |  | 
**CancellationDate** | **NullableString** |  | 
**RetainedStackId** | **float32** |  | 
**AllowGCloudTrial** | [**FormattedOrgMembershipAllowGCloudTrial**](FormattedOrgMembershipAllowGCloudTrial.md) |  | 
**PluginSignatureType** | **string** |  | 
**ContractType** | **string** |  | 
**ContractTypeId** | **float32** |  | 
**LiveChatEnabled** | **bool** |  | 
**DisableTokenExpirationEmails** | **bool** |  | 
**MaxTokenExpirationDays** | **float32** |  | 
**Links** | [**[]LinksInner1**](LinksInner1.md) |  | 
**Subscriptions** | [**Subscriptions**](Subscriptions.md) |  | 

## Methods

### NewFormattedApiOrgPublic

`func NewFormattedApiOrgPublic(id float32, slug string, name string, url string, createdAt string, createdBy NullableString, updatedAt NullableString, updatedBy NullableString, avatar NullableString, checksPerMonth float32, wpPlan string, hgInstanceLimit float32, hmInstanceLimit float32, hlInstanceLimit float32, userQuota float32, supportPlan string, creditApproved float32, msaSignedAt NullableString, msaSignedBy NullableString, enterprisePlugins float32, licenseProducts []string, grafanaCloud float32, privacy string, reseller string, resellerId NullableFloat32, resellerName NullableString, emergencySupport bool, gcloudMonthlyCost float32, hgUsage float32, hgCurrentActiveUsers float32, hgGrafanaUsage float32, hgOnCallUsage float32, hmUsage float32, hmCurrentUsage float32, hmGraphiteUsage float32, hlUsage float32, hlRetentionUsage float32, htUsage float32, hpUsage float32, irmUsage float32, k6VuhUsage float32, k6IPUsage float32, feO11YUsage float32, appO11YUsage float32, smUsage float32, infraO11YHostsUsage float32, infraO11YContainersUsage float32, geUsersUsage float32, geInstancesUsage float32, smBrowserUsage float32, awsMarketplaceSupport float32, trialStartDate NullableString, trialEndDate NullableString, trialLengthDays NullableFloat32, trialNoticeDate NullableString, cancellationDate NullableString, retainedStackId float32, allowGCloudTrial FormattedOrgMembershipAllowGCloudTrial, pluginSignatureType string, contractType string, contractTypeId float32, liveChatEnabled bool, disableTokenExpirationEmails bool, maxTokenExpirationDays float32, links []LinksInner1, subscriptions Subscriptions, ) *FormattedApiOrgPublic`

NewFormattedApiOrgPublic instantiates a new FormattedApiOrgPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiOrgPublicWithDefaults

`func NewFormattedApiOrgPublicWithDefaults() *FormattedApiOrgPublic`

NewFormattedApiOrgPublicWithDefaults instantiates a new FormattedApiOrgPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormattedApiOrgPublic) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiOrgPublic) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiOrgPublic) SetId(v float32)`

SetId sets Id field to given value.


### GetSlug

`func (o *FormattedApiOrgPublic) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormattedApiOrgPublic) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormattedApiOrgPublic) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *FormattedApiOrgPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiOrgPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiOrgPublic) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *FormattedApiOrgPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormattedApiOrgPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormattedApiOrgPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *FormattedApiOrgPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiOrgPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiOrgPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *FormattedApiOrgPublic) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FormattedApiOrgPublic) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FormattedApiOrgPublic) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *FormattedApiOrgPublic) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *FormattedApiOrgPublic) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedAt

`func (o *FormattedApiOrgPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiOrgPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiOrgPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FormattedApiOrgPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiOrgPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *FormattedApiOrgPublic) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FormattedApiOrgPublic) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FormattedApiOrgPublic) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### SetUpdatedByNil

`func (o *FormattedApiOrgPublic) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *FormattedApiOrgPublic) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetAvatar

`func (o *FormattedApiOrgPublic) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *FormattedApiOrgPublic) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *FormattedApiOrgPublic) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *FormattedApiOrgPublic) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *FormattedApiOrgPublic) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetChecksPerMonth

`func (o *FormattedApiOrgPublic) GetChecksPerMonth() float32`

GetChecksPerMonth returns the ChecksPerMonth field if non-nil, zero value otherwise.

### GetChecksPerMonthOk

`func (o *FormattedApiOrgPublic) GetChecksPerMonthOk() (*float32, bool)`

GetChecksPerMonthOk returns a tuple with the ChecksPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksPerMonth

`func (o *FormattedApiOrgPublic) SetChecksPerMonth(v float32)`

SetChecksPerMonth sets ChecksPerMonth field to given value.


### GetWpPlan

`func (o *FormattedApiOrgPublic) GetWpPlan() string`

GetWpPlan returns the WpPlan field if non-nil, zero value otherwise.

### GetWpPlanOk

`func (o *FormattedApiOrgPublic) GetWpPlanOk() (*string, bool)`

GetWpPlanOk returns a tuple with the WpPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpPlan

`func (o *FormattedApiOrgPublic) SetWpPlan(v string)`

SetWpPlan sets WpPlan field to given value.


### GetHgInstanceLimit

`func (o *FormattedApiOrgPublic) GetHgInstanceLimit() float32`

GetHgInstanceLimit returns the HgInstanceLimit field if non-nil, zero value otherwise.

### GetHgInstanceLimitOk

`func (o *FormattedApiOrgPublic) GetHgInstanceLimitOk() (*float32, bool)`

GetHgInstanceLimitOk returns a tuple with the HgInstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgInstanceLimit

`func (o *FormattedApiOrgPublic) SetHgInstanceLimit(v float32)`

SetHgInstanceLimit sets HgInstanceLimit field to given value.


### GetHmInstanceLimit

`func (o *FormattedApiOrgPublic) GetHmInstanceLimit() float32`

GetHmInstanceLimit returns the HmInstanceLimit field if non-nil, zero value otherwise.

### GetHmInstanceLimitOk

`func (o *FormattedApiOrgPublic) GetHmInstanceLimitOk() (*float32, bool)`

GetHmInstanceLimitOk returns a tuple with the HmInstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceLimit

`func (o *FormattedApiOrgPublic) SetHmInstanceLimit(v float32)`

SetHmInstanceLimit sets HmInstanceLimit field to given value.


### GetHlInstanceLimit

`func (o *FormattedApiOrgPublic) GetHlInstanceLimit() float32`

GetHlInstanceLimit returns the HlInstanceLimit field if non-nil, zero value otherwise.

### GetHlInstanceLimitOk

`func (o *FormattedApiOrgPublic) GetHlInstanceLimitOk() (*float32, bool)`

GetHlInstanceLimitOk returns a tuple with the HlInstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceLimit

`func (o *FormattedApiOrgPublic) SetHlInstanceLimit(v float32)`

SetHlInstanceLimit sets HlInstanceLimit field to given value.


### GetUserQuota

`func (o *FormattedApiOrgPublic) GetUserQuota() float32`

GetUserQuota returns the UserQuota field if non-nil, zero value otherwise.

### GetUserQuotaOk

`func (o *FormattedApiOrgPublic) GetUserQuotaOk() (*float32, bool)`

GetUserQuotaOk returns a tuple with the UserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuota

`func (o *FormattedApiOrgPublic) SetUserQuota(v float32)`

SetUserQuota sets UserQuota field to given value.


### GetSupportPlan

`func (o *FormattedApiOrgPublic) GetSupportPlan() string`

GetSupportPlan returns the SupportPlan field if non-nil, zero value otherwise.

### GetSupportPlanOk

`func (o *FormattedApiOrgPublic) GetSupportPlanOk() (*string, bool)`

GetSupportPlanOk returns a tuple with the SupportPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPlan

`func (o *FormattedApiOrgPublic) SetSupportPlan(v string)`

SetSupportPlan sets SupportPlan field to given value.


### GetCreditApproved

`func (o *FormattedApiOrgPublic) GetCreditApproved() float32`

GetCreditApproved returns the CreditApproved field if non-nil, zero value otherwise.

### GetCreditApprovedOk

`func (o *FormattedApiOrgPublic) GetCreditApprovedOk() (*float32, bool)`

GetCreditApprovedOk returns a tuple with the CreditApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditApproved

`func (o *FormattedApiOrgPublic) SetCreditApproved(v float32)`

SetCreditApproved sets CreditApproved field to given value.


### GetMsaSignedAt

`func (o *FormattedApiOrgPublic) GetMsaSignedAt() string`

GetMsaSignedAt returns the MsaSignedAt field if non-nil, zero value otherwise.

### GetMsaSignedAtOk

`func (o *FormattedApiOrgPublic) GetMsaSignedAtOk() (*string, bool)`

GetMsaSignedAtOk returns a tuple with the MsaSignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaSignedAt

`func (o *FormattedApiOrgPublic) SetMsaSignedAt(v string)`

SetMsaSignedAt sets MsaSignedAt field to given value.


### SetMsaSignedAtNil

`func (o *FormattedApiOrgPublic) SetMsaSignedAtNil(b bool)`

 SetMsaSignedAtNil sets the value for MsaSignedAt to be an explicit nil

### UnsetMsaSignedAt
`func (o *FormattedApiOrgPublic) UnsetMsaSignedAt()`

UnsetMsaSignedAt ensures that no value is present for MsaSignedAt, not even an explicit nil
### GetMsaSignedBy

`func (o *FormattedApiOrgPublic) GetMsaSignedBy() string`

GetMsaSignedBy returns the MsaSignedBy field if non-nil, zero value otherwise.

### GetMsaSignedByOk

`func (o *FormattedApiOrgPublic) GetMsaSignedByOk() (*string, bool)`

GetMsaSignedByOk returns a tuple with the MsaSignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaSignedBy

`func (o *FormattedApiOrgPublic) SetMsaSignedBy(v string)`

SetMsaSignedBy sets MsaSignedBy field to given value.


### SetMsaSignedByNil

`func (o *FormattedApiOrgPublic) SetMsaSignedByNil(b bool)`

 SetMsaSignedByNil sets the value for MsaSignedBy to be an explicit nil

### UnsetMsaSignedBy
`func (o *FormattedApiOrgPublic) UnsetMsaSignedBy()`

UnsetMsaSignedBy ensures that no value is present for MsaSignedBy, not even an explicit nil
### GetEnterprisePlugins

`func (o *FormattedApiOrgPublic) GetEnterprisePlugins() float32`

GetEnterprisePlugins returns the EnterprisePlugins field if non-nil, zero value otherwise.

### GetEnterprisePluginsOk

`func (o *FormattedApiOrgPublic) GetEnterprisePluginsOk() (*float32, bool)`

GetEnterprisePluginsOk returns a tuple with the EnterprisePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePlugins

`func (o *FormattedApiOrgPublic) SetEnterprisePlugins(v float32)`

SetEnterprisePlugins sets EnterprisePlugins field to given value.


### GetLicenseProducts

`func (o *FormattedApiOrgPublic) GetLicenseProducts() []string`

GetLicenseProducts returns the LicenseProducts field if non-nil, zero value otherwise.

### GetLicenseProductsOk

`func (o *FormattedApiOrgPublic) GetLicenseProductsOk() (*[]string, bool)`

GetLicenseProductsOk returns a tuple with the LicenseProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProducts

`func (o *FormattedApiOrgPublic) SetLicenseProducts(v []string)`

SetLicenseProducts sets LicenseProducts field to given value.


### GetGrafanaCloud

`func (o *FormattedApiOrgPublic) GetGrafanaCloud() float32`

GetGrafanaCloud returns the GrafanaCloud field if non-nil, zero value otherwise.

### GetGrafanaCloudOk

`func (o *FormattedApiOrgPublic) GetGrafanaCloudOk() (*float32, bool)`

GetGrafanaCloudOk returns a tuple with the GrafanaCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaCloud

`func (o *FormattedApiOrgPublic) SetGrafanaCloud(v float32)`

SetGrafanaCloud sets GrafanaCloud field to given value.


### GetPrivacy

`func (o *FormattedApiOrgPublic) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *FormattedApiOrgPublic) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *FormattedApiOrgPublic) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.


### GetReseller

`func (o *FormattedApiOrgPublic) GetReseller() string`

GetReseller returns the Reseller field if non-nil, zero value otherwise.

### GetResellerOk

`func (o *FormattedApiOrgPublic) GetResellerOk() (*string, bool)`

GetResellerOk returns a tuple with the Reseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReseller

`func (o *FormattedApiOrgPublic) SetReseller(v string)`

SetReseller sets Reseller field to given value.


### GetResellerId

`func (o *FormattedApiOrgPublic) GetResellerId() float32`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *FormattedApiOrgPublic) GetResellerIdOk() (*float32, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *FormattedApiOrgPublic) SetResellerId(v float32)`

SetResellerId sets ResellerId field to given value.


### SetResellerIdNil

`func (o *FormattedApiOrgPublic) SetResellerIdNil(b bool)`

 SetResellerIdNil sets the value for ResellerId to be an explicit nil

### UnsetResellerId
`func (o *FormattedApiOrgPublic) UnsetResellerId()`

UnsetResellerId ensures that no value is present for ResellerId, not even an explicit nil
### GetResellerName

`func (o *FormattedApiOrgPublic) GetResellerName() string`

GetResellerName returns the ResellerName field if non-nil, zero value otherwise.

### GetResellerNameOk

`func (o *FormattedApiOrgPublic) GetResellerNameOk() (*string, bool)`

GetResellerNameOk returns a tuple with the ResellerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerName

`func (o *FormattedApiOrgPublic) SetResellerName(v string)`

SetResellerName sets ResellerName field to given value.


### SetResellerNameNil

`func (o *FormattedApiOrgPublic) SetResellerNameNil(b bool)`

 SetResellerNameNil sets the value for ResellerName to be an explicit nil

### UnsetResellerName
`func (o *FormattedApiOrgPublic) UnsetResellerName()`

UnsetResellerName ensures that no value is present for ResellerName, not even an explicit nil
### GetEmergencySupport

`func (o *FormattedApiOrgPublic) GetEmergencySupport() bool`

GetEmergencySupport returns the EmergencySupport field if non-nil, zero value otherwise.

### GetEmergencySupportOk

`func (o *FormattedApiOrgPublic) GetEmergencySupportOk() (*bool, bool)`

GetEmergencySupportOk returns a tuple with the EmergencySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencySupport

`func (o *FormattedApiOrgPublic) SetEmergencySupport(v bool)`

SetEmergencySupport sets EmergencySupport field to given value.


### GetGcloudMonthlyCost

`func (o *FormattedApiOrgPublic) GetGcloudMonthlyCost() float32`

GetGcloudMonthlyCost returns the GcloudMonthlyCost field if non-nil, zero value otherwise.

### GetGcloudMonthlyCostOk

`func (o *FormattedApiOrgPublic) GetGcloudMonthlyCostOk() (*float32, bool)`

GetGcloudMonthlyCostOk returns a tuple with the GcloudMonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcloudMonthlyCost

`func (o *FormattedApiOrgPublic) SetGcloudMonthlyCost(v float32)`

SetGcloudMonthlyCost sets GcloudMonthlyCost field to given value.


### GetHgUsage

`func (o *FormattedApiOrgPublic) GetHgUsage() float32`

GetHgUsage returns the HgUsage field if non-nil, zero value otherwise.

### GetHgUsageOk

`func (o *FormattedApiOrgPublic) GetHgUsageOk() (*float32, bool)`

GetHgUsageOk returns a tuple with the HgUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgUsage

`func (o *FormattedApiOrgPublic) SetHgUsage(v float32)`

SetHgUsage sets HgUsage field to given value.


### GetHgCurrentActiveUsers

`func (o *FormattedApiOrgPublic) GetHgCurrentActiveUsers() float32`

GetHgCurrentActiveUsers returns the HgCurrentActiveUsers field if non-nil, zero value otherwise.

### GetHgCurrentActiveUsersOk

`func (o *FormattedApiOrgPublic) GetHgCurrentActiveUsersOk() (*float32, bool)`

GetHgCurrentActiveUsersOk returns a tuple with the HgCurrentActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgCurrentActiveUsers

`func (o *FormattedApiOrgPublic) SetHgCurrentActiveUsers(v float32)`

SetHgCurrentActiveUsers sets HgCurrentActiveUsers field to given value.


### GetHgGrafanaUsage

`func (o *FormattedApiOrgPublic) GetHgGrafanaUsage() float32`

GetHgGrafanaUsage returns the HgGrafanaUsage field if non-nil, zero value otherwise.

### GetHgGrafanaUsageOk

`func (o *FormattedApiOrgPublic) GetHgGrafanaUsageOk() (*float32, bool)`

GetHgGrafanaUsageOk returns a tuple with the HgGrafanaUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgGrafanaUsage

`func (o *FormattedApiOrgPublic) SetHgGrafanaUsage(v float32)`

SetHgGrafanaUsage sets HgGrafanaUsage field to given value.


### GetHgOnCallUsage

`func (o *FormattedApiOrgPublic) GetHgOnCallUsage() float32`

GetHgOnCallUsage returns the HgOnCallUsage field if non-nil, zero value otherwise.

### GetHgOnCallUsageOk

`func (o *FormattedApiOrgPublic) GetHgOnCallUsageOk() (*float32, bool)`

GetHgOnCallUsageOk returns a tuple with the HgOnCallUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgOnCallUsage

`func (o *FormattedApiOrgPublic) SetHgOnCallUsage(v float32)`

SetHgOnCallUsage sets HgOnCallUsage field to given value.


### GetHmUsage

`func (o *FormattedApiOrgPublic) GetHmUsage() float32`

GetHmUsage returns the HmUsage field if non-nil, zero value otherwise.

### GetHmUsageOk

`func (o *FormattedApiOrgPublic) GetHmUsageOk() (*float32, bool)`

GetHmUsageOk returns a tuple with the HmUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmUsage

`func (o *FormattedApiOrgPublic) SetHmUsage(v float32)`

SetHmUsage sets HmUsage field to given value.


### GetHmCurrentUsage

`func (o *FormattedApiOrgPublic) GetHmCurrentUsage() float32`

GetHmCurrentUsage returns the HmCurrentUsage field if non-nil, zero value otherwise.

### GetHmCurrentUsageOk

`func (o *FormattedApiOrgPublic) GetHmCurrentUsageOk() (*float32, bool)`

GetHmCurrentUsageOk returns a tuple with the HmCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmCurrentUsage

`func (o *FormattedApiOrgPublic) SetHmCurrentUsage(v float32)`

SetHmCurrentUsage sets HmCurrentUsage field to given value.


### GetHmGraphiteUsage

`func (o *FormattedApiOrgPublic) GetHmGraphiteUsage() float32`

GetHmGraphiteUsage returns the HmGraphiteUsage field if non-nil, zero value otherwise.

### GetHmGraphiteUsageOk

`func (o *FormattedApiOrgPublic) GetHmGraphiteUsageOk() (*float32, bool)`

GetHmGraphiteUsageOk returns a tuple with the HmGraphiteUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteUsage

`func (o *FormattedApiOrgPublic) SetHmGraphiteUsage(v float32)`

SetHmGraphiteUsage sets HmGraphiteUsage field to given value.


### GetHlUsage

`func (o *FormattedApiOrgPublic) GetHlUsage() float32`

GetHlUsage returns the HlUsage field if non-nil, zero value otherwise.

### GetHlUsageOk

`func (o *FormattedApiOrgPublic) GetHlUsageOk() (*float32, bool)`

GetHlUsageOk returns a tuple with the HlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlUsage

`func (o *FormattedApiOrgPublic) SetHlUsage(v float32)`

SetHlUsage sets HlUsage field to given value.


### GetHlRetentionUsage

`func (o *FormattedApiOrgPublic) GetHlRetentionUsage() float32`

GetHlRetentionUsage returns the HlRetentionUsage field if non-nil, zero value otherwise.

### GetHlRetentionUsageOk

`func (o *FormattedApiOrgPublic) GetHlRetentionUsageOk() (*float32, bool)`

GetHlRetentionUsageOk returns a tuple with the HlRetentionUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlRetentionUsage

`func (o *FormattedApiOrgPublic) SetHlRetentionUsage(v float32)`

SetHlRetentionUsage sets HlRetentionUsage field to given value.


### GetHtUsage

`func (o *FormattedApiOrgPublic) GetHtUsage() float32`

GetHtUsage returns the HtUsage field if non-nil, zero value otherwise.

### GetHtUsageOk

`func (o *FormattedApiOrgPublic) GetHtUsageOk() (*float32, bool)`

GetHtUsageOk returns a tuple with the HtUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtUsage

`func (o *FormattedApiOrgPublic) SetHtUsage(v float32)`

SetHtUsage sets HtUsage field to given value.


### GetHpUsage

`func (o *FormattedApiOrgPublic) GetHpUsage() float32`

GetHpUsage returns the HpUsage field if non-nil, zero value otherwise.

### GetHpUsageOk

`func (o *FormattedApiOrgPublic) GetHpUsageOk() (*float32, bool)`

GetHpUsageOk returns a tuple with the HpUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpUsage

`func (o *FormattedApiOrgPublic) SetHpUsage(v float32)`

SetHpUsage sets HpUsage field to given value.


### GetIrmUsage

`func (o *FormattedApiOrgPublic) GetIrmUsage() float32`

GetIrmUsage returns the IrmUsage field if non-nil, zero value otherwise.

### GetIrmUsageOk

`func (o *FormattedApiOrgPublic) GetIrmUsageOk() (*float32, bool)`

GetIrmUsageOk returns a tuple with the IrmUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrmUsage

`func (o *FormattedApiOrgPublic) SetIrmUsage(v float32)`

SetIrmUsage sets IrmUsage field to given value.


### GetK6VuhUsage

`func (o *FormattedApiOrgPublic) GetK6VuhUsage() float32`

GetK6VuhUsage returns the K6VuhUsage field if non-nil, zero value otherwise.

### GetK6VuhUsageOk

`func (o *FormattedApiOrgPublic) GetK6VuhUsageOk() (*float32, bool)`

GetK6VuhUsageOk returns a tuple with the K6VuhUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6VuhUsage

`func (o *FormattedApiOrgPublic) SetK6VuhUsage(v float32)`

SetK6VuhUsage sets K6VuhUsage field to given value.


### GetK6IPUsage

`func (o *FormattedApiOrgPublic) GetK6IPUsage() float32`

GetK6IPUsage returns the K6IPUsage field if non-nil, zero value otherwise.

### GetK6IPUsageOk

`func (o *FormattedApiOrgPublic) GetK6IPUsageOk() (*float32, bool)`

GetK6IPUsageOk returns a tuple with the K6IPUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6IPUsage

`func (o *FormattedApiOrgPublic) SetK6IPUsage(v float32)`

SetK6IPUsage sets K6IPUsage field to given value.


### GetFeO11YUsage

`func (o *FormattedApiOrgPublic) GetFeO11YUsage() float32`

GetFeO11YUsage returns the FeO11YUsage field if non-nil, zero value otherwise.

### GetFeO11YUsageOk

`func (o *FormattedApiOrgPublic) GetFeO11YUsageOk() (*float32, bool)`

GetFeO11YUsageOk returns a tuple with the FeO11YUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeO11YUsage

`func (o *FormattedApiOrgPublic) SetFeO11YUsage(v float32)`

SetFeO11YUsage sets FeO11YUsage field to given value.


### GetAppO11YUsage

`func (o *FormattedApiOrgPublic) GetAppO11YUsage() float32`

GetAppO11YUsage returns the AppO11YUsage field if non-nil, zero value otherwise.

### GetAppO11YUsageOk

`func (o *FormattedApiOrgPublic) GetAppO11YUsageOk() (*float32, bool)`

GetAppO11YUsageOk returns a tuple with the AppO11YUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppO11YUsage

`func (o *FormattedApiOrgPublic) SetAppO11YUsage(v float32)`

SetAppO11YUsage sets AppO11YUsage field to given value.


### GetSmUsage

`func (o *FormattedApiOrgPublic) GetSmUsage() float32`

GetSmUsage returns the SmUsage field if non-nil, zero value otherwise.

### GetSmUsageOk

`func (o *FormattedApiOrgPublic) GetSmUsageOk() (*float32, bool)`

GetSmUsageOk returns a tuple with the SmUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmUsage

`func (o *FormattedApiOrgPublic) SetSmUsage(v float32)`

SetSmUsage sets SmUsage field to given value.


### GetInfraO11YHostsUsage

`func (o *FormattedApiOrgPublic) GetInfraO11YHostsUsage() float32`

GetInfraO11YHostsUsage returns the InfraO11YHostsUsage field if non-nil, zero value otherwise.

### GetInfraO11YHostsUsageOk

`func (o *FormattedApiOrgPublic) GetInfraO11YHostsUsageOk() (*float32, bool)`

GetInfraO11YHostsUsageOk returns a tuple with the InfraO11YHostsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraO11YHostsUsage

`func (o *FormattedApiOrgPublic) SetInfraO11YHostsUsage(v float32)`

SetInfraO11YHostsUsage sets InfraO11YHostsUsage field to given value.


### GetInfraO11YContainersUsage

`func (o *FormattedApiOrgPublic) GetInfraO11YContainersUsage() float32`

GetInfraO11YContainersUsage returns the InfraO11YContainersUsage field if non-nil, zero value otherwise.

### GetInfraO11YContainersUsageOk

`func (o *FormattedApiOrgPublic) GetInfraO11YContainersUsageOk() (*float32, bool)`

GetInfraO11YContainersUsageOk returns a tuple with the InfraO11YContainersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraO11YContainersUsage

`func (o *FormattedApiOrgPublic) SetInfraO11YContainersUsage(v float32)`

SetInfraO11YContainersUsage sets InfraO11YContainersUsage field to given value.


### GetGeUsersUsage

`func (o *FormattedApiOrgPublic) GetGeUsersUsage() float32`

GetGeUsersUsage returns the GeUsersUsage field if non-nil, zero value otherwise.

### GetGeUsersUsageOk

`func (o *FormattedApiOrgPublic) GetGeUsersUsageOk() (*float32, bool)`

GetGeUsersUsageOk returns a tuple with the GeUsersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeUsersUsage

`func (o *FormattedApiOrgPublic) SetGeUsersUsage(v float32)`

SetGeUsersUsage sets GeUsersUsage field to given value.


### GetGeInstancesUsage

`func (o *FormattedApiOrgPublic) GetGeInstancesUsage() float32`

GetGeInstancesUsage returns the GeInstancesUsage field if non-nil, zero value otherwise.

### GetGeInstancesUsageOk

`func (o *FormattedApiOrgPublic) GetGeInstancesUsageOk() (*float32, bool)`

GetGeInstancesUsageOk returns a tuple with the GeInstancesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeInstancesUsage

`func (o *FormattedApiOrgPublic) SetGeInstancesUsage(v float32)`

SetGeInstancesUsage sets GeInstancesUsage field to given value.


### GetSmBrowserUsage

`func (o *FormattedApiOrgPublic) GetSmBrowserUsage() float32`

GetSmBrowserUsage returns the SmBrowserUsage field if non-nil, zero value otherwise.

### GetSmBrowserUsageOk

`func (o *FormattedApiOrgPublic) GetSmBrowserUsageOk() (*float32, bool)`

GetSmBrowserUsageOk returns a tuple with the SmBrowserUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmBrowserUsage

`func (o *FormattedApiOrgPublic) SetSmBrowserUsage(v float32)`

SetSmBrowserUsage sets SmBrowserUsage field to given value.


### GetAwsMarketplaceSupport

`func (o *FormattedApiOrgPublic) GetAwsMarketplaceSupport() float32`

GetAwsMarketplaceSupport returns the AwsMarketplaceSupport field if non-nil, zero value otherwise.

### GetAwsMarketplaceSupportOk

`func (o *FormattedApiOrgPublic) GetAwsMarketplaceSupportOk() (*float32, bool)`

GetAwsMarketplaceSupportOk returns a tuple with the AwsMarketplaceSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarketplaceSupport

`func (o *FormattedApiOrgPublic) SetAwsMarketplaceSupport(v float32)`

SetAwsMarketplaceSupport sets AwsMarketplaceSupport field to given value.


### GetTrialStartDate

`func (o *FormattedApiOrgPublic) GetTrialStartDate() string`

GetTrialStartDate returns the TrialStartDate field if non-nil, zero value otherwise.

### GetTrialStartDateOk

`func (o *FormattedApiOrgPublic) GetTrialStartDateOk() (*string, bool)`

GetTrialStartDateOk returns a tuple with the TrialStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartDate

`func (o *FormattedApiOrgPublic) SetTrialStartDate(v string)`

SetTrialStartDate sets TrialStartDate field to given value.


### SetTrialStartDateNil

`func (o *FormattedApiOrgPublic) SetTrialStartDateNil(b bool)`

 SetTrialStartDateNil sets the value for TrialStartDate to be an explicit nil

### UnsetTrialStartDate
`func (o *FormattedApiOrgPublic) UnsetTrialStartDate()`

UnsetTrialStartDate ensures that no value is present for TrialStartDate, not even an explicit nil
### GetTrialEndDate

`func (o *FormattedApiOrgPublic) GetTrialEndDate() string`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *FormattedApiOrgPublic) GetTrialEndDateOk() (*string, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *FormattedApiOrgPublic) SetTrialEndDate(v string)`

SetTrialEndDate sets TrialEndDate field to given value.


### SetTrialEndDateNil

`func (o *FormattedApiOrgPublic) SetTrialEndDateNil(b bool)`

 SetTrialEndDateNil sets the value for TrialEndDate to be an explicit nil

### UnsetTrialEndDate
`func (o *FormattedApiOrgPublic) UnsetTrialEndDate()`

UnsetTrialEndDate ensures that no value is present for TrialEndDate, not even an explicit nil
### GetTrialLengthDays

`func (o *FormattedApiOrgPublic) GetTrialLengthDays() float32`

GetTrialLengthDays returns the TrialLengthDays field if non-nil, zero value otherwise.

### GetTrialLengthDaysOk

`func (o *FormattedApiOrgPublic) GetTrialLengthDaysOk() (*float32, bool)`

GetTrialLengthDaysOk returns a tuple with the TrialLengthDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialLengthDays

`func (o *FormattedApiOrgPublic) SetTrialLengthDays(v float32)`

SetTrialLengthDays sets TrialLengthDays field to given value.


### SetTrialLengthDaysNil

`func (o *FormattedApiOrgPublic) SetTrialLengthDaysNil(b bool)`

 SetTrialLengthDaysNil sets the value for TrialLengthDays to be an explicit nil

### UnsetTrialLengthDays
`func (o *FormattedApiOrgPublic) UnsetTrialLengthDays()`

UnsetTrialLengthDays ensures that no value is present for TrialLengthDays, not even an explicit nil
### GetTrialNoticeDate

`func (o *FormattedApiOrgPublic) GetTrialNoticeDate() string`

GetTrialNoticeDate returns the TrialNoticeDate field if non-nil, zero value otherwise.

### GetTrialNoticeDateOk

`func (o *FormattedApiOrgPublic) GetTrialNoticeDateOk() (*string, bool)`

GetTrialNoticeDateOk returns a tuple with the TrialNoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialNoticeDate

`func (o *FormattedApiOrgPublic) SetTrialNoticeDate(v string)`

SetTrialNoticeDate sets TrialNoticeDate field to given value.


### SetTrialNoticeDateNil

`func (o *FormattedApiOrgPublic) SetTrialNoticeDateNil(b bool)`

 SetTrialNoticeDateNil sets the value for TrialNoticeDate to be an explicit nil

### UnsetTrialNoticeDate
`func (o *FormattedApiOrgPublic) UnsetTrialNoticeDate()`

UnsetTrialNoticeDate ensures that no value is present for TrialNoticeDate, not even an explicit nil
### GetCancellationDate

`func (o *FormattedApiOrgPublic) GetCancellationDate() string`

GetCancellationDate returns the CancellationDate field if non-nil, zero value otherwise.

### GetCancellationDateOk

`func (o *FormattedApiOrgPublic) GetCancellationDateOk() (*string, bool)`

GetCancellationDateOk returns a tuple with the CancellationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationDate

`func (o *FormattedApiOrgPublic) SetCancellationDate(v string)`

SetCancellationDate sets CancellationDate field to given value.


### SetCancellationDateNil

`func (o *FormattedApiOrgPublic) SetCancellationDateNil(b bool)`

 SetCancellationDateNil sets the value for CancellationDate to be an explicit nil

### UnsetCancellationDate
`func (o *FormattedApiOrgPublic) UnsetCancellationDate()`

UnsetCancellationDate ensures that no value is present for CancellationDate, not even an explicit nil
### GetRetainedStackId

`func (o *FormattedApiOrgPublic) GetRetainedStackId() float32`

GetRetainedStackId returns the RetainedStackId field if non-nil, zero value otherwise.

### GetRetainedStackIdOk

`func (o *FormattedApiOrgPublic) GetRetainedStackIdOk() (*float32, bool)`

GetRetainedStackIdOk returns a tuple with the RetainedStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedStackId

`func (o *FormattedApiOrgPublic) SetRetainedStackId(v float32)`

SetRetainedStackId sets RetainedStackId field to given value.


### GetAllowGCloudTrial

`func (o *FormattedApiOrgPublic) GetAllowGCloudTrial() FormattedOrgMembershipAllowGCloudTrial`

GetAllowGCloudTrial returns the AllowGCloudTrial field if non-nil, zero value otherwise.

### GetAllowGCloudTrialOk

`func (o *FormattedApiOrgPublic) GetAllowGCloudTrialOk() (*FormattedOrgMembershipAllowGCloudTrial, bool)`

GetAllowGCloudTrialOk returns a tuple with the AllowGCloudTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGCloudTrial

`func (o *FormattedApiOrgPublic) SetAllowGCloudTrial(v FormattedOrgMembershipAllowGCloudTrial)`

SetAllowGCloudTrial sets AllowGCloudTrial field to given value.


### GetPluginSignatureType

`func (o *FormattedApiOrgPublic) GetPluginSignatureType() string`

GetPluginSignatureType returns the PluginSignatureType field if non-nil, zero value otherwise.

### GetPluginSignatureTypeOk

`func (o *FormattedApiOrgPublic) GetPluginSignatureTypeOk() (*string, bool)`

GetPluginSignatureTypeOk returns a tuple with the PluginSignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginSignatureType

`func (o *FormattedApiOrgPublic) SetPluginSignatureType(v string)`

SetPluginSignatureType sets PluginSignatureType field to given value.


### GetContractType

`func (o *FormattedApiOrgPublic) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *FormattedApiOrgPublic) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *FormattedApiOrgPublic) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetContractTypeId

`func (o *FormattedApiOrgPublic) GetContractTypeId() float32`

GetContractTypeId returns the ContractTypeId field if non-nil, zero value otherwise.

### GetContractTypeIdOk

`func (o *FormattedApiOrgPublic) GetContractTypeIdOk() (*float32, bool)`

GetContractTypeIdOk returns a tuple with the ContractTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypeId

`func (o *FormattedApiOrgPublic) SetContractTypeId(v float32)`

SetContractTypeId sets ContractTypeId field to given value.


### GetLiveChatEnabled

`func (o *FormattedApiOrgPublic) GetLiveChatEnabled() bool`

GetLiveChatEnabled returns the LiveChatEnabled field if non-nil, zero value otherwise.

### GetLiveChatEnabledOk

`func (o *FormattedApiOrgPublic) GetLiveChatEnabledOk() (*bool, bool)`

GetLiveChatEnabledOk returns a tuple with the LiveChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveChatEnabled

`func (o *FormattedApiOrgPublic) SetLiveChatEnabled(v bool)`

SetLiveChatEnabled sets LiveChatEnabled field to given value.


### GetDisableTokenExpirationEmails

`func (o *FormattedApiOrgPublic) GetDisableTokenExpirationEmails() bool`

GetDisableTokenExpirationEmails returns the DisableTokenExpirationEmails field if non-nil, zero value otherwise.

### GetDisableTokenExpirationEmailsOk

`func (o *FormattedApiOrgPublic) GetDisableTokenExpirationEmailsOk() (*bool, bool)`

GetDisableTokenExpirationEmailsOk returns a tuple with the DisableTokenExpirationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTokenExpirationEmails

`func (o *FormattedApiOrgPublic) SetDisableTokenExpirationEmails(v bool)`

SetDisableTokenExpirationEmails sets DisableTokenExpirationEmails field to given value.


### GetMaxTokenExpirationDays

`func (o *FormattedApiOrgPublic) GetMaxTokenExpirationDays() float32`

GetMaxTokenExpirationDays returns the MaxTokenExpirationDays field if non-nil, zero value otherwise.

### GetMaxTokenExpirationDaysOk

`func (o *FormattedApiOrgPublic) GetMaxTokenExpirationDaysOk() (*float32, bool)`

GetMaxTokenExpirationDaysOk returns a tuple with the MaxTokenExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenExpirationDays

`func (o *FormattedApiOrgPublic) SetMaxTokenExpirationDays(v float32)`

SetMaxTokenExpirationDays sets MaxTokenExpirationDays field to given value.


### GetLinks

`func (o *FormattedApiOrgPublic) GetLinks() []LinksInner1`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiOrgPublic) GetLinksOk() (*[]LinksInner1, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiOrgPublic) SetLinks(v []LinksInner1)`

SetLinks sets Links field to given value.


### GetSubscriptions

`func (o *FormattedApiOrgPublic) GetSubscriptions() Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FormattedApiOrgPublic) GetSubscriptionsOk() (*Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FormattedApiOrgPublic) SetSubscriptions(v Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


