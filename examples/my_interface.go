package examples

type explainer interface {
	value() string
	pointer() string
	_type() string
}
