package openapi

import (
	"encoding/json"
)

// checks if the GetRecurringInvoices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecurringInvoices200Response{}

// GetRecurringInvoices200Response struct for GetRecurringInvoices200Response
type GetRecurringInvoices200Response struct {
	Data []RecurringInvoice `json:"data,omitempty"`
	Meta Meta               `json:"meta,omitempty"`
}

// NewGetRecurringInvoices200Response instantiates a new GetRecurringInvoices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecurringInvoices200Response() *GetRecurringInvoices200Response {
	this := GetRecurringInvoices200Response{}
	return &this
}

// NewGetRecurringInvoices200ResponseWithDefaults instantiates a new GetRecurringInvoices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecurringInvoices200ResponseWithDefaults() *GetRecurringInvoices200Response {
	this := GetRecurringInvoices200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRecurringInvoices200Response) GetData() []RecurringInvoice {
	if o == nil || IsNil(o.Data) {
		var ret []RecurringInvoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringInvoices200Response) GetDataOk() ([]RecurringInvoice, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRecurringInvoices200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RecurringInvoice and assigns it to the Data field.
func (o *GetRecurringInvoices200Response) SetData(v []RecurringInvoice) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetRecurringInvoices200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringInvoices200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetRecurringInvoices200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetRecurringInvoices200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetRecurringInvoices200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecurringInvoices200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetRecurringInvoices200Response struct {
	value *GetRecurringInvoices200Response
	isSet bool
}

func (v NullableGetRecurringInvoices200Response) Get() *GetRecurringInvoices200Response {
	return v.value
}

func (v *NullableGetRecurringInvoices200Response) Set(val *GetRecurringInvoices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecurringInvoices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecurringInvoices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecurringInvoices200Response(val *GetRecurringInvoices200Response) *NullableGetRecurringInvoices200Response {
	return &NullableGetRecurringInvoices200Response{value: val, isSet: true}
}

func (v NullableGetRecurringInvoices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecurringInvoices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
