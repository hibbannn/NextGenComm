package security

import "github.com/hibbannn/grpc-http-boilerplate/app/core/domain"

// Security adalah struct untuk mengelola JWT
type Security struct {
	jwtConfig     domain.JWTConfig
	encryptConfig domain.EncryptionConfig
}

// NewJWTService adalah konstruktor untuk Security
func NewSecurity(cfg domain.JWTConfig, encryptConfig domain.EncryptionConfig) *Security {
	return &Security{jwtConfig: cfg, encryptConfig: encryptConfig}
}
