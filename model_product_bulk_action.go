package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProductBulkAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductBulkAction{}

// ProductBulkAction struct for ProductBulkAction
type ProductBulkAction struct {
	// The tax rate id to set on the list of products  The following constants are available (default = '1')  ``` PRODUCT_TYPE_PHYSICAL = '1' PRODUCT_TYPE_SERVICE = '2' PRODUCT_TYPE_DIGITAL = '3' PRODUCT_TYPE_SHIPPING = '4' PRODUCT_TYPE_EXEMPT = '5' PRODUCT_TYPE_REDUCED_TAX = '6' PRODUCT_TYPE_OVERRIDE_TAX = '7' PRODUCT_TYPE_ZERO_RATED = '8' PRODUCT_TYPE_REVERSE_TAX = '9' ```
	TaxId *string `json:"tax_id,omitempty"`
	// The action to perform ie. archive / restore / delete / set_tax_id
	Action string   `json:"action"`
	Ids    []string `json:"ids"`
}

type _ProductBulkAction ProductBulkAction

// NewProductBulkAction instantiates a new ProductBulkAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBulkAction(action string, ids []string) *ProductBulkAction {
	this := ProductBulkAction{}
	this.Action = action
	this.Ids = ids
	return &this
}

// NewProductBulkActionWithDefaults instantiates a new ProductBulkAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBulkActionWithDefaults() *ProductBulkAction {
	this := ProductBulkAction{}
	return &this
}

// GetAction returns the Action field value
func (o *ProductBulkAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ProductBulkAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ProductBulkAction) SetAction(v string) {
	o.Action = v
}

// GetIds returns the Ids field value
func (o *ProductBulkAction) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ProductBulkAction) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ProductBulkAction) SetIds(v []string) {
	o.Ids = v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *ProductBulkAction) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBulkAction) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *ProductBulkAction) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *ProductBulkAction) SetTaxId(v string) {
	o.TaxId = &v
}

func (o ProductBulkAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductBulkAction) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	toSerialize["action"] = o.Action
	toSerialize["ids"] = o.Ids
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	return toSerialize, nil
}

func (o *ProductBulkAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"ids",
	}

	allProperties := make(map[string]any)

	err = json.Unmarshal(data, &allProperties)
	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProductBulkAction := _ProductBulkAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductBulkAction)
	if err != nil {
		return err
	}

	*o = ProductBulkAction(varProductBulkAction)

	return err
}

type NullableProductBulkAction struct {
	value *ProductBulkAction
	isSet bool
}

func (v NullableProductBulkAction) Get() *ProductBulkAction {
	return v.value
}

func (v *NullableProductBulkAction) Set(val *ProductBulkAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBulkAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBulkAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBulkAction(val *ProductBulkAction) *NullableProductBulkAction {
	return &NullableProductBulkAction{value: val, isSet: true}
}

func (v NullableProductBulkAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBulkAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
