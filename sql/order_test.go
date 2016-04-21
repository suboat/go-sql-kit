package gosql

import "testing"

func TestSQLOrder_StringJSONtoSQL(t *testing.T) {
	example := []string{"key1", "+key2", "+key3", "-key4", "-key5"} // 定义Example，注意前缀设定
	order := NewSQLOrder()                                          // 初始化
	order.Allow("key1", "key2", "key5")                             // 设置关键字过滤规则
	result, err := order.SQLString(example)                         // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
