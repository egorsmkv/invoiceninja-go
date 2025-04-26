# UserRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the user | [optional] [readonly] 
**FirstName** | Pointer to **string** | The first name of the user | [optional] 
**LastName** | Pointer to **string** | The last name of the user | [optional] 
**Email** | Pointer to **string** | The users email address | [optional] 
**Phone** | Pointer to **string** | The users phone number | [optional] 
**Signature** | Pointer to **string** | The users sign off signature | [optional] 
**Avatar** | Pointer to **string** | The users avatar | [optional] 
**AcceptedTermsVersion** | Pointer to **string** | The version of the invoice ninja terms that has been accepted by the user | [optional] [readonly] 
**OauthUserId** | Pointer to **string** | The provider id of the oauth entity | [optional] [readonly] 
**OauthProviderId** | Pointer to **string** | The oauth entity id | [optional] [readonly] 
**LanguageId** | Pointer to **string** | The language id of the user | [optional] 
**VerifiedPhoneNumber** | Pointer to **bool** | Boolean flag if the user has their phone verified. Required to settings up 2FA | [optional] [readonly] 
**SmsVerificationCode** | Pointer to **string** | The sms verification code for the user. Required to settings up 2FA | [optional] [readonly] 
**OauthUserTokenExpiry** | Pointer to **string** | The expiry date of the oauth token | [optional] [readonly] 
**HasPassword** | Pointer to **bool** | Boolean flag determining if the user has a password | [optional] [readonly] 
**LastConfirmedEmailAddress** | Pointer to **string** | The last confirmed email address of the user | [optional] [readonly] 
**CustomValue1** | Pointer to **string** | A custom value | [optional] 
**CustomValue2** | Pointer to **string** | A custom value | [optional] 
**CustomValue3** | Pointer to **string** | A custom value | [optional] 
**CustomValue4** | Pointer to **string** | A custom value | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the user has been deleted | [optional] [readonly] 
**Google2faSecret** | Pointer to **string** | The google 2fa secret for the user | [optional] [readonly] 

## Methods

### NewUserRef

`func NewUserRef() *UserRef`

NewUserRef instantiates a new UserRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRefWithDefaults

`func NewUserRefWithDefaults() *UserRef`

NewUserRefWithDefaults instantiates a new UserRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *UserRef) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserRef) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserRef) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserRef) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserRef) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserRef) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserRef) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserRef) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserRef) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRef) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRef) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRef) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UserRef) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserRef) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserRef) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserRef) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSignature

`func (o *UserRef) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserRef) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserRef) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UserRef) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetAvatar

`func (o *UserRef) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserRef) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserRef) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserRef) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAcceptedTermsVersion

`func (o *UserRef) GetAcceptedTermsVersion() string`

GetAcceptedTermsVersion returns the AcceptedTermsVersion field if non-nil, zero value otherwise.

### GetAcceptedTermsVersionOk

`func (o *UserRef) GetAcceptedTermsVersionOk() (*string, bool)`

GetAcceptedTermsVersionOk returns a tuple with the AcceptedTermsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTermsVersion

`func (o *UserRef) SetAcceptedTermsVersion(v string)`

SetAcceptedTermsVersion sets AcceptedTermsVersion field to given value.

### HasAcceptedTermsVersion

`func (o *UserRef) HasAcceptedTermsVersion() bool`

HasAcceptedTermsVersion returns a boolean if a field has been set.

### GetOauthUserId

`func (o *UserRef) GetOauthUserId() string`

GetOauthUserId returns the OauthUserId field if non-nil, zero value otherwise.

### GetOauthUserIdOk

`func (o *UserRef) GetOauthUserIdOk() (*string, bool)`

GetOauthUserIdOk returns a tuple with the OauthUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUserId

`func (o *UserRef) SetOauthUserId(v string)`

SetOauthUserId sets OauthUserId field to given value.

### HasOauthUserId

`func (o *UserRef) HasOauthUserId() bool`

HasOauthUserId returns a boolean if a field has been set.

### GetOauthProviderId

`func (o *UserRef) GetOauthProviderId() string`

GetOauthProviderId returns the OauthProviderId field if non-nil, zero value otherwise.

### GetOauthProviderIdOk

`func (o *UserRef) GetOauthProviderIdOk() (*string, bool)`

GetOauthProviderIdOk returns a tuple with the OauthProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthProviderId

`func (o *UserRef) SetOauthProviderId(v string)`

