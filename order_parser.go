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
	for k, v := range m {
		if v == nil {
		} else if !IsOrderKey(k) {
		} else if strs, ok := v.([]interface{}); ok && len(strs) != 0 {
			if o.Values == nil {
				o.Values = make([]IOrder, 0, len(strs))
			}
			for _, str := range strs {
				value := &OrderValue{}
				if err := value.Parse(str); err == nil {
					o.Values = append(o.Values, value)
				}
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
		switch str[:1] {
		case OrderKeyASC:
			o.Key = OrderKeyASC
			o.Field = str[1:]
		case OrderKeyDESC:
			o.Key = OrderKeyDESC
			o.Field = str[1:]
		default:
			o.Key = OrderKeyASC
			o.Field = str
		}
	}
	return nil
}
