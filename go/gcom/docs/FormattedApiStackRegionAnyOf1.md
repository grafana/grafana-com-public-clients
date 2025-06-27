# FormattedApiStackRegionAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackStateServiceUrl** | **string** |  | 
**SyntheticMonitoringApiUrl** | **string** |  | 
**InsightsApiUrl** | **string** |  | 
**IntegrationsApiUrl** | **string** |  | 
**HostedExportersApiUrl** | **string** |  | 
**MachineLearningApiUrl** | **string** |  | 
**LlmGatewayUrl** | **string** |  | 
**AssistantUrl** | **string** |  | 
**IncidentApiUrl** | **string** |  | 
**OncallApiUrl** | **string** |  | 
**FaroEndpointUrl** | **string** |  | 
**PdcClusterSlug** | **string** |  | 
**PdcPrivateConnectivityInfo** | [**PdcPrivateConnectivityInfo**](PdcPrivateConnectivityInfo.md) |  | 
**OtlpHttpUrl** | **NullableString** |  | 
**OtlpPrivateConnectivityInfo** | [**OtlpPrivateConnectivityInfo**](OtlpPrivateConnectivityInfo.md) |  | 
**AuthApiUrl** | **string** |  | 
**AuthApiTokenSet** | **bool** |  | 
**HgClusterId** | **float32** |  | 
**HgClusterSlug** | **string** |  | 
**HgClusterName** | **string** |  | 
**HgClusterUrl** | **string** |  | 
**HmPromClusterId** | **float32** |  | 
**HmPromClusterSlug** | **string** |  | 
**HmPromClusterName** | **string** |  | 
**HmPromClusterUrl** | **string** |  | 
**HmGraphiteClusterId** | **float32** |  | 
**HmGraphiteClusterSlug** | **string** |  | 
**HmGraphiteClusterName** | **string** |  | 
**HmGraphiteClusterUrl** | **string** |  | 
**HlClusterId** | **float32** |  | 
**HlClusterSlug** | **string** |  | 
**HlClusterName** | **string** |  | 
**HlClusterUrl** | **string** |  | 
**AmClusterId** | **float32** |  | 
**AmClusterSlug** | **string** |  | 
**AmClusterName** | **string** |  | 
**AmClusterUrl** | **string** |  | 
**HtClusterId** | **float32** |  | 
**HtClusterSlug** | **string** |  | 
**HtClusterName** | **string** |  | 
**HtClusterUrl** | **string** |  | 
**HpClusterId** | **float32** |  | 
**HpClusterSlug** | **string** |  | 
**HpClusterName** | **string** |  | 
**HpClusterUrl** | **string** |  | 
**AgmClusterId** | **float32** |  | 
**AgmClusterSlug** | **string** |  | 
**AgmClusterName** | **string** |  | 
**AgmClusterUrl** | **string** |  | 
**AssertsGraphClusterId** | **float32** |  | 
**AssertsGraphClusterSlug** | **string** |  | 
**AssertsGraphClusterName** | **string** |  | 
**AssertsGraphClusterApiUrl** | Pointer to **NullableString** |  | [optional] 
**ProviderRegion** | **string** |  | 
**IsStub** | **bool** |  | 
**GeoLocation** | Pointer to **[]float32** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Complete** | **bool** |  | 
**Id** | **float32** |  | 
**Status** | **string** |  | 
**Visibility** | **string** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**PublicName** | **string** |  | 
**Description** | **string** |  | 
**Provider** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 
**SortOrder** | **float32** |  | 

## Methods

### NewFormattedApiStackRegionAnyOf1

