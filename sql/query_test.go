package gosql

import "testing"

func TestSQLQuery_Example1(t *testing.T) {
	example := map[string]interface{}{ // 定义Example
		"%eq": map[string]interface{}{
			"t1": "111",
			"t2": 222,
		},
		"%ne": map[string]interface{}{
			"t3": 333,
			"t4": "444",
		},
	}
	order := NewSQLQuery()                  // 初始化
	order.Allow("t1", "t3")                 // 设置关键字过滤规则
	result, err := order.SQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSQLQuery_Example2(t *testing.T) {
	example := map[string]interface{}{
		"%and": map[string]interface{}{
			"%eq": map[string]interface{}{
				"t1": 111,
				"t2": "222",
			},
			"%ne": map[string]interface{}{
				"t3": 333,
				"t4": "444",
			},
		},
	}
	order := NewSQLQuery()
	order.Allow("t1", "t2", "t3")
	result, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSQLQuery_Example3(t *testing.T) {
	example := map[string]interface{}{
		"%and": map[string]interface{}{
			"%eq": map[string]interface{}{
				"t1": "111",
				"t2": 222,
			},
			"%ne": map[string]interface{}{
				"t3": 333,
				"t4": "444",
			},
			"%or": map[string]interface{}{
				"%lt": map[string]interface{}{
					"t11": 1111,
					"t12": "1222",
				},
				"%gte": map[string]interface{}{
					"t13": 1333,
					"t14": 1444,
				},
			},
		},
	}
	order := NewSQLQuery()
	order.Allow("t1", "t2", "t3", "t12", "t13")
	result, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
