package utils

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"gorm.io/gorm"
	"strconv"
)

// Configurations wraps all the config variables required by the auth service
type Configurations struct {
	ServerAddress              string
	DBHost                     string
	DBName                     string
	DBUser                     string
	DBPass                     string
	DBPort                     string
	DBConn                     *sqlx.DB
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int // in minutes
	SendGridApiKey             string
	MailVerifCodeExpiration    int // in hours
	PassResetCodeExpiration    int // in minutes
	MailVerifTemplateID        string
	PassResetTemplateID        string
	DBGormConn                 *gorm.DB
}

// NewConfigurations returns a new Configuration object
func NewConfigurations(logger hclog.Logger) (*Configurations, *errs.AppError) {

	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return nil, errs.NewUnexpectError(err.Error())
	}
	JwtExpiration, err := strconv.Atoi(myEnv["JWT_EXPIRATION"])
	if err != nil {
		return nil, errs.NewUnexpectError("error during JWT_EXPIRATION  convertion: " + err.Error())
	}
	MailVerifCodeExpiration, err := StringToInt(logger, myEnv["MAIL_VERIF_CODE_EXPIRATION"])
	if err != nil {
		return nil, errs.NewUnexpectError("error during MailVerifCodeExpiration convertion: " + err.Error())
	}
	PassResetCodeExpiration, err := StringToInt(logger, myEnv["PASS_RESET_CODE_EXPIRATION"])
	if err != nil {
		return nil, errs.NewUnexpectError("error during PassResetCodeExpiration convertion: " + err.Error())
	}
	configs := &Configurations{
		ServerAddress:           myEnv["DB_ADDRESS"],
		DBHost:                  myEnv["DB_ADDR"],
		DBName:                  myEnv["DB_NAME"],
		DBUser:                  myEnv["DB_USER"],
		DBPass:                  myEnv["DB_PASSWORD"],
		DBPort:                  myEnv["DB_PORT"],
		JwtExpiration:           JwtExpiration,
		SendGridApiKey:          myEnv["SENDGRID_API_KEY"],
		MailVerifCodeExpiration: *MailVerifCodeExpiration,
		PassResetCodeExpiration: *PassResetCodeExpiration,
		MailVerifTemplateID:     myEnv["MAIL_VERIF_TEMPLATE_ID"],
		PassResetTemplateID:     myEnv["PASS_RESET_TEMPLATE_ID"],
	}

	logger.Debug("serve port", configs.ServerAddress)
	logger.Debug("db host", configs.DBHost)
	logger.Debug("db name", configs.DBName)
	logger.Debug("db port", configs.DBPort)
	logger.Debug("jwt expiration", configs.JwtExpiration)

	return configs, nil
}
