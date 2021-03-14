package db

import (
	"fmt"
	"github.com/romaxa83/skeleton/internal/console"
)

const (
	mysqlDriver = "mysql"
	pgsqlDriver = "pgsql"
	titleDriver = "Choice DB driver"
	titleVersion = "Choice DB version"
	titleUsername = "Enter username for db"
	titlePassword = "Enter password for db"
	titleDbName = "Enter db name"
)

var mysqlVersion = []string{"8.0", "5.7", "5.6"}
var pgsqlVersion = []string{"12", "11", "10"}
var dbDriver = []string{mysqlDriver, pgsqlDriver}

type Data struct {
	Driver string
	Version string
	UserName string
	DbName string
	Password string
}

func GetData() *Data {
	// запрашиваем тип бд
	d, err := console.Select(console.NewDataForSelect(titleDriver, dbDriver))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	// запрашиваем версию бд
	versions := getVersionForDb(d)
	v, err := console.Select(console.NewDataForSelect(titleVersion, versions))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	// запрашиваем имя пользователя для бд
	u, err := console.Input(console.NewDataForInput(titleUsername, "root"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	// запрашиваем пароль для бд
	p, err := console.Input(console.NewDataForInput(titlePassword, "root"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	// запрашиваем название бд
	dn, err := console.Input(console.NewDataForInput(titleDbName, "db_name"))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return &Data{
		Driver: d,
		Version: v,
		UserName: u,
		Password: p,
		DbName: dn,
	}
}

func getVersionForDb(driver string) []string {
	if driver == mysqlDriver {
		return mysqlVersion
	}

	return pgsqlVersion
}
