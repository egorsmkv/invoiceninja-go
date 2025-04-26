package openapi

import (
	"encoding/json"
)

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payment{}

// Payment struct for Payment
type Payment struct {
	// The payment hashed id
	Id *string `json:"id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The invitation hashed id
	InvitationId *string `json:"invitation_id,omitempty"`
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
	// Flags whether the payment was made manually or processed via a gateway
	IsManual *bool `json:"is_manual,omitempty"`
	// Defines if the payment has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The amount of this payment
	Amount *float32 `json:"amount,omitempty"`
	// The refunded amount of this payment
	Refunded *float32 `json:"refunded,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// The company gateway id
	CompanyGatewayId *string      `json:"company_gateway_id,omitempty"`
	Paymentables     *Paymentable `json:"paymentables,omitempty"`
	//
	Invoices []InvoicePaymentable `json:"invoices,omitempty"`
	//
	Credits []CreditPaymentable `json:"credits,omitempty"`
	// The payment number - is a unique alpha numeric number per payment per company
	Number *string `json:"number,omitempty"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment() *Payment {
	this := Payment{}
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Payment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Payment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Payment) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Payment) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Payment) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Payment) SetClientId(v string) {
	o.ClientId = &v
}

// GetInvitationId returns the InvitationId field value if set, zero value otherwise.
func (o *Payment) GetInvitationId() string {
	if o == nil || IsNil(o.InvitationId) {
		var ret string
		return ret
	}
	return *o.InvitationId
}

// GetInvitationIdOk returns a tuple with the InvitationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetInvitationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationId) {
		return nil, false
	}
	return o.InvitationId, true
}

// HasInvitationId returns a boolean if a field has been set.
func (o *Payment) HasInvitationId() bool {
	if o != nil && !IsNil(o.InvitationId) {
		return true
	}

	return false
}

// SetInvitationId gets a reference to the given string and assigns it to the InvitationId field.
func (o *Payment) SetInvitationId(v string) {
	o.InvitationId = &v
}

// GetClientContactId returns the ClientContactId field value if set, zero value otherwise.
func (o *Payment) GetClientContactId() string {
	if o == nil || IsNil(o.ClientContactId) {
		var ret string
		return ret
	}
	return *o.ClientContactId
}

// GetClientContactIdOk returns a tuple with the ClientContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetClientContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientContactId) {
		return nil, false
	}
	return o.ClientContactId, true
}

// HasClientContactId returns a boolean if a field has been set.
func (o *Payment) HasClientContactId() bool {
	if o != nil && !IsNil(o.ClientContactId) {
		return true
	}

	return false
}

// SetClientContactId gets a reference to the given string and assigns it to the ClientContactId field.
func (o *Payment) SetClientContactId(v string) {
	o.ClientContactId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Payment) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Payment) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Payment) SetUserId(v string) {
	o.UserId = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *Payment) GetTypeId() PaymentType {
	if o == nil || IsNil(o.TypeId) {
		var ret PaymentType
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTypeIdOk() (*PaymentType, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *Payment) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given PaymentType and assigns it to the TypeId field.
func (o *Payment) SetTypeId(v PaymentType) {
	o.TypeId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Payment) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Payment) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Payment) SetDate(v string) {
	o.Date = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *Payment) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *Payment) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *Payment) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Payment) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Payment) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Payment) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Payment) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Payment) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Payment) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetIsManual returns the IsManual field value if set, zero value otherwise.
func (o *Payment) GetIsManual() bool {
	if o == nil || IsNil(o.IsManual) {
		var ret bool
		return ret
	}
	return *o.IsManual
}

// GetIsManualOk returns a tuple with the IsManual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetIsManualOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManual) {
		return nil, false
	}
	return o.IsManual, true
}

// HasIsManual returns a boolean if a field has been set.
func (o *Payment) HasIsManual() bool {
	if o != nil && !IsNil(o.IsManual) {
		return true
	}

	return false
}

// SetIsManual gets a reference to the given bool and assigns it to the IsManual field.
func (o *Payment) SetIsManual(v bool) {
	o.IsManual = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Payment) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Payment) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Payment) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Payment) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Payment) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Payment) SetAmount(v float32) {
	o.Amount = &v
}

// GetRefunded returns the Refunded field value if set, zero value otherwise.
func (o *Payment) GetRefunded() float32 {
	if o == nil || IsNil(o.Refunded) {
		var ret float32
		return ret
	}
	return *o.Refunded
}

// GetRefundedOk returns a tuple with the Refunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetRefundedOk() (*float32, bool) {
	if o == nil || IsNil(o.Refunded) {
		return nil, false
	}
	return o.Refunded, true
}

// HasRefunded returns a boolean if a field has been set.
func (o *Payment) HasRefunded() bool {
	if o != nil && !IsNil(o.Refunded) {
		return true
	}

	return false
}

// SetRefunded gets a reference to the given float32 and assigns it to the Refunded field.
func (o *Payment) SetRefunded(v float32) {
	o.Refunded = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Payment) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Payment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Payment) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Payment) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Payment) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Payment) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetCompanyGatewayId returns the CompanyGatewayId field value if set, zero value otherwise.
func (o *Payment) GetCompanyGatewayId() string {
	if o == nil || IsNil(o.CompanyGatewayId) {
		var ret string
		return ret
	}
	return *o.CompanyGatewayId
}

// GetCompanyGatewayIdOk returns a tuple with the CompanyGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetCompanyGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyGatewayId) {
		return nil, false
	}
	return o.CompanyGatewayId, true
}

// HasCompanyGatewayId returns a boolean if a field has been set.
func (o *Payment) HasCompanyGatewayId() bool {
	if o != nil && !IsNil(o.CompanyGatewayId) {
		return true
	}

	return false
}

// SetCompanyGatewayId gets a reference to the given string and assigns it to the CompanyGatewayId field.
func (o *Payment) SetCompanyGatewayId(v string) {
	o.CompanyGatewayId = &v
}

// GetPaymentables returns the Paymentables field value if set, zero value otherwise.
func (o *Payment) GetPaymentables() Paymentable {
	if o == nil || IsNil(o.Paymentables) {
		var ret Paymentable
		return ret
	}
	return *o.Paymentables
}

// GetPaymentablesOk returns a tuple with the Paymentables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetPaymentablesOk() (*Paymentable, bool) {
	if o == nil || IsNil(o.Paymentables) {
		return nil, false
	}
	return o.Paymentables, true
}

// HasPaymentables returns a boolean if a field has been set.
func (o *Payment) HasPaymentables() bool {
	if o != nil && !IsNil(o.Paymentables) {
		return true
	}

	return false
}

// SetPaymentables gets a reference to the given Paymentable and assigns it to the Paymentables field.
func (o *Payment) SetPaymentables(v Paymentable) {
	o.Paymentables = &v
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *Payment) GetInvoices() []InvoicePaymentable {
	if o == nil || IsNil(o.Invoices) {
		var ret []InvoicePaymentable
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetInvoicesOk() ([]InvoicePaymentable, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *Payment) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []InvoicePaymentable and assigns it to the Invoices field.
func (o *Payment) SetInvoices(v []InvoicePaymentable) {
	o.Invoices = v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *Payment) GetCredits() []CreditPaymentable {
	if o == nil || IsNil(o.Credits) {
		var ret []CreditPaymentable
		return ret
	}
	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetCreditsOk() ([]CreditPaymentable, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *Payment) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given []CreditPaymentable and assigns it to the Credits field.
func (o *Payment) SetCredits(v []CreditPaymentable) {
	o.Credits = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Payment) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Payment) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Payment) SetNumber(v string) {
	o.Number = &v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.InvitationId) {
		toSerialize["invitation_id"] = o.InvitationId
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
	if !IsNil(o.IsManual) {
		toSerialize["is_manual"] = o.IsManual
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Refunded) {
		toSerialize["refunded"] = o.Refunded
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	if !IsNil(o.CompanyGatewayId) {
		toSerialize["company_gateway_id"] = o.CompanyGatewayId
	}
	if !IsNil(o.Paymentables) {
		toSerialize["paymentables"] = o.Paymentables
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

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
