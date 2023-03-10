package main

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) (resp *douyinuser.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinuser.CreateUserResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (resp *douyinuser.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinuser.CheckUserResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *douyinuser.GetUserRequest) (resp *douyinuser.GetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinuser.GetUserResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.User = user
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUserName implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUserName(ctx context.Context, req *douyinuser.MGetUserNameRequest) (resp *douyinuser.MGetUserNameResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinuser.MGetUserNameResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userNames, err := service.NewMGetUserNameService(ctx).MGetUserName(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Usernames = userNames
	return resp, nil
}
