# FormattedApiStackRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublicName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **float32** |  | [optional] 
**StackStateServiceUrl** | Pointer to **string** |  | [optional] 
**SyntheticMonitoringApiUrl** | Pointer to **string** |  | [optional] 
**InsightsApiUrl** | Pointer to **string** |  | [optional] 
**IntegrationsApiUrl** | Pointer to **string** |  | [optional] 
**HostedExportersApiUrl** | Pointer to **string** |  | [optional] 
**MachineLearningApiUrl** | Pointer to **string** |  | [optional] 
**LlmGatewayUrl** | Pointer to **string** |  | [optional] 
**IncidentApiUrl** | Pointer to **string** |  | [optional] 
**OncallApiUrl** | Pointer to **string** |  | [optional] 
**FaroEndpointUrl** | Pointer to **string** |  | [optional] 
**PdcClusterSlug** | Pointer to **string** |  | [optional] 
**OtlpHttpUrl** | Pointer to **NullableString** |  | [optional] 
**OtlpPrivateConnectivityInfo** | Pointer to [**NullableOtlpPrivateConnectivityInfo**](OtlpPrivateConnectivityInfo.md) |  | [optional] 
**AuthApiUrl** | Pointer to **string** |  | [optional] 
**AuthApiTokenSet** | Pointer to **bool** |  | [optional] 
**HgClusterId** | Pointer to **float32** |  | [optional] 
**HgClusterSlug** | Pointer to **string** |  | [optional] 
**HgClusterName** | Pointer to **string** |  | [optional] 
**HgClusterUrl** | Pointer to **string** |  | [optional] 
**HmPromClusterId** | Pointer to **float32** |  | [optional] 
**HmPromClusterSlug** | Pointer to **string** |  | [optional] 
**HmPromClusterName** | Pointer to **string** |  | [optional] 
**HmPromClusterUrl** | Pointer to **string** |  | [optional] 
**HmGraphiteClusterId** | Pointer to **float32** |  | [optional] 
**HmGraphiteClusterSlug** | Pointer to **string** |  | [optional] 
**HmGraphiteClusterName** | Pointer to **string** |  | [optional] 
**HmGraphiteClusterUrl** | Pointer to **string** |  | [optional] 
**HlClusterId** | Pointer to **float32** |  | [optional] 
**HlClusterSlug** | Pointer to **string** |  | [optional] 
**HlClusterName** | Pointer to **string** |  | [optional] 
**HlClusterUrl** | Pointer to **string** |  | [optional] 
**AmClusterId** | Pointer to **float32** |  | [optional] 
**AmClusterSlug** | Pointer to **string** |  | [optional] 
**AmClusterName** | Pointer to **string** |  | [optional] 
**AmClusterUrl** | Pointer to **string** |  | [optional] 
**HtClusterId** | Pointer to **float32** |  | [optional] 
**HtClusterSlug** | Pointer to **string** |  | [optional] 
**HtClusterName** | Pointer to **string** |  | [optional] 
**HtClusterUrl** | Pointer to **string** |  | [optional] 
**HpClusterId** | Pointer to **float32** |  | [optional] 
**HpClusterSlug** | Pointer to **string** |  | [optional] 
**HpClusterName** | Pointer to **string** |  | [optional] 
**HpClusterUrl** | Pointer to **string** |  | [optional] 
**AgmClusterId** | Pointer to **float32** |  | [optional] 
**AgmClusterSlug** | Pointer to **string** |  | [optional] 
**AgmClusterName** | Pointer to **string** |  | [optional] 
**AgmClusterUrl** | Pointer to **string** |  | [optional] 
**ProviderRegion** | Pointer to **string** |  | [optional] 

## Methods

### NewFormattedApiStackRegion

`func NewFormattedApiStackRegion() *FormattedApiStackRegion`

NewFormattedApiStackRegion instantiates a new FormattedApiStackRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiStackRegionWithDefaults

`func NewFormattedApiStackRegionWithDefaults() *FormattedApiStackRegion`

NewFormattedApiStackRegionWithDefaults instantiates a new FormattedApiStackRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormattedApiStackRegion) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiStackRegion) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiStackRegion) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *FormattedApiStackRegion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FormattedApiStackRegion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormattedApiStackRegion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormattedApiStackRegion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FormattedApiStackRegion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *FormattedApiStackRegion) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormattedApiStackRegion) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormattedApiStackRegion) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormattedApiStackRegion) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSlug

`func (o *FormattedApiStackRegion) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormattedApiStackRegion) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormattedApiStackRegion) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FormattedApiStackRegion) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *FormattedApiStackRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiStackRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiStackRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormattedApiStackRegion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicName

`func (o *FormattedApiStackRegion) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *FormattedApiStackRegion) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *FormattedApiStackRegion) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *FormattedApiStackRegion) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### GetDescription

`func (o *FormattedApiStackRegion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormattedApiStackRegion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormattedApiStackRegion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormattedApiStackRegion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvider

`func (o *FormattedApiStackRegion) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FormattedApiStackRegion) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FormattedApiStackRegion) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *FormattedApiStackRegion) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FormattedApiStackRegion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiStackRegion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiStackRegion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormattedApiStackRegion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormattedApiStackRegion) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiStackRegion) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiStackRegion) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormattedApiStackRegion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *FormattedApiStackRegion) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiStackRegion) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSortOrder

