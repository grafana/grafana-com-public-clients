# Json

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique name of the plugin. If the plugin is published on grafana.com, then the plugin &#x60;id&#x60; has to follow the naming conventions. | 
**Type** | **string** | Plugin type. | 
**Info** | [**Info1**](Info1.md) |  | 
**Name** | **string** | Human-readable name of the plugin that is shown to the user in the UI. | 
**Dependencies** | [**Dependencies**](Dependencies.md) |  | 
**Schema** | Pointer to **string** | Schema definition for the plugin.json file. Used primarily for schema validation. | [optional] 
**Alerting** | Pointer to **bool** | For data source plugins, if the plugin supports alerting. Requires &#x60;backend&#x60; to be set to &#x60;true&#x60;. | [optional] 
**Annotations** | Pointer to **bool** | For data source plugins, if the plugin supports annotation queries. | [optional] 
**AutoEnabled** | Pointer to **bool** | Set to true for app plugins that should be enabled and pinned to the navigation bar in all orgs. | [optional] 
**Backend** | Pointer to **bool** | If the plugin has a backend component. | [optional] 
**BuildMode** | Pointer to **string** | The build mode of the plugin. This field is set automatically at build time, so it should not be provided manually. | [optional] 
**BuiltIn** | Pointer to **bool** | [internal only] Indicates whether the plugin is developed and shipped as part of Grafana. Also known as a &#39;core plugin&#39;. | [optional] 
**Category** | Pointer to **string** | Plugin category used on the Add data source page. | [optional] 
**EnterpriseFeatures** | Pointer to [**EnterpriseFeatures**](EnterpriseFeatures.md) |  | [optional] 
**Executable** | Pointer to **string** | The first part of the file name of the backend component executable. There can be multiple executables built for different operating system and architecture. Grafana will check for executables named &#x60;&lt;executable&gt;_&lt;$GOOS&gt;_&lt;lower case $GOARCH&gt;&lt;.exe for Windows&gt;&#x60;, e.g. &#x60;plugin_linux_amd64&#x60;. Combination of $GOOS and $GOARCH can be found here: https://golang.org/doc/install/source#environment. | [optional] 
**HideFromList** | Pointer to **bool** | [internal only] Excludes the plugin from listings in Grafana&#39;s UI. Only allowed for &#x60;builtIn&#x60; plugins. | [optional] 
**Includes** | Pointer to [**[]IncludesInner**](IncludesInner.md) | Resources to include in plugin. | [optional] 
**Logs** | Pointer to **bool** | For data source plugins, if the plugin supports logs. It may be used to filter logs only features. | [optional] 
**Metrics** | Pointer to **bool** | For data source plugins, if the plugin supports metric queries. Used to enable the plugin in the panel editor. | [optional] 
**MultiValueFilterOperators** | Pointer to **bool** | For data source plugins, if the plugin supports multi value operators in adhoc filters. | [optional] 
**PascalName** | Pointer to **string** | [internal only] The PascalCase name for the plugin. Used for creating machine-friendly identifiers, typically in code generation. If not provided, defaults to name, but title-cased and sanitized (only alphabetical characters allowed). | [optional] 
**Preload** | Pointer to **bool** | Initialize plugin on startup. By default, the plugin initializes on first use, but when preload is set to true the plugin loads when the Grafana web app loads the first time. Only applicable to app plugins. When setting to &#x60;true&#x60;, implement [frontend code splitting](https://grafana.com/developers/plugin-tools/get-started/best-practices#app-plugins) to minimise performance implications. | [optional] 
**QueryOptions** | Pointer to [**QueryOptions**](QueryOptions.md) |  | [optional] 
**Routes** | Pointer to [**[]RoutesInner**](RoutesInner.md) | For data source plugins. Proxy routes used for plugin authentication and adding headers to HTTP requests made by the plugin. For more information, refer to [Authentication for data source plugins](https://grafana.com/developers/plugin-tools/how-to-guides/data-source-plugins/add-authentication-for-data-source-plugins). | [optional] 
**SkipDataQuery** | Pointer to **bool** | For panel plugins. Hides the query editor. | [optional] 
**State** | Pointer to **string** | Describes plugins life cycle status | [optional] 
**Streaming** | Pointer to **bool** | For data source plugins, if the plugin supports streaming. Used in Explore to start live streaming. | [optional] 
**Tracing** | Pointer to **bool** | For data source plugins, if the plugin supports tracing. Used for example to link logs (e.g. Loki logs) with tracing plugins. | [optional] 
**Iam** | Pointer to [**Iam**](Iam.md) |  | [optional] 
**Roles** | Pointer to [**[]RolesInner**](RolesInner.md) | List of RBAC roles defined by the plugin and their default assignments to basic roles (&#x60;Viewer&#x60;, &#x60;Editor&#x60;, &#x60;Admin&#x60;, &#x60;Grafana Admin&#x60;). Requires Grafana version 9.4.0 or later. | [optional] 
**Extensions** | Pointer to [**Extensions1**](Extensions1.md) |  | [optional] 
**Languages** | Pointer to **[]string** | The list of languages supported by the plugin. Each entry should be a locale identifier in the format &#x60;language-COUNTRY&#x60; (for example &#x60;en-US&#x60;, &#x60;es-ES&#x60;, &#x60;de-DE&#x60;). | [optional] 

## Methods

### NewJson

`func NewJson(id string, type_ string, info Info1, name string, dependencies Dependencies, ) *Json`

NewJson instantiates a new Json object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWithDefaults

`func NewJsonWithDefaults() *Json`

NewJsonWithDefaults instantiates a new Json object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Json) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Json) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Json) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Json) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Json) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Json) SetType(v string)`

SetType sets Type field to given value.


### GetInfo

`func (o *Json) GetInfo() Info1`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Json) GetInfoOk() (*Info1, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Json) SetInfo(v Info1)`

