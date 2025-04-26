package openapi

import (
	"encoding/json"
)

// checks if the GetPaymentTerms200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentTerms200Response{}

// GetPaymentTerms200Response struct for GetPaymentTerms200Response
type GetPaymentTerms200Response struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []PaymentTerm `json:"data,omitempty"`
}

// NewGetPaymentTerms200Response instantiates a new GetPaymentTerms200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentTerms200Response() *GetPaymentTerms200Response {
	this := GetPaymentTerms200Response{}
	return &this
}

// NewGetPaymentTerms200ResponseWithDefaults instantiates a new GetPaymentTerms200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentTerms200ResponseWithDefaults() *GetPaymentTerms200Response {
	this := GetPaymentTerms200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPaymentTerms200Response) GetData() []PaymentTerm {
	if o == nil || IsNil(o.Data) {
		var ret []PaymentTerm
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentTerms200Response) GetDataOk() ([]PaymentTerm, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPaymentTerms200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PaymentTerm and assigns it to the Data field.
func (o *GetPaymentTerms200Response) SetData(v []PaymentTerm) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetPaymentTerms200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentTerms200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetPaymentTerms200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetPaymentTerms200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetPaymentTerms200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaymentTerms200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetPaymentTerms200Response struct {
	value *GetPaymentTerms200Response
	isSet bool
}

func (v NullableGetPaymentTerms200Response) Get() *GetPaymentTerms200Response {
	return v.value
}

func (v *NullableGetPaymentTerms200Response) Set(val *GetPaymentTerms200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentTerms200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentTerms200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentTerms200Response(val *GetPaymentTerms200Response) *NullableGetPaymentTerms200Response {
	return &NullableGetPaymentTerms200Response{value: val, isSet: true}
}

func (v NullableGetPaymentTerms200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentTerms200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
