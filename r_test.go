package sqlr

import (
	"fmt"
	"testing"
)

type Demo struct {
	Avatar string `db:"avatar"`
	Cover  string `db:"cover"`
}

func TestNew(t *testing.T) {
	var SQL = "(`username`,`gender`,`created`,`update`,`password`,`avatar`,`cover`) VALUES (?,?,?,?,?,?,?)"
	var Vals = []any{"张三", 1, 86400, "123456", "https://www.demo.com/avatar.png", "https://www.demo.com/cover.png"}

	var r = new(R).
		Append("username", "张三").
		Append("gender", 1).
		AppendV(86400).
		AppendK("created", "update").
		Map(map[string]any{"password": "123456"}).
		Struct(Demo{
			Avatar: "https://www.demo.com/avatar.png",
			Cover:  "https://www.demo.com/cover.png",
		}, "db")

	var partSQL = r.Insert()
	if SQL != partSQL {
		t.Fail()
	}

	var values = r.V()
	for i := 0; i < len(Vals); i++ {
		if Vals[i] != values[i] {
			t.Fail()
		}
	}
	fmt.Println(t.Name())
}
