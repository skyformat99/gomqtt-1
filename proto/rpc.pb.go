// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	PubTextMsg
	PubAckMsg
	PubAckRet
	PubTextRet
	PubMsg
	PubRet
	BPushMsg
	BPushMsgs
	SPushMsg
	SPushMsgs
	PChatMsg
	PChatMsgs
	GChatMsg
	GChatMsgs
	LoginMsg
	LoginRet
	LogoutMsg
	LogoutRet
	Topic
	SubMsg
	SubRet
	UnSubMsg
	UnSubRet
	BPushRet
	SPushRet
	PChatRet
	GChatRet
	NickMsg
	NickRet
	ApnsMsg
	ApnsRet
	LabelMsg
	LabelRet
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

// 推送text消息结构
type PubTextMsg struct {
	Cid   int64  `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
	ToAcc []byte `protobuf:"bytes,2,opt,name=toAcc,proto3" json:"toAcc,omitempty"`
	Ttp   []byte `protobuf:"bytes,3,opt,name=ttp,proto3" json:"ttp,omitempty"`
	Qos   int32  `protobuf:"varint,4,opt,name=qos" json:"qos,omitempty"`
	Mid   []byte `protobuf:"bytes,5,opt,name=mid,proto3" json:"mid,omitempty"`
	Msg   []byte `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *PubTextMsg) Reset()                    { *m = PubTextMsg{} }
