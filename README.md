# go-sql-kit

# v0.4

## This project is still in development :)

## Overview

* 只需通过字符串(string)即可实现**结果排序(ORDER BY)**和**条件筛选(WHERE)**两大常用功能
* JSON格式
* TODO: 提供快速便捷方案，对接开发前端(JS)模块，实现上述**排序**和**筛选**功能

## Documents

* 当前规则均基于JSON格式
* Order(**结果排序(ORDER BY)**)
* Query(**条件筛选(WHERE)**)
* Rule

### 1. Order (`./order.go`)

#### 关键字

```golang
OrderKeyASC  string = "+" // 正序
OrderKeyDESC        = "-" // 反序
```

#### 实例说明

* JSON实例：["key1", "+key2", "+key3", "-key4", "-key5"]
* 正序：例如对字段"key1"正向排序，可写为"+key1"，也可以"key1"
* 反序：例如对字段"key4"反向排序，需写为"-key4"

### 2. Query (`./query.go`)

#### 关键字

```golang
QueryKeyAnd string = "%and" // 与
QueryKeyOr         = "%or"  // 或
```

```golang
QueryKeyEq      string = "%eq"   // 等于
QueryKeyNe             = "%ne"   // 不等于
QueryKeyLt             = "%lt"   // 小于
QueryKeyLte            = "%lte"  // 小于等于
QueryKeyGt             = "%gt"   // 大于
QueryKeyGte            = "%gte"  // 大于等于
QueryKeyLike           = "%like" // 模糊搜索
QueryKeyIn             = "%in"   // TODO: 暂时不支持
QueryKeyBetween        = "%bt"   // TODO: 暂时不支持
```

#### 实例说明

* JSON实例：{"%and":{"%eq":{"key1":"A12","key2":"B23"},"%ne":{"key3":"C34","key4":"D45"}}}
* 条件描述： (key1 == "A12" && key2 == "B23") && (key3 != "C34" && key4 != "D45")
* JSON实例：{"%or":{"%lt":{"key1":12,"key2":23},"%gte":{"key3":34,"key4":45}}}
* 条件描述： (key1 < 12 && key2 < 23) || (key3 >= 34 && key4 >= 45)

### 3. Rule (`./rule.go`)

## TODO
