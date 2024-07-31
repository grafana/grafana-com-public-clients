# GrafanaServiceAccountDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | Pointer to **map[string]bool** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int64** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to **int64** |  | [optional] 

## Methods

### NewGrafanaServiceAccountDTO

`func NewGrafanaServiceAccountDTO() *GrafanaServiceAccountDTO`

NewGrafanaServiceAccountDTO instantiates a new GrafanaServiceAccountDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaServiceAccountDTOWithDefaults

`func NewGrafanaServiceAccountDTOWithDefaults() *GrafanaServiceAccountDTO`

NewGrafanaServiceAccountDTOWithDefaults instantiates a new GrafanaServiceAccountDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *GrafanaServiceAccountDTO) GetAccessControl() map[string]bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *GrafanaServiceAccountDTO) GetAccessControlOk() (*map[string]bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *GrafanaServiceAccountDTO) SetAccessControl(v map[string]bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *GrafanaServiceAccountDTO) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GrafanaServiceAccountDTO) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GrafanaServiceAccountDTO) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GrafanaServiceAccountDTO) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GrafanaServiceAccountDTO) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetId

`func (o *GrafanaServiceAccountDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrafanaServiceAccountDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrafanaServiceAccountDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GrafanaServiceAccountDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDisabled

`func (o *GrafanaServiceAccountDTO) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *GrafanaServiceAccountDTO) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *GrafanaServiceAccountDTO) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *GrafanaServiceAccountDTO) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsExternal

`func (o *GrafanaServiceAccountDTO) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *GrafanaServiceAccountDTO) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *GrafanaServiceAccountDTO) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *GrafanaServiceAccountDTO) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetLogin

`func (o *GrafanaServiceAccountDTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GrafanaServiceAccountDTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GrafanaServiceAccountDTO) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GrafanaServiceAccountDTO) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *GrafanaServiceAccountDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrafanaServiceAccountDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrafanaServiceAccountDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GrafanaServiceAccountDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *GrafanaServiceAccountDTO) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GrafanaServiceAccountDTO) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GrafanaServiceAccountDTO) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GrafanaServiceAccountDTO) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRole

`func (o *GrafanaServiceAccountDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GrafanaServiceAccountDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GrafanaServiceAccountDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GrafanaServiceAccountDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTokens

`func (o *GrafanaServiceAccountDTO) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GrafanaServiceAccountDTO) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GrafanaServiceAccountDTO) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *GrafanaServiceAccountDTO) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


