package metric

type statusCode uint

// Log is nginx log struct
type Log struct {
	RemoteAddr    string
	RemoteUser    string
	TimeLocal     string
	Request       string
	Status        statusCode
	BodyBytes     uint
	HTTPReferer   string
	HTTPUserAgent string
}

type Item struct {
	Log
	Error error
}
