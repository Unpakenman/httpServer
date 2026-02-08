package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"httpServer/internal/app/constants"
	"os"
	"time"
)

type Values struct {
	HttpServer      *HTTPServerConfig      `envconfig:"HTTP" required:"true"`
	ClinicsDB       *DBConfig              `envconfig:"DB_CLINICS" required:"true"`
	HttpClient      *HTTPClientConfig      `envconfig:"HTTP_CLIENT" required:"true"`
	SomeHttpService *SomeHttpServiceConfig `envconfig:"SOME_HTTP_SERVICE" required:"true"`
}

type HTTPServerConfig struct {
	ServerPort     int32  `envconfig:"HTTP_SERVER_PORT" required:"true"`
	ApiDefaultPath string `envconfig:"HTTP_API_DEFAULT_PATH" required:"true"`
}

type DBConfig struct {
	DBName   string `envconfig:"DB_CLINICS_NAME" required:"true"`
	User     string `envconfig:"DB_CLINICS_USER" required:"true"`
	Password string `envconfig:"DB_CLINICS_PASSWORD" required:"true"`
	Hostname string `envconfig:"DB_CLINICS_HOSTNAME" required:"true"`
	SSLMode  string `envconfig:"DB_CLINICS_SSLMODE" required:"false"`
	Port     int32  `envconfig:"DB_CLINICS_PORT" required:"true"`

	MaxOpenConns                    int           `envconfig:"DB_CLINICS_MAX_OPEN_CONNS" required:"true"`
	MaxIdleConns                    int           `envconfig:"DB_CLINICS_MAX_IDLE_CONNS" required:"true"`
	MaxLifeTimeConns                time.Duration `envconfig:"DB_CLINICS_MAX_LIFETIME_CONNS" required:"true"`
	StatementTimeout                time.Duration `envconfig:"DB_CLINICS_STATEMENT_TIMEOUT" required:"false"`
	IdleInTransactionSessionTimeout time.Duration `envconfig:"DB_CLINICS_IDLE_IN_TRANSACTION_SESSION_TIMEOUT" required:"false"`
	LockTimeout                     time.Duration `envconfig:"DB_CLINICS_LOCK_TIMEOUT" required:"false"`
}

type HTTPClientConfig struct {
	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`
}

type SomeHttpServiceConfig struct {
	URL string `envconfig:"URL" required:"true"`
}

var Config *Values

func New() (*Values, error) {
	err := LoadEnvFile()
	if err != nil {
		return nil, err
	}

	cfg := &Values{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadEnvFile() error {
	if needUseLocalEnvFile() {
		err := godotenv.Load(constants.DefaultEnvFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func needUseLocalEnvFile() bool {
	for _, arg := range os.Args {
		if arg == constants.UseLocalEnvFileArg {
			return true
		}
	}
	return false
}
