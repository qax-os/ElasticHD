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
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/elasticHD/main)
> ElasticHD 是一款 ElasticSearch的可视化应用。不依赖ES的插件安装，更便捷；导航栏直接填写对应的ES IP和端口就可以操作Es了。目前支持如下功能：
 * ES 实时搜索
 * ES DashBoard 数据可视化
 * ES Index Template (在线修改、查看、上传）
 * SQL Converts to DSL
 * ES 基本查询文档

## Installation

[Precompiled binaries](https://github.com/farmerx/elasticHD/releases) for supported operating systems are available.

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## ElasticSearch 基本查询语法

### 基本搜索
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

### Group BY
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


### Distinct Count
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
### 全文搜索
{
    "query" : {
        "query_string" : {"query" : "name:rcx"}
    }
}
### match查询
{
    "query": {
        "match": {
            "title": "crime and punishment"
        }
    }
}
### 通配符查询
{
    "query": {
        "wildcard": {
             "title": "cr?me"
        }
    }
}
### 范围查询
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
### 正则表达式查询
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
### 布尔查询
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

