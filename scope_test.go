package sqlr

import (
	"fmt"
	"testing"
)

func TestScope_Build(t *testing.T) {

	var status = 1
	var process *int
	var SQL = "WHERE username = ? AND status = ? AND gender = ? AND sex = ? AND favor IN (?,?,?)"
	var Vals = []any{"John", &status, 1, 2, 1, 2, 3}

	var s = new(Scope).
		When(true, "username = ?", "John").
		When(false, "name = ?", "Leo").
		Pointer("status = ?", &status).
		Pointer("process = ?", process).
		Raw("gender = ? AND sex = ?", 1, 2).
		In("favor", 1, 2, 3)

	var partSQL = s.Where()
	fmt.Println(partSQL)
	if SQL != partSQL {
		t.Error("SQL拼接错误", partSQL)
	}

	var values = s.V()
	for i := 0; i < len(Vals); i++ {
		if Vals[i] != values[i] {
			t.Error("参数结果列表错误")
		}
	}

}
