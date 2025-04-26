# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The location hashed id | [optional] 
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
**IsDeleted** | Pointer to **bool** | Indicates if the location has been deleted | [optional] 
**IsShippingLocation** | Pointer to **bool** | Indicates if this is a shipping location | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] [readonly] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**VendorId** | Pointer to **string** | The vendor hashed id | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp of when the location was created | [optional] 
**UpdatedAt** | Pointer to **int32** | Timestamp of when the location was last updated | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress1

`func (o *Location) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *Location) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *Location) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *Location) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *Location) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *Location) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *Location) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *Location) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Location) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Location) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Location) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Location) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *Location) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Location) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Location) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Location) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryId

`func (o *Location) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Location) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Location) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Location) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Location) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Location) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Location) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Location) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### SetCustomValue1Nil

`func (o *Location) SetCustomValue1Nil(b bool)`

 SetCustomValue1Nil sets the value for CustomValue1 to be an explicit nil

### UnsetCustomValue1
`func (o *Location) UnsetCustomValue1()`

UnsetCustomValue1 ensures that no value is present for CustomValue1, not even an explicit nil
### GetCustomValue2

`func (o *Location) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Location) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Location) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Location) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### SetCustomValue2Nil

`func (o *Location) SetCustomValue2Nil(b bool)`

 SetCustomValue2Nil sets the value for CustomValue2 to be an explicit nil

### UnsetCustomValue2
`func (o *Location) UnsetCustomValue2()`

UnsetCustomValue2 ensures that no value is present for CustomValue2, not even an explicit nil
### GetCustomValue3

`func (o *Location) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Location) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Location) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Location) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### SetCustomValue3Nil

`func (o *Location) SetCustomValue3Nil(b bool)`

 SetCustomValue3Nil sets the value for CustomValue3 to be an explicit nil

### UnsetCustomValue3
`func (o *Location) UnsetCustomValue3()`

UnsetCustomValue3 ensures that no value is present for CustomValue3, not even an explicit nil
### GetCustomValue4

`func (o *Location) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Location) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Location) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Location) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### SetCustomValue4Nil

`func (o *Location) SetCustomValue4Nil(b bool)`

 SetCustomValue4Nil sets the value for CustomValue4 to be an explicit nil

### UnsetCustomValue4
`func (o *Location) UnsetCustomValue4()`

UnsetCustomValue4 ensures that no value is present for CustomValue4, not even an explicit nil
### GetIsDeleted

`func (o *Location) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Location) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Location) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Location) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsShippingLocation

`func (o *Location) GetIsShippingLocation() bool`

GetIsShippingLocation returns the IsShippingLocation field if non-nil, zero value otherwise.

### GetIsShippingLocationOk

`func (o *Location) GetIsShippingLocationOk() (*bool, bool)`

GetIsShippingLocationOk returns a tuple with the IsShippingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingLocation

`func (o *Location) SetIsShippingLocation(v bool)`

SetIsShippingLocation sets IsShippingLocation field to given value.

### HasIsShippingLocation

`func (o *Location) HasIsShippingLocation() bool`

HasIsShippingLocation returns a boolean if a field has been set.

### GetUserId

`func (o *Location) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Location) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Location) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Location) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Location) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Location) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Location) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Location) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *Location) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Location) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Location) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Location) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetVendorId

`func (o *Location) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *Location) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *Location) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *Location) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Location) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Location) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Location) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Location) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Location) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Location) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Location) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Location) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


