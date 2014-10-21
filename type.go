package json

type CheckedValue struct {
	Value   interface{}
	IsValid bool
}

func Check(val interface{}, err error) CheckedValue {
	if err != nil {
		return CheckedValue{
			Value:   nil,
			IsValid: false,
		}
	}

	return CheckedValue{
		Value:   val,
		IsValid: true,
	}
}

func Value(val interface{}) CheckedValue {
	return CheckedValue{
		Value:   val,
		IsValid: true,
	}
}
