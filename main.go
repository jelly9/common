package main

import (
	"fmt"
	// "common/redis"
	"common/mysql"
)

type Test struct {
	Id		int			`xorm:"id"`
	Name	string		`xorm:"name"`
}

func main() {
	// opt := &redis.Options{
	// 	Addr: "39.105.149.213:6379",
	// }
	// client := redis.NewRedisClient(opt)
	// //client.Set("aaa", "10")
	// value, _ := client.Get("bbb").Result()
	// fmt.Println("value: ", value)

	opt := &mysql.Options{
		DriverName: "mysql",
		Source: "xiaoguodong:m9@tcp(39.105.149.213:3306)/recruit",
		DefaultTable: "test",
		ShowSQL: true,
	}
	dao := mysql.NewMySQLDao(opt)
	var res Test
	dao.NewSession().Where("id=?", 1).Get(&res)
	fmt.Printf("res: %+v\n", res)
}