`func (o *FormattedApiStackRegion) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FormattedApiStackRegion) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FormattedApiStackRegion) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FormattedApiStackRegion) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStackStateServiceUrl

`func (o *FormattedApiStackRegion) GetStackStateServiceUrl() string`

GetStackStateServiceUrl returns the StackStateServiceUrl field if non-nil, zero value otherwise.

### GetStackStateServiceUrlOk

`func (o *FormattedApiStackRegion) GetStackStateServiceUrlOk() (*string, bool)`

GetStackStateServiceUrlOk returns a tuple with the StackStateServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStateServiceUrl

`func (o *FormattedApiStackRegion) SetStackStateServiceUrl(v string)`

SetStackStateServiceUrl sets StackStateServiceUrl field to given value.

### HasStackStateServiceUrl

`func (o *FormattedApiStackRegion) HasStackStateServiceUrl() bool`

HasStackStateServiceUrl returns a boolean if a field has been set.

### GetSyntheticMonitoringApiUrl

`func (o *FormattedApiStackRegion) GetSyntheticMonitoringApiUrl() string`

GetSyntheticMonitoringApiUrl returns the SyntheticMonitoringApiUrl field if non-nil, zero value otherwise.

### GetSyntheticMonitoringApiUrlOk

`func (o *FormattedApiStackRegion) GetSyntheticMonitoringApiUrlOk() (*string, bool)`

GetSyntheticMonitoringApiUrlOk returns a tuple with the SyntheticMonitoringApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMonitoringApiUrl

`func (o *FormattedApiStackRegion) SetSyntheticMonitoringApiUrl(v string)`

SetSyntheticMonitoringApiUrl sets SyntheticMonitoringApiUrl field to given value.

### HasSyntheticMonitoringApiUrl

`func (o *FormattedApiStackRegion) HasSyntheticMonitoringApiUrl() bool`

HasSyntheticMonitoringApiUrl returns a boolean if a field has been set.

### GetInsightsApiUrl

`func (o *FormattedApiStackRegion) GetInsightsApiUrl() string`

GetInsightsApiUrl returns the InsightsApiUrl field if non-nil, zero value otherwise.

### GetInsightsApiUrlOk

`func (o *FormattedApiStackRegion) GetInsightsApiUrlOk() (*string, bool)`

GetInsightsApiUrlOk returns a tuple with the InsightsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsApiUrl

`func (o *FormattedApiStackRegion) SetInsightsApiUrl(v string)`

SetInsightsApiUrl sets InsightsApiUrl field to given value.

### HasInsightsApiUrl

`func (o *FormattedApiStackRegion) HasInsightsApiUrl() bool`

HasInsightsApiUrl returns a boolean if a field has been set.

### GetIntegrationsApiUrl

`func (o *FormattedApiStackRegion) GetIntegrationsApiUrl() string`

GetIntegrationsApiUrl returns the IntegrationsApiUrl field if non-nil, zero value otherwise.

### GetIntegrationsApiUrlOk

`func (o *FormattedApiStackRegion) GetIntegrationsApiUrlOk() (*string, bool)`

GetIntegrationsApiUrlOk returns a tuple with the IntegrationsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsApiUrl

`func (o *FormattedApiStackRegion) SetIntegrationsApiUrl(v string)`

SetIntegrationsApiUrl sets IntegrationsApiUrl field to given value.

### HasIntegrationsApiUrl

`func (o *FormattedApiStackRegion) HasIntegrationsApiUrl() bool`

HasIntegrationsApiUrl returns a boolean if a field has been set.

### GetHostedExportersApiUrl

`func (o *FormattedApiStackRegion) GetHostedExportersApiUrl() string`

GetHostedExportersApiUrl returns the HostedExportersApiUrl field if non-nil, zero value otherwise.

### GetHostedExportersApiUrlOk

`func (o *FormattedApiStackRegion) GetHostedExportersApiUrlOk() (*string, bool)`

GetHostedExportersApiUrlOk returns a tuple with the HostedExportersApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedExportersApiUrl

`func (o *FormattedApiStackRegion) SetHostedExportersApiUrl(v string)`

SetHostedExportersApiUrl sets HostedExportersApiUrl field to given value.

### HasHostedExportersApiUrl

`func (o *FormattedApiStackRegion) HasHostedExportersApiUrl() bool`

HasHostedExportersApiUrl returns a boolean if a field has been set.

### GetMachineLearningApiUrl

`func (o *FormattedApiStackRegion) GetMachineLearningApiUrl() string`

GetMachineLearningApiUrl returns the MachineLearningApiUrl field if non-nil, zero value otherwise.

### GetMachineLearningApiUrlOk

`func (o *FormattedApiStackRegion) GetMachineLearningApiUrlOk() (*string, bool)`

GetMachineLearningApiUrlOk returns a tuple with the MachineLearningApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningApiUrl

`func (o *FormattedApiStackRegion) SetMachineLearningApiUrl(v string)`

SetMachineLearningApiUrl sets MachineLearningApiUrl field to given value.

### HasMachineLearningApiUrl

`func (o *FormattedApiStackRegion) HasMachineLearningApiUrl() bool`

HasMachineLearningApiUrl returns a boolean if a field has been set.

### GetLlmGatewayUrl

`func (o *FormattedApiStackRegion) GetLlmGatewayUrl() string`

GetLlmGatewayUrl returns the LlmGatewayUrl field if non-nil, zero value otherwise.

### GetLlmGatewayUrlOk

`func (o *FormattedApiStackRegion) GetLlmGatewayUrlOk() (*string, bool)`

GetLlmGatewayUrlOk returns a tuple with the LlmGatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmGatewayUrl

`func (o *FormattedApiStackRegion) SetLlmGatewayUrl(v string)`

SetLlmGatewayUrl sets LlmGatewayUrl field to given value.

### HasLlmGatewayUrl

`func (o *FormattedApiStackRegion) HasLlmGatewayUrl() bool`

HasLlmGatewayUrl returns a boolean if a field has been set.

### GetIncidentApiUrl

`func (o *FormattedApiStackRegion) GetIncidentApiUrl() string`

GetIncidentApiUrl returns the IncidentApiUrl field if non-nil, zero value otherwise.

### GetIncidentApiUrlOk

`func (o *FormattedApiStackRegion) GetIncidentApiUrlOk() (*string, bool)`

GetIncidentApiUrlOk returns a tuple with the IncidentApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentApiUrl

`func (o *FormattedApiStackRegion) SetIncidentApiUrl(v string)`

SetIncidentApiUrl sets IncidentApiUrl field to given value.

### HasIncidentApiUrl

`func (o *FormattedApiStackRegion) HasIncidentApiUrl() bool`

HasIncidentApiUrl returns a boolean if a field has been set.

### GetOncallApiUrl

`func (o *FormattedApiStackRegion) GetOncallApiUrl() string`

GetOncallApiUrl returns the OncallApiUrl field if non-nil, zero value otherwise.

### GetOncallApiUrlOk

`func (o *FormattedApiStackRegion) GetOncallApiUrlOk() (*string, bool)`

GetOncallApiUrlOk returns a tuple with the OncallApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncallApiUrl

`func (o *FormattedApiStackRegion) SetOncallApiUrl(v string)`

SetOncallApiUrl sets OncallApiUrl field to given value.

### HasOncallApiUrl

`func (o *FormattedApiStackRegion) HasOncallApiUrl() bool`

HasOncallApiUrl returns a boolean if a field has been set.

### GetFaroEndpointUrl

`func (o *FormattedApiStackRegion) GetFaroEndpointUrl() string`

GetFaroEndpointUrl returns the FaroEndpointUrl field if non-nil, zero value otherwise.

### GetFaroEndpointUrlOk

`func (o *FormattedApiStackRegion) GetFaroEndpointUrlOk() (*string, bool)`

GetFaroEndpointUrlOk returns a tuple with the FaroEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaroEndpointUrl

`func (o *FormattedApiStackRegion) SetFaroEndpointUrl(v string)`

SetFaroEndpointUrl sets FaroEndpointUrl field to given value.

### HasFaroEndpointUrl

`func (o *FormattedApiStackRegion) HasFaroEndpointUrl() bool`

HasFaroEndpointUrl returns a boolean if a field has been set.

### GetPdcClusterSlug

`func (o *FormattedApiStackRegion) GetPdcClusterSlug() string`

GetPdcClusterSlug returns the PdcClusterSlug field if non-nil, zero value otherwise.

### GetPdcClusterSlugOk

`func (o *FormattedApiStackRegion) GetPdcClusterSlugOk() (*string, bool)`

GetPdcClusterSlugOk returns a tuple with the PdcClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdcClusterSlug

`func (o *FormattedApiStackRegion) SetPdcClusterSlug(v string)`

SetPdcClusterSlug sets PdcClusterSlug field to given value.

### HasPdcClusterSlug

`func (o *FormattedApiStackRegion) HasPdcClusterSlug() bool`

HasPdcClusterSlug returns a boolean if a field has been set.

### GetOtlpHttpUrl

`func (o *FormattedApiStackRegion) GetOtlpHttpUrl() string`

GetOtlpHttpUrl returns the OtlpHttpUrl field if non-nil, zero value otherwise.

### GetOtlpHttpUrlOk

`func (o *FormattedApiStackRegion) GetOtlpHttpUrlOk() (*string, bool)`

GetOtlpHttpUrlOk returns a tuple with the OtlpHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpHttpUrl

`func (o *FormattedApiStackRegion) SetOtlpHttpUrl(v string)`

SetOtlpHttpUrl sets OtlpHttpUrl field to given value.

### HasOtlpHttpUrl

`func (o *FormattedApiStackRegion) HasOtlpHttpUrl() bool`

HasOtlpHttpUrl returns a boolean if a field has been set.

### SetOtlpHttpUrlNil

`func (o *FormattedApiStackRegion) SetOtlpHttpUrlNil(b bool)`

 SetOtlpHttpUrlNil sets the value for OtlpHttpUrl to be an explicit nil

### UnsetOtlpHttpUrl
`func (o *FormattedApiStackRegion) UnsetOtlpHttpUrl()`

UnsetOtlpHttpUrl ensures that no value is present for OtlpHttpUrl, not even an explicit nil
### GetOtlpPrivateConnectivityInfo

`func (o *FormattedApiStackRegion) GetOtlpPrivateConnectivityInfo() OtlpPrivateConnectivityInfo`

GetOtlpPrivateConnectivityInfo returns the OtlpPrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetOtlpPrivateConnectivityInfoOk

`func (o *FormattedApiStackRegion) GetOtlpPrivateConnectivityInfoOk() (*OtlpPrivateConnectivityInfo, bool)`

GetOtlpPrivateConnectivityInfoOk returns a tuple with the OtlpPrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpPrivateConnectivityInfo

`func (o *FormattedApiStackRegion) SetOtlpPrivateConnectivityInfo(v OtlpPrivateConnectivityInfo)`

SetOtlpPrivateConnectivityInfo sets OtlpPrivateConnectivityInfo field to given value.

### HasOtlpPrivateConnectivityInfo

`func (o *FormattedApiStackRegion) HasOtlpPrivateConnectivityInfo() bool`

HasOtlpPrivateConnectivityInfo returns a boolean if a field has been set.

### SetOtlpPrivateConnectivityInfoNil

`func (o *FormattedApiStackRegion) SetOtlpPrivateConnectivityInfoNil(b bool)`

 SetOtlpPrivateConnectivityInfoNil sets the value for OtlpPrivateConnectivityInfo to be an explicit nil

### UnsetOtlpPrivateConnectivityInfo
`func (o *FormattedApiStackRegion) UnsetOtlpPrivateConnectivityInfo()`

UnsetOtlpPrivateConnectivityInfo ensures that no value is present for OtlpPrivateConnectivityInfo, not even an explicit nil
### GetAuthApiUrl

`func (o *FormattedApiStackRegion) GetAuthApiUrl() string`

GetAuthApiUrl returns the AuthApiUrl field if non-nil, zero value otherwise.

### GetAuthApiUrlOk

`func (o *FormattedApiStackRegion) GetAuthApiUrlOk() (*string, bool)`

GetAuthApiUrlOk returns a tuple with the AuthApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthApiUrl

`func (o *FormattedApiStackRegion) SetAuthApiUrl(v string)`

SetAuthApiUrl sets AuthApiUrl field to given value.

### HasAuthApiUrl

`func (o *FormattedApiStackRegion) HasAuthApiUrl() bool`

HasAuthApiUrl returns a boolean if a field has been set.

### GetAuthApiTokenSet

`func (o *FormattedApiStackRegion) GetAuthApiTokenSet() bool`

GetAuthApiTokenSet returns the AuthApiTokenSet field if non-nil, zero value otherwise.

### GetAuthApiTokenSetOk

`func (o *FormattedApiStackRegion) GetAuthApiTokenSetOk() (*bool, bool)`

GetAuthApiTokenSetOk returns a tuple with the AuthApiTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthApiTokenSet

`func (o *FormattedApiStackRegion) SetAuthApiTokenSet(v bool)`

SetAuthApiTokenSet sets AuthApiTokenSet field to given value.

### HasAuthApiTokenSet

`func (o *FormattedApiStackRegion) HasAuthApiTokenSet() bool`

HasAuthApiTokenSet returns a boolean if a field has been set.

### GetHgClusterId

`func (o *FormattedApiStackRegion) GetHgClusterId() float32`

GetHgClusterId returns the HgClusterId field if non-nil, zero value otherwise.

### GetHgClusterIdOk

`func (o *FormattedApiStackRegion) GetHgClusterIdOk() (*float32, bool)`

GetHgClusterIdOk returns a tuple with the HgClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterId

`func (o *FormattedApiStackRegion) SetHgClusterId(v float32)`

SetHgClusterId sets HgClusterId field to given value.

### HasHgClusterId

`func (o *FormattedApiStackRegion) HasHgClusterId() bool`

HasHgClusterId returns a boolean if a field has been set.

### GetHgClusterSlug

`func (o *FormattedApiStackRegion) GetHgClusterSlug() string`

GetHgClusterSlug returns the HgClusterSlug field if non-nil, zero value otherwise.

### GetHgClusterSlugOk

`func (o *FormattedApiStackRegion) GetHgClusterSlugOk() (*string, bool)`

GetHgClusterSlugOk returns a tuple with the HgClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterSlug

`func (o *FormattedApiStackRegion) SetHgClusterSlug(v string)`

SetHgClusterSlug sets HgClusterSlug field to given value.

### HasHgClusterSlug

`func (o *FormattedApiStackRegion) HasHgClusterSlug() bool`

HasHgClusterSlug returns a boolean if a field has been set.

### GetHgClusterName

`func (o *FormattedApiStackRegion) GetHgClusterName() string`

GetHgClusterName returns the HgClusterName field if non-nil, zero value otherwise.

### GetHgClusterNameOk

`func (o *FormattedApiStackRegion) GetHgClusterNameOk() (*string, bool)`

GetHgClusterNameOk returns a tuple with the HgClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterName

`func (o *FormattedApiStackRegion) SetHgClusterName(v string)`

SetHgClusterName sets HgClusterName field to given value.

### HasHgClusterName

`func (o *FormattedApiStackRegion) HasHgClusterName() bool`

HasHgClusterName returns a boolean if a field has been set.

### GetHgClusterUrl

`func (o *FormattedApiStackRegion) GetHgClusterUrl() string`

GetHgClusterUrl returns the HgClusterUrl field if non-nil, zero value otherwise.

### GetHgClusterUrlOk

`func (o *FormattedApiStackRegion) GetHgClusterUrlOk() (*string, bool)`

GetHgClusterUrlOk returns a tuple with the HgClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterUrl

`func (o *FormattedApiStackRegion) SetHgClusterUrl(v string)`

SetHgClusterUrl sets HgClusterUrl field to given value.

### HasHgClusterUrl

`func (o *FormattedApiStackRegion) HasHgClusterUrl() bool`

HasHgClusterUrl returns a boolean if a field has been set.

### GetHmPromClusterId

`func (o *FormattedApiStackRegion) GetHmPromClusterId() float32`

GetHmPromClusterId returns the HmPromClusterId field if non-nil, zero value otherwise.

### GetHmPromClusterIdOk

`func (o *FormattedApiStackRegion) GetHmPromClusterIdOk() (*float32, bool)`

GetHmPromClusterIdOk returns a tuple with the HmPromClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterId

`func (o *FormattedApiStackRegion) SetHmPromClusterId(v float32)`

SetHmPromClusterId sets HmPromClusterId field to given value.

### HasHmPromClusterId

`func (o *FormattedApiStackRegion) HasHmPromClusterId() bool`

HasHmPromClusterId returns a boolean if a field has been set.

### GetHmPromClusterSlug

`func (o *FormattedApiStackRegion) GetHmPromClusterSlug() string`

GetHmPromClusterSlug returns the HmPromClusterSlug field if non-nil, zero value otherwise.

### GetHmPromClusterSlugOk

`func (o *FormattedApiStackRegion) GetHmPromClusterSlugOk() (*string, bool)`

GetHmPromClusterSlugOk returns a tuple with the HmPromClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterSlug

`func (o *FormattedApiStackRegion) SetHmPromClusterSlug(v string)`

SetHmPromClusterSlug sets HmPromClusterSlug field to given value.

### HasHmPromClusterSlug

`func (o *FormattedApiStackRegion) HasHmPromClusterSlug() bool`

HasHmPromClusterSlug returns a boolean if a field has been set.

### GetHmPromClusterName

`func (o *FormattedApiStackRegion) GetHmPromClusterName() string`

GetHmPromClusterName returns the HmPromClusterName field if non-nil, zero value otherwise.

### GetHmPromClusterNameOk

`func (o *FormattedApiStackRegion) GetHmPromClusterNameOk() (*string, bool)`

GetHmPromClusterNameOk returns a tuple with the HmPromClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterName

`func (o *FormattedApiStackRegion) SetHmPromClusterName(v string)`

SetHmPromClusterName sets HmPromClusterName field to given value.

### HasHmPromClusterName

`func (o *FormattedApiStackRegion) HasHmPromClusterName() bool`

HasHmPromClusterName returns a boolean if a field has been set.

### GetHmPromClusterUrl

`func (o *FormattedApiStackRegion) GetHmPromClusterUrl() string`

GetHmPromClusterUrl returns the HmPromClusterUrl field if non-nil, zero value otherwise.

### GetHmPromClusterUrlOk

`func (o *FormattedApiStackRegion) GetHmPromClusterUrlOk() (*string, bool)`

GetHmPromClusterUrlOk returns a tuple with the HmPromClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterUrl

`func (o *FormattedApiStackRegion) SetHmPromClusterUrl(v string)`

SetHmPromClusterUrl sets HmPromClusterUrl field to given value.

### HasHmPromClusterUrl

`func (o *FormattedApiStackRegion) HasHmPromClusterUrl() bool`

HasHmPromClusterUrl returns a boolean if a field has been set.

### GetHmGraphiteClusterId

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterId() float32`

