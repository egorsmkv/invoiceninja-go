# ClientGatewayToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the client gateway token | [optional] 
**ClientId** | Pointer to **string** | The hashed_id of the client | [optional] 
**Token** | Pointer to **string** | The payment token | [optional] 
**RoutingNumber** | Pointer to **string** | THe bank account routing number | [optional] 
**CompanyGatewayId** | Pointer to **string** | The hashed id of the company gateway | [optional] 
**IsDefault** | Pointer to **bool** | Flag determining if the token is the default payment method | [optional] 

## Methods

### NewClientGatewayToken

`func NewClientGatewayToken() *ClientGatewayToken`

NewClientGatewayToken instantiates a new ClientGatewayToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGatewayTokenWithDefaults

`func NewClientGatewayTokenWithDefaults() *ClientGatewayToken`

NewClientGatewayTokenWithDefaults instantiates a new ClientGatewayToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientGatewayToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientGatewayToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientGatewayToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientGatewayToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *ClientGatewayToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientGatewayToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientGatewayToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientGatewayToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetToken

`func (o *ClientGatewayToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientGatewayToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientGatewayToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientGatewayToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *ClientGatewayToken) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *ClientGatewayToken) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *ClientGatewayToken) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *ClientGatewayToken) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetCompanyGatewayId

`func (o *ClientGatewayToken) GetCompanyGatewayId() string`

GetCompanyGatewayId returns the CompanyGatewayId field if non-nil, zero value otherwise.

### GetCompanyGatewayIdOk

`func (o *ClientGatewayToken) GetCompanyGatewayIdOk() (*string, bool)`

GetCompanyGatewayIdOk returns a tuple with the CompanyGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGatewayId

`func (o *ClientGatewayToken) SetCompanyGatewayId(v string)`

SetCompanyGatewayId sets CompanyGatewayId field to given value.

### HasCompanyGatewayId

`func (o *ClientGatewayToken) HasCompanyGatewayId() bool`

HasCompanyGatewayId returns a boolean if a field has been set.

### GetIsDefault

`func (o *ClientGatewayToken) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ClientGatewayToken) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ClientGatewayToken) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ClientGatewayToken) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


