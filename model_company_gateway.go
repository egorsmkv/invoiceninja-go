package openapi

import (
	"encoding/json"
)

// checks if the CompanyGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyGateway{}

// CompanyGateway struct for CompanyGateway
type CompanyGateway struct {
	// The hashed id of the company gateway
	Id *string `json:"id,omitempty"`
	// The gateway key (hash)
	GatewayKey *string `json:"gateway_key,omitempty"`
	// Bitmask representation of cards
	AcceptedCreditCards *int32 `json:"accepted_credit_cards,omitempty"`
	// Determines if the the billing address is required prior to payment.
	RequireBillingAddress *bool `json:"require_billing_address,omitempty"`
	// Determines if the the billing address is required prior to payment.
	RequireShippingAddress *bool `json:"require_shipping_address,omitempty"`
	// The configuration map for the gateway
	Config *string `json:"config,omitempty"`
	// Determines if the client details should be updated.
	UpdateDetails *bool `json:"update_details,omitempty"`
	// A mapped collection of the fees and limits for the configured gateway
	FeesAndLimits []FeesAndLimits `json:"fees_and_limits,omitempty"`
}

// NewCompanyGateway instantiates a new CompanyGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyGateway() *CompanyGateway {
	this := CompanyGateway{}
	return &this
}

// NewCompanyGatewayWithDefaults instantiates a new CompanyGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyGatewayWithDefaults() *CompanyGateway {
	this := CompanyGateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyGateway) SetId(v string) {
	o.Id = &v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *CompanyGateway) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *CompanyGateway) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *CompanyGateway) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetAcceptedCreditCards returns the AcceptedCreditCards field value if set, zero value otherwise.
func (o *CompanyGateway) GetAcceptedCreditCards() int32 {
	if o == nil || IsNil(o.AcceptedCreditCards) {
		var ret int32
		return ret
	}
	return *o.AcceptedCreditCards
}

// GetAcceptedCreditCardsOk returns a tuple with the AcceptedCreditCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetAcceptedCreditCardsOk() (*int32, bool) {
	if o == nil || IsNil(o.AcceptedCreditCards) {
		return nil, false
	}
	return o.AcceptedCreditCards, true
}

// HasAcceptedCreditCards returns a boolean if a field has been set.
func (o *CompanyGateway) HasAcceptedCreditCards() bool {
	if o != nil && !IsNil(o.AcceptedCreditCards) {
		return true
	}

	return false
}

// SetAcceptedCreditCards gets a reference to the given int32 and assigns it to the AcceptedCreditCards field.
func (o *CompanyGateway) SetAcceptedCreditCards(v int32) {
	o.AcceptedCreditCards = &v
}

// GetRequireBillingAddress returns the RequireBillingAddress field value if set, zero value otherwise.
func (o *CompanyGateway) GetRequireBillingAddress() bool {
	if o == nil || IsNil(o.RequireBillingAddress) {
		var ret bool
		return ret
	}
	return *o.RequireBillingAddress
}

// GetRequireBillingAddressOk returns a tuple with the RequireBillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetRequireBillingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBillingAddress) {
		return nil, false
	}
	return o.RequireBillingAddress, true
}

// HasRequireBillingAddress returns a boolean if a field has been set.
func (o *CompanyGateway) HasRequireBillingAddress() bool {
	if o != nil && !IsNil(o.RequireBillingAddress) {
		return true
	}

	return false
}

// SetRequireBillingAddress gets a reference to the given bool and assigns it to the RequireBillingAddress field.
func (o *CompanyGateway) SetRequireBillingAddress(v bool) {
	o.RequireBillingAddress = &v
}

// GetRequireShippingAddress returns the RequireShippingAddress field value if set, zero value otherwise.
func (o *CompanyGateway) GetRequireShippingAddress() bool {
	if o == nil || IsNil(o.RequireShippingAddress) {
		var ret bool
		return ret
	}
	return *o.RequireShippingAddress
}

