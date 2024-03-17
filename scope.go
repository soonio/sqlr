package sqlr

// 主要用于mysql的where条件构建

import (
	"reflect"
)

type Scope struct {
	R // 基于R驱动
}

func (w *Scope) Clone() *Scope {
	return &Scope{*w.R.Clone()}
}

// Pointer 判断指针中是否有值
func (w *Scope) Pointer(sql string, value any) *Scope {
	return w.When(!reflect.ValueOf(value).IsNil(), sql, value)
}

// When do条件成立时，sql条件生效
func (w *Scope) When(do bool, sql string, value ...any) *Scope {
	if do {
		w.ShareK(sql, value...)
	}
	return w
}

// Raw 直接拼接原生sql部分
func (w *Scope) Raw(sql string, value ...any) *Scope {
	w.ShareK(sql, value...)
	return w
}

// In 构建mysql in语句
func (w *Scope) In(field string, item ...any) *Scope {
	var l = len(item)
	if l == 0 {
		return w
	}
	w.ShareK(In(field, l), item...)

	return w
}

// Where 构建where语句
func (w *Scope) Where() string {
	if len(w.k) > 0 {
		return Join("WHERE", w.R.Where())
	}
	return ""
}
