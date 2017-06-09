<template>
  <section id="help">
     <div v-html="compiledMarkdown" style="background-color:white;"></div>
  </section>
</template>

<script>
import { mapGetters } from 'vuex'
import Marked from 'marked'
export default {
  name: 'help',
  data: function () {
    return {
      content: `
ElasticHD
-----------
[![Build Status](https://travis-ci.org/farmerx/ElasticHD.svg?branch=master)](https://travis-ci.org/farmerx/ElasticHD)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/ElasticHD/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/aurora)](https://goreportcard.com/report/github.com/farmerx/elasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/farmerx/elasticHD/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/elasticHD/elasticHD.svg?label=Release)](https://github.com/farmerx/elasticHD/releases)

>ElasticHD 是一款 ElasticSearch的可视化应用。不依赖ES的插件安装，更便捷；导航栏直接填写对应的ES IP和端口就可以操作Es了。目前支持如下功能：
 * ES 实时搜索
 * ES DashBoard 数据可视化
 * ES Index Template (在线修改、查看、上传）
 * SQL Converts to DSL
 * ES 基本查询文档

### Basic Usage
 * linux and MacOs use ElasticHD
\`\`\`
 下载对应的elasticHD版本，unzip xxx_elasticHd_xxx.zip
 修改权限 chmod 0777 ElasticHD
 可指定ip端口运行elastichd ./ElasticHD -p 127.0.0.1:9800 默认ip和端口也是这个
 \`\`\`
 * windows
\`\`\`
 直接下载对应windows版本,解压，双击运行。当然想指定端口的话同linux
 \`\`\`

### ElasticHD SQL Converts to ElasticSearch DSL Usage
#### SQL Features Support:
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

#### Beyond SQL Features Support：
- [x] ES TopHits
- [x] ES date_histogram
- [x] ES STATS
- [x] ES RANGE
- [x] ES DATE_RANGE

#### SQL Usage
Query
 \`\`\`
SELECT * FROM test WHERE a=1 and b="c" and create_time between '2015-01-01T00:00:00+0800' and '2016-01-01T00:00:00+0800' and process_id > 1 order by id desc limit 100,10
 \`\`\`
Aggregation
 \`\`\`
SELECT avg(age),min(age),max(age),count(student),count(distinct student) FROM test group by grade,class limit 10
 \`\`\`
Beyond SQL
 * range age group 20-25,25-30,30-35,35-40
    \`\`\`
SELECT COUNT(age) FROM bank GROUP BY range(age, 20,25,30,35,40)
    \`\`\`
 * range date group by your config

    \`\`\`
SELECT online FROM online GROUP BY date_range(field="insert_time",format="yyyy-MM-dd" ,"2014-08-18","2014-08-17","now-8d","now-7d","now-6d","now")
    \`\`\`
 * range date group by day

    \`\`\`
SELECT * FROM test GROUP BY date_histogram(field="changeTime",interval="1h",format="yyyy-MM-dd HH:mm:ss")
    \`\`\`
 * stats

    \`\`\`
SELECT online FROM online GROUP BY stats(field="grade")
    \`\`\`
 * topHits

    \`\`\`
SELECT top_hits(field="class", hitssort="age:desc", taglimit = "10", hitslimit = "1", _source="name,age,class,gender") FROM school
    \`\`\`

*Improvement : now the query DSL is much more flat*

### Contributing

> Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

### ElasticSearch 基本查询语法

### 基本搜索
 \`\`\`
{
    "query": {
        "bool": {
            "must": [
                {
                    "match_all": {}
                }
            ]
        }
    },
    "from": 0,
    "size": 1
}
 \`\`\`
### Group BY
 \`\`\`
{
    "query": {
        "bool": {
            "must": [
                {
                    "match_all": {}
                }
            ]
        }
    },
    "from": 0,
    "size": 0,
    "aggregations": {
        "mid": {
            "aggregations": {
                "terminal": {
                    "terms": {
                        "field": "terminal",
                        "size": 0
                    }
                }
            },
            "terms": {
                "field": "mid",
                "size": "1"
            }
        }
    }
}
 \`\`\`

### Distinct Count
 \`\`\`
{
    "query": {
        "bool": {
            "must": [
                {
                    "match_all": {}
                }
            ]
        }
    },
    "from": 0,
    "size": 0,
    "aggregations": {
        "COUNT(distinct (mid))": {
            "cardinality": {
                "field": "(mid)"
            }
        }
    }
}
 \`\`\`
### 全文搜索
 \`\`\`
{
    "query" : {
        "query_string" : {"query" : "name:rcx"}
    }
}
 \`\`\`
### match查询
 \`\`\`
{
    "query": {
        "match": {
            "title": "crime and punishment"
        }
    }
}
 \`\`\`
### 通配符查询
 \`\`\`
{
    "query": {
        "wildcard": {
             "title": "cr?me"
        }
    }
}
 \`\`\`
### 范围查询
 \`\`\`
{
    "query": {
        "range": {
             "year": {
                  "gte" :1890,
                  "lte":1900
              }
        }
    }
}
 \`\`\`
### 正则表达式查询
 \`\`\`
{
    "query": {
        "regexp": {
             "title": {
                  "value" :"cr.m[ae]",
                  "boost":10.0
              }
        }
    }
}
 \`\`\`
### 布尔查询
 \`\`\`
{
    "query": {
        "bool": {
            "must": {
                "term": {
                    "title": "crime"
                }
            },
            "should": {
                "range": {
                    "year": {
                        "from": 1900,
                        "to": 2000
                    }
                }
            },
            "must_not": {
                "term": {
                    "otitle": "nothing"
                }
            }
        }
    }
}
 \`\`\`
`
    }
  },
  computed: {
    ...mapGetters([
    ]),
    compiledMarkdown: function () {
      return Marked(this.content, { sanitize: true })
    }
  },
  created: function () {
  },
  mounted () {
  }
}
</script>

<style>
</style>

