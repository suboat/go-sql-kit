package gosql

import "testing"

func TestSQLXLimit_Example1(t *testing.T) {
	example := `5%12%1`                        // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetDefaultLimit(10).SetMaxLimit(100) // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example2(t *testing.T) {
	example := `2`                             // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetDefaultLimit(10)                  // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example3(t *testing.T) {
	example := `%-12%2`                        // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetDefaultLimit(10)                  // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}

func TestSQLXLimit_Example4(t *testing.T) {
	example := `-2%0%2`                        // 定义Example
	limit := NewSQLXLimit()                    // 初始化
	limit.SetMaxLimit(10)                      // 设置限制规则
	sql, err := limit.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
