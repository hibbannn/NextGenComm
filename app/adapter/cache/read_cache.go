package cache

import "context"

func (c *Cache) GetUserDetailFromRedisSession(ctx context.Context, sessionKey string) (string, error) {
	userDetail, err := c.redis.Get(ctx, sessionKey)
	if err != nil {
		return "", err
	}
	return userDetail, nil
}