`func NewFormattedApiStackRegionAnyOf1(stackStateServiceUrl string, syntheticMonitoringApiUrl string, insightsApiUrl string, integrationsApiUrl string, hostedExportersApiUrl string, machineLearningApiUrl string, llmGatewayUrl string, assistantUrl string, incidentApiUrl string, oncallApiUrl string, faroEndpointUrl string, pdcClusterSlug string, pdcPrivateConnectivityInfo PdcPrivateConnectivityInfo, otlpHttpUrl NullableString, otlpPrivateConnectivityInfo OtlpPrivateConnectivityInfo, authApiUrl string, authApiTokenSet bool, hgClusterId float32, hgClusterSlug string, hgClusterName string, hgClusterUrl string, hmPromClusterId float32, hmPromClusterSlug string, hmPromClusterName string, hmPromClusterUrl string, hmGraphiteClusterId float32, hmGraphiteClusterSlug string, hmGraphiteClusterName string, hmGraphiteClusterUrl string, hlClusterId float32, hlClusterSlug string, hlClusterName string, hlClusterUrl string, amClusterId float32, amClusterSlug string, amClusterName string, amClusterUrl string, htClusterId float32, htClusterSlug string, htClusterName string, htClusterUrl string, hpClusterId float32, hpClusterSlug string, hpClusterName string, hpClusterUrl string, agmClusterId float32, agmClusterSlug string, agmClusterName string, agmClusterUrl string, assertsGraphClusterId float32, assertsGraphClusterSlug string, assertsGraphClusterName string, providerRegion string, isStub bool, complete bool, id float32, status string, visibility string, slug string, name string, publicName string, description string, provider string, createdAt string, updatedAt NullableString, sortOrder float32, ) *FormattedApiStackRegionAnyOf1`

NewFormattedApiStackRegionAnyOf1 instantiates a new FormattedApiStackRegionAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiStackRegionAnyOf1WithDefaults

`func NewFormattedApiStackRegionAnyOf1WithDefaults() *FormattedApiStackRegionAnyOf1`

NewFormattedApiStackRegionAnyOf1WithDefaults instantiates a new FormattedApiStackRegionAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackStateServiceUrl

`func (o *FormattedApiStackRegionAnyOf1) GetStackStateServiceUrl() string`

GetStackStateServiceUrl returns the StackStateServiceUrl field if non-nil, zero value otherwise.

### GetStackStateServiceUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetStackStateServiceUrlOk() (*string, bool)`

GetStackStateServiceUrlOk returns a tuple with the StackStateServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStateServiceUrl

`func (o *FormattedApiStackRegionAnyOf1) SetStackStateServiceUrl(v string)`

SetStackStateServiceUrl sets StackStateServiceUrl field to given value.


### GetSyntheticMonitoringApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetSyntheticMonitoringApiUrl() string`

GetSyntheticMonitoringApiUrl returns the SyntheticMonitoringApiUrl field if non-nil, zero value otherwise.

### GetSyntheticMonitoringApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetSyntheticMonitoringApiUrlOk() (*string, bool)`

GetSyntheticMonitoringApiUrlOk returns a tuple with the SyntheticMonitoringApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMonitoringApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetSyntheticMonitoringApiUrl(v string)`

SetSyntheticMonitoringApiUrl sets SyntheticMonitoringApiUrl field to given value.


### GetInsightsApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetInsightsApiUrl() string`

GetInsightsApiUrl returns the InsightsApiUrl field if non-nil, zero value otherwise.

### GetInsightsApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetInsightsApiUrlOk() (*string, bool)`

GetInsightsApiUrlOk returns a tuple with the InsightsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetInsightsApiUrl(v string)`

SetInsightsApiUrl sets InsightsApiUrl field to given value.


### GetIntegrationsApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetIntegrationsApiUrl() string`

GetIntegrationsApiUrl returns the IntegrationsApiUrl field if non-nil, zero value otherwise.

### GetIntegrationsApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetIntegrationsApiUrlOk() (*string, bool)`

GetIntegrationsApiUrlOk returns a tuple with the IntegrationsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetIntegrationsApiUrl(v string)`

SetIntegrationsApiUrl sets IntegrationsApiUrl field to given value.


### GetHostedExportersApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHostedExportersApiUrl() string`

GetHostedExportersApiUrl returns the HostedExportersApiUrl field if non-nil, zero value otherwise.

### GetHostedExportersApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHostedExportersApiUrlOk() (*string, bool)`

GetHostedExportersApiUrlOk returns a tuple with the HostedExportersApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedExportersApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHostedExportersApiUrl(v string)`

SetHostedExportersApiUrl sets HostedExportersApiUrl field to given value.


### GetMachineLearningApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetMachineLearningApiUrl() string`

GetMachineLearningApiUrl returns the MachineLearningApiUrl field if non-nil, zero value otherwise.

### GetMachineLearningApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetMachineLearningApiUrlOk() (*string, bool)`

