# PaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDays** | Pointer to **int32** | The payment term length in days | [optional] 
**Name** | Pointer to **string** | The payment term length in string format | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewPaymentTerm

`func NewPaymentTerm() *PaymentTerm`

NewPaymentTerm instantiates a new PaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermWithDefaults

`func NewPaymentTermWithDefaults() *PaymentTerm`

NewPaymentTermWithDefaults instantiates a new PaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDays

`func (o *PaymentTerm) GetNumDays() int32`

GetNumDays returns the NumDays field if non-nil, zero value otherwise.

### GetNumDaysOk

`func (o *PaymentTerm) GetNumDaysOk() (*int32, bool)`

GetNumDaysOk returns a tuple with the NumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDays

`func (o *PaymentTerm) SetNumDays(v int32)`

SetNumDays sets NumDays field to given value.

### HasNumDays

`func (o *PaymentTerm) HasNumDays() bool`

HasNumDays returns a boolean if a field has been set.

### GetName

`func (o *PaymentTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentTerm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentTerm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentTerm) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentTerm) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentTerm) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentTerm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentTerm) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentTerm) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentTerm) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentTerm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *PaymentTerm) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PaymentTerm) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PaymentTerm) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PaymentTerm) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


