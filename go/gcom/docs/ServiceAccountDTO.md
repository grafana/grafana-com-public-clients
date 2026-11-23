# ServiceAccountDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | Pointer to **map[string]bool** | resolved access control permissions | [optional] 
**AvatarUrl** | **string** | the service account&#39;s avatar url | 
**Id** | **int64** | service account id | 
**IsDisabled** | **bool** | whether the service account is disabled | 
**IsExternal** | Pointer to **bool** | whether the service account is externally managed | [optional] 
**Login** | **string** | service account login | 
**Name** | **string** | service account name | 
**OrgId** | **int64** | grafana org id the service account belongs to | 
**Role** | **string** | the service account&#39;s role | 
**Tokens** | **int64** | number of tokens on the service account | 
**Uid** | Pointer to **string** | service account uid | [optional] 

## Methods

### NewServiceAccountDTO

`func NewServiceAccountDTO(avatarUrl string, id int64, isDisabled bool, login string, name string, orgId int64, role string, tokens int64, ) *ServiceAccountDTO`

NewServiceAccountDTO instantiates a new ServiceAccountDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountDTOWithDefaults

`func NewServiceAccountDTOWithDefaults() *ServiceAccountDTO`

NewServiceAccountDTOWithDefaults instantiates a new ServiceAccountDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *ServiceAccountDTO) GetAccessControl() map[string]bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ServiceAccountDTO) GetAccessControlOk() (*map[string]bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ServiceAccountDTO) SetAccessControl(v map[string]bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ServiceAccountDTO) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *ServiceAccountDTO) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ServiceAccountDTO) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ServiceAccountDTO) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetId

`func (o *ServiceAccountDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetIsDisabled

`func (o *ServiceAccountDTO) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ServiceAccountDTO) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ServiceAccountDTO) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetIsExternal

`func (o *ServiceAccountDTO) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ServiceAccountDTO) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ServiceAccountDTO) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ServiceAccountDTO) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetLogin

`func (o *ServiceAccountDTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ServiceAccountDTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ServiceAccountDTO) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *ServiceAccountDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountDTO) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ServiceAccountDTO) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ServiceAccountDTO) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ServiceAccountDTO) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *ServiceAccountDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceAccountDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceAccountDTO) SetRole(v string)`

SetRole sets Role field to given value.


### GetTokens

`func (o *ServiceAccountDTO) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ServiceAccountDTO) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ServiceAccountDTO) SetTokens(v int64)`

SetTokens sets Tokens field to given value.


### GetUid

`func (o *ServiceAccountDTO) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ServiceAccountDTO) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ServiceAccountDTO) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ServiceAccountDTO) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