// GetRequireShippingAddressOk returns a tuple with the RequireShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetRequireShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireShippingAddress) {
		return nil, false
	}
	return o.RequireShippingAddress, true
}

// HasRequireShippingAddress returns a boolean if a field has been set.
func (o *CompanyGateway) HasRequireShippingAddress() bool {
	if o != nil && !IsNil(o.RequireShippingAddress) {
		return true
	}

	return false
}

// SetRequireShippingAddress gets a reference to the given bool and assigns it to the RequireShippingAddress field.
func (o *CompanyGateway) SetRequireShippingAddress(v bool) {
	o.RequireShippingAddress = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CompanyGateway) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CompanyGateway) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *CompanyGateway) SetConfig(v string) {
	o.Config = &v
}

// GetUpdateDetails returns the UpdateDetails field value if set, zero value otherwise.
func (o *CompanyGateway) GetUpdateDetails() bool {
	if o == nil || IsNil(o.UpdateDetails) {
		var ret bool
		return ret
	}
	return *o.UpdateDetails
}

// GetUpdateDetailsOk returns a tuple with the UpdateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetUpdateDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDetails) {
		return nil, false
	}
	return o.UpdateDetails, true
}

// HasUpdateDetails returns a boolean if a field has been set.
func (o *CompanyGateway) HasUpdateDetails() bool {
	if o != nil && !IsNil(o.UpdateDetails) {
		return true
	}

	return false
}

// SetUpdateDetails gets a reference to the given bool and assigns it to the UpdateDetails field.
func (o *CompanyGateway) SetUpdateDetails(v bool) {
	o.UpdateDetails = &v
}

// GetFeesAndLimits returns the FeesAndLimits field value if set, zero value otherwise.
func (o *CompanyGateway) GetFeesAndLimits() []FeesAndLimits {
	if o == nil || IsNil(o.FeesAndLimits) {
		var ret []FeesAndLimits
		return ret
	}
	return o.FeesAndLimits
}

// GetFeesAndLimitsOk returns a tuple with the FeesAndLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetFeesAndLimitsOk() ([]FeesAndLimits, bool) {
	if o == nil || IsNil(o.FeesAndLimits) {
		return nil, false
	}
	return o.FeesAndLimits, true
}

// HasFeesAndLimits returns a boolean if a field has been set.
func (o *CompanyGateway) HasFeesAndLimits() bool {
	if o != nil && !IsNil(o.FeesAndLimits) {
		return true
	}

	return false
}

// SetFeesAndLimits gets a reference to the given []FeesAndLimits and assigns it to the FeesAndLimits field.
func (o *CompanyGateway) SetFeesAndLimits(v []FeesAndLimits) {
	o.FeesAndLimits = v
}

func (o CompanyGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GatewayKey) {
		toSerialize["gateway_key"] = o.GatewayKey
	}
	if !IsNil(o.AcceptedCreditCards) {
		toSerialize["accepted_credit_cards"] = o.AcceptedCreditCards
	}
	if !IsNil(o.RequireBillingAddress) {
		toSerialize["require_billing_address"] = o.RequireBillingAddress
	}
	if !IsNil(o.RequireShippingAddress) {
		toSerialize["require_shipping_address"] = o.RequireShippingAddress
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.UpdateDetails) {
		toSerialize["update_details"] = o.UpdateDetails
	}
	if !IsNil(o.FeesAndLimits) {
		toSerialize["fees_and_limits"] = o.FeesAndLimits
	}
	return toSerialize, nil
}

type NullableCompanyGateway struct {
	value *CompanyGateway
	isSet bool
}

func (v NullableCompanyGateway) Get() *CompanyGateway {
	return v.value
}

func (v *NullableCompanyGateway) Set(val *CompanyGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyGateway(val *CompanyGateway) *NullableCompanyGateway {
	return &NullableCompanyGateway{value: val, isSet: true}
}

func (v NullableCompanyGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
