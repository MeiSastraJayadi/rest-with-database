package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-hclog"
)

type Connection struct {
	logger            hclog.Logger
	connection        string
	driver            string
	maxOpenConnection int
	idleConnection    int
	idleTime          time.Duration
	lifeTime          time.Duration
}

type UseDB struct {
	logger hclog.Logger
	db     *sql.DB
}

func (ub *UseDB) GetDB() *sql.DB {
	return ub.db
}

func NewUseDB(logger hclog.Logger, db *sql.DB) *UseDB {
	return &UseDB{
		logger: logger,
		db:     db,
	}
}

func NewConnection(logger hclog.Logger, connection string, driver string, openConnectionMax int, idleConn int, idleTime time.Duration, lifeTime time.Duration) *Connection {
	return &Connection{
		logger:            logger,
		driver:            driver,
		connection:        connection,
		maxOpenConnection: openConnectionMax,
		idleConnection:    idleConn,
		idleTime:          idleTime,
		lifeTime:          lifeTime,
	}
}

func (conn *Connection) CreateConnection() (*sql.DB, error) {
	db, err := sql.Open(conn.driver, conn.connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conn.maxOpenConnection)
	db.SetMaxIdleConns(conn.idleConnection)
	db.SetConnMaxIdleTime(conn.idleTime)
	db.SetConnMaxLifetime(conn.lifeTime)

	conn.logger.Info("Success create connection to database", "address", conn.connection)
	return db, nil
}
