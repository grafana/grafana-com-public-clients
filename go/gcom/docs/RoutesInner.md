# RoutesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | For data source plugins. The route path that is replaced by the route URL field when proxying the call. | [optional] 
**Method** | Pointer to **string** | For data source plugins. Route method matches the HTTP verb like &#x60;GET&#x60; or &#x60;POST&#x60;. Multiple methods can be provided as a comma-separated list. | [optional] 
**Url** | Pointer to **string** | For data source plugins. Route URL is where the request is proxied to. | [optional] 
**ReqSignedIn** | Pointer to **bool** |  | [optional] 
**ReqRole** | Pointer to **string** |  | [optional] 
**ReqAction** | Pointer to **string** | The RBAC action a user must have to use this route. **Warning**: unless the action targets the plugin (or a nested datasource plugin), only the action is verified, not what it applies to. | [optional] 
**Headers** | Pointer to **[]interface{}** | For data source plugins. Route headers adds HTTP headers to the proxied request. | [optional] 
**Body** | Pointer to **map[string]interface{}** | For data source plugins. Route headers set the body content and length to the proxied request. | [optional] 
**TokenAuth** | Pointer to [**TokenAuth**](TokenAuth.md) |  | [optional] 
**JwtTokenAuth** | Pointer to [**JwtTokenAuth**](JwtTokenAuth.md) |  | [optional] 
**UrlParams** | Pointer to [**[]UrlParamsInner**](UrlParamsInner.md) | Add URL parameters to a proxy route | [optional] 

## Methods

### NewRoutesInner

`func NewRoutesInner() *RoutesInner`

NewRoutesInner instantiates a new RoutesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesInnerWithDefaults

`func NewRoutesInnerWithDefaults() *RoutesInner`

NewRoutesInnerWithDefaults instantiates a new RoutesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RoutesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RoutesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RoutesInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RoutesInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *RoutesInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RoutesInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RoutesInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RoutesInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *RoutesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RoutesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RoutesInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RoutesInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetReqSignedIn

`func (o *RoutesInner) GetReqSignedIn() bool`

GetReqSignedIn returns the ReqSignedIn field if non-nil, zero value otherwise.

### GetReqSignedInOk

`func (o *RoutesInner) GetReqSignedInOk() (*bool, bool)`

GetReqSignedInOk returns a tuple with the ReqSignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSignedIn

`func (o *RoutesInner) SetReqSignedIn(v bool)`

SetReqSignedIn sets ReqSignedIn field to given value.

### HasReqSignedIn

`func (o *RoutesInner) HasReqSignedIn() bool`

HasReqSignedIn returns a boolean if a field has been set.

### GetReqRole

`func (o *RoutesInner) GetReqRole() string`

GetReqRole returns the ReqRole field if non-nil, zero value otherwise.

### GetReqRoleOk

`func (o *RoutesInner) GetReqRoleOk() (*string, bool)`

GetReqRoleOk returns a tuple with the ReqRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqRole

`func (o *RoutesInner) SetReqRole(v string)`

SetReqRole sets ReqRole field to given value.

### HasReqRole

`func (o *RoutesInner) HasReqRole() bool`

HasReqRole returns a boolean if a field has been set.

### GetReqAction

`func (o *RoutesInner) GetReqAction() string`

GetReqAction returns the ReqAction field if non-nil, zero value otherwise.

### GetReqActionOk

`func (o *RoutesInner) GetReqActionOk() (*string, bool)`

GetReqActionOk returns a tuple with the ReqAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAction

`func (o *RoutesInner) SetReqAction(v string)`

SetReqAction sets ReqAction field to given value.

### HasReqAction

`func (o *RoutesInner) HasReqAction() bool`

HasReqAction returns a boolean if a field has been set.

### GetHeaders

`func (o *RoutesInner) GetHeaders() []interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RoutesInner) GetHeadersOk() (*[]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RoutesInner) SetHeaders(v []interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RoutesInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *RoutesInner) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RoutesInner) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RoutesInner) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *RoutesInner) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTokenAuth

`func (o *RoutesInner) GetTokenAuth() TokenAuth`

GetTokenAuth returns the TokenAuth field if non-nil, zero value otherwise.

### GetTokenAuthOk

`func (o *RoutesInner) GetTokenAuthOk() (*TokenAuth, bool)`

GetTokenAuthOk returns a tuple with the TokenAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuth

`func (o *RoutesInner) SetTokenAuth(v TokenAuth)`

SetTokenAuth sets TokenAuth field to given value.

### HasTokenAuth

`func (o *RoutesInner) HasTokenAuth() bool`

HasTokenAuth returns a boolean if a field has been set.

### GetJwtTokenAuth

`func (o *RoutesInner) GetJwtTokenAuth() JwtTokenAuth`

GetJwtTokenAuth returns the JwtTokenAuth field if non-nil, zero value otherwise.

### GetJwtTokenAuthOk

`func (o *RoutesInner) GetJwtTokenAuthOk() (*JwtTokenAuth, bool)`

GetJwtTokenAuthOk returns a tuple with the JwtTokenAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTokenAuth

`func (o *RoutesInner) SetJwtTokenAuth(v JwtTokenAuth)`

SetJwtTokenAuth sets JwtTokenAuth field to given value.

### HasJwtTokenAuth

`func (o *RoutesInner) HasJwtTokenAuth() bool`

HasJwtTokenAuth returns a boolean if a field has been set.

### GetUrlParams

`func (o *RoutesInner) GetUrlParams() []UrlParamsInner`

GetUrlParams returns the UrlParams field if non-nil, zero value otherwise.

### GetUrlParamsOk

`func (o *RoutesInner) GetUrlParamsOk() (*[]UrlParamsInner, bool)`

GetUrlParamsOk returns a tuple with the UrlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParams

`func (o *RoutesInner) SetUrlParams(v []UrlParamsInner)`

SetUrlParams sets UrlParams field to given value.

### HasUrlParams

`func (o *RoutesInner) HasUrlParams() bool`

HasUrlParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


