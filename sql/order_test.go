package gosql

import "testing"

func TestSQLOrder_ExampleJSON1(t *testing.T) {
	example := `{"%asc":"key1","%desc":"key2"}` // 定义Example
	order := NewSQLOrder()                      // 初始化
	order.Allow("key1", "key2")                 // 设置关键字过滤规则
	sql, err := order.JSONtoSQLString(example)  // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLOrder_ExampleJSON2(t *testing.T) {
	example := `{"%asc":"key1","%desc":"key2","%desc":"key5"}` // 定义Example，注意前缀设定
	order := NewSQLOrder()                                     // 初始化
	order.Allow("key1", "key2", "key5")                        // 设置关键字过滤规则
	sql, err := order.JSONtoSQLString(example)                 // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
