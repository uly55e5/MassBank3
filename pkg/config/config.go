package config

import (
	"errors"
	"flag"
	"github.com/MassBank/MassBank3/pkg/common"
	"github.com/MassBank/MassBank3/pkg/database"
	"log"
	"strconv"
)

type ToolConfig struct {
	database.DBConfig
	GitRepo   string
	GitBranch string
	DataDir   string
	Drop      bool
}

type ServerConfig struct {
	database.DBConfig
	ServerPort   uint
	CdkDepictUrl string
}

const (
	dbDefault              = "postgres"
	dbUserDefault          = "massbank3"
	dbPasswordDefault      = "massbank3password"
	dbHostDefault          = "localhost"
	dbPortDefault          = "0"
	dbNameDefault          = "massbank3"
	dbConnStringDefault    = ""
	mbGitRepoDefault       = "https://github.com/MassBank/MassBank-data"
	mbGitBranchDefault     = "main"
	mbDataDirectoryDefault = ""
	mbDropAllDefault       = "true"
	serverPortDefault      = "8080"
	cdkdepictUrlDefault    = "http://cdkdepict"
)

var toolConfig *ToolConfig = nil
var serverConfig *ServerConfig = nil

func GetToolConfig() ToolConfig {
	if toolConfig != nil {
		return *toolConfig
	}
	toolConfig = &ToolConfig{DBConfig: getDBConfig()}
	var err error
	toolConfig.GitRepo = common.GetEnv("MB_GIT_REPO", mbGitRepoDefault)
	toolConfig.GitBranch = common.GetEnv("MB_GIT_BRANCH", mbGitBranchDefault)
	var drop = common.GetEnv("MB_DROP_ALL", mbDropAllDefault)
	toolConfig.DataDir = common.GetEnv("MB_DATA_DIRECTORY", mbDataDirectoryDefault)
	toolConfig.Drop, err = strconv.ParseBool(drop)
	if err != nil {
		log.Println(err.Error())
	}
	flag.StringVar(&toolConfig.GitRepo, "git", toolConfig.GitRepo, "git repository. Overwrites environment variable MB_GIT_REPO")
	flag.StringVar(&toolConfig.GitBranch, "branch", toolConfig.GitBranch, "git branch. Overwrites environment variable MB_GIT_BRANCH")
	flag.StringVar(&toolConfig.DataDir, "dir", toolConfig.DataDir, "data directory. Overwrites environment variable MB_DATA_DIRECTORY")
	flag.BoolVar(&toolConfig.Drop, "dropall", toolConfig.Drop, "drop all data. Overwrites environment variable MB_DROP_ALL")
	flag.Parse()
	if len(toolConfig.GitRepo) > 0 && len(toolConfig.DataDir) > 0 {
		println("Git repo and data directory are set. Using data directory as default and git repo as fallback.")
	}
	return *toolConfig
}

func GetServerConfig() *ServerConfig {
	if serverConfig != nil {
		return serverConfig
	}
	serverConfig = &ServerConfig{DBConfig: getDBConfig()}
	var err error
	var serverPortEnv = common.GetEnv("MB3_SERVER_PORT", serverPortDefault)
	var serverPort uint64
	serverPort, err = strconv.ParseUint(serverPortEnv, 10, 16)
	if err != nil {
		panic(errors.New("Could not read port variable: DB_PORT=" + serverPortEnv))
	}
	serverConfig.ServerPort = uint(serverPort)
	serverConfig.CdkDepictUrl = common.GetEnv("CDKDEPICT_URL", cdkdepictUrlDefault)
	flag.StringVar(&serverConfig.CdkDepictUrl, "cdkdepict_url", serverConfig.CdkDepictUrl, "Base URL of the CDK Depict api server. Overwrites environment variable CDKDEPICT_URL")
	flag.UintVar(&serverConfig.ServerPort, "server_port", serverConfig.ServerPort, "Listen on this port. Overwrites environment variable SERVER_PORT")
	flag.Parse()
	return serverConfig
}

func getDBConfig() database.DBConfig {
	var c = database.DBConfig{}
	var err error
	c.DbUser = common.GetEnv("DB_USER", dbUserDefault)
	c.DbPwd = common.GetEnv("DB_PASSWORD", dbPasswordDefault)
	c.DbHost = common.GetEnv("DB_HOST", dbHostDefault)
	c.DbName = common.GetEnv("DB_NAME", dbNameDefault)
	c.DbConnStr = common.GetEnv("DB_CONN_STRING", dbConnStringDefault)
	var databaseType = common.GetEnv("DB_TYPE", dbDefault)
	var dbPortEnv = common.GetEnv("DB_PORT", dbPortDefault)
	var dbPort uint64
	dbPort, err = strconv.ParseUint(dbPortEnv, 10, 16)
	if err != nil {
		log.Panicln(errors.New("Could not read port variable: DB_PORT=" + dbPortEnv))
	}
	c.DbPort = uint(dbPort)
	flag.StringVar(&databaseType, "db_type", databaseType, "Database type must be postgres or mongodb. Overwrites environment variable DB_TYPE")
	flag.StringVar(&c.DbUser, "db_user", c.DbUser, "database user name. Overwrites environment variable DB_USER")
	flag.StringVar(&c.DbPwd, "db_pwd", c.DbPwd, "database user password. Overwrites environment variable DB_PASSWORD")
	flag.StringVar(&c.DbHost, "db_host", c.DbHost, "database host. Overwrites environment variable DB_HOST")
	flag.StringVar(&c.DbName, "db", c.DbName, "database name. Overwrites environment variable DB_NAME")
	flag.UintVar(&c.DbPort, "db_port", c.DbPort, "database port. Overwrites environment variable DB_PORT")
	flag.StringVar(&c.DbConnStr, "db_connstr", c.DbConnStr, "database connection string. Overwrites environment variable DB_CONN_STRING")
	flag.Parse()
	if databaseType == "postgres" {
		c.Database = database.Postgres
	} else {
		panic("Database must be postgres")
	}
	if c.DbPort == 0 {
		if c.Database == database.Postgres {
			c.DbPort = 5432
		} else {
			c.DbPort = 27017
		}
	}
	return c
}
