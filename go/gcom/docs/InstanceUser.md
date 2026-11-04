# InstanceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | when the user was created | 
**Email** | **string** | email of the user | 
**Id** | **int64** | id of the user | 
**IsServiceAccount** | **bool** | whether the user is a service account | 
**LastSeenAt** | **NullableString** | when the user was last seen, null if never | 
**Login** | **string** | login/username of the user | 
**Name** | **string** | full name of the user | 
**Role** | **string** | role of the user on the instance | 
**UpdatedAt** | **string** | when the user was last updated | 

## Methods

### NewInstanceUser

`func NewInstanceUser(createdAt string, email string, id int64, isServiceAccount bool, lastSeenAt NullableString, login string, name string, role string, updatedAt string, ) *InstanceUser`

NewInstanceUser instantiates a new InstanceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUserWithDefaults

`func NewInstanceUserWithDefaults() *InstanceUser`

NewInstanceUserWithDefaults instantiates a new InstanceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InstanceUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *InstanceUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InstanceUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InstanceUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *InstanceUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceUser) SetId(v int64)`

SetId sets Id field to given value.


### GetIsServiceAccount

`func (o *InstanceUser) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *InstanceUser) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *InstanceUser) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.


### GetLastSeenAt

`func (o *InstanceUser) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *InstanceUser) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *InstanceUser) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.


### SetLastSeenAtNil

`func (o *InstanceUser) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *InstanceUser) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetLogin

`func (o *InstanceUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *InstanceUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *InstanceUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *InstanceUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceUser) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *InstanceUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InstanceUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InstanceUser) SetRole(v string)`

SetRole sets Role field to given value.


### GetUpdatedAt

`func (o *InstanceUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


