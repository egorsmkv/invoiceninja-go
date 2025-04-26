package openapi

import (
	"encoding/json"
	"fmt"
)

// PaymentType The payment type used to complete this payment
type PaymentType string

// List of PaymentType
const (
	_1  PaymentType = "1"
	_2  PaymentType = "2"
	_3  PaymentType = "3"
	_4  PaymentType = "4"
	_5  PaymentType = "5"
	_6  PaymentType = "6"
	_7  PaymentType = "7"
	_8  PaymentType = "8"
	_9  PaymentType = "9"
	_10 PaymentType = "10"
	_11 PaymentType = "11"
	_12 PaymentType = "12"
	_13 PaymentType = "13"
	_14 PaymentType = "14"
	_15 PaymentType = "15"
	_16 PaymentType = "16"
	_17 PaymentType = "17"
	_18 PaymentType = "18"
	_19 PaymentType = "19"
	_20 PaymentType = "20"
	_21 PaymentType = "21"
	_22 PaymentType = "22"
	_23 PaymentType = "23"
	_24 PaymentType = "24"
	_25 PaymentType = "25"
	_26 PaymentType = "26"
	_27 PaymentType = "27"
	_28 PaymentType = "28"
	_29 PaymentType = "29"
	_30 PaymentType = "30"
	_31 PaymentType = "31"
	_32 PaymentType = "32"
	_33 PaymentType = "33"
)

// All allowed values of PaymentType enum
var AllowedPaymentTypeEnumValues = []PaymentType{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"11",
	"12",
	"13",
	"14",
	"15",
	"16",
	"17",
	"18",
	"19",
	"20",
	"21",
	"22",
	"23",
	"24",
	"25",
	"26",
	"27",
	"28",
	"29",
	"30",
	"31",
	"32",
	"33",
}

func (v *PaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentType(value)
	for _, existing := range AllowedPaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentType", value)
}

// NewPaymentTypeFromValue returns a pointer to a valid PaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentTypeFromValue(v string) (*PaymentType, error) {
	ev := PaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentType: valid values are %v", v, AllowedPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentType) IsValid() bool {
	for _, existing := range AllowedPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentType value
func (v PaymentType) Ptr() *PaymentType {
	return &v
}

type NullablePaymentType struct {
	value *PaymentType
	isSet bool
}

func (v NullablePaymentType) Get() *PaymentType {
	return v.value
}

func (v *NullablePaymentType) Set(val *PaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentType(val *PaymentType) *NullablePaymentType {
	return &NullablePaymentType{value: val, isSet: true}
}

func (v NullablePaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
