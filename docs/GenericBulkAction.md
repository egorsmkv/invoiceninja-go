# GenericBulkAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to perform ie. archive / restore / delete | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGenericBulkAction

`func NewGenericBulkAction() *GenericBulkAction`

NewGenericBulkAction instantiates a new GenericBulkAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericBulkActionWithDefaults

`func NewGenericBulkActionWithDefaults() *GenericBulkAction`

NewGenericBulkActionWithDefaults instantiates a new GenericBulkAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GenericBulkAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GenericBulkAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GenericBulkAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GenericBulkAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIds

`func (o *GenericBulkAction) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GenericBulkAction) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GenericBulkAction) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *GenericBulkAction) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


