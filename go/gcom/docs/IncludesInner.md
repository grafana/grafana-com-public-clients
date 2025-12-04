# IncludesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Unique identifier of the included resource | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **string** | (Legacy) The Angular component to use for a page. | [optional] 
**Role** | Pointer to **string** | The minimum role a user must have to see this page in the navigation menu. | [optional] 
**Action** | Pointer to **string** | The RBAC action a user must have to see this page in the navigation menu. **Warning**: unless the action targets the plugin, only the action is verified, not what it applies to. | [optional] 
**Path** | Pointer to **string** | Used for app plugins. | [optional] 
**AddToNav** | Pointer to **bool** | Add the include to the navigation menu. | [optional] 
**DefaultNav** | Pointer to **bool** | Page or dashboard when user clicks the icon in the side menu. | [optional] 
**Icon** | Pointer to **string** | Icon to use in the side menu. For information on available icon, refer to [Icons Overview](https://developers.grafana.com/ui/latest/index.html?path&#x3D;/story/docs-overview-icon--icons-overview). | [optional] 

## Methods

### NewIncludesInner

`func NewIncludesInner() *IncludesInner`

NewIncludesInner instantiates a new IncludesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludesInnerWithDefaults

`func NewIncludesInnerWithDefaults() *IncludesInner`

NewIncludesInnerWithDefaults instantiates a new IncludesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *IncludesInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IncludesInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IncludesInner) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IncludesInner) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *IncludesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncludesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IncludesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncludesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncludesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncludesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComponent

`func (o *IncludesInner) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IncludesInner) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IncludesInner) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IncludesInner) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetRole

`func (o *IncludesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IncludesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IncludesInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IncludesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAction

`func (o *IncludesInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IncludesInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IncludesInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IncludesInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPath

`func (o *IncludesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IncludesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IncludesInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IncludesInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAddToNav

`func (o *IncludesInner) GetAddToNav() bool`

GetAddToNav returns the AddToNav field if non-nil, zero value otherwise.

### GetAddToNavOk

`func (o *IncludesInner) GetAddToNavOk() (*bool, bool)`

GetAddToNavOk returns a tuple with the AddToNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToNav

`func (o *IncludesInner) SetAddToNav(v bool)`

SetAddToNav sets AddToNav field to given value.

### HasAddToNav

`func (o *IncludesInner) HasAddToNav() bool`

HasAddToNav returns a boolean if a field has been set.

### GetDefaultNav

`func (o *IncludesInner) GetDefaultNav() bool`

GetDefaultNav returns the DefaultNav field if non-nil, zero value otherwise.

### GetDefaultNavOk

`func (o *IncludesInner) GetDefaultNavOk() (*bool, bool)`

GetDefaultNavOk returns a tuple with the DefaultNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNav

`func (o *IncludesInner) SetDefaultNav(v bool)`

SetDefaultNav sets DefaultNav field to given value.

### HasDefaultNav

`func (o *IncludesInner) HasDefaultNav() bool`

HasDefaultNav returns a boolean if a field has been set.

### GetIcon

`func (o *IncludesInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IncludesInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IncludesInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IncludesInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


