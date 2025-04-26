package openapi

import (
	"encoding/json"
)

// checks if the BankTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransaction{}

// BankTransaction struct for BankTransaction
type BankTransaction struct {
	// The bank integration hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The id of the transaction rule
	TransactionId *int32 `json:"transaction_id,omitempty"`
	// The transaction amount
	Amount *float32 `json:"amount,omitempty"`
	// The currency ID of the currency
	CurrencyId *string `json:"currency_id,omitempty"`
	// The account type
	AccountType *string `json:"account_type,omitempty"`
	// The description of the transaction
	Description *string `json:"description,omitempty"`
	// The category id
	CategoryId *int32 `json:"category_id,omitempty"`
	// The category description
	CategoryType *string `json:"category_type,omitempty"`
	// Either CREDIT or DEBIT
	BaseType *string `json:"base_type,omitempty"`
	// The date of the transaction
	Date *string `json:"date,omitempty"`
	// The ID number of the bank account
	BankAccountId *int32 `json:"bank_account_id,omitempty"`
}

// NewBankTransaction instantiates a new BankTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransaction() *BankTransaction {
	this := BankTransaction{}
	return &this
}

// NewBankTransactionWithDefaults instantiates a new BankTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransactionWithDefaults() *BankTransaction {
	this := BankTransaction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankTransaction) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BankTransaction) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BankTransaction) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BankTransaction) SetUserId(v string) {
	o.UserId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BankTransaction) GetTransactionId() int32 {
	if o == nil || IsNil(o.TransactionId) {
		var ret int32
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetTransactionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BankTransaction) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given int32 and assigns it to the TransactionId field.
func (o *BankTransaction) SetTransactionId(v int32) {
	o.TransactionId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BankTransaction) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BankTransaction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BankTransaction) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *BankTransaction) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *BankTransaction) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *BankTransaction) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *BankTransaction) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *BankTransaction) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *BankTransaction) SetAccountType(v string) {
	o.AccountType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BankTransaction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BankTransaction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BankTransaction) SetDescription(v string) {
	o.Description = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *BankTransaction) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *BankTransaction) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *BankTransaction) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCategoryType returns the CategoryType field value if set, zero value otherwise.
func (o *BankTransaction) GetCategoryType() string {
	if o == nil || IsNil(o.CategoryType) {
		var ret string
		return ret
	}
	return *o.CategoryType
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetCategoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryType) {
		return nil, false
	}
	return o.CategoryType, true
}

// HasCategoryType returns a boolean if a field has been set.
func (o *BankTransaction) HasCategoryType() bool {
	if o != nil && !IsNil(o.CategoryType) {
		return true
	}

	return false
}

// SetCategoryType gets a reference to the given string and assigns it to the CategoryType field.
func (o *BankTransaction) SetCategoryType(v string) {
	o.CategoryType = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *BankTransaction) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *BankTransaction) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *BankTransaction) SetBaseType(v string) {
	o.BaseType = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BankTransaction) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BankTransaction) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *BankTransaction) SetDate(v string) {
	o.Date = &v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *BankTransaction) GetBankAccountId() int32 {
	if o == nil || IsNil(o.BankAccountId) {
		var ret int32
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetBankAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *BankTransaction) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given int32 and assigns it to the BankAccountId field.
func (o *BankTransaction) SetBankAccountId(v int32) {
	o.BankAccountId = &v
}

func (o BankTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransaction) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CategoryType) {
		toSerialize["category_type"] = o.CategoryType
	}
	if !IsNil(o.BaseType) {
		toSerialize["base_type"] = o.BaseType
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.BankAccountId) {
		toSerialize["bank_account_id"] = o.BankAccountId
	}
	return toSerialize, nil
}

type NullableBankTransaction struct {
	value *BankTransaction
	isSet bool
}

func (v NullableBankTransaction) Get() *BankTransaction {
	return v.value
}

func (v *NullableBankTransaction) Set(val *BankTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransaction(val *BankTransaction) *NullableBankTransaction {
	return &NullableBankTransaction{value: val, isSet: true}
}

func (v NullableBankTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
