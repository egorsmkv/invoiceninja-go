package openapi

import (
	"encoding/json"
	"fmt"
)

// PaymentType The payment type used to complete this payment
type PaymentType string

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
