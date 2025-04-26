package openapi

import (
	"encoding/json"
)

// checks if the BankIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankIntegration{}

// BankIntegration struct for BankIntegration
type BankIntegration struct {
	// The bank integration hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The providers bank name
	ProviderBankName *string `json:"provider_bank_name,omitempty"`
	// The bank account id
	BankAccountId *int32 `json:"bank_account_id,omitempty"`
	// The name of the account
	BankAccountName *string `json:"bank_account_name,omitempty"`
	// The account number
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// The status of the bank account
	BankAccountStatus *string `json:"bank_account_status,omitempty"`
	// The type of account
	BankAccountType *string `json:"bank_account_type,omitempty"`
	// The current bank balance if available
	Balance *float32 `json:"balance,omitempty"`
	// iso_3166_3 code
	Currency *string `json:"currency,omitempty"`
}

// NewBankIntegration instantiates a new BankIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankIntegration() *BankIntegration {
	this := BankIntegration{}
	return &this
}

// NewBankIntegrationWithDefaults instantiates a new BankIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankIntegrationWithDefaults() *BankIntegration {
	this := BankIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankIntegration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankIntegration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankIntegration) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BankIntegration) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BankIntegration) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BankIntegration) SetUserId(v string) {
	o.UserId = &v
}

// GetProviderBankName returns the ProviderBankName field value if set, zero value otherwise.
func (o *BankIntegration) GetProviderBankName() string {
	if o == nil || IsNil(o.ProviderBankName) {
		var ret string
		return ret
	}
	return *o.ProviderBankName
}

// GetProviderBankNameOk returns a tuple with the ProviderBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetProviderBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderBankName) {
		return nil, false
	}
	return o.ProviderBankName, true
}

// HasProviderBankName returns a boolean if a field has been set.
func (o *BankIntegration) HasProviderBankName() bool {
	if o != nil && !IsNil(o.ProviderBankName) {
		return true
	}

	return false
}

// SetProviderBankName gets a reference to the given string and assigns it to the ProviderBankName field.
func (o *BankIntegration) SetProviderBankName(v string) {
	o.ProviderBankName = &v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountId() int32 {
	if o == nil || IsNil(o.BankAccountId) {
		var ret int32
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given int32 and assigns it to the BankAccountId field.
func (o *BankIntegration) SetBankAccountId(v int32) {
	o.BankAccountId = &v
}

// GetBankAccountName returns the BankAccountName field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountName() string {
	if o == nil || IsNil(o.BankAccountName) {
		var ret string
		return ret
	}
	return *o.BankAccountName
}

// GetBankAccountNameOk returns a tuple with the BankAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountName) {
		return nil, false
	}
	return o.BankAccountName, true
}

// HasBankAccountName returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountName() bool {
	if o != nil && !IsNil(o.BankAccountName) {
		return true
	}

	return false
}

// SetBankAccountName gets a reference to the given string and assigns it to the BankAccountName field.
func (o *BankIntegration) SetBankAccountName(v string) {
	o.BankAccountName = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountNumber() string {
	if o == nil || IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountNumber() bool {
	if o != nil && !IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *BankIntegration) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankAccountStatus returns the BankAccountStatus field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountStatus() string {
	if o == nil || IsNil(o.BankAccountStatus) {
		var ret string
		return ret
	}
	return *o.BankAccountStatus
}

// GetBankAccountStatusOk returns a tuple with the BankAccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountStatus) {
		return nil, false
	}
	return o.BankAccountStatus, true
}

// HasBankAccountStatus returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountStatus() bool {
	if o != nil && !IsNil(o.BankAccountStatus) {
		return true
	}

	return false
}

// SetBankAccountStatus gets a reference to the given string and assigns it to the BankAccountStatus field.
func (o *BankIntegration) SetBankAccountStatus(v string) {
	o.BankAccountStatus = &v
}

// GetBankAccountType returns the BankAccountType field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountType() string {
	if o == nil || IsNil(o.BankAccountType) {
		var ret string
		return ret
	}
	return *o.BankAccountType
}

// GetBankAccountTypeOk returns a tuple with the BankAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountType) {
		return nil, false
	}
	return o.BankAccountType, true
}

// HasBankAccountType returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountType() bool {
	if o != nil && !IsNil(o.BankAccountType) {
		return true
	}

	return false
}

// SetBankAccountType gets a reference to the given string and assigns it to the BankAccountType field.
func (o *BankIntegration) SetBankAccountType(v string) {
	o.BankAccountType = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *BankIntegration) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *BankIntegration) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *BankIntegration) SetBalance(v float32) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BankIntegration) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BankIntegration) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BankIntegration) SetCurrency(v string) {
	o.Currency = &v
}

func (o BankIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankIntegration) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ProviderBankName) {
		toSerialize["provider_bank_name"] = o.ProviderBankName
	}
	if !IsNil(o.BankAccountId) {
		toSerialize["bank_account_id"] = o.BankAccountId
	}
	if !IsNil(o.BankAccountName) {
		toSerialize["bank_account_name"] = o.BankAccountName
	}
	if !IsNil(o.BankAccountNumber) {
		toSerialize["bank_account_number"] = o.BankAccountNumber
	}
	if !IsNil(o.BankAccountStatus) {
		toSerialize["bank_account_status"] = o.BankAccountStatus
	}
	if !IsNil(o.BankAccountType) {
		toSerialize["bank_account_type"] = o.BankAccountType
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableBankIntegration struct {
	value *BankIntegration
	isSet bool
}

func (v NullableBankIntegration) Get() *BankIntegration {
	return v.value
}

func (v *NullableBankIntegration) Set(val *BankIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableBankIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableBankIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankIntegration(val *BankIntegration) *NullableBankIntegration {
	return &NullableBankIntegration{value: val, isSet: true}
}

func (v NullableBankIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
