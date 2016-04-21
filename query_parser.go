package gosql

import "strings"

func (q *QueryRoot) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return nil
	}
	q.Value = make([]IQuery, 0, len(m))
	for k, v := range m {
		if v == nil {
		} else if IsQueryKey1(k) || IsQueryKey2(k) {
			elem := &QueryElem{anonymous: IsQueryKey2(k), Key: k}
			if err := elem.Parse(v); err == nil {
				q.Value = append(q.Value, elem)
			}
		}
	}
	return nil
}

func (q *QueryElem) Parse(obj interface{}) error {
	if m, ok := obj.(map[string]interface{}); !ok {
		return ErrTypeMap
	} else if m == nil || len(m) == 0 {
		return ErrTypeMap
	} else {
		q.Value = make([]IQuery, 0, len(m))
		for k, v := range m {
			if v == nil {
			} else if IsQueryKey1(k) || IsQueryKey2(k) {
				elem := &QueryElem{anonymous: IsQueryKey2(k), Key: k}
				if err := elem.Parse(v); err == nil {
					q.Value = append(q.Value, elem)
				}
			} else {
				value := &QueryValue{Key: q.Key, Field: k}
				if err := value.Parse(v); err == nil {
					q.Value = append(q.Value, value)
				}
			}
		}
	}
	if q.Value == nil || len(q.Value) == 0 {
		return ErrTypeMap
	}
	return nil
}

func (q *QueryValue) Parse(obj interface{}) error {
	if obj == nil {
		return ErrTypeString
	} else if strings.HasPrefix(q.Field, "%") {
		return ErrTypeString
	}
	switch v := obj.(type) {
	case int:
		q.Value = v
	case int64:
		q.Value = v
	case float32:
		q.Value = v
	case float64:
		q.Value = v
	case string:
		q.Value = v
	default:
		return ErrTypeString
	}
	return nil
}
