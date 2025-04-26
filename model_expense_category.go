package openapi

import (
	"encoding/json"
)

// checks if the ExpenseCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpenseCategory{}

// ExpenseCategory struct for ExpenseCategory
type ExpenseCategory struct {
	// The expense hashed id
	Id *string `json:"id,omitempty"`
	// The expense category name
	Name *string `json:"name,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// Flag determining whether the expense category has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The updated at timestamp
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// The created at timestamp
	CreatedAt *int32 `json:"created_at,omitempty"`
}

// NewExpenseCategory instantiates a new ExpenseCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseCategory() *ExpenseCategory {
	this := ExpenseCategory{}
	return &this
}

// NewExpenseCategoryWithDefaults instantiates a new ExpenseCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseCategoryWithDefaults() *ExpenseCategory {
	this := ExpenseCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExpenseCategory) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExpenseCategory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExpenseCategory) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExpenseCategory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExpenseCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExpenseCategory) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ExpenseCategory) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ExpenseCategory) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ExpenseCategory) SetUserId(v string) {
	o.UserId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *ExpenseCategory) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ExpenseCategory) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *ExpenseCategory) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExpenseCategory) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExpenseCategory) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *ExpenseCategory) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExpenseCategory) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseCategory) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExpenseCategory) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *ExpenseCategory) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

func (o ExpenseCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpenseCategory) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableExpenseCategory struct {
	value *ExpenseCategory
	isSet bool
}

func (v NullableExpenseCategory) Get() *ExpenseCategory {
	return v.value
}

func (v *NullableExpenseCategory) Set(val *ExpenseCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseCategory(val *ExpenseCategory) *NullableExpenseCategory {
	return &NullableExpenseCategory{value: val, isSet: true}
}

func (v NullableExpenseCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
