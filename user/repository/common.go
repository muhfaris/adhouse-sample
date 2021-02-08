package repository

// QueryResult is to wrap query result
type QueryResult struct {
	Result interface{}
	Error  error
}
