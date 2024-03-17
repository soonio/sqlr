package sqlr

import (
	"testing"
)

func TestSorter(t *testing.T) {
	var partSQL = "order by name desc, score desc"
	var s = new(Sorter).
		Allow("name", "gender", "score").
		When(true, "name", "Desc").
		When(false, "gender", "dEsc").
		Add("score", "DESc").
		String()

	if s != partSQL {
		t.Error("SQL拼接错误", s)
	}
}
