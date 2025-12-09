# Logos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Small** | **string** | Link to the \&quot;small\&quot; version of the plugin logo, which must be an SVG image. \&quot;Large\&quot; and \&quot;small\&quot; logos can be the same image. | 
**Large** | **string** | Link to the \&quot;large\&quot; version of the plugin logo, which must be an SVG image. \&quot;Large\&quot; and \&quot;small\&quot; logos can be the same image. | 

## Methods

### NewLogos

`func NewLogos(small string, large string, ) *Logos`

NewLogos instantiates a new Logos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogosWithDefaults

`func NewLogosWithDefaults() *Logos`

NewLogosWithDefaults instantiates a new Logos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmall

`func (o *Logos) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *Logos) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *Logos) SetSmall(v string)`

SetSmall sets Small field to given value.


### GetLarge

`func (o *Logos) GetLarge() string`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *Logos) GetLargeOk() (*string, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *Logos) SetLarge(v string)`

SetLarge sets Large field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


