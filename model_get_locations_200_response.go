package openapi

import (
	"encoding/json"
)

// checks if the GetLocations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLocations200Response{}

// GetLocations200Response struct for GetLocations200Response
type GetLocations200Response struct {
	Data []Location `json:"data,omitempty"`
}

// NewGetLocations200Response instantiates a new GetLocations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLocations200Response() *GetLocations200Response {
	this := GetLocations200Response{}
	return &this
}

// NewGetLocations200ResponseWithDefaults instantiates a new GetLocations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLocations200ResponseWithDefaults() *GetLocations200Response {
	this := GetLocations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetLocations200Response) GetData() []Location {
	if o == nil || IsNil(o.Data) {
		var ret []Location
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLocations200Response) GetDataOk() ([]Location, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetLocations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Location and assigns it to the Data field.
func (o *GetLocations200Response) SetData(v []Location) {
	o.Data = v
}

func (o GetLocations200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLocations200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetLocations200Response struct {
	value *GetLocations200Response
	isSet bool
}

func (v NullableGetLocations200Response) Get() *GetLocations200Response {
	return v.value
}

func (v *NullableGetLocations200Response) Set(val *GetLocations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLocations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLocations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLocations200Response(val *GetLocations200Response) *NullableGetLocations200Response {
	return &NullableGetLocations200Response{value: val, isSet: true}
}

func (v NullableGetLocations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLocations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
