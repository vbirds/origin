// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/dynamicdiscover.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NodeInfo struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	ListenAddr           string   `protobuf:"bytes,3,opt,name=ListenAddr,proto3" json:"ListenAddr,omitempty"`
	Private              bool     `protobuf:"varint,4,opt,name=Private,proto3" json:"Private,omitempty"`
	PublicServiceList    []string `protobuf:"bytes,5,rep,name=PublicServiceList,proto3" json:"PublicServiceList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bfdd3ec0419520f, []int{0}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeInfo) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func (m *NodeInfo) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *NodeInfo) GetPublicServiceList() []string {
	if m != nil {
		return m.PublicServiceList
	}
	return nil
}

//Client->Master
type ServiceDiscoverReq struct {
	NodeInfo             *NodeInfo `protobuf:"bytes,1,opt,name=nodeInfo,proto3" json:"nodeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServiceDiscoverReq) Reset()         { *m = ServiceDiscoverReq{} }
func (m *ServiceDiscoverReq) String() string { return proto.CompactTextString(m) }
func (*ServiceDiscoverReq) ProtoMessage()    {}
func (*ServiceDiscoverReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bfdd3ec0419520f, []int{1}
}
func (m *ServiceDiscoverReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceDiscoverReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceDiscoverReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceDiscoverReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceDiscoverReq.Merge(m, src)
}
func (m *ServiceDiscoverReq) XXX_Size() int {
	return m.Size()
}
func (m *ServiceDiscoverReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceDiscoverReq.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceDiscoverReq proto.InternalMessageInfo

func (m *ServiceDiscoverReq) GetNodeInfo() *NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

//Master->Client
type SubscribeDiscoverNotify struct {
	DelNodeId            int32     `protobuf:"varint,1,opt,name=DelNodeId,proto3" json:"DelNodeId,omitempty"`
	NodeInfo             *NodeInfo `protobuf:"bytes,2,opt,name=nodeInfo,proto3" json:"nodeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SubscribeDiscoverNotify) Reset()         { *m = SubscribeDiscoverNotify{} }
func (m *SubscribeDiscoverNotify) String() string { return proto.CompactTextString(m) }
func (*SubscribeDiscoverNotify) ProtoMessage()    {}
func (*SubscribeDiscoverNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bfdd3ec0419520f, []int{2}
}
func (m *SubscribeDiscoverNotify) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeDiscoverNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeDiscoverNotify.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeDiscoverNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeDiscoverNotify.Merge(m, src)
}
func (m *SubscribeDiscoverNotify) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeDiscoverNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeDiscoverNotify.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeDiscoverNotify proto.InternalMessageInfo

func (m *SubscribeDiscoverNotify) GetDelNodeId() int32 {
	if m != nil {
		return m.DelNodeId
	}
	return 0
}

func (m *SubscribeDiscoverNotify) GetNodeInfo() *NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

//Master->Client
type ServiceDiscoverRes struct {
	NodeInfo             []*NodeInfo `protobuf:"bytes,1,rep,name=nodeInfo,proto3" json:"nodeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ServiceDiscoverRes) Reset()         { *m = ServiceDiscoverRes{} }
func (m *ServiceDiscoverRes) String() string { return proto.CompactTextString(m) }
func (*ServiceDiscoverRes) ProtoMessage()    {}
func (*ServiceDiscoverRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bfdd3ec0419520f, []int{3}
}
func (m *ServiceDiscoverRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceDiscoverRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceDiscoverRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceDiscoverRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceDiscoverRes.Merge(m, src)
}
func (m *ServiceDiscoverRes) XXX_Size() int {
	return m.Size()
}
func (m *ServiceDiscoverRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceDiscoverRes.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceDiscoverRes proto.InternalMessageInfo

func (m *ServiceDiscoverRes) GetNodeInfo() []*NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "rpc.NodeInfo")
	proto.RegisterType((*ServiceDiscoverReq)(nil), "rpc.ServiceDiscoverReq")
	proto.RegisterType((*SubscribeDiscoverNotify)(nil), "rpc.SubscribeDiscoverNotify")
	proto.RegisterType((*ServiceDiscoverRes)(nil), "rpc.ServiceDiscoverRes")
}

func init() { proto.RegisterFile("rpc/dynamicdiscover.proto", fileDescriptor_9bfdd3ec0419520f) }

var fileDescriptor_9bfdd3ec0419520f = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2a, 0x48, 0xd6,
	0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x4e, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x5a, 0xc6, 0xc8, 0xc5, 0xe1,
	0x97, 0x9f, 0x92, 0xea, 0x99, 0x97, 0x96, 0x2f, 0x24, 0xc6, 0xc5, 0x06, 0x66, 0xa7, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x42, 0x52, 0x10, 0x35, 0x7e, 0x89, 0xb9, 0xa9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x1c, 0x17, 0x97, 0x4f, 0x66, 0x71, 0x49,
	0x6a, 0x9e, 0x63, 0x4a, 0x4a, 0x91, 0x04, 0x33, 0x58, 0x16, 0x49, 0x44, 0x48, 0x82, 0x8b, 0x3d,
	0xa0, 0x28, 0xb3, 0x2c, 0xb1, 0x24, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc6, 0x15,
	0xd2, 0xe1, 0x12, 0x0c, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0x0e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x05, 0x69, 0x92, 0x60, 0x55, 0x60, 0xd6, 0xe0, 0x0c, 0xc2, 0x94, 0x50, 0xb2, 0xe7, 0x12, 0x82,
	0x72, 0x5d, 0xa0, 0xde, 0x08, 0x4a, 0x2d, 0x14, 0xd2, 0xe4, 0xe2, 0xc8, 0x83, 0xba, 0x1e, 0xec,
	0x66, 0x6e, 0x23, 0x5e, 0xbd, 0xa2, 0x82, 0x64, 0x3d, 0x98, 0x97, 0x82, 0xe0, 0xd2, 0x4a, 0x49,
	0x5c, 0xe2, 0xc1, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x49, 0x70, 0x23, 0xfc, 0xf2, 0x4b, 0x32,
	0xd3, 0x2a, 0x85, 0x64, 0xb8, 0x38, 0x5d, 0x52, 0x73, 0x50, 0xbc, 0x8e, 0x10, 0x40, 0xb1, 0x83,
	0x09, 0xbf, 0x1d, 0xd8, 0x1c, 0x59, 0x8c, 0xe6, 0x48, 0x66, 0x3c, 0x06, 0x38, 0x09, 0x9f, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x51, 0xac, 0x7a, 0xfa, 0x45,
	0x05, 0xc9, 0x49, 0x6c, 0xe0, 0xf8, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x96, 0xd9,
	0x54, 0xcc, 0x01, 0x00, 0x00,
}

