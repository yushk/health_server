package convert

import (
	v1 "github.com/yushk/health_backend/apiserver/v1"
	"github.com/yushk/health_backend/pkg/pb"
)

func UserPb2V1(pbUser *pb.User) (v1User *v1.User) {
	if pbUser == nil {
		return
	}

	v1User = &v1.User{
		ID:    pbUser.Id,
		Name:  pbUser.Name,
		Ps:    pbUser.Ps,
		Role:  pbUser.Role,
		Phone: pbUser.Phone,
		Email: pbUser.Email,
	}
	return
}
