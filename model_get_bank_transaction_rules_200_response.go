package openapi

import (
	"encoding/json"
)

// checks if the GetBankTransactionRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBankTransactionRules200Response{}

// GetBankTransactionRules200Response struct for GetBankTransactionRules200Response
type GetBankTransactionRules200Response struct {
	Data []BankTransactionRule `json:"data,omitempty"`
	Meta Meta                  `json:"meta,omitempty"`
}

// NewGetBankTransactionRules200Response instantiates a new GetBankTransactionRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBankTransactionRules200Response() *GetBankTransactionRules200Response {
	this := GetBankTransactionRules200Response{}
	return &this
}

// NewGetBankTransactionRules200ResponseWithDefaults instantiates a new GetBankTransactionRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBankTransactionRules200ResponseWithDefaults() *GetBankTransactionRules200Response {
	this := GetBankTransactionRules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetBankTransactionRules200Response) GetData() []BankTransactionRule {
	if o == nil || IsNil(o.Data) {
		var ret []BankTransactionRule
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankTransactionRules200Response) GetDataOk() ([]BankTransactionRule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetBankTransactionRules200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BankTransactionRule and assigns it to the Data field.
func (o *GetBankTransactionRules200Response) SetData(v []BankTransactionRule) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetBankTransactionRules200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankTransactionRules200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetBankTransactionRules200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetBankTransactionRules200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetBankTransactionRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBankTransactionRules200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetBankTransactionRules200Response struct {
	value *GetBankTransactionRules200Response
	isSet bool
}

func (v NullableGetBankTransactionRules200Response) Get() *GetBankTransactionRules200Response {
	return v.value
}

func (v *NullableGetBankTransactionRules200Response) Set(val *GetBankTransactionRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBankTransactionRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBankTransactionRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBankTransactionRules200Response(val *GetBankTransactionRules200Response) *NullableGetBankTransactionRules200Response {
	return &NullableGetBankTransactionRules200Response{value: val, isSet: true}
}

func (v NullableGetBankTransactionRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBankTransactionRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
