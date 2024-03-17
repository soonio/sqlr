package sqlr

import (
	"testing"
)

func TestR_Insert(t *testing.T) {
	var SQL = "(`username`,`gender`,`password`,`avatar`,`cover`,`created`,`update`) VALUES (?,?,?,?,?,?,?)"
	var Vals = []any{"张三", 1, "123456", "https://www.demo.com/avatar.png", "https://www.demo.com/cover.png", 86400, 86400}

	var r = new(R).
		Pair("username", "张三").
		Pair("gender", 1).
		Map(map[string]any{"password": "123456"}).
		Struct(struct {
			Avatar string `db:"avatar"`
			Cover  string `db:"cover"`
		}{
			Avatar: "https://www.demo.com/avatar.png",
			Cover:  "https://www.demo.com/cover.png",
		}, "db").
		ShareV(86400, "created", "update")

	var partSQL = r.Insert()
	if SQL != partSQL {
		t.Error("SQL拼接错误")
	}

	var values = r.V()
	for i := 0; i < len(Vals); i++ {
		if Vals[i] != values[i] {
			t.Error("参数结果列表错误")
		}
	}
}

func TestR_Update(t *testing.T) {
	var SQL = "`username` = ?,`gender` = ?,`password` = ?,`avatar` = ?,`cover` = ?,`update` = ?"
	var Vals = []any{"张三", 1, "123456", "https://www.demo.com/avatar.png", "https://www.demo.com/cover.png", 86400}

	var r = new(R).
		Pair("username", "张三").
		Pair("gender", 1).
		Map(map[string]any{"password": "123456"}).
		Struct(struct {
			Avatar string `db:"avatar"`
			Cover  string `db:"cover"`
		}{
			Avatar: "https://www.demo.com/avatar.png",
			Cover:  "https://www.demo.com/cover.png",
		}, "db").
		Pair("update", 86400)

	var partSQL = r.Update()
	if SQL != partSQL {
		t.Error("SQL拼接错误", partSQL)
	}

	var values = r.V()
	for i := 0; i < len(Vals); i++ {
		if Vals[i] != values[i] {
			t.Error("参数结果列表错误")
		}
	}
}

func TestR_Where(t *testing.T) {
	var SQL = "username = ? AND gender = ? AND update > ?"
	var Vals = []any{"张三", 1, 86400}

	var r = new(R).
		Pair("username = ?", "张三").
		Pair("gender = ?", 1).
		Pair("update > ?", 86400)

	var partSQL = r.Where()
	if SQL != partSQL {
		t.Error("SQL拼接错误", partSQL)
	}

	var values = r.V()
	for i := 0; i < len(Vals); i++ {
		if Vals[i] != values[i] {
			t.Error("参数结果列表错误")
		}
	}
}
