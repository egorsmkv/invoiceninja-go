package openapi

import (
	"encoding/json"
)

// checks if the FillableInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FillableInvoice{}

// FillableInvoice struct for FillableInvoice
type FillableInvoice struct {
	// The assigned user's hashed ID
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The client's hashed ID
	ClientId *string `json:"client_id,omitempty"`
	// The unique alphanumeric invoice number for each invoice per company
	Number *string `json:"number,omitempty"`
	// The purchase order number associated with the invoice
	PoNumber *string `json:"po_number,omitempty"`
	// The terms and conditions for the invoice
	Terms *string `json:"terms,omitempty"`
	// Public notes visible to the client on the invoice
	PublicNotes *string `json:"public_notes,omitempty"`
	// Private notes for internal use only
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The footer text displayed on the invoice
	Footer *string `json:"footer,omitempty"`
	// First custom value for additional information
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Second custom value for additional information
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Third custom value for additional information
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Fourth custom value for additional information
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// Name of the first tax applied to the invoice
	TaxName1 *string `json:"tax_name1,omitempty"`
	// Name of the second tax applied to the invoice
	TaxName2 *string `json:"tax_name2,omitempty"`
	// Rate of the first tax applied to the invoice
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// Rate of the second tax applied to the invoice
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// Name of the third tax applied to the invoice
	TaxName3 *string `json:"tax_name3,omitempty"`
	// Rate of the third tax applied to the invoice
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The discount applied to the invoice
	Discount *float32 `json:"discount,omitempty"`
	// The partial amount applied to the invoice
	Partial *float32 `json:"partial,omitempty"`
	// Indicates whether the discount applied is a fixed amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Indicates whether the tax rates applied to the invoice are inclusive or exclusive
	UsesInclusiveTaxes *bool `json:"uses_inclusive_taxes,omitempty"`
	// The date the invoice was issued
	Date *string `json:"date,omitempty"`
	// The due date for the partial payment
	PartialDueDate *string `json:"partial_due_date,omitempty"`
	// The due date for the invoice
	DueDate *string `json:"due_date,omitempty"`
	// First custom surcharge applied to the invoice
	CustomSurcharge1 *float32 `json:"custom_surcharge1,omitempty"`
	// Second custom surcharge applied to the invoice
	CustomSurcharge2 *float32 `json:"custom_surcharge2,omitempty"`
	// Third custom surcharge applied to the invoice
	CustomSurcharge3 *float32 `json:"custom_surcharge3,omitempty"`
	// Fourth custom surcharge applied to the invoice
	CustomSurcharge4 *float32 `json:"custom_surcharge4,omitempty"`
	// An array of objects which define the line items of the invoice
	LineItems []InvoiceItem `json:"line_items,omitempty"`
}

// NewFillableInvoice instantiates a new FillableInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFillableInvoice() *FillableInvoice {
	this := FillableInvoice{}
	return &this
}

// NewFillableInvoiceWithDefaults instantiates a new FillableInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFillableInvoiceWithDefaults() *FillableInvoice {
	this := FillableInvoice{}
	return &this
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *FillableInvoice) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *FillableInvoice) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *FillableInvoice) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *FillableInvoice) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *FillableInvoice) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *FillableInvoice) SetClientId(v string) {
	o.ClientId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *FillableInvoice) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *FillableInvoice) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *FillableInvoice) SetNumber(v string) {
	o.Number = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *FillableInvoice) GetPoNumber() string {
	if o == nil || IsNil(o.PoNumber) {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetPoNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PoNumber) {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *FillableInvoice) HasPoNumber() bool {
	if o != nil && !IsNil(o.PoNumber) {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *FillableInvoice) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *FillableInvoice) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *FillableInvoice) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *FillableInvoice) SetTerms(v string) {
	o.Terms = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *FillableInvoice) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *FillableInvoice) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *FillableInvoice) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *FillableInvoice) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *FillableInvoice) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *FillableInvoice) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *FillableInvoice) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *FillableInvoice) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *FillableInvoice) SetFooter(v string) {
	o.Footer = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *FillableInvoice) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *FillableInvoice) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *FillableInvoice) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *FillableInvoice) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *FillableInvoice) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *FillableInvoice) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *FillableInvoice) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *FillableInvoice) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *FillableInvoice) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *FillableInvoice) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *FillableInvoice) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *FillableInvoice) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *FillableInvoice) GetLineItems() []InvoiceItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []InvoiceItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetLineItemsOk() ([]InvoiceItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *FillableInvoice) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceItem and assigns it to the LineItems field.
func (o *FillableInvoice) SetLineItems(v []InvoiceItem) {
	o.LineItems = v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *FillableInvoice) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *FillableInvoice) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *FillableInvoice) SetDiscount(v float32) {
	o.Discount = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *FillableInvoice) GetPartial() float32 {
	if o == nil || IsNil(o.Partial) {
		var ret float32
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetPartialOk() (*float32, bool) {
	if o == nil || IsNil(o.Partial) {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *FillableInvoice) HasPartial() bool {
	if o != nil && !IsNil(o.Partial) {
		return true
	}

	return false
}

// SetPartial gets a reference to the given float32 and assigns it to the Partial field.
func (o *FillableInvoice) SetPartial(v float32) {
	o.Partial = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *FillableInvoice) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *FillableInvoice) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *FillableInvoice) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field value if set, zero value otherwise.
func (o *FillableInvoice) GetUsesInclusiveTaxes() bool {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.UsesInclusiveTaxes
}

// GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetUsesInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		return nil, false
	}
	return o.UsesInclusiveTaxes, true
}

