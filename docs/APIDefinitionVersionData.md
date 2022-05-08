# APIDefinitionVersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **string** |  | [optional] 
**NotVersioned** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**map[string]VersionInfo**](VersionInfo.md) |  | [optional] 

## Methods

### NewAPIDefinitionVersionData

`func NewAPIDefinitionVersionData() *APIDefinitionVersionData`

NewAPIDefinitionVersionData instantiates a new APIDefinitionVersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionVersionDataWithDefaults

`func NewAPIDefinitionVersionDataWithDefaults() *APIDefinitionVersionData`

NewAPIDefinitionVersionDataWithDefaults instantiates a new APIDefinitionVersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *APIDefinitionVersionData) GetDefaultVersion() string`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *APIDefinitionVersionData) GetDefaultVersionOk() (*string, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *APIDefinitionVersionData) SetDefaultVersion(v string)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *APIDefinitionVersionData) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetNotVersioned

`func (o *APIDefinitionVersionData) GetNotVersioned() bool`

GetNotVersioned returns the NotVersioned field if non-nil, zero value otherwise.

### GetNotVersionedOk

`func (o *APIDefinitionVersionData) GetNotVersionedOk() (*bool, bool)`

GetNotVersionedOk returns a tuple with the NotVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotVersioned

`func (o *APIDefinitionVersionData) SetNotVersioned(v bool)`

SetNotVersioned sets NotVersioned field to given value.

### HasNotVersioned

`func (o *APIDefinitionVersionData) HasNotVersioned() bool`

HasNotVersioned returns a boolean if a field has been set.

### GetVersions

`func (o *APIDefinitionVersionData) GetVersions() map[string]VersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *APIDefinitionVersionData) GetVersionsOk() (*map[string]VersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *APIDefinitionVersionData) SetVersions(v map[string]VersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *APIDefinitionVersionData) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


