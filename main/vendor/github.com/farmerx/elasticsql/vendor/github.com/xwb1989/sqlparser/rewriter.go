package sqlparser

import (
	_ "fmt"
	"reflect"
)

var typeOfBytes = reflect.TypeOf([]byte(nil))
var typeOfSQLNode = reflect.TypeOf((*SQLNode)(nil)).Elem()

type Rewriter func([]byte) []byte

func Rewrite(node SQLNode, rewriter Rewriter) {
	rewrite(reflect.ValueOf(node), rewriter)
}
func rewrite(nodeVal reflect.Value, rewriter Rewriter) {
	if !nodeVal.IsValid() {
		return
	}
	nodeTyp := nodeVal.Type()
	switch nodeTyp.Kind() {
	case reflect.Slice:
		if nodeTyp == typeOfBytes && !nodeVal.IsNil() {
			val := rewriter(nodeVal.Bytes()) //use rewriter to rewrite the bytes
			nodeVal.SetBytes(val)
		} else if nodeTyp.Implements(typeOfSQLNode) {
			for i := 0; i < nodeVal.Len(); i++ {
				m := nodeVal.Index(i)
				rewrite(m, rewriter)
			}
		}
	case reflect.Struct:
		for i := 0; i < nodeVal.NumField(); i++ {
			f := nodeVal.Field(i)
			rewrite(f, rewriter)
		}
	case reflect.Ptr, reflect.Interface:
		rewrite(nodeVal.Elem(), rewriter)
	}
}
