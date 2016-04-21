package gosql

import "testing"

func TestSQLOrder_StringJSONtoSQL(t *testing.T) {
	example := []string{"key1", "+key2", "+key3", "-key4", "-key5"}
	order := NewSQLOrder()
	order.Allow("key1", "key2", "key5")
	result, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
