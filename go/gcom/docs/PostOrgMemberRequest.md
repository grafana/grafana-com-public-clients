# PostOrgMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to **int32** |  | [optional] 
**DefaultOrg** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewPostOrgMemberRequest

`func NewPostOrgMemberRequest() *PostOrgMemberRequest`

NewPostOrgMemberRequest instantiates a new PostOrgMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOrgMemberRequestWithDefaults

`func NewPostOrgMemberRequestWithDefaults() *PostOrgMemberRequest`

NewPostOrgMemberRequestWithDefaults instantiates a new PostOrgMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *PostOrgMemberRequest) GetBilling() int32`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *PostOrgMemberRequest) GetBillingOk() (*int32, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *PostOrgMemberRequest) SetBilling(v int32)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *PostOrgMemberRequest) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetDefaultOrg

`func (o *PostOrgMemberRequest) GetDefaultOrg() string`

GetDefaultOrg returns the DefaultOrg field if non-nil, zero value otherwise.

### GetDefaultOrgOk

`func (o *PostOrgMemberRequest) GetDefaultOrgOk() (*string, bool)`

GetDefaultOrgOk returns a tuple with the DefaultOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrg

`func (o *PostOrgMemberRequest) SetDefaultOrg(v string)`

SetDefaultOrg sets DefaultOrg field to given value.

### HasDefaultOrg

`func (o *PostOrgMemberRequest) HasDefaultOrg() bool`

HasDefaultOrg returns a boolean if a field has been set.

### GetPrivacy

`func (o *PostOrgMemberRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *PostOrgMemberRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *PostOrgMemberRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *PostOrgMemberRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRole

`func (o *PostOrgMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostOrgMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostOrgMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostOrgMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