GetHmGraphiteClusterId returns the HmGraphiteClusterId field if non-nil, zero value otherwise.

### GetHmGraphiteClusterIdOk

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterIdOk() (*float32, bool)`

GetHmGraphiteClusterIdOk returns a tuple with the HmGraphiteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterId

`func (o *FormattedApiStackRegion) SetHmGraphiteClusterId(v float32)`

SetHmGraphiteClusterId sets HmGraphiteClusterId field to given value.

### HasHmGraphiteClusterId

`func (o *FormattedApiStackRegion) HasHmGraphiteClusterId() bool`

HasHmGraphiteClusterId returns a boolean if a field has been set.

### GetHmGraphiteClusterSlug

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterSlug() string`

GetHmGraphiteClusterSlug returns the HmGraphiteClusterSlug field if non-nil, zero value otherwise.

### GetHmGraphiteClusterSlugOk

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterSlugOk() (*string, bool)`

GetHmGraphiteClusterSlugOk returns a tuple with the HmGraphiteClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterSlug

`func (o *FormattedApiStackRegion) SetHmGraphiteClusterSlug(v string)`

SetHmGraphiteClusterSlug sets HmGraphiteClusterSlug field to given value.

### HasHmGraphiteClusterSlug

`func (o *FormattedApiStackRegion) HasHmGraphiteClusterSlug() bool`

HasHmGraphiteClusterSlug returns a boolean if a field has been set.

### GetHmGraphiteClusterName

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterName() string`

