# FormattedApiPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Id** | **float32** |  | 
**TypeId** | **float32** |  | 
**TypeName** | **string** |  | 
**TypeCode** | **string** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Version** | **string** |  | 
**VersionStatus** | **string** |  | 
**VersionSignatureType** | **string** |  | 
**VersionSignedByOrg** | **string** |  | 
**VersionSignedByOrgName** | **string** |  | 
**UserId** | **float32** |  | 
**OrgId** | **float32** |  | 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**OrgUrl** | **string** |  | 
**Url** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Downloads** | **float32** |  | 
**Verified** | **bool** |  | 
**Featured** | **float32** |  | 
**Internal** | **bool** |  | 
**DownloadSlug** | **string** |  | 
**Popularity** | **float32** |  | 
**SignatureType** | **string** |  | 
**Packages** | **map[string]interface{}** |  | 
**Links** | [**[]LinksInner1**](LinksInner1.md) |  | 
**AngularDetected** | **bool** |  | 
**LicenseUrl** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewFormattedApiPlugin

`func NewFormattedApiPlugin(status string, id float32, typeId float32, typeName string, typeCode string, slug string, name string, description string, version string, versionStatus string, versionSignatureType string, versionSignedByOrg string, versionSignedByOrgName string, userId float32, orgId float32, orgName string, orgSlug string, orgUrl string, url string, createdAt string, updatedAt string, downloads float32, verified bool, featured float32, internal bool, downloadSlug string, popularity float32, signatureType string, packages map[string]interface{}, links []LinksInner1, angularDetected bool, ) *FormattedApiPlugin`

NewFormattedApiPlugin instantiates a new FormattedApiPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiPluginWithDefaults

`func NewFormattedApiPluginWithDefaults() *FormattedApiPlugin`

NewFormattedApiPluginWithDefaults instantiates a new FormattedApiPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FormattedApiPlugin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormattedApiPlugin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormattedApiPlugin) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetId

`func (o *FormattedApiPlugin) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiPlugin) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiPlugin) SetId(v float32)`

SetId sets Id field to given value.


### GetTypeId

`func (o *FormattedApiPlugin) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *FormattedApiPlugin) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *FormattedApiPlugin) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetTypeName

`func (o *FormattedApiPlugin) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FormattedApiPlugin) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FormattedApiPlugin) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetTypeCode

`func (o *FormattedApiPlugin) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *FormattedApiPlugin) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *FormattedApiPlugin) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetSlug

`func (o *FormattedApiPlugin) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormattedApiPlugin) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormattedApiPlugin) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *FormattedApiPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiPlugin) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FormattedApiPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormattedApiPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormattedApiPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *FormattedApiPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormattedApiPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormattedApiPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionStatus

`func (o *FormattedApiPlugin) GetVersionStatus() string`

GetVersionStatus returns the VersionStatus field if non-nil, zero value otherwise.

### GetVersionStatusOk

`func (o *FormattedApiPlugin) GetVersionStatusOk() (*string, bool)`

GetVersionStatusOk returns a tuple with the VersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStatus

`func (o *FormattedApiPlugin) SetVersionStatus(v string)`

SetVersionStatus sets VersionStatus field to given value.


### GetVersionSignatureType

`func (o *FormattedApiPlugin) GetVersionSignatureType() string`

GetVersionSignatureType returns the VersionSignatureType field if non-nil, zero value otherwise.

### GetVersionSignatureTypeOk

`func (o *FormattedApiPlugin) GetVersionSignatureTypeOk() (*string, bool)`

GetVersionSignatureTypeOk returns a tuple with the VersionSignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSignatureType

`func (o *FormattedApiPlugin) SetVersionSignatureType(v string)`

SetVersionSignatureType sets VersionSignatureType field to given value.


### GetVersionSignedByOrg

`func (o *FormattedApiPlugin) GetVersionSignedByOrg() string`

GetVersionSignedByOrg returns the VersionSignedByOrg field if non-nil, zero value otherwise.

### GetVersionSignedByOrgOk

`func (o *FormattedApiPlugin) GetVersionSignedByOrgOk() (*string, bool)`

GetVersionSignedByOrgOk returns a tuple with the VersionSignedByOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSignedByOrg

`func (o *FormattedApiPlugin) SetVersionSignedByOrg(v string)`

SetVersionSignedByOrg sets VersionSignedByOrg field to given value.


### GetVersionSignedByOrgName

`func (o *FormattedApiPlugin) GetVersionSignedByOrgName() string`

GetVersionSignedByOrgName returns the VersionSignedByOrgName field if non-nil, zero value otherwise.

### GetVersionSignedByOrgNameOk

`func (o *FormattedApiPlugin) GetVersionSignedByOrgNameOk() (*string, bool)`

