package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"github.com/spf13/viper"
	"time"
)

// Struktur untuk menyimpan refresh tokens
var refreshTokensStore = make(map[string]string)

// GenerateJWT menghasilkan access token baru
func (s *Security) GenerateJWT(jwtID, sessionKey string) (string, error) {
	claims := &domain.JWTClaims{
		Session: sessionKey, // Session key dari Redis sebagai bagian dari claim
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jwtID,
			Audience:  jwt.ClaimStrings{s.jwtConfig.Audience},
			Issuer:    s.jwtConfig.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.jwtConfig.Expiration) * time.Minute)),
		},
	}

	signingMethod, err := s.JWTSigningMethod()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(viper.GetViper().GetString(s.jwtConfig.SecretKey)))
}

// StoreRefreshToken menyimpan refresh token
func (s *Security) StoreRefreshToken(username, refreshToken string) {
	refreshTokensStore[username] = refreshToken
}

// ValidateRefreshToken memvalidasi refresh token
func (s *Security) ValidateRefreshToken(refreshToken string) (string, bool) {
	for username, token := range refreshTokensStore {
		if token == refreshToken {
			return username, true
		}
	}
	return constants.EmptyString, false
}

// RevokeRefreshToken mencabut refresh token
func (s *Security) RevokeRefreshToken(username string) {
	delete(refreshTokensStore, username)
}

// JWTSigningMethod menentukan JWT signing method berdasarkan konfigurasi
func (s *Security) JWTSigningMethod() (jwt.SigningMethod, error) {
	switch s.jwtConfig.Algorithm {
	case constants.AlgorithmHS256:
		return jwt.SigningMethodHS256, nil
	case constants.AlgorithmHS512:
		return jwt.SigningMethodHS512, nil
	case constants.AlgorithmHS384:
		return jwt.SigningMethodHS384, nil
	case constants.AlgorithmRS256:
		return jwt.SigningMethodRS256, nil
	case constants.AlgorithmRS512:
		return jwt.SigningMethodRS512, nil
	case constants.AlgorithmRS384:
		return jwt.SigningMethodRS384, nil
	case constants.AlgorithmES256:
		return jwt.SigningMethodES256, nil
	case constants.AlgorithmES512:
		return jwt.SigningMethodES512, nil
	case constants.AlgorithmES384:
		return jwt.SigningMethodES384, nil
	case constants.AlgorithmPS256:
		return jwt.SigningMethodPS256, nil
	case constants.AlgorithmPS512:
		return jwt.SigningMethodPS512, nil
	case constants.AlgorithmPS384:
		return jwt.SigningMethodPS384, nil
	case constants.AlgorithmEdDSA:
		return jwt.SigningMethodEdDSA, nil
	case constants.AlgorithmNone:
		return jwt.SigningMethodNone, nil
	default:
		return nil, fmt.Errorf("invalid signing method: %s", s.jwtConfig.Algorithm)
	}
}
