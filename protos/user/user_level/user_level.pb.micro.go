// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_level.proto

package user_level

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserLevel service

type UserLevelService interface {
	//获取用户等级信息
	GetUserLevelInfo(ctx context.Context, in *In_GetUserLevelInfo, opts ...client.CallOption) (*Out_GetUserLevelInfo, error)
	//修改等级信息
	UpdateUserLevel(ctx context.Context, in *In_UpdateUserLevel, opts ...client.CallOption) (*Out_UpdateUserLevel, error)
	//获取等级列表
	GetUserLevels(ctx context.Context, in *In_GetUserLevels, opts ...client.CallOption) (*Out_GetUserLevels, error)
	//删除等级列表
	DeleteUserLevels(ctx context.Context, in *In_DeleteUserLevels, opts ...client.CallOption) (*Out_DeleteUserLevels, error)
	//新建等级信息
	CreateUserLevel(ctx context.Context, in *In_CreateUserLevel, opts ...client.CallOption) (*Out_CreateUserLevel, error)
}

type userLevelService struct {
	c    client.Client
	name string
}

func NewUserLevelService(name string, c client.Client) UserLevelService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user_level"
	}
	return &userLevelService{
		c:    c,
		name: name,
	}
}

func (c *userLevelService) GetUserLevelInfo(ctx context.Context, in *In_GetUserLevelInfo, opts ...client.CallOption) (*Out_GetUserLevelInfo, error) {
	req := c.c.NewRequest(c.name, "UserLevel.GetUserLevelInfo", in)
	out := new(Out_GetUserLevelInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLevelService) UpdateUserLevel(ctx context.Context, in *In_UpdateUserLevel, opts ...client.CallOption) (*Out_UpdateUserLevel, error) {
	req := c.c.NewRequest(c.name, "UserLevel.UpdateUserLevel", in)
	out := new(Out_UpdateUserLevel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLevelService) GetUserLevels(ctx context.Context, in *In_GetUserLevels, opts ...client.CallOption) (*Out_GetUserLevels, error) {
	req := c.c.NewRequest(c.name, "UserLevel.GetUserLevels", in)
	out := new(Out_GetUserLevels)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLevelService) DeleteUserLevels(ctx context.Context, in *In_DeleteUserLevels, opts ...client.CallOption) (*Out_DeleteUserLevels, error) {
	req := c.c.NewRequest(c.name, "UserLevel.DeleteUserLevels", in)
	out := new(Out_DeleteUserLevels)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLevelService) CreateUserLevel(ctx context.Context, in *In_CreateUserLevel, opts ...client.CallOption) (*Out_CreateUserLevel, error) {
	req := c.c.NewRequest(c.name, "UserLevel.CreateUserLevel", in)
	out := new(Out_CreateUserLevel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserLevel service

type UserLevelHandler interface {
	//获取用户等级信息
	GetUserLevelInfo(context.Context, *In_GetUserLevelInfo, *Out_GetUserLevelInfo) error
	//修改等级信息
	UpdateUserLevel(context.Context, *In_UpdateUserLevel, *Out_UpdateUserLevel) error
	//获取等级列表
	GetUserLevels(context.Context, *In_GetUserLevels, *Out_GetUserLevels) error
	//删除等级列表
	DeleteUserLevels(context.Context, *In_DeleteUserLevels, *Out_DeleteUserLevels) error
	//新建等级信息
	CreateUserLevel(context.Context, *In_CreateUserLevel, *Out_CreateUserLevel) error
}

func RegisterUserLevelHandler(s server.Server, hdlr UserLevelHandler, opts ...server.HandlerOption) error {
	type userLevel interface {
		GetUserLevelInfo(ctx context.Context, in *In_GetUserLevelInfo, out *Out_GetUserLevelInfo) error
		UpdateUserLevel(ctx context.Context, in *In_UpdateUserLevel, out *Out_UpdateUserLevel) error
		GetUserLevels(ctx context.Context, in *In_GetUserLevels, out *Out_GetUserLevels) error
		DeleteUserLevels(ctx context.Context, in *In_DeleteUserLevels, out *Out_DeleteUserLevels) error
		CreateUserLevel(ctx context.Context, in *In_CreateUserLevel, out *Out_CreateUserLevel) error
	}
	type UserLevel struct {
		userLevel
	}
	h := &userLevelHandler{hdlr}
	return s.Handle(s.NewHandler(&UserLevel{h}, opts...))
}

type userLevelHandler struct {
	UserLevelHandler
}

func (h *userLevelHandler) GetUserLevelInfo(ctx context.Context, in *In_GetUserLevelInfo, out *Out_GetUserLevelInfo) error {
	return h.UserLevelHandler.GetUserLevelInfo(ctx, in, out)
}

func (h *userLevelHandler) UpdateUserLevel(ctx context.Context, in *In_UpdateUserLevel, out *Out_UpdateUserLevel) error {
	return h.UserLevelHandler.UpdateUserLevel(ctx, in, out)
}

func (h *userLevelHandler) GetUserLevels(ctx context.Context, in *In_GetUserLevels, out *Out_GetUserLevels) error {
	return h.UserLevelHandler.GetUserLevels(ctx, in, out)
}

func (h *userLevelHandler) DeleteUserLevels(ctx context.Context, in *In_DeleteUserLevels, out *Out_DeleteUserLevels) error {
	return h.UserLevelHandler.DeleteUserLevels(ctx, in, out)
}

func (h *userLevelHandler) CreateUserLevel(ctx context.Context, in *In_CreateUserLevel, out *Out_CreateUserLevel) error {
	return h.UserLevelHandler.CreateUserLevel(ctx, in, out)
}