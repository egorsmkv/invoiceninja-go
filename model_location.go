package openapi

import (
	"encoding/json"
)

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location struct for Location
type Location struct {
	// The location hashed id
	Id *string `json:"id,omitempty"`
	// The location name
	Name *string `json:"name,omitempty"`
	// The first line of the address
	Address1 *string `json:"address1,omitempty"`
	// The second line of the address
	Address2 *string `json:"address2,omitempty"`
	// The city name
	City *string `json:"city,omitempty"`
	// The state or region
	State *string `json:"state,omitempty"`
	// The postal or zip code
	PostalCode *string `json:"postal_code,omitempty"`
	// The ID of the associated country
	CountryId *string `json:"country_id,omitempty"`
	// Indicates if the location has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Indicates if this is a shipping location
	IsShippingLocation *bool `json:"is_shipping_location,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The vendor hashed id
	VendorId *string `json:"vendor_id,omitempty"`
	// Timestamp of when the location was created
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Timestamp of when the location was last updated
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// Custom field value 1
	CustomValue1 NullableString `json:"custom_value1,omitempty"`
	// Custom field value 2
	CustomValue2 NullableString `json:"custom_value2,omitempty"`
	// Custom field value 3
	CustomValue3 NullableString `json:"custom_value3,omitempty"`
	// Custom field value 4
	CustomValue4 NullableString `json:"custom_value4,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Location) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Location) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Location) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Location) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Location) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Location) SetName(v string) {
	o.Name = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *Location) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *Location) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *Location) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *Location) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *Location) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *Location) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Location) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Location) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Location) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Location) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Location) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Location) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Location) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Location) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Location) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *Location) GetCountryId() string {
	if o == nil || IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCountryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *Location) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *Location) SetCountryId(v string) {
	o.CountryId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue1.Get()
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCustomValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue1.Get(), o.CustomValue1.IsSet()
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Location) HasCustomValue1() bool {
	if o != nil && o.CustomValue1.IsSet() {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given NullableString and assigns it to the CustomValue1 field.
func (o *Location) SetCustomValue1(v string) {
	o.CustomValue1.Set(&v)
}

// SetCustomValue1Nil sets the value for CustomValue1 to be an explicit nil
func (o *Location) SetCustomValue1Nil() {
	o.CustomValue1.Set(nil)
}

// UnsetCustomValue1 ensures that no value is present for CustomValue1, not even an explicit nil
func (o *Location) UnsetCustomValue1() {
	o.CustomValue1.Unset()
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue2.Get()
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCustomValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue2.Get(), o.CustomValue2.IsSet()
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Location) HasCustomValue2() bool {
	if o != nil && o.CustomValue2.IsSet() {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given NullableString and assigns it to the CustomValue2 field.
func (o *Location) SetCustomValue2(v string) {
	o.CustomValue2.Set(&v)
}

// SetCustomValue2Nil sets the value for CustomValue2 to be an explicit nil
func (o *Location) SetCustomValue2Nil() {
	o.CustomValue2.Set(nil)
}

// UnsetCustomValue2 ensures that no value is present for CustomValue2, not even an explicit nil
func (o *Location) UnsetCustomValue2() {
	o.CustomValue2.Unset()
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue3.Get()
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCustomValue3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue3.Get(), o.CustomValue3.IsSet()
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Location) HasCustomValue3() bool {
	if o != nil && o.CustomValue3.IsSet() {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given NullableString and assigns it to the CustomValue3 field.
func (o *Location) SetCustomValue3(v string) {
	o.CustomValue3.Set(&v)
}

// SetCustomValue3Nil sets the value for CustomValue3 to be an explicit nil
func (o *Location) SetCustomValue3Nil() {
	o.CustomValue3.Set(nil)
}

// UnsetCustomValue3 ensures that no value is present for CustomValue3, not even an explicit nil
func (o *Location) UnsetCustomValue3() {
	o.CustomValue3.Unset()
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue4.Get()
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCustomValue4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue4.Get(), o.CustomValue4.IsSet()
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Location) HasCustomValue4() bool {
	if o != nil && o.CustomValue4.IsSet() {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given NullableString and assigns it to the CustomValue4 field.
func (o *Location) SetCustomValue4(v string) {
	o.CustomValue4.Set(&v)
}

// SetCustomValue4Nil sets the value for CustomValue4 to be an explicit nil
func (o *Location) SetCustomValue4Nil() {
	o.CustomValue4.Set(nil)
}

// UnsetCustomValue4 ensures that no value is present for CustomValue4, not even an explicit nil
func (o *Location) UnsetCustomValue4() {
	o.CustomValue4.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Location) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Location) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Location) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetIsShippingLocation returns the IsShippingLocation field value if set, zero value otherwise.
func (o *Location) GetIsShippingLocation() bool {
	if o == nil || IsNil(o.IsShippingLocation) {
		var ret bool
		return ret
	}
	return *o.IsShippingLocation
}

// GetIsShippingLocationOk returns a tuple with the IsShippingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIsShippingLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShippingLocation) {
		return nil, false
	}
	return o.IsShippingLocation, true
}

// HasIsShippingLocation returns a boolean if a field has been set.
func (o *Location) HasIsShippingLocation() bool {
	if o != nil && !IsNil(o.IsShippingLocation) {
		return true
	}

	return false
}

// SetIsShippingLocation gets a reference to the given bool and assigns it to the IsShippingLocation field.
func (o *Location) SetIsShippingLocation(v bool) {
	o.IsShippingLocation = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Location) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Location) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Location) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Location) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Location) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Location) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Location) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Location) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Location) SetClientId(v string) {
	o.ClientId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *Location) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *Location) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *Location) SetVendorId(v string) {
	o.VendorId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Location) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Location) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Location) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Location) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Location) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Location) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
	}
	if o.CustomValue1.IsSet() {
		toSerialize["custom_value1"] = o.CustomValue1.Get()
	}
	if o.CustomValue2.IsSet() {
		toSerialize["custom_value2"] = o.CustomValue2.Get()
	}
	if o.CustomValue3.IsSet() {
		toSerialize["custom_value3"] = o.CustomValue3.Get()
	}
	if o.CustomValue4.IsSet() {
		toSerialize["custom_value4"] = o.CustomValue4.Get()
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.IsShippingLocation) {
		toSerialize["is_shipping_location"] = o.IsShippingLocation
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
