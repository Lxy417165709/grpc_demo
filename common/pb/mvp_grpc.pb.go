// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MvpClient is the client API for Mvp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MvpClient interface {
	AddGoodClass(ctx context.Context, in *AddGoodClassReq, opts ...grpc.CallOption) (*AddGoodClassRes, error)
	AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*AddGoodRes, error)
	AddElement(ctx context.Context, in *AddElementReq, opts ...grpc.CallOption) (*AddElementRes, error)
	AddDesk(ctx context.Context, in *AddDeskReq, opts ...grpc.CallOption) (*AddDeskRes, error)
	OrderGood(ctx context.Context, in *OrderGoodReq, opts ...grpc.CallOption) (*OrderGoodRes, error)
	OpenDesk(ctx context.Context, in *OpenDeskReq, opts ...grpc.CallOption) (*OpenDeskRes, error)
	CloseDesk(ctx context.Context, in *CloseDeskReq, opts ...grpc.CallOption) (*CloseDeskRes, error)
	GetDesk(ctx context.Context, in *GetDeskReq, opts ...grpc.CallOption) (*GetDeskRes, error)
	FormExpense(ctx context.Context, in *FormExpenseReq, opts ...grpc.CallOption) (*FormExpenseRes, error)
	CheckOut(ctx context.Context, in *CheckOutReq, opts ...grpc.CallOption) (*CheckOutRes, error)
	GetAllGoodClasses(ctx context.Context, in *GetAllGoodClassesReq, opts ...grpc.CallOption) (*GetAllGoodClassesRes, error)
}

type mvpClient struct {
	cc grpc.ClientConnInterface
}

func NewMvpClient(cc grpc.ClientConnInterface) MvpClient {
	return &mvpClient{cc}
}

