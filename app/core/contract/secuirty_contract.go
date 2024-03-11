package contract

type SecurityContract interface {
	Encrypt(data interface{}) (string, error)
	Decrypt(cipherText string) ([]byte, error)
	GenerateJWT(jwtID, sessionKey string) (string, error)
}