GetHmGraphiteClusterName returns the HmGraphiteClusterName field if non-nil, zero value otherwise.

### GetHmGraphiteClusterNameOk

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterNameOk() (*string, bool)`

GetHmGraphiteClusterNameOk returns a tuple with the HmGraphiteClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterName

`func (o *FormattedApiStackRegion) SetHmGraphiteClusterName(v string)`

SetHmGraphiteClusterName sets HmGraphiteClusterName field to given value.

### HasHmGraphiteClusterName

`func (o *FormattedApiStackRegion) HasHmGraphiteClusterName() bool`

HasHmGraphiteClusterName returns a boolean if a field has been set.

### GetHmGraphiteClusterUrl

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterUrl() string`

GetHmGraphiteClusterUrl returns the HmGraphiteClusterUrl field if non-nil, zero value otherwise.

### GetHmGraphiteClusterUrlOk

`func (o *FormattedApiStackRegion) GetHmGraphiteClusterUrlOk() (*string, bool)`

GetHmGraphiteClusterUrlOk returns a tuple with the HmGraphiteClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterUrl

`func (o *FormattedApiStackRegion) SetHmGraphiteClusterUrl(v string)`

SetHmGraphiteClusterUrl sets HmGraphiteClusterUrl field to given value.

### HasHmGraphiteClusterUrl

