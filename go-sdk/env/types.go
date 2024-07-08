package env

type ENV string

const (
	DEVELOPMENT ENV = "development"
	STAGING     ENV = "staging"
	PRODUCTION  ENV = "production"
)
