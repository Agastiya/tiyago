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
