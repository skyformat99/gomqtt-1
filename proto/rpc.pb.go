// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	BPushMsg
	SPushMsg
	PChatMsg
	GChatMsg
	AccMsg
	TcMsg
	Reply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// 广播
type BPushMsg struct {
}

func (m *BPushMsg) Reset()                    { *m = BPushMsg{} }
func (m *BPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*BPushMsg) ProtoMessage()               {}
func (*BPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 单播
type SPushMsg struct {
	TTopics [][]byte `protobuf:"bytes,1,rep,name=tTopics,proto3" json:"tTopics,omitempty"`
	TOpids  [][]byte `protobuf:"bytes,2,rep,name=tOpids,proto3" json:"tOpids,omitempty"`
}

func (m *SPushMsg) Reset()                    { *m = SPushMsg{} }
func (m *SPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*SPushMsg) ProtoMessage()               {}
func (*SPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 私聊
type PChatMsg struct {
}

func (m *PChatMsg) Reset()                    { *m = PChatMsg{} }
func (m *PChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*PChatMsg) ProtoMessage()               {}
func (*PChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 群播
type GChatMsg struct {
}

func (m *GChatMsg) Reset()                    { *m = GChatMsg{} }
func (m *GChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*GChatMsg) ProtoMessage()               {}
func (*GChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 用户账户
type AccMsg struct {
	An     string `protobuf:"bytes,1,opt,name=an" json:"an,omitempty"`
	Un     string `protobuf:"bytes,2,opt,name=un" json:"un,omitempty"`
	ConVer int32  `protobuf:"varint,3,opt,name=conVer" json:"conVer,omitempty"`
	Gip    string `protobuf:"bytes,4,opt,name=gip" json:"gip,omitempty"`
}

func (m *AccMsg) Reset()                    { *m = AccMsg{} }
func (m *AccMsg) String() string            { return proto1.CompactTextString(m) }
func (*AccMsg) ProtoMessage()               {}
func (*AccMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 主题消息
type TcMsg struct {
}

func (m *TcMsg) Reset()                    { *m = TcMsg{} }
func (m *TcMsg) String() string            { return proto1.CompactTextString(m) }
func (*TcMsg) ProtoMessage()               {}
func (*TcMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Reply struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto1.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto1.RegisterType((*BPushMsg)(nil), "proto.BPushMsg")
	proto1.RegisterType((*SPushMsg)(nil), "proto.SPushMsg")
	proto1.RegisterType((*PChatMsg)(nil), "proto.PChatMsg")
	proto1.RegisterType((*GChatMsg)(nil), "proto.GChatMsg")
	proto1.RegisterType((*AccMsg)(nil), "proto.AccMsg")
	proto1.RegisterType((*TcMsg)(nil), "proto.TcMsg")
	proto1.RegisterType((*Reply)(nil), "proto.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	// 推送接口
	BPush(ctx context.Context, in *BPushMsg, opts ...grpc.CallOption) (*Reply, error)
	SPush(ctx context.Context, in *SPushMsg, opts ...grpc.CallOption) (*Reply, error)
	PChat(ctx context.Context, in *PChatMsg, opts ...grpc.CallOption) (*Reply, error)
	GChat(ctx context.Context, in *GChatMsg, opts ...grpc.CallOption) (*Reply, error)
	// 用户相关接口
	Login(ctx context.Context, in *AccMsg, opts ...grpc.CallOption) (*Reply, error)
	Logout(ctx context.Context, in *AccMsg, opts ...grpc.CallOption) (*Reply, error)
	// 用户订阅相关
	Subscribe(ctx context.Context, in *TcMsg, opts ...grpc.CallOption) (*Reply, error)
	UnSubscribe(ctx context.Context, in *TcMsg, opts ...grpc.CallOption) (*Reply, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) BPush(ctx context.Context, in *BPushMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/BPush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) SPush(ctx context.Context, in *SPushMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/SPush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) PChat(ctx context.Context, in *PChatMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/PChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GChat(ctx context.Context, in *GChatMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/GChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Login(ctx context.Context, in *AccMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Logout(ctx context.Context, in *AccMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Subscribe(ctx context.Context, in *TcMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UnSubscribe(ctx context.Context, in *TcMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/UnSubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rpc service

type RpcServer interface {
	// 推送接口
	BPush(context.Context, *BPushMsg) (*Reply, error)
	SPush(context.Context, *SPushMsg) (*Reply, error)
	PChat(context.Context, *PChatMsg) (*Reply, error)
	GChat(context.Context, *GChatMsg) (*Reply, error)
	// 用户相关接口
	Login(context.Context, *AccMsg) (*Reply, error)
	Logout(context.Context, *AccMsg) (*Reply, error)
	// 用户订阅相关
	Subscribe(context.Context, *TcMsg) (*Reply, error)
	UnSubscribe(context.Context, *TcMsg) (*Reply, error)
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_BPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BPushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).BPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/BPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).BPush(ctx, req.(*BPushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_SPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).SPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/SPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).SPush(ctx, req.(*SPushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_PChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).PChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/PChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).PChat(ctx, req.(*PChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/GChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GChat(ctx, req.(*GChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Login(ctx, req.(*AccMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Logout(ctx, req.(*AccMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TcMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Subscribe(ctx, req.(*TcMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TcMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UnSubscribe(ctx, req.(*TcMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BPush",
			Handler:    _Rpc_BPush_Handler,
		},
		{
			MethodName: "SPush",
			Handler:    _Rpc_SPush_Handler,
		},
		{
			MethodName: "PChat",
			Handler:    _Rpc_PChat_Handler,
		},
		{
			MethodName: "GChat",
			Handler:    _Rpc_GChat_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Rpc_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Rpc_Logout_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Rpc_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Rpc_UnSubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x69, 0x82, 0xd3, 0xe6, 0x52, 0x7e, 0xe4, 0x01, 0x05, 0x26, 0x94, 0x01, 0x82, 0x90,
	0x3a, 0xc0, 0xca, 0x02, 0x0c, 0x59, 0x40, 0x54, 0x4e, 0x60, 0x6f, 0x4c, 0x94, 0x46, 0x02, 0xdb,
	0x8a, 0x93, 0x81, 0x17, 0xe5, 0x79, 0xf0, 0x75, 0x6d, 0xa4, 0x8a, 0xf2, 0x33, 0xe5, 0x1c, 0xdf,
	0x4f, 0x9f, 0x74, 0x02, 0x71, 0xa7, 0xf8, 0x4c, 0x75, 0xb2, 0x97, 0x94, 0xd8, 0x4f, 0x0a, 0x30,
	0xb9, 0x9d, 0x0f, 0x7a, 0xf9, 0xa0, 0x9b, 0xf4, 0x1a, 0x26, 0x85, 0xcb, 0x34, 0x81, 0x71, 0x5f,
	0x4a, 0xd5, 0x72, 0x9d, 0x8c, 0x4e, 0xc2, 0x6c, 0xca, 0x7c, 0xa5, 0x87, 0x10, 0xf5, 0x8f, 0xaa,
	0x7d, 0xd1, 0x49, 0x60, 0x0f, 0xae, 0xa1, 0x69, 0x7e, 0xb7, 0x5c, 0xf4, 0x68, 0x32, 0x39, 0xf7,
	0x99, 0x41, 0x74, 0xc3, 0x39, 0x3a, 0xf7, 0x20, 0x58, 0x08, 0xa3, 0x1b, 0x65, 0x31, 0x33, 0x09,
	0xfb, 0x20, 0x8c, 0xc5, 0xf6, 0x41, 0xa0, 0x99, 0x4b, 0xf1, 0x5c, 0x77, 0x49, 0x68, 0xde, 0x08,
	0x73, 0x8d, 0x1e, 0x40, 0xd8, 0xb4, 0x2a, 0xd9, 0xb6, 0x20, 0xc6, 0x74, 0x0c, 0xa4, 0x44, 0x65,
	0x7a, 0x04, 0x84, 0xd5, 0xea, 0xf5, 0x1d, 0x99, 0x37, 0xdd, 0x38, 0x39, 0xc6, 0xcb, 0x8f, 0x00,
	0x42, 0xa6, 0x38, 0xcd, 0x80, 0xd8, 0x85, 0x74, 0x7f, 0xb5, 0x7c, 0xe6, 0xf7, 0x1e, 0x4f, 0xdd,
	0x83, 0x35, 0xa4, 0x5b, 0x48, 0x16, 0x6b, 0x64, 0xf1, 0x0b, 0x69, 0xb7, 0x7e, 0x91, 0x7e, 0xf9,
	0x26, 0x32, 0x5f, 0x23, 0xf3, 0x9f, 0xc8, 0x53, 0x20, 0xf7, 0xb2, 0x69, 0x05, 0xdd, 0x75, 0x87,
	0xd5, 0x5f, 0xfb, 0xc6, 0x9d, 0x41, 0x64, 0x38, 0x39, 0xf4, 0x7f, 0x81, 0xe7, 0x10, 0x17, 0x43,
	0xa5, 0x79, 0xd7, 0x56, 0x35, 0xf5, 0xc7, 0x72, 0x23, 0x7a, 0x01, 0x3b, 0x4f, 0xe2, 0x9f, 0x70,
	0x15, 0xd9, 0x7a, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x50, 0xcc, 0x06, 0xf4, 0x4d, 0x02, 0x00,
	0x00,
}
