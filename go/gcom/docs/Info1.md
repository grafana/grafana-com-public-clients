# Info1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**Author**](Author.md) |  | [optional] 
**Build** | Pointer to [**Build**](Build.md) |  | [optional] 
**Description** | Pointer to **string** | Description of plugin. Used on the plugins page in Grafana and for search on grafana.com. | [optional] 
**Keywords** | **[]string** | Array of plugin keywords. Used for search on grafana.com. | 
**Links** | Pointer to [**[]LinksInner2**](LinksInner2.md) | An array of link objects to be displayed on this plugin&#39;s project page in the form &#x60;{name: &#39;foo&#39;, url: &#39;http://example.com&#39;}&#x60; | [optional] 
**Logos** | [**Logos**](Logos.md) |  | 
**Screenshots** | Pointer to [**[]ScreenshotsInner**](ScreenshotsInner.md) | An array of screenshot objects in the form &#x60;{name: &#39;bar&#39;, path: &#39;img/screenshot.png&#39;}&#x60; | [optional] 
**Updated** | **string** | Date when this plugin was built. | 
**Version** | **string** | [SemVer](https://semver.org/) version of this commit, e.g. &#x60;6.7.1&#x60;. | 

## Methods

### NewInfo1

`func NewInfo1(keywords []string, logos Logos, updated string, version string, ) *Info1`

NewInfo1 instantiates a new Info1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfo1WithDefaults

`func NewInfo1WithDefaults() *Info1`

NewInfo1WithDefaults instantiates a new Info1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Info1) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Info1) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Info1) SetAuthor(v Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Info1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuild

`func (o *Info1) GetBuild() Build`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *Info1) GetBuildOk() (*Build, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *Info1) SetBuild(v Build)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *Info1) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetDescription

`func (o *Info1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Info1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Info1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Info1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeywords

`func (o *Info1) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Info1) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Info1) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetLinks

`func (o *Info1) GetLinks() []LinksInner2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Info1) GetLinksOk() (*[]LinksInner2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Info1) SetLinks(v []LinksInner2)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Info1) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLogos

`func (o *Info1) GetLogos() Logos`

GetLogos returns the Logos field if non-nil, zero value otherwise.

### GetLogosOk

`func (o *Info1) GetLogosOk() (*Logos, bool)`

GetLogosOk returns a tuple with the Logos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogos

`func (o *Info1) SetLogos(v Logos)`

SetLogos sets Logos field to given value.


### GetScreenshots

`func (o *Info1) GetScreenshots() []ScreenshotsInner`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *Info1) GetScreenshotsOk() (*[]ScreenshotsInner, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *Info1) SetScreenshots(v []ScreenshotsInner)`

SetScreenshots sets Screenshots field to given value.

### HasScreenshots

`func (o *Info1) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.

### GetUpdated

`func (o *Info1) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Info1) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Info1) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetVersion

`func (o *Info1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Info1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Info1) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


