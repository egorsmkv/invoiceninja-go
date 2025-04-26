package openapi

import (
	"encoding/json"
)

// checks if the Design type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Design{}

// Design struct for Design
type Design struct {
	// The design hashed id
	Id *string `json:"id,omitempty"`
	// The design name
	Name *string `json:"name,omitempty"`
	// The design HTML
	Design *string `json:"design,omitempty"`
	// Flag to determine if the design is a custom user design
	IsCustom *bool `json:"is_custom,omitempty"`
	// Flag to determine if the design is available for use
	IsActive *bool `json:"is_active,omitempty"`
	// Flag to determine if the design is deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	DeletedAt *float32 `json:"deleted_at,omitempty"`
}

// NewDesign instantiates a new Design object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesign() *Design {
	this := Design{}
	return &this
}

// NewDesignWithDefaults instantiates a new Design object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesignWithDefaults() *Design {
	this := Design{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Design) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Design) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Design) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Design) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Design) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Design) SetName(v string) {
	o.Name = &v
}

// GetDesign returns the Design field value if set, zero value otherwise.
func (o *Design) GetDesign() string {
	if o == nil || IsNil(o.Design) {
		var ret string
		return ret
	}
	return *o.Design
}

// GetDesignOk returns a tuple with the Design field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetDesignOk() (*string, bool) {
	if o == nil || IsNil(o.Design) {
		return nil, false
	}
	return o.Design, true
}

// HasDesign returns a boolean if a field has been set.
func (o *Design) HasDesign() bool {
	if o != nil && !IsNil(o.Design) {
		return true
	}

	return false
}

// SetDesign gets a reference to the given string and assigns it to the Design field.
func (o *Design) SetDesign(v string) {
	o.Design = &v
}

// GetIsCustom returns the IsCustom field value if set, zero value otherwise.
func (o *Design) GetIsCustom() bool {
	if o == nil || IsNil(o.IsCustom) {
		var ret bool
		return ret
	}
	return *o.IsCustom
}

// GetIsCustomOk returns a tuple with the IsCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetIsCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustom) {
		return nil, false
	}
	return o.IsCustom, true
}

// HasIsCustom returns a boolean if a field has been set.
func (o *Design) HasIsCustom() bool {
	if o != nil && !IsNil(o.IsCustom) {
		return true
	}

	return false
}

// SetIsCustom gets a reference to the given bool and assigns it to the IsCustom field.
func (o *Design) SetIsCustom(v bool) {
	o.IsCustom = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Design) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Design) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Design) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Design) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Design) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Design) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Design) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Design) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Design) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Design) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Design) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Design) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Design) GetDeletedAt() float32 {
	if o == nil || IsNil(o.DeletedAt) {
		var ret float32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Design) GetDeletedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Design) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given float32 and assigns it to the DeletedAt field.
func (o *Design) SetDeletedAt(v float32) {
	o.DeletedAt = &v
}

func (o Design) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Design) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Design) {
		toSerialize["design"] = o.Design
	}
	if !IsNil(o.IsCustom) {
		toSerialize["is_custom"] = o.IsCustom
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return toSerialize, nil
}

type NullableDesign struct {
	value *Design
	isSet bool
}

func (v NullableDesign) Get() *Design {
	return v.value
}

func (v *NullableDesign) Set(val *Design) {
	v.value = val
	v.isSet = true
}

func (v NullableDesign) IsSet() bool {
	return v.isSet
}

func (v *NullableDesign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesign(val *Design) *NullableDesign {
	return &NullableDesign{value: val, isSet: true}
}

func (v NullableDesign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
