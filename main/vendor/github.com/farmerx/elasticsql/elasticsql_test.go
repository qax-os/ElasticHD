package elasticsql

import (
	"fmt"
	"testing"
)

var esql *ElasticSQL

func init() {
	esql = NewElasticSQL(InitOptions{})
}

func Test_SQLConver(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`select * from test where a="b" and c ="2" or d="1" and age between 2 and 3 group by mid,gid order by mid desc `)
	fmt.Println(table, dsl, err)
}

// select count(distinct mid) from test
// SELECT * FROM bank GROUP BY range(pid, 20,25,30,35,40) limit 20
func Test_SQLConverGroup(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`SELECT online FROM online GROUP BY date_range(field="insert_time",format="yyyy-MM-dd" ,"2014-08-18","2014-08-17","now-8d","now-7d","now-6d","now")`)
	fmt.Println(table, dsl, err)
}

func Test_SQLConverLike(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`SELECT online FROM online Where  a like 'x'`)
	fmt.Println(table, dsl, err)
}

func Test_SQLConverQueryString(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`SELECT online FROM online Where query_string = 'sip:"xxx" AND dip:"xxx"'`)
	fmt.Println(table, dsl, err)
}

func Test_SQLConverStats(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`SELECT online FROM online group by stats(field="pid")`)
	fmt.Println(table, dsl, err)
}

// top_hits(field="class", hitssort="age:desc", taglimit = "10", hitslimit = "1", _source="name,age,class")
func Test_SQLConvertsTopHits(t *testing.T) {
	table, dsl, err := esql.SQLConvert(`SELECT top_hits(field="mid") FROM online `)
	fmt.Println(table, dsl, err)
}
