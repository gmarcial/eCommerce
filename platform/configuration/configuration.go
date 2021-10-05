package configuration

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gmarcial/eCommerce/platform/infrastructure/filepathutil"
	"time"
)

const (
	envFileRelativePath  = ".env"
	errLoadConfiguration = "an error occurred to load the configurations"
)

//Configuration the configuration utilized in application
type Configuration struct {
	GrpcServerAddress string    `env:"GRPC_SERVER_ADDRESS"`
	BlackFridayDay    time.Time `env:"BLACK_FRIDAY_DAY"`
	HttpServerPort    string    `env:"HTTP_SERVER_PORT"`
}

//LoadConfiguration read the .env file and the use to construct the configuration
func LoadConfiguration() *Configuration {
	envFilePath, err := filepathutil.JoinWithRootDir(envFileRelativePath)
	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errLoadConfiguration, err.Error())
		panic(errorMessage)
	}

	configuration := new(Configuration)
	err = cleanenv.ReadConfig(envFilePath, configuration)
	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errLoadConfiguration, err.Error())
		panic(errorMessage)
	}

	return configuration
}
