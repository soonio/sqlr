package sqlr

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {

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

	var sql, vs = Insert("user", r)

	fmt.Println(sql)
	fmt.Println(vs)
}
