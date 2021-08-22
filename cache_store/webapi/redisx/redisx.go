// brian taylor vann
// redisx

package redisx

import (
	"errors"
	"fmt"

	"webapi/details"

	"github.com/gomodule/redigo/redis"
)

const (
	DELIMITER = ":"
)

var (
	errNilConfig = errors.New(
		"redix.Create() - nil config provided",
	)
	errOneOrFewerArgs = errors.New("1 or fewer arguments provided")

	pool, errPool = createRedisPool(&details.Details.Cache)
)

func createRedisPool(config *details.CacheDetails) (*redis.Pool, error) {
	if config == nil {
		return nil, errNilConfig
	}

	redisAddress := fmt.Sprint(config.Host, DELIMITER, config.Port)

	pool := redis.Pool{
		MaxIdle:     int(config.MaxIdle),
		IdleTimeout: config.IdleTimeout,
		MaxActive:   int(config.MaxActive),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				config.Protocol,
				redisAddress,
			)

			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	return &pool, nil
}

func Exec(args *[]interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	if errPool != nil {
		return nil, errPool
	}

	if len(*args) < 2 {
		return nil, errOneOrFewerArgs
	}

	conn := pool.Get()
	defer conn.Close()

	upckdArgs := *args
	firstArg := fmt.Sprint(upckdArgs[0])
	restArgs := upckdArgs[1:]

	return conn.Do(firstArg, restArgs...)
}
