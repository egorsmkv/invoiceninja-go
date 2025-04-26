package openapi

import (
	"encoding/json"
)

// checks if the BankTransactionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransactionRule{}

// BankTransactionRule struct for BankTransactionRule
type BankTransactionRule struct {
	// The bank transaction rules hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The name of the transaction
	Name *string `json:"name,omitempty"`
	// A mapped collection of the sub rules for the BankTransactionRule
	Rules []BTRules `json:"rules,omitempty"`
	// Flags whether the rule converts the transaction automatically
	AutoConvert *bool `json:"auto_convert,omitempty"`
	// Flags whether all subrules are required for the match
	MatchesOnAll *bool `json:"matches_on_all,omitempty"`
	// Flags whether the rule applies to a CREDIT or DEBIT
	AppliesTo *string `json:"applies_to,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The vendor hashed id
	VendorId *string `json:"vendor_id,omitempty"`
	// The category hashed id
	CategoryId *string `json:"category_id,omitempty"`
}

// NewBankTransactionRule instantiates a new BankTransactionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransactionRule() *BankTransactionRule {
	this := BankTransactionRule{}
	return &this
}

// NewBankTransactionRuleWithDefaults instantiates a new BankTransactionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransactionRuleWithDefaults() *BankTransactionRule {
	this := BankTransactionRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankTransactionRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankTransactionRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankTransactionRule) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BankTransactionRule) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BankTransactionRule) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BankTransactionRule) SetUserId(v string) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BankTransactionRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BankTransactionRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BankTransactionRule) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *BankTransactionRule) GetRules() []BTRules {
	if o == nil || IsNil(o.Rules) {
		var ret []BTRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetRulesOk() ([]BTRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *BankTransactionRule) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []BTRules and assigns it to the Rules field.
func (o *BankTransactionRule) SetRules(v []BTRules) {
	o.Rules = v
}

// GetAutoConvert returns the AutoConvert field value if set, zero value otherwise.
func (o *BankTransactionRule) GetAutoConvert() bool {
	if o == nil || IsNil(o.AutoConvert) {
		var ret bool
		return ret
	}
	return *o.AutoConvert
}

// GetAutoConvertOk returns a tuple with the AutoConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetAutoConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConvert) {
		return nil, false
	}
	return o.AutoConvert, true
}

// HasAutoConvert returns a boolean if a field has been set.
func (o *BankTransactionRule) HasAutoConvert() bool {
	if o != nil && !IsNil(o.AutoConvert) {
		return true
	}

	return false
}

// SetAutoConvert gets a reference to the given bool and assigns it to the AutoConvert field.
func (o *BankTransactionRule) SetAutoConvert(v bool) {
	o.AutoConvert = &v
}

// GetMatchesOnAll returns the MatchesOnAll field value if set, zero value otherwise.
func (o *BankTransactionRule) GetMatchesOnAll() bool {
	if o == nil || IsNil(o.MatchesOnAll) {
		var ret bool
		return ret
	}
	return *o.MatchesOnAll
}

// GetMatchesOnAllOk returns a tuple with the MatchesOnAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetMatchesOnAllOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchesOnAll) {
		return nil, false
	}
	return o.MatchesOnAll, true
}

// HasMatchesOnAll returns a boolean if a field has been set.
func (o *BankTransactionRule) HasMatchesOnAll() bool {
	if o != nil && !IsNil(o.MatchesOnAll) {
		return true
	}

	return false
}

// SetMatchesOnAll gets a reference to the given bool and assigns it to the MatchesOnAll field.
func (o *BankTransactionRule) SetMatchesOnAll(v bool) {
	o.MatchesOnAll = &v
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *BankTransactionRule) GetAppliesTo() string {
	if o == nil || IsNil(o.AppliesTo) {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetAppliesToOk() (*string, bool) {
	if o == nil || IsNil(o.AppliesTo) {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *BankTransactionRule) HasAppliesTo() bool {
	if o != nil && !IsNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *BankTransactionRule) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransactionRule) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransactionRule) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransactionRule) SetClientId(v string) {
	o.ClientId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *BankTransactionRule) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *BankTransactionRule) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *BankTransactionRule) SetVendorId(v string) {
	o.VendorId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *BankTransactionRule) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransactionRule) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *BankTransactionRule) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *BankTransactionRule) SetCategoryId(v string) {
	o.CategoryId = &v
}

func (o BankTransactionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransactionRule) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.AutoConvert) {
		toSerialize["auto_convert"] = o.AutoConvert
	}
	if !IsNil(o.MatchesOnAll) {
		toSerialize["matches_on_all"] = o.MatchesOnAll
	}
	if !IsNil(o.AppliesTo) {
		toSerialize["applies_to"] = o.AppliesTo
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	return toSerialize, nil
}

type NullableBankTransactionRule struct {
	value *BankTransactionRule
	isSet bool
}

func (v NullableBankTransactionRule) Get() *BankTransactionRule {
	return v.value
}

func (v *NullableBankTransactionRule) Set(val *BankTransactionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransactionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransactionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransactionRule(val *BankTransactionRule) *NullableBankTransactionRule {
	return &NullableBankTransactionRule{value: val, isSet: true}
}

func (v NullableBankTransactionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransactionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
