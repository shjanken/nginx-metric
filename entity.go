package log

type statusCode uint

type Log struct {
	RemoteAddr string
	RemoteUser string
	TimeLocal string
	Request string
	Status statusCode
	BodyBytes uint
	HttpReferer string
	HttpUserAgent string
}
