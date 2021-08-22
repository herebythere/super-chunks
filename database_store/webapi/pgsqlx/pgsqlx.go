package pgsqlx

import (
	"context"
	"errors"
	"fmt"

	"webapi/details"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Statement struct {
	Sql    string
	Values []interface{}
}

const (
	connectionString = "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
)

var (
	errNilConfig = errors.New("pgsqlx.Create() - nil config provided")

	pool, errPool = create(&details.Details.DB)
)

func getConnectionStr(config *details.DBDetails) string {
	return fmt.Sprintf(
		connectionString,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)
}

func create(config *details.DBDetails) (*pgxpool.Pool, error) {
	if config == nil {
		return nil, errNilConfig
	}

	connStr := getConnectionStr(config)

	return pgxpool.Connect(context.Background(), connStr)
}

func Query(args *Statement, err error) (*[][]interface{}, error) {
	if err != nil {
		return nil, err
	}

	if errPool != nil {
		return nil, errPool
	}

	results, errResults := pool.Query(context.Background(), args.Sql, args.Values...)
	if errResults != nil {
		return nil, errResults
	}

	defer results.Close()

	var parsedRows [][]interface{}
	for results.Next() {
		values, errValues := results.Values()
		if errValues != nil {
			return nil, errValues
		}

		parsedRows = append(parsedRows, values)
	}

	return &parsedRows, nil
}
