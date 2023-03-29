package manager

import (
	"stage/http_server/request"
	"stage/internal/redis"
)

func RedisSet(params request.KeyValueRequest) error {
	err := redis.RedisSet(params.Key, params.Value)
	return err
}

func RedisUnset(params request.KeyValueRequest) {

}
