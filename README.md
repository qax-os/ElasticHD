ElasticHD
-----------
[![Build Status](https://travis-ci.org/farmerx/ElasticHD.svg?branch=master)](https://travis-ci.org/farmerx/ElasticHD)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/ElasticHD/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/aurora)](https://goreportcard.com/report/github.com/farmerx/elasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/farmerx/elasticHD/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/elasticHD/elasticHD.svg?label=Release)](https://github.com/farmerx/elasticHD/releases)
[![GitHub issues](https://img.shields.io/github/issues/farmerx/ElasticHD.svg)](https://github.com/farmerx/ElasticHD/issues)
> ElasticHD is a ElasticSearch visual management tool. It does not require any software. It works in your web browser, allowing you to manage and monitor your ElasticSearch clusters from anywhere at any time. Built on responsive CSS design, ElasticHD adjusts itself to any screen size on any device.The following functions are supported：
 * ES Real time data search and query
 * ES Dashboard data visualization
 * ES Indices Management
 * Managing Type Mappings (在线修改、查看、上传）
 * SQL Converts to Elasticsearch DSL
 * ES 基本查询文档
 * Device Friendly
 
## 支持权限认证
> 支持有权限认证的ElasticSearch服务器，url格式：http://user:password@host:port

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
## No Software to Install
ElasticHD does not require any software. It works in your web browser, allowing you to manage and monitor your ElasticSearch clusters from anywhere at any time. Built on responsive CSS design, ElasticHD adjusts itself to any screen size on any device.
## Es version support
> Compatible with all ES versions
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
git Clone https://github.com/farmerx/ElasticHD
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
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(1).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(2).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(3).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(4).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(5).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(6).png)
![image](https://github.com/farmerx/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(7).png)

## Todo
* More plugins support
* The indices list supports search, better sorting, detailed viewing, and more
* Program logo design
* Monitoring information collection

## Licenses

This program is under the terms of the MIT License. See [LICENSE](https://github.com/farmerx/elasticHD/blob/master/LICENSE) for the full license text.


