package openapi

import (
	"encoding/json"
)

// checks if the GetRecurringQuotes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecurringQuotes200Response{}

// GetRecurringQuotes200Response struct for GetRecurringQuotes200Response
type GetRecurringQuotes200Response struct {
	Data []RecurringQuote `json:"data,omitempty"`
	Meta Meta             `json:"meta,omitempty"`
}

// NewGetRecurringQuotes200Response instantiates a new GetRecurringQuotes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecurringQuotes200Response() *GetRecurringQuotes200Response {
	this := GetRecurringQuotes200Response{}
	return &this
}

// NewGetRecurringQuotes200ResponseWithDefaults instantiates a new GetRecurringQuotes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecurringQuotes200ResponseWithDefaults() *GetRecurringQuotes200Response {
	this := GetRecurringQuotes200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRecurringQuotes200Response) GetData() []RecurringQuote {
	if o == nil || IsNil(o.Data) {
		var ret []RecurringQuote
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringQuotes200Response) GetDataOk() ([]RecurringQuote, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRecurringQuotes200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RecurringQuote and assigns it to the Data field.
func (o *GetRecurringQuotes200Response) SetData(v []RecurringQuote) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetRecurringQuotes200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringQuotes200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetRecurringQuotes200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetRecurringQuotes200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetRecurringQuotes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecurringQuotes200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetRecurringQuotes200Response struct {
	value *GetRecurringQuotes200Response
	isSet bool
}

func (v NullableGetRecurringQuotes200Response) Get() *GetRecurringQuotes200Response {
	return v.value
}

func (v *NullableGetRecurringQuotes200Response) Set(val *GetRecurringQuotes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecurringQuotes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecurringQuotes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecurringQuotes200Response(val *GetRecurringQuotes200Response) *NullableGetRecurringQuotes200Response {
	return &NullableGetRecurringQuotes200Response{value: val, isSet: true}
}

func (v NullableGetRecurringQuotes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecurringQuotes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
