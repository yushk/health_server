package mongodb

import (
	"context"
	"fmt"

	config2 "github.com/yushk/health_backend/pkg/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Mongo数据库名和表名定义
const (
	dbName         = "health"
	collectionUser = "user" // 用户表
	collectionAuth = "auth" // 用户认证(含认证信息)
)

type mongoDB struct {
	client *mongo.Client
}

func NewClient() (*mongoDB, error) {
	mongoURI := config2.GetString(config2.DBAddress)
	username := config2.GetString(config2.DBUsername)
	password := config2.GetString(config2.DBPassword)
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	if username != "" {
		auth := options.Credential{
			Username: username,
			Password: password,
		}
		clientOptions.SetAuth(auth)
	}
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logrus.Fatalln(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	// Create index
	err = createEnsureIndex(client)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln("Connected to MongoDB!")

	return &mongoDB{
		client: client,
	}, nil
}

func (m *mongoDB) Close() error {
	if m.client != nil {
		return m.client.Disconnect(context.TODO())
	}
	return nil
}

// DB 数据库句柄
func (m *mongoDB) DB(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return m.client.Database(name)
}

// C 集合句柄
func (m *mongoDB) C(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.client.Database(dbName).Collection(name, opts...)
}

// CAuth 认证集合句柄
func (m *mongoDB) CAuth() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionAuth)
}

// CUser 用户集合句柄
func (m *mongoDB) CUser() *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionUser)
}

// createEnsureIndex  创建唯一索引
func createEnsureIndex(client *mongo.Client) (err error) {
	collections := map[string][]string{
		collectionUser: {"name"},
	}

	for collect, rowArr := range collections {
		switch len(rowArr) {
		case 1:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		case 2:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
						{Key: rowArr[1], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		case 3:
			_, err = client.Database(dbName).Collection(collect).Indexes().CreateOne(
				context.Background(),
				mongo.IndexModel{
					Keys: bsonx.Doc{
						{Key: rowArr[0], Value: bsonx.Int32(1)},
						{Key: rowArr[1], Value: bsonx.Int32(1)},
						{Key: rowArr[2], Value: bsonx.Int32(1)},
					},
					Options: options.Index().SetUnique(true),
				},
			)
		default:
			err = fmt.Errorf("should implement more")
		}
		if err != nil {
			logrus.Fatalln(err)
		}
	}
	return
}
