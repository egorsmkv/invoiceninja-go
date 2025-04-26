# VendorContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the vendor contact | [optional] [readonly] 
**UserId** | Pointer to **string** | The hashed id of the user id | [optional] [readonly] 
**VendorId** | Pointer to **string** | The hashed id of the vendor | [optional] [readonly] 
**FirstName** | Pointer to **string** | The first name of the contact | [optional] 
**LastName** | Pointer to **string** | The last name of the contact | [optional] 
**ContactKey** | Pointer to **string** | A unique identifier for the contact | [optional] [readonly] 
**ConfirmationCode** | Pointer to **string** | The confirmation code used to authenticate the contacts email address | [optional] [readonly] 
**Phone** | Pointer to **string** | The contacts phone number | [optional] 
**CustomValue1** | Pointer to **string** | A custom value | [optional] 
**CustomValue2** | Pointer to **string** | A custom value | [optional] 
**CustomValue3** | Pointer to **string** | A custom value | [optional] 
**CustomValue4** | Pointer to **string** | A custom value | [optional] 
**Email** | Pointer to **string** | The contact email address | [optional] 
**EmailVerifiedAt** | Pointer to **float32** | The date which the contact confirmed their email | [optional] [readonly] 
**Password** | Pointer to **string** | The hashed password of the contact | [optional] 
**IsPrimary** | Pointer to **bool** | Boolean flag determining if the contact is the primary contact for the vendor | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**DeletedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 

## Methods

### NewVendorContact

`func NewVendorContact() *VendorContact`

NewVendorContact instantiates a new VendorContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorContactWithDefaults

`func NewVendorContactWithDefaults() *VendorContact`

NewVendorContactWithDefaults instantiates a new VendorContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VendorContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *VendorContact) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VendorContact) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VendorContact) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *VendorContact) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVendorId

`func (o *VendorContact) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *VendorContact) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *VendorContact) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *VendorContact) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetFirstName

`func (o *VendorContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VendorContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VendorContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *VendorContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *VendorContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VendorContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VendorContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *VendorContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetContactKey

`func (o *VendorContact) GetContactKey() string`

GetContactKey returns the ContactKey field if non-nil, zero value otherwise.

### GetContactKeyOk

`func (o *VendorContact) GetContactKeyOk() (*string, bool)`

GetContactKeyOk returns a tuple with the ContactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactKey

`func (o *VendorContact) SetContactKey(v string)`

SetContactKey sets ContactKey field to given value.

### HasContactKey

`func (o *VendorContact) HasContactKey() bool`

HasContactKey returns a boolean if a field has been set.

### GetConfirmationCode

`func (o *VendorContact) GetConfirmationCode() string`

GetConfirmationCode returns the ConfirmationCode field if non-nil, zero value otherwise.

### GetConfirmationCodeOk

`func (o *VendorContact) GetConfirmationCodeOk() (*string, bool)`

GetConfirmationCodeOk returns a tuple with the ConfirmationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationCode

`func (o *VendorContact) SetConfirmationCode(v string)`

SetConfirmationCode sets ConfirmationCode field to given value.

### HasConfirmationCode

`func (o *VendorContact) HasConfirmationCode() bool`

HasConfirmationCode returns a boolean if a field has been set.

### GetPhone

`func (o *VendorContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VendorContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VendorContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *VendorContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCustomValue1

`func (o *VendorContact) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *VendorContact) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *VendorContact) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *VendorContact) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *VendorContact) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *VendorContact) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *VendorContact) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *VendorContact) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *VendorContact) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *VendorContact) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *VendorContact) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *VendorContact) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *VendorContact) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *VendorContact) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *VendorContact) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *VendorContact) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetEmail

`func (o *VendorContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VendorContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VendorContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VendorContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerifiedAt

`func (o *VendorContact) GetEmailVerifiedAt() float32`

GetEmailVerifiedAt returns the EmailVerifiedAt field if non-nil, zero value otherwise.

### GetEmailVerifiedAtOk

`func (o *VendorContact) GetEmailVerifiedAtOk() (*float32, bool)`

GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerifiedAt

`func (o *VendorContact) SetEmailVerifiedAt(v float32)`

SetEmailVerifiedAt sets EmailVerifiedAt field to given value.

### HasEmailVerifiedAt

`func (o *VendorContact) HasEmailVerifiedAt() bool`

HasEmailVerifiedAt returns a boolean if a field has been set.

### GetPassword

`func (o *VendorContact) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VendorContact) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VendorContact) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VendorContact) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIsPrimary

`func (o *VendorContact) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *VendorContact) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *VendorContact) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *VendorContact) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VendorContact) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VendorContact) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VendorContact) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VendorContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VendorContact) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VendorContact) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VendorContact) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VendorContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *VendorContact) GetDeletedAt() float32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VendorContact) GetDeletedAtOk() (*float32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VendorContact) SetDeletedAt(v float32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VendorContact) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


