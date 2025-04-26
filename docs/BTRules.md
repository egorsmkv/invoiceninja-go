# BTRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataKey** | Pointer to **string** | The key to search | [optional] 
**Operator** | Pointer to **string** | The operator flag of the search | [optional] 
**Value** | Pointer to **string** | The value to search for | [optional] 

## Methods

### NewBTRules

`func NewBTRules() *BTRules`

NewBTRules instantiates a new BTRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRulesWithDefaults

`func NewBTRulesWithDefaults() *BTRules`

NewBTRulesWithDefaults instantiates a new BTRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataKey

`func (o *BTRules) GetDataKey() string`

GetDataKey returns the DataKey field if non-nil, zero value otherwise.

### GetDataKeyOk

`func (o *BTRules) GetDataKeyOk() (*string, bool)`

GetDataKeyOk returns a tuple with the DataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKey

`func (o *BTRules) SetDataKey(v string)`

SetDataKey sets DataKey field to given value.

### HasDataKey

`func (o *BTRules) HasDataKey() bool`

HasDataKey returns a boolean if a field has been set.

### GetOperator

`func (o *BTRules) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BTRules) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BTRules) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *BTRules) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *BTRules) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTRules) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTRules) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BTRules) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


