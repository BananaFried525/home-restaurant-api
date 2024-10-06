package ports

type TableService interface {
	AddTable(Number int) error
	// GetListTable(limit int, offset int) error
	// GetTableDetail(ID uint) error
	// UpdateTable(ID uint, table domain.Table) error
	// DeleteTable(ID uint) error
}
