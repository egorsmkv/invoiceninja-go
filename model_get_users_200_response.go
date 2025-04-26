package openapi

import (
	"encoding/json"
)

// checks if the GetUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsers200Response{}

// GetUsers200Response struct for GetUsers200Response
type GetUsers200Response struct {
	Data []User `json:"data,omitempty"`
	Meta Meta   `json:"meta,omitempty"`
}

// NewGetUsers200Response instantiates a new GetUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsers200Response() *GetUsers200Response {
	this := GetUsers200Response{}
	return &this
}

// NewGetUsers200ResponseWithDefaults instantiates a new GetUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsers200ResponseWithDefaults() *GetUsers200Response {
	this := GetUsers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUsers200Response) GetData() []User {
	if o == nil || IsNil(o.Data) {
		var ret []User
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200Response) GetDataOk() ([]User, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUsers200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []User and assigns it to the Data field.
func (o *GetUsers200Response) SetData(v []User) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetUsers200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetUsers200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetUsers200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetUsers200Response struct {
	value *GetUsers200Response
	isSet bool
}

func (v NullableGetUsers200Response) Get() *GetUsers200Response {
	return v.value
}

func (v *NullableGetUsers200Response) Set(val *GetUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsers200Response(val *GetUsers200Response) *NullableGetUsers200Response {
	return &NullableGetUsers200Response{value: val, isSet: true}
}

func (v NullableGetUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
