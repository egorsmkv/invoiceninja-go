package openapi

import (
	"encoding/json"
)

// checks if the GroupSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSetting{}

// GroupSetting struct for GroupSetting
type GroupSetting struct {
	// The group setting hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The name of the group
	Name *string `json:"name,omitempty"`
	// The settings object
	Settings map[string]any `json:"settings,omitempty"`
}

// NewGroupSetting instantiates a new GroupSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSetting() *GroupSetting {
	this := GroupSetting{}
	return &this
}

// NewGroupSettingWithDefaults instantiates a new GroupSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSettingWithDefaults() *GroupSetting {
	this := GroupSetting{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupSetting) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSetting) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupSetting) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GroupSetting) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSetting) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GroupSetting) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GroupSetting) SetUserId(v string) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupSetting) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSetting) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupSetting) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupSetting) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *GroupSetting) GetSettings() map[string]any {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]any
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSetting) GetSettingsOk() (map[string]any, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]any{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *GroupSetting) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]any and assigns it to the Settings field.
func (o *GroupSetting) SetSettings(v map[string]any) {
	o.Settings = v
}

func (o GroupSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSetting) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableGroupSetting struct {
	value *GroupSetting
	isSet bool
}

func (v NullableGroupSetting) Get() *GroupSetting {
	return v.value
}

func (v *NullableGroupSetting) Set(val *GroupSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSetting(val *GroupSetting) *NullableGroupSetting {
	return &NullableGroupSetting{value: val, isSet: true}
}

func (v NullableGroupSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
