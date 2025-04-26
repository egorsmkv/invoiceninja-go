# BulkInvoicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailType** | Pointer to [**Enum**](enum.md) | The email type to be sent, when bulk emailing invoices | [optional] 
**Action** | Pointer to **string** | The action to be performed, options include:   - &#x60;bulk_download&#x60;     Bulk download an array of invoice PDFs (These are sent to the admin via email.)   - &#x60;download&#x60;     Download a single PDF. (Returns a single PDF object)   - &#x60;bulk_print&#x60;     Merges an array of Invoice PDFs for easy one click printing.   - &#x60;auto_bill&#x60;     Attempts to automatically bill the invoices with the payment method on file.   - &#x60;clone_to_invoice&#x60;     Returns a clone of the invoice.   - &#x60;clone_to_quote&#x60;     Returns a quote cloned using the properties of the given invoice.   - &#x60;mark_paid&#x60;     Marks an array of invoices as paid.   - &#x60;mark_sent&#x60;     Marks an array of invoices as sent.   - &#x60;restore&#x60;     Restores an array of invoices   - &#x60;delete&#x60;     Deletes an array of invoices   - &#x60;archive&#x60;     Archives an array of invoices   - &#x60;cancel&#x60;     Cancels an array of invoices   - &#x60;email&#x60;     Emails an array of invoices   - &#x60;send_email&#x60;     Emails an array of invoices. Requires additional properties to be sent. &#x60;email_type&#x60;    | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkInvoicesRequest

`func NewBulkInvoicesRequest() *BulkInvoicesRequest`

NewBulkInvoicesRequest instantiates a new BulkInvoicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInvoicesRequestWithDefaults

`func NewBulkInvoicesRequestWithDefaults() *BulkInvoicesRequest`

NewBulkInvoicesRequestWithDefaults instantiates a new BulkInvoicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailType

`func (o *BulkInvoicesRequest) GetEmailType() Enum`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *BulkInvoicesRequest) GetEmailTypeOk() (*Enum, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *BulkInvoicesRequest) SetEmailType(v Enum)`

SetEmailType sets EmailType field to given value.

### HasEmailType

`func (o *BulkInvoicesRequest) HasEmailType() bool`

HasEmailType returns a boolean if a field has been set.

### GetAction

`func (o *BulkInvoicesRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkInvoicesRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkInvoicesRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkInvoicesRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIds

`func (o *BulkInvoicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkInvoicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkInvoicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkInvoicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


