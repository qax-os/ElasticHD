package elasticsql

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/xwb1989/sqlparser"
)

// InitOptions ...
type InitOptions struct {
}

// ElasticSQL ..
type ElasticSQL struct {
}

// NewElasticSQL ...
func NewElasticSQL(options InitOptions) *ElasticSQL {
	esql := new(ElasticSQL)
	return esql
}

// SQLConvert sql convert to elasticsearch dsl
func (esql *ElasticSQL) SQLConvert(sql string) (table string, dsl string, err error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return ``, ``, err
	}
	switch v := stmt.(type) {
	case *sqlparser.Select:
		return handleSelect(v)
	default:
		return ``, ``, errors.New(`不支持此类查询`)
	}
}

// handleSelect 处理select 类型的查询
func handleSelect(Select *sqlparser.Select) (table string, dsl string, err error) {
	var rootParent sqlparser.BoolExpr
	var defaultQueryMapStr = `{"bool" : {"must": [{"match_all" : {}}]}}`
	var queryMapStr string
	// if select where include conditions
	if Select.Where != nil {
		queryMapStr, err = handleSelectWhere(&Select.Where.Expr, true, &rootParent)
		if err != nil {
			return ``, ``, err
		}
	}
	if queryMapStr == "" {
		queryMapStr = defaultQueryMapStr
	}
	// 处理
	if len(Select.From) != 1 {
		return ``, ``, errors.New("elasticsql: multiple from currently not supported")
	}
	// require Select from and convert to string
	table = sqlparser.String(Select.From)
	table = strings.Replace(table, "`", "", -1)

	// require queryFrom and query Size
	from, size := "0", "1"
	if Select.Limit != nil {
		if Select.Limit.Offset != nil {
			from = sqlparser.String(Select.Limit.Offset)
		}
		size = sqlparser.String(Select.Limit.Rowcount)
	}
	var aggFlag = false
	var aggStr []byte
	funcArr, _, err := handleSelectExpr(Select.SelectExprs)
	if err != nil {
		return ``, ``, err
	}

	if len(Select.GroupBy) > 0 || len(funcArr) > 0 {
		aggFlag = true
		if aggStr, err = handleSelectGroupBy(Select, size); err != nil {
			return ``, ``, err
		}
		size = "0"
	}
	// Handle order by
	// when executating aggregations, order by is useless
	var orderByArr []string
	if aggFlag == false {
		for _, orderByExpr := range Select.OrderBy {
			orderByStr := fmt.Sprintf(`{"%v": "%v"}`, sqlparser.String(orderByExpr.Expr), orderByExpr.Direction)
			orderByArr = append(orderByArr, orderByStr)
		}
	}
	return table, buildDSL(queryMapStr, from, size, string(aggStr), orderByArr), nil
}

func handleSelectWhere(expr *sqlparser.BoolExpr, topLevel bool, parent *sqlparser.BoolExpr) (string, error) {
	if expr == nil {
		return "", errors.New("elasticsql: error expression cannot be nil here")
	}
	switch exprValue := (*expr).(type) {
	case *sqlparser.AndExpr:
		resultStr, err := handleAndExpr(exprValue, expr)
		if err != nil {
			return ``, err
		}
		if _, ok := (*parent).(*sqlparser.AndExpr); ok {
			return resultStr, nil
		}
		return fmt.Sprintf(`{"bool" : {"must" : [%v]}}`, resultStr), nil
	case *sqlparser.OrExpr:
		resultStr, err := handleOrExpr(exprValue, expr)
		if err != nil {
			return ``, err
		}
		if _, ok := (*parent).(*sqlparser.OrExpr); ok {
			return resultStr, nil
		}
		return fmt.Sprintf(`{"bool" : {"should" : [%v]}}`, resultStr), nil
	case *sqlparser.ComparisonExpr:
		return handleComparisonExpr(exprValue, topLevel)
	case *sqlparser.NullCheck:
		return "", errors.New("elasticsql: null check expression currently not supported")
	case *sqlparser.RangeCond:
		// between a and b
		colName, ok := exprValue.Left.(*sqlparser.ColName)
		if !ok {
			return "", errors.New("elasticsql: range column name missing")
		}
		colNameStr := sqlparser.String(colName)
		fromStr := strings.Trim(sqlparser.String(exprValue.From), `'`)
		toStr := strings.Trim(sqlparser.String(exprValue.To), `'`)
		resultStr := fmt.Sprintf(`{"range" : {"%v" : {"from" : "%v", "to" : "%v"}}}`, colNameStr, fromStr, toStr)
		if topLevel {
			resultStr = fmt.Sprintf(`{"bool" : {"must" : [%v]}}`, resultStr)
		}
		return resultStr, nil
	case *sqlparser.ParenBoolExpr:
		parentBoolExpr := (*expr).(*sqlparser.ParenBoolExpr)
		boolExpr := parentBoolExpr.Expr
		return handleSelectWhere(&boolExpr, false, parent)
	case *sqlparser.NotExpr:
		return "", errors.New("elasticsql: not expression currently not supported")
	default:
		return ``, errors.New("elaticsql: logically cannot reached here")
	}
}

// handleSelectGroupBy 处置Select 里面的group by
func handleSelectGroupBy(Select *sqlparser.Select, size string) ([]byte, error) {
	var aggMap = make(map[string]interface{}, 0)
	var parentNode *map[string]interface{}
	// 解析各路groupby
	for index, item := range Select.GroupBy {
		switch itemValue := item.(type) {
		case *sqlparser.ColName:
			innerMap := handleGroupByColName(itemValue, index, size)
			if index == 0 {
				aggMap[string(itemValue.Name)] = innerMap
				parentNode = &innerMap
			} else {
				(*parentNode)["aggregations"] = map[string]interface{}{
					string(itemValue.Name): innerMap,
				}
				parentNode = &innerMap
			}
		case *sqlparser.FuncExpr:
			innerMap, err := handleGroupByFuncExpr(itemValue, size)
			if err != nil {
				return nil, err
			}
			keyName := sqlparser.String(itemValue)
			keyName = strings.Replace(keyName, `'`, ``, -1)
			keyName = strings.Replace(keyName, ` `, ``, -1)
			aggMap[keyName] = innerMap
			parentNode = &innerMap
		}
	}
	// if parentNode == nil {
	// 	return "", errors.New("elasticsql: agg not supported yet")
	// }
	// 解析avg, sum， count...等等 distint
	innerAggMap, err := handleSelectExprGroupBy(Select.SelectExprs)
	if err != nil {
		return nil, err
	}

	if len(innerAggMap) > 0 && parentNode != nil {
		(*parentNode)["aggregations"] = innerAggMap
	} else if len(innerAggMap) > 0 {
		return json.Marshal(innerAggMap)
	}
	return json.Marshal(aggMap)
}
