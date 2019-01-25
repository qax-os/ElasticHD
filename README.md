ElasticHD
-----------
[![Build Status](https://travis-ci.org/360EntSecGroup-Skylar/ElasticHD.svg?branch=master)](https://travis-ci.org/360EntSecGroup-Skylar/ElasticHD)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/360EntSecGroup-Skylar/ElasticHD/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/aurora)](https://goreportcard.com/report/github.com/360EntSecGroup-Skylar/elasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/360EntSecGroup-Skylar/elasticHD/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/360EntSecGroup-Skylar/ElasticHD.svg?label=Release)](https://github.com/360EntSecGroup-Skylar/ElasticHD/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/360EntSecGroup-Skylar/ElasticHD.svg)](https://github.com/360EntSecGroup-Skylar/ElasticHD/issues)
> ElasticHD is a ElasticSearch visual management tool. It does not require any software. It works in your web browser, allowing you to manage and monitor your ElasticSearch clusters from anywhere at any time. Built on responsive CSS design, ElasticHD adjusts itself to any screen size on any device.The following functions are supported：
 * ES Real time data search and query
 * ES Dashboard data visualization
 * ES Indices Management
 * Managing Type Mappings
 * SQL Converts to Elasticsearch DSL
 * Device Friendly
 
 

## ElasticHD Application Pages

![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard.png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(2).png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(3).png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(4).png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(5).png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(6).png)
![image](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/Elastic%20HD%20Dashboard%20(7).png)
 
## Authentication
> ElasticSearch server that supports privileged authentication, URL format：http://user:password@host:port

## Installation

[Precompiled binaries]( https://github.com/360EntSecGroup-Skylar/ElasticHD/releases/) for supported operating systems are available.

## Basic Usage

 * Linux and MacOS use ElasticHD 
 ```
 Step1: Download the corresponding elasticHD version，unzip xxx_elasticHd_xxx.zip
 Step2: chmod 0777 ElasticHD
 Step3: exec elastichd ./ElasticHD -p 127.0.0.1:9800 
 ```
 * windows
 ```
 Step1: Download the corresponding elasticHD version，Double click zip package to unzip
 Step2: exec elastichd ./ElasticHD -p 127.0.0.1:9800 
 ```
 
## Standalone Executable 

> ElasticHD does not require any software. It works in your web browser, allowing you to manage and monitor your ElasticSearch clusters from anywhere at any time. Built on responsive CSS design, ElasticHD adjusts itself to any screen size on any device.

## Es version support

> Compatible with all ES versions

## Contributing

> Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

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



* Improvement : now the query DSL is much more flat*


### SQL Usage

Query
```
SELECT * FROM test WHERE a=1 AND b="c" AND create_time BETWEEN '2015-01-01T00:00:00+0800' AND '2016-01-01T00:00:00+0800' AND process_id > 1 ORDER BY id DESC LIMIT 100,10
```
Aggregation
```
SELECT avg(age), min(age), max(age), count(student), count(distinct student) FROM test GROUP BY grade,class LIMIT 10
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
    SELECT * FROM test GROUP BY date_histogram(field="changeTime",interval="1h",format="yyyy-MM-dd HH:mm:ss")
    ```
 * stats
    ```
    SELECT online FROM online GROUP BY stats(field="grade")
    ```
 * topHits
    ```
    SELECT top_hits(field="class", hitssort="age:desc", taglimit = "10", hitslimit = "1", _source="name,age,class,gender") FROM school
    ```


## Source code compilation

```
git clone https://github.com/360EntSecGroup-Skylar/ElasticHD
cd ElasticHD
npm install
npm run build
cd ./main
statik -src=../dist
# go build
GO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o elasticHD.exe github.com/elasticHD/main
```

## Docker Quick Start:

> Image link: [docker images](https://hub.docker.com/r/containerize/elastichd/)

* Make Docker Images
docker build -t elastichd:latest .

* Docker Usage:

```
docker run -p 9200:9200 -d --name elasticsearch elasticsearch
docker run -p 9800:9800 -d --link elasticsearch:demo containerize/elastichd
Open http://localhost:9800 in Browser
Connect with http://demo:9200

```


## Todo
* More plugins support
* The indices list supports search, better sorting, detailed viewing, and more
* Program logo design
* Monitoring information collection

## Licenses

This program is under the terms of the MIT License. See [LICENSE](https://github.com/360EntSecGroup-Skylar/ElasticHD/blob/master/LICENSE) for the full license text.
