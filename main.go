package main

import (
	"fmt"
	"time"
	"common/redis"
	"common/mysql"
	"common/mongo"

	"github.com/globalsign/mgo/bson"
)

type Test struct {
	Id		int			`xorm:"id"`
	Name	string		`xorm:"name"`
}

// 简历表结构体
type Resume struct {
	ResumeID 			uint  		`xorm:"resume_id not null pk autoincr INT(8)"`
	Name 				string		`xorm:"name not null default '' VARCHAR(64)" json:"name" binding:"required"`
	IDNo 				string		`xorm:"id_no not null default '' VARCHAR(20)" json:"id_no" binding:"required"`
	Mobile 				string		`xorm:"mobile not null default '' VARCHAR(16)" json:"mobile" binding:"required"`
	Email 				string		`xorm:"email not null default '' VARCHAR(64)" json:"email" binding:"required"`
	School 				string		`xorm:"school not null default '' VARCHAR(128)" json:"school" binding:"required"`
	Major 				string		`xorm:"major not null default '' VARCHAR(128)" json:"major" binding:"required"`
	Post 				string		`xorm:"post not null default '' VARCHAR(128)" json:"post" binding:"required"`
	CertifyName 		string		`xorm:"certify_name not null default '' VARCHAR(64)" json:"certify_name" binding:"required"`
	CertifyRelation 	string		`xorm:"certify_relation not null default '' VARCHAR(64)" json:"certify_relation" binding:"required"`
	CertifyMobile		string		`xorm:"certify_mobile not null default '' VARCHAR(16)" json:"certify_mobile" binding:"required"`
	OSSAddr 			string		`xorm:"oss_addr not null default '' VARCHAR(255)" json:"oss_addr" binding:"required"`
	CreateTime			time.Time	`xorm:"ctime created"`
	UpdateTime			time.Time	`xorm:"utime updated"`
}

// 岗位列表信息
type PostCategory struct {
	Id			interface{}	`json:"_id" bson:"_id"`
	Name 		string		`json:"name" bson:"name"`
	Achieve		bool		`json:"achieve" bson:"achieve"`
}

func testRedis(){
	opt := redis.Options{
		Addr: "39.105.149.213:6379",
	}
	client := redis.NewClient(&opt)
	//client.Set("aaa", "10")
	value, _ := client.Get("bbb").Result()
	fmt.Println("value: ", value)	
}

func testMysql() {
	opt := mysql.Options{
		DriverName: "mysql",
		Source: "xiaoguodong:m9@tcp(39.105.149.213:3306)/recruit",
		DefaultTable: "test",
		ShowSQL: true,
	}
	dao := mysql.NewEngine(&opt)
	var res Test
	dao.NewSession().Where("id=?", 1).Get(&res)
	fmt.Printf("res: %+v\n", res)
}

func testMgo() {
	opt := mongo.Options{
		Url: "39.105.149.213",
		Database:"recruit",
		Collection: "post_categorys",
	}

	s := mongo.NewSession(&opt)
	result := make([]*PostCategory, 0)
	// query := bson.M{"achieve": false}
	query := bson.M{}
	err := s.Collection().Find(query).All(&result)
	if err != nil {
		fmt.Printf("Insert error: %v\n", err)
	}

	for k, v := range result {
		fmt.Printf("result[%d] %+v\n", k, v)
	}

	fmt.Println("done...")
}

func main() {
	testMgo()
}