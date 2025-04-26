# FeesAndLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLimit** | Pointer to **string** | The minimum amount accepted for this gateway | [optional] 
**MaxLimit** | Pointer to **string** | The maximum amount accepted for this gateway | [optional] 
**FeeAmount** | Pointer to **float32** | The gateway fee amount | [optional] 
**FeePercent** | Pointer to **float32** | The gateway fee percentage | [optional] 
**FeeTaxName1** | Pointer to **string** | Fee tax name | [optional] 
**FeeTaxName2** | Pointer to **string** | Fee tax name | [optional] 
**FeeTaxName3** | Pointer to **string** | Fee tax name | [optional] 
**FeeTaxRate1** | Pointer to **float32** | The tax rate | [optional] 
**FeeTaxRate2** | Pointer to **float32** | The tax rate | [optional] 
**FeeTaxRate3** | Pointer to **float32** | The tax rate | [optional] 
**FeeCap** | Pointer to **float32** | If set the fee amount will be no higher than this amount | [optional] 
**AdjustFeePercent** | Pointer to **bool** | Adjusts the fee to match the exact gateway fee. | [optional] 

## Methods

### NewFeesAndLimits

`func NewFeesAndLimits() *FeesAndLimits`

NewFeesAndLimits instantiates a new FeesAndLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeesAndLimitsWithDefaults

`func NewFeesAndLimitsWithDefaults() *FeesAndLimits`

NewFeesAndLimitsWithDefaults instantiates a new FeesAndLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLimit

`func (o *FeesAndLimits) GetMinLimit() string`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *FeesAndLimits) GetMinLimitOk() (*string, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *FeesAndLimits) SetMinLimit(v string)`

SetMinLimit sets MinLimit field to given value.

### HasMinLimit

`func (o *FeesAndLimits) HasMinLimit() bool`

HasMinLimit returns a boolean if a field has been set.

### GetMaxLimit

`func (o *FeesAndLimits) GetMaxLimit() string`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *FeesAndLimits) GetMaxLimitOk() (*string, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *FeesAndLimits) SetMaxLimit(v string)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *FeesAndLimits) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.

### GetFeeAmount

`func (o *FeesAndLimits) GetFeeAmount() float32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *FeesAndLimits) GetFeeAmountOk() (*float32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *FeesAndLimits) SetFeeAmount(v float32)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *FeesAndLimits) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeePercent

`func (o *FeesAndLimits) GetFeePercent() float32`

GetFeePercent returns the FeePercent field if non-nil, zero value otherwise.

### GetFeePercentOk

`func (o *FeesAndLimits) GetFeePercentOk() (*float32, bool)`

GetFeePercentOk returns a tuple with the FeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercent

`func (o *FeesAndLimits) SetFeePercent(v float32)`

SetFeePercent sets FeePercent field to given value.

### HasFeePercent

`func (o *FeesAndLimits) HasFeePercent() bool`

HasFeePercent returns a boolean if a field has been set.

### GetFeeTaxName1

`func (o *FeesAndLimits) GetFeeTaxName1() string`

GetFeeTaxName1 returns the FeeTaxName1 field if non-nil, zero value otherwise.

### GetFeeTaxName1Ok

`func (o *FeesAndLimits) GetFeeTaxName1Ok() (*string, bool)`

GetFeeTaxName1Ok returns a tuple with the FeeTaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxName1

`func (o *FeesAndLimits) SetFeeTaxName1(v string)`

SetFeeTaxName1 sets FeeTaxName1 field to given value.

### HasFeeTaxName1

`func (o *FeesAndLimits) HasFeeTaxName1() bool`

HasFeeTaxName1 returns a boolean if a field has been set.

### GetFeeTaxName2

`func (o *FeesAndLimits) GetFeeTaxName2() string`

GetFeeTaxName2 returns the FeeTaxName2 field if non-nil, zero value otherwise.

### GetFeeTaxName2Ok

`func (o *FeesAndLimits) GetFeeTaxName2Ok() (*string, bool)`

GetFeeTaxName2Ok returns a tuple with the FeeTaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxName2

`func (o *FeesAndLimits) SetFeeTaxName2(v string)`

SetFeeTaxName2 sets FeeTaxName2 field to given value.

### HasFeeTaxName2

`func (o *FeesAndLimits) HasFeeTaxName2() bool`

HasFeeTaxName2 returns a boolean if a field has been set.

