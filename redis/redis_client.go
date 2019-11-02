package redis
import (
    "github.com/go-redis/redis"
)

type Options struct {
    Addr    string
    expired int64
}

type Client struct {
    client  *redis.Client
    expired int
}

func NewClient(opt *Options) *Client{
    tmp := &redis.Options{
	    Addr: opt.Addr,
    }
    return &Client{
        client: redis.NewClient(tmp),
        expired: opt.expired,
    }
}

func (c *Client) Set (key, value string, expired ...int) *redis.StringCmd {
    tm := c.expired
    if len(expired) != 0 {
        tm = expired[0]
    }
    cmd := redis.NewStringCmd("set", key, value, "EX", tm)
    c.client.Process(cmd)
    return cmd
}

func (c *Client) Get (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("get", key)
    c.client.Process(cmd)
    return cmd
}

func (c *Client) Incr (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("Incr", key)
    c.client.Process(cmd)
    return cmd
}

func (c *Client) Del (key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("Del", key)
    c.client.Process(cmd)
    return cmd
}
