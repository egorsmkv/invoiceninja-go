package openapi

import (
	"encoding/json"
)

// checks if the GetBankTransactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBankTransactions200Response{}

// GetBankTransactions200Response struct for GetBankTransactions200Response
type GetBankTransactions200Response struct {
	Meta Meta              `json:"meta,omitempty"`
	Data []BankTransaction `json:"data,omitempty"`
}

// NewGetBankTransactions200Response instantiates a new GetBankTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBankTransactions200Response() *GetBankTransactions200Response {
	this := GetBankTransactions200Response{}
	return &this
}

// NewGetBankTransactions200ResponseWithDefaults instantiates a new GetBankTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBankTransactions200ResponseWithDefaults() *GetBankTransactions200Response {
	this := GetBankTransactions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetBankTransactions200Response) GetData() []BankTransaction {
	if o == nil || IsNil(o.Data) {
		var ret []BankTransaction
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankTransactions200Response) GetDataOk() ([]BankTransaction, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetBankTransactions200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BankTransaction and assigns it to the Data field.
func (o *GetBankTransactions200Response) SetData(v []BankTransaction) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetBankTransactions200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankTransactions200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetBankTransactions200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetBankTransactions200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetBankTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBankTransactions200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetBankTransactions200Response struct {
	value *GetBankTransactions200Response
	isSet bool
}

func (v NullableGetBankTransactions200Response) Get() *GetBankTransactions200Response {
	return v.value
}

func (v *NullableGetBankTransactions200Response) Set(val *GetBankTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBankTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBankTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBankTransactions200Response(val *GetBankTransactions200Response) *NullableGetBankTransactions200Response {
	return &NullableGetBankTransactions200Response{value: val, isSet: true}
}

func (v NullableGetBankTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBankTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
