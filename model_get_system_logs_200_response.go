package openapi

import (
	"encoding/json"
)

// checks if the GetSystemLogs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSystemLogs200Response{}

// GetSystemLogs200Response struct for GetSystemLogs200Response
type GetSystemLogs200Response struct {
	Data []SystemLog `json:"data,omitempty"`
	Meta Meta        `json:"meta,omitempty"`
}

// NewGetSystemLogs200Response instantiates a new GetSystemLogs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSystemLogs200Response() *GetSystemLogs200Response {
	this := GetSystemLogs200Response{}
	return &this
}

// NewGetSystemLogs200ResponseWithDefaults instantiates a new GetSystemLogs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSystemLogs200ResponseWithDefaults() *GetSystemLogs200Response {
	this := GetSystemLogs200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSystemLogs200Response) GetData() []SystemLog {
	if o == nil || IsNil(o.Data) {
		var ret []SystemLog
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSystemLogs200Response) GetDataOk() ([]SystemLog, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSystemLogs200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SystemLog and assigns it to the Data field.
func (o *GetSystemLogs200Response) SetData(v []SystemLog) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetSystemLogs200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSystemLogs200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetSystemLogs200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetSystemLogs200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetSystemLogs200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSystemLogs200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetSystemLogs200Response struct {
	value *GetSystemLogs200Response
	isSet bool
}

func (v NullableGetSystemLogs200Response) Get() *GetSystemLogs200Response {
	return v.value
}

func (v *NullableGetSystemLogs200Response) Set(val *GetSystemLogs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSystemLogs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSystemLogs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSystemLogs200Response(val *GetSystemLogs200Response) *NullableGetSystemLogs200Response {
	return &NullableGetSystemLogs200Response{value: val, isSet: true}
}

func (v NullableGetSystemLogs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSystemLogs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
