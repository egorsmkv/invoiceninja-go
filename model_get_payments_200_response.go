package openapi

import (
	"encoding/json"
)

// checks if the GetPayments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPayments200Response{}

// GetPayments200Response struct for GetPayments200Response
type GetPayments200Response struct {
	Meta Meta      `json:"meta,omitempty"`
	Data []Payment `json:"data,omitempty"`
}

// NewGetPayments200Response instantiates a new GetPayments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayments200Response() *GetPayments200Response {
	this := GetPayments200Response{}
	return &this
}

// NewGetPayments200ResponseWithDefaults instantiates a new GetPayments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayments200ResponseWithDefaults() *GetPayments200Response {
	this := GetPayments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPayments200Response) GetData() []Payment {
	if o == nil || IsNil(o.Data) {
		var ret []Payment
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayments200Response) GetDataOk() ([]Payment, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPayments200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Payment and assigns it to the Data field.
func (o *GetPayments200Response) SetData(v []Payment) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetPayments200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayments200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetPayments200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetPayments200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetPayments200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayments200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetPayments200Response struct {
	value *GetPayments200Response
	isSet bool
}

func (v NullableGetPayments200Response) Get() *GetPayments200Response {
	return v.value
}

func (v *NullableGetPayments200Response) Set(val *GetPayments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayments200Response(val *GetPayments200Response) *NullableGetPayments200Response {
	return &NullableGetPayments200Response{value: val, isSet: true}
}

func (v NullableGetPayments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
