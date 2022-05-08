# MiddlewareDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RequireSession** | Pointer to **bool** |  | [optional] 

## Methods

### NewMiddlewareDefinition

`func NewMiddlewareDefinition() *MiddlewareDefinition`

NewMiddlewareDefinition instantiates a new MiddlewareDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareDefinitionWithDefaults

`func NewMiddlewareDefinitionWithDefaults() *MiddlewareDefinition`

NewMiddlewareDefinitionWithDefaults instantiates a new MiddlewareDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MiddlewareDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MiddlewareDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MiddlewareDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MiddlewareDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *MiddlewareDefinition) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MiddlewareDefinition) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MiddlewareDefinition) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MiddlewareDefinition) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRequireSession

`func (o *MiddlewareDefinition) GetRequireSession() bool`

GetRequireSession returns the RequireSession field if non-nil, zero value otherwise.

### GetRequireSessionOk

`func (o *MiddlewareDefinition) GetRequireSessionOk() (*bool, bool)`

GetRequireSessionOk returns a tuple with the RequireSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSession

`func (o *MiddlewareDefinition) SetRequireSession(v bool)`

SetRequireSession sets RequireSession field to given value.

### HasRequireSession

`func (o *MiddlewareDefinition) HasRequireSession() bool`

HasRequireSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


