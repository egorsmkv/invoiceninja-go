package openapi

import (
	"encoding/json"
)

// checks if the BulkQuotesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkQuotesRequest{}

// BulkQuotesRequest struct for BulkQuotesRequest
type BulkQuotesRequest struct {
	// The action to be performed, options include:   - `approve`     Bulk approve an array of quotes   - `convert`     Bulk convert an array of quotes to invoices   - `send_email`     Bulk send an array of quotes as emails   - `mark_sent`     Bulk mark an array of quotes as sent   - `restore`     Restores an array of quotes   - `delete`     Deletes an array of invoices   - `archive`     Archives an array of invoices
	Action *string  `json:"action,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

// NewBulkQuotesRequest instantiates a new BulkQuotesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkQuotesRequest() *BulkQuotesRequest {
	this := BulkQuotesRequest{}
	return &this
}

// NewBulkQuotesRequestWithDefaults instantiates a new BulkQuotesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkQuotesRequestWithDefaults() *BulkQuotesRequest {
	this := BulkQuotesRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkQuotesRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQuotesRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkQuotesRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkQuotesRequest) SetAction(v string) {
	o.Action = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BulkQuotesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkQuotesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BulkQuotesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BulkQuotesRequest) SetIds(v []string) {
	o.Ids = v
}

func (o BulkQuotesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkQuotesRequest) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableBulkQuotesRequest struct {
	value *BulkQuotesRequest
	isSet bool
}

func (v NullableBulkQuotesRequest) Get() *BulkQuotesRequest {
	return v.value
}

func (v *NullableBulkQuotesRequest) Set(val *BulkQuotesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkQuotesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkQuotesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkQuotesRequest(val *BulkQuotesRequest) *NullableBulkQuotesRequest {
	return &NullableBulkQuotesRequest{value: val, isSet: true}
}

func (v NullableBulkQuotesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkQuotesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
