package env

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

const (
	jwtSecretEnvVarName = "TEMPLATE_JWT_SECRET"
	apiTokenEnvVarName  = "TEMPLATE_API_TOKEN"
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
