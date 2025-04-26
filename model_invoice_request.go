package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceRequest{}

// InvoiceRequest struct for InvoiceRequest
type InvoiceRequest struct {
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The client hashed id
	ClientId string `json:"client_id"`
	// The invoice number - is a unique alpha numeric number per invoice per company
	Number *string `json:"number,omitempty"`
	// The purchase order associated with this invoice
	PoNumber *string `json:"po_number,omitempty"`
	// The invoice terms
	Terms *string `json:"terms,omitempty"`
	// The public notes of the invoice
	PublicNotes *string `json:"public_notes,omitempty"`
	// The private notes of the invoice
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The invoice footer notes
	Footer *string `json:"footer,omitempty"`
	// A custom field value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom field value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom field value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom field value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The tax name
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The tax name
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The tax rate
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The tax rate
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The tax name
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The tax rate
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// An array of objects which define the line items of the invoice
	LineItems []InvoiceItem `json:"line_items,omitempty"`
	// An array of objects which define the invitations of the invoice
	Invitations []InvoiceInvitationRequest `json:"invitations,omitempty"`
	// The invoice discount, can be an amount or a percentage
	Discount *float32 `json:"discount,omitempty"`
	// The deposit/partial amount
	Partial *float32 `json:"partial,omitempty"`
	// Flag determining if the discount is an amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Defines the type of taxes used as either inclusive or exclusive
	UsesInclusiveTaxes *bool `json:"uses_inclusive_taxes,omitempty"`
	// The Invoice Date
	Date *string `json:"date,omitempty"`
	// The due date for the deposit/partial amount
	PartialDueDate *string `json:"partial_due_date,omitempty"`
	// The due date of the invoice
	DueDate *string `json:"due_date,omitempty"`
	// First Custom Surcharge
	CustomSurcharge1 *float32 `json:"custom_surcharge1,omitempty"`
	// Second Custom Surcharge
	CustomSurcharge2 *float32 `json:"custom_surcharge2,omitempty"`
	// Third Custom Surcharge
	CustomSurcharge3 *float32 `json:"custom_surcharge3,omitempty"`
	// Fourth Custom Surcharge
	CustomSurcharge4 *float32 `json:"custom_surcharge4,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax1 *bool `json:"custom_surcharge_tax1,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax2 *bool `json:"custom_surcharge_tax2,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax3 *bool `json:"custom_surcharge_tax3,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax4 *bool `json:"custom_surcharge_tax4,omitempty"`
	// The project associated with this invoice
	ProjectId *string `json:"project_id,omitempty"`
}

type _InvoiceRequest InvoiceRequest

// NewInvoiceRequest instantiates a new InvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceRequest(clientId string) *InvoiceRequest {
	this := InvoiceRequest{}
	this.ClientId = clientId
	return &this
}

// NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRequestWithDefaults() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InvoiceRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InvoiceRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *InvoiceRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *InvoiceRequest) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *InvoiceRequest) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *InvoiceRequest) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value
func (o *InvoiceRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *InvoiceRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InvoiceRequest) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InvoiceRequest) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *InvoiceRequest) SetNumber(v string) {
	o.Number = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPoNumber() string {
	if o == nil || IsNil(o.PoNumber) {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPoNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PoNumber) {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPoNumber() bool {
	if o != nil && !IsNil(o.PoNumber) {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *InvoiceRequest) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *InvoiceRequest) SetTerms(v string) {
	o.Terms = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *InvoiceRequest) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *InvoiceRequest) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *InvoiceRequest) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *InvoiceRequest) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *InvoiceRequest) SetFooter(v string) {
	o.Footer = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *InvoiceRequest) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *InvoiceRequest) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *InvoiceRequest) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *InvoiceRequest) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *InvoiceRequest) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *InvoiceRequest) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *InvoiceRequest) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *InvoiceRequest) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *InvoiceRequest) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *InvoiceRequest) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *InvoiceRequest) GetLineItems() []InvoiceItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []InvoiceItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetLineItemsOk() ([]InvoiceItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *InvoiceRequest) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceItem and assigns it to the LineItems field.
func (o *InvoiceRequest) SetLineItems(v []InvoiceItem) {
	o.LineItems = v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *InvoiceRequest) GetInvitations() []InvoiceInvitationRequest {
	if o == nil || IsNil(o.Invitations) {
		var ret []InvoiceInvitationRequest
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetInvitationsOk() ([]InvoiceInvitationRequest, bool) {
	if o == nil || IsNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *InvoiceRequest) HasInvitations() bool {
	if o != nil && !IsNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []InvoiceInvitationRequest and assigns it to the Invitations field.
func (o *InvoiceRequest) SetInvitations(v []InvoiceInvitationRequest) {
	o.Invitations = v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *InvoiceRequest) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *InvoiceRequest) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *InvoiceRequest) SetDiscount(v float32) {
	o.Discount = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPartial() float32 {
	if o == nil || IsNil(o.Partial) {
		var ret float32
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPartialOk() (*float32, bool) {
	if o == nil || IsNil(o.Partial) {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPartial() bool {
	if o != nil && !IsNil(o.Partial) {
		return true
	}

	return false
}

// SetPartial gets a reference to the given float32 and assigns it to the Partial field.
func (o *InvoiceRequest) SetPartial(v float32) {
	o.Partial = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *InvoiceRequest) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *InvoiceRequest) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *InvoiceRequest) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field value if set, zero value otherwise.
func (o *InvoiceRequest) GetUsesInclusiveTaxes() bool {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.UsesInclusiveTaxes
}

// GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetUsesInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		return nil, false
	}
	return o.UsesInclusiveTaxes, true
}

// HasUsesInclusiveTaxes returns a boolean if a field has been set.
func (o *InvoiceRequest) HasUsesInclusiveTaxes() bool {
	if o != nil && !IsNil(o.UsesInclusiveTaxes) {
		return true
	}

	return false
}

// SetUsesInclusiveTaxes gets a reference to the given bool and assigns it to the UsesInclusiveTaxes field.
func (o *InvoiceRequest) SetUsesInclusiveTaxes(v bool) {
	o.UsesInclusiveTaxes = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *InvoiceRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *InvoiceRequest) SetDate(v string) {
	o.Date = &v
}

// GetPartialDueDate returns the PartialDueDate field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPartialDueDate() string {
	if o == nil || IsNil(o.PartialDueDate) {
		var ret string
		return ret
	}
	return *o.PartialDueDate
}

// GetPartialDueDateOk returns a tuple with the PartialDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPartialDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PartialDueDate) {
		return nil, false
	}
	return o.PartialDueDate, true
}

// HasPartialDueDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPartialDueDate() bool {
	if o != nil && !IsNil(o.PartialDueDate) {
		return true
	}

	return false
}

// SetPartialDueDate gets a reference to the given string and assigns it to the PartialDueDate field.
func (o *InvoiceRequest) SetPartialDueDate(v string) {
	o.PartialDueDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoiceRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InvoiceRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetCustomSurcharge1 returns the CustomSurcharge1 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurcharge1() float32 {
	if o == nil || IsNil(o.CustomSurcharge1) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge1
}

// GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurcharge1Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge1) {
		return nil, false
	}
	return o.CustomSurcharge1, true
}

// HasCustomSurcharge1 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurcharge1() bool {
	if o != nil && !IsNil(o.CustomSurcharge1) {
		return true
	}

	return false
}

// SetCustomSurcharge1 gets a reference to the given float32 and assigns it to the CustomSurcharge1 field.
func (o *InvoiceRequest) SetCustomSurcharge1(v float32) {
	o.CustomSurcharge1 = &v
}

// GetCustomSurcharge2 returns the CustomSurcharge2 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurcharge2() float32 {
	if o == nil || IsNil(o.CustomSurcharge2) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge2
}

// GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurcharge2Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge2) {
		return nil, false
	}
	return o.CustomSurcharge2, true
}

// HasCustomSurcharge2 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurcharge2() bool {
	if o != nil && !IsNil(o.CustomSurcharge2) {
		return true
	}

	return false
}

// SetCustomSurcharge2 gets a reference to the given float32 and assigns it to the CustomSurcharge2 field.
func (o *InvoiceRequest) SetCustomSurcharge2(v float32) {
	o.CustomSurcharge2 = &v
}

// GetCustomSurcharge3 returns the CustomSurcharge3 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurcharge3() float32 {
	if o == nil || IsNil(o.CustomSurcharge3) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge3
}

// GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurcharge3Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge3) {
		return nil, false
	}
	return o.CustomSurcharge3, true
}

// HasCustomSurcharge3 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurcharge3() bool {
	if o != nil && !IsNil(o.CustomSurcharge3) {
		return true
	}

	return false
}

// SetCustomSurcharge3 gets a reference to the given float32 and assigns it to the CustomSurcharge3 field.
func (o *InvoiceRequest) SetCustomSurcharge3(v float32) {
	o.CustomSurcharge3 = &v
}

// GetCustomSurcharge4 returns the CustomSurcharge4 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurcharge4() float32 {
	if o == nil || IsNil(o.CustomSurcharge4) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge4
}

// GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurcharge4Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge4) {
		return nil, false
	}
	return o.CustomSurcharge4, true
}

// HasCustomSurcharge4 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurcharge4() bool {
	if o != nil && !IsNil(o.CustomSurcharge4) {
		return true
	}

	return false
}

// SetCustomSurcharge4 gets a reference to the given float32 and assigns it to the CustomSurcharge4 field.
func (o *InvoiceRequest) SetCustomSurcharge4(v float32) {
	o.CustomSurcharge4 = &v
}

// GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurchargeTax1() bool {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax1
}

// GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurchargeTax1Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		return nil, false
	}
	return o.CustomSurchargeTax1, true
}

// HasCustomSurchargeTax1 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurchargeTax1() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax1) {
		return true
	}

	return false
}

// SetCustomSurchargeTax1 gets a reference to the given bool and assigns it to the CustomSurchargeTax1 field.
func (o *InvoiceRequest) SetCustomSurchargeTax1(v bool) {
	o.CustomSurchargeTax1 = &v
}

// GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurchargeTax2() bool {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax2
}

// GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurchargeTax2Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		return nil, false
	}
	return o.CustomSurchargeTax2, true
}

// HasCustomSurchargeTax2 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurchargeTax2() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax2) {
		return true
	}

	return false
}

// SetCustomSurchargeTax2 gets a reference to the given bool and assigns it to the CustomSurchargeTax2 field.
func (o *InvoiceRequest) SetCustomSurchargeTax2(v bool) {
	o.CustomSurchargeTax2 = &v
}

// GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurchargeTax3() bool {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax3
}

// GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurchargeTax3Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		return nil, false
	}
	return o.CustomSurchargeTax3, true
}

// HasCustomSurchargeTax3 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurchargeTax3() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax3) {
		return true
	}

	return false
}

// SetCustomSurchargeTax3 gets a reference to the given bool and assigns it to the CustomSurchargeTax3 field.
func (o *InvoiceRequest) SetCustomSurchargeTax3(v bool) {
	o.CustomSurchargeTax3 = &v
}

// GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field value if set, zero value otherwise.
func (o *InvoiceRequest) GetCustomSurchargeTax4() bool {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax4
}

// GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetCustomSurchargeTax4Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		return nil, false
	}
	return o.CustomSurchargeTax4, true
}

// HasCustomSurchargeTax4 returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCustomSurchargeTax4() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax4) {
		return true
	}

	return false
}

// SetCustomSurchargeTax4 gets a reference to the given bool and assigns it to the CustomSurchargeTax4 field.
func (o *InvoiceRequest) SetCustomSurchargeTax4(v bool) {
	o.CustomSurchargeTax4 = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InvoiceRequest) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InvoiceRequest) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InvoiceRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o InvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceRequest) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	toSerialize["client_id"] = o.ClientId
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.PoNumber) {
		toSerialize["po_number"] = o.PoNumber
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.PublicNotes) {
		toSerialize["public_notes"] = o.PublicNotes
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !IsNil(o.CustomValue1) {
		toSerialize["custom_value1"] = o.CustomValue1
	}
	if !IsNil(o.CustomValue2) {
		toSerialize["custom_value2"] = o.CustomValue2
	}
	if !IsNil(o.CustomValue3) {
		toSerialize["custom_value3"] = o.CustomValue3
	}
	if !IsNil(o.CustomValue4) {
		toSerialize["custom_value4"] = o.CustomValue4
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxRate2) {
		toSerialize["tax_rate2"] = o.TaxRate2
	}
	if !IsNil(o.TaxName3) {
		toSerialize["tax_name3"] = o.TaxName3
	}
	if !IsNil(o.TaxRate3) {
		toSerialize["tax_rate3"] = o.TaxRate3
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Invitations) {
		toSerialize["invitations"] = o.Invitations
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Partial) {
		toSerialize["partial"] = o.Partial
	}
	if !IsNil(o.IsAmountDiscount) {
		toSerialize["is_amount_discount"] = o.IsAmountDiscount
	}
	if !IsNil(o.UsesInclusiveTaxes) {
		toSerialize["uses_inclusive_taxes"] = o.UsesInclusiveTaxes
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.PartialDueDate) {
		toSerialize["partial_due_date"] = o.PartialDueDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.CustomSurcharge1) {
		toSerialize["custom_surcharge1"] = o.CustomSurcharge1
	}
	if !IsNil(o.CustomSurcharge2) {
		toSerialize["custom_surcharge2"] = o.CustomSurcharge2
	}
	if !IsNil(o.CustomSurcharge3) {
		toSerialize["custom_surcharge3"] = o.CustomSurcharge3
	}
	if !IsNil(o.CustomSurcharge4) {
		toSerialize["custom_surcharge4"] = o.CustomSurcharge4
	}
	if !IsNil(o.CustomSurchargeTax1) {
		toSerialize["custom_surcharge_tax1"] = o.CustomSurchargeTax1
	}
	if !IsNil(o.CustomSurchargeTax2) {
		toSerialize["custom_surcharge_tax2"] = o.CustomSurchargeTax2
	}
	if !IsNil(o.CustomSurchargeTax3) {
		toSerialize["custom_surcharge_tax3"] = o.CustomSurchargeTax3
	}
	if !IsNil(o.CustomSurchargeTax4) {
		toSerialize["custom_surcharge_tax4"] = o.CustomSurchargeTax4
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	return toSerialize, nil
}

func (o *InvoiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
	}

	allProperties := make(map[string]any)

	err = json.Unmarshal(data, &allProperties)
	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceRequest := _InvoiceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceRequest)
	if err != nil {
		return err
	}

	*o = InvoiceRequest(varInvoiceRequest)

	return err
}

type NullableInvoiceRequest struct {
	value *InvoiceRequest
	isSet bool
}

func (v NullableInvoiceRequest) Get() *InvoiceRequest {
	return v.value
}

func (v *NullableInvoiceRequest) Set(val *InvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceRequest(val *InvoiceRequest) *NullableInvoiceRequest {
	return &NullableInvoiceRequest{value: val, isSet: true}
}

func (v NullableInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
