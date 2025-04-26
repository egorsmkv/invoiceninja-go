package openapi

import (
	"encoding/json"
)

// checks if the GetDocuments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDocuments200Response{}

// GetDocuments200Response struct for GetDocuments200Response
type GetDocuments200Response struct {
	Data []Document `json:"data,omitempty"`
	Meta Meta       `json:"meta,omitempty"`
}

// NewGetDocuments200Response instantiates a new GetDocuments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocuments200Response() *GetDocuments200Response {
	this := GetDocuments200Response{}
	return &this
}

// NewGetDocuments200ResponseWithDefaults instantiates a new GetDocuments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocuments200ResponseWithDefaults() *GetDocuments200Response {
	this := GetDocuments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDocuments200Response) GetData() []Document {
	if o == nil || IsNil(o.Data) {
		var ret []Document
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocuments200Response) GetDataOk() ([]Document, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDocuments200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Document and assigns it to the Data field.
func (o *GetDocuments200Response) SetData(v []Document) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetDocuments200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocuments200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetDocuments200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetDocuments200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetDocuments200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDocuments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetDocuments200Response struct {
	value *GetDocuments200Response
	isSet bool
}

func (v NullableGetDocuments200Response) Get() *GetDocuments200Response {
	return v.value
}

func (v *NullableGetDocuments200Response) Set(val *GetDocuments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocuments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocuments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocuments200Response(val *GetDocuments200Response) *NullableGetDocuments200Response {
	return &NullableGetDocuments200Response{value: val, isSet: true}
}

func (v NullableGetDocuments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocuments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
