package openapi

import (
	"encoding/json"
)

// checks if the Vendor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vendor{}

// Vendor struct for Vendor
type Vendor struct {
	// The hashed id of the vendor. This is a unique identifier for the vendor.
	Id *string `json:"id,omitempty"`
	// The hashed id of the user who created the vendor. This is a unique identifier for the user.
	UserId *string `json:"user_id,omitempty"`
	// The hashed id of the assigned user to this vendor. This is a unique identifier for the user.
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The name of the vendor.
	Name *string `json:"name,omitempty"`
	// The classification of the vendor.
	Classification *string `json:"classification,omitempty"`
	// The website of the vendor.
	Website *string `json:"website,omitempty"`
	// The private notes of the vendor. These notes are only visible to users with appropriate permissions.
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The industry id of the vendor. This is a unique identifier for the industry.
	IndustryId *string `json:"industry_id,omitempty"`
	// The size id of the vendor. This is a unique identifier for the size of the vendor.
	SizeId *string `json:"size_id,omitempty"`
	// The first line of the vendor's address.
	Address1 *string `json:"address1,omitempty"`
	// The second line of the vendor's address.
	Address2 *string `json:"address2,omitempty"`
	// The city of the vendor's address.
	City *string `json:"city,omitempty"`
	// The state of the vendor's address.
	State *string `json:"state,omitempty"`
	// The postal code of the vendor's address.
	PostalCode *string `json:"postal_code,omitempty"`
	// The phone number of the vendor.
	Phone *string `json:"phone,omitempty"`
	// The country id of the vendor. This is a unique identifier for the country.
	CountryId *string `json:"country_id,omitempty"`
	// The currency id of the vendor. This is a unique identifier for the currency.
	CurrencyId *string `json:"currency_id,omitempty"`
	// The value of the first custom field for the vendor.
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// The value of the second custom field for the vendor.
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// The value of the third custom field for the vendor.
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// The value of the fourth custom field for the vendor.
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The VAT number of the vendor.
	VatNumber *string `json:"vat_number,omitempty"`
	// The ID number of the vendor.
	IdNumber *string `json:"id_number,omitempty"`
	// The number of the vendor
	Number *string `json:"number,omitempty"`
	// Boolean flag determining if the vendor has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The language id of the vendor. This is a unique identifier for the language.
	LanguageId *string `json:"language_id,omitempty"`
	// The vendor hash of the vendor. This is a unique identifier for the vendor.
	VendorHash *string `json:"vendor_hash,omitempty"`
	// The transaction name of the vendor.
	TransactionName *string `json:"transaction_name,omitempty"`
	// Timestamp
	LastLogin *float32 `json:"last_login,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// The display name of the vendor.
	DisplayName *string `json:"display_name,omitempty"`
	// An array of contacts associated with the vendor.
	Contacts []VendorContact `json:"contacts,omitempty"`
}

// NewVendor instantiates a new Vendor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendor() *Vendor {
	this := Vendor{}
	return &this
}

// NewVendorWithDefaults instantiates a new Vendor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorWithDefaults() *Vendor {
	this := Vendor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vendor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vendor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vendor) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Vendor) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Vendor) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Vendor) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Vendor) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Vendor) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Vendor) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Vendor) GetContacts() []VendorContact {
	if o == nil || IsNil(o.Contacts) {
		var ret []VendorContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetContactsOk() ([]VendorContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Vendor) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []VendorContact and assigns it to the Contacts field.
func (o *Vendor) SetContacts(v []VendorContact) {
	o.Contacts = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vendor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vendor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vendor) SetName(v string) {
	o.Name = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *Vendor) GetClassification() string {
	if o == nil || IsNil(o.Classification) {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *Vendor) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *Vendor) SetClassification(v string) {
	o.Classification = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Vendor) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Vendor) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Vendor) SetWebsite(v string) {
	o.Website = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Vendor) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Vendor) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Vendor) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetIndustryId returns the IndustryId field value if set, zero value otherwise.
func (o *Vendor) GetIndustryId() string {
	if o == nil || IsNil(o.IndustryId) {
		var ret string
		return ret
	}
	return *o.IndustryId
}

// GetIndustryIdOk returns a tuple with the IndustryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetIndustryIdOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryId) {
		return nil, false
	}
	return o.IndustryId, true
}

// HasIndustryId returns a boolean if a field has been set.
func (o *Vendor) HasIndustryId() bool {
	if o != nil && !IsNil(o.IndustryId) {
		return true
	}

	return false
}

// SetIndustryId gets a reference to the given string and assigns it to the IndustryId field.
func (o *Vendor) SetIndustryId(v string) {
	o.IndustryId = &v
}

// GetSizeId returns the SizeId field value if set, zero value otherwise.
func (o *Vendor) GetSizeId() string {
	if o == nil || IsNil(o.SizeId) {
		var ret string
		return ret
	}
	return *o.SizeId
}

// GetSizeIdOk returns a tuple with the SizeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetSizeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SizeId) {
		return nil, false
	}
	return o.SizeId, true
}

// HasSizeId returns a boolean if a field has been set.
func (o *Vendor) HasSizeId() bool {
	if o != nil && !IsNil(o.SizeId) {
		return true
	}

	return false
}

// SetSizeId gets a reference to the given string and assigns it to the SizeId field.
func (o *Vendor) SetSizeId(v string) {
	o.SizeId = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *Vendor) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *Vendor) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *Vendor) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *Vendor) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *Vendor) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *Vendor) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Vendor) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Vendor) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Vendor) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Vendor) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Vendor) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Vendor) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Vendor) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Vendor) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Vendor) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Vendor) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Vendor) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Vendor) SetPhone(v string) {
	o.Phone = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *Vendor) GetCountryId() string {
	if o == nil || IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCountryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *Vendor) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *Vendor) SetCountryId(v string) {
	o.CountryId = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *Vendor) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *Vendor) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *Vendor) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Vendor) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Vendor) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Vendor) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Vendor) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Vendor) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Vendor) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Vendor) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Vendor) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Vendor) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Vendor) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Vendor) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Vendor) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Vendor) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Vendor) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Vendor) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *Vendor) GetIdNumber() string {
	if o == nil || IsNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IdNumber) {
		return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *Vendor) HasIdNumber() bool {
	if o != nil && !IsNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *Vendor) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Vendor) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Vendor) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Vendor) SetNumber(v string) {
	o.Number = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Vendor) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Vendor) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Vendor) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise.
func (o *Vendor) GetLanguageId() string {
	if o == nil || IsNil(o.LanguageId) {
		var ret string
		return ret
	}
	return *o.LanguageId
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetLanguageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageId) {
		return nil, false
	}
	return o.LanguageId, true
}

// HasLanguageId returns a boolean if a field has been set.
func (o *Vendor) HasLanguageId() bool {
	if o != nil && !IsNil(o.LanguageId) {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given string and assigns it to the LanguageId field.
func (o *Vendor) SetLanguageId(v string) {
	o.LanguageId = &v
}

// GetVendorHash returns the VendorHash field value if set, zero value otherwise.
func (o *Vendor) GetVendorHash() string {
	if o == nil || IsNil(o.VendorHash) {
		var ret string
		return ret
	}
	return *o.VendorHash
}

// GetVendorHashOk returns a tuple with the VendorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetVendorHashOk() (*string, bool) {
	if o == nil || IsNil(o.VendorHash) {
		return nil, false
	}
	return o.VendorHash, true
}

// HasVendorHash returns a boolean if a field has been set.
func (o *Vendor) HasVendorHash() bool {
	if o != nil && !IsNil(o.VendorHash) {
		return true
	}

	return false
}

// SetVendorHash gets a reference to the given string and assigns it to the VendorHash field.
func (o *Vendor) SetVendorHash(v string) {
	o.VendorHash = &v
}

// GetTransactionName returns the TransactionName field value if set, zero value otherwise.
func (o *Vendor) GetTransactionName() string {
	if o == nil || IsNil(o.TransactionName) {
		var ret string
		return ret
	}
	return *o.TransactionName
}

// GetTransactionNameOk returns a tuple with the TransactionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetTransactionNameOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionName) {
		return nil, false
	}
	return o.TransactionName, true
}

// HasTransactionName returns a boolean if a field has been set.
func (o *Vendor) HasTransactionName() bool {
	if o != nil && !IsNil(o.TransactionName) {
		return true
	}

	return false
}

// SetTransactionName gets a reference to the given string and assigns it to the TransactionName field.
func (o *Vendor) SetTransactionName(v string) {
	o.TransactionName = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *Vendor) GetLastLogin() float32 {
	if o == nil || IsNil(o.LastLogin) {
		var ret float32
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetLastLoginOk() (*float32, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *Vendor) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given float32 and assigns it to the LastLogin field.
func (o *Vendor) SetLastLogin(v float32) {
	o.LastLogin = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Vendor) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Vendor) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Vendor) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Vendor) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Vendor) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Vendor) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Vendor) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Vendor) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Vendor) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o Vendor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vendor) ToMap() (map[string]any, error) {
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
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.IndustryId) {
		toSerialize["industry_id"] = o.IndustryId
	}
	if !IsNil(o.SizeId) {
		toSerialize["size_id"] = o.SizeId
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
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
	if !IsNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !IsNil(o.IdNumber) {
		toSerialize["id_number"] = o.IdNumber
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.LanguageId) {
		toSerialize["language_id"] = o.LanguageId
	}
	if !IsNil(o.VendorHash) {
		toSerialize["vendor_hash"] = o.VendorHash
	}
	if !IsNil(o.TransactionName) {
		toSerialize["transaction_name"] = o.TransactionName
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableVendor struct {
	value *Vendor
	isSet bool
}

func (v NullableVendor) Get() *Vendor {
	return v.value
}

func (v *NullableVendor) Set(val *Vendor) {
	v.value = val
	v.isSet = true
}

func (v NullableVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendor(val *Vendor) *NullableVendor {
	return &NullableVendor{value: val, isSet: true}
}

func (v NullableVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