GetMachineLearningApiUrlOk returns a tuple with the MachineLearningApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetMachineLearningApiUrl(v string)`

SetMachineLearningApiUrl sets MachineLearningApiUrl field to given value.


### GetLlmGatewayUrl

`func (o *FormattedApiStackRegionAnyOf1) GetLlmGatewayUrl() string`

GetLlmGatewayUrl returns the LlmGatewayUrl field if non-nil, zero value otherwise.

### GetLlmGatewayUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetLlmGatewayUrlOk() (*string, bool)`

GetLlmGatewayUrlOk returns a tuple with the LlmGatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmGatewayUrl

`func (o *FormattedApiStackRegionAnyOf1) SetLlmGatewayUrl(v string)`

SetLlmGatewayUrl sets LlmGatewayUrl field to given value.


### GetAssistantUrl

`func (o *FormattedApiStackRegionAnyOf1) GetAssistantUrl() string`

GetAssistantUrl returns the AssistantUrl field if non-nil, zero value otherwise.

### GetAssistantUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetAssistantUrlOk() (*string, bool)`

GetAssistantUrlOk returns a tuple with the AssistantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantUrl

`func (o *FormattedApiStackRegionAnyOf1) SetAssistantUrl(v string)`

SetAssistantUrl sets AssistantUrl field to given value.


### GetIncidentApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetIncidentApiUrl() string`

GetIncidentApiUrl returns the IncidentApiUrl field if non-nil, zero value otherwise.

### GetIncidentApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetIncidentApiUrlOk() (*string, bool)`

GetIncidentApiUrlOk returns a tuple with the IncidentApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetIncidentApiUrl(v string)`

SetIncidentApiUrl sets IncidentApiUrl field to given value.


### GetOncallApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetOncallApiUrl() string`

GetOncallApiUrl returns the OncallApiUrl field if non-nil, zero value otherwise.

### GetOncallApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetOncallApiUrlOk() (*string, bool)`

GetOncallApiUrlOk returns a tuple with the OncallApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncallApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetOncallApiUrl(v string)`

SetOncallApiUrl sets OncallApiUrl field to given value.


### GetFaroEndpointUrl

`func (o *FormattedApiStackRegionAnyOf1) GetFaroEndpointUrl() string`

GetFaroEndpointUrl returns the FaroEndpointUrl field if non-nil, zero value otherwise.

### GetFaroEndpointUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetFaroEndpointUrlOk() (*string, bool)`

GetFaroEndpointUrlOk returns a tuple with the FaroEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaroEndpointUrl

`func (o *FormattedApiStackRegionAnyOf1) SetFaroEndpointUrl(v string)`

SetFaroEndpointUrl sets FaroEndpointUrl field to given value.


### GetPdcClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetPdcClusterSlug() string`

GetPdcClusterSlug returns the PdcClusterSlug field if non-nil, zero value otherwise.

### GetPdcClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetPdcClusterSlugOk() (*string, bool)`

GetPdcClusterSlugOk returns a tuple with the PdcClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdcClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetPdcClusterSlug(v string)`

SetPdcClusterSlug sets PdcClusterSlug field to given value.


### GetPdcPrivateConnectivityInfo

`func (o *FormattedApiStackRegionAnyOf1) GetPdcPrivateConnectivityInfo() PdcPrivateConnectivityInfo`

GetPdcPrivateConnectivityInfo returns the PdcPrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetPdcPrivateConnectivityInfoOk

`func (o *FormattedApiStackRegionAnyOf1) GetPdcPrivateConnectivityInfoOk() (*PdcPrivateConnectivityInfo, bool)`

GetPdcPrivateConnectivityInfoOk returns a tuple with the PdcPrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdcPrivateConnectivityInfo

`func (o *FormattedApiStackRegionAnyOf1) SetPdcPrivateConnectivityInfo(v PdcPrivateConnectivityInfo)`

SetPdcPrivateConnectivityInfo sets PdcPrivateConnectivityInfo field to given value.


### GetOtlpHttpUrl

`func (o *FormattedApiStackRegionAnyOf1) GetOtlpHttpUrl() string`

GetOtlpHttpUrl returns the OtlpHttpUrl field if non-nil, zero value otherwise.

### GetOtlpHttpUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetOtlpHttpUrlOk() (*string, bool)`

