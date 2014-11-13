package json

type CheckedValue interface {
	value() interface{}
	isValid() bool
}

type checkedValue struct {
	Value   interface{}
	IsValid bool
}

func Check(val interface{}, err error) CheckedValue {
	if err != nil {
		return checkedValue{
			Value:   nil,
			IsValid: false,
		}
	}

	return checkedValue{
		Value:   val,
		IsValid: true,
	}
}

func Value(val interface{}) CheckedValue {
	return checkedValue {
		Value:   val,
		IsValid: true,
	}
}

func (m checkedValue) value() interface{} {
	return m.Value
}

func (m checkedValue) isValid() bool {
	return m.IsValid
}
