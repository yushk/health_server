package convert

import (
	v1 "github.com/yushk/health_backend/apiserver/v1"
	"github.com/yushk/health_backend/pkg/pb"
)

func UserV12Pb(v1User *v1.User) (pbUser *pb.User) {
	pbUser = &pb.User{
		Id:    v1User.ID,
		Name:  v1User.Name,
		Ps:    v1User.Ps,
		Role:  v1User.Role,
		Phone: v1User.Phone,
		Email: v1User.Email,
	}
	return
}
