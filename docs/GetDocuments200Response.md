# GetDocuments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetDocuments200Response

`func NewGetDocuments200Response() *GetDocuments200Response`

NewGetDocuments200Response instantiates a new GetDocuments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocuments200ResponseWithDefaults

`func NewGetDocuments200ResponseWithDefaults() *GetDocuments200Response`

NewGetDocuments200ResponseWithDefaults instantiates a new GetDocuments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDocuments200Response) GetData() []Document`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDocuments200Response) GetDataOk() (*[]Document, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDocuments200Response) SetData(v []Document)`

SetData sets Data field to given value.

### HasData

`func (o *GetDocuments200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetDocuments200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDocuments200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDocuments200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDocuments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


