package dao

import (
	"github.com/sillyhatxu/mysql-client"
	"github.com/sirupsen/logrus"
)

var client *dbclient.MysqlClient

func Initial(dataSourceName, ddlPath string) error {
	logrus.Infof("initial db client;dataSourceName : %s", dataSourceName)
	mysqlClient, err := dbclient.NewMysqlClient(dataSourceName, dbclient.DDLPath(ddlPath))
	if err != nil {
		return err
	}
	client = mysqlClient
	return nil
}
