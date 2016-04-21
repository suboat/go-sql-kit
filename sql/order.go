package gosql

import . "github.com/suboat/go-sql-kit"

func OrderStringJSONtoSQL(str string) (string, error) {
	return "", nil
}

type SQLOrderRoot struct {
	OrderRoot
}

func (o *SQLOrderRoot) String() (statement string) {
	if o.Value == nil || len(o.Value) == 0 {
		return
	}
	for _, iv := range o.Value {
		if v, ok := iv.(OrderValue); ok {
			v.Field
		}
	}

	return ""
}