GetOtlpHttpUrlOk returns a tuple with the OtlpHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpHttpUrl

`func (o *FormattedApiStackRegionAnyOf1) SetOtlpHttpUrl(v string)`

SetOtlpHttpUrl sets OtlpHttpUrl field to given value.


### SetOtlpHttpUrlNil

`func (o *FormattedApiStackRegionAnyOf1) SetOtlpHttpUrlNil(b bool)`

 SetOtlpHttpUrlNil sets the value for OtlpHttpUrl to be an explicit nil

### UnsetOtlpHttpUrl
`func (o *FormattedApiStackRegionAnyOf1) UnsetOtlpHttpUrl()`

UnsetOtlpHttpUrl ensures that no value is present for OtlpHttpUrl, not even an explicit nil
### GetOtlpPrivateConnectivityInfo

`func (o *FormattedApiStackRegionAnyOf1) GetOtlpPrivateConnectivityInfo() OtlpPrivateConnectivityInfo`

GetOtlpPrivateConnectivityInfo returns the OtlpPrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetOtlpPrivateConnectivityInfoOk

`func (o *FormattedApiStackRegionAnyOf1) GetOtlpPrivateConnectivityInfoOk() (*OtlpPrivateConnectivityInfo, bool)`

GetOtlpPrivateConnectivityInfoOk returns a tuple with the OtlpPrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpPrivateConnectivityInfo

`func (o *FormattedApiStackRegionAnyOf1) SetOtlpPrivateConnectivityInfo(v OtlpPrivateConnectivityInfo)`

SetOtlpPrivateConnectivityInfo sets OtlpPrivateConnectivityInfo field to given value.


### GetAuthApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetAuthApiUrl() string`

GetAuthApiUrl returns the AuthApiUrl field if non-nil, zero value otherwise.

### GetAuthApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetAuthApiUrlOk() (*string, bool)`

GetAuthApiUrlOk returns a tuple with the AuthApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetAuthApiUrl(v string)`

SetAuthApiUrl sets AuthApiUrl field to given value.


### GetAuthApiTokenSet

`func (o *FormattedApiStackRegionAnyOf1) GetAuthApiTokenSet() bool`

GetAuthApiTokenSet returns the AuthApiTokenSet field if non-nil, zero value otherwise.

### GetAuthApiTokenSetOk

`func (o *FormattedApiStackRegionAnyOf1) GetAuthApiTokenSetOk() (*bool, bool)`

GetAuthApiTokenSetOk returns a tuple with the AuthApiTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthApiTokenSet

`func (o *FormattedApiStackRegionAnyOf1) SetAuthApiTokenSet(v bool)`

SetAuthApiTokenSet sets AuthApiTokenSet field to given value.


### GetHgClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterId() float32`

GetHgClusterId returns the HgClusterId field if non-nil, zero value otherwise.

### GetHgClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterIdOk() (*float32, bool)`

GetHgClusterIdOk returns a tuple with the HgClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHgClusterId(v float32)`

SetHgClusterId sets HgClusterId field to given value.


### GetHgClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterSlug() string`

GetHgClusterSlug returns the HgClusterSlug field if non-nil, zero value otherwise.

### GetHgClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterSlugOk() (*string, bool)`

GetHgClusterSlugOk returns a tuple with the HgClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHgClusterSlug(v string)`

SetHgClusterSlug sets HgClusterSlug field to given value.


### GetHgClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterName() string`

GetHgClusterName returns the HgClusterName field if non-nil, zero value otherwise.

### GetHgClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterNameOk() (*string, bool)`

GetHgClusterNameOk returns a tuple with the HgClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHgClusterName(v string)`

SetHgClusterName sets HgClusterName field to given value.


### GetHgClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterUrl() string`

GetHgClusterUrl returns the HgClusterUrl field if non-nil, zero value otherwise.

### GetHgClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHgClusterUrlOk() (*string, bool)`

GetHgClusterUrlOk returns a tuple with the HgClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHgClusterUrl(v string)`

SetHgClusterUrl sets HgClusterUrl field to given value.


### GetHmPromClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterId() float32`

GetHmPromClusterId returns the HmPromClusterId field if non-nil, zero value otherwise.

### GetHmPromClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterIdOk() (*float32, bool)`

