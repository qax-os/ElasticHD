ElasticHD
-----------
[![Build Status](https://travis-ci.org/farmerx/ElasticHD.svg?branch=master)](https://travis-ci.org/farmerx/ElasticHD)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/ElasticHD/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/aurora)](https://goreportcard.com/report/github.com/farmerx/elasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/farmerx/elasticHD/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/elasticHD/elasticHD.svg?label=Release)](https://github.com/farmerx/elasticHD/releases)
> ElasticHD 是一款 ElasticSearch的可视化应用。不依赖ES的插件安装，更便捷；导航栏直接填写对应的ES IP和端口就可以操作Es了。目前支持如下功能：
 * ES 实时搜索
 * ES DashBoard 数据可视化
 * ES Index Template (在线修改、查看、上传）
 * SQL Converts to DSL
 * ES 基本查询文档
 

## Installation

[Precompiled binaries](https://github.com/farmerx/elasticHD/releases) for supported operating systems are available.

## Basic Usage
 * linux and MacOs use ElasticHD 
 ```  
 下载对应的elasticHD版本，unzip xxx_elasticHd_xxx.zip
 修改权限 chmod 0777 ElasticHD
 可指定ip端口运行elastichd ./ElasticHD -p 127.0.0.1:9800 默认 ip和端口也是这个
 ```
 * windows
 ```
 直接下载对应windows版本,解压，双击运行。当然想指定端口的话同linux
 ```

## Es version support
> 测试过elasticsearch 1.5版本到5.2.1的版本都能正常使用。 关于 sql 转化成 dsl　马上会出elasticHD 1.1版本修复一些兼容性错误。

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## ElasticHD SQL Converts to ElasticSearch DSL Usage

### SQL Features Support:

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

### Beyond SQL Features Support：
- [x] ES TopHits
- [x] ES date_histogram
- [x] ES STATS
- [x] ES RANGE
- [x] ES DATE_RANGE



*Improvement : now the query DSL is much more flat*


### SQL Usage
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


## 源码编译
```
# 需要go环境
go get https://github.com/farmerx/ElasticHD
# 进入到应用目录下
cd ElasticHD
npm install
# build vue 源码
npm run build
# 进入到服务端程序目录
cd ./main
# 使用statik 压缩编译好的程序(github上的一款go应用)
statik -src=../dist
# go build
GO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o elasticHD.exe github.com/elasticHD/main
```
## ElasticHD应用页面
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120044177.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120114338.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120147401.png)
![image](https://raw.githubusercontent.com/farmerx/ElasticHD/master/WX20170605-105148.png)

## Todo
...


## Licenses

This program is under the terms of the MIT License. See [LICENSE](https://github.com/farmerx/elasticHD/blob/master/LICENSE) for the full license text.


