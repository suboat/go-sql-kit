# go-sql-kit

# v0.3

## This project is still in development :)

## Overview

* 只需通过字符串(string)即可实现**值排序(ORDER)**和**条件筛选(WHERE)**两大常用功能
* JSON格式
* 容易对接前段(JS)模块，实现上述**排序**和**筛选**功能

## Document

* 当前规则均基于JSON格式

### Order

#### 关键字

```golang
OrderKeyASC  string = "+" // 正序
OrderKeyDESC        = "-" // 反序
```


#### 说明

* JSON实例：[{"key1", "+key2", "+key3", "-key4", "-key5"}]
* 正序：例如对字段"key1"正向排序，可写为"+key1"，也可以"key1"
* 反序：例如对字段"key4"反向排序，需写为"-key4"

### Query

### Rule

## TODO
