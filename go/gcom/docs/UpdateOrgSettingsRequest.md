# UpdateOrgSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiDataSharingOptOut** | Pointer to **NullableBool** |  | [optional] 
**DisableTokenExpirationEmails** | Pointer to **NullableBool** |  | [optional] 
**MaxTokenExpirationDays** | Pointer to **NullableInt32** |  | [optional] 
**MfaAdminRecoveryOnly** | Pointer to **NullableBool** |  | [optional] 
**MfaRequired** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateOrgSettingsRequest

`func NewUpdateOrgSettingsRequest() *UpdateOrgSettingsRequest`

NewUpdateOrgSettingsRequest instantiates a new UpdateOrgSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgSettingsRequestWithDefaults

`func NewUpdateOrgSettingsRequestWithDefaults() *UpdateOrgSettingsRequest`

NewUpdateOrgSettingsRequestWithDefaults instantiates a new UpdateOrgSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiDataSharingOptOut

`func (o *UpdateOrgSettingsRequest) GetAiDataSharingOptOut() bool`

GetAiDataSharingOptOut returns the AiDataSharingOptOut field if non-nil, zero value otherwise.

### GetAiDataSharingOptOutOk

`func (o *UpdateOrgSettingsRequest) GetAiDataSharingOptOutOk() (*bool, bool)`

GetAiDataSharingOptOutOk returns a tuple with the AiDataSharingOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDataSharingOptOut

`func (o *UpdateOrgSettingsRequest) SetAiDataSharingOptOut(v bool)`

SetAiDataSharingOptOut sets AiDataSharingOptOut field to given value.

### HasAiDataSharingOptOut

`func (o *UpdateOrgSettingsRequest) HasAiDataSharingOptOut() bool`

HasAiDataSharingOptOut returns a boolean if a field has been set.

### SetAiDataSharingOptOutNil

`func (o *UpdateOrgSettingsRequest) SetAiDataSharingOptOutNil(b bool)`

 SetAiDataSharingOptOutNil sets the value for AiDataSharingOptOut to be an explicit nil

### UnsetAiDataSharingOptOut
`func (o *UpdateOrgSettingsRequest) UnsetAiDataSharingOptOut()`

UnsetAiDataSharingOptOut ensures that no value is present for AiDataSharingOptOut, not even an explicit nil
### GetDisableTokenExpirationEmails

`func (o *UpdateOrgSettingsRequest) GetDisableTokenExpirationEmails() bool`

GetDisableTokenExpirationEmails returns the DisableTokenExpirationEmails field if non-nil, zero value otherwise.

### GetDisableTokenExpirationEmailsOk

`func (o *UpdateOrgSettingsRequest) GetDisableTokenExpirationEmailsOk() (*bool, bool)`

GetDisableTokenExpirationEmailsOk returns a tuple with the DisableTokenExpirationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTokenExpirationEmails

`func (o *UpdateOrgSettingsRequest) SetDisableTokenExpirationEmails(v bool)`

SetDisableTokenExpirationEmails sets DisableTokenExpirationEmails field to given value.

### HasDisableTokenExpirationEmails

`func (o *UpdateOrgSettingsRequest) HasDisableTokenExpirationEmails() bool`

HasDisableTokenExpirationEmails returns a boolean if a field has been set.

### SetDisableTokenExpirationEmailsNil

`func (o *UpdateOrgSettingsRequest) SetDisableTokenExpirationEmailsNil(b bool)`

 SetDisableTokenExpirationEmailsNil sets the value for DisableTokenExpirationEmails to be an explicit nil

### UnsetDisableTokenExpirationEmails
`func (o *UpdateOrgSettingsRequest) UnsetDisableTokenExpirationEmails()`

UnsetDisableTokenExpirationEmails ensures that no value is present for DisableTokenExpirationEmails, not even an explicit nil
### GetMaxTokenExpirationDays

`func (o *UpdateOrgSettingsRequest) GetMaxTokenExpirationDays() int32`

GetMaxTokenExpirationDays returns the MaxTokenExpirationDays field if non-nil, zero value otherwise.

### GetMaxTokenExpirationDaysOk

`func (o *UpdateOrgSettingsRequest) GetMaxTokenExpirationDaysOk() (*int32, bool)`

GetMaxTokenExpirationDaysOk returns a tuple with the MaxTokenExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenExpirationDays

`func (o *UpdateOrgSettingsRequest) SetMaxTokenExpirationDays(v int32)`

SetMaxTokenExpirationDays sets MaxTokenExpirationDays field to given value.

### HasMaxTokenExpirationDays

`func (o *UpdateOrgSettingsRequest) HasMaxTokenExpirationDays() bool`

HasMaxTokenExpirationDays returns a boolean if a field has been set.

### SetMaxTokenExpirationDaysNil

`func (o *UpdateOrgSettingsRequest) SetMaxTokenExpirationDaysNil(b bool)`

 SetMaxTokenExpirationDaysNil sets the value for MaxTokenExpirationDays to be an explicit nil

### UnsetMaxTokenExpirationDays
`func (o *UpdateOrgSettingsRequest) UnsetMaxTokenExpirationDays()`

UnsetMaxTokenExpirationDays ensures that no value is present for MaxTokenExpirationDays, not even an explicit nil
### GetMfaAdminRecoveryOnly

`func (o *UpdateOrgSettingsRequest) GetMfaAdminRecoveryOnly() bool`

GetMfaAdminRecoveryOnly returns the MfaAdminRecoveryOnly field if non-nil, zero value otherwise.

### GetMfaAdminRecoveryOnlyOk

`func (o *UpdateOrgSettingsRequest) GetMfaAdminRecoveryOnlyOk() (*bool, bool)`

GetMfaAdminRecoveryOnlyOk returns a tuple with the MfaAdminRecoveryOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaAdminRecoveryOnly

`func (o *UpdateOrgSettingsRequest) SetMfaAdminRecoveryOnly(v bool)`

SetMfaAdminRecoveryOnly sets MfaAdminRecoveryOnly field to given value.

### HasMfaAdminRecoveryOnly

`func (o *UpdateOrgSettingsRequest) HasMfaAdminRecoveryOnly() bool`

HasMfaAdminRecoveryOnly returns a boolean if a field has been set.

### SetMfaAdminRecoveryOnlyNil

`func (o *UpdateOrgSettingsRequest) SetMfaAdminRecoveryOnlyNil(b bool)`

 SetMfaAdminRecoveryOnlyNil sets the value for MfaAdminRecoveryOnly to be an explicit nil

### UnsetMfaAdminRecoveryOnly
`func (o *UpdateOrgSettingsRequest) UnsetMfaAdminRecoveryOnly()`

UnsetMfaAdminRecoveryOnly ensures that no value is present for MfaAdminRecoveryOnly, not even an explicit nil
### GetMfaRequired

`func (o *UpdateOrgSettingsRequest) GetMfaRequired() bool`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *UpdateOrgSettingsRequest) GetMfaRequiredOk() (*bool, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *UpdateOrgSettingsRequest) SetMfaRequired(v bool)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *UpdateOrgSettingsRequest) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### SetMfaRequiredNil

`func (o *UpdateOrgSettingsRequest) SetMfaRequiredNil(b bool)`

 SetMfaRequiredNil sets the value for MfaRequired to be an explicit nil

### UnsetMfaRequired
`func (o *UpdateOrgSettingsRequest) UnsetMfaRequired()`

UnsetMfaRequired ensures that no value is present for MfaRequired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