`func (o *FormattedApiStackRegion) HasHmGraphiteClusterUrl() bool`

HasHmGraphiteClusterUrl returns a boolean if a field has been set.

### GetHlClusterId

`func (o *FormattedApiStackRegion) GetHlClusterId() float32`

GetHlClusterId returns the HlClusterId field if non-nil, zero value otherwise.

### GetHlClusterIdOk

`func (o *FormattedApiStackRegion) GetHlClusterIdOk() (*float32, bool)`

GetHlClusterIdOk returns a tuple with the HlClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterId

`func (o *FormattedApiStackRegion) SetHlClusterId(v float32)`

SetHlClusterId sets HlClusterId field to given value.

### HasHlClusterId

`func (o *FormattedApiStackRegion) HasHlClusterId() bool`

HasHlClusterId returns a boolean if a field has been set.

### GetHlClusterSlug

`func (o *FormattedApiStackRegion) GetHlClusterSlug() string`

GetHlClusterSlug returns the HlClusterSlug field if non-nil, zero value otherwise.

### GetHlClusterSlugOk

`func (o *FormattedApiStackRegion) GetHlClusterSlugOk() (*string, bool)`

GetHlClusterSlugOk returns a tuple with the HlClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterSlug

`func (o *FormattedApiStackRegion) SetHlClusterSlug(v string)`