SetOauthProviderId sets OauthProviderId field to given value.

### HasOauthProviderId

`func (o *UserRef) HasOauthProviderId() bool`

HasOauthProviderId returns a boolean if a field has been set.

### GetLanguageId

`func (o *UserRef) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *UserRef) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *UserRef) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *UserRef) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetVerifiedPhoneNumber

`func (o *UserRef) GetVerifiedPhoneNumber() bool`

GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field if non-nil, zero value otherwise.

### GetVerifiedPhoneNumberOk

`func (o *UserRef) GetVerifiedPhoneNumberOk() (*bool, bool)`

GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPhoneNumber

`func (o *UserRef) SetVerifiedPhoneNumber(v bool)`

SetVerifiedPhoneNumber sets VerifiedPhoneNumber field to given value.

### HasVerifiedPhoneNumber

`func (o *UserRef) HasVerifiedPhoneNumber() bool`

HasVerifiedPhoneNumber returns a boolean if a field has been set.

### GetSmsVerificationCode

`func (o *UserRef) GetSmsVerificationCode() string`

GetSmsVerificationCode returns the SmsVerificationCode field if non-nil, zero value otherwise.

### GetSmsVerificationCodeOk

`func (o *UserRef) GetSmsVerificationCodeOk() (*string, bool)`

GetSmsVerificationCodeOk returns a tuple with the SmsVerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsVerificationCode

`func (o *UserRef) SetSmsVerificationCode(v string)`

SetSmsVerificationCode sets SmsVerificationCode field to given value.

### HasSmsVerificationCode

`func (o *UserRef) HasSmsVerificationCode() bool`

HasSmsVerificationCode returns a boolean if a field has been set.

### GetOauthUserTokenExpiry

`func (o *UserRef) GetOauthUserTokenExpiry() string`

GetOauthUserTokenExpiry returns the OauthUserTokenExpiry field if non-nil, zero value otherwise.

### GetOauthUserTokenExpiryOk

`func (o *UserRef) GetOauthUserTokenExpiryOk() (*string, bool)`

GetOauthUserTokenExpiryOk returns a tuple with the OauthUserTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUserTokenExpiry

`func (o *UserRef) SetOauthUserTokenExpiry(v string)`

SetOauthUserTokenExpiry sets OauthUserTokenExpiry field to given value.

### HasOauthUserTokenExpiry

`func (o *UserRef) HasOauthUserTokenExpiry() bool`

HasOauthUserTokenExpiry returns a boolean if a field has been set.

### GetHasPassword

`func (o *UserRef) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *UserRef) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *UserRef) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *UserRef) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetLastConfirmedEmailAddress

`func (o *UserRef) GetLastConfirmedEmailAddress() string`

GetLastConfirmedEmailAddress returns the LastConfirmedEmailAddress field if non-nil, zero value otherwise.

### GetLastConfirmedEmailAddressOk

`func (o *UserRef) GetLastConfirmedEmailAddressOk() (*string, bool)`

GetLastConfirmedEmailAddressOk returns a tuple with the LastConfirmedEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfirmedEmailAddress

`func (o *UserRef) SetLastConfirmedEmailAddress(v string)`

SetLastConfirmedEmailAddress sets LastConfirmedEmailAddress field to given value.

### HasLastConfirmedEmailAddress

`func (o *UserRef) HasLastConfirmedEmailAddress() bool`

HasLastConfirmedEmailAddress returns a boolean if a field has been set.

### GetCustomValue1

`func (o *UserRef) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *UserRef) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *UserRef) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *UserRef) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *UserRef) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *UserRef) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *UserRef) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *UserRef) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *UserRef) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *UserRef) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *UserRef) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *UserRef) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *UserRef) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *UserRef) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *UserRef) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *UserRef) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UserRef) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UserRef) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UserRef) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UserRef) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetGoogle2faSecret

`func (o *UserRef) GetGoogle2faSecret() string`

GetGoogle2faSecret returns the Google2faSecret field if non-nil, zero value otherwise.

### GetGoogle2faSecretOk

`func (o *UserRef) GetGoogle2faSecretOk() (*string, bool)`

GetGoogle2faSecretOk returns a tuple with the Google2faSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle2faSecret

`func (o *UserRef) SetGoogle2faSecret(v string)`

SetGoogle2faSecret sets Google2faSecret field to given value.

### HasGoogle2faSecret

`func (o *UserRef) HasGoogle2faSecret() bool`

HasGoogle2faSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


