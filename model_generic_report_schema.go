package openapi

import (
	"encoding/json"
)

// checks if the GenericReportSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericReportSchema{}

// GenericReportSchema struct for GenericReportSchema
type GenericReportSchema struct {
	// The string representation of the date range of data to be returned
	DateRange *string `json:"date_range,omitempty"`
	// The date column to search between.
	DateKey *string `json:"date_key,omitempty"`
	// The start date to search between
	StartDate *string `json:"start_date,omitempty"`
	// The end date to search between
	EndDate    *string  `json:"end_date,omitempty"`
	ReportKeys []string `json:"report_keys,omitempty"`
}

// NewGenericReportSchema instantiates a new GenericReportSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericReportSchema() *GenericReportSchema {
	this := GenericReportSchema{}
	return &this
}

// NewGenericReportSchemaWithDefaults instantiates a new GenericReportSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericReportSchemaWithDefaults() *GenericReportSchema {
	this := GenericReportSchema{}
	return &this
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *GenericReportSchema) GetDateRange() string {
	if o == nil || IsNil(o.DateRange) {
		var ret string
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericReportSchema) GetDateRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *GenericReportSchema) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given string and assigns it to the DateRange field.
func (o *GenericReportSchema) SetDateRange(v string) {
	o.DateRange = &v
}

// GetDateKey returns the DateKey field value if set, zero value otherwise.
func (o *GenericReportSchema) GetDateKey() string {
	if o == nil || IsNil(o.DateKey) {
		var ret string
		return ret
	}
	return *o.DateKey
}

// GetDateKeyOk returns a tuple with the DateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericReportSchema) GetDateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DateKey) {
		return nil, false
	}
	return o.DateKey, true
}

// HasDateKey returns a boolean if a field has been set.
func (o *GenericReportSchema) HasDateKey() bool {
	if o != nil && !IsNil(o.DateKey) {
		return true
	}

	return false
}

// SetDateKey gets a reference to the given string and assigns it to the DateKey field.
func (o *GenericReportSchema) SetDateKey(v string) {
	o.DateKey = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GenericReportSchema) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericReportSchema) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GenericReportSchema) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *GenericReportSchema) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GenericReportSchema) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericReportSchema) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GenericReportSchema) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *GenericReportSchema) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReportKeys returns the ReportKeys field value if set, zero value otherwise.
func (o *GenericReportSchema) GetReportKeys() []string {
	if o == nil || IsNil(o.ReportKeys) {
		var ret []string
		return ret
	}
	return o.ReportKeys
}

// GetReportKeysOk returns a tuple with the ReportKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericReportSchema) GetReportKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ReportKeys) {
		return nil, false
	}
	return o.ReportKeys, true
}

// HasReportKeys returns a boolean if a field has been set.
func (o *GenericReportSchema) HasReportKeys() bool {
	if o != nil && !IsNil(o.ReportKeys) {
		return true
	}

	return false
}

// SetReportKeys gets a reference to the given []string and assigns it to the ReportKeys field.
func (o *GenericReportSchema) SetReportKeys(v []string) {
	o.ReportKeys = v
}

func (o GenericReportSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericReportSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateRange) {
		toSerialize["date_range"] = o.DateRange
	}
	if !IsNil(o.DateKey) {
		toSerialize["date_key"] = o.DateKey
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ReportKeys) {
		toSerialize["report_keys"] = o.ReportKeys
	}
	return toSerialize, nil
}

type NullableGenericReportSchema struct {
	value *GenericReportSchema
	isSet bool
}

func (v NullableGenericReportSchema) Get() *GenericReportSchema {
	return v.value
}

func (v *NullableGenericReportSchema) Set(val *GenericReportSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericReportSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericReportSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericReportSchema(val *GenericReportSchema) *NullableGenericReportSchema {
	return &NullableGenericReportSchema{value: val, isSet: true}
}

func (v NullableGenericReportSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericReportSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