SetInfo sets Info field to given value.


### GetName

`func (o *Json) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Json) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Json) SetName(v string)`

SetName sets Name field to given value.


### GetDependencies

`func (o *Json) GetDependencies() Dependencies`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Json) GetDependenciesOk() (*Dependencies, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Json) SetDependencies(v Dependencies)`

SetDependencies sets Dependencies field to given value.


### GetSchema

`func (o *Json) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Json) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Json) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Json) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAlerting

`func (o *Json) GetAlerting() bool`

GetAlerting returns the Alerting field if non-nil, zero value otherwise.

### GetAlertingOk

`func (o *Json) GetAlertingOk() (*bool, bool)`

GetAlertingOk returns a tuple with the Alerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerting

`func (o *Json) SetAlerting(v bool)`

SetAlerting sets Alerting field to given value.

### HasAlerting

`func (o *Json) HasAlerting() bool`

HasAlerting returns a boolean if a field has been set.

### GetAnnotations

`func (o *Json) GetAnnotations() bool`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Json) GetAnnotationsOk() (*bool, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Json) SetAnnotations(v bool)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Json) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAutoEnabled

`func (o *Json) GetAutoEnabled() bool`

GetAutoEnabled returns the AutoEnabled field if non-nil, zero value otherwise.

### GetAutoEnabledOk

`func (o *Json) GetAutoEnabledOk() (*bool, bool)`

GetAutoEnabledOk returns a tuple with the AutoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnabled

`func (o *Json) SetAutoEnabled(v bool)`

SetAutoEnabled sets AutoEnabled field to given value.

### HasAutoEnabled

`func (o *Json) HasAutoEnabled() bool`

HasAutoEnabled returns a boolean if a field has been set.

### GetBackend

`func (o *Json) GetBackend() bool`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *Json) GetBackendOk() (*bool, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *Json) SetBackend(v bool)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *Json) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetBuildMode

`func (o *Json) GetBuildMode() string`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *Json) GetBuildModeOk() (*string, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *Json) SetBuildMode(v string)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *Json) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetBuiltIn

`func (o *Json) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *Json) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *Json) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *Json) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetCategory

`func (o *Json) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Json) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Json) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Json) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnterpriseFeatures

`func (o *Json) GetEnterpriseFeatures() EnterpriseFeatures`

GetEnterpriseFeatures returns the EnterpriseFeatures field if non-nil, zero value otherwise.

### GetEnterpriseFeaturesOk

`func (o *Json) GetEnterpriseFeaturesOk() (*EnterpriseFeatures, bool)`

GetEnterpriseFeaturesOk returns a tuple with the EnterpriseFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseFeatures

`func (o *Json) SetEnterpriseFeatures(v EnterpriseFeatures)`

SetEnterpriseFeatures sets EnterpriseFeatures field to given value.

### HasEnterpriseFeatures

`func (o *Json) HasEnterpriseFeatures() bool`

HasEnterpriseFeatures returns a boolean if a field has been set.

### GetExecutable