### GetFeeTaxName3

`func (o *FeesAndLimits) GetFeeTaxName3() string`

GetFeeTaxName3 returns the FeeTaxName3 field if non-nil, zero value otherwise.

### GetFeeTaxName3Ok

`func (o *FeesAndLimits) GetFeeTaxName3Ok() (*string, bool)`

GetFeeTaxName3Ok returns a tuple with the FeeTaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxName3

`func (o *FeesAndLimits) SetFeeTaxName3(v string)`

SetFeeTaxName3 sets FeeTaxName3 field to given value.

### HasFeeTaxName3

`func (o *FeesAndLimits) HasFeeTaxName3() bool`

HasFeeTaxName3 returns a boolean if a field has been set.

### GetFeeTaxRate1

`func (o *FeesAndLimits) GetFeeTaxRate1() float32`

GetFeeTaxRate1 returns the FeeTaxRate1 field if non-nil, zero value otherwise.

### GetFeeTaxRate1Ok

`func (o *FeesAndLimits) GetFeeTaxRate1Ok() (*float32, bool)`

GetFeeTaxRate1Ok returns a tuple with the FeeTaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxRate1

`func (o *FeesAndLimits) SetFeeTaxRate1(v float32)`

SetFeeTaxRate1 sets FeeTaxRate1 field to given value.

### HasFeeTaxRate1

`func (o *FeesAndLimits) HasFeeTaxRate1() bool`

HasFeeTaxRate1 returns a boolean if a field has been set.

### GetFeeTaxRate2

`func (o *FeesAndLimits) GetFeeTaxRate2() float32`

GetFeeTaxRate2 returns the FeeTaxRate2 field if non-nil, zero value otherwise.

### GetFeeTaxRate2Ok

`func (o *FeesAndLimits) GetFeeTaxRate2Ok() (*float32, bool)`

GetFeeTaxRate2Ok returns a tuple with the FeeTaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxRate2

`func (o *FeesAndLimits) SetFeeTaxRate2(v float32)`

SetFeeTaxRate2 sets FeeTaxRate2 field to given value.

### HasFeeTaxRate2

`func (o *FeesAndLimits) HasFeeTaxRate2() bool`

HasFeeTaxRate2 returns a boolean if a field has been set.

### GetFeeTaxRate3

`func (o *FeesAndLimits) GetFeeTaxRate3() float32`

GetFeeTaxRate3 returns the FeeTaxRate3 field if non-nil, zero value otherwise.

### GetFeeTaxRate3Ok

`func (o *FeesAndLimits) GetFeeTaxRate3Ok() (*float32, bool)`

GetFeeTaxRate3Ok returns a tuple with the FeeTaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTaxRate3

`func (o *FeesAndLimits) SetFeeTaxRate3(v float32)`

SetFeeTaxRate3 sets FeeTaxRate3 field to given value.

### HasFeeTaxRate3

`func (o *FeesAndLimits) HasFeeTaxRate3() bool`

HasFeeTaxRate3 returns a boolean if a field has been set.

### GetFeeCap

`func (o *FeesAndLimits) GetFeeCap() float32`

GetFeeCap returns the FeeCap field if non-nil, zero value otherwise.

### GetFeeCapOk

`func (o *FeesAndLimits) GetFeeCapOk() (*float32, bool)`

GetFeeCapOk returns a tuple with the FeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCap

`func (o *FeesAndLimits) SetFeeCap(v float32)`

SetFeeCap sets FeeCap field to given value.

### HasFeeCap

`func (o *FeesAndLimits) HasFeeCap() bool`

HasFeeCap returns a boolean if a field has been set.

### GetAdjustFeePercent

`func (o *FeesAndLimits) GetAdjustFeePercent() bool`

GetAdjustFeePercent returns the AdjustFeePercent field if non-nil, zero value otherwise.

### GetAdjustFeePercentOk

`func (o *FeesAndLimits) GetAdjustFeePercentOk() (*bool, bool)`

GetAdjustFeePercentOk returns a tuple with the AdjustFeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustFeePercent

`func (o *FeesAndLimits) SetAdjustFeePercent(v bool)`

SetAdjustFeePercent sets AdjustFeePercent field to given value.

### HasAdjustFeePercent

`func (o *FeesAndLimits) HasAdjustFeePercent() bool`

HasAdjustFeePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


