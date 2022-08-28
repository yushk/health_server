package pb

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/yushk/health_backend/pkg/config"
	"google.golang.org/grpc"
)

type GRPCClient interface {
	Host() string
	Port() string
	User() UserManagerClient
	Close() error
}

type grpcClient struct {
	host string
	port string
	conn *grpc.ClientConn
	user UserManagerClient
}

func NewGRPCClient() (GRPCClient, error) {
	host := config.GetString(config.ManagerHost)
	port := config.GetString(config.ManagerPort)
	address := net.JoinHostPort(host, port)

	logrus.Infoln("address", address)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		logrus.Errorf("did not connect: %v", err)
		return nil, err
	}

	return &grpcClient{
		host: host,
		port: port,
		conn: conn,
		user: NewUserManagerClient(conn),
	}, nil
}

func (c *grpcClient) Host() string {
	return c.host
}

func (c *grpcClient) Port() string {
	return c.port
}

func (c *grpcClient) User() UserManagerClient {
	return c.user
}

func (c *grpcClient) Close() error {
	if c.conn != nil {
		logrus.Debugln("client: ", c.conn.GetState(), "->closed")
		return c.conn.Close()
	}
	return nil
}
