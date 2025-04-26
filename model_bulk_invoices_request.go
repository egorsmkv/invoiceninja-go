package openapi

import (
	"encoding/json"
)

// checks if the BulkInvoicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkInvoicesRequest{}

// BulkInvoicesRequest struct for BulkInvoicesRequest
type BulkInvoicesRequest struct {
	// The email type to be sent, when bulk emailing invoices
	EmailType *Enum `json:"email_type,omitempty"`
	// The action to be performed, options include:   - `bulk_download`     Bulk download an array of invoice PDFs (These are sent to the admin via email.)   - `download`     Download a single PDF. (Returns a single PDF object)   - `bulk_print`     Merges an array of Invoice PDFs for easy one click printing.   - `auto_bill`     Attempts to automatically bill the invoices with the payment method on file.   - `clone_to_invoice`     Returns a clone of the invoice.   - `clone_to_quote`     Returns a quote cloned using the properties of the given invoice.   - `mark_paid`     Marks an array of invoices as paid.   - `mark_sent`     Marks an array of invoices as sent.   - `restore`     Restores an array of invoices   - `delete`     Deletes an array of invoices   - `archive`     Archives an array of invoices   - `cancel`     Cancels an array of invoices   - `email`     Emails an array of invoices   - `send_email`     Emails an array of invoices. Requires additional properties to be sent. `email_type`
	Action *string  `json:"action,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

// NewBulkInvoicesRequest instantiates a new BulkInvoicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkInvoicesRequest() *BulkInvoicesRequest {
	this := BulkInvoicesRequest{}
	return &this
}

// NewBulkInvoicesRequestWithDefaults instantiates a new BulkInvoicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkInvoicesRequestWithDefaults() *BulkInvoicesRequest {
	this := BulkInvoicesRequest{}
	return &this
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetEmailType() Enum {
	if o == nil || IsNil(o.EmailType) {
		var ret Enum
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetEmailTypeOk() (*Enum, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given Enum and assigns it to the EmailType field.
func (o *BulkInvoicesRequest) SetEmailType(v Enum) {
	o.EmailType = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkInvoicesRequest) SetAction(v string) {
	o.Action = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BulkInvoicesRequest) SetIds(v []string) {
	o.Ids = v
}

func (o BulkInvoicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkInvoicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailType) {
		toSerialize["email_type"] = o.EmailType
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableBulkInvoicesRequest struct {
	value *BulkInvoicesRequest
	isSet bool
}

func (v NullableBulkInvoicesRequest) Get() *BulkInvoicesRequest {
	return v.value
}

func (v *NullableBulkInvoicesRequest) Set(val *BulkInvoicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkInvoicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkInvoicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkInvoicesRequest(val *BulkInvoicesRequest) *NullableBulkInvoicesRequest {
	return &NullableBulkInvoicesRequest{value: val, isSet: true}
}

func (v NullableBulkInvoicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkInvoicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
