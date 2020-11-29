package ds

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/v-blazhko/blazhko/backend/api"
	"os"
)

type DS struct {
	dbMap *gorp.DbMap
}

func (ds *DS) Close() error {
	return ds.dbMap.Db.Close()
}

func NewStorage(db *sql.DB) *DS {
	dbMap := &gorp.DbMap{
		Db: db, Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	return &DS{
		dbMap: dbMap,
	}
}

func NewDB() (*sql.DB, error) {
	var (
		dBPort     string
		dBAddress  string
		dBUser     string
		dBPassword string
		dBName     string
	)
	if dBPort = os.Getenv("MYSQL_PORT"); dBPort == "" {
		dBPort = "3306"
	}
	if dBAddress = os.Getenv("MYSQL_ADDRESS"); dBAddress == "" {
		dBAddress = "localhost"
	}
	if dBName = os.Getenv("MYSQL_DATABASE"); dBName == "" {
		dBName = "db"
	}
	if dBUser = os.Getenv("MYSQL_USER"); dBUser == "" {
		dBUser = "user"
	}
	if dBPassword = os.Getenv("MYSQL_PASSWORD"); dBPassword == "" {
		dBPassword = "password"
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dBUser, dBPassword, dBAddress, dBPort, dBName))

	if err != nil {
		fmt.Print(err)
	}
	return db, err
}

func (ds *DS) CreateNewClientContact(c api.ClientContact) error {
	query := "INSERT INTO contact(client_name, email, message) VALUES(?, ?, ?)"
	params := []interface{}{c.Name, c.Email, c.Message}

	_, err := ds.dbMap.Exec(query, params...)
	return err
}