SetHlClusterSlug sets HlClusterSlug field to given value.

### HasHlClusterSlug

`func (o *FormattedApiStackRegion) HasHlClusterSlug() bool`

HasHlClusterSlug returns a boolean if a field has been set.

### GetHlClusterName

`func (o *FormattedApiStackRegion) GetHlClusterName() string`

GetHlClusterName returns the HlClusterName field if non-nil, zero value otherwise.

### GetHlClusterNameOk

`func (o *FormattedApiStackRegion) GetHlClusterNameOk() (*string, bool)`

GetHlClusterNameOk returns a tuple with the HlClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterName

`func (o *FormattedApiStackRegion) SetHlClusterName(v string)`

SetHlClusterName sets HlClusterName field to given value.

### HasHlClusterName

`func (o *FormattedApiStackRegion) HasHlClusterName() bool`

HasHlClusterName returns a boolean if a field has been set.

### GetHlClusterUrl

`func (o *FormattedApiStackRegion) GetHlClusterUrl() string`

GetHlClusterUrl returns the HlClusterUrl field if non-nil, zero value otherwise.

### GetHlClusterUrlOk

`func (o *FormattedApiStackRegion) GetHlClusterUrlOk() (*string, bool)`

GetHlClusterUrlOk returns a tuple with the HlClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterUrl

`func (o *FormattedApiStackRegion) SetHlClusterUrl(v string)`

SetHlClusterUrl sets HlClusterUrl field to given value.

### HasHlClusterUrl

`func (o *FormattedApiStackRegion) HasHlClusterUrl() bool`

HasHlClusterUrl returns a boolean if a field has been set.

### GetAmClusterId

`func (o *FormattedApiStackRegion) GetAmClusterId() float32`

GetAmClusterId returns the AmClusterId field if non-nil, zero value otherwise.

### GetAmClusterIdOk

`func (o *FormattedApiStackRegion) GetAmClusterIdOk() (*float32, bool)`

GetAmClusterIdOk returns a tuple with the AmClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterId

`func (o *FormattedApiStackRegion) SetAmClusterId(v float32)`

SetAmClusterId sets AmClusterId field to given value.

### HasAmClusterId

`func (o *FormattedApiStackRegion) HasAmClusterId() bool`

HasAmClusterId returns a boolean if a field has been set.

### GetAmClusterSlug

`func (o *FormattedApiStackRegion) GetAmClusterSlug() string`

GetAmClusterSlug returns the AmClusterSlug field if non-nil, zero value otherwise.

### GetAmClusterSlugOk

`func (o *FormattedApiStackRegion) GetAmClusterSlugOk() (*string, bool)`

GetAmClusterSlugOk returns a tuple with the AmClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterSlug

`func (o *FormattedApiStackRegion) SetAmClusterSlug(v string)`

SetAmClusterSlug sets AmClusterSlug field to given value.

### HasAmClusterSlug

