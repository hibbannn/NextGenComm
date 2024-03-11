package domain

import "google.golang.org/grpc/credentials"

type Config struct {
	AppName     string            `mapstructure:"app_name"`
	Environment string            `mapstructure:"ENVIRONMENT"`
	Database    DatabaseConfig    `mapstructure:"database"`
	Logging     LoggingConfig     `mapstructure:"logging"`
	RateLimiter RateLimiterConfig `mapstructure:"rate_limiter"`
	Features    FeatureFlags      `mapstructure:"features"`
	Server      ServerConfig      `mapstructure:"server"`
	SSL         SSLConfig         `mapstructure:"ssl"`
	Encryption  EncryptionConfig  `mapstructure:"encryption"`
	Mail        MailConfig        `mapstructure:"mail"`
	Cache       CacheConfig       `mapstructure:"cache"`
	Services    ServicesConfig    `mapstructure:"services"`
	JWT         JWTConfig         `mapstructure:"jwt"`
}

type CacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Enabled  bool   `mapstructure:"enabled"`
	Timeout  int    `mapstructure:"timeout"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxAge   int    `mapstructure:"max_age"`
	MaxSize  int    `mapstructure:"max_size"`
	MaxConn  int    `mapstructure:"max_conn"`
	SSL      bool   `mapstructure:"ssl"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type MailConfig struct {
	Server         string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Timeout        int    `mapstructure:"timeout"`
	MaxRetries     int    `mapstructure:"max_retries"`
	MaxMessageSize int    `mapstructure:"max_message_size"`
	RetryInterval  int    `mapstructure:"retry_interval"`
	TLS            bool   `mapstructure:"tls"`
	SSL            bool   `mapstructure:"ssl"`
}

type SSLConfig struct {
	GlobalTLSCredentials credentials.TransportCredentials
	Enabled              bool   `mapstructure:"enabled"`
	CertFile             string `mapstructure:"cert_file"`
	KeyFile              string `mapstructure:"key_file"`
}

type DatabaseConfig struct {
	Driver      string `mapstructure:"driver"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Name        string `mapstructure:"name"`
	Path        string `mapstructure:"path"`
	SSLMode     string `mapstructure:"ssl_mode"`
	Schema      string `mapstructure:"schema"`
	Migration   bool   `mapstructure:"migration"`
	MaxIdle     int    `mapstructure:"max_idle_conns"`
	MaxOpen     int    `mapstructure:"max_open_conns"`
	MaxLife     int    `mapstructure:"max_life_time"`
	MaxIdleLife int    `mapstructure:"max_idle_life_time"`
	Password    string `mapstructure:"password"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Save   bool   `mapstructure:"save_to_file"`
	Path   string `mapstructure:"file_path"`
}

type RateLimiterConfig struct {
	Enabled           bool `mapstructure:"enabled"`
	RequestsPerMinute int  `mapstructure:"requests_per_minute"`
}

type FeatureFlags struct {
	NewFeatureEnabled  bool `mapstructure:"new_feature_enabled"`
	BetaFeatureEnabled bool `mapstructure:"beta_feature_enabled"`
}

type ServerConfig struct {
	GRPC GRPCConfig `mapstructure:"grpc"`
	HTTP HTTPConfig `mapstructure:"http"`
}

type GRPCConfig struct {
	Port       int  `mapstructure:"port"`
	Reflection bool `mapstructure:"reflection"`
	Gateway    int  `mapstructure:"gateway"`
}

type HTTPConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	IdleTimeout  int    `mapstructure:"idle_timeout"`
}

type EncryptionConfig struct {
	Salt int    `mapstructure:"key"`
	Aes  string `mapstructure:"aes_key"`
}

type ServicesConfig struct {
	UserService UserService `mapstructure:"user_service"`
	MenuService MenuService `mapstructure:"menu_service"`
}

type UserService struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MenuService struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type JWTConfig struct {
	Algorithm        string `mapstructure:"algorithm"`
	SigningKey       string `mapstructure:"signing_key"`
	SecretKey        string `mapstructure:"secret_key"`
	SecretSessionKey string `mapstructure:"secret_session_key"`
	EncryptionKey    string `mapstructure:"encryption_key"`
	Issuer           string `mapstructure:"issuer"`
	Audience         string `mapstructure:"audience"`
	Expiration       int    `mapstructure:"expiration"`
	Refresh          int    `mapstructure:"refresh_token_expiration"`
}
