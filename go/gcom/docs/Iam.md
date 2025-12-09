# Iam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**[]PermissionsInner**](PermissionsInner.md) | Permissions are the permissions that the plugin needs its associated service account to have to query Grafana. | [optional] 

## Methods

### NewIam

`func NewIam() *Iam`

NewIam instantiates a new Iam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamWithDefaults

`func NewIamWithDefaults() *Iam`

NewIamWithDefaults instantiates a new Iam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *Iam) GetPermissions() []PermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Iam) GetPermissionsOk() (*[]PermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Iam) SetPermissions(v []PermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Iam) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


