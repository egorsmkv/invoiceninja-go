package openapi

import (
	"encoding/json"
)

// checks if the GetPurchaseOrders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPurchaseOrders200Response{}

// GetPurchaseOrders200Response struct for GetPurchaseOrders200Response
type GetPurchaseOrders200Response struct {
	Meta Meta            `json:"meta,omitempty"`
	Data []PurchaseOrder `json:"data,omitempty"`
}

// NewGetPurchaseOrders200Response instantiates a new GetPurchaseOrders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPurchaseOrders200Response() *GetPurchaseOrders200Response {
	this := GetPurchaseOrders200Response{}
	return &this
}

// NewGetPurchaseOrders200ResponseWithDefaults instantiates a new GetPurchaseOrders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPurchaseOrders200ResponseWithDefaults() *GetPurchaseOrders200Response {
	this := GetPurchaseOrders200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPurchaseOrders200Response) GetData() []PurchaseOrder {
	if o == nil || IsNil(o.Data) {
		var ret []PurchaseOrder
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPurchaseOrders200Response) GetDataOk() ([]PurchaseOrder, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPurchaseOrders200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PurchaseOrder and assigns it to the Data field.
func (o *GetPurchaseOrders200Response) SetData(v []PurchaseOrder) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetPurchaseOrders200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPurchaseOrders200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetPurchaseOrders200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetPurchaseOrders200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetPurchaseOrders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPurchaseOrders200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetPurchaseOrders200Response struct {
	value *GetPurchaseOrders200Response
	isSet bool
}

func (v NullableGetPurchaseOrders200Response) Get() *GetPurchaseOrders200Response {
	return v.value
}

func (v *NullableGetPurchaseOrders200Response) Set(val *GetPurchaseOrders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPurchaseOrders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPurchaseOrders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPurchaseOrders200Response(val *GetPurchaseOrders200Response) *NullableGetPurchaseOrders200Response {
	return &NullableGetPurchaseOrders200Response{value: val, isSet: true}
}

func (v NullableGetPurchaseOrders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPurchaseOrders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