`func (o *FormattedApiStackRegion) HasAmClusterSlug() bool`

HasAmClusterSlug returns a boolean if a field has been set.

### GetAmClusterName

`func (o *FormattedApiStackRegion) GetAmClusterName() string`

GetAmClusterName returns the AmClusterName field if non-nil, zero value otherwise.

### GetAmClusterNameOk

`func (o *FormattedApiStackRegion) GetAmClusterNameOk() (*string, bool)`

GetAmClusterNameOk returns a tuple with the AmClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterName

`func (o *FormattedApiStackRegion) SetAmClusterName(v string)`

SetAmClusterName sets AmClusterName field to given value.

### HasAmClusterName

`func (o *FormattedApiStackRegion) HasAmClusterName() bool`

HasAmClusterName returns a boolean if a field has been set.

### GetAmClusterUrl

`func (o *FormattedApiStackRegion) GetAmClusterUrl() string`

GetAmClusterUrl returns the AmClusterUrl field if non-nil, zero value otherwise.

### GetAmClusterUrlOk

`func (o *FormattedApiStackRegion) GetAmClusterUrlOk() (*string, bool)`

GetAmClusterUrlOk returns a tuple with the AmClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterUrl

`func (o *FormattedApiStackRegion) SetAmClusterUrl(v string)`

SetAmClusterUrl sets AmClusterUrl field to given value.

### HasAmClusterUrl

`func (o *FormattedApiStackRegion) HasAmClusterUrl() bool`

HasAmClusterUrl returns a boolean if a field has been set.

### GetHtClusterId

`func (o *FormattedApiStackRegion) GetHtClusterId() float32`

GetHtClusterId returns the HtClusterId field if non-nil, zero value otherwise.

### GetHtClusterIdOk

`func (o *FormattedApiStackRegion) GetHtClusterIdOk() (*float32, bool)`

GetHtClusterIdOk returns a tuple with the HtClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterId

`func (o *FormattedApiStackRegion) SetHtClusterId(v float32)`

SetHtClusterId sets HtClusterId field to given value.

### HasHtClusterId

`func (o *FormattedApiStackRegion) HasHtClusterId() bool`

HasHtClusterId returns a boolean if a field has been set.

### GetHtClusterSlug

`func (o *FormattedApiStackRegion) GetHtClusterSlug() string`

GetHtClusterSlug returns the HtClusterSlug field if non-nil, zero value otherwise.

### GetHtClusterSlugOk

`func (o *FormattedApiStackRegion) GetHtClusterSlugOk() (*string, bool)`

GetHtClusterSlugOk returns a tuple with the HtClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterSlug

`func (o *FormattedApiStackRegion) SetHtClusterSlug(v string)`

SetHtClusterSlug sets HtClusterSlug field to given value.

### HasHtClusterSlug

`func (o *FormattedApiStackRegion) HasHtClusterSlug() bool`

HasHtClusterSlug returns a boolean if a field has been set.

### GetHtClusterName

`func (o *FormattedApiStackRegion) GetHtClusterName() string`

GetHtClusterName returns the HtClusterName field if non-nil, zero value otherwise.

### GetHtClusterNameOk

`func (o *FormattedApiStackRegion) GetHtClusterNameOk() (*string, bool)`

GetHtClusterNameOk returns a tuple with the HtClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterName

`func (o *FormattedApiStackRegion) SetHtClusterName(v string)`

SetHtClusterName sets HtClusterName field to given value.

### HasHtClusterName

`func (o *FormattedApiStackRegion) HasHtClusterName() bool`

HasHtClusterName returns a boolean if a field has been set.

### GetHtClusterUrl

`func (o *FormattedApiStackRegion) GetHtClusterUrl() string`

GetHtClusterUrl returns the HtClusterUrl field if non-nil, zero value otherwise.

### GetHtClusterUrlOk

`func (o *FormattedApiStackRegion) GetHtClusterUrlOk() (*string, bool)`

GetHtClusterUrlOk returns a tuple with the HtClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterUrl

`func (o *FormattedApiStackRegion) SetHtClusterUrl(v string)`

SetHtClusterUrl sets HtClusterUrl field to given value.

### HasHtClusterUrl

`func (o *FormattedApiStackRegion) HasHtClusterUrl() bool`

HasHtClusterUrl returns a boolean if a field has been set.

### GetHpClusterId

`func (o *FormattedApiStackRegion) GetHpClusterId() float32`

GetHpClusterId returns the HpClusterId field if non-nil, zero value otherwise.

### GetHpClusterIdOk

`func (o *FormattedApiStackRegion) GetHpClusterIdOk() (*float32, bool)`

GetHpClusterIdOk returns a tuple with the HpClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterId

`func (o *FormattedApiStackRegion) SetHpClusterId(v float32)`

SetHpClusterId sets HpClusterId field to given value.

### HasHpClusterId

`func (o *FormattedApiStackRegion) HasHpClusterId() bool`

HasHpClusterId returns a boolean if a field has been set.

### GetHpClusterSlug

`func (o *FormattedApiStackRegion) GetHpClusterSlug() string`

GetHpClusterSlug returns the HpClusterSlug field if non-nil, zero value otherwise.

### GetHpClusterSlugOk

`func (o *FormattedApiStackRegion) GetHpClusterSlugOk() (*string, bool)`

