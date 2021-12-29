package config

import (
	"fiber-101/models"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// default
var C = Configuration{
	APP_ENV:  "development",
	APP_PORT: 9090,

	DATABASE_TEST: false,

	LOG_LEVEL: "debug",
}

type Configuration struct {
	APP_ENV  string `mapstructure:"APP_ENV"`
	APP_PORT int    `mapstructure:"APP_PORT"`

	DATABASE_DSN  string `mapstructure:"DATABASE_DSN"`
	DATABASE_TEST bool   `mapstructure:"DATABASE_TEST"`

	DATABASE_HOST     string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT     int    `mapstructure:"DATABASE_PORT"`
	DATABASE_USER     string `mapstructure:"DATABASE_USER"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	DATABASE_SSLMODE  string `mapstructure:"DATABASE_SSLMODE"`
	DATABASE_TIMEZONE string `mapstructure:"DATABASE_TIMEZONE"`

	LOG_LEVEL string `mapstructure:"LOG_LEVEL"`
}

func Load(filePath string) {
	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindEnvs(C)

	err = viper.Unmarshal(&C)
	if err != nil {
		logrus.Errorf("unable to decode into struct, %v", err)
		panic("[CONFIG] unable to decode into struct")
	}

	e := validate(C)
	if e != nil {
		for _, s := range e {
			logrus.Errorf("[CONFIG] %s: %s\n", s.FailedField, s.Tag)
		}
		panic("[CONFIG] invalid configuration")
	}

	PrintConfig()
}

func PrintConfig() {
	logrus.Debugf("[CONFIG]: %#v\n", C)
}

func validate(c Configuration) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				fmt.Printf("can't bind config from ENV, %v\n", err)
			}
		}
	}
}