`func (o *Json) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *Json) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *Json) SetExecutable(v string)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *Json) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetHideFromList

`func (o *Json) GetHideFromList() bool`

GetHideFromList returns the HideFromList field if non-nil, zero value otherwise.

### GetHideFromListOk

`func (o *Json) GetHideFromListOk() (*bool, bool)`

GetHideFromListOk returns a tuple with the HideFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromList

`func (o *Json) SetHideFromList(v bool)`

SetHideFromList sets HideFromList field to given value.

### HasHideFromList

`func (o *Json) HasHideFromList() bool`

HasHideFromList returns a boolean if a field has been set.

### GetIncludes

`func (o *Json) GetIncludes() []IncludesInner`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *Json) GetIncludesOk() (*[]IncludesInner, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *Json) SetIncludes(v []IncludesInner)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *Json) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetLogs

`func (o *Json) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *Json) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *Json) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *Json) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *Json) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Json) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Json) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Json) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMultiValueFilterOperators

`func (o *Json) GetMultiValueFilterOperators() bool`

GetMultiValueFilterOperators returns the MultiValueFilterOperators field if non-nil, zero value otherwise.

### GetMultiValueFilterOperatorsOk

`func (o *Json) GetMultiValueFilterOperatorsOk() (*bool, bool)`

GetMultiValueFilterOperatorsOk returns a tuple with the MultiValueFilterOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueFilterOperators

`func (o *Json) SetMultiValueFilterOperators(v bool)`

SetMultiValueFilterOperators sets MultiValueFilterOperators field to given value.

### HasMultiValueFilterOperators

`func (o *Json) HasMultiValueFilterOperators() bool`

HasMultiValueFilterOperators returns a boolean if a field has been set.

### GetPascalName

`func (o *Json) GetPascalName() string`

GetPascalName returns the PascalName field if non-nil, zero value otherwise.

### GetPascalNameOk

`func (o *Json) GetPascalNameOk() (*string, bool)`

GetPascalNameOk returns a tuple with the PascalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPascalName

`func (o *Json) SetPascalName(v string)`

SetPascalName sets PascalName field to given value.

### HasPascalName

`func (o *Json) HasPascalName() bool`

HasPascalName returns a boolean if a field has been set.

### GetPreload

`func (o *Json) GetPreload() bool`

GetPreload returns the Preload field if non-nil, zero value otherwise.

### GetPreloadOk

`func (o *Json) GetPreloadOk() (*bool, bool)`

GetPreloadOk returns a tuple with the Preload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreload

`func (o *Json) SetPreload(v bool)`

SetPreload sets Preload field to given value.

### HasPreload

`func (o *Json) HasPreload() bool`

HasPreload returns a boolean if a field has been set.

### GetQueryOptions

`func (o *Json) GetQueryOptions() QueryOptions`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *Json) GetQueryOptionsOk() (*QueryOptions, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *Json) SetQueryOptions(v QueryOptions)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *Json) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.

### GetRoutes

`func (o *Json) GetRoutes() []RoutesInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Json) GetRoutesOk() (*[]RoutesInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Json) SetRoutes(v []RoutesInner)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Json) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetSkipDataQuery

`func (o *Json) GetSkipDataQuery() bool`

GetSkipDataQuery returns the SkipDataQuery field if non-nil, zero value otherwise.

### GetSkipDataQueryOk

`func (o *Json) GetSkipDataQueryOk() (*bool, bool)`

GetSkipDataQueryOk returns a tuple with the SkipDataQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDataQuery

`func (o *Json) SetSkipDataQuery(v bool)`

SetSkipDataQuery sets SkipDataQuery field to given value.

### HasSkipDataQuery

`func (o *Json) HasSkipDataQuery() bool`

HasSkipDataQuery returns a boolean if a field has been set.

### GetState

`func (o *Json) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Json) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Json) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Json) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreaming

`func (o *Json) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *Json) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *Json) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *Json) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetTracing

`func (o *Json) GetTracing() bool`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *Json) GetTracingOk() (*bool, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *Json) SetTracing(v bool)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *Json) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetIam

`func (o *Json) GetIam() Iam`

GetIam returns the Iam field if non-nil, zero value otherwise.

### GetIamOk

`func (o *Json) GetIamOk() (*Iam, bool)`

GetIamOk returns a tuple with the Iam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIam

`func (o *Json) SetIam(v Iam)`

SetIam sets Iam field to given value.

### HasIam

`func (o *Json) HasIam() bool`

HasIam returns a boolean if a field has been set.

### GetRoles

`func (o *Json) GetRoles() []RolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Json) GetRolesOk() (*[]RolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Json) SetRoles(v []RolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Json) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetExtensions

`func (o *Json) GetExtensions() Extensions1`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Json) GetExtensionsOk() (*Extensions1, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Json) SetExtensions(v Extensions1)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Json) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetLanguages

`func (o *Json) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Json) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Json) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Json) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


