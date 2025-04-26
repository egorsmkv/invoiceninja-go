package openapi

import (
	"encoding/json"
)

// checks if the ClientGatewayToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientGatewayToken{}

// ClientGatewayToken struct for ClientGatewayToken
type ClientGatewayToken struct {
	// The hashed id of the client gateway token
	Id *string `json:"id,omitempty"`
	// The hashed_id of the client
	ClientId *string `json:"client_id,omitempty"`
	// The payment token
	Token *string `json:"token,omitempty"`
	// THe bank account routing number
	RoutingNumber *string `json:"routing_number,omitempty"`
	// The hashed id of the company gateway
	CompanyGatewayId *string `json:"company_gateway_id,omitempty"`
	// Flag determining if the token is the default payment method
	IsDefault *bool `json:"is_default,omitempty"`
}

// NewClientGatewayToken instantiates a new ClientGatewayToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGatewayToken() *ClientGatewayToken {
	this := ClientGatewayToken{}
	return &this
}

// NewClientGatewayTokenWithDefaults instantiates a new ClientGatewayToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGatewayTokenWithDefaults() *ClientGatewayToken {
	this := ClientGatewayToken{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientGatewayToken) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientGatewayToken) SetClientId(v string) {
	o.ClientId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ClientGatewayToken) SetToken(v string) {
	o.Token = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *ClientGatewayToken) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetCompanyGatewayId returns the CompanyGatewayId field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetCompanyGatewayId() string {
	if o == nil || IsNil(o.CompanyGatewayId) {
		var ret string
		return ret
	}
	return *o.CompanyGatewayId
}

// GetCompanyGatewayIdOk returns a tuple with the CompanyGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetCompanyGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyGatewayId) {
		return nil, false
	}
	return o.CompanyGatewayId, true
}

// HasCompanyGatewayId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasCompanyGatewayId() bool {
	if o != nil && !IsNil(o.CompanyGatewayId) {
		return true
	}

	return false
}

// SetCompanyGatewayId gets a reference to the given string and assigns it to the CompanyGatewayId field.
func (o *ClientGatewayToken) SetCompanyGatewayId(v string) {
	o.CompanyGatewayId = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ClientGatewayToken) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o ClientGatewayToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientGatewayToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.RoutingNumber) {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	if !IsNil(o.CompanyGatewayId) {
		toSerialize["company_gateway_id"] = o.CompanyGatewayId
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableClientGatewayToken struct {
	value *ClientGatewayToken
	isSet bool
}

func (v NullableClientGatewayToken) Get() *ClientGatewayToken {
	return v.value
}

func (v *NullableClientGatewayToken) Set(val *ClientGatewayToken) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGatewayToken) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGatewayToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGatewayToken(val *ClientGatewayToken) *NullableClientGatewayToken {
	return &NullableClientGatewayToken{value: val, isSet: true}
}

func (v NullableClientGatewayToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGatewayToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
