# ExpenseCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The expense hashed id | [optional] 
**Name** | Pointer to **string** | The expense category name | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**IsDeleted** | Pointer to **bool** | Flag determining whether the expense category has been deleted | [optional] 
**UpdatedAt** | Pointer to **int32** | The updated at timestamp | [optional] 
**CreatedAt** | Pointer to **int32** | The created at timestamp | [optional] 

## Methods

### NewExpenseCategory

`func NewExpenseCategory() *ExpenseCategory`

NewExpenseCategory instantiates a new ExpenseCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseCategoryWithDefaults

`func NewExpenseCategoryWithDefaults() *ExpenseCategory`

NewExpenseCategoryWithDefaults instantiates a new ExpenseCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExpenseCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpenseCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ExpenseCategory) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExpenseCategory) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExpenseCategory) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExpenseCategory) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ExpenseCategory) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ExpenseCategory) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ExpenseCategory) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ExpenseCategory) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExpenseCategory) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExpenseCategory) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExpenseCategory) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExpenseCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExpenseCategory) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExpenseCategory) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExpenseCategory) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExpenseCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


