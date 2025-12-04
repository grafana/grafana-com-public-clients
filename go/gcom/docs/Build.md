# Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **float32** | Time when the plugin was built, as a Unix timestamp. | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** | Git branch the plugin was built from. | [optional] 
**Hash** | Pointer to **string** | Git hash of the commit the plugin was built from | [optional] 
**Number** | Pointer to **float32** |  | [optional] 
**Pr** | Pointer to **float32** | GitHub pull request the plugin was built from | [optional] 
**Build** | Pointer to **float32** | Build job number used to build this plugin. | [optional] 

## Methods

### NewBuild

`func NewBuild() *Build`

NewBuild instantiates a new Build object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildWithDefaults

`func NewBuildWithDefaults() *Build`

NewBuildWithDefaults instantiates a new Build object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *Build) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Build) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Build) SetTime(v float32)`

SetTime sets Time field to given value.

### HasTime

`func (o *Build) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRepo

`func (o *Build) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Build) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Build) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Build) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetBranch

`func (o *Build) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Build) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Build) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Build) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetHash

`func (o *Build) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Build) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Build) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Build) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetNumber

`func (o *Build) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Build) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Build) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Build) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPr

`func (o *Build) GetPr() float32`

GetPr returns the Pr field if non-nil, zero value otherwise.

### GetPrOk

`func (o *Build) GetPrOk() (*float32, bool)`

GetPrOk returns a tuple with the Pr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPr

`func (o *Build) SetPr(v float32)`

SetPr sets Pr field to given value.

### HasPr

`func (o *Build) HasPr() bool`

HasPr returns a boolean if a field has been set.

### GetBuild

`func (o *Build) GetBuild() float32`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *Build) GetBuildOk() (*float32, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *Build) SetBuild(v float32)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *Build) HasBuild() bool`

HasBuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


