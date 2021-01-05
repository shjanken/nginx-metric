package metric

import "time"

type statusCode uint

// Log is nginx log struct
type Log struct {
	RemoteAddr    string
	RemoteUser    string
	TimeLocal     time.Time
	Request       string
	Status        statusCode
	BodyBytes     uint
	HTTPReferer   string
	HTTPUserAgent string
}

// Item is metric item
type Item struct {
	Log
	Error error
}
