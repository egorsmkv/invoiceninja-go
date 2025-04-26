package openapi

import (
	"encoding/json"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	// The hashed id of the user
	Id *string `json:"id,omitempty"`
	// The first name of the user
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the user
	LastName *string `json:"last_name,omitempty"`
	// The users email address
	Email *string `json:"email,omitempty"`
	// The users phone number
	Phone *string `json:"phone,omitempty"`
	// The users sign off signature
	Signature *string `json:"signature,omitempty"`
	// The users avatar
	Avatar *string `json:"avatar,omitempty"`
	// The version of the invoice ninja terms that has been accepted by the user
	AcceptedTermsVersion *string `json:"accepted_terms_version,omitempty"`
	// The provider id of the oauth entity
	OauthUserId *string `json:"oauth_user_id,omitempty"`
	// The oauth entity id
	OauthProviderId *string `json:"oauth_provider_id,omitempty"`
	// The language id of the user
	LanguageId *string `json:"language_id,omitempty"`
	// Boolean flag if the user has their phone verified. Required to settings up 2FA
	VerifiedPhoneNumber *bool `json:"verified_phone_number,omitempty"`
	// The sms verification code for the user. Required to settings up 2FA
	SmsVerificationCode *string `json:"sms_verification_code,omitempty"`
	// The expiry date of the oauth token
	OauthUserTokenExpiry *string `json:"oauth_user_token_expiry,omitempty"`
	// Boolean flag determining if the user has a password
	HasPassword *bool `json:"has_password,omitempty"`
	// The last confirmed email address of the user
	LastConfirmedEmailAddress *string `json:"last_confirmed_email_address,omitempty"`
	// A custom value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// Boolean flag determining if the user has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The google 2fa secret for the user
	Google2faSecret *string         `json:"google_2fa_secret,omitempty"`
	CompanyUser     *CompanyUserRef `json:"company_user,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *User) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *User) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *User) SetSignature(v string) {
	o.Signature = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *User) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *User) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *User) SetAvatar(v string) {
	o.Avatar = &v
}

// GetAcceptedTermsVersion returns the AcceptedTermsVersion field value if set, zero value otherwise.
func (o *User) GetAcceptedTermsVersion() string {
	if o == nil || IsNil(o.AcceptedTermsVersion) {
		var ret string
		return ret
	}
	return *o.AcceptedTermsVersion
}

// GetAcceptedTermsVersionOk returns a tuple with the AcceptedTermsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAcceptedTermsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptedTermsVersion) {
		return nil, false
	}
	return o.AcceptedTermsVersion, true
}

// HasAcceptedTermsVersion returns a boolean if a field has been set.
func (o *User) HasAcceptedTermsVersion() bool {
	if o != nil && !IsNil(o.AcceptedTermsVersion) {
		return true
	}

	return false
}

// SetAcceptedTermsVersion gets a reference to the given string and assigns it to the AcceptedTermsVersion field.
func (o *User) SetAcceptedTermsVersion(v string) {
	o.AcceptedTermsVersion = &v
}

// GetOauthUserId returns the OauthUserId field value if set, zero value otherwise.
func (o *User) GetOauthUserId() string {
	if o == nil || IsNil(o.OauthUserId) {
		var ret string
		return ret
	}
	return *o.OauthUserId
}

// GetOauthUserIdOk returns a tuple with the OauthUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOauthUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.OauthUserId) {
		return nil, false
	}
	return o.OauthUserId, true
}

// HasOauthUserId returns a boolean if a field has been set.
func (o *User) HasOauthUserId() bool {
	if o != nil && !IsNil(o.OauthUserId) {
		return true
	}

	return false
}

// SetOauthUserId gets a reference to the given string and assigns it to the OauthUserId field.
func (o *User) SetOauthUserId(v string) {
	o.OauthUserId = &v
}

// GetOauthProviderId returns the OauthProviderId field value if set, zero value otherwise.
func (o *User) GetOauthProviderId() string {
	if o == nil || IsNil(o.OauthProviderId) {
		var ret string
		return ret
	}
	return *o.OauthProviderId
}

// GetOauthProviderIdOk returns a tuple with the OauthProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOauthProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OauthProviderId) {
		return nil, false
	}
	return o.OauthProviderId, true
}

// HasOauthProviderId returns a boolean if a field has been set.
func (o *User) HasOauthProviderId() bool {
	if o != nil && !IsNil(o.OauthProviderId) {
		return true
	}

	return false
}

// SetOauthProviderId gets a reference to the given string and assigns it to the OauthProviderId field.
func (o *User) SetOauthProviderId(v string) {
	o.OauthProviderId = &v
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise.
func (o *User) GetLanguageId() string {
	if o == nil || IsNil(o.LanguageId) {
		var ret string
		return ret
	}
	return *o.LanguageId
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLanguageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageId) {
		return nil, false
	}
	return o.LanguageId, true
}

// HasLanguageId returns a boolean if a field has been set.
func (o *User) HasLanguageId() bool {
	if o != nil && !IsNil(o.LanguageId) {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given string and assigns it to the LanguageId field.
func (o *User) SetLanguageId(v string) {
	o.LanguageId = &v
}

// GetVerifiedPhoneNumber returns the VerifiedPhoneNumber field value if set, zero value otherwise.
func (o *User) GetVerifiedPhoneNumber() bool {
	if o == nil || IsNil(o.VerifiedPhoneNumber) {
		var ret bool
		return ret
	}
	return *o.VerifiedPhoneNumber
}

// GetVerifiedPhoneNumberOk returns a tuple with the VerifiedPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedPhoneNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifiedPhoneNumber) {
		return nil, false
	}
	return o.VerifiedPhoneNumber, true
}

// HasVerifiedPhoneNumber returns a boolean if a field has been set.
func (o *User) HasVerifiedPhoneNumber() bool {
	if o != nil && !IsNil(o.VerifiedPhoneNumber) {
		return true
	}

	return false
}

// SetVerifiedPhoneNumber gets a reference to the given bool and assigns it to the VerifiedPhoneNumber field.
func (o *User) SetVerifiedPhoneNumber(v bool) {
	o.VerifiedPhoneNumber = &v
}

// GetSmsVerificationCode returns the SmsVerificationCode field value if set, zero value otherwise.
func (o *User) GetSmsVerificationCode() string {
	if o == nil || IsNil(o.SmsVerificationCode) {
		var ret string
		return ret
	}
	return *o.SmsVerificationCode
}

// GetSmsVerificationCodeOk returns a tuple with the SmsVerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSmsVerificationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SmsVerificationCode) {
		return nil, false
	}
	return o.SmsVerificationCode, true
}

// HasSmsVerificationCode returns a boolean if a field has been set.
func (o *User) HasSmsVerificationCode() bool {
	if o != nil && !IsNil(o.SmsVerificationCode) {
		return true
	}

	return false
}

// SetSmsVerificationCode gets a reference to the given string and assigns it to the SmsVerificationCode field.
func (o *User) SetSmsVerificationCode(v string) {
	o.SmsVerificationCode = &v
}

// GetOauthUserTokenExpiry returns the OauthUserTokenExpiry field value if set, zero value otherwise.
func (o *User) GetOauthUserTokenExpiry() string {
	if o == nil || IsNil(o.OauthUserTokenExpiry) {
		var ret string
		return ret
	}
	return *o.OauthUserTokenExpiry
}

// GetOauthUserTokenExpiryOk returns a tuple with the OauthUserTokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOauthUserTokenExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.OauthUserTokenExpiry) {
		return nil, false
	}
	return o.OauthUserTokenExpiry, true
}

// HasOauthUserTokenExpiry returns a boolean if a field has been set.
func (o *User) HasOauthUserTokenExpiry() bool {
	if o != nil && !IsNil(o.OauthUserTokenExpiry) {
		return true
	}

	return false
}

// SetOauthUserTokenExpiry gets a reference to the given string and assigns it to the OauthUserTokenExpiry field.
func (o *User) SetOauthUserTokenExpiry(v string) {
	o.OauthUserTokenExpiry = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *User) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *User) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *User) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetLastConfirmedEmailAddress returns the LastConfirmedEmailAddress field value if set, zero value otherwise.
func (o *User) GetLastConfirmedEmailAddress() string {
	if o == nil || IsNil(o.LastConfirmedEmailAddress) {
		var ret string
		return ret
	}
	return *o.LastConfirmedEmailAddress
}

// GetLastConfirmedEmailAddressOk returns a tuple with the LastConfirmedEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastConfirmedEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LastConfirmedEmailAddress) {
		return nil, false
	}
	return o.LastConfirmedEmailAddress, true
}

// HasLastConfirmedEmailAddress returns a boolean if a field has been set.
func (o *User) HasLastConfirmedEmailAddress() bool {
	if o != nil && !IsNil(o.LastConfirmedEmailAddress) {
		return true
	}

	return false
}

// SetLastConfirmedEmailAddress gets a reference to the given string and assigns it to the LastConfirmedEmailAddress field.
func (o *User) SetLastConfirmedEmailAddress(v string) {
	o.LastConfirmedEmailAddress = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *User) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *User) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *User) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *User) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *User) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *User) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *User) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *User) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *User) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *User) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *User) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *User) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *User) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *User) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *User) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetGoogle2faSecret returns the Google2faSecret field value if set, zero value otherwise.
func (o *User) GetGoogle2faSecret() string {
	if o == nil || IsNil(o.Google2faSecret) {
		var ret string
		return ret
	}
	return *o.Google2faSecret
}

// GetGoogle2faSecretOk returns a tuple with the Google2faSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGoogle2faSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Google2faSecret) {
		return nil, false
	}
	return o.Google2faSecret, true
}

// HasGoogle2faSecret returns a boolean if a field has been set.
func (o *User) HasGoogle2faSecret() bool {
	if o != nil && !IsNil(o.Google2faSecret) {
		return true
	}

	return false
}

// SetGoogle2faSecret gets a reference to the given string and assigns it to the Google2faSecret field.
func (o *User) SetGoogle2faSecret(v string) {
	o.Google2faSecret = &v
}

// GetCompanyUser returns the CompanyUser field value if set, zero value otherwise.
func (o *User) GetCompanyUser() CompanyUserRef {
	if o == nil || IsNil(o.CompanyUser) {
		var ret CompanyUserRef
		return ret
	}
	return *o.CompanyUser
}

// GetCompanyUserOk returns a tuple with the CompanyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCompanyUserOk() (*CompanyUserRef, bool) {
	if o == nil || IsNil(o.CompanyUser) {
		return nil, false
	}
	return o.CompanyUser, true
}

// HasCompanyUser returns a boolean if a field has been set.
func (o *User) HasCompanyUser() bool {
	if o != nil && !IsNil(o.CompanyUser) {
		return true
	}

	return false
}

// SetCompanyUser gets a reference to the given CompanyUserRef and assigns it to the CompanyUser field.
func (o *User) SetCompanyUser(v CompanyUserRef) {
	o.CompanyUser = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.AcceptedTermsVersion) {
		toSerialize["accepted_terms_version"] = o.AcceptedTermsVersion
	}
	if !IsNil(o.OauthUserId) {
		toSerialize["oauth_user_id"] = o.OauthUserId
	}
	if !IsNil(o.OauthProviderId) {
		toSerialize["oauth_provider_id"] = o.OauthProviderId
	}
	if !IsNil(o.LanguageId) {
		toSerialize["language_id"] = o.LanguageId
	}
	if !IsNil(o.VerifiedPhoneNumber) {
		toSerialize["verified_phone_number"] = o.VerifiedPhoneNumber
	}
	if !IsNil(o.SmsVerificationCode) {
		toSerialize["sms_verification_code"] = o.SmsVerificationCode
	}
	if !IsNil(o.OauthUserTokenExpiry) {
		toSerialize["oauth_user_token_expiry"] = o.OauthUserTokenExpiry
	}
	if !IsNil(o.HasPassword) {
		toSerialize["has_password"] = o.HasPassword
	}
	if !IsNil(o.LastConfirmedEmailAddress) {
		toSerialize["last_confirmed_email_address"] = o.LastConfirmedEmailAddress
	}
	if !IsNil(o.CustomValue1) {
		toSerialize["custom_value1"] = o.CustomValue1
	}
	if !IsNil(o.CustomValue2) {
		toSerialize["custom_value2"] = o.CustomValue2
	}
	if !IsNil(o.CustomValue3) {
		toSerialize["custom_value3"] = o.CustomValue3
	}
	if !IsNil(o.CustomValue4) {
		toSerialize["custom_value4"] = o.CustomValue4
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Google2faSecret) {
		toSerialize["google_2fa_secret"] = o.Google2faSecret
	}
	if !IsNil(o.CompanyUser) {
		toSerialize["company_user"] = o.CompanyUser
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
