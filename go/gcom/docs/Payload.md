# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithAddons** | Pointer to **[]string** |  | [optional] 
**LicenseAllPlugins** | Pointer to **bool** |  | [optional] 
**CaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewPayload

`func NewPayload() *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithAddons

`func (o *Payload) GetWithAddons() []string`

GetWithAddons returns the WithAddons field if non-nil, zero value otherwise.

### GetWithAddonsOk

`func (o *Payload) GetWithAddonsOk() (*[]string, bool)`

GetWithAddonsOk returns a tuple with the WithAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAddons

`func (o *Payload) SetWithAddons(v []string)`

SetWithAddons sets WithAddons field to given value.

### HasWithAddons

`func (o *Payload) HasWithAddons() bool`

HasWithAddons returns a boolean if a field has been set.

### GetLicenseAllPlugins

`func (o *Payload) GetLicenseAllPlugins() bool`

GetLicenseAllPlugins returns the LicenseAllPlugins field if non-nil, zero value otherwise.

### GetLicenseAllPluginsOk

`func (o *Payload) GetLicenseAllPluginsOk() (*bool, bool)`

GetLicenseAllPluginsOk returns a tuple with the LicenseAllPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAllPlugins

`func (o *Payload) SetLicenseAllPlugins(v bool)`

SetLicenseAllPlugins sets LicenseAllPlugins field to given value.

### HasLicenseAllPlugins

`func (o *Payload) HasLicenseAllPlugins() bool`

HasLicenseAllPlugins returns a boolean if a field has been set.

### GetCaseId

`func (o *Payload) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *Payload) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *Payload) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *Payload) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


