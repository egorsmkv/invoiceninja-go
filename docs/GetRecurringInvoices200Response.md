# GetRecurringInvoices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RecurringInvoice**](RecurringInvoice.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetRecurringInvoices200Response

`func NewGetRecurringInvoices200Response() *GetRecurringInvoices200Response`

NewGetRecurringInvoices200Response instantiates a new GetRecurringInvoices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecurringInvoices200ResponseWithDefaults

`func NewGetRecurringInvoices200ResponseWithDefaults() *GetRecurringInvoices200Response`

NewGetRecurringInvoices200ResponseWithDefaults instantiates a new GetRecurringInvoices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRecurringInvoices200Response) GetData() []RecurringInvoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRecurringInvoices200Response) GetDataOk() (*[]RecurringInvoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRecurringInvoices200Response) SetData(v []RecurringInvoice)`

SetData sets Data field to given value.

### HasData

`func (o *GetRecurringInvoices200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetRecurringInvoices200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRecurringInvoices200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRecurringInvoices200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetRecurringInvoices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


