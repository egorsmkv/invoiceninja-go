# ClientContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed if of the contact | [optional] [readonly] 
**UserId** | Pointer to **string** | The hashed id of the user who created the contact | [optional] [readonly] 
**ClientId** | Pointer to **string** | The hashed id of the client | [optional] [readonly] 
**FirstName** | Pointer to **string** | The first name of the contact | [optional] 
**LastName** | Pointer to **string** | The last name of the contact | [optional] 
**Phone** | Pointer to **string** | The phone number of the contact | [optional] 
**CustomValue1** | Pointer to **string** | A Custom field value | [optional] 
**CustomValue2** | Pointer to **string** | A Custom field value | [optional] 
**CustomValue3** | Pointer to **string** | A Custom field value | [optional] 
**CustomValue4** | Pointer to **string** | A Custom field value | [optional] 
**Email** | Pointer to **string** | The email of the contact | [optional] 
**AcceptedTermsVersion** | Pointer to **string** | The terms of service which the contact has accpeted | [optional] [readonly] 
**Password** | Pointer to **string** | The hashed password of the contact | [optional] 
**ConfirmationCode** | Pointer to **string** | The confirmation code used to authenticate the contacts email address | [optional] [readonly] 
**Token** | Pointer to **string** | A uuid based token. | [optional] [readonly] 
**ContactKey** | Pointer to **string** | A unique identifier for the contact | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** | Defines is this contact is the primary contact for the client | [optional] 
**Confirmed** | Pointer to **bool** | Boolean value confirms the user has confirmed their account. | [optional] 
**IsLocked** | Pointer to **bool** | Boolean value defines if the contact has been locked out. | [optional] 
**SendEmail** | Pointer to **bool** | Boolean value determines is this contact should receive emails | [optional] 
**FailedLogins** | Pointer to **float32** | The number of failed logins the contact has had | [optional] [readonly] 
**EmailVerifiedAt** | Pointer to **float32** | The date which the contact confirmed their email | [optional] [readonly] 
**LastLogin** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**DeletedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 

## Methods

### NewClientContact

`func NewClientContact() *ClientContact`

NewClientContact instantiates a new ClientContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientContactWithDefaults

`func NewClientContactWithDefaults() *ClientContact`

NewClientContactWithDefaults instantiates a new ClientContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *ClientContact) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ClientContact) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ClientContact) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ClientContact) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetClientId

`func (o *ClientContact) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientContact) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientContact) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientContact) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFirstName

`func (o *ClientContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ClientContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ClientContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ClientContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ClientContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ClientContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ClientContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ClientContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *ClientContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClientContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClientContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClientContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCustomValue1

`func (o *ClientContact) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *ClientContact) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *ClientContact) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *ClientContact) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *ClientContact) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *ClientContact) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *ClientContact) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *ClientContact) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *ClientContact) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *ClientContact) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *ClientContact) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *ClientContact) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *ClientContact) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *ClientContact) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *ClientContact) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *ClientContact) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetEmail

`func (o *ClientContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClientContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClientContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClientContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAcceptedTermsVersion

`func (o *ClientContact) GetAcceptedTermsVersion() string`

GetAcceptedTermsVersion returns the AcceptedTermsVersion field if non-nil, zero value otherwise.

### GetAcceptedTermsVersionOk

`func (o *ClientContact) GetAcceptedTermsVersionOk() (*string, bool)`

GetAcceptedTermsVersionOk returns a tuple with the AcceptedTermsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTermsVersion

`func (o *ClientContact) SetAcceptedTermsVersion(v string)`

SetAcceptedTermsVersion sets AcceptedTermsVersion field to given value.

### HasAcceptedTermsVersion

`func (o *ClientContact) HasAcceptedTermsVersion() bool`

HasAcceptedTermsVersion returns a boolean if a field has been set.

### GetPassword

`func (o *ClientContact) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClientContact) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClientContact) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClientContact) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConfirmationCode

`func (o *ClientContact) GetConfirmationCode() string`

GetConfirmationCode returns the ConfirmationCode field if non-nil, zero value otherwise.

### GetConfirmationCodeOk

`func (o *ClientContact) GetConfirmationCodeOk() (*string, bool)`

GetConfirmationCodeOk returns a tuple with the ConfirmationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationCode

`func (o *ClientContact) SetConfirmationCode(v string)`

SetConfirmationCode sets ConfirmationCode field to given value.

### HasConfirmationCode

`func (o *ClientContact) HasConfirmationCode() bool`

HasConfirmationCode returns a boolean if a field has been set.

### GetToken

`func (o *ClientContact) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientContact) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientContact) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientContact) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetContactKey

`func (o *ClientContact) GetContactKey() string`

GetContactKey returns the ContactKey field if non-nil, zero value otherwise.

### GetContactKeyOk

`func (o *ClientContact) GetContactKeyOk() (*string, bool)`

GetContactKeyOk returns a tuple with the ContactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactKey

`func (o *ClientContact) SetContactKey(v string)`

SetContactKey sets ContactKey field to given value.

### HasContactKey

`func (o *ClientContact) HasContactKey() bool`

HasContactKey returns a boolean if a field has been set.

### GetIsPrimary

`func (o *ClientContact) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ClientContact) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ClientContact) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *ClientContact) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetConfirmed

`func (o *ClientContact) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *ClientContact) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *ClientContact) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *ClientContact) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetIsLocked

`func (o *ClientContact) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ClientContact) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ClientContact) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ClientContact) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetSendEmail

`func (o *ClientContact) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *ClientContact) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *ClientContact) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *ClientContact) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetFailedLogins

`func (o *ClientContact) GetFailedLogins() float32`

GetFailedLogins returns the FailedLogins field if non-nil, zero value otherwise.

### GetFailedLoginsOk

`func (o *ClientContact) GetFailedLoginsOk() (*float32, bool)`

GetFailedLoginsOk returns a tuple with the FailedLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLogins

`func (o *ClientContact) SetFailedLogins(v float32)`

SetFailedLogins sets FailedLogins field to given value.

### HasFailedLogins

`func (o *ClientContact) HasFailedLogins() bool`

HasFailedLogins returns a boolean if a field has been set.

### GetEmailVerifiedAt

`func (o *ClientContact) GetEmailVerifiedAt() float32`

GetEmailVerifiedAt returns the EmailVerifiedAt field if non-nil, zero value otherwise.

### GetEmailVerifiedAtOk

`func (o *ClientContact) GetEmailVerifiedAtOk() (*float32, bool)`

GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerifiedAt

`func (o *ClientContact) SetEmailVerifiedAt(v float32)`

SetEmailVerifiedAt sets EmailVerifiedAt field to given value.

### HasEmailVerifiedAt

`func (o *ClientContact) HasEmailVerifiedAt() bool`

HasEmailVerifiedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *ClientContact) GetLastLogin() float32`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *ClientContact) GetLastLoginOk() (*float32, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *ClientContact) SetLastLogin(v float32)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *ClientContact) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClientContact) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientContact) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientContact) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClientContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClientContact) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientContact) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientContact) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClientContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ClientContact) GetDeletedAt() float32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ClientContact) GetDeletedAtOk() (*float32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ClientContact) SetDeletedAt(v float32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ClientContact) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


