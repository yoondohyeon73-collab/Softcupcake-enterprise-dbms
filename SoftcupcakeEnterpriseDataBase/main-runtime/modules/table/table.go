package table

type Column_type int

const (
	CT_none Column_type = iota
	CT_number
	CT_text
)

type Column struct {
	Type     Column_type
	Name     string
	Is_key   bool
	Not_null bool
}

type Table struct {
	Table_name     string
	Columns_struct []Column
}
