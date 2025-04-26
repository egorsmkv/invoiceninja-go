# TaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The task status hashed id | [optional] [readonly] 
**Name** | Pointer to **string** | The task status name | [optional] 
**Color** | Pointer to **string** | The task status color - hex value | [optional] 
**TaskStatusOrder** | Pointer to **int32** | The order of the task status | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**IsDeleted** | Pointer to **bool** | A boolean flag determining if the task status has been deleted | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 

## Methods

### NewTaskStatus

`func NewTaskStatus() *TaskStatus`

NewTaskStatus instantiates a new TaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusWithDefaults

`func NewTaskStatusWithDefaults() *TaskStatus`

NewTaskStatusWithDefaults instantiates a new TaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaskStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *TaskStatus) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TaskStatus) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TaskStatus) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TaskStatus) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetTaskStatusOrder

`func (o *TaskStatus) GetTaskStatusOrder() int32`

GetTaskStatusOrder returns the TaskStatusOrder field if non-nil, zero value otherwise.

### GetTaskStatusOrderOk

`func (o *TaskStatus) GetTaskStatusOrderOk() (*int32, bool)`

GetTaskStatusOrderOk returns a tuple with the TaskStatusOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatusOrder

`func (o *TaskStatus) SetTaskStatusOrder(v int32)`

SetTaskStatusOrder sets TaskStatusOrder field to given value.

### HasTaskStatusOrder

`func (o *TaskStatus) HasTaskStatusOrder() bool`

HasTaskStatusOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskStatus) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskStatus) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskStatus) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TaskStatus) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TaskStatus) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TaskStatus) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TaskStatus) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskStatus) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskStatus) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskStatus) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *TaskStatus) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *TaskStatus) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *TaskStatus) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *TaskStatus) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


