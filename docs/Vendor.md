# Vendor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the vendor. This is a unique identifier for the vendor. | [optional] [readonly] 
**UserId** | Pointer to **string** | The hashed id of the user who created the vendor. This is a unique identifier for the user. | [optional] 
**AssignedUserId** | Pointer to **string** | The hashed id of the assigned user to this vendor. This is a unique identifier for the user. | [optional] 
**Contacts** | Pointer to [**[]VendorContact**](VendorContact.md) | An array of contacts associated with the vendor. | [optional] 
**Name** | Pointer to **string** | The name of the vendor. | [optional] 
**Classification** | Pointer to **string** | The classification of the vendor. | [optional] 
**Website** | Pointer to **string** | The website of the vendor. | [optional] 
**PrivateNotes** | Pointer to **string** | The private notes of the vendor. These notes are only visible to users with appropriate permissions. | [optional] 
**IndustryId** | Pointer to **string** | The industry id of the vendor. This is a unique identifier for the industry. | [optional] 
**SizeId** | Pointer to **string** | The size id of the vendor. This is a unique identifier for the size of the vendor. | [optional] 
**Address1** | Pointer to **string** | The first line of the vendor&#39;s address. | [optional] 
**Address2** | Pointer to **string** | The second line of the vendor&#39;s address. | [optional] 
**City** | Pointer to **string** | The city of the vendor&#39;s address. | [optional] 
**State** | Pointer to **string** | The state of the vendor&#39;s address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the vendor&#39;s address. | [optional] 
**Phone** | Pointer to **string** | The phone number of the vendor. | [optional] 
**CountryId** | Pointer to **string** | The country id of the vendor. This is a unique identifier for the country. | [optional] 
**CurrencyId** | Pointer to **string** | The currency id of the vendor. This is a unique identifier for the currency. | [optional] 
**CustomValue1** | Pointer to **string** | The value of the first custom field for the vendor. | [optional] 
**CustomValue2** | Pointer to **string** | The value of the second custom field for the vendor. | [optional] 
**CustomValue3** | Pointer to **string** | The value of the third custom field for the vendor. | [optional] 
**CustomValue4** | Pointer to **string** | The value of the fourth custom field for the vendor. | [optional] 
**VatNumber** | Pointer to **string** | The VAT number of the vendor. | [optional] 
**IdNumber** | Pointer to **string** | The ID number of the vendor. | [optional] 
**Number** | Pointer to **string** | The number of the vendor | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the vendor has been deleted | [optional] 
**LanguageId** | Pointer to **string** | The language id of the vendor. This is a unique identifier for the language. | [optional] 
**VendorHash** | Pointer to **string** | The vendor hash of the vendor. This is a unique identifier for the vendor. | [optional] [readonly] 
**TransactionName** | Pointer to **string** | The transaction name of the vendor. | [optional] 
**LastLogin** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the vendor. | [optional] [readonly] 

## Methods

### NewVendor

`func NewVendor() *Vendor`

NewVendor instantiates a new Vendor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorWithDefaults

`func NewVendorWithDefaults() *Vendor`

NewVendorWithDefaults instantiates a new Vendor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vendor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vendor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vendor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vendor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Vendor) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Vendor) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Vendor) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Vendor) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Vendor) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Vendor) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Vendor) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Vendor) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetContacts

`func (o *Vendor) GetContacts() []VendorContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Vendor) GetContactsOk() (*[]VendorContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Vendor) SetContacts(v []VendorContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Vendor) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetName

`func (o *Vendor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vendor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vendor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vendor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClassification

`func (o *Vendor) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Vendor) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Vendor) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Vendor) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetWebsite

`func (o *Vendor) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Vendor) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Vendor) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Vendor) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *Vendor) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *Vendor) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *Vendor) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *Vendor) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetIndustryId

`func (o *Vendor) GetIndustryId() string`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Vendor) GetIndustryIdOk() (*string, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Vendor) SetIndustryId(v string)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Vendor) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetSizeId

`func (o *Vendor) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *Vendor) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *Vendor) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *Vendor) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### GetAddress1

`func (o *Vendor) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *Vendor) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *Vendor) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *Vendor) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *Vendor) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *Vendor) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *Vendor) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *Vendor) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Vendor) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Vendor) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Vendor) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Vendor) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Vendor) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Vendor) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Vendor) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Vendor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *Vendor) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Vendor) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Vendor) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Vendor) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPhone

`func (o *Vendor) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Vendor) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Vendor) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Vendor) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCountryId

`func (o *Vendor) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Vendor) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Vendor) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Vendor) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCurrencyId

`func (o *Vendor) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *Vendor) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *Vendor) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *Vendor) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Vendor) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Vendor) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Vendor) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Vendor) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Vendor) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Vendor) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Vendor) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Vendor) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Vendor) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Vendor) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Vendor) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Vendor) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Vendor) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Vendor) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Vendor) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Vendor) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetVatNumber

`func (o *Vendor) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Vendor) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Vendor) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Vendor) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetIdNumber

`func (o *Vendor) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *Vendor) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *Vendor) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *Vendor) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetNumber

`func (o *Vendor) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Vendor) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Vendor) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Vendor) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Vendor) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Vendor) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Vendor) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Vendor) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLanguageId

`func (o *Vendor) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *Vendor) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *Vendor) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *Vendor) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetVendorHash

`func (o *Vendor) GetVendorHash() string`

GetVendorHash returns the VendorHash field if non-nil, zero value otherwise.

### GetVendorHashOk

`func (o *Vendor) GetVendorHashOk() (*string, bool)`

GetVendorHashOk returns a tuple with the VendorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorHash

`func (o *Vendor) SetVendorHash(v string)`

SetVendorHash sets VendorHash field to given value.

### HasVendorHash

`func (o *Vendor) HasVendorHash() bool`

HasVendorHash returns a boolean if a field has been set.

### GetTransactionName

`func (o *Vendor) GetTransactionName() string`

GetTransactionName returns the TransactionName field if non-nil, zero value otherwise.

### GetTransactionNameOk

`func (o *Vendor) GetTransactionNameOk() (*string, bool)`

GetTransactionNameOk returns a tuple with the TransactionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionName

`func (o *Vendor) SetTransactionName(v string)`

SetTransactionName sets TransactionName field to given value.

### HasTransactionName

`func (o *Vendor) HasTransactionName() bool`

HasTransactionName returns a boolean if a field has been set.

### GetLastLogin

`func (o *Vendor) GetLastLogin() float32`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Vendor) GetLastLoginOk() (*float32, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Vendor) SetLastLogin(v float32)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Vendor) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Vendor) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vendor) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vendor) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Vendor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Vendor) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Vendor) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Vendor) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Vendor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *Vendor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Vendor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Vendor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Vendor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


