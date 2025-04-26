package openapi

import (
	"encoding/json"
)

// checks if the GetActivities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActivities200Response{}

// GetActivities200Response struct for GetActivities200Response
type GetActivities200Response struct {
	Data []Activity `json:"data,omitempty"`
	Meta Meta       `json:"meta,omitempty"`
}

// NewGetActivities200Response instantiates a new GetActivities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActivities200Response() *GetActivities200Response {
	this := GetActivities200Response{}
	return &this
}

// NewGetActivities200ResponseWithDefaults instantiates a new GetActivities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActivities200ResponseWithDefaults() *GetActivities200Response {
	this := GetActivities200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetActivities200Response) GetData() []Activity {
	if o == nil || IsNil(o.Data) {
		var ret []Activity
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivities200Response) GetDataOk() ([]Activity, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetActivities200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Activity and assigns it to the Data field.
func (o *GetActivities200Response) SetData(v []Activity) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetActivities200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivities200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetActivities200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetActivities200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetActivities200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActivities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetActivities200Response struct {
	value *GetActivities200Response
	isSet bool
}

func (v NullableGetActivities200Response) Get() *GetActivities200Response {
	return v.value
}

func (v *NullableGetActivities200Response) Set(val *GetActivities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActivities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActivities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActivities200Response(val *GetActivities200Response) *NullableGetActivities200Response {
	return &NullableGetActivities200Response{value: val, isSet: true}
}

func (v NullableGetActivities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActivities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
