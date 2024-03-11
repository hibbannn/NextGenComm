package cache

import "context"

func (c *Cache) SaveUserSessionToRedis(ctx context.Context, sessionKey string, encryptedUserData string) error {
	return c.redis.Set(ctx, sessionKey, encryptedUserData, 0)
}
