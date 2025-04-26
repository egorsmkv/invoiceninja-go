package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the InvoiceInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceInvitationRequest{}

// InvoiceInvitationRequest struct for InvoiceInvitationRequest
type InvoiceInvitationRequest struct {
	// The entity invitation hashed id
	Id *string `json:"id,omitempty"`
	// The client contact hashed id
	ClientContactId string `json:"client_contact_id"`
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

type _InvoiceInvitationRequest InvoiceInvitationRequest

// NewInvoiceInvitationRequest instantiates a new InvoiceInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceInvitationRequest(clientContactId string) *InvoiceInvitationRequest {
	this := InvoiceInvitationRequest{}
	this.ClientContactId = clientContactId
	return &this
}

// NewInvoiceInvitationRequestWithDefaults instantiates a new InvoiceInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceInvitationRequestWithDefaults() *InvoiceInvitationRequest {
	this := InvoiceInvitationRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvoiceInvitationRequest) SetId(v string) {
	o.Id = &v
}

// GetClientContactId returns the ClientContactId field value
func (o *InvoiceInvitationRequest) GetClientContactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientContactId
}

// GetClientContactIdOk returns a tuple with the ClientContactId field value
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetClientContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientContactId, true
}

// SetClientContactId sets field value
func (o *InvoiceInvitationRequest) SetClientContactId(v string) {
	o.ClientContactId = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InvoiceInvitationRequest) SetKey(v string) {
	o.Key = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *InvoiceInvitationRequest) SetLink(v string) {
	o.Link = &v
}

// GetSentDate returns the SentDate field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetSentDate() time.Time {
	if o == nil || IsNil(o.SentDate) {
		var ret time.Time
		return ret
	}
	return *o.SentDate
}

// GetSentDateOk returns a tuple with the SentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetSentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentDate) {
		return nil, false
	}
	return o.SentDate, true
}

// HasSentDate returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasSentDate() bool {
	if o != nil && !IsNil(o.SentDate) {
		return true
	}

	return false
}

// SetSentDate gets a reference to the given time.Time and assigns it to the SentDate field.
func (o *InvoiceInvitationRequest) SetSentDate(v time.Time) {
	o.SentDate = &v
}

// GetViewedDate returns the ViewedDate field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetViewedDate() time.Time {
	if o == nil || IsNil(o.ViewedDate) {
		var ret time.Time
		return ret
	}
	return *o.ViewedDate
}

// GetViewedDateOk returns a tuple with the ViewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetViewedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ViewedDate) {
		return nil, false
	}
	return o.ViewedDate, true
}

// HasViewedDate returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasViewedDate() bool {
	if o != nil && !IsNil(o.ViewedDate) {
		return true
	}

	return false
}

// SetViewedDate gets a reference to the given time.Time and assigns it to the ViewedDate field.
func (o *InvoiceInvitationRequest) SetViewedDate(v time.Time) {
	o.ViewedDate = &v
}

// GetOpenedDate returns the OpenedDate field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetOpenedDate() time.Time {
	if o == nil || IsNil(o.OpenedDate) {
		var ret time.Time
		return ret
	}
	return *o.OpenedDate
}

// GetOpenedDateOk returns a tuple with the OpenedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetOpenedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpenedDate) {
		return nil, false
	}
	return o.OpenedDate, true
}

// HasOpenedDate returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasOpenedDate() bool {
	if o != nil && !IsNil(o.OpenedDate) {
		return true
	}

	return false
}

// SetOpenedDate gets a reference to the given time.Time and assigns it to the OpenedDate field.
func (o *InvoiceInvitationRequest) SetOpenedDate(v time.Time) {
	o.OpenedDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *InvoiceInvitationRequest) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *InvoiceInvitationRequest) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetEmailError returns the EmailError field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetEmailError() string {
	if o == nil || IsNil(o.EmailError) {
		var ret string
		return ret
	}
	return *o.EmailError
}

// GetEmailErrorOk returns a tuple with the EmailError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetEmailErrorOk() (*string, bool) {
	if o == nil || IsNil(o.EmailError) {
		return nil, false
	}
	return o.EmailError, true
}

// HasEmailError returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasEmailError() bool {
	if o != nil && !IsNil(o.EmailError) {
		return true
	}

	return false
}

// SetEmailError gets a reference to the given string and assigns it to the EmailError field.
func (o *InvoiceInvitationRequest) SetEmailError(v string) {
	o.EmailError = &v
}

// GetEmailStatus returns the EmailStatus field value if set, zero value otherwise.
func (o *InvoiceInvitationRequest) GetEmailStatus() string {
	if o == nil || IsNil(o.EmailStatus) {
		var ret string
		return ret
	}
	return *o.EmailStatus
}

// GetEmailStatusOk returns a tuple with the EmailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInvitationRequest) GetEmailStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmailStatus) {
		return nil, false
	}
	return o.EmailStatus, true
}

// HasEmailStatus returns a boolean if a field has been set.
func (o *InvoiceInvitationRequest) HasEmailStatus() bool {
	if o != nil && !IsNil(o.EmailStatus) {
		return true
	}

	return false
}

// SetEmailStatus gets a reference to the given string and assigns it to the EmailStatus field.
func (o *InvoiceInvitationRequest) SetEmailStatus(v string) {
	o.EmailStatus = &v
}

func (o InvoiceInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["client_contact_id"] = o.ClientContactId
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

func (o *InvoiceInvitationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_contact_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)
	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceInvitationRequest := _InvoiceInvitationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceInvitationRequest)
	if err != nil {
		return err
	}

	*o = InvoiceInvitationRequest(varInvoiceInvitationRequest)

	return err
}

type NullableInvoiceInvitationRequest struct {
	value *InvoiceInvitationRequest
	isSet bool
}

func (v NullableInvoiceInvitationRequest) Get() *InvoiceInvitationRequest {
	return v.value
}

func (v *NullableInvoiceInvitationRequest) Set(val *InvoiceInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceInvitationRequest(val *InvoiceInvitationRequest) *NullableInvoiceInvitationRequest {
	return &NullableInvoiceInvitationRequest{value: val, isSet: true}
}

func (v NullableInvoiceInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
