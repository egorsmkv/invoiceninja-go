# CompanyGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the company gateway | [optional] 
**GatewayKey** | Pointer to **string** | The gateway key (hash) | [optional] 
**AcceptedCreditCards** | Pointer to **int32** | Bitmask representation of cards | [optional] 
**RequireBillingAddress** | Pointer to **bool** | Determines if the the billing address is required prior to payment. | [optional] 
**RequireShippingAddress** | Pointer to **bool** | Determines if the the billing address is required prior to payment. | [optional] 
**Config** | Pointer to **string** | The configuration map for the gateway | [optional] 
**UpdateDetails** | Pointer to **bool** | Determines if the client details should be updated. | [optional] 
**FeesAndLimits** | Pointer to [**[]FeesAndLimits**](FeesAndLimits.md) | A mapped collection of the fees and limits for the configured gateway | [optional] 

## Methods

### NewCompanyGateway

`func NewCompanyGateway() *CompanyGateway`

NewCompanyGateway instantiates a new CompanyGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyGatewayWithDefaults

`func NewCompanyGatewayWithDefaults() *CompanyGateway`

NewCompanyGatewayWithDefaults instantiates a new CompanyGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGatewayKey

`func (o *CompanyGateway) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *CompanyGateway) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *CompanyGateway) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *CompanyGateway) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetAcceptedCreditCards

`func (o *CompanyGateway) GetAcceptedCreditCards() int32`

GetAcceptedCreditCards returns the AcceptedCreditCards field if non-nil, zero value otherwise.

### GetAcceptedCreditCardsOk

`func (o *CompanyGateway) GetAcceptedCreditCardsOk() (*int32, bool)`

GetAcceptedCreditCardsOk returns a tuple with the AcceptedCreditCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCreditCards

`func (o *CompanyGateway) SetAcceptedCreditCards(v int32)`

SetAcceptedCreditCards sets AcceptedCreditCards field to given value.

### HasAcceptedCreditCards

`func (o *CompanyGateway) HasAcceptedCreditCards() bool`

HasAcceptedCreditCards returns a boolean if a field has been set.

### GetRequireBillingAddress

`func (o *CompanyGateway) GetRequireBillingAddress() bool`

GetRequireBillingAddress returns the RequireBillingAddress field if non-nil, zero value otherwise.

### GetRequireBillingAddressOk

`func (o *CompanyGateway) GetRequireBillingAddressOk() (*bool, bool)`

GetRequireBillingAddressOk returns a tuple with the RequireBillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBillingAddress

`func (o *CompanyGateway) SetRequireBillingAddress(v bool)`

SetRequireBillingAddress sets RequireBillingAddress field to given value.

### HasRequireBillingAddress

`func (o *CompanyGateway) HasRequireBillingAddress() bool`

HasRequireBillingAddress returns a boolean if a field has been set.

### GetRequireShippingAddress

`func (o *CompanyGateway) GetRequireShippingAddress() bool`

GetRequireShippingAddress returns the RequireShippingAddress field if non-nil, zero value otherwise.

### GetRequireShippingAddressOk

`func (o *CompanyGateway) GetRequireShippingAddressOk() (*bool, bool)`

GetRequireShippingAddressOk returns a tuple with the RequireShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireShippingAddress

`func (o *CompanyGateway) SetRequireShippingAddress(v bool)`

SetRequireShippingAddress sets RequireShippingAddress field to given value.

### HasRequireShippingAddress

`func (o *CompanyGateway) HasRequireShippingAddress() bool`

HasRequireShippingAddress returns a boolean if a field has been set.

### GetConfig

`func (o *CompanyGateway) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CompanyGateway) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CompanyGateway) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CompanyGateway) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUpdateDetails

`func (o *CompanyGateway) GetUpdateDetails() bool`

GetUpdateDetails returns the UpdateDetails field if non-nil, zero value otherwise.

### GetUpdateDetailsOk

`func (o *CompanyGateway) GetUpdateDetailsOk() (*bool, bool)`

GetUpdateDetailsOk returns a tuple with the UpdateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDetails

`func (o *CompanyGateway) SetUpdateDetails(v bool)`

SetUpdateDetails sets UpdateDetails field to given value.

### HasUpdateDetails

`func (o *CompanyGateway) HasUpdateDetails() bool`

HasUpdateDetails returns a boolean if a field has been set.

### GetFeesAndLimits

`func (o *CompanyGateway) GetFeesAndLimits() []FeesAndLimits`

GetFeesAndLimits returns the FeesAndLimits field if non-nil, zero value otherwise.

### GetFeesAndLimitsOk

`func (o *CompanyGateway) GetFeesAndLimitsOk() (*[]FeesAndLimits, bool)`

GetFeesAndLimitsOk returns a tuple with the FeesAndLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAndLimits

`func (o *CompanyGateway) SetFeesAndLimits(v []FeesAndLimits)`

SetFeesAndLimits sets FeesAndLimits field to given value.

### HasFeesAndLimits

`func (o *CompanyGateway) HasFeesAndLimits() bool`

HasFeesAndLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


