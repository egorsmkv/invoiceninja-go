package openapi

import (
	"encoding/json"
)

// checks if the PurchaseOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseOrder{}

// PurchaseOrder struct for PurchaseOrder
type PurchaseOrder struct {
	// The unique hashed identifier for the purchase order
	Id *string `json:"id,omitempty"`
	// The unique hashed identifier for the user who created the purchase order
	UserId *string `json:"user_id,omitempty"`
	// The unique hashed identifier for the user assigned to the purchase order
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The unique hashed identifier for the vendor associated with the purchase order
	VendorId *string `json:"vendor_id,omitempty"`
	// The status of the purchase order represented by a unique identifier
	StatusId *string `json:"status_id,omitempty"`
	// The unique alpha-numeric purchase order number per company
	Number *string `json:"number,omitempty"`
	// The quote number associated with this purchase order
	QuoteNumber *string `json:"quote_number,omitempty"`
	// The terms and conditions for the purchase order
	Terms *string `json:"terms,omitempty"`
	// Publicly visible notes associated with the purchase order
	PublicNotes *string `json:"public_notes,omitempty"`
	// Privately visible notes associated with the purchase order, not disclosed to the vendor
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The footer text of the purchase order
	Footer *string `json:"footer,omitempty"`
	// First custom value field for additional information
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Second custom value field for additional information
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Third custom value field for additional information
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Fourth custom value field for additional information
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The name of the first tax applied to the purchase order
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The name of the second tax applied to the purchase order
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The rate of the first tax applied to the purchase order
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The rate of the second tax applied to the purchase order
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The name of the third tax applied to the purchase order
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The rate of the third tax applied to the purchase order
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The total amount of taxes applied to the purchase order
	TotalTaxes *float32 `json:"total_taxes,omitempty"`
	// The total amount of the purchase order before taxes and discounts
	Amount *float32 `json:"amount,omitempty"`
	// The balance due for the purchase order after accounting for payments
	Balance *float32 `json:"balance,omitempty"`
	// The total amount paid on the purchase order so far
	PaidToDate *float32 `json:"paid_to_date,omitempty"`
	// The discount amount or percentage applied to the purchase order
	Discount *float32 `json:"discount,omitempty"`
	// The partial or deposit amount for the purchase order
	Partial *float32 `json:"partial,omitempty"`
	// Boolean flag indicating if the discount is a fixed amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Boolean flag indicating if the purchase order has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Boolean flag indicating if the taxes used are inclusive or exclusive
	UsesInclusiveTaxes *bool `json:"uses_inclusive_taxes,omitempty"`
	// The date the purchase order was created
	Date *string `json:"date,omitempty"`
	// The last date the purchase order was sent to the vendor
	LastSentDate *string `json:"last_sent_date,omitempty"`
	// The next scheduled date for sending a reminder for the purchase order
	NextSendDate *string `json:"next_send_date,omitempty"`
	// The due date for the partial or deposit amount
	PartialDueDate *string `json:"partial_due_date,omitempty"`
	// The due date for the total amount of the purchase order
	DueDate  *string          `json:"due_date,omitempty"`
	Settings *CompanySettings `json:"settings,omitempty"`
	// Timestamp
	LastViewed *float32 `json:"last_viewed,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// First custom surcharge amount for the purchase order
	CustomSurcharge1 *float32 `json:"custom_surcharge1,omitempty"`
	// Second custom surcharge amount for the purchase order
	CustomSurcharge2 *float32 `json:"custom_surcharge2,omitempty"`
	// Third custom surcharge amount for the purchase order
	CustomSurcharge3 *float32 `json:"custom_surcharge3,omitempty"`
	// Fourth custom surcharge amount for the purchase order
	CustomSurcharge4 *float32 `json:"custom_surcharge4,omitempty"`
	// Boolean flag indicating if taxes are charged on the first custom surcharge amount
	CustomSurchargeTax1 *bool `json:"custom_surcharge_tax1,omitempty"`
	// Boolean flag indicating if taxes are charged on the second custom surcharge amount
	CustomSurchargeTax2 *bool `json:"custom_surcharge_tax2,omitempty"`
	// Boolean flag indicating if taxes are charged on the third custom surcharge amount
	CustomSurchargeTax3 *bool `json:"custom_surcharge_tax3,omitempty"`
	// Boolean flag indicating if taxes are charged on the fourth custom surcharge amount
	CustomSurchargeTax4 *bool `json:"custom_surcharge_tax4,omitempty"`
	// An array of objects which define the line items of the purchase order
	LineItems []InvoiceItem `json:"line_items,omitempty"`
}

// NewPurchaseOrder instantiates a new PurchaseOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrder() *PurchaseOrder {
	this := PurchaseOrder{}
	return &this
}

// NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderWithDefaults() *PurchaseOrder {
	this := PurchaseOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PurchaseOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PurchaseOrder) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PurchaseOrder) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PurchaseOrder) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *PurchaseOrder) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *PurchaseOrder) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *PurchaseOrder) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *PurchaseOrder) SetVendorId(v string) {
	o.VendorId = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *PurchaseOrder) GetStatusId() string {
	if o == nil || IsNil(o.StatusId) {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetStatusIdOk() (*string, bool) {
	if o == nil || IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasStatusId() bool {
	if o != nil && !IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *PurchaseOrder) SetStatusId(v string) {
	o.StatusId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PurchaseOrder) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PurchaseOrder) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *PurchaseOrder) SetNumber(v string) {
	o.Number = &v
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *PurchaseOrder) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *PurchaseOrder) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *PurchaseOrder) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *PurchaseOrder) SetTerms(v string) {
	o.Terms = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *PurchaseOrder) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *PurchaseOrder) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *PurchaseOrder) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *PurchaseOrder) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *PurchaseOrder) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *PurchaseOrder) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *PurchaseOrder) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *PurchaseOrder) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *PurchaseOrder) SetFooter(v string) {
	o.Footer = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *PurchaseOrder) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *PurchaseOrder) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *PurchaseOrder) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *PurchaseOrder) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *PurchaseOrder) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *PurchaseOrder) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *PurchaseOrder) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *PurchaseOrder) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *PurchaseOrder) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *PurchaseOrder) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetTotalTaxes returns the TotalTaxes field value if set, zero value otherwise.
func (o *PurchaseOrder) GetTotalTaxes() float32 {
	if o == nil || IsNil(o.TotalTaxes) {
		var ret float32
		return ret
	}
	return *o.TotalTaxes
}

// GetTotalTaxesOk returns a tuple with the TotalTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetTotalTaxesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTaxes) {
		return nil, false
	}
	return o.TotalTaxes, true
}

// HasTotalTaxes returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTotalTaxes() bool {
	if o != nil && !IsNil(o.TotalTaxes) {
		return true
	}

	return false
}

// SetTotalTaxes gets a reference to the given float32 and assigns it to the TotalTaxes field.
func (o *PurchaseOrder) SetTotalTaxes(v float32) {
	o.TotalTaxes = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PurchaseOrder) GetLineItems() []InvoiceItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []InvoiceItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetLineItemsOk() ([]InvoiceItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PurchaseOrder) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceItem and assigns it to the LineItems field.
func (o *PurchaseOrder) SetLineItems(v []InvoiceItem) {
	o.LineItems = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PurchaseOrder) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PurchaseOrder) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PurchaseOrder) SetAmount(v float32) {
	o.Amount = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *PurchaseOrder) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *PurchaseOrder) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *PurchaseOrder) SetBalance(v float32) {
	o.Balance = &v
}

// GetPaidToDate returns the PaidToDate field value if set, zero value otherwise.
func (o *PurchaseOrder) GetPaidToDate() float32 {
	if o == nil || IsNil(o.PaidToDate) {
		var ret float32
		return ret
	}
	return *o.PaidToDate
}

// GetPaidToDateOk returns a tuple with the PaidToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetPaidToDateOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidToDate) {
		return nil, false
	}
	return o.PaidToDate, true
}

// HasPaidToDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasPaidToDate() bool {
	if o != nil && !IsNil(o.PaidToDate) {
		return true
	}

	return false
}

// SetPaidToDate gets a reference to the given float32 and assigns it to the PaidToDate field.
func (o *PurchaseOrder) SetPaidToDate(v float32) {
	o.PaidToDate = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *PurchaseOrder) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *PurchaseOrder) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *PurchaseOrder) SetDiscount(v float32) {
	o.Discount = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *PurchaseOrder) GetPartial() float32 {
	if o == nil || IsNil(o.Partial) {
		var ret float32
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetPartialOk() (*float32, bool) {
	if o == nil || IsNil(o.Partial) {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *PurchaseOrder) HasPartial() bool {
	if o != nil && !IsNil(o.Partial) {
		return true
	}

	return false
}

// SetPartial gets a reference to the given float32 and assigns it to the Partial field.
func (o *PurchaseOrder) SetPartial(v float32) {
	o.Partial = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *PurchaseOrder) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *PurchaseOrder) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *PurchaseOrder) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *PurchaseOrder) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *PurchaseOrder) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *PurchaseOrder) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field value if set, zero value otherwise.
func (o *PurchaseOrder) GetUsesInclusiveTaxes() bool {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.UsesInclusiveTaxes
}

// GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetUsesInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		return nil, false
	}
	return o.UsesInclusiveTaxes, true
}

// HasUsesInclusiveTaxes returns a boolean if a field has been set.
func (o *PurchaseOrder) HasUsesInclusiveTaxes() bool {
	if o != nil && !IsNil(o.UsesInclusiveTaxes) {
		return true
	}

	return false
}

// SetUsesInclusiveTaxes gets a reference to the given bool and assigns it to the UsesInclusiveTaxes field.
func (o *PurchaseOrder) SetUsesInclusiveTaxes(v bool) {
	o.UsesInclusiveTaxes = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PurchaseOrder) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *PurchaseOrder) SetDate(v string) {
	o.Date = &v
}

// GetLastSentDate returns the LastSentDate field value if set, zero value otherwise.
func (o *PurchaseOrder) GetLastSentDate() string {
	if o == nil || IsNil(o.LastSentDate) {
		var ret string
		return ret
	}
	return *o.LastSentDate
}

// GetLastSentDateOk returns a tuple with the LastSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetLastSentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSentDate) {
		return nil, false
	}
	return o.LastSentDate, true
}

// HasLastSentDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasLastSentDate() bool {
	if o != nil && !IsNil(o.LastSentDate) {
		return true
	}

	return false
}

// SetLastSentDate gets a reference to the given string and assigns it to the LastSentDate field.
func (o *PurchaseOrder) SetLastSentDate(v string) {
	o.LastSentDate = &v
}

// GetNextSendDate returns the NextSendDate field value if set, zero value otherwise.
func (o *PurchaseOrder) GetNextSendDate() string {
	if o == nil || IsNil(o.NextSendDate) {
		var ret string
		return ret
	}
	return *o.NextSendDate
}

// GetNextSendDateOk returns a tuple with the NextSendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetNextSendDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextSendDate) {
		return nil, false
	}
	return o.NextSendDate, true
}

// HasNextSendDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasNextSendDate() bool {
	if o != nil && !IsNil(o.NextSendDate) {
		return true
	}

	return false
}

// SetNextSendDate gets a reference to the given string and assigns it to the NextSendDate field.
func (o *PurchaseOrder) SetNextSendDate(v string) {
	o.NextSendDate = &v
}

// GetPartialDueDate returns the PartialDueDate field value if set, zero value otherwise.
func (o *PurchaseOrder) GetPartialDueDate() string {
	if o == nil || IsNil(o.PartialDueDate) {
		var ret string
		return ret
	}
	return *o.PartialDueDate
}

// GetPartialDueDateOk returns a tuple with the PartialDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetPartialDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PartialDueDate) {
		return nil, false
	}
	return o.PartialDueDate, true
}

// HasPartialDueDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasPartialDueDate() bool {
	if o != nil && !IsNil(o.PartialDueDate) {
		return true
	}

	return false
}

// SetPartialDueDate gets a reference to the given string and assigns it to the PartialDueDate field.
func (o *PurchaseOrder) SetPartialDueDate(v string) {
	o.PartialDueDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *PurchaseOrder) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *PurchaseOrder) SetDueDate(v string) {
	o.DueDate = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *PurchaseOrder) GetSettings() CompanySettings {
	if o == nil || IsNil(o.Settings) {
		var ret CompanySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetSettingsOk() (*CompanySettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PurchaseOrder) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CompanySettings and assigns it to the Settings field.
func (o *PurchaseOrder) SetSettings(v CompanySettings) {
	o.Settings = &v
}

// GetLastViewed returns the LastViewed field value if set, zero value otherwise.
func (o *PurchaseOrder) GetLastViewed() float32 {
	if o == nil || IsNil(o.LastViewed) {
		var ret float32
		return ret
	}
	return *o.LastViewed
}

// GetLastViewedOk returns a tuple with the LastViewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetLastViewedOk() (*float32, bool) {
	if o == nil || IsNil(o.LastViewed) {
		return nil, false
	}
	return o.LastViewed, true
}

// HasLastViewed returns a boolean if a field has been set.
func (o *PurchaseOrder) HasLastViewed() bool {
	if o != nil && !IsNil(o.LastViewed) {
		return true
	}

	return false
}

// SetLastViewed gets a reference to the given float32 and assigns it to the LastViewed field.
func (o *PurchaseOrder) SetLastViewed(v float32) {
	o.LastViewed = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PurchaseOrder) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PurchaseOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *PurchaseOrder) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *PurchaseOrder) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *PurchaseOrder) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *PurchaseOrder) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetCustomSurcharge1 returns the CustomSurcharge1 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurcharge1() float32 {
	if o == nil || IsNil(o.CustomSurcharge1) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge1
}

// GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurcharge1Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge1) {
		return nil, false
	}
	return o.CustomSurcharge1, true
}

// HasCustomSurcharge1 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurcharge1() bool {
	if o != nil && !IsNil(o.CustomSurcharge1) {
		return true
	}

	return false
}

// SetCustomSurcharge1 gets a reference to the given float32 and assigns it to the CustomSurcharge1 field.
func (o *PurchaseOrder) SetCustomSurcharge1(v float32) {
	o.CustomSurcharge1 = &v
}

// GetCustomSurcharge2 returns the CustomSurcharge2 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurcharge2() float32 {
	if o == nil || IsNil(o.CustomSurcharge2) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge2
}

// GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurcharge2Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge2) {
		return nil, false
	}
	return o.CustomSurcharge2, true
}

// HasCustomSurcharge2 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurcharge2() bool {
	if o != nil && !IsNil(o.CustomSurcharge2) {
		return true
	}

	return false
}

// SetCustomSurcharge2 gets a reference to the given float32 and assigns it to the CustomSurcharge2 field.
func (o *PurchaseOrder) SetCustomSurcharge2(v float32) {
	o.CustomSurcharge2 = &v
}

// GetCustomSurcharge3 returns the CustomSurcharge3 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurcharge3() float32 {
	if o == nil || IsNil(o.CustomSurcharge3) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge3
}

// GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurcharge3Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge3) {
		return nil, false
	}
	return o.CustomSurcharge3, true
}

// HasCustomSurcharge3 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurcharge3() bool {
	if o != nil && !IsNil(o.CustomSurcharge3) {
		return true
	}

	return false
}

// SetCustomSurcharge3 gets a reference to the given float32 and assigns it to the CustomSurcharge3 field.
func (o *PurchaseOrder) SetCustomSurcharge3(v float32) {
	o.CustomSurcharge3 = &v
}

// GetCustomSurcharge4 returns the CustomSurcharge4 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurcharge4() float32 {
	if o == nil || IsNil(o.CustomSurcharge4) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge4
}

// GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurcharge4Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge4) {
		return nil, false
	}
	return o.CustomSurcharge4, true
}

// HasCustomSurcharge4 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurcharge4() bool {
	if o != nil && !IsNil(o.CustomSurcharge4) {
		return true
	}

	return false
}

// SetCustomSurcharge4 gets a reference to the given float32 and assigns it to the CustomSurcharge4 field.
func (o *PurchaseOrder) SetCustomSurcharge4(v float32) {
	o.CustomSurcharge4 = &v
}

// GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurchargeTax1() bool {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax1
}

// GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurchargeTax1Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		return nil, false
	}
	return o.CustomSurchargeTax1, true
}

// HasCustomSurchargeTax1 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurchargeTax1() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax1) {
		return true
	}

	return false
}

// SetCustomSurchargeTax1 gets a reference to the given bool and assigns it to the CustomSurchargeTax1 field.
func (o *PurchaseOrder) SetCustomSurchargeTax1(v bool) {
	o.CustomSurchargeTax1 = &v
}

// GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurchargeTax2() bool {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax2
}

// GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurchargeTax2Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		return nil, false
	}
	return o.CustomSurchargeTax2, true
}

// HasCustomSurchargeTax2 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurchargeTax2() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax2) {
		return true
	}

	return false
}

// SetCustomSurchargeTax2 gets a reference to the given bool and assigns it to the CustomSurchargeTax2 field.
func (o *PurchaseOrder) SetCustomSurchargeTax2(v bool) {
	o.CustomSurchargeTax2 = &v
}

// GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurchargeTax3() bool {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax3
}

// GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurchargeTax3Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		return nil, false
	}
	return o.CustomSurchargeTax3, true
}

// HasCustomSurchargeTax3 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurchargeTax3() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax3) {
		return true
	}

	return false
}

// SetCustomSurchargeTax3 gets a reference to the given bool and assigns it to the CustomSurchargeTax3 field.
func (o *PurchaseOrder) SetCustomSurchargeTax3(v bool) {
	o.CustomSurchargeTax3 = &v
}

// GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field value if set, zero value otherwise.
func (o *PurchaseOrder) GetCustomSurchargeTax4() bool {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax4
}

// GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetCustomSurchargeTax4Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		return nil, false
	}
	return o.CustomSurchargeTax4, true
}

// HasCustomSurchargeTax4 returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomSurchargeTax4() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax4) {
		return true
	}

	return false
}

// SetCustomSurchargeTax4 gets a reference to the given bool and assigns it to the CustomSurchargeTax4 field.
func (o *PurchaseOrder) SetCustomSurchargeTax4(v bool) {
	o.CustomSurchargeTax4 = &v
}

func (o PurchaseOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseOrder) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.StatusId) {
		toSerialize["status_id"] = o.StatusId
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quote_number"] = o.QuoteNumber
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
	if !IsNil(o.TotalTaxes) {
		toSerialize["total_taxes"] = o.TotalTaxes
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.PaidToDate) {
		toSerialize["paid_to_date"] = o.PaidToDate
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
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.UsesInclusiveTaxes) {
		toSerialize["uses_inclusive_taxes"] = o.UsesInclusiveTaxes
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.LastSentDate) {
		toSerialize["last_sent_date"] = o.LastSentDate
	}
	if !IsNil(o.NextSendDate) {
		toSerialize["next_send_date"] = o.NextSendDate
	}
	if !IsNil(o.PartialDueDate) {
		toSerialize["partial_due_date"] = o.PartialDueDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.LastViewed) {
		toSerialize["last_viewed"] = o.LastViewed
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
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
	return toSerialize, nil
}

type NullablePurchaseOrder struct {
	value *PurchaseOrder
	isSet bool
}

func (v NullablePurchaseOrder) Get() *PurchaseOrder {
	return v.value
}

func (v *NullablePurchaseOrder) Set(val *PurchaseOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrder(val *PurchaseOrder) *NullablePurchaseOrder {
	return &NullablePurchaseOrder{value: val, isSet: true}
}

func (v NullablePurchaseOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
