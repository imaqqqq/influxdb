// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

/*
	Package storage is a generated protocol buffer package.

	It is generated from these files:
		storage.proto

	It has these top-level messages:
		ReadRequest
		ReadResponse
		CapabilitiesResponse
		HintsResponse
		TimestampRange
*/
package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ReadResponse_FrameType int32

const (
	FrameTypeSeries ReadResponse_FrameType = 0
	FrameTypePoints ReadResponse_FrameType = 1
)

var ReadResponse_FrameType_name = map[int32]string{
	0: "SERIES",
	1: "POINTS",
}
var ReadResponse_FrameType_value = map[string]int32{
	"SERIES": 0,
	"POINTS": 1,
}

func (x ReadResponse_FrameType) String() string {
	return proto.EnumName(ReadResponse_FrameType_name, int32(x))
}
func (ReadResponse_FrameType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorStorage, []int{1, 0}
}

type ReadRequest struct {
	Database       string         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	TimestampRange TimestampRange `protobuf:"bytes,2,opt,name=timestamp_range,json=timestampRange" json:"timestamp_range"`
	Descending     bool           `protobuf:"varint,3,opt,name=descending,proto3" json:"descending,omitempty"`
	// Grouping specifies a list of tags used to order the data
	Grouping []string `protobuf:"bytes,4,rep,name=grouping" json:"grouping,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

type ReadResponse struct {
	Frames []byte `protobuf:"bytes,1,opt,name=frames,proto3" json:"frames,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{1} }

type ReadResponse_SeriesFrame struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ReadResponse_SeriesFrame) Reset()         { *m = ReadResponse_SeriesFrame{} }
func (m *ReadResponse_SeriesFrame) String() string { return proto.CompactTextString(m) }
func (*ReadResponse_SeriesFrame) ProtoMessage()    {}
func (*ReadResponse_SeriesFrame) Descriptor() ([]byte, []int) {
	return fileDescriptorStorage, []int{1, 0}
}

type ReadResponse_PointsFrame struct {
	Points []byte `protobuf:"bytes,1,opt,name=points,proto3" json:"points,omitempty"`
}

func (m *ReadResponse_PointsFrame) Reset()         { *m = ReadResponse_PointsFrame{} }
func (m *ReadResponse_PointsFrame) String() string { return proto.CompactTextString(m) }
func (*ReadResponse_PointsFrame) ProtoMessage()    {}
func (*ReadResponse_PointsFrame) Descriptor() ([]byte, []int) {
	return fileDescriptorStorage, []int{1, 1}
}