GetHmPromClusterIdOk returns a tuple with the HmPromClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHmPromClusterId(v float32)`

SetHmPromClusterId sets HmPromClusterId field to given value.


### GetHmPromClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterSlug() string`

GetHmPromClusterSlug returns the HmPromClusterSlug field if non-nil, zero value otherwise.

### GetHmPromClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterSlugOk() (*string, bool)`

GetHmPromClusterSlugOk returns a tuple with the HmPromClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHmPromClusterSlug(v string)`

SetHmPromClusterSlug sets HmPromClusterSlug field to given value.


### GetHmPromClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterName() string`

GetHmPromClusterName returns the HmPromClusterName field if non-nil, zero value otherwise.

### GetHmPromClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterNameOk() (*string, bool)`

GetHmPromClusterNameOk returns a tuple with the HmPromClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHmPromClusterName(v string)`

SetHmPromClusterName sets HmPromClusterName field to given value.


### GetHmPromClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterUrl() string`

GetHmPromClusterUrl returns the HmPromClusterUrl field if non-nil, zero value otherwise.

### GetHmPromClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmPromClusterUrlOk() (*string, bool)`

GetHmPromClusterUrlOk returns a tuple with the HmPromClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHmPromClusterUrl(v string)`

SetHmPromClusterUrl sets HmPromClusterUrl field to given value.


### GetHmGraphiteClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterId() float32`

GetHmGraphiteClusterId returns the HmGraphiteClusterId field if non-nil, zero value otherwise.

### GetHmGraphiteClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterIdOk() (*float32, bool)`

GetHmGraphiteClusterIdOk returns a tuple with the HmGraphiteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHmGraphiteClusterId(v float32)`

SetHmGraphiteClusterId sets HmGraphiteClusterId field to given value.


### GetHmGraphiteClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterSlug() string`

GetHmGraphiteClusterSlug returns the HmGraphiteClusterSlug field if non-nil, zero value otherwise.

### GetHmGraphiteClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterSlugOk() (*string, bool)`

GetHmGraphiteClusterSlugOk returns a tuple with the HmGraphiteClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHmGraphiteClusterSlug(v string)`

SetHmGraphiteClusterSlug sets HmGraphiteClusterSlug field to given value.


### GetHmGraphiteClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterName() string`

GetHmGraphiteClusterName returns the HmGraphiteClusterName field if non-nil, zero value otherwise.

### GetHmGraphiteClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterNameOk() (*string, bool)`

GetHmGraphiteClusterNameOk returns a tuple with the HmGraphiteClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHmGraphiteClusterName(v string)`

SetHmGraphiteClusterName sets HmGraphiteClusterName field to given value.


### GetHmGraphiteClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterUrl() string`

GetHmGraphiteClusterUrl returns the HmGraphiteClusterUrl field if non-nil, zero value otherwise.

### GetHmGraphiteClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHmGraphiteClusterUrlOk() (*string, bool)`

GetHmGraphiteClusterUrlOk returns a tuple with the HmGraphiteClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHmGraphiteClusterUrl(v string)`

SetHmGraphiteClusterUrl sets HmGraphiteClusterUrl field to given value.


### GetHlClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterId() float32`

GetHlClusterId returns the HlClusterId field if non-nil, zero value otherwise.

### GetHlClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterIdOk() (*float32, bool)`

GetHlClusterIdOk returns a tuple with the HlClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHlClusterId(v float32)`

SetHlClusterId sets HlClusterId field to given value.


### GetHlClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterSlug() string`

GetHlClusterSlug returns the HlClusterSlug field if non-nil, zero value otherwise.

### GetHlClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterSlugOk() (*string, bool)`

GetHlClusterSlugOk returns a tuple with the HlClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHlClusterSlug(v string)`

SetHlClusterSlug sets HlClusterSlug field to given value.


### GetHlClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterName() string`

GetHlClusterName returns the HlClusterName field if non-nil, zero value otherwise.

### GetHlClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterNameOk() (*string, bool)`

GetHlClusterNameOk returns a tuple with the HlClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHlClusterName(v string)`

SetHlClusterName sets HlClusterName field to given value.


### GetHlClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterUrl() string`

