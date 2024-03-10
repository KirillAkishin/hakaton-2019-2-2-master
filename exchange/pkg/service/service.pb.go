// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type OHLCV struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Time                 int32    `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Interval             int32    `protobuf:"varint,3,opt,name=Interval,proto3" json:"Interval,omitempty"`
	Open                 float32  `protobuf:"fixed32,4,opt,name=Open,proto3" json:"Open,omitempty"`
	High                 float32  `protobuf:"fixed32,5,opt,name=High,proto3" json:"High,omitempty"`
	Low                  float32  `protobuf:"fixed32,6,opt,name=Low,proto3" json:"Low,omitempty"`
	Close                float32  `protobuf:"fixed32,7,opt,name=Close,proto3" json:"Close,omitempty"`
	Volume               uint32   `protobuf:"varint,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Ticker               string   `protobuf:"bytes,9,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OHLCV) Reset()         { *m = OHLCV{} }
func (m *OHLCV) String() string { return proto.CompactTextString(m) }
func (*OHLCV) ProtoMessage()    {}
func (*OHLCV) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *OHLCV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OHLCV.Unmarshal(m, b)
}
func (m *OHLCV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OHLCV.Marshal(b, m, deterministic)
}
func (m *OHLCV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OHLCV.Merge(m, src)
}
func (m *OHLCV) XXX_Size() int {
	return xxx_messageInfo_OHLCV.Size(m)
}
func (m *OHLCV) XXX_DiscardUnknown() {
	xxx_messageInfo_OHLCV.DiscardUnknown(m)
}

var xxx_messageInfo_OHLCV proto.InternalMessageInfo

func (m *OHLCV) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *OHLCV) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *OHLCV) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *OHLCV) GetOpen() float32 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *OHLCV) GetHigh() float32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *OHLCV) GetLow() float32 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *OHLCV) GetClose() float32 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *OHLCV) GetVolume() uint32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *OHLCV) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

type Deal struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BrokerID             int32    `protobuf:"varint,2,opt,name=BrokerID,proto3" json:"BrokerID,omitempty"`
	ClientID             int32    `protobuf:"varint,3,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Ticker               string   `protobuf:"bytes,4,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	Amount               int32    `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Partial              bool     `protobuf:"varint,6,opt,name=Partial,proto3" json:"Partial,omitempty"`
	Time                 int32    `protobuf:"varint,7,opt,name=Time,proto3" json:"Time,omitempty"`
	Price                float32  `protobuf:"fixed32,8,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Deal) GetBrokerID() int32 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

func (m *Deal) GetClientID() int32 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *Deal) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Deal) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Deal) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *Deal) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Deal) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type DealID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BrokerID             int64    `protobuf:"varint,2,opt,name=BrokerID,proto3" json:"BrokerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealID) Reset()         { *m = DealID{} }
func (m *DealID) String() string { return proto.CompactTextString(m) }
func (*DealID) ProtoMessage()    {}
func (*DealID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *DealID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealID.Unmarshal(m, b)
}
func (m *DealID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealID.Marshal(b, m, deterministic)
}
func (m *DealID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealID.Merge(m, src)
}
func (m *DealID) XXX_Size() int {
	return xxx_messageInfo_DealID.Size(m)
}
func (m *DealID) XXX_DiscardUnknown() {
	xxx_messageInfo_DealID.DiscardUnknown(m)
}

var xxx_messageInfo_DealID proto.InternalMessageInfo

func (m *DealID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *DealID) GetBrokerID() int64 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

type BrokerID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerID) Reset()         { *m = BrokerID{} }
func (m *BrokerID) String() string { return proto.CompactTextString(m) }
func (*BrokerID) ProtoMessage()    {}
func (*BrokerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *BrokerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerID.Unmarshal(m, b)
}
func (m *BrokerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerID.Marshal(b, m, deterministic)
}
func (m *BrokerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerID.Merge(m, src)
}
func (m *BrokerID) XXX_Size() int {
	return xxx_messageInfo_BrokerID.Size(m)
}
func (m *BrokerID) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerID.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerID proto.InternalMessageInfo

