package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type config struct {
	database.DBConfig
	gitRepo   string
	gitBranch string
	dataDir   string
}

const (
	DB_USER_DEFAULT           = "massbank3"
	DB_PASSWORD_DEFAULT       = "massbank3password"
	DB_HOST_DEFAULT           = "localhost"
	DB_PORT_DEFAULT           = "27017"
	DB_NAME_DEFAULT           = "massbank3"
	DB_CONN_STRING_DEFAULT    = ""
	MB_GIT_REPO_DEFAULT       = "https://github.com/MassBank/MassBank-data"
	MB_GIT_BRANCH_DEFAULT     = "main"
	MB_DATA_DIRECTORY_DEFAULT = ""
)

func main() {
	var userConfig = getConfig()
	var db database.MB3Database
	var err error
	db, err = database.NewMongoDB(userConfig.DBConfig)
	if err != nil {
		panic(err)
	}
	err = db.Connect()
	if err != nil {
		panic(err)
	}
	var mbfiles []*massbank.Massbank
	if len(userConfig.dataDir) > 0 {
		mbfiles, err = readDirectoryData(userConfig.dataDir)
		if err != nil {
			println(err.Error())
		}
	}
	if mbfiles == nil && len(userConfig.gitRepo) > 0 {
		mbfiles, err = readGitData(userConfig.gitRepo, userConfig.gitBranch)
		if err != nil {
			panic(err)
		}
	}
	empty, err := db.IsEmpty()
	if err != nil {
		println(err.Error())
	}
	if empty {
		err = db.AddRecords(mbfiles)
		if err != nil {
			panic(err)
		}
	} else {
		_, _, err := db.UpdateRecords(mbfiles, true)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		println("Could not add records: " + err.Error())
	}
	if mbfiles == nil {
		panic("No files found")
	}
}

func getEnv(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return fallback
}

func getConfig() config {
	var c = config{}
	c.DbUser = getEnv("DB_USER", DB_USER_DEFAULT)
	c.DbPwd = getEnv("DB_PASSWORD", DB_PASSWORD_DEFAULT)
	c.DbHost = getEnv("DB_HOST", DB_HOST_DEFAULT)
	c.DbName = getEnv("DB_NAME", DB_NAME_DEFAULT)
	c.DbConnStr = getEnv("DB_CONN_STRING", DB_CONN_STRING_DEFAULT)
	c.gitRepo = getEnv("MB_GIT_REPO", MB_GIT_REPO_DEFAULT)
	c.gitBranch = getEnv("MB_GIT_BRANCH", MB_GIT_BRANCH_DEFAULT)
	c.dataDir = getEnv("MB_DATA_DIRECTORY", MB_DATA_DIRECTORY_DEFAULT)
	var dbPortEnv = getEnv("DB_PORT", DB_PORT_DEFAULT)
	dbPort, err := strconv.ParseUint(dbPortEnv, 10, 16)
	if err != nil {
		panic(errors.New("Could not read port variable: DB_PORT=" + dbPortEnv))
	}
	c.DbPort = uint(dbPort)
	flag.StringVar(&c.DbUser, "user", c.DbUser, "database user name")
	flag.StringVar(&c.DbPwd, "pwd", c.DbPwd, "database user password")
	flag.StringVar(&c.DbHost, "host", c.DbHost, "database host")
	flag.StringVar(&c.DbName, "db", c.DbName, "database name")
	flag.UintVar(&c.DbPort, "port", c.DbPort, "database port")
	flag.StringVar(&c.DbConnStr, "connstr", c.DbConnStr, "database connection string")
	flag.StringVar(&c.gitRepo, "git", c.gitRepo, "git host")
	flag.StringVar(&c.gitBranch, "branch", c.gitBranch, "git branch")
	flag.StringVar(&c.dataDir, "dir", c.dataDir, "data directory")
	flag.Parse()
	if len(c.gitRepo) > 0 && len(c.dataDir) > 0 {
		println("Git repo and data directory are set. Using data directory as default and git repo as fallback.")
	}
	return c
}

func readDirectoryData(dir string) ([]*massbank.Massbank, error) {
	filesNames, err := filepath.Glob(dir + "/**/*.txt")
	if err != nil {
		return nil, err
	}
	var mbfiles = []*massbank.Massbank{}
	for _, name := range filesNames {
		file, err := os.Open(name)
		if err != nil {
			return nil, err
		}
		mb, err := massbank.ScanMbFile(file, name)
		mbfiles = append(mbfiles, mb)
	}
	return mbfiles, nil
}

func readGitData(repo string, branch string) ([]*massbank.Massbank, error) {
	c := http.Client{}
	var url = fmt.Sprintf("%v/archive/refs/heads/%v.zip", repo, branch)
	println("Downloading file " + url)
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	println("Download finished")
	if err != nil {
		return nil, err
	}
	zReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Panicln(err)
	}
	var mbfiles = []*massbank.Massbank{}
	for _, zFile := range zReader.File {
		if strings.HasSuffix(zFile.Name, ".txt") {
			file, err := zFile.Open()
			if err != nil {
				return nil, err
			}
			mb, err := massbank.ScanMbFile(file, zFile.Name)
			mbfiles = append(mbfiles, mb)
		}
	}
	return mbfiles, nil
}
