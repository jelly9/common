package redis
import (
    "github.com/go-redis/redis"
)

type Client struct {
    client *redis.Client
}

type Options struct {
    Addr string
}

func NewRedisClient(opt *Options) *Client{
    tmp := &redis.Options{
	Addr: opt.Addr,
    }
    return &Client{
	client: redis.NewClient(tmp),
    }
}

func (rc *Client) Set (key, value string) *redis.StringCmd {
    cmd := redis.NewStringCmd("set", key, value)
    rc.client.Process(cmd)
    return cmd
}

func (rc *Client) Get (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("get", key)
    rc.client.Process(cmd)
    return cmd
}

func (rc *Client) Incr (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("Incr", key)
    rc.client.Process(cmd)
    return cmd
}

func (rc *Client) Del (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("Del", key)
    rc.client.Process(cmd)
    return cmd
}