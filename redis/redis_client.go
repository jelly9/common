package redis
import (
    "github.com/go-redis/redis"
)

type Options struct {
    Addr    string
    Timeout int
}

type Client struct {
    client  *redis.Client
    timeout int
}

func NewClient(opt *Options) *Client{
    tmp := &redis.Options{
	    Addr: opt.Addr,
    }
    return &Client{
        client: redis.NewClient(tmp),
        timeout: opt.Timeout,
    }
}

func (c *Client) Set (key, value string, timeout ...int) *redis.StringCmd {
    tm := c.timeout
    if len(timeout) != 0 {
        tm = timeout[0]
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