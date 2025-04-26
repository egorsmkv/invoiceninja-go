package openapi

import (
	"encoding/json"
)

// checks if the GetRecurringExpenses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecurringExpenses200Response{}

// GetRecurringExpenses200Response struct for GetRecurringExpenses200Response
type GetRecurringExpenses200Response struct {
	Meta Meta               `json:"meta,omitempty"`
	Data []RecurringExpense `json:"data,omitempty"`
}

// NewGetRecurringExpenses200Response instantiates a new GetRecurringExpenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecurringExpenses200Response() *GetRecurringExpenses200Response {
	this := GetRecurringExpenses200Response{}
	return &this
}

// NewGetRecurringExpenses200ResponseWithDefaults instantiates a new GetRecurringExpenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecurringExpenses200ResponseWithDefaults() *GetRecurringExpenses200Response {
	this := GetRecurringExpenses200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRecurringExpenses200Response) GetData() []RecurringExpense {
	if o == nil || IsNil(o.Data) {
		var ret []RecurringExpense
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringExpenses200Response) GetDataOk() ([]RecurringExpense, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRecurringExpenses200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RecurringExpense and assigns it to the Data field.
func (o *GetRecurringExpenses200Response) SetData(v []RecurringExpense) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetRecurringExpenses200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecurringExpenses200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetRecurringExpenses200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetRecurringExpenses200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetRecurringExpenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecurringExpenses200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetRecurringExpenses200Response struct {
	value *GetRecurringExpenses200Response
	isSet bool
}

func (v NullableGetRecurringExpenses200Response) Get() *GetRecurringExpenses200Response {
	return v.value
}

func (v *NullableGetRecurringExpenses200Response) Set(val *GetRecurringExpenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecurringExpenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecurringExpenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecurringExpenses200Response(val *GetRecurringExpenses200Response) *NullableGetRecurringExpenses200Response {
	return &NullableGetRecurringExpenses200Response{value: val, isSet: true}
}

func (v NullableGetRecurringExpenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecurringExpenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
