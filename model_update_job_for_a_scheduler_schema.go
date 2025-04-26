package openapi

import (
	"encoding/json"
)

// checks if the UpdateJobForASchedulerSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateJobForASchedulerSchema{}

// UpdateJobForASchedulerSchema struct for UpdateJobForASchedulerSchema
type UpdateJobForASchedulerSchema struct {
	// Set action name, action names can be found in Scheduler Model
	Job *string `json:"job,omitempty"`
}

// NewUpdateJobForASchedulerSchema instantiates a new UpdateJobForASchedulerSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateJobForASchedulerSchema() *UpdateJobForASchedulerSchema {
	this := UpdateJobForASchedulerSchema{}
	return &this
}

// NewUpdateJobForASchedulerSchemaWithDefaults instantiates a new UpdateJobForASchedulerSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateJobForASchedulerSchemaWithDefaults() *UpdateJobForASchedulerSchema {
	this := UpdateJobForASchedulerSchema{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateJobForASchedulerSchema) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateJobForASchedulerSchema) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateJobForASchedulerSchema) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *UpdateJobForASchedulerSchema) SetJob(v string) {
	o.Job = &v
}

func (o UpdateJobForASchedulerSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateJobForASchedulerSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateJobForASchedulerSchema struct {
	value *UpdateJobForASchedulerSchema
	isSet bool
}

func (v NullableUpdateJobForASchedulerSchema) Get() *UpdateJobForASchedulerSchema {
	return v.value
}

func (v *NullableUpdateJobForASchedulerSchema) Set(val *UpdateJobForASchedulerSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateJobForASchedulerSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateJobForASchedulerSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateJobForASchedulerSchema(val *UpdateJobForASchedulerSchema) *NullableUpdateJobForASchedulerSchema {
	return &NullableUpdateJobForASchedulerSchema{value: val, isSet: true}
}

func (v NullableUpdateJobForASchedulerSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateJobForASchedulerSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
