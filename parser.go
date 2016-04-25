package gosql

import "encoding/json"

type IParser interface {
	ParseJSONString(string) error
	Parse(map[string]interface{}) error
}

type Parser struct {
	values []IParser
}

func (p *Parser) Add(v IParser) {
	if v != nil {
		p.values = append(p.values, v)
	}
}

func (p *Parser) ParseJSONString(str string) error {
	var objs []interface{}
	if err := json.Unmarshal([]byte(str), &objs); err != nil {
		return err
	} else if len(objs) == 0 {
		return nil
	}
	for _, obj := range objs {
		if obj == nil {
		} else if m, ok := obj.(map[string]interface{}); ok {
			p.Parse(m)
		}
	}
	return nil
}

func (p *Parser) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
	} else if p.values == nil || len(p.values) == 0 {
	} else {
		for _, v := range p.values {
			if v == nil {
			} else if err := v.Parse(m); err != nil {
				return err
			}
		}
	}
	return nil
}