GetHlClusterUrl returns the HlClusterUrl field if non-nil, zero value otherwise.

### GetHlClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHlClusterUrlOk() (*string, bool)`

GetHlClusterUrlOk returns a tuple with the HlClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHlClusterUrl(v string)`

SetHlClusterUrl sets HlClusterUrl field to given value.


### GetAmClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterId() float32`

GetAmClusterId returns the AmClusterId field if non-nil, zero value otherwise.

### GetAmClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterIdOk() (*float32, bool)`

GetAmClusterIdOk returns a tuple with the AmClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetAmClusterId(v float32)`

SetAmClusterId sets AmClusterId field to given value.


### GetAmClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterSlug() string`

GetAmClusterSlug returns the AmClusterSlug field if non-nil, zero value otherwise.

### GetAmClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterSlugOk() (*string, bool)`

GetAmClusterSlugOk returns a tuple with the AmClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetAmClusterSlug(v string)`

SetAmClusterSlug sets AmClusterSlug field to given value.


### GetAmClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterName() string`

GetAmClusterName returns the AmClusterName field if non-nil, zero value otherwise.

### GetAmClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterNameOk() (*string, bool)`

GetAmClusterNameOk returns a tuple with the AmClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetAmClusterName(v string)`

SetAmClusterName sets AmClusterName field to given value.


### GetAmClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterUrl() string`

GetAmClusterUrl returns the AmClusterUrl field if non-nil, zero value otherwise.

### GetAmClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetAmClusterUrlOk() (*string, bool)`

GetAmClusterUrlOk returns a tuple with the AmClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetAmClusterUrl(v string)`

SetAmClusterUrl sets AmClusterUrl field to given value.


### GetHtClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterId() float32`

GetHtClusterId returns the HtClusterId field if non-nil, zero value otherwise.

### GetHtClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterIdOk() (*float32, bool)`

GetHtClusterIdOk returns a tuple with the HtClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHtClusterId(v float32)`

SetHtClusterId sets HtClusterId field to given value.


### GetHtClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterSlug() string`

GetHtClusterSlug returns the HtClusterSlug field if non-nil, zero value otherwise.

### GetHtClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterSlugOk() (*string, bool)`

GetHtClusterSlugOk returns a tuple with the HtClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHtClusterSlug(v string)`

SetHtClusterSlug sets HtClusterSlug field to given value.


### GetHtClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterName() string`

GetHtClusterName returns the HtClusterName field if non-nil, zero value otherwise.

### GetHtClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterNameOk() (*string, bool)`

GetHtClusterNameOk returns a tuple with the HtClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHtClusterName(v string)`

SetHtClusterName sets HtClusterName field to given value.


### GetHtClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterUrl() string`

GetHtClusterUrl returns the HtClusterUrl field if non-nil, zero value otherwise.

### GetHtClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHtClusterUrlOk() (*string, bool)`

GetHtClusterUrlOk returns a tuple with the HtClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHtClusterUrl(v string)`

SetHtClusterUrl sets HtClusterUrl field to given value.


### GetHpClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterId() float32`

GetHpClusterId returns the HpClusterId field if non-nil, zero value otherwise.

### GetHpClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterIdOk() (*float32, bool)`

GetHpClusterIdOk returns a tuple with the HpClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetHpClusterId(v float32)`

SetHpClusterId sets HpClusterId field to given value.


### GetHpClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterSlug() string`

GetHpClusterSlug returns the HpClusterSlug field if non-nil, zero value otherwise.

### GetHpClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterSlugOk() (*string, bool)`

GetHpClusterSlugOk returns a tuple with the HpClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetHpClusterSlug(v string)`

SetHpClusterSlug sets HpClusterSlug field to given value.


### GetHpClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterName() string`

GetHpClusterName returns the HpClusterName field if non-nil, zero value otherwise.

### GetHpClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterNameOk() (*string, bool)`

GetHpClusterNameOk returns a tuple with the HpClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetHpClusterName(v string)`

SetHpClusterName sets HpClusterName field to given value.


### GetHpClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterUrl() string`

GetHpClusterUrl returns the HpClusterUrl field if non-nil, zero value otherwise.

### GetHpClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetHpClusterUrlOk() (*string, bool)`

