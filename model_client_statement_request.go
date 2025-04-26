package openapi

import (
	"encoding/json"
)

// checks if the ClientStatementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientStatementRequest{}

// ClientStatementRequest struct for ClientStatementRequest
type ClientStatementRequest struct {
	// The start date of the statement period - format Y-m-d
	StartDate *string `json:"start_date,omitempty"`
	// The start date of the statement period - format Y-m-d
	EndDate *string `json:"end_date,omitempty"`
	// The hashed ID of the client
	ClientId *string `json:"client_id,omitempty"`
	// Flag which determines if the payments table is shown
	ShowPaymentsTable *bool `json:"show_payments_table,omitempty"`
	// Flag which determines if the credits table is shown
	ShowCreditsTable *bool `json:"show_credits_table,omitempty"`
	// Flag which determines if the aging table is shown
	ShowAgingTable *bool `json:"show_aging_table,omitempty"`
}

// NewClientStatementRequest instantiates a new ClientStatementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientStatementRequest() *ClientStatementRequest {
	this := ClientStatementRequest{}
	return &this
}

// NewClientStatementRequestWithDefaults instantiates a new ClientStatementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientStatementRequestWithDefaults() *ClientStatementRequest {
	this := ClientStatementRequest{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ClientStatementRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ClientStatementRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientStatementRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetShowPaymentsTable returns the ShowPaymentsTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowPaymentsTable() bool {
	if o == nil || IsNil(o.ShowPaymentsTable) {
		var ret bool
		return ret
	}
	return *o.ShowPaymentsTable
}

// GetShowPaymentsTableOk returns a tuple with the ShowPaymentsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowPaymentsTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPaymentsTable) {
		return nil, false
	}
	return o.ShowPaymentsTable, true
}

// HasShowPaymentsTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowPaymentsTable() bool {
	if o != nil && !IsNil(o.ShowPaymentsTable) {
		return true
	}

	return false
}

// SetShowPaymentsTable gets a reference to the given bool and assigns it to the ShowPaymentsTable field.
func (o *ClientStatementRequest) SetShowPaymentsTable(v bool) {
	o.ShowPaymentsTable = &v
}

// GetShowCreditsTable returns the ShowCreditsTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowCreditsTable() bool {
	if o == nil || IsNil(o.ShowCreditsTable) {
		var ret bool
		return ret
	}
	return *o.ShowCreditsTable
}

// GetShowCreditsTableOk returns a tuple with the ShowCreditsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowCreditsTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCreditsTable) {
		return nil, false
	}
	return o.ShowCreditsTable, true
}

// HasShowCreditsTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowCreditsTable() bool {
	if o != nil && !IsNil(o.ShowCreditsTable) {
		return true
	}

	return false
}

// SetShowCreditsTable gets a reference to the given bool and assigns it to the ShowCreditsTable field.
func (o *ClientStatementRequest) SetShowCreditsTable(v bool) {
	o.ShowCreditsTable = &v
}

// GetShowAgingTable returns the ShowAgingTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowAgingTable() bool {
	if o == nil || IsNil(o.ShowAgingTable) {
		var ret bool
		return ret
	}
	return *o.ShowAgingTable
}

// GetShowAgingTableOk returns a tuple with the ShowAgingTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowAgingTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAgingTable) {
		return nil, false
	}
	return o.ShowAgingTable, true
}

// HasShowAgingTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowAgingTable() bool {
	if o != nil && !IsNil(o.ShowAgingTable) {
		return true
	}

	return false
}

// SetShowAgingTable gets a reference to the given bool and assigns it to the ShowAgingTable field.
func (o *ClientStatementRequest) SetShowAgingTable(v bool) {
	o.ShowAgingTable = &v
}

func (o ClientStatementRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientStatementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ShowPaymentsTable) {
		toSerialize["show_payments_table"] = o.ShowPaymentsTable
	}
	if !IsNil(o.ShowCreditsTable) {
		toSerialize["show_credits_table"] = o.ShowCreditsTable
	}
	if !IsNil(o.ShowAgingTable) {
		toSerialize["show_aging_table"] = o.ShowAgingTable
	}
	return toSerialize, nil
}

type NullableClientStatementRequest struct {
	value *ClientStatementRequest
	isSet bool
}

func (v NullableClientStatementRequest) Get() *ClientStatementRequest {
	return v.value
}

func (v *NullableClientStatementRequest) Set(val *ClientStatementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientStatementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientStatementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientStatementRequest(val *ClientStatementRequest) *NullableClientStatementRequest {
	return &NullableClientStatementRequest{value: val, isSet: true}
}

func (v NullableClientStatementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientStatementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
