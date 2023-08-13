package config

import (
	"fmt"
	"go-notif/src/helper"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./../..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Warningf("%v", err)
	}
	log.Info("Using config file: ", viper.ConfigFileUsed())
}

func Port() string {
	if !viper.IsSet("port") {
		return "8080"
	}
	return viper.GetString("port")
}

func DBHost() string {
	return viper.GetString("database.host")
}

func DBDatabase() string {
	return viper.GetString("database.database")
}

func DBUser() string {
	return viper.GetString("database.username")
}

func DBPassword() string {
	return viper.GetString("database.password")
}

func DBDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser(), DBPassword(), DBHost(), DBDatabase())
}

func MaxIdleConns() int {
	if !viper.IsSet("database.maxIdleConns") {
		return 3
	}
	return viper.GetInt("database.maxIdleConns")
}

func MaxOpenConns() int {
	if !viper.IsSet("database.maxOpenConns") {
		return 15
	}
	return viper.GetInt("database.maxOpenConns")
}

func ConnMaxLifeTime() time.Duration {
	time := viper.GetString("database.connMaxLifeTime")
	return helper.ParseTimeDuration(time, DefaultConnMaxLifeTime)
}

func ConnMaxIdleTime() time.Duration {
	time := viper.GetString("database.connMaxIdleTime")
	return helper.ParseTimeDuration(time, DefaultConnMaxIdleTime)
}

func RedisHost() string {
	return viper.GetString("redis.host")
}

func RedisExp() time.Duration {
	time := viper.GetString("redis.exp")
	return helper.ParseTimeDuration(time, DefaultRedisExpiredDuration)
}