// HasUsesInclusiveTaxes returns a boolean if a field has been set.
func (o *FillableInvoice) HasUsesInclusiveTaxes() bool {
	if o != nil && !IsNil(o.UsesInclusiveTaxes) {
		return true
	}

	return false
}

// SetUsesInclusiveTaxes gets a reference to the given bool and assigns it to the UsesInclusiveTaxes field.
func (o *FillableInvoice) SetUsesInclusiveTaxes(v bool) {
	o.UsesInclusiveTaxes = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *FillableInvoice) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *FillableInvoice) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *FillableInvoice) SetDate(v string) {
	o.Date = &v
}

// GetPartialDueDate returns the PartialDueDate field value if set, zero value otherwise.
func (o *FillableInvoice) GetPartialDueDate() string {
	if o == nil || IsNil(o.PartialDueDate) {
		var ret string
		return ret
	}
	return *o.PartialDueDate
}

// GetPartialDueDateOk returns a tuple with the PartialDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetPartialDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PartialDueDate) {
		return nil, false
	}
	return o.PartialDueDate, true
}

// HasPartialDueDate returns a boolean if a field has been set.
func (o *FillableInvoice) HasPartialDueDate() bool {
	if o != nil && !IsNil(o.PartialDueDate) {
		return true
	}

	return false
}

// SetPartialDueDate gets a reference to the given string and assigns it to the PartialDueDate field.
func (o *FillableInvoice) SetPartialDueDate(v string) {
	o.PartialDueDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *FillableInvoice) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *FillableInvoice) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *FillableInvoice) SetDueDate(v string) {
	o.DueDate = &v
}

// GetCustomSurcharge1 returns the CustomSurcharge1 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomSurcharge1() float32 {
	if o == nil || IsNil(o.CustomSurcharge1) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge1
}

// GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomSurcharge1Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge1) {
		return nil, false
	}
	return o.CustomSurcharge1, true
}

// HasCustomSurcharge1 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomSurcharge1() bool {
	if o != nil && !IsNil(o.CustomSurcharge1) {
		return true
	}

	return false
}

// SetCustomSurcharge1 gets a reference to the given float32 and assigns it to the CustomSurcharge1 field.
func (o *FillableInvoice) SetCustomSurcharge1(v float32) {
	o.CustomSurcharge1 = &v
}

// GetCustomSurcharge2 returns the CustomSurcharge2 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomSurcharge2() float32 {
	if o == nil || IsNil(o.CustomSurcharge2) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge2
}

// GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomSurcharge2Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge2) {
		return nil, false
	}
	return o.CustomSurcharge2, true
}

// HasCustomSurcharge2 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomSurcharge2() bool {
	if o != nil && !IsNil(o.CustomSurcharge2) {
		return true
	}

	return false
}

// SetCustomSurcharge2 gets a reference to the given float32 and assigns it to the CustomSurcharge2 field.
func (o *FillableInvoice) SetCustomSurcharge2(v float32) {
	o.CustomSurcharge2 = &v
}

// GetCustomSurcharge3 returns the CustomSurcharge3 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomSurcharge3() float32 {
	if o == nil || IsNil(o.CustomSurcharge3) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge3
}

// GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomSurcharge3Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge3) {
		return nil, false
	}
	return o.CustomSurcharge3, true
}

// HasCustomSurcharge3 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomSurcharge3() bool {
	if o != nil && !IsNil(o.CustomSurcharge3) {
		return true
	}

	return false
}

// SetCustomSurcharge3 gets a reference to the given float32 and assigns it to the CustomSurcharge3 field.
func (o *FillableInvoice) SetCustomSurcharge3(v float32) {
	o.CustomSurcharge3 = &v
}

// GetCustomSurcharge4 returns the CustomSurcharge4 field value if set, zero value otherwise.
func (o *FillableInvoice) GetCustomSurcharge4() float32 {
	if o == nil || IsNil(o.CustomSurcharge4) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge4
}

// GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FillableInvoice) GetCustomSurcharge4Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge4) {
		return nil, false
	}
	return o.CustomSurcharge4, true
}

// HasCustomSurcharge4 returns a boolean if a field has been set.
func (o *FillableInvoice) HasCustomSurcharge4() bool {
	if o != nil && !IsNil(o.CustomSurcharge4) {
		return true
	}

	return false
}

// SetCustomSurcharge4 gets a reference to the given float32 and assigns it to the CustomSurcharge4 field.
func (o *FillableInvoice) SetCustomSurcharge4(v float32) {
	o.CustomSurcharge4 = &v
}

func (o FillableInvoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FillableInvoice) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
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
	return toSerialize, nil
}

type NullableFillableInvoice struct {
	value *FillableInvoice
	isSet bool
}

func (v NullableFillableInvoice) Get() *FillableInvoice {
	return v.value
}

func (v *NullableFillableInvoice) Set(val *FillableInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableFillableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableFillableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFillableInvoice(val *FillableInvoice) *NullableFillableInvoice {
	return &NullableFillableInvoice{value: val, isSet: true}
}

func (v NullableFillableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFillableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
