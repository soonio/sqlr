package sqlr

import (
	"fmt"
	"strings"
)

// Update 直接构建更新语句
func Update(table string, data *R, scope ...*R) (string, []any) {
	table = fmt.Sprintf("`%s`", table)
	var bs strings.Builder
	bs.WriteString("UPDATE ")
	bs.WriteString(table)
	bs.WriteString(" SET ")
	bs.WriteString(data.Update())

	if len(scope) > 0 {
		bs.WriteString(" WHERE ")
		bs.WriteString(scope[0].Where())

		var all []any
		all = append(all, data.v...)
		all = append(all, scope[0].v...)

		return bs.String(), all
	}
	return bs.String(), data.V()
}

// Insert 直接构建Insert语句
func Insert(table string, data *R) (string, []any) {
	table = fmt.Sprintf("`%s`", table)
	var bs strings.Builder
	bs.WriteString("INSERT INTO ")
	bs.WriteString(table)
	bs.WriteString(data.Insert())
	return bs.String(), data.V()
}

// InsertBatch 批量插入 ⚠️未做异常判断，务必保证 columns的列数和rows[n]的列数相等
func InsertBatch(table string, columns []string, rows [][]any) (string, []any) {
	table = fmt.Sprintf("`%s`", table)
	var bs strings.Builder
	var ss strings.Builder
	var vs []any
	bs.WriteString("INSERT INTO ")
	bs.WriteString(table)

	var lc = len(columns)

	bs.WriteString("(`")
	bs.WriteString(columns[0])
	ss.WriteString("(?")
	for i := 1; i < lc; i++ {
		bs.WriteString("`,`")
		bs.WriteString(columns[i])
		ss.WriteString(" ,?")
	}
	bs.WriteString("`) VALUES ")
	ss.WriteString(")")

	var lr = len(rows)

	var sss = ss.String()

	if lr > 0 {
		bs.WriteString(sss)
		vs = append(vs, rows[0]...)
		for i := 1; i < lr; i++ {
			bs.WriteString(",")
			bs.WriteString(sss)
			vs = append(vs, rows[i]...)
		}
	}
	return bs.String(), vs
}

// In 直接构建In语句模版
func In(filed string, count int) string {
	var sb strings.Builder
	sb.WriteString(filed)
	sb.WriteString(" IN (?")
	for i := 1; i < count; i++ {
		sb.WriteString(",?")
	}
	sb.WriteString(")")
	return sb.String()
}

// Join 把多段sql使用string builder连接起来
func Join(parts ...string) string {
	var bs strings.Builder

	for i := 0; i < len(parts); i++ {
		bs.WriteString(parts[i])
		bs.WriteString(" ")
	}
	return bs.String()[:bs.Len()-1]
}
