package openapi

import (
	"encoding/json"
)

// checks if the BulkRecurringInvoicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkRecurringInvoicesRequest{}

// BulkRecurringInvoicesRequest struct for BulkRecurringInvoicesRequest
type BulkRecurringInvoicesRequest struct {
	// The action to be performed, options include:   - `start`     Starts (or restarts) the recurring invoice. **note** if the recurring invoice has been stopped for a long time, it will attempt to catch back up firing a new Invoice every hour per interval that has been missed.     If you do not wish to have the recurring invoice catch up, you should set the next_send_date to the correct date you wish the recurring invoice to commence from. - `stop`     Stops the recurring invoice.  - `send_now`     Force sends the recurring invoice - this option is only available when the recurring invoice is in a draft state.   - `restore`     Restores the recurring invoice from an archived or deleted state. - `archive`     Archives the recurring invoice. The recurring invoice will not fire in this state. - `delete`     Deletes a recurring invoice.
	Action *string  `json:"action,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

// NewBulkRecurringInvoicesRequest instantiates a new BulkRecurringInvoicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRecurringInvoicesRequest() *BulkRecurringInvoicesRequest {
	this := BulkRecurringInvoicesRequest{}
	return &this
}

// NewBulkRecurringInvoicesRequestWithDefaults instantiates a new BulkRecurringInvoicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRecurringInvoicesRequestWithDefaults() *BulkRecurringInvoicesRequest {
	this := BulkRecurringInvoicesRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkRecurringInvoicesRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRecurringInvoicesRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkRecurringInvoicesRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkRecurringInvoicesRequest) SetAction(v string) {
	o.Action = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BulkRecurringInvoicesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRecurringInvoicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BulkRecurringInvoicesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BulkRecurringInvoicesRequest) SetIds(v []string) {
	o.Ids = v
}

func (o BulkRecurringInvoicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkRecurringInvoicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableBulkRecurringInvoicesRequest struct {
	value *BulkRecurringInvoicesRequest
	isSet bool
}

func (v NullableBulkRecurringInvoicesRequest) Get() *BulkRecurringInvoicesRequest {
	return v.value
}

func (v *NullableBulkRecurringInvoicesRequest) Set(val *BulkRecurringInvoicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRecurringInvoicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRecurringInvoicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRecurringInvoicesRequest(val *BulkRecurringInvoicesRequest) *NullableBulkRecurringInvoicesRequest {
	return &NullableBulkRecurringInvoicesRequest{value: val, isSet: true}
}

func (v NullableBulkRecurringInvoicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRecurringInvoicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
