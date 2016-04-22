package gosql

import (
	"strconv"
	"strings"
)

func (l *LimitRoot) ParseJSONString(str string) error {
	return l.Parse(str)
}

func (l *LimitRoot) Parse(str string) error {
	v := new(LimitValue)
	l.Value = v
	if len(str) == 0 {
	} else if strs := strings.Split(str, LimitKey); len(strs) == 0 {
	} else {
		for idx, s := range strs {
			if len(s) == 0 {
				continue
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			switch idx {
			case 0:
				v.Limit = i
			case 1:
				v.Skip = i
			case 2:
				v.Page = i
			}
		}
	}
	return nil
}
