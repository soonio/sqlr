package sqlr

type m uint

const (
	SelectSQL m = 0 // 查询
	InsertSQL m = 1 // 插入
	UpdateSQL m = 2 // 更新
	DeleteSQL m = 3 // 删除
)

type Builder struct {
	table string
	r     *R
	s     *Sorter
	m     m
}

func New() *Builder {
	return &Builder{r: new(R), m: SelectSQL}
}

func (b *Builder) Select(table string) *Builder {
	b.table = table
	b.m = SelectSQL
	return b
}

func (b *Builder) Update(table string) *Builder {
	b.table = table
	b.m = UpdateSQL
	return b
}

func (b *Builder) Insert(table string) *Builder {
	b.table = table
	b.m = InsertSQL
	return b
}

func (b *Builder) Delete(table string) *Builder {
	b.table = table
	b.m = DeleteSQL
	return b
}

func (b *Builder) ToSQL() (string, error) {
	return "", nil
}
