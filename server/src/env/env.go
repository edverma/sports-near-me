package env

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

const (
	jwtSecretEnvVarName = "sports-near-me_JWT_SECRET"
	apiTokenEnvVarName  = "sports-near-me_API_TOKEN"
)

var (
	SigningMethod *jwt.SigningMethodHMAC
	JwtSecret     []byte
	ApiToken      string
)

func init() {
	SigningMethod = jwt.SigningMethodHS256

	JwtSecret = []byte(os.Getenv(jwtSecretEnvVarName))
	if len(JwtSecret) == 0 {
		panic(fmt.Errorf("jwt secret is empty, please set %s env variable", jwtSecretEnvVarName))
	}

	ApiToken = os.Getenv(apiTokenEnvVarName)
	if len(JwtSecret) == 0 {
		panic(fmt.Errorf("api token is empty, please set %s env variable", apiTokenEnvVarName))
	}
}
