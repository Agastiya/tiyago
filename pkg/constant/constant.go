package constant

// Path File
const (
	Environment = "../environment/"
)

// Environment
const (
	Local       = "local"
	Development = "development"
	Production  = "production"
)

// Database
const (
	Postgres = "postgres"
	Mongo    = "mongo"
	Redis    = "redis"
)

// Api Headers
const (
	ContentTypeJSON = "application/json"
)

// Time / Location
const (
	DateTimeFormat = "2006-01-02 15:04:05"
	TimeLocation   = "Asia/Jakarta"

	THIRTYSECOND   = 30          // 30 seconds
	THIRTYMINUTES  = 60 * 30     // 1,800 seconds = 30 minutes
	ONEHOUR        = 60 * 60     // 3,600 seconds = 1 hour
	ONEANDHALFHOUR = 60 * 90     // 5,400 seconds = 1.5 hours
	TWOHOURS       = 60 * 60 * 2 // 7,200 seconds = 2 hours
	ONEDAY         = 3600 * 24   // 86,400 seconds = 24 hours = 1 day
	THREEHOURS     = 60 * 60 * 3 // 10,800 seconds = 3 hours
)

// Cron Job
const (
	CronEvery1AM = "CRON_TZ=Asia/Jakarta 00 01 * * *"
	CronEvery3AM = "CRON_TZ=Asia/Jakarta 00 03 * * *"
)

type contextKey string

const (
	ClaimsKey contextKey = "claims_value"
)
