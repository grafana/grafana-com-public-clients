# StackCheckRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | org owning the stack | [optional] 
**Url** | **string** | url for the stack | 

## Methods

### NewStackCheckRequestV1

`func NewStackCheckRequestV1(url string, ) *StackCheckRequestV1`

NewStackCheckRequestV1 instantiates a new StackCheckRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackCheckRequestV1WithDefaults

`func NewStackCheckRequestV1WithDefaults() *StackCheckRequestV1`

NewStackCheckRequestV1WithDefaults instantiates a new StackCheckRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *StackCheckRequestV1) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StackCheckRequestV1) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StackCheckRequestV1) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *StackCheckRequestV1) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUrl

`func (o *StackCheckRequestV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackCheckRequestV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackCheckRequestV1) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