GetVersionSignedByOrgNameOk returns a tuple with the VersionSignedByOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSignedByOrgName

`func (o *FormattedApiPlugin) SetVersionSignedByOrgName(v string)`

SetVersionSignedByOrgName sets VersionSignedByOrgName field to given value.


### GetUserId

`func (o *FormattedApiPlugin) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FormattedApiPlugin) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FormattedApiPlugin) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetOrgId

`func (o *FormattedApiPlugin) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FormattedApiPlugin) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FormattedApiPlugin) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *FormattedApiPlugin) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FormattedApiPlugin) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FormattedApiPlugin) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *FormattedApiPlugin) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *FormattedApiPlugin) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *FormattedApiPlugin) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetOrgUrl

`func (o *FormattedApiPlugin) GetOrgUrl() string`

GetOrgUrl returns the OrgUrl field if non-nil, zero value otherwise.

### GetOrgUrlOk

`func (o *FormattedApiPlugin) GetOrgUrlOk() (*string, bool)`

GetOrgUrlOk returns a tuple with the OrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUrl

`func (o *FormattedApiPlugin) SetOrgUrl(v string)`

SetOrgUrl sets OrgUrl field to given value.


### GetUrl

`func (o *FormattedApiPlugin) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormattedApiPlugin) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormattedApiPlugin) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *FormattedApiPlugin) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiPlugin) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiPlugin) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FormattedApiPlugin) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiPlugin) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiPlugin) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDownloads

`func (o *FormattedApiPlugin) GetDownloads() float32`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *FormattedApiPlugin) GetDownloadsOk() (*float32, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *FormattedApiPlugin) SetDownloads(v float32)`

SetDownloads sets Downloads field to given value.


### GetVerified

`func (o *FormattedApiPlugin) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FormattedApiPlugin) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FormattedApiPlugin) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetFeatured

`func (o *FormattedApiPlugin) GetFeatured() float32`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *FormattedApiPlugin) GetFeaturedOk() (*float32, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *FormattedApiPlugin) SetFeatured(v float32)`

SetFeatured sets Featured field to given value.


### GetInternal

`func (o *FormattedApiPlugin) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *FormattedApiPlugin) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *FormattedApiPlugin) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetDownloadSlug

`func (o *FormattedApiPlugin) GetDownloadSlug() string`

GetDownloadSlug returns the DownloadSlug field if non-nil, zero value otherwise.

### GetDownloadSlugOk

`func (o *FormattedApiPlugin) GetDownloadSlugOk() (*string, bool)`

GetDownloadSlugOk returns a tuple with the DownloadSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSlug

`func (o *FormattedApiPlugin) SetDownloadSlug(v string)`

SetDownloadSlug sets DownloadSlug field to given value.


### GetPopularity

`func (o *FormattedApiPlugin) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *FormattedApiPlugin) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *FormattedApiPlugin) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.


### GetSignatureType

`func (o *FormattedApiPlugin) GetSignatureType() string`

GetSignatureType returns the SignatureType field if non-nil, zero value otherwise.

### GetSignatureTypeOk

`func (o *FormattedApiPlugin) GetSignatureTypeOk() (*string, bool)`

GetSignatureTypeOk returns a tuple with the SignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureType

`func (o *FormattedApiPlugin) SetSignatureType(v string)`

SetSignatureType sets SignatureType field to given value.


### GetPackages

`func (o *FormattedApiPlugin) GetPackages() map[string]interface{}`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *FormattedApiPlugin) GetPackagesOk() (*map[string]interface{}, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *FormattedApiPlugin) SetPackages(v map[string]interface{})`

SetPackages sets Packages field to given value.


### GetLinks

`func (o *FormattedApiPlugin) GetLinks() []LinksInner1`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiPlugin) GetLinksOk() (*[]LinksInner1, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiPlugin) SetLinks(v []LinksInner1)`

SetLinks sets Links field to given value.


### GetAngularDetected

`func (o *FormattedApiPlugin) GetAngularDetected() bool`

GetAngularDetected returns the AngularDetected field if non-nil, zero value otherwise.

### GetAngularDetectedOk

`func (o *FormattedApiPlugin) GetAngularDetectedOk() (*bool, bool)`

GetAngularDetectedOk returns a tuple with the AngularDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngularDetected

`func (o *FormattedApiPlugin) SetAngularDetected(v bool)`

SetAngularDetected sets AngularDetected field to given value.


### GetLicenseUrl

`func (o *FormattedApiPlugin) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *FormattedApiPlugin) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *FormattedApiPlugin) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *FormattedApiPlugin) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *FormattedApiPlugin) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *FormattedApiPlugin) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *FormattedApiPlugin) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *FormattedApiPlugin) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