func (c *mvpClient) AddGoodClass(ctx context.Context, in *AddGoodClassReq, opts ...grpc.CallOption) (*AddGoodClassRes, error) {
	out := new(AddGoodClassRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddGoodClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*AddGoodRes, error) {
	out := new(AddGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddElement(ctx context.Context, in *AddElementReq, opts ...grpc.CallOption) (*AddElementRes, error) {
	out := new(AddElementRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddDesk(ctx context.Context, in *AddDeskReq, opts ...grpc.CallOption) (*AddDeskRes, error) {
	out := new(AddDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) OrderGood(ctx context.Context, in *OrderGoodReq, opts ...grpc.CallOption) (*OrderGoodRes, error) {
	out := new(OrderGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/OrderGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) OpenDesk(ctx context.Context, in *OpenDeskReq, opts ...grpc.CallOption) (*OpenDeskRes, error) {
	out := new(OpenDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/OpenDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) CloseDesk(ctx context.Context, in *CloseDeskReq, opts ...grpc.CallOption) (*CloseDeskRes, error) {
	out := new(CloseDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/CloseDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetDesk(ctx context.Context, in *GetDeskReq, opts ...grpc.CallOption) (*GetDeskRes, error) {
	out := new(GetDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) FormExpense(ctx context.Context, in *FormExpenseReq, opts ...grpc.CallOption) (*FormExpenseRes, error) {
	out := new(FormExpenseRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/FormExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) CheckOut(ctx context.Context, in *CheckOutReq, opts ...grpc.CallOption) (*CheckOutRes, error) {
	out := new(CheckOutRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/CheckOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetAllGoodClasses(ctx context.Context, in *GetAllGoodClassesReq, opts ...grpc.CallOption) (*GetAllGoodClassesRes, error) {
	out := new(GetAllGoodClassesRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllGoodClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MvpServer is the server API for Mvp service.
// All implementations must embed UnimplementedMvpServer
// for forward compatibility
type MvpServer interface {
	AddGoodClass(context.Context, *AddGoodClassReq) (*AddGoodClassRes, error)
	AddGood(context.Context, *AddGoodReq) (*AddGoodRes, error)
	AddElement(context.Context, *AddElementReq) (*AddElementRes, error)
	AddDesk(context.Context, *AddDeskReq) (*AddDeskRes, error)
	OrderGood(context.Context, *OrderGoodReq) (*OrderGoodRes, error)
	OpenDesk(context.Context, *OpenDeskReq) (*OpenDeskRes, error)
	CloseDesk(context.Context, *CloseDeskReq) (*CloseDeskRes, error)
	GetDesk(context.Context, *GetDeskReq) (*GetDeskRes, error)
	FormExpense(context.Context, *FormExpenseReq) (*FormExpenseRes, error)
	CheckOut(context.Context, *CheckOutReq) (*CheckOutRes, error)
	GetAllGoodClasses(context.Context, *GetAllGoodClassesReq) (*GetAllGoodClassesRes, error)
	mustEmbedUnimplementedMvpServer()
}

// UnimplementedMvpServer must be embedded to have forward compatible implementations.
type UnimplementedMvpServer struct {
}

func (UnimplementedMvpServer) AddGoodClass(context.Context, *AddGoodClassReq) (*AddGoodClassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGoodClass not implemented")
}
func (UnimplementedMvpServer) AddGood(context.Context, *AddGoodReq) (*AddGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGood not implemented")
}
func (UnimplementedMvpServer) AddElement(context.Context, *AddElementReq) (*AddElementRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddElement not implemented")
}
func (UnimplementedMvpServer) AddDesk(context.Context, *AddDeskReq) (*AddDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDesk not implemented")
}
func (UnimplementedMvpServer) OrderGood(context.Context, *OrderGoodReq) (*OrderGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderGood not implemented")
}
func (UnimplementedMvpServer) OpenDesk(context.Context, *OpenDeskReq) (*OpenDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenDesk not implemented")
}
func (UnimplementedMvpServer) CloseDesk(context.Context, *CloseDeskReq) (*CloseDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDesk not implemented")
}
func (UnimplementedMvpServer) GetDesk(context.Context, *GetDeskReq) (*GetDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesk not implemented")
}
func (UnimplementedMvpServer) FormExpense(context.Context, *FormExpenseReq) (*FormExpenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormExpense not implemented")
}
func (UnimplementedMvpServer) CheckOut(context.Context, *CheckOutReq) (*CheckOutRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOut not implemented")
}
func (UnimplementedMvpServer) GetAllGoodClasses(context.Context, *GetAllGoodClassesReq) (*GetAllGoodClassesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGoodClasses not implemented")
}
func (UnimplementedMvpServer) mustEmbedUnimplementedMvpServer() {}

// UnsafeMvpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MvpServer will
// result in compilation errors.
type UnsafeMvpServer interface {
	mustEmbedUnimplementedMvpServer()
}

func RegisterMvpServer(s grpc.ServiceRegistrar, srv MvpServer) {
	s.RegisterService(&_Mvp_serviceDesc, srv)
}

func _Mvp_AddGoodClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodClassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddGoodClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddGoodClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddGoodClass(ctx, req.(*AddGoodClassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddGood(ctx, req.(*AddGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddElementReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddElement(ctx, req.(*AddElementReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddDesk(ctx, req.(*AddDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_OrderGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).OrderGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/OrderGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).OrderGood(ctx, req.(*OrderGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_OpenDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).OpenDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/OpenDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).OpenDesk(ctx, req.(*OpenDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_CloseDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).CloseDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/CloseDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).CloseDesk(ctx, req.(*CloseDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetDesk(ctx, req.(*GetDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_FormExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormExpenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).FormExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/FormExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).FormExpense(ctx, req.(*FormExpenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_CheckOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).CheckOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/CheckOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).CheckOut(ctx, req.(*CheckOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetAllGoodClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGoodClassesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllGoodClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllGoodClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllGoodClasses(ctx, req.(*GetAllGoodClassesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mvp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Mvp",
	HandlerType: (*MvpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGoodClass",
			Handler:    _Mvp_AddGoodClass_Handler,
		},
		{
			MethodName: "AddGood",
			Handler:    _Mvp_AddGood_Handler,
		},
		{
			MethodName: "AddElement",
			Handler:    _Mvp_AddElement_Handler,
		},
		{
			MethodName: "AddDesk",
			Handler:    _Mvp_AddDesk_Handler,
		},
		{
			MethodName: "OrderGood",
			Handler:    _Mvp_OrderGood_Handler,
		},
		{
			MethodName: "OpenDesk",
			Handler:    _Mvp_OpenDesk_Handler,
		},
		{
			MethodName: "CloseDesk",
			Handler:    _Mvp_CloseDesk_Handler,
		},
		{
			MethodName: "GetDesk",
			Handler:    _Mvp_GetDesk_Handler,
		},
		{
			MethodName: "FormExpense",
			Handler:    _Mvp_FormExpense_Handler,
		},
		{
			MethodName: "CheckOut",
			Handler:    _Mvp_CheckOut_Handler,
		},
		{
			MethodName: "GetAllGoodClasses",
			Handler:    _Mvp_GetAllGoodClasses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/pb/mvp.proto",
}
