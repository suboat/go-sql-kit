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
	order := NewSQLXQuery()                      // 初始化
	order.Allow("t1", "t3")                      // 设置关键字过滤规则
	sql, values, err := order.SQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
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
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3").SetMapping("t1", "ttt111").SetMapping("t2", "ttt222")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
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
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3", "t12", "t13")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLQuery_Example4(t *testing.T) {
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
			"%like": map[string]interface{}{
				"t21": 2111,
				"t22": "2222",
				"t23": 2333,
				"t24": 2444,
			},
		},
	}
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3", "t12", "t13", "t21", "t24").Allow("%like")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLQuery_ErrorExample1(t *testing.T) {
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
					"%lt": map[string]interface{}{
						"t11": 1111,
						"t12": "1222",
					},
				},
			},
			"%like": map[string]interface{}{
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
		},
	}
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3", "t12", "t13").Allow("%like")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}