func (m *BrokerID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type CancelResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResult) Reset()         { *m = CancelResult{} }
func (m *CancelResult) String() string { return proto.CompactTextString(m) }
func (*CancelResult) ProtoMessage()    {}
func (*CancelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *CancelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelResult.Unmarshal(m, b)
}
func (m *CancelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelResult.Marshal(b, m, deterministic)
}
func (m *CancelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResult.Merge(m, src)
}
func (m *CancelResult) XXX_Size() int {
	return xxx_messageInfo_CancelResult.Size(m)
}
func (m *CancelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResult.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResult proto.InternalMessageInfo

func (m *CancelResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*OHLCV)(nil), "OHLCV")
	proto.RegisterType((*Deal)(nil), "Deal")
	proto.RegisterType((*DealID)(nil), "DealID")
	proto.RegisterType((*BrokerID)(nil), "BrokerID")
	proto.RegisterType((*CancelResult)(nil), "CancelResult")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6a, 0xdc, 0x30,
	0x14, 0x85, 0x47, 0x9e, 0xf1, 0xcf, 0x5c, 0x3a, 0xa5, 0x88, 0x50, 0x84, 0x37, 0x35, 0x5a, 0x79,
	0x65, 0x4a, 0xdb, 0x17, 0x68, 0xed, 0x42, 0x0c, 0x81, 0x04, 0x35, 0x64, 0xef, 0x8a, 0x4b, 0x22,
	0xa2, 0xb1, 0x83, 0xa4, 0x49, 0xfb, 0x16, 0x7d, 0xa7, 0x6e, 0xfa, 0x5a, 0x45, 0x92, 0x3d, 0x35,
	0xed, 0xa2, 0xbb, 0xf3, 0x1d, 0xfd, 0x9c, 0x7b, 0x84, 0xe0, 0x60, 0xd1, 0x3c, 0x2b, 0x89, 0xcd,
	0x93, 0x99, 0xdc, 0xc4, 0x7f, 0x11, 0x48, 0xaf, 0x2f, 0xaf, 0xda, 0x3b, 0xfa, 0x12, 0x92, 0xbe,
	0x63, 0xa4, 0x22, 0xf5, 0x56, 0x24, 0x7d, 0x47, 0x29, 0xec, 0x6e, 0xd5, 0x11, 0x59, 0x52, 0x91,
	0x3a, 0x15, 0x41, 0xd3, 0x12, 0x8a, 0x7e, 0x74, 0x68, 0x9e, 0x07, 0xcd, 0xb6, 0xc1, 0x3f, 0xb3,
	0xdf, 0x7f, 0xfd, 0x84, 0x23, 0xdb, 0x55, 0xa4, 0x4e, 0x44, 0xd0, 0xde, 0xbb, 0x54, 0xf7, 0x0f,
	0x2c, 0x8d, 0x9e, 0xd7, 0xf4, 0x15, 0x6c, 0xaf, 0xa6, 0x6f, 0x2c, 0x0b, 0x96, 0x97, 0xf4, 0x02,
	0xd2, 0x56, 0x4f, 0x16, 0x59, 0x1e, 0xbc, 0x08, 0xf4, 0x35, 0x64, 0x77, 0x93, 0x3e, 0x1d, 0x91,
	0x15, 0x15, 0xa9, 0x0f, 0x62, 0x26, 0xef, 0xdf, 0x2a, 0xf9, 0x88, 0x86, 0xed, 0x2b, 0x52, 0xef,
	0xc5, 0x4c, 0xfc, 0x27, 0x81, 0x5d, 0x87, 0x83, 0xfe, 0xa7, 0x48, 0x09, 0xc5, 0x27, 0x33, 0x3d,
	0xa2, 0xe9, 0xbb, 0xb9, 0xcc, 0x99, 0xfd, 0x5a, 0xab, 0x15, 0x8e, 0xae, 0xef, 0x96, 0x42, 0x0b,
	0xaf, 0x82, 0x76, 0xeb, 0x20, 0xef, 0x7f, 0x3c, 0x4e, 0xa7, 0xd1, 0x85, 0x5a, 0xa9, 0x98, 0x89,
	0x32, 0xc8, 0x6f, 0x06, 0xe3, 0xd4, 0xa0, 0x43, 0xb9, 0x42, 0x2c, 0x78, 0x7e, 0xca, 0x7c, 0xf5,
	0x94, 0x17, 0x90, 0xde, 0x18, 0x25, 0x63, 0xbb, 0x44, 0x44, 0xe0, 0x1f, 0x20, 0xf3, 0x1d, 0xfa,
	0xee, 0xbf, 0x2d, 0xb6, 0x7f, 0x5a, 0xf0, 0xd5, 0xda, 0xdf, 0xe7, 0x78, 0x0d, 0x2f, 0xda, 0x61,
	0x94, 0xa8, 0x05, 0xda, 0x93, 0x0e, 0x53, 0xda, 0x93, 0x94, 0x68, 0x6d, 0xd8, 0x54, 0x88, 0x05,
	0xdf, 0xfd, 0x20, 0x50, 0x7c, 0xfe, 0x2e, 0x1f, 0x86, 0xf1, 0x1e, 0x29, 0x87, 0xfd, 0x17, 0x37,
	0x38, 0x65, 0x9d, 0x92, 0x74, 0xdf, 0x2c, 0xd7, 0x97, 0x59, 0x13, 0x7e, 0x0b, 0xdf, 0xbc, 0x25,
	0xb4, 0x84, 0xac, 0x35, 0x38, 0x38, 0xa4, 0x69, 0xe3, 0xa7, 0x2e, 0xf3, 0x26, 0x0e, 0xcf, 0x37,
	0x94, 0x43, 0x16, 0x63, 0xe9, 0x62, 0x96, 0x87, 0x66, 0x3d, 0x08, 0xdf, 0xd0, 0x37, 0x90, 0x47,
	0x6d, 0xd7, 0x09, 0xf1, 0x2e, 0x1f, 0xf0, 0x35, 0x0b, 0x7f, 0xf4, 0xfd, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x18, 0xf9, 0xdd, 0xb4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	// поток ценовых данных от биржи к брокеру
	// мы каждую секнуду будем получать отсюда событие с ценами, которые броке аггрегирует у себя в минуты и показывает клиентам
	// устанавливается 1 раз брокером
	Statistic(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_StatisticClient, error)
	// отправка на биржу заявки от брокера
	Create(ctx context.Context, in *Deal, opts ...grpc.CallOption) (*DealID, error)
	// отмена заявки
	Cancel(ctx context.Context, in *DealID, opts ...grpc.CallOption) (*CancelResult, error)
	// исполнение заявок от биржи к брокеру
	// устанавливается 1 раз брокером и при исполнении какой-то заявки
	Results(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_ResultsClient, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Statistic(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_StatisticClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exchange_serviceDesc.Streams[0], "/Exchange/Statistic", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeStatisticClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exchange_StatisticClient interface {
	Recv() (*OHLCV, error)
	grpc.ClientStream
}

type exchangeStatisticClient struct {
	grpc.ClientStream
}

func (x *exchangeStatisticClient) Recv() (*OHLCV, error) {
	m := new(OHLCV)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exchangeClient) Create(ctx context.Context, in *Deal, opts ...grpc.CallOption) (*DealID, error) {
	out := new(DealID)
	err := c.cc.Invoke(ctx, "/Exchange/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Cancel(ctx context.Context, in *DealID, opts ...grpc.CallOption) (*CancelResult, error) {
	out := new(CancelResult)
	err := c.cc.Invoke(ctx, "/Exchange/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Results(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_ResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exchange_serviceDesc.Streams[1], "/Exchange/Results", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeResultsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exchange_ResultsClient interface {
	Recv() (*Deal, error)
	grpc.ClientStream
}

type exchangeResultsClient struct {
	grpc.ClientStream
}

func (x *exchangeResultsClient) Recv() (*Deal, error) {
	m := new(Deal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	// поток ценовых данных от биржи к брокеру
	// мы каждую секнуду будем получать отсюда событие с ценами, которые броке аггрегирует у себя в минуты и показывает клиентам
	// устанавливается 1 раз брокером
	Statistic(*BrokerID, Exchange_StatisticServer) error
	// отправка на биржу заявки от брокера
	Create(context.Context, *Deal) (*DealID, error)
	// отмена заявки
	Cancel(context.Context, *DealID) (*CancelResult, error)
	// исполнение заявок от биржи к брокеру
	// устанавливается 1 раз брокером и при исполнении какой-то заявки
	Results(*BrokerID, Exchange_ResultsServer) error
}

// UnimplementedExchangeServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (*UnimplementedExchangeServer) Statistic(req *BrokerID, srv Exchange_StatisticServer) error {
	return status.Errorf(codes.Unimplemented, "method Statistic not implemented")
}
func (*UnimplementedExchangeServer) Create(ctx context.Context, req *Deal) (*DealID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedExchangeServer) Cancel(ctx context.Context, req *DealID) (*CancelResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedExchangeServer) Results(req *BrokerID, srv Exchange_ResultsServer) error {
	return status.Errorf(codes.Unimplemented, "method Results not implemented")
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_Statistic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BrokerID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeServer).Statistic(m, &exchangeStatisticServer{stream})
}

type Exchange_StatisticServer interface {
	Send(*OHLCV) error
	grpc.ServerStream
}

type exchangeStatisticServer struct {
	grpc.ServerStream
}

func (x *exchangeStatisticServer) Send(m *OHLCV) error {
	return x.ServerStream.SendMsg(m)
}

func _Exchange_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Create(ctx, req.(*Deal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Cancel(ctx, req.(*DealID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Results_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BrokerID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeServer).Results(m, &exchangeResultsServer{stream})
}

type Exchange_ResultsServer interface {
	Send(*Deal) error
	grpc.ServerStream
}

type exchangeResultsServer struct {
	grpc.ServerStream
}

func (x *exchangeResultsServer) Send(m *Deal) error {
	return x.ServerStream.SendMsg(m)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Exchange_Create_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Exchange_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Statistic",
			Handler:       _Exchange_Statistic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Results",
			Handler:       _Exchange_Results_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
