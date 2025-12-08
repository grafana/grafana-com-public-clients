# RolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**Grants** | Pointer to **[]string** | Default assignments of the role to Grafana basic roles (&#x60;Viewer&#x60;, &#x60;Editor&#x60;, &#x60;Admin&#x60;, &#x60;Grafana Admin&#x60;). | [optional] 

## Methods

### NewRolesInner

`func NewRolesInner() *RolesInner`

NewRolesInner instantiates a new RolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesInnerWithDefaults

`func NewRolesInnerWithDefaults() *RolesInner`

NewRolesInnerWithDefaults instantiates a new RolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RolesInner) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RolesInner) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RolesInner) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *RolesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetGrants

`func (o *RolesInner) GetGrants() []string`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *RolesInner) GetGrantsOk() (*[]string, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *RolesInner) SetGrants(v []string)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *RolesInner) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


