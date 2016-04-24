package gosql

import "encoding/json"

func (o *OrderRoot) ParseJSONString(str string) error {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return err
	}
	return o.Parse(m)
}

func (o *OrderRoot) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return nil
	}
	o.Values = make([]IOrder, 0, len(m))
	for k, v := range m {
		if v == nil {
		} else if IsOrderKey(k) {
			value := &OrderValue{Key: k}
			if err := value.Parse(v); err == nil {
				o.Values = append(o.Values, value)
			}
		}
	}
	return nil
}

func (o *OrderValue) Parse(obj interface{}) error {
	if str, ok := obj.(string); !ok {
		return ErrTypeString
	} else if len(str) == 0 {
		return ErrTypeString
	} else {
		o.Field = str
	}
	return nil
}
