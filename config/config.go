package config

import (
	"WALLET-AS-A-SERVICE/logger"
	"WALLET-AS-A-SERVICE/utils"
	"encoding/json"
	"io/ioutil"
)

type Db struct {
	Mysql *Mysql `json:"mysql"`
}

type Mysql struct {
	Dsn string `json:"dsn"`
}

var DbConfig Db

func GetAllConfiguration() string {
	file, err := ioutil.ReadFile("./config-ci.json")
	utils.CheckError(err)
	json.Unmarshal(file, &DbConfig)
	RequiredDsn := DbConfig.Mysql.Dsn
	logger.GetMyLogger().Info(RequiredDsn)
	return RequiredDsn
}
