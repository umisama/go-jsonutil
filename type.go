package json

type CheckedValue interface {
	Value() interface{}
	IsValid() bool
}

type checkedValue struct {
	value   interface{}
	isValid bool
}

func Check(val interface{}, err error) CheckedValue {
	if err != nil {
		return checkedValue{
			value:   nil,
			isValid: false,
		}
	}

	return checkedValue{
		value:   val,
		isValid: true,
	}
}

func Value(val interface{}) CheckedValue {
	return checkedValue {
		value:   val,
		isValid: true,
	}
}

func (m checkedValue) Value() interface{} {
	return m.value
}

func (m checkedValue) IsValid() bool {
	return m.isValid
}
