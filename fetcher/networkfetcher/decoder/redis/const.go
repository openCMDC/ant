package redis

type redisCmdType byte

const (
	SIMPLE_STRING redisCmdType = '+'
	ERR_MSG       redisCmdType = '-'
	INTEGER       redisCmdType = ':'
	BUIK_STRING   redisCmdType = '$'
	ARRAY         redisCmdType = '*'

	redisName = "redis"
)