type CapabilitiesResponse struct {
	Caps map[string]string `protobuf:"bytes,1,rep,name=caps" json:"caps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *CapabilitiesResponse) Reset()                    { *m = CapabilitiesResponse{} }
func (m *CapabilitiesResponse) String() string            { return proto.CompactTextString(m) }
func (*CapabilitiesResponse) ProtoMessage()               {}
func (*CapabilitiesResponse) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{2} }

type HintsResponse struct {
}

func (m *HintsResponse) Reset()                    { *m = HintsResponse{} }
func (m *HintsResponse) String() string            { return proto.CompactTextString(m) }
func (*HintsResponse) ProtoMessage()               {}
func (*HintsResponse) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{3} }

// Specifies a continuous range of nanosecond timestamps.
type TimestampRange struct {
	// Start defines the inclusive lower bound.
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End defines the inclusive upper bound.
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *TimestampRange) Reset()                    { *m = TimestampRange{} }
func (m *TimestampRange) String() string            { return proto.CompactTextString(m) }
func (*TimestampRange) ProtoMessage()               {}
func (*TimestampRange) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{4} }

func init() {
	proto.RegisterType((*ReadRequest)(nil), "storage.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "storage.ReadResponse")
	proto.RegisterType((*ReadResponse_SeriesFrame)(nil), "storage.ReadResponse.SeriesFrame")
	proto.RegisterType((*ReadResponse_PointsFrame)(nil), "storage.ReadResponse.PointsFrame")
	proto.RegisterType((*CapabilitiesResponse)(nil), "storage.CapabilitiesResponse")
	proto.RegisterType((*HintsResponse)(nil), "storage.HintsResponse")
	proto.RegisterType((*TimestampRange)(nil), "storage.TimestampRange")
	proto.RegisterEnum("storage.ReadResponse_FrameType", ReadResponse_FrameType_name, ReadResponse_FrameType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
	// Read performs a read operation using the given ReadRequest
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (Storage_ReadClient, error)
	// Capabilities returns a map of keys and values identifying the capabilities supported by the storage engine
	Capabilities(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
	Hints(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*HintsResponse, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (Storage_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[0], c.cc, "/storage.Storage/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type storageReadClient struct {
	grpc.ClientStream
}

func (x *storageReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Capabilities(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := grpc.Invoke(ctx, "/storage.Storage/Capabilities", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Hints(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*HintsResponse, error) {
	out := new(HintsResponse)
	err := grpc.Invoke(ctx, "/storage.Storage/Hints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Storage service

type StorageServer interface {
	// Read performs a read operation using the given ReadRequest
	Read(*ReadRequest, Storage_ReadServer) error
	// Capabilities returns a map of keys and values identifying the capabilities supported by the storage engine
	Capabilities(context.Context, *google_protobuf1.Empty) (*CapabilitiesResponse, error)
	Hints(context.Context, *google_protobuf1.Empty) (*HintsResponse, error)
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Read(m, &storageReadServer{stream})
}

type Storage_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type storageReadServer struct {
	grpc.ServerStream
}

func (x *storageReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Capabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Capabilities(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Hints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Hints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Hints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Hints(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capabilities",
			Handler:    _Storage_Capabilities_Handler,
		},
		{
			MethodName: "Hints",
			Handler:    _Storage_Hints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _Storage_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}

func (m *ReadRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Database) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Database)))
		i += copy(dAtA[i:], m.Database)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintStorage(dAtA, i, uint64(m.TimestampRange.Size()))
	n1, err := m.TimestampRange.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Descending {
		dAtA[i] = 0x18
		i++
		if m.Descending {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Grouping) > 0 {
		for _, s := range m.Grouping {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *ReadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Frames) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Frames)))
		i += copy(dAtA[i:], m.Frames)
	}
	return i, nil
}

func (m *ReadResponse_SeriesFrame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadResponse_SeriesFrame) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *ReadResponse_PointsFrame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadResponse_PointsFrame) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Points) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Points)))
		i += copy(dAtA[i:], m.Points)
	}
	return i, nil
}

func (m *CapabilitiesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapabilitiesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Caps) > 0 {
		for k, _ := range m.Caps {
			dAtA[i] = 0xa
			i++
			v := m.Caps[k]
			mapSize := 1 + len(k) + sovStorage(uint64(len(k))) + 1 + len(v) + sovStorage(uint64(len(v)))
			i = encodeVarintStorage(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintStorage(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintStorage(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *HintsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HintsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *TimestampRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimestampRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Start))
	}
	if m.End != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.End))
	}
	return i, nil
}

func encodeFixed64Storage(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Storage(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReadRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Database)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = m.TimestampRange.Size()
	n += 1 + l + sovStorage(uint64(l))
	if m.Descending {
		n += 2
	}
	if len(m.Grouping) > 0 {
		for _, s := range m.Grouping {
			l = len(s)
			n += 1 + l + sovStorage(uint64(l))
		}
	}
	return n
}

func (m *ReadResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Frames)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *ReadResponse_SeriesFrame) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *ReadResponse_PointsFrame) Size() (n int) {
	var l int
	_ = l
	l = len(m.Points)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *CapabilitiesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Caps) > 0 {
		for k, v := range m.Caps {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStorage(uint64(len(k))) + 1 + len(v) + sovStorage(uint64(len(v)))
			n += mapEntrySize + 1 + sovStorage(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *HintsResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *TimestampRange) Size() (n int) {
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovStorage(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovStorage(uint64(m.End))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReadRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimestampRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descending", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Descending = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grouping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grouping = append(m.Grouping, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frames", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Frames = append(m.Frames[:0], dAtA[iNdEx:postIndex]...)
			if m.Frames == nil {
				m.Frames = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadResponse_SeriesFrame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeriesFrame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeriesFrame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadResponse_PointsFrame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PointsFrame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PointsFrame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points[:0], dAtA[iNdEx:postIndex]...)
			if m.Points == nil {
				m.Points = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CapabilitiesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CapabilitiesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapabilitiesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStorage
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Caps == nil {
				m.Caps = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthStorage
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Caps[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Caps[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HintsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HintsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HintsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimestampRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimestampRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimestampRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStorage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage.proto", fileDescriptorStorage) }

var fileDescriptorStorage = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8e, 0xd2, 0x40,
	0x1c, 0xc7, 0x3b, 0x5b, 0x96, 0x5d, 0x7e, 0xb0, 0x0b, 0x19, 0x11, 0x49, 0x8d, 0x05, 0x9b, 0x18,
	0xb9, 0xd8, 0x35, 0x78, 0xd8, 0x55, 0x6f, 0x6b, 0x50, 0xf7, 0xa2, 0x9b, 0x29, 0x07, 0x6f, 0x66,
	0x80, 0xd9, 0xda, 0x48, 0x3b, 0xb5, 0x33, 0x98, 0xf0, 0x06, 0xc6, 0x78, 0xf0, 0x05, 0x3c, 0xf9,
	0x20, 0x9e, 0x34, 0x1c, 0x7d, 0x82, 0x8d, 0xe2, 0x8b, 0x98, 0x99, 0xa1, 0x15, 0x88, 0x7a, 0xfb,
	0x7d, 0x7e, 0xff, 0xbf, 0x33, 0x3f, 0x38, 0x10, 0x92, 0x67, 0x34, 0x64, 0x7e, 0x9a, 0x71, 0xc9,
	0xf1, 0xde, 0x0a, 0x9d, 0x3b, 0x61, 0x24, 0x5f, 0xcd, 0x46, 0xfe, 0x98, 0xc7, 0x47, 0x21, 0x0f,
	0xf9, 0x91, 0x8e, 0x8f, 0x66, 0x17, 0x9a, 0x34, 0x68, 0xcb, 0xd4, 0x39, 0xd7, 0x43, 0xce, 0xc3,
	0x29, 0xfb, 0x93, 0xc5, 0xe2, 0x54, 0xce, 0x4d, 0xd0, 0xfb, 0x82, 0xa0, 0x4a, 0x18, 0x9d, 0x10,
	0xf6, 0x66, 0xc6, 0x84, 0xc4, 0x0e, 0xec, 0x4f, 0xa8, 0xa4, 0x23, 0x2a, 0x58, 0x1b, 0x75, 0x51,
	0xaf, 0x42, 0x0a, 0xc6, 0x2f, 0xa0, 0x2e, 0xa3, 0x98, 0x09, 0x49, 0xe3, 0xf4, 0x65, 0x46, 0x93,
	0x90, 0xb5, 0x77, 0xba, 0xa8, 0x57, 0xed, 0x5f, 0xf3, 0xf3, 0x4d, 0x87, 0x79, 0x9c, 0xa8, 0xf0,
	0x69, 0x6b, 0x71, 0xd9, 0xb1, 0x96, 0x97, 0x9d, 0xc3, 0x4d, 0x3f, 0x39, 0x94, 0x1b, 0x8c, 0x5d,
	0x80, 0x09, 0x13, 0x63, 0x96, 0x4c, 0xa2, 0x24, 0x6c, 0xdb, 0x5d, 0xd4, 0xdb, 0x27, 0x6b, 0x1e,
	0xb5, 0x55, 0x98, 0xf1, 0x59, 0xaa, 0xa2, 0xa5, 0xae, 0xad, 0xb6, 0xca, 0xd9, 0xfb, 0x8a, 0xa0,
	0x66, 0x14, 0x88, 0x94, 0x27, 0x82, 0xe1, 0x16, 0x94, 0x2f, 0x32, 0x1a, 0x33, 0xa1, 0x05, 0xd4,
	0xc8, 0x8a, 0x9c, 0x9b, 0x50, 0x0d, 0x58, 0x16, 0x31, 0xf1, 0x58, 0x31, 0xc6, 0x50, 0x4a, 0x68,
	0x9c, 0xab, 0xd4, 0xb6, 0x73, 0x0b, 0xaa, 0xe7, 0x3c, 0x4a, 0xe4, 0x2a, 0xa5, 0x05, 0xe5, 0x54,
	0x63, 0xde, 0xc9, 0x90, 0x17, 0x40, 0x45, 0x27, 0x0c, 0xe7, 0x29, 0xc3, 0x1d, 0x28, 0x07, 0x03,
	0x72, 0x36, 0x08, 0x1a, 0x96, 0x73, 0xe5, 0xfd, 0xa7, 0x6e, 0xbd, 0x08, 0x99, 0x69, 0x2a, 0xe1,
	0xfc, 0xf9, 0xd9, 0xb3, 0x61, 0xd0, 0x40, 0x5b, 0x09, 0x66, 0x96, 0x53, 0x7a, 0xf7, 0xd9, 0xb5,
	0xbc, 0x0f, 0x08, 0x9a, 0x8f, 0x68, 0x4a, 0x47, 0xd1, 0x34, 0x92, 0x11, 0x13, 0x85, 0x9e, 0x87,
	0x50, 0x1a, 0xd3, 0x54, 0xed, 0x60, 0xf7, 0xaa, 0xfd, 0xdb, 0xc5, 0x5b, 0xff, 0x2d, 0x59, 0x39,
	0xc5, 0x20, 0x91, 0xd9, 0x9c, 0xe8, 0x22, 0xe7, 0x18, 0x2a, 0x85, 0x0b, 0x37, 0xc0, 0x7e, 0xcd,
	0xe6, 0x2b, 0xc5, 0xca, 0xc4, 0x4d, 0xd8, 0x7d, 0x4b, 0xa7, 0x33, 0xf3, 0x91, 0x15, 0x62, 0xe0,
	0xc1, 0xce, 0x09, 0xf2, 0xea, 0x70, 0xf0, 0x54, 0x6d, 0x97, 0x77, 0xf6, 0x4e, 0x60, 0xeb, 0x17,
	0x55, 0xb1, 0x90, 0x34, 0x93, 0xba, 0xa1, 0x4d, 0x0c, 0xa8, 0x21, 0x2c, 0x99, 0xe8, 0x86, 0x36,
	0x51, 0x66, 0xff, 0x1b, 0x82, 0xbd, 0xc0, 0x2c, 0x8d, 0x8f, 0xa1, 0xa4, 0x3e, 0x0b, 0x37, 0x0b,
	0x19, 0x6b, 0xd7, 0xe7, 0x5c, 0xdd, 0xf2, 0xae, 0x46, 0x5b, 0x77, 0x11, 0x7e, 0x02, 0xb5, 0x75,
	0xc1, 0xb8, 0xe5, 0x9b, 0xb3, 0xf6, 0xf3, 0xb3, 0xf6, 0x07, 0xea, 0xac, 0x9d, 0x1b, 0xff, 0x7d,
	0x1f, 0xcf, 0xc2, 0xf7, 0x61, 0x57, 0x0b, 0xfb, 0x67, 0x87, 0x56, 0xd1, 0x61, 0xf3, 0x01, 0xac,
	0xd3, 0xe6, 0xe2, 0xa7, 0x6b, 0x2d, 0x96, 0x2e, 0xfa, 0xbe, 0x74, 0xd1, 0x8f, 0xa5, 0x8b, 0x3e,
	0xfe, 0x72, 0xad, 0x51, 0x59, 0xd7, 0xdf, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x83, 0x09,
	0x29, 0xaf, 0x03, 0x00, 0x00,
}
