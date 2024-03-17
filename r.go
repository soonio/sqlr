package sqlr

import (
	"reflect"
	"strings"
)

// R 参与构建SQL的原始数据(raw data)
type R struct {
	k []string // 数据表字段的键
	v []any    // 数据表字段的值列表
}

// Clone 复制一份
func (r *R) Clone() *R {
	k := make([]string, len(r.k))
	v := make([]any, len(r.v))

	copy(k, r.k)
	copy(v, r.v)

	return &R{k: k, v: v}
}

// Pair 添加 k和v一对数据 如 k="name = ?" v="john"
func (r *R) Pair(k string, v any) *R {
	r.k = append(r.k, k)
	r.v = append(r.v, v)
	return r
}

// ShareK 添加多个key
func (r *R) ShareK(k string, v ...any) *R {
	r.k = append(r.k, k)
	r.v = append(r.v, v...)
	return r
}

// ShareV 添加多个value
func (r *R) ShareV(v any, k ...string) *R {
	r.k = append(r.k, k...)
	for i := 0; i < len(k); i++ {
		r.v = append(r.v, v)
	}
	return r
}

// Map 通过map添加 {"name": "张三", "gender": "男"}
func (r *R) Map(m map[string]any) *R {
	for k, v := range m {
		r.k = append(r.k, k)
		r.v = append(r.v, v)
	}
	return r
}

func (r *R) Struct(s any, tag ...string) *R {
	tag = append(tag, "db")
	var t = reflect.TypeOf(s)
	var v = reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fields := t.NumField()
	for i := 0; i < fields; i++ {
		key := t.Field(i).Name
		if t.Field(i).Tag.Get(tag[0]) != "" {
			key = t.Field(i).Tag.Get(tag[0])
		}
		r.k = append(r.k, key)
		r.v = append(r.v, v.Field(i).Interface())
	}
	return r
}

// K 获取所有的键
func (r *R) K() []string {
	return r.k
}

// V 获取所有的value值
func (r *R) V() []any {
	return r.v
}

// Insert 把键值对构建成SQL insert时可用的字符串
func (r *R) Insert() string {
	var bs strings.Builder
	bs.WriteString("(`")
	bs.WriteString(r.k[0])
	for i := 1; i < len(r.k); i++ {
		bs.WriteString("`,`")
		bs.WriteString(r.k[i])
	}
	bs.WriteString("`) VALUES (?")
	for i := 1; i < len(r.k); i++ {
		bs.WriteString(",?")
	}
	bs.WriteString(")")
	return bs.String()
}

// Update 把键值对构建成SQL update时可用的字符串
func (r *R) Update() string {
	var sb strings.Builder
	for i := 0; i < len(r.k); i++ {
		sb.WriteString("`")
		sb.WriteString(r.k[i])
		sb.WriteString("`")
		sb.WriteString(" = ?,")
	}
	return sb.String()[:sb.Len()-1]
}

// Where 把键值对构建成SQL where时可用的字符串
func (r *R) Where() string {
	var sb strings.Builder
	for i := 0; i < len(r.k); i++ {
		sb.WriteString(r.k[i])
		sb.WriteString(" AND ")
	}

	return sb.String()[:sb.Len()-5]
}
