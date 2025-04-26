# Paymentable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The paymentable hashed id | [optional] 
**InvoiceId** | Pointer to **string** | The invoice hashed id | [optional] 
**CreditId** | Pointer to **string** | The credit hashed id | [optional] 
**Refunded** | Pointer to **float32** | The amount that has been refunded for this payment | [optional] 
**Amount** | Pointer to **float32** | The amount that has been applied to the payment | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewPaymentable

`func NewPaymentable() *Paymentable`

NewPaymentable instantiates a new Paymentable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentableWithDefaults

`func NewPaymentableWithDefaults() *Paymentable`

NewPaymentableWithDefaults instantiates a new Paymentable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Paymentable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Paymentable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Paymentable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Paymentable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Paymentable) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Paymentable) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Paymentable) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Paymentable) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCreditId

`func (o *Paymentable) GetCreditId() string`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *Paymentable) GetCreditIdOk() (*string, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *Paymentable) SetCreditId(v string)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *Paymentable) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetRefunded

`func (o *Paymentable) GetRefunded() float32`

GetRefunded returns the Refunded field if non-nil, zero value otherwise.

### GetRefundedOk

`func (o *Paymentable) GetRefundedOk() (*float32, bool)`

GetRefundedOk returns a tuple with the Refunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunded

`func (o *Paymentable) SetRefunded(v float32)`

SetRefunded sets Refunded field to given value.

### HasRefunded

`func (o *Paymentable) HasRefunded() bool`

HasRefunded returns a boolean if a field has been set.

### GetAmount

`func (o *Paymentable) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Paymentable) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Paymentable) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Paymentable) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Paymentable) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Paymentable) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Paymentable) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Paymentable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Paymentable) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Paymentable) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Paymentable) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Paymentable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


