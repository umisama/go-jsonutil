package json

import (
	"encoding/json"
)

type ObjectT map[string]CheckedValue
type ArrayT []CheckedValue

func UnmarshalObject(req []byte)(ObjectT, error) {
	buf := make(map[string]interface{})
	err := json.Unmarshal(req, &buf)
	if err != nil {
		return nil, err
	}

	ret := make(ObjectT)
	for k, v := range buf {
		ret[k] = Value(v)
	}

	return ret, nil
}

func UnmarshalArray(req []byte)(ArrayT, error) {
	buf := make([]interface{}, 0)
	err := json.Unmarshal(req, &buf)
	if err != nil {
		return nil, err
	}

	ret := make(ArrayT, 0)
	for _, v := range buf {
		ret = append(ret, Value(v))
	}

	return ret, nil
}

func (m ObjectT) Marshal() []byte {
	return marshal(m.clean())
}

func (m ObjectT) clean() map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v.isValid() {
			switch vt := v.value().(type) {
			case ObjectT:
				ret[k] = vt.clean()
			case ArrayT:
				ret[k] = vt.clean()
			default:
				ret[k] = vt
			}
		}
	}
	return ret
}

func (m ArrayT) Marshal() []byte {
	return marshal(m.clean())
}

func (m ArrayT) clean() []interface{} {
	ret := make([]interface{}, 0, len(m))
	for _, v := range m {
		if v.isValid() {
			switch vt := v.value().(type) {
			case ObjectT:
				ret = append(ret, vt.clean())
			case ArrayT:
				ret = append(ret, vt.clean())
			default:
				ret = append(ret, vt)
			}
		}
	}
	return ret
}

func marshal(i interface{}) []byte {
	buf, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return buf
}

func (m *ObjectT) isValid() bool {
	return true
}

func (m *ObjectT) value() interface{} {
	return m
}

func (m *ArrayT) isValid() bool {
	return true
}

func (m *ArrayT) value() interface{} {
	return m
}
