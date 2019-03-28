package config

import (
    "fmt"
    "github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func init()  {
	projectName := "go-mega"
	dbType := GetDBType()
	log.Println("OS DBTYPE:",dbType)
	if IsHeroku() {
		log.Println("Get Env from os.env")
	}else {
		log.Println("Init viper")
		getConfig(projectName)
	}
}

func getConfig(projectName string)  {
	viper.SetConfigName("config") //name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s",projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s",projectName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file :%s",err))
	}
}
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true",usr,pwd,host,db,charset)
}
func GetSMTPConfig()(server string,port int,user,pwd string) {
	if IsHeroku() {
		server = os.Getenv("MAIL_SMTP")
		port, _ = strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
		user = os.Getenv("MAIL_USER")
		pwd = os.Getenv("MAIL_PASSWORD")
		return
	}
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}
func GetServerURL() (url string) {
	if IsHeroku() {
		url = os.Getenv("SERVER_URL")
		return
	}
	url = viper.GetString("server.url")
	return
}
// GetHerokuConnectingString func

func GetHerokuConnectingString() string {
	return os.Getenv("DATABASE_URL")
}
// GetDBType func
func GetDBType() string {
	dbtype := os.Getenv("DBTYPE")
	return dbtype
}

// IsHeroku func
func IsHeroku() bool {
	return GetDBType() == "heroku"
}