func (m *NodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PublicServiceList) > 0 {
		for iNdEx := len(m.PublicServiceList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PublicServiceList[iNdEx])
			copy(dAtA[i:], m.PublicServiceList[iNdEx])
			i = encodeVarintDynamicdiscover(dAtA, i, uint64(len(m.PublicServiceList[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Private {
		i--
		if m.Private {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.ListenAddr) > 0 {
		i -= len(m.ListenAddr)
		copy(dAtA[i:], m.ListenAddr)
		i = encodeVarintDynamicdiscover(dAtA, i, uint64(len(m.ListenAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeName) > 0 {
		i -= len(m.NodeName)
		copy(dAtA[i:], m.NodeName)
		i = encodeVarintDynamicdiscover(dAtA, i, uint64(len(m.NodeName)))
		i--
		dAtA[i] = 0x12
	}
	if m.NodeId != 0 {
		i = encodeVarintDynamicdiscover(dAtA, i, uint64(m.NodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ServiceDiscoverReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceDiscoverReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceDiscoverReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NodeInfo != nil {
		{
			size, err := m.NodeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDynamicdiscover(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubscribeDiscoverNotify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeDiscoverNotify) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeDiscoverNotify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NodeInfo != nil {
		{
			size, err := m.NodeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDynamicdiscover(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.DelNodeId != 0 {
		i = encodeVarintDynamicdiscover(dAtA, i, uint64(m.DelNodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ServiceDiscoverRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceDiscoverRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceDiscoverRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NodeInfo) > 0 {
		for iNdEx := len(m.NodeInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDynamicdiscover(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDynamicdiscover(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicdiscover(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeId != 0 {
		n += 1 + sovDynamicdiscover(uint64(m.NodeId))
	}
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovDynamicdiscover(uint64(l))
	}
	l = len(m.ListenAddr)
	if l > 0 {
		n += 1 + l + sovDynamicdiscover(uint64(l))
	}
	if m.Private {
		n += 2
	}
	if len(m.PublicServiceList) > 0 {
		for _, s := range m.PublicServiceList {
			l = len(s)
			n += 1 + l + sovDynamicdiscover(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServiceDiscoverReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovDynamicdiscover(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubscribeDiscoverNotify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DelNodeId != 0 {
		n += 1 + sovDynamicdiscover(uint64(m.DelNodeId))
	}
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovDynamicdiscover(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServiceDiscoverRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NodeInfo) > 0 {
		for _, e := range m.NodeInfo {
			l = e.Size()
			n += 1 + l + sovDynamicdiscover(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDynamicdiscover(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicdiscover(x uint64) (n int) {
	return sovDynamicdiscover(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicdiscover
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			m.NodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Private", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Private = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicServiceList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicServiceList = append(m.PublicServiceList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicdiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceDiscoverReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicdiscover
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceDiscoverReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceDiscoverReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicdiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubscribeDiscoverNotify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicdiscover
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubscribeDiscoverNotify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeDiscoverNotify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelNodeId", wireType)
			}
			m.DelNodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelNodeId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicdiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceDiscoverRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicdiscover
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceDiscoverRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceDiscoverRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeInfo = append(m.NodeInfo, &NodeInfo{})
			if err := m.NodeInfo[len(m.NodeInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicdiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicdiscover
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDynamicdiscover(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicdiscover
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicdiscover
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDynamicdiscover
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicdiscover
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicdiscover
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicdiscover        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicdiscover          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicdiscover = fmt.Errorf("proto: unexpected end of group")
)
