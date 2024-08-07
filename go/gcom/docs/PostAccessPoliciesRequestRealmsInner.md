# PostAccessPoliciesRequestRealmsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**LabelPolicies** | Pointer to [**[]PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner**](PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPostAccessPoliciesRequestRealmsInner

`func NewPostAccessPoliciesRequestRealmsInner(identifier string, type_ string, ) *PostAccessPoliciesRequestRealmsInner`

NewPostAccessPoliciesRequestRealmsInner instantiates a new PostAccessPoliciesRequestRealmsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAccessPoliciesRequestRealmsInnerWithDefaults

`func NewPostAccessPoliciesRequestRealmsInnerWithDefaults() *PostAccessPoliciesRequestRealmsInner`

NewPostAccessPoliciesRequestRealmsInnerWithDefaults instantiates a new PostAccessPoliciesRequestRealmsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PostAccessPoliciesRequestRealmsInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PostAccessPoliciesRequestRealmsInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PostAccessPoliciesRequestRealmsInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLabelPolicies

`func (o *PostAccessPoliciesRequestRealmsInner) GetLabelPolicies() []PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner`

GetLabelPolicies returns the LabelPolicies field if non-nil, zero value otherwise.

### GetLabelPoliciesOk

`func (o *PostAccessPoliciesRequestRealmsInner) GetLabelPoliciesOk() (*[]PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner, bool)`

GetLabelPoliciesOk returns a tuple with the LabelPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPolicies

`func (o *PostAccessPoliciesRequestRealmsInner) SetLabelPolicies(v []PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner)`

SetLabelPolicies sets LabelPolicies field to given value.

### HasLabelPolicies

`func (o *PostAccessPoliciesRequestRealmsInner) HasLabelPolicies() bool`

HasLabelPolicies returns a boolean if a field has been set.

### GetType

`func (o *PostAccessPoliciesRequestRealmsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAccessPoliciesRequestRealmsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAccessPoliciesRequestRealmsInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


