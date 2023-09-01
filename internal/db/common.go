package db

// DBError defines db error
type DBError struct {
	Msg  string // DB error message
	Type int
}

func (e *DBError) Error() string {
	return e.Msg
}

const (
	InvalidContent = 100 // Error returned if data isn't correct
	AlreadyExist   = 101 // Error returned when data already exist

	InvoiceAlreadyPaid    = 200
	InvoiceAmountNotFound = 201
	InvoiceNotFound       = 203
)
