package query

// Token for type query
type Token interface {
	Set(str *string) Token
	Get() Token
}
