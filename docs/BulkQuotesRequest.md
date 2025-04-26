# BulkQuotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to be performed, options include:   - &#x60;approve&#x60;     Bulk approve an array of quotes   - &#x60;convert&#x60;     Bulk convert an array of quotes to invoices   - &#x60;send_email&#x60;     Bulk send an array of quotes as emails   - &#x60;mark_sent&#x60;     Bulk mark an array of quotes as sent   - &#x60;restore&#x60;     Restores an array of quotes   - &#x60;delete&#x60;     Deletes an array of invoices   - &#x60;archive&#x60;     Archives an array of invoices    | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkQuotesRequest

`func NewBulkQuotesRequest() *BulkQuotesRequest`

NewBulkQuotesRequest instantiates a new BulkQuotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkQuotesRequestWithDefaults

`func NewBulkQuotesRequestWithDefaults() *BulkQuotesRequest`

NewBulkQuotesRequestWithDefaults instantiates a new BulkQuotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkQuotesRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkQuotesRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkQuotesRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkQuotesRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIds

`func (o *BulkQuotesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkQuotesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkQuotesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkQuotesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


