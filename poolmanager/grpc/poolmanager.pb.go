// Code generated by protoc-gen-go. DO NOT EDIT.
// source: poolmanager.proto

/*
Package poolmanager is a generated protocol buffer package.

It is generated from these files:
	poolmanager.proto

It has these top-level messages:
	Runner
	LBGroupId
	LBGroupMembership
	CapacitySnapshot
	CapacitySnapshotList
*/
package poolmanager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Runner struct {
	Address   string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	ClientKey []byte `protobuf:"bytes,2,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
}

func (m *Runner) Reset()                    { *m = Runner{} }
func (m *Runner) String() string            { return proto.CompactTextString(m) }
func (*Runner) ProtoMessage()               {}
func (*Runner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Runner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Runner) GetClientKey() []byte {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

type LBGroupId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *LBGroupId) Reset()                    { *m = LBGroupId{} }
func (m *LBGroupId) String() string            { return proto.CompactTextString(m) }
func (*LBGroupId) ProtoMessage()               {}
func (*LBGroupId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LBGroupId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LBGroupMembership struct {
	GroupId *LBGroupId `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Runners []*Runner  `protobuf:"bytes,2,rep,name=runners" json:"runners,omitempty"`
}

func (m *LBGroupMembership) Reset()                    { *m = LBGroupMembership{} }
func (m *LBGroupMembership) String() string            { return proto.CompactTextString(m) }
func (*LBGroupMembership) ProtoMessage()               {}
func (*LBGroupMembership) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LBGroupMembership) GetGroupId() *LBGroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *LBGroupMembership) GetRunners() []*Runner {
	if m != nil {
		return m.Runners
	}
	return nil
}

