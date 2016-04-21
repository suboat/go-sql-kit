package gosql

import "encoding/json"

func (o *OrderRoot) ParseJSON(str string) error {
	var set []string
	if err := json.Unmarshal([]byte(str), set); err != nil {
		return err
	} else if len(set) == 0 {
		o.Value = []IOrder{}
		return nil
	}
	o.Value = make([]IOrder, 0, len(set))
	for _, s := range set {
		if len(s) == 0 {
			continue
		}
		var v OrderValue
		switch s[:1] {
		case OrderKey_asc:
			v.ASC = true
			v.Field = s[1:]
		case OrderKey_desc:
			v.ASC = false
			v.Field = s[1:]
		default:
			v.ASC = true
			v.Field = s
		}
		if len(v.Field) != 0 {
			o.Value = append(o.Value, v)
		}
	}
	return nil
}
