package constants

const (
	ConfigPath        = "./doc/config"
	ConfigNameDev     = "config.dev"
	ConfigNameStaging = "config.staging"
	ConfigNameProd    = "config.prod"
)

const (
	FailedConfigRead      = "failed to read config : %s"
	FailedConfigUnmarshal = "failed to unmarshal config : %s"
	FailedConfigValidate  = "failed to validate config : %s"
	FailedEnvRead         = "failed to read .env: %v"
)

const (
	ExtensionEnv = ".env"
	ExtensionYml = ".yml"
)
