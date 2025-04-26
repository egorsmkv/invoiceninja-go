package openapi

import (
	"encoding/json"
)

// checks if the SystemLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemLog{}

// SystemLog struct for SystemLog
type SystemLog struct {
	// The account hashed id
	Id *string `json:"id,omitempty"`
	// The user_id hashed id
	UserId *string `json:"user_id,omitempty"`
	// The client_id hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The Log Type ID
	EventId *int32 `json:"event_id,omitempty"`
	// The Category Type ID
	CategoryId *int32 `json:"category_id,omitempty"`
	// The Type Type ID
	TypeId *int32 `json:"type_id,omitempty"`
	// The json object of the error
	Log map[string]any `json:"log,omitempty"`
	// Timestamp
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Timestamp
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewSystemLog instantiates a new SystemLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemLog() *SystemLog {
	this := SystemLog{}
	return &this
}

// NewSystemLogWithDefaults instantiates a new SystemLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemLogWithDefaults() *SystemLog {
	this := SystemLog{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SystemLog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SystemLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SystemLog) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SystemLog) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SystemLog) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SystemLog) SetUserId(v string) {
	o.UserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SystemLog) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SystemLog) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SystemLog) SetClientId(v string) {
	o.ClientId = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *SystemLog) GetEventId() int32 {
	if o == nil || IsNil(o.EventId) {
		var ret int32
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetEventIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *SystemLog) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int32 and assigns it to the EventId field.
func (o *SystemLog) SetEventId(v int32) {
	o.EventId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *SystemLog) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *SystemLog) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *SystemLog) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *SystemLog) GetTypeId() int32 {
	if o == nil || IsNil(o.TypeId) {
		var ret int32
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *SystemLog) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.
func (o *SystemLog) SetTypeId(v int32) {
	o.TypeId = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *SystemLog) GetLog() map[string]any {
	if o == nil || IsNil(o.Log) {
		var ret map[string]any
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetLogOk() (map[string]any, bool) {
	if o == nil || IsNil(o.Log) {
		return map[string]any{}, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *SystemLog) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given map[string]interface{} and assigns it to the Log field.
func (o *SystemLog) SetLog(v map[string]any) {
	o.Log = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SystemLog) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SystemLog) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SystemLog) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SystemLog) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SystemLog) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SystemLog) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o SystemLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemLog) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.TypeId) {
		toSerialize["type_id"] = o.TypeId
	}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableSystemLog struct {
	value *SystemLog
	isSet bool
}

func (v NullableSystemLog) Get() *SystemLog {
	return v.value
}

func (v *NullableSystemLog) Set(val *SystemLog) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemLog) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemLog(val *SystemLog) *NullableSystemLog {
	return &NullableSystemLog{value: val, isSet: true}
}

func (v NullableSystemLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
