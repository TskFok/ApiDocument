package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
)

//go:embed conf.yaml
var conf []byte

func InitConfig() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewReader(conf))

	if nil != err {
		panic(err)
	}

}

func GetSwaggerMap() interface{} {
	return viper.Get("swagger")
}
