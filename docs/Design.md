# Design

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The design hashed id | [optional] 
**Name** | Pointer to **string** | The design name | [optional] 
**Design** | Pointer to **string** | The design HTML | [optional] 
**IsCustom** | Pointer to **bool** | Flag to determine if the design is a custom user design | [optional] 
**IsActive** | Pointer to **bool** | Flag to determine if the design is available for use | [optional] 
**IsDeleted** | Pointer to **bool** | Flag to determine if the design is deleted | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**DeletedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewDesign

`func NewDesign() *Design`

NewDesign instantiates a new Design object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignWithDefaults

`func NewDesignWithDefaults() *Design`

NewDesignWithDefaults instantiates a new Design object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Design) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Design) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Design) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Design) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Design) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Design) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Design) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Design) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesign

`func (o *Design) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *Design) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *Design) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *Design) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetIsCustom

`func (o *Design) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *Design) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *Design) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *Design) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetIsActive

`func (o *Design) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Design) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Design) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Design) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Design) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Design) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Design) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Design) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Design) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Design) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Design) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Design) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Design) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Design) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Design) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Design) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Design) GetDeletedAt() float32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Design) GetDeletedAtOk() (*float32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Design) SetDeletedAt(v float32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Design) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


