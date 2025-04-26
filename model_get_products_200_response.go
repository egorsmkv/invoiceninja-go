package openapi

import (
	"encoding/json"
)

// checks if the GetProducts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProducts200Response{}

// GetProducts200Response struct for GetProducts200Response
type GetProducts200Response struct {
	Meta Meta      `json:"meta,omitempty"`
	Data []Product `json:"data,omitempty"`
}

// NewGetProducts200Response instantiates a new GetProducts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProducts200Response() *GetProducts200Response {
	this := GetProducts200Response{}
	return &this
}

// NewGetProducts200ResponseWithDefaults instantiates a new GetProducts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProducts200ResponseWithDefaults() *GetProducts200Response {
	this := GetProducts200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProducts200Response) GetData() []Product {
	if o == nil || IsNil(o.Data) {
		var ret []Product
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProducts200Response) GetDataOk() ([]Product, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProducts200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Product and assigns it to the Data field.
func (o *GetProducts200Response) SetData(v []Product) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetProducts200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProducts200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetProducts200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetProducts200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetProducts200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProducts200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetProducts200Response struct {
	value *GetProducts200Response
	isSet bool
}

func (v NullableGetProducts200Response) Get() *GetProducts200Response {
	return v.value
}

func (v *NullableGetProducts200Response) Set(val *GetProducts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProducts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProducts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProducts200Response(val *GetProducts200Response) *NullableGetProducts200Response {
	return &NullableGetProducts200Response{value: val, isSet: true}
}

func (v NullableGetProducts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProducts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
