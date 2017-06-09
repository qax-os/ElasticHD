ElasticSQL
-----------
[![Build Status](https://travis-ci.org/farmerx/elasticsql.svg?branch=master)](https://travis-ci.org/farmerx/elasticsql)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/elasticsql)
[![Coverage Status](https://coveralls.io/repos/github/farmerx/elasticsql/badge.svg?branch=master)](https://coveralls.io/github/farmerx/elasticsql?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/farmerx/elasticsql)](https://goreportcard.com/report/github.com/farmerx/elasticsql)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/farmerx/elasticsql/blob/master/LICENSE)

> ElasticSQL package converts SQL to ElasticSearch DSL

## SQL Features Support:

- [x] SQL Select
- [x] SQL Where
- [x] SQL Order BySQL
- [x] SQL Group By
- [x] SQL AND & OR
- [x] SQL Like & NOT Like
- [x] SQL COUNT distinct
- [x] SQL In & Not In
- [x] SQL Between
- [x] SQL avg()、count(*), count(field), min(field), max(field)

## Beyond SQL Features Support：
- [x] ES TopHits
- [x] ES date_histogram
- [x] ES STATS
- [x] ES RANGE
- [x] ES DATE_RANGE



*Improvement : now the query DSL is much more flat*


## SQL Usage
Query
```
select * from test where a=1 and b="c" and create_time between '2015-01-01T00:00:00+0800' and '2016-01-01T00:00:00+0800' and process_id > 1 order by id desc limit 100,10
```
Aggregation
```
select avg(age),min(age),max(age),count(student),count(distinct student) from test group by grade,class limit 10
```
Beyond SQL
 * range age group 20-25,25-30,30-35,35-40
	```
	SELECT COUNT(age) FROM bank GROUP BY range(age, 20,25,30,35,40)
	```
 * range date group by your config
 	```
	SELECT online FROM online GROUP BY date_range(field="insert_time",format="yyyy-MM-dd" ,"2014-08-18","2014-08-17","now-8d","now-7d","now-6d","now")
	```
 * range date group by day

	```
	select * from test group by date_histogram(field="changeTime",interval="1h",format="yyyy-MM-dd HH:mm:ss")
	```
 * stats
 	```
	 SELECT online FROM online group by stats(field="grade")
	```
 * topHits
 	```
	  select top_hits(field="class", hitssort="age:desc", taglimit = "10", hitslimit = "1", _source="name,age,class,gender") from school
	```


## PKG Usage
-------------

> go get github.com/farmerx/elasticsql

Demo :
```go
package main

import (
    "fmt"
    "github.com/farmerx/elasticsql"
)

var sql = `
select * from test where a=1 and b="c" and create_time between '2015-01-01T00:00:00+0800' and '2016-01-01T00:00:00+0800' and process_id > 1 order by id desc limit 100,10
`

var sql2= `
  select avg(age),min(age),max(age),count(student) from test group by class limit 10
`
var sql3= `
  select * from test group by class,student limit 10
`
var sql4 = `
  select * from test group by date_histogram(field="changeTime",interval="1h",format="yyyy-MM-dd HH:mm:ss")
`


func main() {
    esql := elasticsql.NewElasticSQL(eslasticsql.InitOptions{})
    table, dsl, err := esql.SQLConvert(sql)
	fmt.Println(table, dsl, err)
}

```

## Licenses

This program is under the terms of the MIT License. See [LICENSE](https://github.com/farmerx/elasticsql/blob/master/LICENSE) for the full license text.

