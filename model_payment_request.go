package openapi

import (
	"encoding/json"
)

// checks if the PaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentRequest{}

// PaymentRequest struct for PaymentRequest
type PaymentRequest struct {
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The client contact hashed id
	ClientContactId *string `json:"client_contact_id,omitempty"`
	// The user hashed id
	UserId *string      `json:"user_id,omitempty"`
	TypeId *PaymentType `json:"type_id,omitempty"`
	// The Payment date
	Date *string `json:"date,omitempty"`
	// The transaction reference as defined by the payment gateway
	TransactionReference *string `json:"transaction_reference,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The private notes of the payment
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The amount of this payment
	Amount *float32 `json:"amount,omitempty"`
	// The payment number - is a unique alpha numeric number per payment per company
	Number *string `json:"number,omitempty"`
	//
	Invoices []InvoicePaymentable `json:"invoices,omitempty"`
	//
	Credits []CreditPaymentable `json:"credits,omitempty"`
}

// NewPaymentRequest instantiates a new PaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequest() *PaymentRequest {
	this := PaymentRequest{}
	return &this
}

// NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestWithDefaults() *PaymentRequest {
	this := PaymentRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientContactId returns the ClientContactId field value if set, zero value otherwise.
func (o *PaymentRequest) GetClientContactId() string {
	if o == nil || IsNil(o.ClientContactId) {
		var ret string
		return ret
	}
	return *o.ClientContactId
}

// GetClientContactIdOk returns a tuple with the ClientContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetClientContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientContactId) {
		return nil, false
	}
	return o.ClientContactId, true
}

// HasClientContactId returns a boolean if a field has been set.
func (o *PaymentRequest) HasClientContactId() bool {
	if o != nil && !IsNil(o.ClientContactId) {
		return true
	}

	return false
}

// SetClientContactId gets a reference to the given string and assigns it to the ClientContactId field.
func (o *PaymentRequest) SetClientContactId(v string) {
	o.ClientContactId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PaymentRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PaymentRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PaymentRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *PaymentRequest) GetTypeId() PaymentType {
	if o == nil || IsNil(o.TypeId) {
		var ret PaymentType
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetTypeIdOk() (*PaymentType, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *PaymentRequest) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given PaymentType and assigns it to the TypeId field.
func (o *PaymentRequest) SetTypeId(v PaymentType) {
	o.TypeId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PaymentRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PaymentRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *PaymentRequest) SetDate(v string) {
	o.Date = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *PaymentRequest) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *PaymentRequest) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *PaymentRequest) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *PaymentRequest) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *PaymentRequest) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *PaymentRequest) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *PaymentRequest) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *PaymentRequest) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *PaymentRequest) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentRequest) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PaymentRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *PaymentRequest) GetInvoices() []InvoicePaymentable {
	if o == nil || IsNil(o.Invoices) {
		var ret []InvoicePaymentable
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetInvoicesOk() ([]InvoicePaymentable, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *PaymentRequest) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []InvoicePaymentable and assigns it to the Invoices field.
func (o *PaymentRequest) SetInvoices(v []InvoicePaymentable) {
	o.Invoices = v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *PaymentRequest) GetCredits() []CreditPaymentable {
	if o == nil || IsNil(o.Credits) {
		var ret []CreditPaymentable
		return ret
	}
	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetCreditsOk() ([]CreditPaymentable, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *PaymentRequest) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given []CreditPaymentable and assigns it to the Credits field.
func (o *PaymentRequest) SetCredits(v []CreditPaymentable) {
	o.Credits = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PaymentRequest) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequest) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PaymentRequest) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *PaymentRequest) SetNumber(v string) {
	o.Number = &v
}

func (o PaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequest) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientContactId) {
		toSerialize["client_contact_id"] = o.ClientContactId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.TypeId) {
		toSerialize["type_id"] = o.TypeId
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transaction_reference"] = o.TransactionReference
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullablePaymentRequest struct {
	value *PaymentRequest
	isSet bool
}

func (v NullablePaymentRequest) Get() *PaymentRequest {
	return v.value
}

func (v *NullablePaymentRequest) Set(val *PaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequest(val *PaymentRequest) *NullablePaymentRequest {
	return &NullablePaymentRequest{value: val, isSet: true}
}

func (v NullablePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