type CapacitySnapshot struct {
	GroupId        *LBGroupId `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	MemMbCommitted int32      `protobuf:"varint,2,opt,name=mem_mb_committed,json=memMbCommitted" json:"mem_mb_committed,omitempty"`
	MemMbTotal     int32      `protobuf:"varint,3,opt,name=mem_mb_total,json=memMbTotal" json:"mem_mb_total,omitempty"`
}

func (m *CapacitySnapshot) Reset()                    { *m = CapacitySnapshot{} }
func (m *CapacitySnapshot) String() string            { return proto.CompactTextString(m) }
func (*CapacitySnapshot) ProtoMessage()               {}
func (*CapacitySnapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CapacitySnapshot) GetGroupId() *LBGroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *CapacitySnapshot) GetMemMbCommitted() int32 {
	if m != nil {
		return m.MemMbCommitted
	}
	return 0
}

func (m *CapacitySnapshot) GetMemMbTotal() int32 {
	if m != nil {
		return m.MemMbTotal
	}
	return 0
}

type CapacitySnapshotList struct {
	Ts        *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=ts" json:"ts,omitempty"`
	LbId      string                     `protobuf:"bytes,2,opt,name=lb_id,json=lbId" json:"lb_id,omitempty"`
	Snapshots []*CapacitySnapshot        `protobuf:"bytes,3,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *CapacitySnapshotList) Reset()                    { *m = CapacitySnapshotList{} }
func (m *CapacitySnapshotList) String() string            { return proto.CompactTextString(m) }
func (*CapacitySnapshotList) ProtoMessage()               {}
func (*CapacitySnapshotList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CapacitySnapshotList) GetTs() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *CapacitySnapshotList) GetLbId() string {
	if m != nil {
		return m.LbId
	}
	return ""
}

func (m *CapacitySnapshotList) GetSnapshots() []*CapacitySnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func init() {
	proto.RegisterType((*Runner)(nil), "Runner")
	proto.RegisterType((*LBGroupId)(nil), "LBGroupId")
	proto.RegisterType((*LBGroupMembership)(nil), "LBGroupMembership")
	proto.RegisterType((*CapacitySnapshot)(nil), "CapacitySnapshot")
	proto.RegisterType((*CapacitySnapshotList)(nil), "CapacitySnapshotList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodePoolScaler service

type NodePoolScalerClient interface {
	AdvertiseCapacity(ctx context.Context, in *CapacitySnapshotList, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type nodePoolScalerClient struct {
	cc *grpc.ClientConn
}

func NewNodePoolScalerClient(cc *grpc.ClientConn) NodePoolScalerClient {
	return &nodePoolScalerClient{cc}
}

func (c *nodePoolScalerClient) AdvertiseCapacity(ctx context.Context, in *CapacitySnapshotList, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/NodePoolScaler/AdvertiseCapacity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodePoolScaler service

type NodePoolScalerServer interface {
	AdvertiseCapacity(context.Context, *CapacitySnapshotList) (*google_protobuf1.Empty, error)
}

func RegisterNodePoolScalerServer(s *grpc.Server, srv NodePoolScalerServer) {
	s.RegisterService(&_NodePoolScaler_serviceDesc, srv)
}

func _NodePoolScaler_AdvertiseCapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapacitySnapshotList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePoolScalerServer).AdvertiseCapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodePoolScaler/AdvertiseCapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePoolScalerServer).AdvertiseCapacity(ctx, req.(*CapacitySnapshotList))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodePoolScaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NodePoolScaler",
	HandlerType: (*NodePoolScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdvertiseCapacity",
			Handler:    _NodePoolScaler_AdvertiseCapacity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poolmanager.proto",
}

// Client API for RunnerManager service

type RunnerManagerClient interface {
	GetLBGroup(ctx context.Context, in *LBGroupId, opts ...grpc.CallOption) (*LBGroupMembership, error)
}

type runnerManagerClient struct {
	cc *grpc.ClientConn
}

func NewRunnerManagerClient(cc *grpc.ClientConn) RunnerManagerClient {
	return &runnerManagerClient{cc}
}

func (c *runnerManagerClient) GetLBGroup(ctx context.Context, in *LBGroupId, opts ...grpc.CallOption) (*LBGroupMembership, error) {
	out := new(LBGroupMembership)
	err := grpc.Invoke(ctx, "/RunnerManager/GetLBGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RunnerManager service

type RunnerManagerServer interface {
	GetLBGroup(context.Context, *LBGroupId) (*LBGroupMembership, error)
}

func RegisterRunnerManagerServer(s *grpc.Server, srv RunnerManagerServer) {
	s.RegisterService(&_RunnerManager_serviceDesc, srv)
}

func _RunnerManager_GetLBGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LBGroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).GetLBGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/GetLBGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).GetLBGroup(ctx, req.(*LBGroupId))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunnerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RunnerManager",
	HandlerType: (*RunnerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLBGroup",
			Handler:    _RunnerManager_GetLBGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poolmanager.proto",
}

func init() { proto.RegisterFile("poolmanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xed, 0x6e, 0x6c, 0x63, 0x6e, 0x6b, 0x68, 0xc6, 0x0f, 0x96, 0x14, 0x31, 0x2e, 0x08, 0xc1,
	0x87, 0x09, 0xc4, 0x1f, 0x20, 0xb5, 0x4a, 0x29, 0x36, 0x22, 0xdb, 0xe2, 0x9b, 0x84, 0xd9, 0xcc,
	0x75, 0x3b, 0xb8, 0xb3, 0x33, 0xcc, 0xdc, 0x08, 0xf9, 0x01, 0x82, 0x3f, 0x5b, 0x32, 0x3b, 0xab,
	0x92, 0xf8, 0xd0, 0xc7, 0x7b, 0xee, 0xd9, 0xb3, 0x67, 0xce, 0xb9, 0x30, 0xb2, 0xc6, 0xd4, 0x5a,
	0x34, 0xa2, 0x42, 0xc7, 0xad, 0x33, 0x64, 0xc6, 0x2f, 0x2a, 0x63, 0xaa, 0x1a, 0x67, 0x61, 0x2a,
	0xd7, 0xdf, 0x66, 0xa4, 0x34, 0x7a, 0x12, 0xda, 0x46, 0xc2, 0xd9, 0x2e, 0x01, 0xb5, 0xa5, 0x4d,
	0xbb, 0xcc, 0xcf, 0xe1, 0xa8, 0x58, 0x37, 0x0d, 0x3a, 0x96, 0x41, 0x5f, 0x48, 0xe9, 0xd0, 0xfb,
	0x2c, 0x99, 0x24, 0xd3, 0x41, 0xd1, 0x8d, 0xec, 0x39, 0xc0, 0xaa, 0x56, 0xd8, 0xd0, 0xf2, 0x3b,
	0x6e, 0xb2, 0x74, 0x92, 0x4c, 0x4f, 0x8a, 0x41, 0x8b, 0x7c, 0xc4, 0x4d, 0x7e, 0x06, 0x83, 0xeb,
	0x77, 0x97, 0xce, 0xac, 0xed, 0x95, 0x64, 0x43, 0x48, 0x95, 0x8c, 0x02, 0xa9, 0x92, 0xf9, 0x57,
	0x18, 0xc5, 0xe5, 0x02, 0x75, 0x89, 0xce, 0xdf, 0x29, 0xcb, 0x5e, 0xc1, 0xc3, 0x6a, 0x0b, 0x2d,
	0x23, 0xf5, 0x78, 0x0e, 0xfc, 0x8f, 0x44, 0xd1, 0xaf, 0xa2, 0xd6, 0x4b, 0xe8, 0xbb, 0xe0, 0xcd,
	0x67, 0xe9, 0xa4, 0x37, 0x3d, 0x9e, 0xf7, 0x79, 0xeb, 0xb5, 0xe8, 0xf0, 0xfc, 0x67, 0x02, 0xa7,
	0x17, 0xc2, 0x8a, 0x95, 0xa2, 0xcd, 0x4d, 0x23, 0xac, 0xbf, 0x33, 0x74, 0x5f, 0xf9, 0x29, 0x9c,
	0x6a, 0xd4, 0x4b, 0x5d, 0x2e, 0x57, 0x46, 0x6b, 0x45, 0x84, 0x32, 0x3c, 0xee, 0xb0, 0x18, 0x6a,
	0xd4, 0x8b, 0xf2, 0xa2, 0x43, 0xd9, 0x04, 0x4e, 0x22, 0x93, 0x0c, 0x89, 0x3a, 0xeb, 0x05, 0x16,
	0x04, 0xd6, 0xed, 0x16, 0xc9, 0x7f, 0x25, 0xf0, 0x64, 0xd7, 0xc7, 0xb5, 0xf2, 0xc4, 0x5e, 0x43,
	0x4a, 0x3e, 0xba, 0x18, 0xf3, 0xb6, 0x09, 0xde, 0x35, 0xc1, 0x6f, 0xbb, 0xaa, 0x8a, 0x94, 0x3c,
	0x7b, 0x0c, 0x87, 0x75, 0xb9, 0x35, 0x9d, 0x86, 0xf8, 0x1e, 0xd4, 0xe5, 0x95, 0x64, 0x33, 0x18,
	0xf8, 0x28, 0xe8, 0xb3, 0x5e, 0x88, 0x61, 0xc4, 0x77, 0x7f, 0x55, 0xfc, 0xe5, 0xcc, 0xbf, 0xc0,
	0xf0, 0x93, 0x91, 0xf8, 0xd9, 0x98, 0xfa, 0x66, 0x25, 0x6a, 0x74, 0xec, 0x3d, 0x8c, 0xce, 0xe5,
	0x0f, 0x74, 0xa4, 0x3c, 0x76, 0x5f, 0xb2, 0xa7, 0xfc, 0x7f, 0x7e, 0xc7, 0xcf, 0xf6, 0x3c, 0x7e,
	0xd8, 0x5e, 0x4b, 0x7e, 0x30, 0x7f, 0x0b, 0x8f, 0xda, 0xf4, 0x17, 0xed, 0xf9, 0x31, 0x0e, 0x70,
	0x89, 0x14, 0x83, 0x65, 0xff, 0x44, 0x3c, 0x66, 0x7c, 0xaf, 0xf3, 0xfc, 0xa0, 0x3c, 0x0a, 0x92,
	0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x83, 0xc3, 0x05, 0xc4, 0x02, 0x00, 0x00,
}