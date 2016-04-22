package gosql

import "encoding/json"

func (l *LimitRoot) ParseJSONString(str string) error {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return err
	}
	return l.Parse(m)
}

func (l *LimitRoot) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return nil
	}
	l.Values = make([]ILimit, 0, len(m))
	for k, v := range m {
		if v == nil {
		} else if IsLimitKey(k) {
			value := &LimitValue{Key: k}
			if err := value.Parse(v); err == nil {
				l.Values = append(l.Values, value)
			}
		}
	}
	return nil
}

func (l *LimitValue) Parse(obj interface{}) error {
	if i, ok := obj.(int); !ok {
		return ErrTypeInt
	} else if l.Value = i; !l.IsLimited() {
		return ErrTypeValue
	}
	return nil
}
