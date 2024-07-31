# AuthAccessPolicyConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSubnets** | Pointer to [**[]AuthAccessPolicyConditionsAllowedSubnetsInner**](AuthAccessPolicyConditionsAllowedSubnetsInner.md) | An array of IP addresses with subnet masks in CIDR notation (both IPv4 or IPv6 are supported) that is used to restrict access to only hosts that are part of at least one of the subnets. Providing an empty array is equivalent to providing an empty &#x60;conditions&#x60; object, and results in the removal of the &#x60;conditions&#x60; object from the Access Policy. Note that an IP address is not valid CIDR notation. For specifying a single IP address use a subnet mask of &#x60;/32&#x60; for IPv4 and &#x60;/128&#x60; for IPv6. Examples: 192.168.0.10/32 2001:db0:82a3:0:0:8a2e:370:1234/128  | [optional] 

## Methods

### NewAuthAccessPolicyConditions

`func NewAuthAccessPolicyConditions() *AuthAccessPolicyConditions`

NewAuthAccessPolicyConditions instantiates a new AuthAccessPolicyConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccessPolicyConditionsWithDefaults

`func NewAuthAccessPolicyConditionsWithDefaults() *AuthAccessPolicyConditions`

NewAuthAccessPolicyConditionsWithDefaults instantiates a new AuthAccessPolicyConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSubnets

`func (o *AuthAccessPolicyConditions) GetAllowedSubnets() []AuthAccessPolicyConditionsAllowedSubnetsInner`

GetAllowedSubnets returns the AllowedSubnets field if non-nil, zero value otherwise.

### GetAllowedSubnetsOk

`func (o *AuthAccessPolicyConditions) GetAllowedSubnetsOk() (*[]AuthAccessPolicyConditionsAllowedSubnetsInner, bool)`

GetAllowedSubnetsOk returns a tuple with the AllowedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubnets

`func (o *AuthAccessPolicyConditions) SetAllowedSubnets(v []AuthAccessPolicyConditionsAllowedSubnetsInner)`

SetAllowedSubnets sets AllowedSubnets field to given value.

### HasAllowedSubnets

`func (o *AuthAccessPolicyConditions) HasAllowedSubnets() bool`

HasAllowedSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


