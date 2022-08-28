package db

import (
	"context"
	"time"

	"github.com/yushk/health_backend/manager/internal/db/mongodb"
	"github.com/yushk/health_backend/pkg/pb"
)

// Type 数据库类型
type Type string

// 数据库类型枚举
const (
	MongoDB Type = "mongodb"
)

// 数据库连接默认的超时时间
const (
	DefaultTimeout = 10 * time.Second
)

// Database 数据库接口
type Database interface {
	// 用户管理接口
	CreateUser(ctx context.Context, user *pb.User) (*pb.User, error)
	GetUser(ctx context.Context, id string) (user *pb.User, err error)
	DeleteUser(ctx context.Context, id string) (user *pb.User, err error)
	Authenticate(ctx context.Context, username, password string) (ret bool)
	UpdateUser(ctx context.Context, id string, user *pb.User) (*pb.User, error)
	GetUsers(ctx context.Context, limit, skip int64, query string) (int64, []*pb.User, error)
	GetUserByName(ctx context.Context, name string) (user *pb.User, err error)
	ChangeAuthorization(ctx context.Context, name, password string) (err error)
	RemoveAuthorization(ctx context.Context, username string) (err error)

	Close() error
}

// New 创建数据库句柄
func New(name Type) (Database, error) {
	switch name {
	case MongoDB:
		return mongodb.NewClient()
	default:
		return mongodb.NewClient()
	}
}
