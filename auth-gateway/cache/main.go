package cache

var RedisClient *redis.Client

func ConnectRedisCache() {
	RedisClient = nil
}
