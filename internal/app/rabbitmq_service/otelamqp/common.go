package otelamqp

type contextKeyType int64

const (
	StartTimeKey contextKeyType = iota
	MetricAttrsKey

	//error types
	ErrorTypeInternalError         = "internal-error"
	ErrorTypeDecodeError           = "decode-error"
	ErrorTypeSchemaValidationError = "schema-validation-error"

	scopeName    = "vcs.bingo-boom.ru/digital_department/ru/application-layer/resources/dependencies/go-modules/otelamqp"
	scopeVersion = "1.0.10"
)

var (
	defaultBounds = []float64{0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1, 2.5, 5, 7.5, 10}
)
