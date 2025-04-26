# CompanyLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | This field will reference one of the following entity hashed ID payment_id, invoice_id or credit_id | [optional] 
**Notes** | Pointer to **string** | The notes which reference this entry of the ledger | [optional] 
**Balance** | Pointer to **float32** | The client balance | [optional] 
**Adjustment** | Pointer to **float32** | The amount the client balance is adjusted by | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewCompanyLedger

`func NewCompanyLedger() *CompanyLedger`

NewCompanyLedger instantiates a new CompanyLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyLedgerWithDefaults

`func NewCompanyLedgerWithDefaults() *CompanyLedger`

NewCompanyLedgerWithDefaults instantiates a new CompanyLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *CompanyLedger) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CompanyLedger) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CompanyLedger) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CompanyLedger) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetNotes

`func (o *CompanyLedger) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CompanyLedger) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CompanyLedger) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CompanyLedger) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetBalance

`func (o *CompanyLedger) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CompanyLedger) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CompanyLedger) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CompanyLedger) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetAdjustment

`func (o *CompanyLedger) GetAdjustment() float32`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *CompanyLedger) GetAdjustmentOk() (*float32, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *CompanyLedger) SetAdjustment(v float32)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *CompanyLedger) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompanyLedger) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyLedger) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyLedger) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompanyLedger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CompanyLedger) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyLedger) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyLedger) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CompanyLedger) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


