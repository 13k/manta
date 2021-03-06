// Code generated by protoc-gen-go.
// source: networksystem_protomessages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetMessageSplitscreenUserChanged struct {
	Slot             *uint32 `protobuf:"varint,1,opt,name=slot" json:"slot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetMessageSplitscreenUserChanged) Reset()         { *m = NetMessageSplitscreenUserChanged{} }
func (m *NetMessageSplitscreenUserChanged) String() string { return proto.CompactTextString(m) }
func (*NetMessageSplitscreenUserChanged) ProtoMessage()    {}
func (*NetMessageSplitscreenUserChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor33, []int{0}
}

func (m *NetMessageSplitscreenUserChanged) GetSlot() uint32 {
	if m != nil && m.Slot != nil {
		return *m.Slot
	}
	return 0
}

type NetMessageConnectionClosed struct {
	Reason           *uint32 `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetMessageConnectionClosed) Reset()                    { *m = NetMessageConnectionClosed{} }
func (m *NetMessageConnectionClosed) String() string            { return proto.CompactTextString(m) }
func (*NetMessageConnectionClosed) ProtoMessage()               {}
func (*NetMessageConnectionClosed) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{1} }

func (m *NetMessageConnectionClosed) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type NetMessageConnectionCrashed struct {
	Reason           *uint32 `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetMessageConnectionCrashed) Reset()                    { *m = NetMessageConnectionCrashed{} }
func (m *NetMessageConnectionCrashed) String() string            { return proto.CompactTextString(m) }
func (*NetMessageConnectionCrashed) ProtoMessage()               {}
func (*NetMessageConnectionCrashed) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{2} }

func (m *NetMessageConnectionCrashed) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type NetMessagePacketStart struct {
	IncomingSequence     *uint32 `protobuf:"varint,1,opt,name=incoming_sequence" json:"incoming_sequence,omitempty"`
	OutgoingAcknowledged *uint32 `protobuf:"varint,2,opt,name=outgoing_acknowledged" json:"outgoing_acknowledged,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *NetMessagePacketStart) Reset()                    { *m = NetMessagePacketStart{} }
func (m *NetMessagePacketStart) String() string            { return proto.CompactTextString(m) }
func (*NetMessagePacketStart) ProtoMessage()               {}
func (*NetMessagePacketStart) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{3} }

func (m *NetMessagePacketStart) GetIncomingSequence() uint32 {
	if m != nil && m.IncomingSequence != nil {
		return *m.IncomingSequence
	}
	return 0
}

func (m *NetMessagePacketStart) GetOutgoingAcknowledged() uint32 {
	if m != nil && m.OutgoingAcknowledged != nil {
		return *m.OutgoingAcknowledged
	}
	return 0
}

type NetMessagePacketEnd struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NetMessagePacketEnd) Reset()                    { *m = NetMessagePacketEnd{} }
func (m *NetMessagePacketEnd) String() string            { return proto.CompactTextString(m) }
func (*NetMessagePacketEnd) ProtoMessage()               {}
func (*NetMessagePacketEnd) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{4} }

func init() {
	proto.RegisterType((*NetMessageSplitscreenUserChanged)(nil), "dota.NetMessageSplitscreenUserChanged")
	proto.RegisterType((*NetMessageConnectionClosed)(nil), "dota.NetMessageConnectionClosed")
	proto.RegisterType((*NetMessageConnectionCrashed)(nil), "dota.NetMessageConnectionCrashed")
	proto.RegisterType((*NetMessagePacketStart)(nil), "dota.NetMessagePacketStart")
	proto.RegisterType((*NetMessagePacketEnd)(nil), "dota.NetMessagePacketEnd")
}

func init() { proto.RegisterFile("networksystem_protomessages.proto", fileDescriptor33) }

var fileDescriptor33 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0xcc, 0xc1, 0x4a, 0x03, 0x41,
	0x10, 0x04, 0x50, 0x57, 0x83, 0x87, 0x46, 0x05, 0x47, 0x16, 0xa2, 0x22, 0xc4, 0x9c, 0x3c, 0xa8,
	0xf8, 0x0d, 0x2e, 0x1e, 0x15, 0x25, 0x78, 0x0e, 0xc3, 0x4c, 0x31, 0x59, 0x76, 0xb6, 0x3b, 0x4e,
	0x77, 0x08, 0xde, 0xfc, 0x74, 0x31, 0x11, 0x02, 0x61, 0x8f, 0x45, 0xbd, 0x2a, 0xba, 0x65, 0xd8,
	0x5a, 0x4a, 0xa7, 0xdf, 0x6a, 0xe8, 0xe7, 0xcb, 0x22, 0x26, 0x3d, 0x54, 0x7d, 0x82, 0x3e, 0x6e,
	0x92, 0x1b, 0x45, 0x31, 0x3f, 0x7d, 0xa2, 0xc9, 0x1b, 0xec, 0x75, 0x5b, 0xcd, 0x96, 0xb9, 0x35,
	0x0d, 0x05, 0xe0, 0x4f, 0x45, 0x69, 0x16, 0x9e, 0x13, 0xa2, 0x3b, 0xa1, 0x91, 0x66, 0xb1, 0x71,
	0x35, 0xa9, 0xee, 0x4e, 0xa7, 0xf7, 0x74, 0xb5, 0x5b, 0x34, 0xc2, 0x8c, 0x60, 0xad, 0x70, 0x93,
	0x45, 0x11, 0xdd, 0x19, 0x1d, 0x17, 0x78, 0x15, 0xfe, 0xd7, 0x0f, 0x74, 0x3d, 0xa8, 0x8b, 0xd7,
	0xc5, 0x00, 0xff, 0xa0, 0x7a, 0xc7, 0xdf, 0x7d, 0xe8, 0x60, 0x33, 0xf3, 0xc5, 0xdc, 0x25, 0x9d,
	0xb7, 0x1c, 0xa4, 0x6f, 0x39, 0xcd, 0x15, 0x5f, 0x2b, 0x70, 0xc0, 0x76, 0xe3, 0x6e, 0xa8, 0x96,
	0x95, 0x25, 0xf9, 0xab, 0x7c, 0xe8, 0x58, 0xd6, 0x19, 0x31, 0x21, 0x8e, 0x0f, 0x37, 0x97, 0x35,
	0x5d, 0xec, 0x5f, 0xbe, 0x70, 0x7c, 0x3e, 0xfa, 0xa9, 0x0e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x71, 0x7e, 0x7a, 0x0f, 0x27, 0x01, 0x00, 0x00,
}
