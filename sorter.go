package sqlr

import (
	"strings"
)

type Sorter struct {
	data  [][2]string // 参与排序的字段及排序方式
	allow []string    // 允许的排序列表
}

// When 当条件成立是参与排序
func (s *Sorter) When(do bool, key, sort string) *Sorter {
	if do {
		s.Add(key, sort)
	}
	return s
}

// Add 添加一个排序条件
func (s *Sorter) Add(key, sort string) *Sorter {
	var lsf = len(s.allow)
	if lsf > 0 {
		for j := 0; j < lsf; j++ {
			if key == s.allow[j] {
				s.data = append(s.data, [2]string{key, sort})
			}
		}
	} else {
		s.data = append(s.data, [2]string{key, sort})
	}

	return s
}

// Allow 允许的排序字段
func (s *Sorter) Allow(filed ...string) *Sorter {
	s.allow = append(s.allow, filed...)
	return s
}

// 生成字排序条件字符串
func (s *Sorter) String() string {
	var lsd = len(s.data)
	if lsd > 0 {
		var bs strings.Builder
		bs.WriteString("order by ")
		bs.WriteString(s.data[0][0])
		if strings.ToLower(s.data[0][1]) == "desc" {
			bs.WriteString(" desc")
		}
		for i := 1; i < lsd; i++ {
			bs.WriteString(", ")
			bs.WriteString(s.data[i][0])
			if strings.ToLower(s.data[i][1]) == "desc" {
				bs.WriteString(" desc")
			}
		}
		return bs.String()
	}
	return ""
}
