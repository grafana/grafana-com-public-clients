# ItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**ItemsInnerId**](ItemsInnerId.md) |  | 
**AccessPolicyId** | **string** |  | 
**OrgId** | **float32** |  | 
**OrgSlug** | **string** |  | 
**OrgName** | **string** |  | 
**InstanceId** | **NullableFloat32** |  | 
**Name** | **string** |  | 
**Role** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 
**FirstUsed** | **NullableString** |  | 
**Token** | Pointer to **string** |  | [optional] 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewItemsInner

`func NewItemsInner(id ItemsInnerId, accessPolicyId string, orgId float32, orgSlug string, orgName string, instanceId NullableFloat32, name string, role string, createdAt string, updatedAt NullableString, firstUsed NullableString, links []LinksInner, ) *ItemsInner`

NewItemsInner instantiates a new ItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsInnerWithDefaults

`func NewItemsInnerWithDefaults() *ItemsInner`

NewItemsInnerWithDefaults instantiates a new ItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemsInner) GetId() ItemsInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemsInner) GetIdOk() (*ItemsInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemsInner) SetId(v ItemsInnerId)`

SetId sets Id field to given value.


### GetAccessPolicyId

`func (o *ItemsInner) GetAccessPolicyId() string`

GetAccessPolicyId returns the AccessPolicyId field if non-nil, zero value otherwise.

### GetAccessPolicyIdOk

`func (o *ItemsInner) GetAccessPolicyIdOk() (*string, bool)`

GetAccessPolicyIdOk returns a tuple with the AccessPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyId

`func (o *ItemsInner) SetAccessPolicyId(v string)`

SetAccessPolicyId sets AccessPolicyId field to given value.


### GetOrgId

`func (o *ItemsInner) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ItemsInner) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ItemsInner) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.


### GetOrgSlug

`func (o *ItemsInner) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *ItemsInner) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *ItemsInner) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetOrgName

`func (o *ItemsInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ItemsInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ItemsInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetInstanceId

`func (o *ItemsInner) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ItemsInner) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ItemsInner) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### SetInstanceIdNil

`func (o *ItemsInner) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ItemsInner) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetName

`func (o *ItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ItemsInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *ItemsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItemsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItemsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ItemsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItemsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItemsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ItemsInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ItemsInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFirstUsed

`func (o *ItemsInner) GetFirstUsed() string`

GetFirstUsed returns the FirstUsed field if non-nil, zero value otherwise.

### GetFirstUsedOk

`func (o *ItemsInner) GetFirstUsedOk() (*string, bool)`

GetFirstUsedOk returns a tuple with the FirstUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUsed

`func (o *ItemsInner) SetFirstUsed(v string)`

SetFirstUsed sets FirstUsed field to given value.


### SetFirstUsedNil

`func (o *ItemsInner) SetFirstUsedNil(b bool)`

 SetFirstUsedNil sets the value for FirstUsed to be an explicit nil

### UnsetFirstUsed
`func (o *ItemsInner) UnsetFirstUsed()`

UnsetFirstUsed ensures that no value is present for FirstUsed, not even an explicit nil
### GetToken

`func (o *ItemsInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ItemsInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ItemsInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ItemsInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLinks

`func (o *ItemsInner) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ItemsInner) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ItemsInner) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


