package db

import (
	"database/sql"
	"fmt"
	"os"
	// "log"
)

type DbConnectionPool struct {
	size      int
	semaphore chan *sql.DB
	open      bool
}

func NewDbConnectionPool(maxConcurrency int) DbConnectionPool {
	return DbConnectionPool{
		size:      maxConcurrency,
		semaphore: make(chan *sql.DB, maxConcurrency),
		open:      false,
	}

}

func (pool *DbConnectionPool) Open(conf SqlConf) {
	for i := 0; i < pool.size; i++ {
		go func(i int) {
			// log.Printf("Spinning up db connection %d of %d", i, pool.size)
			data_dir := os.Getenv("DATA_HOME")
			dbURL := fmt.Sprintf("file:%s/%s", data_dir, conf.Sqlite.DbName)
			dbDriver := "sqlite3"

			db, err := sql.Open(dbDriver, dbURL)
			if err != nil {
				panic(err)
			}
			db.SetMaxOpenConns(1)
			pool.semaphore <- db
		}(i)
	}
	pool.open = true
}

func (pool DbConnectionPool) Acquire() *sql.DB {
	return <-pool.semaphore
}

func (pool DbConnectionPool) Release(db *sql.DB) {
	pool.semaphore <- db
}

func (pool *DbConnectionPool) Close() {
	close(pool.semaphore)
	for d := range pool.semaphore {
		d.Close()
	}
	pool.open = false
}
