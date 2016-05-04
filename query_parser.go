package gosql

import (
	"encoding/json"
	"strings"
)

func (o *QueryRoot) ParseJSONString(str string) error {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return err
	}
	return o.Parse(m)
}

func (q *QueryRoot) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return nil
	}
	for k, v := range m {
		if v == nil {
		} else if IsQueryKey(k) {
			if q.Values == nil {
				q.Values = make([]IQuery, 0, len(m))
			}
			elem := &QueryElem{anonymous: IsQueryAnonymousKey(k), Key: k}
			if err := elem.Parse(v); err == nil {
				q.Values = append(q.Values, elem)
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
		if q.Values == nil {
			q.Values = make([]IQuery, 0, len(m))
		}
		for k, v := range m {
			if v == nil {
			} else if IsQueryKey(k) {
				if !q.IsAnonymous() {
					elem := &QueryElem{anonymous: IsQueryAnonymousKey(k), Key: k}
					if err := elem.Parse(v); err == nil {
						q.Values = append(q.Values, elem)
					}
				}
			} else {
				value := &QueryValue{Key: q.Key, Field: k}
				if err := value.Parse(v); err == nil {
					q.Values = append(q.Values, value)
				}
			}
		}
	}
	if q.Values == nil || len(q.Values) == 0 {
		return ErrTypeMap
	}
	return nil
}

func (q *QueryValue) Parse(obj interface{}) error {
	if obj == nil {
		return ErrTypeNil
	} else if strings.HasPrefix(q.Field, "%") {
		return ErrTypeString
	}
	switch v := obj.(type) {
	case int:
		q.Value = v
	case int8:
		q.Value = v
	case int16:
		q.Value = v
	case int32:
		q.Value = v
	case int64:
		q.Value = v
	case float32:
		q.Value = v
	case float64:
		q.Value = v
	case string:
		q.Value = v
	case []interface{}:
		q.Value = v
	default:
		return ErrTypeString
	}
	return nil
}
