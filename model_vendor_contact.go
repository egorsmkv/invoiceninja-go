package openapi

import (
	"encoding/json"
)

// checks if the VendorContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorContact{}

// VendorContact struct for VendorContact
type VendorContact struct {
	// The hashed id of the vendor contact
	Id *string `json:"id,omitempty"`
	// The hashed id of the user id
	UserId *string `json:"user_id,omitempty"`
	// The hashed id of the vendor
	VendorId *string `json:"vendor_id,omitempty"`
	// The first name of the contact
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the contact
	LastName *string `json:"last_name,omitempty"`
	// A unique identifier for the contact
	ContactKey *string `json:"contact_key,omitempty"`
	// The confirmation code used to authenticate the contacts email address
	ConfirmationCode *string `json:"confirmation_code,omitempty"`
	// The contacts phone number
	Phone *string `json:"phone,omitempty"`
	// A custom value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The contact email address
	Email *string `json:"email,omitempty"`
	// The date which the contact confirmed their email
	EmailVerifiedAt *float32 `json:"email_verified_at,omitempty"`
	// The hashed password of the contact
	Password *string `json:"password,omitempty"`
	// Boolean flag determining if the contact is the primary contact for the vendor
	IsPrimary *bool `json:"is_primary,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	DeletedAt *float32 `json:"deleted_at,omitempty"`
}

// NewVendorContact instantiates a new VendorContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorContact() *VendorContact {
	this := VendorContact{}
	return &this
}

// NewVendorContactWithDefaults instantiates a new VendorContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorContactWithDefaults() *VendorContact {
	this := VendorContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VendorContact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VendorContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VendorContact) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *VendorContact) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *VendorContact) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *VendorContact) SetUserId(v string) {
	o.UserId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *VendorContact) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *VendorContact) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *VendorContact) SetVendorId(v string) {
	o.VendorId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *VendorContact) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *VendorContact) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *VendorContact) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *VendorContact) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *VendorContact) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *VendorContact) SetLastName(v string) {
	o.LastName = &v
}

// GetContactKey returns the ContactKey field value if set, zero value otherwise.
func (o *VendorContact) GetContactKey() string {
	if o == nil || IsNil(o.ContactKey) {
		var ret string
		return ret
	}
	return *o.ContactKey
}

// GetContactKeyOk returns a tuple with the ContactKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetContactKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ContactKey) {
		return nil, false
	}
	return o.ContactKey, true
}

// HasContactKey returns a boolean if a field has been set.
func (o *VendorContact) HasContactKey() bool {
	if o != nil && !IsNil(o.ContactKey) {
		return true
	}

	return false
}

// SetContactKey gets a reference to the given string and assigns it to the ContactKey field.
func (o *VendorContact) SetContactKey(v string) {
	o.ContactKey = &v
}

// GetConfirmationCode returns the ConfirmationCode field value if set, zero value otherwise.
func (o *VendorContact) GetConfirmationCode() string {
	if o == nil || IsNil(o.ConfirmationCode) {
		var ret string
		return ret
	}
	return *o.ConfirmationCode
}

// GetConfirmationCodeOk returns a tuple with the ConfirmationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetConfirmationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmationCode) {
		return nil, false
	}
	return o.ConfirmationCode, true
}

// HasConfirmationCode returns a boolean if a field has been set.
func (o *VendorContact) HasConfirmationCode() bool {
	if o != nil && !IsNil(o.ConfirmationCode) {
		return true
	}

	return false
}

// SetConfirmationCode gets a reference to the given string and assigns it to the ConfirmationCode field.
func (o *VendorContact) SetConfirmationCode(v string) {
	o.ConfirmationCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *VendorContact) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *VendorContact) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *VendorContact) SetPhone(v string) {
	o.Phone = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *VendorContact) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *VendorContact) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *VendorContact) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *VendorContact) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *VendorContact) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *VendorContact) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *VendorContact) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *VendorContact) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *VendorContact) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *VendorContact) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *VendorContact) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *VendorContact) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *VendorContact) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *VendorContact) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *VendorContact) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerifiedAt returns the EmailVerifiedAt field value if set, zero value otherwise.
func (o *VendorContact) GetEmailVerifiedAt() float32 {
	if o == nil || IsNil(o.EmailVerifiedAt) {
		var ret float32
		return ret
	}
	return *o.EmailVerifiedAt
}

// GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetEmailVerifiedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.EmailVerifiedAt) {
		return nil, false
	}
	return o.EmailVerifiedAt, true
}

// HasEmailVerifiedAt returns a boolean if a field has been set.
func (o *VendorContact) HasEmailVerifiedAt() bool {
	if o != nil && !IsNil(o.EmailVerifiedAt) {
		return true
	}

	return false
}

// SetEmailVerifiedAt gets a reference to the given float32 and assigns it to the EmailVerifiedAt field.
func (o *VendorContact) SetEmailVerifiedAt(v float32) {
	o.EmailVerifiedAt = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VendorContact) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VendorContact) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VendorContact) SetPassword(v string) {
	o.Password = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *VendorContact) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *VendorContact) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *VendorContact) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VendorContact) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VendorContact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *VendorContact) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VendorContact) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VendorContact) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *VendorContact) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *VendorContact) GetDeletedAt() float32 {
	if o == nil || IsNil(o.DeletedAt) {
		var ret float32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorContact) GetDeletedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *VendorContact) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given float32 and assigns it to the DeletedAt field.
func (o *VendorContact) SetDeletedAt(v float32) {
	o.DeletedAt = &v
}

func (o VendorContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.ContactKey) {
		toSerialize["contact_key"] = o.ContactKey
	}
	if !IsNil(o.ConfirmationCode) {
		toSerialize["confirmation_code"] = o.ConfirmationCode
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailVerifiedAt) {
		toSerialize["email_verified_at"] = o.EmailVerifiedAt
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return toSerialize, nil
}

type NullableVendorContact struct {
	value *VendorContact
	isSet bool
}

func (v NullableVendorContact) Get() *VendorContact {
	return v.value
}

func (v *NullableVendorContact) Set(val *VendorContact) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorContact) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorContact(val *VendorContact) *NullableVendorContact {
	return &NullableVendorContact{value: val, isSet: true}
}

func (v NullableVendorContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
