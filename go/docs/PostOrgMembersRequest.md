# PostOrgMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to **int32** |  | [optional] 
**Privacy** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewPostOrgMembersRequest

`func NewPostOrgMembersRequest(username string, ) *PostOrgMembersRequest`

NewPostOrgMembersRequest instantiates a new PostOrgMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOrgMembersRequestWithDefaults

`func NewPostOrgMembersRequestWithDefaults() *PostOrgMembersRequest`

NewPostOrgMembersRequestWithDefaults instantiates a new PostOrgMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *PostOrgMembersRequest) GetBilling() int32`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *PostOrgMembersRequest) GetBillingOk() (*int32, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *PostOrgMembersRequest) SetBilling(v int32)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *PostOrgMembersRequest) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetPrivacy

`func (o *PostOrgMembersRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *PostOrgMembersRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *PostOrgMembersRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *PostOrgMembersRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRole

`func (o *PostOrgMembersRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostOrgMembersRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostOrgMembersRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostOrgMembersRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUsername

`func (o *PostOrgMembersRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostOrgMembersRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostOrgMembersRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


