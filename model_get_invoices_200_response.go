package openapi

import (
	"encoding/json"
)

// checks if the GetInvoices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoices200Response{}

// GetInvoices200Response struct for GetInvoices200Response
type GetInvoices200Response struct {
	Data []Invoice `json:"data,omitempty"`
	Meta Meta      `json:"meta,omitempty"`
}

// NewGetInvoices200Response instantiates a new GetInvoices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoices200Response() *GetInvoices200Response {
	this := GetInvoices200Response{}
	return &this
}

// NewGetInvoices200ResponseWithDefaults instantiates a new GetInvoices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoices200ResponseWithDefaults() *GetInvoices200Response {
	this := GetInvoices200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetInvoices200Response) GetData() []Invoice {
	if o == nil || IsNil(o.Data) {
		var ret []Invoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoices200Response) GetDataOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetInvoices200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invoice and assigns it to the Data field.
func (o *GetInvoices200Response) SetData(v []Invoice) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetInvoices200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoices200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetInvoices200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetInvoices200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetInvoices200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvoices200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetInvoices200Response struct {
	value *GetInvoices200Response
	isSet bool
}

func (v NullableGetInvoices200Response) Get() *GetInvoices200Response {
	return v.value
}

func (v *NullableGetInvoices200Response) Set(val *GetInvoices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoices200Response(val *GetInvoices200Response) *NullableGetInvoices200Response {
	return &NullableGetInvoices200Response{value: val, isSet: true}
}

func (v NullableGetInvoices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvoices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
