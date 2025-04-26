package openapi

import (
	"encoding/json"
)

// checks if the TaskSchedulerSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskSchedulerSchema{}

// TaskSchedulerSchema struct for TaskSchedulerSchema
type TaskSchedulerSchema struct {
	// The scheduler paused state
	Paused *bool `json:"paused,omitempty"`
	// Accepted values (DAY,WEEK,MONTH,3MONTHS,YEAR)
	RepeatEvery *string `json:"repeat_every,omitempty"`
	// Timestamp when we should start the scheduler, default is today
	StartFrom *int32 `json:"start_from,omitempty"`
	// Job, we can find list of available jobs in Scheduler model
	Job *string `json:"job,omitempty"`
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

// NewTaskSchedulerSchema instantiates a new TaskSchedulerSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskSchedulerSchema() *TaskSchedulerSchema {
	this := TaskSchedulerSchema{}
	return &this
}

// NewTaskSchedulerSchemaWithDefaults instantiates a new TaskSchedulerSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSchedulerSchemaWithDefaults() *TaskSchedulerSchema {
	this := TaskSchedulerSchema{}
	return &this
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *TaskSchedulerSchema) SetPaused(v bool) {
	o.Paused = &v
}

// GetRepeatEvery returns the RepeatEvery field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetRepeatEvery() string {
	if o == nil || IsNil(o.RepeatEvery) {
		var ret string
		return ret
	}
	return *o.RepeatEvery
}

// GetRepeatEveryOk returns a tuple with the RepeatEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetRepeatEveryOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatEvery) {
		return nil, false
	}
	return o.RepeatEvery, true
}

// HasRepeatEvery returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasRepeatEvery() bool {
	if o != nil && !IsNil(o.RepeatEvery) {
		return true
	}

	return false
}

// SetRepeatEvery gets a reference to the given string and assigns it to the RepeatEvery field.
func (o *TaskSchedulerSchema) SetRepeatEvery(v string) {
	o.RepeatEvery = &v
}

// GetStartFrom returns the StartFrom field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetStartFrom() int32 {
	if o == nil || IsNil(o.StartFrom) {
		var ret int32
		return ret
	}
	return *o.StartFrom
}

// GetStartFromOk returns a tuple with the StartFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetStartFromOk() (*int32, bool) {
	if o == nil || IsNil(o.StartFrom) {
		return nil, false
	}
	return o.StartFrom, true
}

// HasStartFrom returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasStartFrom() bool {
	if o != nil && !IsNil(o.StartFrom) {
		return true
	}

	return false
}

// SetStartFrom gets a reference to the given int32 and assigns it to the StartFrom field.
func (o *TaskSchedulerSchema) SetStartFrom(v int32) {
	o.StartFrom = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *TaskSchedulerSchema) SetJob(v string) {
	o.Job = &v
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetDateRange() string {
	if o == nil || IsNil(o.DateRange) {
		var ret string
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetDateRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given string and assigns it to the DateRange field.
func (o *TaskSchedulerSchema) SetDateRange(v string) {
	o.DateRange = &v
}

// GetDateKey returns the DateKey field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetDateKey() string {
	if o == nil || IsNil(o.DateKey) {
		var ret string
		return ret
	}
	return *o.DateKey
}

// GetDateKeyOk returns a tuple with the DateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetDateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DateKey) {
		return nil, false
	}
	return o.DateKey, true
}

// HasDateKey returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasDateKey() bool {
	if o != nil && !IsNil(o.DateKey) {
		return true
	}

	return false
}

// SetDateKey gets a reference to the given string and assigns it to the DateKey field.
func (o *TaskSchedulerSchema) SetDateKey(v string) {
	o.DateKey = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TaskSchedulerSchema) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *TaskSchedulerSchema) SetEndDate(v string) {
	o.EndDate = &v
}

// GetReportKeys returns the ReportKeys field value if set, zero value otherwise.
func (o *TaskSchedulerSchema) GetReportKeys() []string {
	if o == nil || IsNil(o.ReportKeys) {
		var ret []string
		return ret
	}
	return o.ReportKeys
}

// GetReportKeysOk returns a tuple with the ReportKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedulerSchema) GetReportKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ReportKeys) {
		return nil, false
	}
	return o.ReportKeys, true
}

// HasReportKeys returns a boolean if a field has been set.
func (o *TaskSchedulerSchema) HasReportKeys() bool {
	if o != nil && !IsNil(o.ReportKeys) {
		return true
	}

	return false
}

// SetReportKeys gets a reference to the given []string and assigns it to the ReportKeys field.
func (o *TaskSchedulerSchema) SetReportKeys(v []string) {
	o.ReportKeys = v
}

func (o TaskSchedulerSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskSchedulerSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.RepeatEvery) {
		toSerialize["repeat_every"] = o.RepeatEvery
	}
	if !IsNil(o.StartFrom) {
		toSerialize["start_from"] = o.StartFrom
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
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

type NullableTaskSchedulerSchema struct {
	value *TaskSchedulerSchema
	isSet bool
}

func (v NullableTaskSchedulerSchema) Get() *TaskSchedulerSchema {
	return v.value
}

func (v *NullableTaskSchedulerSchema) Set(val *TaskSchedulerSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSchedulerSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSchedulerSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSchedulerSchema(val *TaskSchedulerSchema) *NullableTaskSchedulerSchema {
	return &NullableTaskSchedulerSchema{value: val, isSet: true}
}

func (v NullableTaskSchedulerSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSchedulerSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
