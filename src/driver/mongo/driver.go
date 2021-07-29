package mongo

import (
	"context"
	"time"

	"anydbconnect/src/driver"

	"github.com/shzy2012/common/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
}

func NewDB() driver.DB {
	return &Client{}
}

func (x *Client) Connect(conn string) {

	var err error
	//初始化 MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//mongodb的用户名和密码是基于特定数据库的，而不是基于整个系统的。所有所有数据库db都需要设置密码
	//mongodb://youruser2:yourpassword2@localhost/yourdatabase
	connectString := conn
	log.Infoln("[mongo URI]=>", connectString)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	if err != nil {
		log.Fatalf("[mongo]=>%s\n", err)
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err == nil {
		log.Println("[mongo]=>int ok")
	} else {
		log.Fatalf("[mongo]=>int fail %s\n", err.Error())
	}
}