GetHpClusterSlugOk returns a tuple with the HpClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterSlug

`func (o *FormattedApiStackRegion) SetHpClusterSlug(v string)`

SetHpClusterSlug sets HpClusterSlug field to given value.

### HasHpClusterSlug

`func (o *FormattedApiStackRegion) HasHpClusterSlug() bool`

HasHpClusterSlug returns a boolean if a field has been set.

### GetHpClusterName

`func (o *FormattedApiStackRegion) GetHpClusterName() string`

GetHpClusterName returns the HpClusterName field if non-nil, zero value otherwise.

### GetHpClusterNameOk

`func (o *FormattedApiStackRegion) GetHpClusterNameOk() (*string, bool)`

GetHpClusterNameOk returns a tuple with the HpClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterName

`func (o *FormattedApiStackRegion) SetHpClusterName(v string)`

SetHpClusterName sets HpClusterName field to given value.

### HasHpClusterName

`func (o *FormattedApiStackRegion) HasHpClusterName() bool`

HasHpClusterName returns a boolean if a field has been set.

### GetHpClusterUrl

`func (o *FormattedApiStackRegion) GetHpClusterUrl() string`

GetHpClusterUrl returns the HpClusterUrl field if non-nil, zero value otherwise.

### GetHpClusterUrlOk

`func (o *FormattedApiStackRegion) GetHpClusterUrlOk() (*string, bool)`

GetHpClusterUrlOk returns a tuple with the HpClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterUrl

`func (o *FormattedApiStackRegion) SetHpClusterUrl(v string)`

SetHpClusterUrl sets HpClusterUrl field to given value.

### HasHpClusterUrl

`func (o *FormattedApiStackRegion) HasHpClusterUrl() bool`

HasHpClusterUrl returns a boolean if a field has been set.

### GetAgmClusterId

`func (o *FormattedApiStackRegion) GetAgmClusterId() float32`

GetAgmClusterId returns the AgmClusterId field if non-nil, zero value otherwise.

### GetAgmClusterIdOk

`func (o *FormattedApiStackRegion) GetAgmClusterIdOk() (*float32, bool)`

GetAgmClusterIdOk returns a tuple with the AgmClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterId

`func (o *FormattedApiStackRegion) SetAgmClusterId(v float32)`

SetAgmClusterId sets AgmClusterId field to given value.

### HasAgmClusterId

`func (o *FormattedApiStackRegion) HasAgmClusterId() bool`

HasAgmClusterId returns a boolean if a field has been set.

### GetAgmClusterSlug

`func (o *FormattedApiStackRegion) GetAgmClusterSlug() string`

GetAgmClusterSlug returns the AgmClusterSlug field if non-nil, zero value otherwise.

### GetAgmClusterSlugOk

`func (o *FormattedApiStackRegion) GetAgmClusterSlugOk() (*string, bool)`

GetAgmClusterSlugOk returns a tuple with the AgmClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterSlug

`func (o *FormattedApiStackRegion) SetAgmClusterSlug(v string)`

SetAgmClusterSlug sets AgmClusterSlug field to given value.

### HasAgmClusterSlug

`func (o *FormattedApiStackRegion) HasAgmClusterSlug() bool`

HasAgmClusterSlug returns a boolean if a field has been set.

### GetAgmClusterName

`func (o *FormattedApiStackRegion) GetAgmClusterName() string`

GetAgmClusterName returns the AgmClusterName field if non-nil, zero value otherwise.

### GetAgmClusterNameOk

`func (o *FormattedApiStackRegion) GetAgmClusterNameOk() (*string, bool)`

GetAgmClusterNameOk returns a tuple with the AgmClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterName

`func (o *FormattedApiStackRegion) SetAgmClusterName(v string)`

SetAgmClusterName sets AgmClusterName field to given value.

### HasAgmClusterName

`func (o *FormattedApiStackRegion) HasAgmClusterName() bool`

HasAgmClusterName returns a boolean if a field has been set.

### GetAgmClusterUrl

`func (o *FormattedApiStackRegion) GetAgmClusterUrl() string`

GetAgmClusterUrl returns the AgmClusterUrl field if non-nil, zero value otherwise.

### GetAgmClusterUrlOk

`func (o *FormattedApiStackRegion) GetAgmClusterUrlOk() (*string, bool)`

GetAgmClusterUrlOk returns a tuple with the AgmClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterUrl

`func (o *FormattedApiStackRegion) SetAgmClusterUrl(v string)`

SetAgmClusterUrl sets AgmClusterUrl field to given value.

### HasAgmClusterUrl

`func (o *FormattedApiStackRegion) HasAgmClusterUrl() bool`

HasAgmClusterUrl returns a boolean if a field has been set.

### GetProviderRegion

`func (o *FormattedApiStackRegion) GetProviderRegion() string`

GetProviderRegion returns the ProviderRegion field if non-nil, zero value otherwise.

### GetProviderRegionOk

`func (o *FormattedApiStackRegion) GetProviderRegionOk() (*string, bool)`

GetProviderRegionOk returns a tuple with the ProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRegion

`func (o *FormattedApiStackRegion) SetProviderRegion(v string)`

SetProviderRegion sets ProviderRegion field to given value.

### HasProviderRegion

`func (o *FormattedApiStackRegion) HasProviderRegion() bool`

HasProviderRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


