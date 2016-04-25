package gosql

import "testing"

func TestSQLParser_ExampleJSON1(t *testing.T) {
	example := `[{"%and":{"%eq":{"key1":"A12","key2":"B23"},"%ne":{"key3":"C34","key4":"D45"}}},` +
		`{"%o":["+key1","-key2"]},` +
		`{"%l":5,"%s":12,"%p":1}]`
	parser := NewSQLParser().InitALL()      // 初始化
	parser.GetQuery().Allow("key1", "key3") // 设置关键字过滤规则
	parser.GetOrder().Allow("key1", "key2") // 设置关键字过滤规则
	parser.GetLimit().SetMaxLimit(100)      // 设置关键字过滤规则
	err := parser.ParseJSONString(example)  // 解析SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parser.GetQuery().String()) // 生成SQL语句
	t.Log(parser.GetOrder().String()) // 生成SQL语句
	t.Log(parser.GetLimit().String()) // 生成SQL语句
}
