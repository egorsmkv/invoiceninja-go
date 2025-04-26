# GetLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Location**](Location.md) |  | [optional] 

## Methods

### NewGetLocations200Response

`func NewGetLocations200Response() *GetLocations200Response`

NewGetLocations200Response instantiates a new GetLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLocations200ResponseWithDefaults

`func NewGetLocations200ResponseWithDefaults() *GetLocations200Response`

NewGetLocations200ResponseWithDefaults instantiates a new GetLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetLocations200Response) GetData() []Location`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLocations200Response) GetDataOk() (*[]Location, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLocations200Response) SetData(v []Location)`

SetData sets Data field to given value.

### HasData

`func (o *GetLocations200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


