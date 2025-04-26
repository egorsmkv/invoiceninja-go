package openapi

import (
	"encoding/json"
	"time"
)

// checks if the InvoiceInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceInvitation{}

// InvoiceInvitation struct for InvoiceInvitation
type InvoiceInvitation struct {
	// The entity invitation hashed id
	Id *string `json:"id,omitempty"`
	// The client contact hashed id
	ClientContactId *string `json:"client_contact_id,omitempty"`
	// The invitation key
	Key *string `json:"key,omitempty"`
	// The invitation link
	Link *string `json:"link,omitempty"`
	// The invitation sent date
	SentDate *time.Time `json:"sent_date,omitempty"`
	// The invitation viewed date
	ViewedDate *time.Time `json:"viewed_date,omitempty"`
	// The invitation opened date
	OpenedDate *time.Time `json:"opened_date,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// The email error
	EmailError *string `json:"email_error,omitempty"`
	// The email status
	EmailStatus *string `json:"email_status,omitempty"`
}

// NewInvoiceInvitation instantiates a new InvoiceInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceInvitation() *InvoiceInvitation {
	this := InvoiceInvitation{}
	return &this
}

// NewInvoiceInvitationWithDefaults instantiates a new InvoiceInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceInvitationWithDefaults() *InvoiceInvitation {
	this := InvoiceInvitation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvoiceInvitation) SetId(v string) {
	o.Id = &v
}

// GetClientContactId returns the ClientContactId field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetClientContactId() string {
	if o == nil || IsNil(o.ClientContactId) {
		var ret string
		return ret
	}
	return *o.ClientContactId
}

// GetClientContactIdOk returns a tuple with the ClientContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetClientContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientContactId) {
		return nil, false
	}
	return o.ClientContactId, true
}

// HasClientContactId returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasClientContactId() bool {
	if o != nil && !IsNil(o.ClientContactId) {
		return true
	}

	return false
}

// SetClientContactId gets a reference to the given string and assigns it to the ClientContactId field.
func (o *InvoiceInvitation) SetClientContactId(v string) {
	o.ClientContactId = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InvoiceInvitation) SetKey(v string) {
	o.Key = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *InvoiceInvitation) SetLink(v string) {
	o.Link = &v
}

// GetSentDate returns the SentDate field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetSentDate() time.Time {
	if o == nil || IsNil(o.SentDate) {
		var ret time.Time
		return ret
	}
	return *o.SentDate
}

// GetSentDateOk returns a tuple with the SentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetSentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentDate) {
		return nil, false
	}
	return o.SentDate, true
}

// HasSentDate returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasSentDate() bool {
	if o != nil && !IsNil(o.SentDate) {
		return true
	}

	return false
}

// SetSentDate gets a reference to the given time.Time and assigns it to the SentDate field.
func (o *InvoiceInvitation) SetSentDate(v time.Time) {
	o.SentDate = &v
}

// GetViewedDate returns the ViewedDate field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetViewedDate() time.Time {
	if o == nil || IsNil(o.ViewedDate) {
		var ret time.Time
		return ret
	}
	return *o.ViewedDate
}

// GetViewedDateOk returns a tuple with the ViewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetViewedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ViewedDate) {
		return nil, false
	}
	return o.ViewedDate, true
}

// HasViewedDate returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasViewedDate() bool {
	if o != nil && !IsNil(o.ViewedDate) {
		return true
	}

	return false
}

// SetViewedDate gets a reference to the given time.Time and assigns it to the ViewedDate field.
func (o *InvoiceInvitation) SetViewedDate(v time.Time) {
	o.ViewedDate = &v
}

// GetOpenedDate returns the OpenedDate field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetOpenedDate() time.Time {
	if o == nil || IsNil(o.OpenedDate) {
		var ret time.Time
		return ret
	}
	return *o.OpenedDate
}

// GetOpenedDateOk returns a tuple with the OpenedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetOpenedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpenedDate) {
		return nil, false
	}
	return o.OpenedDate, true
}

// HasOpenedDate returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasOpenedDate() bool {
	if o != nil && !IsNil(o.OpenedDate) {
		return true
	}

	return false
}

// SetOpenedDate gets a reference to the given time.Time and assigns it to the OpenedDate field.
func (o *InvoiceInvitation) SetOpenedDate(v time.Time) {
	o.OpenedDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *InvoiceInvitation) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *InvoiceInvitation) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetEmailError returns the EmailError field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetEmailError() string {
	if o == nil || IsNil(o.EmailError) {
		var ret string
		return ret
	}
	return *o.EmailError
}

// GetEmailErrorOk returns a tuple with the EmailError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetEmailErrorOk() (*string, bool) {
	if o == nil || IsNil(o.EmailError) {
		return nil, false
	}
	return o.EmailError, true
}

// HasEmailError returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasEmailError() bool {
	if o != nil && !IsNil(o.EmailError) {
		return true
	}

	return false
}

// SetEmailError gets a reference to the given string and assigns it to the EmailError field.
func (o *InvoiceInvitation) SetEmailError(v string) {
	o.EmailError = &v
}

// GetEmailStatus returns the EmailStatus field value if set, zero value otherwise.
func (o *InvoiceInvitation) GetEmailStatus() string {
	if o == nil || IsNil(o.EmailStatus) {
		var ret string
		return ret
	}
	return *o.EmailStatus
}

// GetEmailStatusOk returns a tuple with the EmailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitation) GetEmailStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmailStatus) {
		return nil, false
	}
	return o.EmailStatus, true
}

// HasEmailStatus returns a boolean if a field has been set.
func (o *InvoiceInvitation) HasEmailStatus() bool {
	if o != nil && !IsNil(o.EmailStatus) {
		return true
	}

	return false
}

// SetEmailStatus gets a reference to the given string and assigns it to the EmailStatus field.
func (o *InvoiceInvitation) SetEmailStatus(v string) {
	o.EmailStatus = &v
}

func (o InvoiceInvitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClientContactId) {
		toSerialize["client_contact_id"] = o.ClientContactId
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.SentDate) {
		toSerialize["sent_date"] = o.SentDate
	}
	if !IsNil(o.ViewedDate) {
		toSerialize["viewed_date"] = o.ViewedDate
	}
	if !IsNil(o.OpenedDate) {
		toSerialize["opened_date"] = o.OpenedDate
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	if !IsNil(o.EmailError) {
		toSerialize["email_error"] = o.EmailError
	}
	if !IsNil(o.EmailStatus) {
		toSerialize["email_status"] = o.EmailStatus
	}
	return toSerialize, nil
}

type NullableInvoiceInvitation struct {
	value *InvoiceInvitation
	isSet bool
}

func (v NullableInvoiceInvitation) Get() *InvoiceInvitation {
	return v.value
}

func (v *NullableInvoiceInvitation) Set(val *InvoiceInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceInvitation(val *InvoiceInvitation) *NullableInvoiceInvitation {
	return &NullableInvoiceInvitation{value: val, isSet: true}
}

func (v NullableInvoiceInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
