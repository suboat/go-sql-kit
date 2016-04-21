package gosql

import "testing"

func TestSQLOrder_StringJSONtoSQL(t *testing.T) {
	order := NewSQLOrder()
	result, err := order.JSONStringsToSQL([]string{"key", "+key_asc", "-key_desc"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
