# BulkRecurringInvoicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to be performed, options include:   - &#x60;start&#x60;     Starts (or restarts) the recurring invoice. **note** if the recurring invoice has been stopped for a long time, it will attempt to catch back up firing a new Invoice every hour per interval that has been missed.     If you do not wish to have the recurring invoice catch up, you should set the next_send_date to the correct date you wish the recurring invoice to commence from. - &#x60;stop&#x60;     Stops the recurring invoice.  - &#x60;send_now&#x60;     Force sends the recurring invoice - this option is only available when the recurring invoice is in a draft state.   - &#x60;restore&#x60;     Restores the recurring invoice from an archived or deleted state. - &#x60;archive&#x60;     Archives the recurring invoice. The recurring invoice will not fire in this state. - &#x60;delete&#x60;     Deletes a recurring invoice.    | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkRecurringInvoicesRequest

`func NewBulkRecurringInvoicesRequest() *BulkRecurringInvoicesRequest`

NewBulkRecurringInvoicesRequest instantiates a new BulkRecurringInvoicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRecurringInvoicesRequestWithDefaults

`func NewBulkRecurringInvoicesRequestWithDefaults() *BulkRecurringInvoicesRequest`

NewBulkRecurringInvoicesRequestWithDefaults instantiates a new BulkRecurringInvoicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkRecurringInvoicesRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkRecurringInvoicesRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkRecurringInvoicesRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkRecurringInvoicesRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIds

`func (o *BulkRecurringInvoicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkRecurringInvoicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkRecurringInvoicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkRecurringInvoicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


