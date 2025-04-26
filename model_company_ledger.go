package openapi

import (
	"encoding/json"
)

// checks if the CompanyLedger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyLedger{}

// CompanyLedger struct for CompanyLedger
type CompanyLedger struct {
	// This field will reference one of the following entity hashed ID payment_id, invoice_id or credit_id
	EntityId *string `json:"entity_id,omitempty"`
	// The notes which reference this entry of the ledger
	Notes *string `json:"notes,omitempty"`
	// The client balance
	Balance *float32 `json:"balance,omitempty"`
	// The amount the client balance is adjusted by
	Adjustment *float32 `json:"adjustment,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
}

// NewCompanyLedger instantiates a new CompanyLedger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyLedger() *CompanyLedger {
	this := CompanyLedger{}
	return &this
}

// NewCompanyLedgerWithDefaults instantiates a new CompanyLedger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyLedgerWithDefaults() *CompanyLedger {
	this := CompanyLedger{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CompanyLedger) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CompanyLedger) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *CompanyLedger) SetEntityId(v string) {
	o.EntityId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CompanyLedger) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CompanyLedger) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CompanyLedger) SetNotes(v string) {
	o.Notes = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *CompanyLedger) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *CompanyLedger) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *CompanyLedger) SetBalance(v float32) {
	o.Balance = &v
}

// GetAdjustment returns the Adjustment field value if set, zero value otherwise.
func (o *CompanyLedger) GetAdjustment() float32 {
	if o == nil || IsNil(o.Adjustment) {
		var ret float32
		return ret
	}
	return *o.Adjustment
}

// GetAdjustmentOk returns a tuple with the Adjustment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetAdjustmentOk() (*float32, bool) {
	if o == nil || IsNil(o.Adjustment) {
		return nil, false
	}
	return o.Adjustment, true
}

// HasAdjustment returns a boolean if a field has been set.
func (o *CompanyLedger) HasAdjustment() bool {
	if o != nil && !IsNil(o.Adjustment) {
		return true
	}

	return false
}

// SetAdjustment gets a reference to the given float32 and assigns it to the Adjustment field.
func (o *CompanyLedger) SetAdjustment(v float32) {
	o.Adjustment = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CompanyLedger) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CompanyLedger) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *CompanyLedger) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CompanyLedger) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLedger) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CompanyLedger) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *CompanyLedger) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o CompanyLedger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyLedger) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Adjustment) {
		toSerialize["adjustment"] = o.Adjustment
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableCompanyLedger struct {
	value *CompanyLedger
	isSet bool
}

func (v NullableCompanyLedger) Get() *CompanyLedger {
	return v.value
}

func (v *NullableCompanyLedger) Set(val *CompanyLedger) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyLedger) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyLedger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyLedger(val *CompanyLedger) *NullableCompanyLedger {
	return &NullableCompanyLedger{value: val, isSet: true}
}

func (v NullableCompanyLedger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyLedger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
