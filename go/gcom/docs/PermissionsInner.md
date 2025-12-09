# PermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action, for example: &#x60;teams:read&#x60;. | [optional] 
**Scope** | Pointer to **string** | The scope that the plugin needs to access e.g: &#x60;teams:*&#x60;. | [optional] 

## Methods

### NewPermissionsInner

`func NewPermissionsInner() *PermissionsInner`

NewPermissionsInner instantiates a new PermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsInnerWithDefaults

`func NewPermissionsInnerWithDefaults() *PermissionsInner`

NewPermissionsInnerWithDefaults instantiates a new PermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PermissionsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PermissionsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PermissionsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PermissionsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetScope

`func (o *PermissionsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PermissionsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PermissionsInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PermissionsInner) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


