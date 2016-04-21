package gosql

import "encoding/json"

func (o *OrderRoot) ParseJSONString(str string) error {
	var strs []string
	if err := json.Unmarshal([]byte(str), strs); err != nil {
		return err
	}
	return o.Parse(strs)
}

func (o *OrderRoot) Parse(strs []string) error {
	if strs == nil || len(strs) == 0 {
		o.Value = []IOrder{}
		return nil
	}
	o.Value = make([]IOrder, 0, len(strs))
	for _, s := range strs {
		if len(s) == 0 {
			continue
		}
		var v OrderValue
		switch s[:1] {
		case OrderKeyASC:
			v.ASC = true
			v.Field = s[1:]
		case OrderKeyDESC:
			v.ASC = false
			v.Field = s[1:]
		default:
			v.ASC = true
			v.Field = s
		}
		if len(v.Field) != 0 {
			o.Value = append(o.Value, &v)
		}
	}
	return nil
}
