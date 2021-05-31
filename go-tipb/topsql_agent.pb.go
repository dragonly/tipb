// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topsql_agent.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorTopsqlAgent, []int{0} }

type CPUTimeRequestTiDB struct {
	TimestampList    []uint64 `protobuf:"varint,1,rep,name=timestamp_list,json=timestampList" json:"timestamp_list,omitempty"`
	CpuTimeMsList    []uint32 `protobuf:"varint,2,rep,name=cpu_time_ms_list,json=cpuTimeMsList" json:"cpu_time_ms_list,omitempty"`
	SqlDigest        *string  `protobuf:"bytes,3,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	NormalizedSql    *string  `protobuf:"bytes,4,opt,name=normalized_sql,json=normalizedSql" json:"normalized_sql,omitempty"`
	PlanDigest       *string  `protobuf:"bytes,5,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	NormalizedPlan   *string  `protobuf:"bytes,6,opt,name=normalized_plan,json=normalizedPlan" json:"normalized_plan,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CPUTimeRequestTiDB) Reset()                    { *m = CPUTimeRequestTiDB{} }
func (m *CPUTimeRequestTiDB) String() string            { return proto.CompactTextString(m) }
func (*CPUTimeRequestTiDB) ProtoMessage()               {}
func (*CPUTimeRequestTiDB) Descriptor() ([]byte, []int) { return fileDescriptorTopsqlAgent, []int{1} }

func (m *CPUTimeRequestTiDB) GetTimestampList() []uint64 {
	if m != nil {
		return m.TimestampList
	}
	return nil
}

func (m *CPUTimeRequestTiDB) GetCpuTimeMsList() []uint32 {
	if m != nil {
		return m.CpuTimeMsList
	}
	return nil
}

func (m *CPUTimeRequestTiDB) GetSqlDigest() string {
	if m != nil && m.SqlDigest != nil {
		return *m.SqlDigest
	}
	return ""
}

func (m *CPUTimeRequestTiDB) GetNormalizedSql() string {
	if m != nil && m.NormalizedSql != nil {
		return *m.NormalizedSql
	}
	return ""
}

func (m *CPUTimeRequestTiDB) GetPlanDigest() string {
	if m != nil && m.PlanDigest != nil {
		return *m.PlanDigest
	}
	return ""
}

func (m *CPUTimeRequestTiDB) GetNormalizedPlan() string {
	if m != nil && m.NormalizedPlan != nil {
		return *m.NormalizedPlan
	}
	return ""
}

type CPUTimeRequestTiKV struct {
	TimestampList    []uint64 `protobuf:"varint,1,rep,name=timestamp_list,json=timestampList" json:"timestamp_list,omitempty"`
	CpuTimeMsList    []uint32 `protobuf:"varint,2,rep,name=cpu_time_ms_list,json=cpuTimeMsList" json:"cpu_time_ms_list,omitempty"`
	SqlDigest        *string  `protobuf:"bytes,3,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	PlanDigest       *string  `protobuf:"bytes,4,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CPUTimeRequestTiKV) Reset()                    { *m = CPUTimeRequestTiKV{} }
func (m *CPUTimeRequestTiKV) String() string            { return proto.CompactTextString(m) }
func (*CPUTimeRequestTiKV) ProtoMessage()               {}
func (*CPUTimeRequestTiKV) Descriptor() ([]byte, []int) { return fileDescriptorTopsqlAgent, []int{2} }

func (m *CPUTimeRequestTiKV) GetTimestampList() []uint64 {
	if m != nil {
		return m.TimestampList
	}
	return nil
}

func (m *CPUTimeRequestTiKV) GetCpuTimeMsList() []uint32 {
	if m != nil {
		return m.CpuTimeMsList
	}
	return nil
}

func (m *CPUTimeRequestTiKV) GetSqlDigest() string {
	if m != nil && m.SqlDigest != nil {
		return *m.SqlDigest
	}
	return ""
}

func (m *CPUTimeRequestTiKV) GetPlanDigest() string {
	if m != nil && m.PlanDigest != nil {
		return *m.PlanDigest
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "tipb.Empty")
	proto.RegisterType((*CPUTimeRequestTiDB)(nil), "tipb.CPUTimeRequestTiDB")
	proto.RegisterType((*CPUTimeRequestTiKV)(nil), "tipb.CPUTimeRequestTiKV")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TopSQLAgent service

type TopSQLAgentClient interface {
	CollectTiDB(ctx context.Context, opts ...grpc.CallOption) (TopSQLAgent_CollectTiDBClient, error)
	CollectTiKV(ctx context.Context, opts ...grpc.CallOption) (TopSQLAgent_CollectTiKVClient, error)
}

type topSQLAgentClient struct {
	cc *grpc.ClientConn
}

func NewTopSQLAgentClient(cc *grpc.ClientConn) TopSQLAgentClient {
	return &topSQLAgentClient{cc}
}

func (c *topSQLAgentClient) CollectTiDB(ctx context.Context, opts ...grpc.CallOption) (TopSQLAgent_CollectTiDBClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TopSQLAgent_serviceDesc.Streams[0], c.cc, "/tipb.TopSQLAgent/CollectTiDB", opts...)
	if err != nil {
		return nil, err
	}
	x := &topSQLAgentCollectTiDBClient{stream}
	return x, nil
}

type TopSQLAgent_CollectTiDBClient interface {
	Send(*CPUTimeRequestTiDB) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type topSQLAgentCollectTiDBClient struct {
	grpc.ClientStream
}

func (x *topSQLAgentCollectTiDBClient) Send(m *CPUTimeRequestTiDB) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topSQLAgentCollectTiDBClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topSQLAgentClient) CollectTiKV(ctx context.Context, opts ...grpc.CallOption) (TopSQLAgent_CollectTiKVClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TopSQLAgent_serviceDesc.Streams[1], c.cc, "/tipb.TopSQLAgent/CollectTiKV", opts...)
	if err != nil {
		return nil, err
	}
	x := &topSQLAgentCollectTiKVClient{stream}
	return x, nil
}

