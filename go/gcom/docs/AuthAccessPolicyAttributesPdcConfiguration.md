# AuthAccessPolicyAttributesPdcConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitedHosts** | Pointer to **[]string** | List of hosts that Grafana is able to access through a PDC. All hosts are allowed if not provided. | [optional] 

## Methods

### NewAuthAccessPolicyAttributesPdcConfiguration

`func NewAuthAccessPolicyAttributesPdcConfiguration() *AuthAccessPolicyAttributesPdcConfiguration`

NewAuthAccessPolicyAttributesPdcConfiguration instantiates a new AuthAccessPolicyAttributesPdcConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccessPolicyAttributesPdcConfigurationWithDefaults

`func NewAuthAccessPolicyAttributesPdcConfigurationWithDefaults() *AuthAccessPolicyAttributesPdcConfiguration`

NewAuthAccessPolicyAttributesPdcConfigurationWithDefaults instantiates a new AuthAccessPolicyAttributesPdcConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitedHosts

`func (o *AuthAccessPolicyAttributesPdcConfiguration) GetLimitedHosts() []string`

GetLimitedHosts returns the LimitedHosts field if non-nil, zero value otherwise.

### GetLimitedHostsOk

`func (o *AuthAccessPolicyAttributesPdcConfiguration) GetLimitedHostsOk() (*[]string, bool)`

GetLimitedHostsOk returns a tuple with the LimitedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedHosts

`func (o *AuthAccessPolicyAttributesPdcConfiguration) SetLimitedHosts(v []string)`

SetLimitedHosts sets LimitedHosts field to given value.

### HasLimitedHosts

`func (o *AuthAccessPolicyAttributesPdcConfiguration) HasLimitedHosts() bool`

HasLimitedHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


