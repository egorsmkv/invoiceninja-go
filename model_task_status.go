package openapi

import (
	"encoding/json"
)

// checks if the TaskStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskStatus{}

// TaskStatus struct for TaskStatus
type TaskStatus struct {
	// The task status hashed id
	Id *string `json:"id,omitempty"`
	// The task status name
	Name *string `json:"name,omitempty"`
	// The task status color - hex value
	Color *string `json:"color,omitempty"`
	// The order of the task status
	TaskStatusOrder *int32 `json:"task_status_order,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// A boolean flag determining if the task status has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
}

// NewTaskStatus instantiates a new TaskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStatus() *TaskStatus {
	this := TaskStatus{}
	return &this
}

// NewTaskStatusWithDefaults instantiates a new TaskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStatusWithDefaults() *TaskStatus {
	this := TaskStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskStatus) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskStatus) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskStatus) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskStatus) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskStatus) SetName(v string) {
	o.Name = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TaskStatus) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TaskStatus) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TaskStatus) SetColor(v string) {
	o.Color = &v
}

// GetTaskStatusOrder returns the TaskStatusOrder field value if set, zero value otherwise.
func (o *TaskStatus) GetTaskStatusOrder() int32 {
	if o == nil || IsNil(o.TaskStatusOrder) {
		var ret int32
		return ret
	}
	return *o.TaskStatusOrder
}

// GetTaskStatusOrderOk returns a tuple with the TaskStatusOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetTaskStatusOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskStatusOrder) {
		return nil, false
	}
	return o.TaskStatusOrder, true
}

// HasTaskStatusOrder returns a boolean if a field has been set.
func (o *TaskStatus) HasTaskStatusOrder() bool {
	if o != nil && !IsNil(o.TaskStatusOrder) {
		return true
	}

	return false
}

// SetTaskStatusOrder gets a reference to the given int32 and assigns it to the TaskStatusOrder field.
func (o *TaskStatus) SetTaskStatusOrder(v int32) {
	o.TaskStatusOrder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TaskStatus) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TaskStatus) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *TaskStatus) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *TaskStatus) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *TaskStatus) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *TaskStatus) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TaskStatus) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TaskStatus) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *TaskStatus) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *TaskStatus) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStatus) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *TaskStatus) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *TaskStatus) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

func (o TaskStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.TaskStatusOrder) {
		toSerialize["task_status_order"] = o.TaskStatusOrder
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullableTaskStatus struct {
	value *TaskStatus
	isSet bool
}

func (v NullableTaskStatus) Get() *TaskStatus {
	return v.value
}

func (v *NullableTaskStatus) Set(val *TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatus(val *TaskStatus) *NullableTaskStatus {
	return &NullableTaskStatus{value: val, isSet: true}
}

func (v NullableTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
