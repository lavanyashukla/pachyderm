// Code generated by protoc-gen-go.
// source: server/pkg/metrics/metrics.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	server/pkg/metrics/metrics.proto

It has these top-level messages:
	Metrics
*/
package metrics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Metrics struct {
	ID        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	MachineID string `protobuf:"bytes,2,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	Nodes     int64  `protobuf:"varint,3,opt,name=nodes" json:"nodes,omitempty"`
	Repos     int64  `protobuf:"varint,4,opt,name=repos" json:"repos,omitempty"`
	Commits   int64  `protobuf:"varint,5,opt,name=commits" json:"commits,omitempty"`
	Files     int64  `protobuf:"varint,6,opt,name=files" json:"files,omitempty"`
	Bytes     int64  `protobuf:"varint,7,opt,name=bytes" json:"bytes,omitempty"`
	Jobs      int64  `protobuf:"varint,8,opt,name=jobs" json:"jobs,omitempty"`
	Pipelines int64  `protobuf:"varint,9,opt,name=pipelines" json:"pipelines,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Metrics)(nil), "Metrics")
}

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0xcf, 0x3d, 0x8e, 0xc2, 0x30,
	0x10, 0x05, 0x60, 0xe5, 0x7f, 0x3d, 0xc5, 0x16, 0xd6, 0x16, 0x2e, 0x16, 0x29, 0xa2, 0xa2, 0x22,
	0x05, 0xa7, 0xa0, 0xa0, 0xc9, 0x05, 0x10, 0x49, 0x0c, 0x18, 0xe2, 0xd8, 0xf2, 0x58, 0x48, 0xdc,
	0x95, 0xc3, 0xe0, 0x8c, 0x89, 0xa8, 0x3c, 0xef, 0x9b, 0x27, 0xcb, 0x86, 0x1a, 0xa5, 0x7b, 0x48,
	0xd7, 0xd8, 0xfb, 0xa5, 0xd1, 0xd2, 0x3b, 0xd5, 0xe3, 0x72, 0x6e, 0xad, 0x33, 0xde, 0xac, 0x5f,
	0x09, 0x54, 0x87, 0x28, 0xfc, 0x17, 0x52, 0x35, 0x88, 0xa4, 0x4e, 0x36, 0xac, 0x0d, 0x13, 0x5f,
	0x01, 0xe8, 0x53, 0x7f, 0x55, 0x93, 0x3c, 0x06, 0x4f, 0xc9, 0xd9, 0x47, 0xf6, 0x03, 0xff, 0x83,
	0x62, 0x32, 0x83, 0x44, 0x91, 0x85, 0x4d, 0xd6, 0xc6, 0x30, 0xab, 0x93, 0xd6, 0xa0, 0xc8, 0xa3,
	0x52, 0xe0, 0x02, 0xaa, 0xde, 0x68, 0xad, 0x3c, 0x8a, 0x82, 0x7c, 0x89, 0x73, 0xff, 0xac, 0xc6,
	0x70, 0x4b, 0x19, 0xfb, 0x14, 0x66, 0xed, 0x9e, 0x3e, 0x68, 0x15, 0x95, 0x02, 0xe7, 0x90, 0xdf,
	0x4c, 0x87, 0xe2, 0x87, 0x90, 0x66, 0xfe, 0x0f, 0xcc, 0x2a, 0x2b, 0xc7, 0xf0, 0x26, 0x14, 0x8c,
	0x16, 0x5f, 0xe8, 0x4a, 0xfa, 0xe5, 0xee, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xd0, 0xfd, 0x50,
	0x09, 0x01, 0x00, 0x00,
}