type TopSQLAgent_CollectTiKVClient interface {
	Send(*CPUTimeRequestTiKV) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type topSQLAgentCollectTiKVClient struct {
	grpc.ClientStream
}

func (x *topSQLAgentCollectTiKVClient) Send(m *CPUTimeRequestTiKV) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topSQLAgentCollectTiKVClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TopSQLAgent service

type TopSQLAgentServer interface {
	CollectTiDB(TopSQLAgent_CollectTiDBServer) error
	CollectTiKV(TopSQLAgent_CollectTiKVServer) error
}

func RegisterTopSQLAgentServer(s *grpc.Server, srv TopSQLAgentServer) {
	s.RegisterService(&_TopSQLAgent_serviceDesc, srv)
}

func _TopSQLAgent_CollectTiDB_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopSQLAgentServer).CollectTiDB(&topSQLAgentCollectTiDBServer{stream})
}

type TopSQLAgent_CollectTiDBServer interface {
	Send(*Empty) error
	Recv() (*CPUTimeRequestTiDB, error)
	grpc.ServerStream
}

type topSQLAgentCollectTiDBServer struct {
	grpc.ServerStream
}

func (x *topSQLAgentCollectTiDBServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topSQLAgentCollectTiDBServer) Recv() (*CPUTimeRequestTiDB, error) {
	m := new(CPUTimeRequestTiDB)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TopSQLAgent_CollectTiKV_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopSQLAgentServer).CollectTiKV(&topSQLAgentCollectTiKVServer{stream})
}

type TopSQLAgent_CollectTiKVServer interface {
	Send(*Empty) error
	Recv() (*CPUTimeRequestTiKV, error)
	grpc.ServerStream
}

type topSQLAgentCollectTiKVServer struct {
	grpc.ServerStream
}

func (x *topSQLAgentCollectTiKVServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topSQLAgentCollectTiKVServer) Recv() (*CPUTimeRequestTiKV, error) {
	m := new(CPUTimeRequestTiKV)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TopSQLAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tipb.TopSQLAgent",
	HandlerType: (*TopSQLAgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CollectTiDB",
			Handler:       _TopSQLAgent_CollectTiDB_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CollectTiKV",
			Handler:       _TopSQLAgent_CollectTiKV_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "topsql_agent.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CPUTimeRequestTiDB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPUTimeRequestTiDB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TimestampList) > 0 {
		for _, num := range m.TimestampList {
			dAtA[i] = 0x8
			i++
			i = encodeVarintTopsqlAgent(dAtA, i, uint64(num))
		}
	}
	if len(m.CpuTimeMsList) > 0 {
		for _, num := range m.CpuTimeMsList {
			dAtA[i] = 0x10
			i++
			i = encodeVarintTopsqlAgent(dAtA, i, uint64(num))
		}
	}
	if m.SqlDigest != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.SqlDigest)))
		i += copy(dAtA[i:], *m.SqlDigest)
	}
	if m.NormalizedSql != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.NormalizedSql)))
		i += copy(dAtA[i:], *m.NormalizedSql)
	}
	if m.PlanDigest != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.PlanDigest)))
		i += copy(dAtA[i:], *m.PlanDigest)
	}
	if m.NormalizedPlan != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.NormalizedPlan)))
		i += copy(dAtA[i:], *m.NormalizedPlan)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CPUTimeRequestTiKV) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPUTimeRequestTiKV) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TimestampList) > 0 {
		for _, num := range m.TimestampList {
			dAtA[i] = 0x8
			i++
			i = encodeVarintTopsqlAgent(dAtA, i, uint64(num))
		}
	}
	if len(m.CpuTimeMsList) > 0 {
		for _, num := range m.CpuTimeMsList {
			dAtA[i] = 0x10
			i++
			i = encodeVarintTopsqlAgent(dAtA, i, uint64(num))
		}
	}
	if m.SqlDigest != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.SqlDigest)))
		i += copy(dAtA[i:], *m.SqlDigest)
	}
	if m.PlanDigest != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTopsqlAgent(dAtA, i, uint64(len(*m.PlanDigest)))
		i += copy(dAtA[i:], *m.PlanDigest)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTopsqlAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CPUTimeRequestTiDB) Size() (n int) {
	var l int
	_ = l
	if len(m.TimestampList) > 0 {
		for _, e := range m.TimestampList {
			n += 1 + sovTopsqlAgent(uint64(e))
		}
	}
	if len(m.CpuTimeMsList) > 0 {
		for _, e := range m.CpuTimeMsList {
			n += 1 + sovTopsqlAgent(uint64(e))
		}
	}
	if m.SqlDigest != nil {
		l = len(*m.SqlDigest)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.NormalizedSql != nil {
		l = len(*m.NormalizedSql)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(*m.PlanDigest)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.NormalizedPlan != nil {
		l = len(*m.NormalizedPlan)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CPUTimeRequestTiKV) Size() (n int) {
	var l int
	_ = l
	if len(m.TimestampList) > 0 {
		for _, e := range m.TimestampList {
			n += 1 + sovTopsqlAgent(uint64(e))
		}
	}
	if len(m.CpuTimeMsList) > 0 {
		for _, e := range m.CpuTimeMsList {
			n += 1 + sovTopsqlAgent(uint64(e))
		}
	}
	if m.SqlDigest != nil {
		l = len(*m.SqlDigest)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(*m.PlanDigest)
		n += 1 + l + sovTopsqlAgent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTopsqlAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTopsqlAgent(x uint64) (n int) {
	return sovTopsqlAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopsqlAgent
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTopsqlAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopsqlAgent
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
func (m *CPUTimeRequestTiDB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopsqlAgent
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
			return fmt.Errorf("proto: CPUTimeRequestTiDB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPUTimeRequestTiDB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TimestampList = append(m.TimestampList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTopsqlAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopsqlAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TimestampList = append(m.TimestampList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampList", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CpuTimeMsList = append(m.CpuTimeMsList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTopsqlAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopsqlAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CpuTimeMsList = append(m.CpuTimeMsList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuTimeMsList", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.SqlDigest = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NormalizedSql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.NormalizedSql = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.PlanDigest = &s
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NormalizedPlan", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.NormalizedPlan = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopsqlAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopsqlAgent
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
func (m *CPUTimeRequestTiKV) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopsqlAgent
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
			return fmt.Errorf("proto: CPUTimeRequestTiKV: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPUTimeRequestTiKV: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TimestampList = append(m.TimestampList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTopsqlAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopsqlAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TimestampList = append(m.TimestampList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampList", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CpuTimeMsList = append(m.CpuTimeMsList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopsqlAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTopsqlAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopsqlAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CpuTimeMsList = append(m.CpuTimeMsList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuTimeMsList", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.SqlDigest = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopsqlAgent
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
				return ErrInvalidLengthTopsqlAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.PlanDigest = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopsqlAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTopsqlAgent
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
func skipTopsqlAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopsqlAgent
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
					return 0, ErrIntOverflowTopsqlAgent
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
					return 0, ErrIntOverflowTopsqlAgent
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
				return 0, ErrInvalidLengthTopsqlAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTopsqlAgent
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
				next, err := skipTopsqlAgent(dAtA[start:])
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
	ErrInvalidLengthTopsqlAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopsqlAgent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("topsql_agent.proto", fileDescriptorTopsqlAgent) }

var fileDescriptorTopsqlAgent = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0x4d, 0x4a, 0xc3, 0x40,
	0x1c, 0xc5, 0x3b, 0x36, 0x55, 0x3a, 0x21, 0x2a, 0x83, 0x42, 0x28, 0x18, 0x4b, 0x41, 0x1a, 0x37,
	0x51, 0x5c, 0xba, 0xb3, 0xad, 0xab, 0x56, 0xa8, 0x69, 0xed, 0x36, 0xa4, 0xe9, 0x10, 0x06, 0x66,
	0x32, 0x93, 0xce, 0x74, 0xa1, 0x7b, 0xef, 0xe0, 0x01, 0x3c, 0x8c, 0x4b, 0x8f, 0x20, 0xf1, 0x0a,
	0x1e, 0x40, 0x66, 0xe2, 0x47, 0xad, 0x75, 0xeb, 0x2e, 0xfc, 0xfe, 0xef, 0xbd, 0xcc, 0xe3, 0x41,
	0xa4, 0xb8, 0x90, 0x39, 0x8d, 0xe2, 0x14, 0x67, 0x2a, 0x10, 0x73, 0xae, 0x38, 0xb2, 0x14, 0x11,
	0xd3, 0xc6, 0x5e, 0xca, 0x53, 0x6e, 0xc0, 0x89, 0xfe, 0x2a, 0x6f, 0xad, 0x2d, 0x58, 0xbb, 0x64,
	0x42, 0xdd, 0xb6, 0xde, 0x00, 0x44, 0xdd, 0xe1, 0xcd, 0x98, 0x30, 0x1c, 0xe2, 0x7c, 0x81, 0xa5,
	0x1a, 0x93, 0x5e, 0x07, 0x1d, 0xc1, 0x6d, 0x45, 0x18, 0x96, 0x2a, 0x66, 0x22, 0xa2, 0x44, 0x2a,
	0x17, 0x34, 0xab, 0xbe, 0x15, 0x3a, 0x5f, 0x74, 0x40, 0xa4, 0x42, 0x6d, 0xb8, 0x9b, 0x88, 0x45,
	0xa4, 0x61, 0xc4, 0x64, 0x29, 0xdc, 0x68, 0x56, 0x7d, 0x27, 0x74, 0x12, 0xb1, 0xd0, 0xa1, 0x57,
	0xd2, 0x08, 0x0f, 0x20, 0xd4, 0xcf, 0x9b, 0x91, 0x14, 0x4b, 0xe5, 0x56, 0x9b, 0xc0, 0xaf, 0x87,
	0x75, 0x99, 0xd3, 0x9e, 0x01, 0xfa, 0x77, 0x19, 0x9f, 0xb3, 0x98, 0x92, 0x3b, 0x3c, 0x8b, 0x64,
	0x4e, 0x5d, 0xcb, 0x48, 0x9c, 0x6f, 0x3a, 0xca, 0x29, 0x3a, 0x84, 0xb6, 0xa0, 0x71, 0xf6, 0x19,
	0x53, 0x33, 0x1a, 0xa8, 0xd1, 0x47, 0x4e, 0x1b, 0xee, 0x2c, 0xe5, 0xe8, 0x83, 0xbb, 0x69, 0x44,
	0x4b, 0xf1, 0x43, 0x1a, 0x67, 0xad, 0xc7, 0x35, 0xb5, 0xfb, 0x93, 0xff, 0xae, 0xbd, 0xd2, 0xc7,
	0x5a, 0xed, 0x73, 0x76, 0x0f, 0xa0, 0x3d, 0xe6, 0x62, 0x74, 0x3d, 0xb8, 0xd0, 0xc3, 0xa2, 0x73,
	0x68, 0x77, 0x39, 0xa5, 0x38, 0x29, 0x57, 0x72, 0x03, 0x3d, 0x71, 0xf0, 0x7b, 0xbf, 0x86, 0x5d,
	0x5e, 0xca, 0x8d, 0x2b, 0x3e, 0x38, 0x05, 0x3f, 0xbc, 0xfd, 0xc9, 0x5f, 0xde, 0xfe, 0x64, 0x8d,
	0xb7, 0x73, 0xfc, 0x54, 0x78, 0xe0, 0xb9, 0xf0, 0xc0, 0x4b, 0xe1, 0x81, 0x87, 0x57, 0xaf, 0x02,
	0xf7, 0x13, 0xce, 0x02, 0x41, 0xb2, 0x34, 0x89, 0x45, 0xa0, 0xc8, 0x6c, 0x6a, 0x4c, 0x43, 0xf0,
	0x1e, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x22, 0x41, 0xfd, 0x8a, 0x02, 0x00, 0x00,
}