func (m *PubTextMsg) String() string            { return proto1.CompactTextString(m) }
func (*PubTextMsg) ProtoMessage()               {}
func (*PubTextMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 消息Ack
type PubAckMsg struct {
	Cid int64    `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
	Mid [][]byte `protobuf:"bytes,2,rep,name=mid,proto3" json:"mid,omitempty"`
}

func (m *PubAckMsg) Reset()                    { *m = PubAckMsg{} }
func (m *PubAckMsg) String() string            { return proto1.CompactTextString(m) }
func (*PubAckMsg) ProtoMessage()               {}
func (*PubAckMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 消息Ack消息结构返回结果
type PubAckRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *PubAckRet) Reset()                    { *m = PubAckRet{} }
func (m *PubAckRet) String() string            { return proto1.CompactTextString(m) }
func (*PubAckRet) ProtoMessage()               {}
func (*PubAckRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 推送text消息结构返回结果
type PubTextRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *PubTextRet) Reset()                    { *m = PubTextRet{} }
func (m *PubTextRet) String() string            { return proto1.CompactTextString(m) }
func (*PubTextRet) ProtoMessage()               {}
func (*PubTextRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PubMsg struct {
	Cid int64  `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
	Msg []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *PubMsg) Reset()                    { *m = PubMsg{} }
func (m *PubMsg) String() string            { return proto1.CompactTextString(m) }
func (*PubMsg) ProtoMessage()               {}
func (*PubMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

//
type PubRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *PubRet) Reset()                    { *m = PubRet{} }
func (m *PubRet) String() string            { return proto1.CompactTextString(m) }
func (*PubRet) ProtoMessage()               {}
func (*PubRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 广播
type BPushMsg struct {
	Tp  []byte `protobuf:"bytes,1,opt,name=tp,proto3" json:"tp,omitempty"`
	Rt  int64  `protobuf:"varint,2,opt,name=rt" json:"rt,omitempty"`
	Ttl int64  `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
	Mid []byte `protobuf:"bytes,4,opt,name=mid,proto3" json:"mid,omitempty"`
	Lb  []byte `protobuf:"bytes,8,opt,name=lb,proto3" json:"lb,omitempty"`
	Ij  []byte `protobuf:"bytes,9,opt,name=ij,proto3" json:"ij,omitempty"`
	Msg []byte `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *BPushMsg) Reset()                    { *m = BPushMsg{} }
func (m *BPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*BPushMsg) ProtoMessage()               {}
func (*BPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 广播打包结构
type BPushMsgs struct {
	Msgs []*BPushMsg `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *BPushMsgs) Reset()                    { *m = BPushMsgs{} }
func (m *BPushMsgs) String() string            { return proto1.CompactTextString(m) }
func (*BPushMsgs) ProtoMessage()               {}
func (*BPushMsgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BPushMsgs) GetMsgs() []*BPushMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// 单播
type SPushMsg struct {
	Tts [][]byte `protobuf:"bytes,1,rep,name=tts,proto3" json:"tts,omitempty"`
	Tus [][]byte `protobuf:"bytes,2,rep,name=tus,proto3" json:"tus,omitempty"`
	Ft  []byte   `protobuf:"bytes,3,opt,name=ft,proto3" json:"ft,omitempty"`
	Fu  []byte   `protobuf:"bytes,4,opt,name=fu,proto3" json:"fu,omitempty"`
	Rt  int64    `protobuf:"varint,5,opt,name=rt" json:"rt,omitempty"`
	Ttl int64    `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
	Mid []byte   `protobuf:"bytes,7,opt,name=mid,proto3" json:"mid,omitempty"`
	Lb  []byte   `protobuf:"bytes,8,opt,name=lb,proto3" json:"lb,omitempty"`
	Ij  []byte   `protobuf:"bytes,9,opt,name=ij,proto3" json:"ij,omitempty"`
	Msg []byte   `protobuf:"bytes,10,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *SPushMsg) Reset()                    { *m = SPushMsg{} }
func (m *SPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*SPushMsg) ProtoMessage()               {}
func (*SPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 单播打包结构
type SPushMsgs struct {
	Msgs []*SPushMsg `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *SPushMsgs) Reset()                    { *m = SPushMsgs{} }
func (m *SPushMsgs) String() string            { return proto1.CompactTextString(m) }
func (*SPushMsgs) ProtoMessage()               {}
func (*SPushMsgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SPushMsgs) GetMsgs() []*SPushMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// 私聊
type PChatMsg struct {
	Tp  []byte `protobuf:"bytes,1,opt,name=tp,proto3" json:"tp,omitempty"`
	Tu  []byte `protobuf:"bytes,2,opt,name=tu,proto3" json:"tu,omitempty"`
	Ft  []byte `protobuf:"bytes,3,opt,name=ft,proto3" json:"ft,omitempty"`
	Fu  []byte `protobuf:"bytes,4,opt,name=fu,proto3" json:"fu,omitempty"`
	Rt  int64  `protobuf:"varint,5,opt,name=rt" json:"rt,omitempty"`
	Ttl int64  `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
	Mid []byte `protobuf:"bytes,7,opt,name=mid,proto3" json:"mid,omitempty"`
	Lb  []byte `protobuf:"bytes,8,opt,name=lb,proto3" json:"lb,omitempty"`
	Ij  []byte `protobuf:"bytes,9,opt,name=ij,proto3" json:"ij,omitempty"`
	Msg []byte `protobuf:"bytes,10,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *PChatMsg) Reset()                    { *m = PChatMsg{} }
func (m *PChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*PChatMsg) ProtoMessage()               {}
func (*PChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// 私聊打包结构
type PChatMsgs struct {
	Msgs []*PChatMsg `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *PChatMsgs) Reset()                    { *m = PChatMsgs{} }
func (m *PChatMsgs) String() string            { return proto1.CompactTextString(m) }
func (*PChatMsgs) ProtoMessage()               {}
func (*PChatMsgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PChatMsgs) GetMsgs() []*PChatMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// 群聊
type GChatMsg struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *GChatMsg) Reset()                    { *m = GChatMsg{} }
func (m *GChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*GChatMsg) ProtoMessage()               {}
func (*GChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// 群聊打包结构
type GChatMsgs struct {
	Msgs []*GChatMsg `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *GChatMsgs) Reset()                    { *m = GChatMsgs{} }
func (m *GChatMsgs) String() string            { return proto1.CompactTextString(m) }
func (*GChatMsgs) ProtoMessage()               {}
func (*GChatMsgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GChatMsgs) GetMsgs() []*GChatMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// 登陆消息
type LoginMsg struct {
	Acc   []byte   `protobuf:"bytes,1,opt,name=acc,proto3" json:"acc,omitempty"`
	AppID []byte   `protobuf:"bytes,2,opt,name=appID,proto3" json:"appID,omitempty"`
	PT    int32    `protobuf:"varint,3,opt,name=pT" json:"pT,omitempty"`
	Cid   int64    `protobuf:"varint,4,opt,name=cid" json:"cid,omitempty"`
	Gip   []byte   `protobuf:"bytes,5,opt,name=gip,proto3" json:"gip,omitempty"`
	Ts    []*Topic `protobuf:"bytes,6,rep,name=ts" json:"ts,omitempty"`
}

func (m *LoginMsg) Reset()                    { *m = LoginMsg{} }
func (m *LoginMsg) String() string            { return proto1.CompactTextString(m) }
func (*LoginMsg) ProtoMessage()               {}
func (*LoginMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *LoginMsg) GetTs() []*Topic {
	if m != nil {
		return m.Ts
	}
	return nil
}

// 登陆消息返回消息
type LoginRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *LoginRet) Reset()                    { *m = LoginRet{} }
func (m *LoginRet) String() string            { return proto1.CompactTextString(m) }
func (*LoginRet) ProtoMessage()               {}
func (*LoginRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// 登出消息
type LogoutMsg struct {
	Cid int64 `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
}

func (m *LogoutMsg) Reset()                    { *m = LogoutMsg{} }
func (m *LogoutMsg) String() string            { return proto1.CompactTextString(m) }
func (*LogoutMsg) ProtoMessage()               {}
func (*LogoutMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// 登出消息返回消息
type LogoutRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *LogoutRet) Reset()                    { *m = LogoutRet{} }
func (m *LogoutRet) String() string            { return proto1.CompactTextString(m) }
func (*LogoutRet) ProtoMessage()               {}
func (*LogoutRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type Topic struct {
	Qos int32  `protobuf:"varint,1,opt,name=qos" json:"qos,omitempty"`
	Tp  []byte `protobuf:"bytes,2,opt,name=tp,proto3" json:"tp,omitempty"`
	Ty  int32  `protobuf:"varint,3,opt,name=ty" json:"ty,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto1.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

// 订阅主题消息
type SubMsg struct {
	Cid int64    `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
	Ts  []*Topic `protobuf:"bytes,2,rep,name=ts" json:"ts,omitempty"`
}

func (m *SubMsg) Reset()                    { *m = SubMsg{} }
func (m *SubMsg) String() string            { return proto1.CompactTextString(m) }
func (*SubMsg) ProtoMessage()               {}
func (*SubMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *SubMsg) GetTs() []*Topic {
	if m != nil {
		return m.Ts
	}
	return nil
}

// 订阅主题消息返回消息
type SubRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *SubRet) Reset()                    { *m = SubRet{} }
func (m *SubRet) String() string            { return proto1.CompactTextString(m) }
func (*SubRet) ProtoMessage()               {}
func (*SubRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

// 取消订阅主题消息
type UnSubMsg struct {
	Cid int64    `protobuf:"varint,1,opt,name=cid" json:"cid,omitempty"`
	Ts  []*Topic `protobuf:"bytes,2,rep,name=ts" json:"ts,omitempty"`
}

func (m *UnSubMsg) Reset()                    { *m = UnSubMsg{} }
func (m *UnSubMsg) String() string            { return proto1.CompactTextString(m) }
func (*UnSubMsg) ProtoMessage()               {}
func (*UnSubMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *UnSubMsg) GetTs() []*Topic {
	if m != nil {
		return m.Ts
	}
	return nil
}

// 取消订阅主题消息返回消息
type UnSubRet struct {
	R bool   `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
	M []byte `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (m *UnSubRet) Reset()                    { *m = UnSubRet{} }
func (m *UnSubRet) String() string            { return proto1.CompactTextString(m) }
func (*UnSubRet) ProtoMessage()               {}
func (*UnSubRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

// 广播返回消息
type BPushRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *BPushRet) Reset()                    { *m = BPushRet{} }
func (m *BPushRet) String() string            { return proto1.CompactTextString(m) }
func (*BPushRet) ProtoMessage()               {}
func (*BPushRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

// 单播返回消息
type SPushRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *SPushRet) Reset()                    { *m = SPushRet{} }
func (m *SPushRet) String() string            { return proto1.CompactTextString(m) }
func (*SPushRet) ProtoMessage()               {}
func (*SPushRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

// 私聊返回消息
type PChatRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *PChatRet) Reset()                    { *m = PChatRet{} }
func (m *PChatRet) String() string            { return proto1.CompactTextString(m) }
func (*PChatRet) ProtoMessage()               {}
func (*PChatRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

// 群播返回消息
type GChatRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *GChatRet) Reset()                    { *m = GChatRet{} }
func (m *GChatRet) String() string            { return proto1.CompactTextString(m) }
func (*GChatRet) ProtoMessage()               {}
func (*GChatRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

// 设置Nick
type NickMsg struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *NickMsg) Reset()                    { *m = NickMsg{} }
func (m *NickMsg) String() string            { return proto1.CompactTextString(m) }
func (*NickMsg) ProtoMessage()               {}
func (*NickMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

// 设置Nick返回消息
type NickRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *NickRet) Reset()                    { *m = NickRet{} }
func (m *NickRet) String() string            { return proto1.CompactTextString(m) }
func (*NickRet) ProtoMessage()               {}
func (*NickRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{28} }

// 设置Apns
type ApnsMsg struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ApnsMsg) Reset()                    { *m = ApnsMsg{} }
func (m *ApnsMsg) String() string            { return proto1.CompactTextString(m) }
func (*ApnsMsg) ProtoMessage()               {}
func (*ApnsMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{29} }

// 设置Apns返回消息
type ApnsRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ApnsRet) Reset()                    { *m = ApnsRet{} }
func (m *ApnsRet) String() string            { return proto1.CompactTextString(m) }
func (*ApnsRet) ProtoMessage()               {}
func (*ApnsRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{30} }

// Label
type LabelMsg struct {
	Msg [][]byte `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
}

func (m *LabelMsg) Reset()                    { *m = LabelMsg{} }
func (m *LabelMsg) String() string            { return proto1.CompactTextString(m) }
func (*LabelMsg) ProtoMessage()               {}
func (*LabelMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{31} }

// 设置Label返回消息
type LabelRet struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *LabelRet) Reset()                    { *m = LabelRet{} }
func (m *LabelRet) String() string            { return proto1.CompactTextString(m) }
func (*LabelRet) ProtoMessage()               {}
func (*LabelRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{32} }

func init() {
	proto1.RegisterType((*PubTextMsg)(nil), "proto.PubTextMsg")
	proto1.RegisterType((*PubAckMsg)(nil), "proto.PubAckMsg")
	proto1.RegisterType((*PubAckRet)(nil), "proto.PubAckRet")
	proto1.RegisterType((*PubTextRet)(nil), "proto.PubTextRet")
	proto1.RegisterType((*PubMsg)(nil), "proto.PubMsg")
	proto1.RegisterType((*PubRet)(nil), "proto.PubRet")
	proto1.RegisterType((*BPushMsg)(nil), "proto.BPushMsg")
	proto1.RegisterType((*BPushMsgs)(nil), "proto.BPushMsgs")
	proto1.RegisterType((*SPushMsg)(nil), "proto.SPushMsg")
	proto1.RegisterType((*SPushMsgs)(nil), "proto.SPushMsgs")
	proto1.RegisterType((*PChatMsg)(nil), "proto.PChatMsg")
	proto1.RegisterType((*PChatMsgs)(nil), "proto.PChatMsgs")
	proto1.RegisterType((*GChatMsg)(nil), "proto.GChatMsg")
	proto1.RegisterType((*GChatMsgs)(nil), "proto.GChatMsgs")
	proto1.RegisterType((*LoginMsg)(nil), "proto.LoginMsg")
	proto1.RegisterType((*LoginRet)(nil), "proto.LoginRet")
	proto1.RegisterType((*LogoutMsg)(nil), "proto.LogoutMsg")
	proto1.RegisterType((*LogoutRet)(nil), "proto.LogoutRet")
	proto1.RegisterType((*Topic)(nil), "proto.Topic")
	proto1.RegisterType((*SubMsg)(nil), "proto.SubMsg")
	proto1.RegisterType((*SubRet)(nil), "proto.SubRet")
	proto1.RegisterType((*UnSubMsg)(nil), "proto.UnSubMsg")
	proto1.RegisterType((*UnSubRet)(nil), "proto.UnSubRet")
	proto1.RegisterType((*BPushRet)(nil), "proto.BPushRet")
	proto1.RegisterType((*SPushRet)(nil), "proto.SPushRet")
	proto1.RegisterType((*PChatRet)(nil), "proto.PChatRet")
	proto1.RegisterType((*GChatRet)(nil), "proto.GChatRet")
	proto1.RegisterType((*NickMsg)(nil), "proto.NickMsg")
	proto1.RegisterType((*NickRet)(nil), "proto.NickRet")
	proto1.RegisterType((*ApnsMsg)(nil), "proto.ApnsMsg")
	proto1.RegisterType((*ApnsRet)(nil), "proto.ApnsRet")
	proto1.RegisterType((*LabelMsg)(nil), "proto.LabelMsg")
	proto1.RegisterType((*LabelRet)(nil), "proto.LabelRet")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	// 用户登陆相关接口
	Login(ctx context.Context, in *LoginMsg, opts ...grpc.CallOption) (*LoginRet, error)
	Logout(ctx context.Context, in *LogoutMsg, opts ...grpc.CallOption) (*LogoutRet, error)
	// 用户订阅相关
	Subscribe(ctx context.Context, in *SubMsg, opts ...grpc.CallOption) (*SubRet, error)
	UnSubscribe(ctx context.Context, in *UnSubMsg, opts ...grpc.CallOption) (*UnSubRet, error)
	// 用户发布
	Publish(ctx context.Context, in *PubMsg, opts ...grpc.CallOption) (*PubRet, error)
	PubText(ctx context.Context, in *PubTextMsg, opts ...grpc.CallOption) (*PubTextRet, error)
	PubAck(ctx context.Context, in *PubAckMsg, opts ...grpc.CallOption) (*PubAckRet, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Login(ctx context.Context, in *LoginMsg, opts ...grpc.CallOption) (*LoginRet, error) {
	out := new(LoginRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Logout(ctx context.Context, in *LogoutMsg, opts ...grpc.CallOption) (*LogoutRet, error) {
	out := new(LogoutRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Subscribe(ctx context.Context, in *SubMsg, opts ...grpc.CallOption) (*SubRet, error) {
	out := new(SubRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UnSubscribe(ctx context.Context, in *UnSubMsg, opts ...grpc.CallOption) (*UnSubRet, error) {
	out := new(UnSubRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/UnSubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Publish(ctx context.Context, in *PubMsg, opts ...grpc.CallOption) (*PubRet, error) {
	out := new(PubRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) PubText(ctx context.Context, in *PubTextMsg, opts ...grpc.CallOption) (*PubTextRet, error) {
	out := new(PubTextRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/PubText", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) PubAck(ctx context.Context, in *PubAckMsg, opts ...grpc.CallOption) (*PubAckRet, error) {
	out := new(PubAckRet)
	err := grpc.Invoke(ctx, "/proto.Rpc/PubAck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rpc service

type RpcServer interface {
	// 用户登陆相关接口
	Login(context.Context, *LoginMsg) (*LoginRet, error)
	Logout(context.Context, *LogoutMsg) (*LogoutRet, error)
	// 用户订阅相关
	Subscribe(context.Context, *SubMsg) (*SubRet, error)
	UnSubscribe(context.Context, *UnSubMsg) (*UnSubRet, error)
	// 用户发布
	Publish(context.Context, *PubMsg) (*PubRet, error)
	PubText(context.Context, *PubTextMsg) (*PubTextRet, error)
	PubAck(context.Context, *PubAckMsg) (*PubAckRet, error)
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMsg)
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
		return srv.(RpcServer).Login(ctx, req.(*LoginMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutMsg)
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
		return srv.(RpcServer).Logout(ctx, req.(*LogoutMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubMsg)
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
		return srv.(RpcServer).Subscribe(ctx, req.(*SubMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubMsg)
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
		return srv.(RpcServer).UnSubscribe(ctx, req.(*UnSubMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Publish(ctx, req.(*PubMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_PubText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubTextMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).PubText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/PubText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).PubText(ctx, req.(*PubTextMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_PubAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubAckMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).PubAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/PubAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).PubAck(ctx, req.(*PubAckMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
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
		{
			MethodName: "Publish",
			Handler:    _Rpc_Publish_Handler,
		},
		{
			MethodName: "PubText",
			Handler:    _Rpc_PubText_Handler,
		},
		{
			MethodName: "PubAck",
			Handler:    _Rpc_PubAck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x76, 0xec, 0x24, 0xd3, 0x02, 0x65, 0xc5, 0xc1, 0x32, 0x41, 0xaa, 0x4c, 0x55, 0x8a,
	0x8a, 0x0a, 0x94, 0x0b, 0x70, 0x2b, 0x20, 0x45, 0x48, 0x05, 0x45, 0x4e, 0x79, 0x80, 0xd8, 0x84,
	0xd4, 0xc5, 0x89, 0x8d, 0x77, 0x2d, 0x51, 0x89, 0x03, 0x07, 0x1e, 0x88, 0x03, 0x0f, 0xc2, 0x23,
	0x31, 0xfb, 0x67, 0xb3, 0x69, 0x2c, 0x0b, 0x4e, 0x9c, 0x32, 0x3b, 0x3f, 0xdf, 0x7e, 0xdf, 0xee,
	0x78, 0x36, 0x30, 0x2c, 0x8b, 0xe4, 0xa8, 0x28, 0x73, 0x96, 0x13, 0x57, 0xfc, 0x84, 0x5f, 0x01,
	0x26, 0x55, 0x7c, 0x36, 0xff, 0xc2, 0xde, 0xd2, 0x05, 0xd9, 0x01, 0x27, 0x49, 0x3f, 0xf8, 0xd6,
	0xae, 0x75, 0xe0, 0x44, 0xdc, 0x24, 0xb7, 0xc1, 0x65, 0xf9, 0x49, 0x92, 0xf8, 0x36, 0xfa, 0xb6,
	0x23, 0xb9, 0xe0, 0x79, 0x8c, 0x15, 0xbe, 0x23, 0x7c, 0xdc, 0xe4, 0x9e, 0xcf, 0x39, 0xf5, 0x7b,
	0xe8, 0x71, 0x23, 0x6e, 0x72, 0xcf, 0x12, 0xb1, 0x5c, 0x99, 0x83, 0xa6, 0xf0, 0xd0, 0x85, 0xef,
	0x29, 0x0f, 0x5d, 0x84, 0x8f, 0x60, 0x88, 0xbb, 0x9f, 0x24, 0x9f, 0x36, 0x6f, 0xae, 0x20, 0xec,
	0x5d, 0x47, 0x41, 0x84, 0xf7, 0x75, 0x41, 0x34, 0x67, 0x64, 0x1b, 0xac, 0x52, 0xa4, 0x0f, 0x22,
	0xab, 0xe4, 0xab, 0xa5, 0x62, 0x69, 0x2d, 0xc3, 0x83, 0x5a, 0x57, 0x57, 0xe6, 0x43, 0xf0, 0x30,
	0xb3, 0x9d, 0x00, 0x32, 0xb6, 0x1b, 0xc6, 0x7b, 0x22, 0xbb, 0x0b, 0xf3, 0x9b, 0x05, 0x83, 0x97,
	0x93, 0x8a, 0x9e, 0x73, 0xd8, 0x1b, 0x60, 0xe3, 0x59, 0x59, 0x22, 0x86, 0x16, 0x5f, 0x97, 0x4c,
	0xe4, 0x3a, 0x11, 0x5a, 0xf2, 0x30, 0x33, 0x71, 0x98, 0x0e, 0x3f, 0xcc, 0x4c, 0xeb, 0xee, 0x35,
	0x47, 0x87, 0x35, 0x59, 0xec, 0x0f, 0x24, 0x46, 0x16, 0xf3, 0x75, 0x7a, 0xe1, 0x0f, 0xe5, 0x3a,
	0xbd, 0xd0, 0x44, 0xfb, 0x0d, 0xd1, 0xc7, 0x30, 0xd4, 0x0c, 0x28, 0xb9, 0x07, 0x3d, 0xf4, 0x51,
	0x24, 0xe1, 0x1c, 0x6c, 0x1d, 0xdf, 0x94, 0x2d, 0x70, 0xa4, 0xe3, 0x91, 0x08, 0x86, 0x3f, 0x91,
	0xf4, 0x54, 0x93, 0x16, 0xa4, 0x64, 0x81, 0xb8, 0x61, 0x71, 0x9f, 0xac, 0xa2, 0xfa, 0x32, 0xd0,
	0xe4, 0x24, 0x3e, 0x32, 0xd5, 0x04, 0x68, 0x89, 0x75, 0xa5, 0x58, 0xa3, 0xa5, 0x84, 0xba, 0xeb,
	0x42, 0xbd, 0x2b, 0x42, 0xfb, 0x7f, 0x2d, 0x14, 0x0c, 0xa1, 0xd3, 0x0e, 0xa1, 0x53, 0x53, 0xe8,
	0x0f, 0x14, 0x3a, 0x79, 0x75, 0x3e, 0x63, 0x2d, 0xb7, 0xc3, 0x2a, 0x75, 0x93, 0x68, 0xfd, 0x3f,
	0x22, 0x35, 0xe3, 0x36, 0x91, 0x3a, 0xae, 0x44, 0x8e, 0x60, 0x30, 0xd6, 0x1a, 0x15, 0x9e, 0x65,
	0xe0, 0x8d, 0x3b, 0xf0, 0xc6, 0x26, 0xde, 0x77, 0x3c, 0xb4, 0xd3, 0x7c, 0x91, 0xae, 0x14, 0xe0,
	0x0c, 0x67, 0x82, 0x02, 0x44, 0x93, 0xcf, 0x89, 0x59, 0x51, 0xbc, 0x79, 0xad, 0xe7, 0x84, 0x58,
	0x70, 0x61, 0xc5, 0x99, 0x38, 0x3c, 0x37, 0x42, 0x4b, 0x7f, 0x61, 0x3d, 0xe3, 0x0b, 0x5b, 0xa4,
	0x85, 0x9e, 0x12, 0x68, 0x92, 0x11, 0x5e, 0x00, 0xc5, 0xf3, 0xe3, 0x5c, 0xb6, 0x15, 0x97, 0xb3,
	0xbc, 0x48, 0x13, 0xbc, 0x0e, 0x1a, 0xee, 0x2b, 0x16, 0x5d, 0x5f, 0xe0, 0x5d, 0x18, 0x62, 0x5e,
	0x5e, 0x6d, 0x1e, 0x6b, 0x7c, 0x8e, 0xc8, 0x70, 0x17, 0xce, 0x73, 0x70, 0xc5, 0xe6, 0x7a, 0xc0,
	0x59, 0xcd, 0x80, 0x93, 0x9d, 0x63, 0x1b, 0x9d, 0x73, 0xa9, 0xc5, 0xb2, 0xcb, 0xf0, 0x19, 0x78,
	0xd3, 0xb6, 0xc1, 0x22, 0x45, 0xda, 0x2d, 0x22, 0xf7, 0x44, 0x65, 0x17, 0xb5, 0x17, 0x30, 0x78,
	0xbf, 0xfa, 0xc7, 0x1d, 0xf6, 0x55, 0x6d, 0xd7, 0x1e, 0x23, 0x35, 0xc7, 0x78, 0xde, 0xd5, 0x2e,
	0x1a, 0xa9, 0x81, 0xd1, 0x1a, 0x15, 0x3d, 0xd9, 0x1a, 0x1d, 0xb7, 0x47, 0xef, 0x40, 0xff, 0x5d,
	0x5a, 0x3f, 0x0b, 0x9b, 0x83, 0xad, 0x95, 0x27, 0xc5, 0x8a, 0xb6, 0x56, 0xf2, 0x60, 0x2b, 0xa3,
	0xd3, 0x59, 0x3c, 0xcf, 0x8c, 0x52, 0x67, 0x3d, 0xba, 0xb1, 0xf6, 0xf8, 0x97, 0x0d, 0x4e, 0x54,
	0x24, 0xe4, 0x10, 0x5c, 0xd1, 0x9e, 0x44, 0x7f, 0x45, 0xfa, 0x93, 0x09, 0x0c, 0x07, 0x82, 0x84,
	0xd7, 0xc8, 0x11, 0x78, 0xb2, 0x09, 0xc9, 0x4e, 0x13, 0x94, 0x2d, 0x1b, 0x98, 0x1e, 0x99, 0x7f,
	0x88, 0x93, 0xae, 0x8a, 0x69, 0x52, 0xa6, 0xf1, 0x9c, 0x5c, 0xd7, 0xb3, 0x4d, 0x34, 0x40, 0xf0,
	0xc7, 0x52, 0x26, 0x3f, 0x81, 0x2d, 0x71, 0xc3, 0x2a, 0x5d, 0x6f, 0xaf, 0x3b, 0x26, 0x30, 0x1c,
	0xb2, 0xe4, 0x01, 0xf4, 0xf1, 0x6d, 0xcb, 0x52, 0x7a, 0x5e, 0xa3, 0x4f, 0x4c, 0xf4, 0x49, 0x83,
	0xde, 0x57, 0xcf, 0x2b, 0xb9, 0xd5, 0xc4, 0xd4, 0xdf, 0x88, 0x60, 0xcd, 0x55, 0xab, 0x95, 0x4f,
	0x77, 0xad, 0xb6, 0x7e, 0xfa, 0x03, 0xd3, 0x23, 0xf2, 0x63, 0x4f, 0xb8, 0x9e, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x93, 0xbe, 0x7c, 0xb4, 0x08, 0x00, 0x00,
}
