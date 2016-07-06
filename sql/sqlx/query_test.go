package gosql

import "testing"

func TestSQLXQuery_ExampleJSON1(t *testing.T) {
	example := `{"%and":{"%eq":{"key1":"A12","key2":"B23"},"%ne":{"key3":"C34","key4":"D45"}}}`
	order := NewSQLXQuery()                            // 初始化
	order.Allow("key1", "key3")                        // 设置关键字过滤规则
	sql, values, err := order.JSONtoSQLString(example) // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLXQuery_ExampleJSON2(t *testing.T) {
	example := `{"%or":{"%lt":{"key1":12,"key2":23},"%gte":{"key3":34,"key4":45}}}`
	order := NewSQLXQuery()                                 // 初始化
	order.Allow("key1", "key3")                             // 设置关键字过滤规则
	sql, values, err := order.JSONtoSQLString(example, "a") // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLXQuery_Example1(t *testing.T) {
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
	order := NewSQLXQuery()                           // 初始化
	order.Allow("t1", "t3")                           // 设置关键字过滤规则
	sql, values, err := order.SQLString(example, "b") // 生成SQL语句
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLXQuery_Example2(t *testing.T) {
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

func TestSQLXQuery_Example3(t *testing.T) {
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
				"%and": map[string]interface{}{
					"%lt": map[string]interface{}{
						"t11": 1111,
						"t12": "1222",
					},
				},
				"%gte": map[string]interface{}{
					"t13": 1333,
					"t14": 1444,
				},
			},
		},
	}
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3", "t11", "t12", "t13", "t14")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLXQuery_Example4(t *testing.T) {
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

func TestSQLXQuery_Example5(t *testing.T) {
	example := map[string]interface{}{
		"%and": map[string]interface{}{
			"%eq": map[string]interface{}{
				"t1": "111",
				"t2": 122,
			},
			"%bt": map[string]interface{}{
				"t3": []interface{}{
					311,
					322,
				},
			},
			"%nbt": map[string]interface{}{
				"t4": []interface{}{
					"433",
					"444",
				},
			},
			"%in": map[string]interface{}{
				"t5": []interface{}{
					511,
					"522",
				},
			},
		},
	}
	order := NewSQLXQuery()
	order.Allow("t1", "t2", "t3", "t4", "t5").Allow("%bt", "%nbt", "%in")
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}

func TestSQLXQuery_ErrorExample1(t *testing.T) {
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
	order.Allow("t1", "t2", "t3", "t12", "t13", "t14").Allow("%like").
		SetMappingFunc("t2", func(key string, value interface{}) (string, interface{}, bool) {
			return "tt2", "2222", true
		}).
		SetMappingFunc("t3", func(key string, value interface{}) (string, interface{}, bool) {
			return key, value, false
		}).
		SetRuleMappingResult("t12", func(key string, value interface{}, method string, alias ...string) (interface{}, bool) {
			if alias != nil && len(alias) != 0 {
				return alias[0] + "t12>1222", true
			}
			return "t12>1222", true
		})
	sql, values, err := order.SQLString(example)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
	t.Logf("%v", values)
}
