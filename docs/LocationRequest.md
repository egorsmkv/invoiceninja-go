# LocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The location name | [optional] 
**Address1** | Pointer to **string** | The first line of the address | [optional] 
**Address2** | Pointer to **string** | The second line of the address | [optional] 
**City** | Pointer to **string** | The city name | [optional] 
**State** | Pointer to **string** | The state or region | [optional] 
**PostalCode** | Pointer to **string** | The postal or zip code | [optional] 
**CountryId** | Pointer to **string** | The ID of the associated country | [optional] 
**CustomValue1** | Pointer to **NullableString** | Custom field value 1 | [optional] 
**CustomValue2** | Pointer to **NullableString** | Custom field value 2 | [optional] 
**CustomValue3** | Pointer to **NullableString** | Custom field value 3 | [optional] 
**CustomValue4** | Pointer to **NullableString** | Custom field value 4 | [optional] 
**IsShippingLocation** | Pointer to **bool** | Indicates if this is a shipping location | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**VendorId** | Pointer to **string** | The vendor hashed id | [optional] 

## Methods

### NewLocationRequest

`func NewLocationRequest() *LocationRequest`

NewLocationRequest instantiates a new LocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRequestWithDefaults

`func NewLocationRequestWithDefaults() *LocationRequest`

NewLocationRequestWithDefaults instantiates a new LocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LocationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress1

`func (o *LocationRequest) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationRequest) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationRequest) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationRequest) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *LocationRequest) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationRequest) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationRequest) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationRequest) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *LocationRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LocationRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *LocationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *LocationRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LocationRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LocationRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LocationRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryId

`func (o *LocationRequest) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *LocationRequest) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *LocationRequest) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *LocationRequest) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *LocationRequest) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *LocationRequest) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *LocationRequest) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *LocationRequest) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### SetCustomValue1Nil

`func (o *LocationRequest) SetCustomValue1Nil(b bool)`

 SetCustomValue1Nil sets the value for CustomValue1 to be an explicit nil

### UnsetCustomValue1
`func (o *LocationRequest) UnsetCustomValue1()`

UnsetCustomValue1 ensures that no value is present for CustomValue1, not even an explicit nil
### GetCustomValue2

`func (o *LocationRequest) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *LocationRequest) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *LocationRequest) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *LocationRequest) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### SetCustomValue2Nil

`func (o *LocationRequest) SetCustomValue2Nil(b bool)`

 SetCustomValue2Nil sets the value for CustomValue2 to be an explicit nil

### UnsetCustomValue2
`func (o *LocationRequest) UnsetCustomValue2()`

UnsetCustomValue2 ensures that no value is present for CustomValue2, not even an explicit nil
### GetCustomValue3

`func (o *LocationRequest) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *LocationRequest) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *LocationRequest) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *LocationRequest) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### SetCustomValue3Nil

`func (o *LocationRequest) SetCustomValue3Nil(b bool)`

 SetCustomValue3Nil sets the value for CustomValue3 to be an explicit nil

### UnsetCustomValue3
`func (o *LocationRequest) UnsetCustomValue3()`

UnsetCustomValue3 ensures that no value is present for CustomValue3, not even an explicit nil
### GetCustomValue4

`func (o *LocationRequest) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *LocationRequest) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *LocationRequest) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *LocationRequest) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### SetCustomValue4Nil

`func (o *LocationRequest) SetCustomValue4Nil(b bool)`

 SetCustomValue4Nil sets the value for CustomValue4 to be an explicit nil

### UnsetCustomValue4
`func (o *LocationRequest) UnsetCustomValue4()`

UnsetCustomValue4 ensures that no value is present for CustomValue4, not even an explicit nil
### GetIsShippingLocation

`func (o *LocationRequest) GetIsShippingLocation() bool`

GetIsShippingLocation returns the IsShippingLocation field if non-nil, zero value otherwise.

### GetIsShippingLocationOk

`func (o *LocationRequest) GetIsShippingLocationOk() (*bool, bool)`

GetIsShippingLocationOk returns a tuple with the IsShippingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingLocation

`func (o *LocationRequest) SetIsShippingLocation(v bool)`

SetIsShippingLocation sets IsShippingLocation field to given value.

### HasIsShippingLocation

`func (o *LocationRequest) HasIsShippingLocation() bool`

HasIsShippingLocation returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *LocationRequest) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *LocationRequest) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *LocationRequest) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *LocationRequest) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *LocationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *LocationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *LocationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *LocationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetVendorId

`func (o *LocationRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *LocationRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *LocationRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *LocationRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


