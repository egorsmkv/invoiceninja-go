package openapi

import (
	"encoding/json"
)

// checks if the ClientContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientContactRequest{}

// ClientContactRequest struct for ClientContactRequest
type ClientContactRequest struct {
	// The first name of the contact
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the contact
	LastName *string `json:"last_name,omitempty"`
	// The phone number of the contact
	Phone *string `json:"phone,omitempty"`
	// A Custom field value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A Custom field value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A Custom field value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A Custom field value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The email of the contact
	Email *string `json:"email,omitempty"`
	// The hashed password of the contact
	Password *string `json:"password,omitempty"`
	// Boolean flag that determines if this contact should receive emails
	SendEmail *bool `json:"send_email,omitempty"`
}

// NewClientContactRequest instantiates a new ClientContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientContactRequest() *ClientContactRequest {
	this := ClientContactRequest{}
	return &this
}

// NewClientContactRequestWithDefaults instantiates a new ClientContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientContactRequestWithDefaults() *ClientContactRequest {
	this := ClientContactRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ClientContactRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ClientContactRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ClientContactRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ClientContactRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ClientContactRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ClientContactRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ClientContactRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ClientContactRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ClientContactRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *ClientContactRequest) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *ClientContactRequest) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *ClientContactRequest) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *ClientContactRequest) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *ClientContactRequest) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *ClientContactRequest) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *ClientContactRequest) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *ClientContactRequest) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *ClientContactRequest) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *ClientContactRequest) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *ClientContactRequest) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *ClientContactRequest) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ClientContactRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ClientContactRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ClientContactRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ClientContactRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ClientContactRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ClientContactRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSendEmail returns the SendEmail field value if set, zero value otherwise.
func (o *ClientContactRequest) GetSendEmail() bool {
	if o == nil || IsNil(o.SendEmail) {
		var ret bool
		return ret
	}
	return *o.SendEmail
}

// GetSendEmailOk returns a tuple with the SendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContactRequest) GetSendEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmail) {
		return nil, false
	}
	return o.SendEmail, true
}

// HasSendEmail returns a boolean if a field has been set.
func (o *ClientContactRequest) HasSendEmail() bool {
	if o != nil && !IsNil(o.SendEmail) {
		return true
	}

	return false
}

// SetSendEmail gets a reference to the given bool and assigns it to the SendEmail field.
func (o *ClientContactRequest) SetSendEmail(v bool) {
	o.SendEmail = &v
}

func (o ClientContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientContactRequest) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
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
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SendEmail) {
		toSerialize["send_email"] = o.SendEmail
	}
	return toSerialize, nil
}

type NullableClientContactRequest struct {
	value *ClientContactRequest
	isSet bool
}

func (v NullableClientContactRequest) Get() *ClientContactRequest {
	return v.value
}

func (v *NullableClientContactRequest) Set(val *ClientContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientContactRequest(val *ClientContactRequest) *NullableClientContactRequest {
	return &NullableClientContactRequest{value: val, isSet: true}
}

func (v NullableClientContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
