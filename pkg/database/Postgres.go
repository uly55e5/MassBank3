package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MassBank/MassBank3/pkg/massbank"
	_ "github.com/lib/pq"
	"strconv"
)

type PostgresSQLDB struct {
	user       string
	dbname     string
	password   string
	host       string
	port       uint
	connString string
	database   *sql.DB
}

func NewPostgresSQLDb(config DBConfig) *PostgresSQLDB {
	var postgres = PostgresSQLDB{
		user:       config.DbUser,
		dbname:     config.DbName,
		password:   config.DbPwd,
		host:       config.DbHost,
		port:       config.DbPort,
		connString: "",
		database:   nil,
	}
	postgres.connString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPwd, config.DbName)
	return &postgres
}

func (p *PostgresSQLDB) Connect() error {
	db, err := sql.Open("postgres", p.connString)
	if err != nil {
		return err
	}
	p.database = db
	return p.init()
}

func (p *PostgresSQLDB) init() error {
	conn, err := p.database.Conn(context.Background())
	println(conn)
	if err := p.database.Ping(); err != nil {
		return err
	}
	q := "DROP TABLE massbank;"
	_, err = p.database.Exec(q)

	q = "DROP TABLE metadata;"
	_, err = p.database.Exec(q)
	q = `
		CREATE TABLE IF NOT EXISTS metadata
			(id INT GENERATED ALWAYS AS IDENTITY,
			commit char(40),
			timestamp timestamp NOT NULL,
			version varchar(10) NOT NULL,
		    PRIMARY KEY (id),
		    UNIQUE (commit,timestamp,version))
		
	`
	_, err = p.database.Exec(q)
	if err != nil {
		return err
	}
	q = `
		CREATE TABLE IF NOT EXISTS massbank 
			(id INT GENERATED ALWAYS AS IDENTITY, 
			 filename VARCHAR, 
			 document jsonb,
			 metadataId INT,
			 FOREIGN KEY  (metadataId) REFERENCES metadata(id), 
			 PRIMARY KEY (id))
 	`
	_, err = p.database.Exec(q)
	return err
}

func (p *PostgresSQLDB) Disconnect() error {
	if p.database == nil {
		return errors.New("database not set")
	}
	return p.database.Close()
}

func (p *PostgresSQLDB) GetRecord(s *string) (*massbank.Massbank, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresSQLDB) GetRecords(filters Filters, limit uint64) ([]*massbank.Massbank, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresSQLDB) UpdateMetadata(meta *massbank.MbMetaData) (string, error) {
	var id = 0
	err := p.database.QueryRow(
		"INSERT INTO metadata(commit,timestamp,version) VALUES  ($1,$2,$3)  "+
			"ON CONFLICT DO NOTHING RETURNING id;", meta.Commit, meta.Timestamp, meta.Version).Scan(&id)
	return strconv.Itoa(id), err

}

func (p *PostgresSQLDB) AddRecord(record *massbank.Massbank, metaDataId string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresSQLDB) AddRecords(records []*massbank.Massbank, metaDataId string) error {
	tx, err := p.database.Begin()
	if err != nil {
		return err
	}
	pStmt, err := tx.Prepare("INSERT INTO massbank(filename,document,metadataid) VALUES ($1,$2,$3);")
	if err != nil {
		tx.Rollback()
		return err
	}
	mid, err := strconv.ParseInt(metaDataId, 10, 64)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, record := range records {
		js, err := json.Marshal(record)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = pStmt.Exec(record.Metadata.FileName, js, mid)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	return err
}

func (p *PostgresSQLDB) UpdateRecord(record *massbank.Massbank, metaDataId string, upsert bool) (uint64, uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresSQLDB) UpdateRecords(records []*massbank.Massbank, metaDataId string, upsert bool) (uint64, uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresSQLDB) IsEmpty() (bool, error) {
	var count uint
	res := p.database.QueryRow(`SELECT COUNT(*) FROM massbank;`)
	err := res.Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
