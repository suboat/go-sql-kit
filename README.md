# go-sql-kit

# v0.6

## This project is still in development :)

## Overview

* JSON格式
* 只需通过字符串(string)，即可实现**条件筛选(WHERE)**/**结果排序(ORDER BY)**/**结果分页(LIMIT)**等常用功能
* 提供快速便捷方案，对接开发前端(JS)模块(Developing...)，实现上述**筛选**/**排序**/**分页**功能

## Usage

```
go get -u github.com/suboat/go-sql-kit
```

## Documents

当前规则均基于JSON格式

1. Query(**条件筛选(WHERE)**)
2. Order(**结果排序(ORDER BY)**)
3. Limit(**结果分页(LIMIT)**)
4. Rule

### 1. Query (`./query.go`)

#### 关键字

```golang
QueryKeyAnd string = "%and" // 与
QueryKeyOr         = "%or"  // 或

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

* 关键字"%and"和"%or"需继续包含关键字

#### 实例说明

* JSON实例：{"%and":{"%eq":{"key1":"A12"}}} 或 {"%eq":{"key1":"A12"}}
* 结果描述： key1 == "A12"
* JSON实例：{"%and":{"%eq":{"key1":"A12","key2":"B23"},"%ne":{"key3":"C34","key4":"D45"}}}
* 结果描述： (key1 == "A12" && key2 == "B23") && (key3 != "C34" && key4 != "D45")
* JSON实例：{"%or":{"%lt":{"key1":12,"key2":23},"%gte":{"key3":34,"key4":45}}}
* 结果描述： (key1 < 12 && key2 < 23) || (key3 >= 34 && key4 >= 45)

### 2. Order (`./order.go`)

#### 关键字

```golang
OrderKeyASC  string = "+" // 正序
OrderKeyDESC        = "-" // 反序
```

* 正序缺省可以不加关键字

#### 实例说明

* JSON实例：["key1", "+key2", "+key3", "-key4", "-key5"]
* 结果描述： 正序("key1", "key2", "key3")，反序("key4", "key5")
* 正序：例如对字段"key1"正向排序，可写为"+key1"，也可以"key1"
* 反序：例如对字段"key4"反向排序，需写为"-key4"

### 3. Limit (`./limit.go`)

#### 关键字

```golang
LimitKey string = "%" // 限制分隔符
```

* 用法描述: limit%skip%page
* 允许缺省值，如: limit%skip, limit%skip%, limit%%page, %%page, %%

#### 实例说明

* JSON实例："5%3%2"
* 结果描述1： 忽略最前面的13个值，返回最多5个值
* 结果描述2： 忽略最前面的3个，并返回第3页的值，每页最多5个值

### 4. Rule (`./rule.go`)

## TODO

* 开发前端(JS)模块
* 不断完善文档说明
