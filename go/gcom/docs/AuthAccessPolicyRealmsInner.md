# AuthAccessPolicyRealmsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the realm | [optional] 
**Identifier** | Pointer to **string** | The unique identifier of a realm (org or stack). | [optional] 
**LabelPolicies** | Pointer to [**[]AuthAccessPolicyRealmsInnerLabelPoliciesInner**](AuthAccessPolicyRealmsInnerLabelPoliciesInner.md) |  | [optional] 

## Methods

### NewAuthAccessPolicyRealmsInner

`func NewAuthAccessPolicyRealmsInner() *AuthAccessPolicyRealmsInner`

NewAuthAccessPolicyRealmsInner instantiates a new AuthAccessPolicyRealmsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccessPolicyRealmsInnerWithDefaults

`func NewAuthAccessPolicyRealmsInnerWithDefaults() *AuthAccessPolicyRealmsInner`

NewAuthAccessPolicyRealmsInnerWithDefaults instantiates a new AuthAccessPolicyRealmsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthAccessPolicyRealmsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthAccessPolicyRealmsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthAccessPolicyRealmsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthAccessPolicyRealmsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthAccessPolicyRealmsInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthAccessPolicyRealmsInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthAccessPolicyRealmsInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthAccessPolicyRealmsInner) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLabelPolicies

`func (o *AuthAccessPolicyRealmsInner) GetLabelPolicies() []AuthAccessPolicyRealmsInnerLabelPoliciesInner`

GetLabelPolicies returns the LabelPolicies field if non-nil, zero value otherwise.

### GetLabelPoliciesOk

`func (o *AuthAccessPolicyRealmsInner) GetLabelPoliciesOk() (*[]AuthAccessPolicyRealmsInnerLabelPoliciesInner, bool)`

GetLabelPoliciesOk returns a tuple with the LabelPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPolicies

`func (o *AuthAccessPolicyRealmsInner) SetLabelPolicies(v []AuthAccessPolicyRealmsInnerLabelPoliciesInner)`

SetLabelPolicies sets LabelPolicies field to given value.

### HasLabelPolicies

`func (o *AuthAccessPolicyRealmsInner) HasLabelPolicies() bool`

HasLabelPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


