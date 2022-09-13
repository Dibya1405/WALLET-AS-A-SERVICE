package config

import (
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

var Dbs []Db

func GetAllConfiguration() string {
	file, err := ioutil.ReadFile("./config-ci.json")
	utils.CheckError(err)
	json.Unmarshal([]byte(file), &Dbs)
	RequiredDsn := Dbs[0].Mysql.Dsn
	return RequiredDsn
}