GetHpClusterUrlOk returns a tuple with the HpClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetHpClusterUrl(v string)`

SetHpClusterUrl sets HpClusterUrl field to given value.


### GetAgmClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterId() float32`

GetAgmClusterId returns the AgmClusterId field if non-nil, zero value otherwise.

### GetAgmClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterIdOk() (*float32, bool)`

GetAgmClusterIdOk returns a tuple with the AgmClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetAgmClusterId(v float32)`

SetAgmClusterId sets AgmClusterId field to given value.


### GetAgmClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterSlug() string`

GetAgmClusterSlug returns the AgmClusterSlug field if non-nil, zero value otherwise.

### GetAgmClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterSlugOk() (*string, bool)`

GetAgmClusterSlugOk returns a tuple with the AgmClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetAgmClusterSlug(v string)`

SetAgmClusterSlug sets AgmClusterSlug field to given value.


### GetAgmClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterName() string`

GetAgmClusterName returns the AgmClusterName field if non-nil, zero value otherwise.

### GetAgmClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterNameOk() (*string, bool)`

GetAgmClusterNameOk returns a tuple with the AgmClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetAgmClusterName(v string)`

SetAgmClusterName sets AgmClusterName field to given value.


### GetAgmClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterUrl() string`

GetAgmClusterUrl returns the AgmClusterUrl field if non-nil, zero value otherwise.

### GetAgmClusterUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetAgmClusterUrlOk() (*string, bool)`

GetAgmClusterUrlOk returns a tuple with the AgmClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgmClusterUrl

`func (o *FormattedApiStackRegionAnyOf1) SetAgmClusterUrl(v string)`

SetAgmClusterUrl sets AgmClusterUrl field to given value.


### GetAssertsGraphClusterId

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterId() float32`

GetAssertsGraphClusterId returns the AssertsGraphClusterId field if non-nil, zero value otherwise.

### GetAssertsGraphClusterIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterIdOk() (*float32, bool)`

GetAssertsGraphClusterIdOk returns a tuple with the AssertsGraphClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertsGraphClusterId

`func (o *FormattedApiStackRegionAnyOf1) SetAssertsGraphClusterId(v float32)`

SetAssertsGraphClusterId sets AssertsGraphClusterId field to given value.


### GetAssertsGraphClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterSlug() string`

GetAssertsGraphClusterSlug returns the AssertsGraphClusterSlug field if non-nil, zero value otherwise.

### GetAssertsGraphClusterSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterSlugOk() (*string, bool)`

GetAssertsGraphClusterSlugOk returns a tuple with the AssertsGraphClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertsGraphClusterSlug

`func (o *FormattedApiStackRegionAnyOf1) SetAssertsGraphClusterSlug(v string)`

SetAssertsGraphClusterSlug sets AssertsGraphClusterSlug field to given value.


### GetAssertsGraphClusterName

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterName() string`

GetAssertsGraphClusterName returns the AssertsGraphClusterName field if non-nil, zero value otherwise.

### GetAssertsGraphClusterNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterNameOk() (*string, bool)`

GetAssertsGraphClusterNameOk returns a tuple with the AssertsGraphClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertsGraphClusterName

`func (o *FormattedApiStackRegionAnyOf1) SetAssertsGraphClusterName(v string)`

SetAssertsGraphClusterName sets AssertsGraphClusterName field to given value.


