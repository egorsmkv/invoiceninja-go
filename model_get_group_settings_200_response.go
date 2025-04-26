package openapi

import (
	"encoding/json"
)

// checks if the GetGroupSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGroupSettings200Response{}

// GetGroupSettings200Response struct for GetGroupSettings200Response
type GetGroupSettings200Response struct {
	Data []GroupSetting `json:"data,omitempty"`
	Meta Meta           `json:"meta,omitempty"`
}

// NewGetGroupSettings200Response instantiates a new GetGroupSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGroupSettings200Response() *GetGroupSettings200Response {
	this := GetGroupSettings200Response{}
	return &this
}

// NewGetGroupSettings200ResponseWithDefaults instantiates a new GetGroupSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGroupSettings200ResponseWithDefaults() *GetGroupSettings200Response {
	this := GetGroupSettings200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetGroupSettings200Response) GetData() []GroupSetting {
	if o == nil || IsNil(o.Data) {
		var ret []GroupSetting
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupSettings200Response) GetDataOk() ([]GroupSetting, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetGroupSettings200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GroupSetting and assigns it to the Data field.
func (o *GetGroupSettings200Response) SetData(v []GroupSetting) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetGroupSettings200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupSettings200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetGroupSettings200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetGroupSettings200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetGroupSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGroupSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetGroupSettings200Response struct {
	value *GetGroupSettings200Response
	isSet bool
}

func (v NullableGetGroupSettings200Response) Get() *GetGroupSettings200Response {
	return v.value
}

func (v *NullableGetGroupSettings200Response) Set(val *GetGroupSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGroupSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGroupSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGroupSettings200Response(val *GetGroupSettings200Response) *NullableGetGroupSettings200Response {
	return &NullableGetGroupSettings200Response{value: val, isSet: true}
}

func (v NullableGetGroupSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGroupSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
