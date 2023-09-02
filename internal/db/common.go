package db

// DBError defines db error
type DBError struct {
	Msg  string // DB error message
	Type int
}

func (e *DBError) Error() string {
	return e.Msg
}