### GetAssertsGraphClusterApiUrl

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterApiUrl() string`

GetAssertsGraphClusterApiUrl returns the AssertsGraphClusterApiUrl field if non-nil, zero value otherwise.

### GetAssertsGraphClusterApiUrlOk

`func (o *FormattedApiStackRegionAnyOf1) GetAssertsGraphClusterApiUrlOk() (*string, bool)`

GetAssertsGraphClusterApiUrlOk returns a tuple with the AssertsGraphClusterApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertsGraphClusterApiUrl

`func (o *FormattedApiStackRegionAnyOf1) SetAssertsGraphClusterApiUrl(v string)`

SetAssertsGraphClusterApiUrl sets AssertsGraphClusterApiUrl field to given value.

### HasAssertsGraphClusterApiUrl

`func (o *FormattedApiStackRegionAnyOf1) HasAssertsGraphClusterApiUrl() bool`

HasAssertsGraphClusterApiUrl returns a boolean if a field has been set.

### SetAssertsGraphClusterApiUrlNil

`func (o *FormattedApiStackRegionAnyOf1) SetAssertsGraphClusterApiUrlNil(b bool)`

 SetAssertsGraphClusterApiUrlNil sets the value for AssertsGraphClusterApiUrl to be an explicit nil

### UnsetAssertsGraphClusterApiUrl
`func (o *FormattedApiStackRegionAnyOf1) UnsetAssertsGraphClusterApiUrl()`

UnsetAssertsGraphClusterApiUrl ensures that no value is present for AssertsGraphClusterApiUrl, not even an explicit nil
### GetProviderRegion

`func (o *FormattedApiStackRegionAnyOf1) GetProviderRegion() string`

GetProviderRegion returns the ProviderRegion field if non-nil, zero value otherwise.

### GetProviderRegionOk

`func (o *FormattedApiStackRegionAnyOf1) GetProviderRegionOk() (*string, bool)`

GetProviderRegionOk returns a tuple with the ProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRegion

`func (o *FormattedApiStackRegionAnyOf1) SetProviderRegion(v string)`

SetProviderRegion sets ProviderRegion field to given value.


### GetIsStub

`func (o *FormattedApiStackRegionAnyOf1) GetIsStub() bool`

GetIsStub returns the IsStub field if non-nil, zero value otherwise.

### GetIsStubOk

`func (o *FormattedApiStackRegionAnyOf1) GetIsStubOk() (*bool, bool)`

GetIsStubOk returns a tuple with the IsStub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStub

`func (o *FormattedApiStackRegionAnyOf1) SetIsStub(v bool)`

SetIsStub sets IsStub field to given value.


### GetGeoLocation

`func (o *FormattedApiStackRegionAnyOf1) GetGeoLocation() []float32`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *FormattedApiStackRegionAnyOf1) GetGeoLocationOk() (*[]float32, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *FormattedApiStackRegionAnyOf1) SetGeoLocation(v []float32)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *FormattedApiStackRegionAnyOf1) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCountryCode

`func (o *FormattedApiStackRegionAnyOf1) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *FormattedApiStackRegionAnyOf1) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *FormattedApiStackRegionAnyOf1) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *FormattedApiStackRegionAnyOf1) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetComplete

`func (o *FormattedApiStackRegionAnyOf1) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *FormattedApiStackRegionAnyOf1) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *FormattedApiStackRegionAnyOf1) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetId

`func (o *FormattedApiStackRegionAnyOf1) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiStackRegionAnyOf1) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiStackRegionAnyOf1) SetId(v float32)`

SetId sets Id field to given value.


### GetStatus

`func (o *FormattedApiStackRegionAnyOf1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormattedApiStackRegionAnyOf1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormattedApiStackRegionAnyOf1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVisibility

`func (o *FormattedApiStackRegionAnyOf1) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormattedApiStackRegionAnyOf1) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormattedApiStackRegionAnyOf1) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSlug

`func (o *FormattedApiStackRegionAnyOf1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormattedApiStackRegionAnyOf1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormattedApiStackRegionAnyOf1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *FormattedApiStackRegionAnyOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiStackRegionAnyOf1) SetName(v string)`

SetName sets Name field to given value.


### GetPublicName

`func (o *FormattedApiStackRegionAnyOf1) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *FormattedApiStackRegionAnyOf1) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *FormattedApiStackRegionAnyOf1) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.


### GetDescription

`func (o *FormattedApiStackRegionAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormattedApiStackRegionAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormattedApiStackRegionAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProvider

`func (o *FormattedApiStackRegionAnyOf1) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FormattedApiStackRegionAnyOf1) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FormattedApiStackRegionAnyOf1) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetCreatedAt

`func (o *FormattedApiStackRegionAnyOf1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiStackRegionAnyOf1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiStackRegionAnyOf1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FormattedApiStackRegionAnyOf1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiStackRegionAnyOf1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiStackRegionAnyOf1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FormattedApiStackRegionAnyOf1) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiStackRegionAnyOf1) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSortOrder

`func (o *FormattedApiStackRegionAnyOf1) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FormattedApiStackRegionAnyOf1) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FormattedApiStackRegionAnyOf1) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


