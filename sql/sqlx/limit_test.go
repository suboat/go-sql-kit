package gosql

import "testing"

func TestSQLXLimit_Example1(t *testing.T) {
	example := `{"%l":5,"%s":12,"%p":1}`       // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetMaxLimit(100)                     // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example2(t *testing.T) {
	example := `{"%l":5,"%s":12}`              // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example3(t *testing.T) {
	example := `{"%l":5,"%s":-12,"%p":3}`      // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example4(t *testing.T) {
	example := `{"%l":-1,"%s":2,"%p":2}`       // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetMaxLimit(10)                      // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
