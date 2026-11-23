# ServiceAccountProfileDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | Pointer to **map[string]bool** | resolved access control permissions | [optional] 
**AvatarUrl** | **string** | the service account&#39;s avatar url | 
**CreatedAt** | **string** | when the service account was created (RFC3339) | 
**Id** | **int64** | service account id | 
**IsDisabled** | **bool** | whether the service account is disabled | 
**IsExternal** | Pointer to **bool** | whether the service account is externally managed | [optional] 
**Login** | **string** | service account login | 
**Name** | **string** | service account name | 
**OrgId** | **int64** | grafana org id the service account belongs to | 
**RequiredBy** | Pointer to **string** | the external service requiring this service account | [optional] 
**Role** | **string** | the service account&#39;s role | 
**Teams** | **[]string** | names of the teams the service account belongs to | 
**Tokens** | Pointer to **int64** | number of tokens on the service account | [optional] 
**Uid** | Pointer to **string** | service account uid | [optional] 
**UpdatedAt** | **string** | when the service account was last updated (RFC3339) | 

## Methods

### NewServiceAccountProfileDTO

`func NewServiceAccountProfileDTO(avatarUrl string, createdAt string, id int64, isDisabled bool, login string, name string, orgId int64, role string, teams []string, updatedAt string, ) *ServiceAccountProfileDTO`

NewServiceAccountProfileDTO instantiates a new ServiceAccountProfileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountProfileDTOWithDefaults

`func NewServiceAccountProfileDTOWithDefaults() *ServiceAccountProfileDTO`

NewServiceAccountProfileDTOWithDefaults instantiates a new ServiceAccountProfileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *ServiceAccountProfileDTO) GetAccessControl() map[string]bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ServiceAccountProfileDTO) GetAccessControlOk() (*map[string]bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ServiceAccountProfileDTO) SetAccessControl(v map[string]bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ServiceAccountProfileDTO) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *ServiceAccountProfileDTO) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ServiceAccountProfileDTO) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ServiceAccountProfileDTO) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetCreatedAt

`func (o *ServiceAccountProfileDTO) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountProfileDTO) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountProfileDTO) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ServiceAccountProfileDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountProfileDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountProfileDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetIsDisabled

`func (o *ServiceAccountProfileDTO) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ServiceAccountProfileDTO) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ServiceAccountProfileDTO) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetIsExternal

`func (o *ServiceAccountProfileDTO) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ServiceAccountProfileDTO) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ServiceAccountProfileDTO) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ServiceAccountProfileDTO) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetLogin

`func (o *ServiceAccountProfileDTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ServiceAccountProfileDTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ServiceAccountProfileDTO) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *ServiceAccountProfileDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountProfileDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountProfileDTO) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ServiceAccountProfileDTO) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ServiceAccountProfileDTO) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ServiceAccountProfileDTO) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.


### GetRequiredBy

`func (o *ServiceAccountProfileDTO) GetRequiredBy() string`

GetRequiredBy returns the RequiredBy field if non-nil, zero value otherwise.

### GetRequiredByOk

`func (o *ServiceAccountProfileDTO) GetRequiredByOk() (*string, bool)`

GetRequiredByOk returns a tuple with the RequiredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBy

`func (o *ServiceAccountProfileDTO) SetRequiredBy(v string)`

SetRequiredBy sets RequiredBy field to given value.

### HasRequiredBy

`func (o *ServiceAccountProfileDTO) HasRequiredBy() bool`

HasRequiredBy returns a boolean if a field has been set.

### GetRole

`func (o *ServiceAccountProfileDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceAccountProfileDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceAccountProfileDTO) SetRole(v string)`

SetRole sets Role field to given value.


### GetTeams

`func (o *ServiceAccountProfileDTO) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ServiceAccountProfileDTO) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ServiceAccountProfileDTO) SetTeams(v []string)`

SetTeams sets Teams field to given value.


### SetTeamsNil

`func (o *ServiceAccountProfileDTO) SetTeamsNil(b bool)`

 SetTeamsNil sets the value for Teams to be an explicit nil

### UnsetTeams
`func (o *ServiceAccountProfileDTO) UnsetTeams()`

UnsetTeams ensures that no value is present for Teams, not even an explicit nil
### GetTokens

`func (o *ServiceAccountProfileDTO) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ServiceAccountProfileDTO) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ServiceAccountProfileDTO) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ServiceAccountProfileDTO) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetUid

`func (o *ServiceAccountProfileDTO) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ServiceAccountProfileDTO) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ServiceAccountProfileDTO) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ServiceAccountProfileDTO) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceAccountProfileDTO) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceAccountProfileDTO) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceAccountProfileDTO) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


