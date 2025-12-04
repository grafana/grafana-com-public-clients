# LinksInner2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Display name of the link. Special names with predefined behavior: &lt;br/&gt;• &#x60;documentation&#x60; - sets Documentation link on plugins detail page&lt;br/&gt;• &#x60;repository&#x60; - used to determine and link to repository of the plugin&lt;br/&gt;• &#x60;license&#x60; - sets License link on plugins detail page&lt;br/&gt;• &#x60;raise issue&#x60; - sets &#x60;Raise an Issue&#x60; link on plugins detail page&lt;br/&gt;• &#x60;sponsorship&#x60; - sets &#x60;Sponsor this developer&#x60; link on plugins detail page to direct users to how they can support your work | [optional] 
**Url** | Pointer to **string** | URL value to use for this specific link. | [optional] 
**Target** | Pointer to **string** | A string that indicates where to display the linked resource | [optional] 

## Methods

### NewLinksInner2

`func NewLinksInner2() *LinksInner2`

NewLinksInner2 instantiates a new LinksInner2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksInner2WithDefaults

`func NewLinksInner2WithDefaults() *LinksInner2`

NewLinksInner2WithDefaults instantiates a new LinksInner2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinksInner2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinksInner2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinksInner2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinksInner2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *LinksInner2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinksInner2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinksInner2) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LinksInner2) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTarget

`func (o *LinksInner2) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LinksInner2) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LinksInner2) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LinksInner2) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


