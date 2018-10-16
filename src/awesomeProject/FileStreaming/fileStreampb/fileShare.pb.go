// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileShare.proto

package FileStreaming

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_OK      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "OK",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"OK":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{0}
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{0}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Folder struct {
	Folder               string   `protobuf:"bytes,1,opt,name=Folder,proto3" json:"Folder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Folder) Reset()         { *m = Folder{} }
func (m *Folder) String() string { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()    {}
func (*Folder) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{1}
}

func (m *Folder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Folder.Unmarshal(m, b)
}
func (m *Folder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Folder.Marshal(b, m, deterministic)
}
func (m *Folder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Folder.Merge(m, src)
}
func (m *Folder) XXX_Size() int {
	return xxx_messageInfo_Folder.Size(m)
}
func (m *Folder) XXX_DiscardUnknown() {
	xxx_messageInfo_Folder.DiscardUnknown(m)
}

var xxx_messageInfo_Folder proto.InternalMessageInfo

func (m *Folder) GetFolder() string {
	if m != nil {
		return m.Folder
	}
	return ""
}

type FileName struct {
	FileName             string   `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Dir                  *Folder  `protobuf:"bytes,2,opt,name=Dir,proto3" json:"Dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileName) Reset()         { *m = FileName{} }
func (m *FileName) String() string { return proto.CompactTextString(m) }
func (*FileName) ProtoMessage()    {}
func (*FileName) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{2}
}

func (m *FileName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileName.Unmarshal(m, b)
}
func (m *FileName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileName.Marshal(b, m, deterministic)
}
func (m *FileName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileName.Merge(m, src)
}
func (m *FileName) XXX_Size() int {
	return xxx_messageInfo_FileName.Size(m)
}
func (m *FileName) XXX_DiscardUnknown() {
	xxx_messageInfo_FileName.DiscardUnknown(m)
}

var xxx_messageInfo_FileName proto.InternalMessageInfo

func (m *FileName) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileName) GetDir() *Folder {
	if m != nil {
		return m.Dir
	}
	return nil
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=FileStreaming.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{3}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

type ChunkPackage struct {
	Batch                *Chunk        `protobuf:"bytes,1,opt,name=Batch,proto3" json:"Batch,omitempty"`
	Status               *UploadStatus `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	TotalSize            int64         `protobuf:"varint,3,opt,name=TotalSize,proto3" json:"TotalSize,omitempty"`
	BytesSent            int64         `protobuf:"varint,4,opt,name=BytesSent,proto3" json:"BytesSent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChunkPackage) Reset()         { *m = ChunkPackage{} }
func (m *ChunkPackage) String() string { return proto.CompactTextString(m) }
func (*ChunkPackage) ProtoMessage()    {}
func (*ChunkPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b00108b092638ad, []int{4}
}

func (m *ChunkPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkPackage.Unmarshal(m, b)
}
func (m *ChunkPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkPackage.Marshal(b, m, deterministic)
}
func (m *ChunkPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkPackage.Merge(m, src)
}
func (m *ChunkPackage) XXX_Size() int {
	return xxx_messageInfo_ChunkPackage.Size(m)
}
func (m *ChunkPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkPackage.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkPackage proto.InternalMessageInfo

func (m *ChunkPackage) GetBatch() *Chunk {
	if m != nil {
		return m.Batch
	}
	return nil
}

func (m *ChunkPackage) GetStatus() *UploadStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ChunkPackage) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ChunkPackage) GetBytesSent() int64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func init() {
	proto.RegisterEnum("FileStreaming.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*Chunk)(nil), "FileStreaming.Chunk")
	proto.RegisterType((*Folder)(nil), "FileStreaming.Folder")
	proto.RegisterType((*FileName)(nil), "FileStreaming.FileName")
	proto.RegisterType((*UploadStatus)(nil), "FileStreaming.UploadStatus")
	proto.RegisterType((*ChunkPackage)(nil), "FileStreaming.ChunkPackage")
}

func init() { proto.RegisterFile("fileShare.proto", fileDescriptor_9b00108b092638ad) }

var fileDescriptor_9b00108b092638ad = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcb, 0x4e, 0xf2, 0x40,
	0x14, 0x80, 0x19, 0x2e, 0xe5, 0xe7, 0xd0, 0x5f, 0x9b, 0x89, 0x97, 0x06, 0x4c, 0xac, 0xdd, 0x48,
	0x58, 0x10, 0xd3, 0x3e, 0x81, 0x60, 0x88, 0x89, 0x51, 0x4c, 0x2b, 0x4b, 0x17, 0x23, 0x1d, 0x69,
	0x43, 0xe9, 0x90, 0xe9, 0x20, 0xd1, 0x97, 0xf1, 0x01, 0x7c, 0x49, 0x33, 0xd3, 0x56, 0xb4, 0x11,
	0x77, 0x73, 0xce, 0xf9, 0xf2, 0xf5, 0x5c, 0x0a, 0xfb, 0xcf, 0x51, 0x4c, 0xfd, 0x90, 0x70, 0x3a,
	0x58, 0x71, 0x26, 0x18, 0xfe, 0x3f, 0x96, 0x09, 0xc1, 0x29, 0x59, 0x46, 0xc9, 0xdc, 0x3e, 0x83,
	0xc6, 0x28, 0x5c, 0x27, 0x0b, 0x6c, 0x42, 0x73, 0xc4, 0x12, 0x41, 0x13, 0x61, 0x22, 0x0b, 0xf5,
	0x74, 0xaf, 0x08, 0x6d, 0x0b, 0xb4, 0x31, 0x8b, 0x03, 0xca, 0xf1, 0x51, 0xf1, 0x52, 0x48, 0xcb,
	0xcb, 0x23, 0x7b, 0x02, 0xff, 0xa4, 0xf5, 0x8e, 0x2c, 0x29, 0xee, 0x6c, 0xdf, 0x39, 0xb5, 0xad,
	0x9d, 0x43, 0xed, 0x2a, 0xe2, 0x66, 0xd5, 0x42, 0xbd, 0xb6, 0x73, 0x38, 0xf8, 0xd1, 0xc9, 0x20,
	0x73, 0x79, 0x92, 0xb0, 0x1f, 0x41, 0x9f, 0xae, 0x62, 0x46, 0x02, 0x5f, 0x10, 0xb1, 0x4e, 0x65,
	0x73, 0xb7, 0x34, 0x4d, 0xc9, 0xbc, 0x70, 0x16, 0x21, 0x76, 0xa1, 0x3e, 0x62, 0x01, 0x55, 0xce,
	0x3d, 0xe7, 0xb4, 0xe4, 0xfc, 0x2e, 0x91, 0x98, 0xa7, 0x60, 0xfb, 0x03, 0x81, 0xae, 0xa6, 0xbe,
	0x27, 0xb3, 0x85, 0xb4, 0xf4, 0xa1, 0x31, 0x24, 0x62, 0x16, 0x2a, 0x7b, 0xdb, 0x39, 0x28, 0x69,
	0x14, 0xeb, 0x65, 0x08, 0x76, 0x41, 0xcb, 0x84, 0xf9, 0x1c, 0xdd, 0x3f, 0xbe, 0xe9, 0xe5, 0x28,
	0x3e, 0x81, 0xd6, 0x03, 0x13, 0x24, 0xf6, 0xa3, 0x37, 0x6a, 0xd6, 0x2c, 0xd4, 0xab, 0x79, 0xdb,
	0x84, 0xac, 0x0e, 0x5f, 0x05, 0x4d, 0x7d, 0xb9, 0xfd, 0x7a, 0x56, 0xfd, 0x4a, 0xf4, 0x5d, 0x30,
	0xca, 0x73, 0xe0, 0x36, 0x34, 0xa7, 0xc9, 0x22, 0x61, 0x9b, 0xc4, 0xa8, 0x60, 0x0d, 0xaa, 0x93,
	0x1b, 0x03, 0x61, 0x00, 0x6d, 0x4c, 0xa2, 0x98, 0x06, 0x46, 0xd5, 0x79, 0x47, 0x60, 0xa8, 0xb3,
	0xab, 0xe6, 0x28, 0x7f, 0x89, 0x66, 0x14, 0x5f, 0x42, 0xcb, 0x0f, 0xd9, 0x46, 0xa6, 0x52, 0x7c,
	0x5c, 0xde, 0x7f, 0x7e, 0xa5, 0xce, 0xae, 0x82, 0x5d, 0xb9, 0x40, 0xf8, 0x1a, 0x74, 0x69, 0xa3,
	0x3c, 0x6b, 0x69, 0xb7, 0xa5, 0xfb, 0xdb, 0x0e, 0xf3, 0x7d, 0x4b, 0xd3, 0x93, 0xa6, 0xfe, 0x47,
	0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0x81, 0x6d, 0xa0, 0xb4, 0xa2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShareFileServiceClient is the client API for ShareFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareFileServiceClient interface {
	// Provide a function that will return the file names of the media folder
	ShowFiles(ctx context.Context, in *FileName, opts ...grpc.CallOption) (ShareFileService_ShowFilesClient, error)
	// 1 step. Client-to-Server streaming. The server streams the file to the client.
	ServerUpload(ctx context.Context, in *FileName, opts ...grpc.CallOption) (ShareFileService_ServerUploadClient, error)
}

type shareFileServiceClient struct {
	cc *grpc.ClientConn
}

func NewShareFileServiceClient(cc *grpc.ClientConn) ShareFileServiceClient {
	return &shareFileServiceClient{cc}
}

func (c *shareFileServiceClient) ShowFiles(ctx context.Context, in *FileName, opts ...grpc.CallOption) (ShareFileService_ShowFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShareFileService_serviceDesc.Streams[0], "/FileStreaming.ShareFileService/ShowFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &shareFileServiceShowFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShareFileService_ShowFilesClient interface {
	Recv() (*FileName, error)
	grpc.ClientStream
}

type shareFileServiceShowFilesClient struct {
	grpc.ClientStream
}

func (x *shareFileServiceShowFilesClient) Recv() (*FileName, error) {
	m := new(FileName)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shareFileServiceClient) ServerUpload(ctx context.Context, in *FileName, opts ...grpc.CallOption) (ShareFileService_ServerUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShareFileService_serviceDesc.Streams[1], "/FileStreaming.ShareFileService/ServerUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &shareFileServiceServerUploadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShareFileService_ServerUploadClient interface {
	Recv() (*ChunkPackage, error)
	grpc.ClientStream
}

type shareFileServiceServerUploadClient struct {
	grpc.ClientStream
}

func (x *shareFileServiceServerUploadClient) Recv() (*ChunkPackage, error) {
	m := new(ChunkPackage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShareFileServiceServer is the server API for ShareFileService service.
type ShareFileServiceServer interface {
	// Provide a function that will return the file names of the media folder
	ShowFiles(*FileName, ShareFileService_ShowFilesServer) error
	// 1 step. Client-to-Server streaming. The server streams the file to the client.
	ServerUpload(*FileName, ShareFileService_ServerUploadServer) error
}

func RegisterShareFileServiceServer(s *grpc.Server, srv ShareFileServiceServer) {
	s.RegisterService(&_ShareFileService_serviceDesc, srv)
}

func _ShareFileService_ShowFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShareFileServiceServer).ShowFiles(m, &shareFileServiceShowFilesServer{stream})
}

type ShareFileService_ShowFilesServer interface {
	Send(*FileName) error
	grpc.ServerStream
}

type shareFileServiceShowFilesServer struct {
	grpc.ServerStream
}

func (x *shareFileServiceShowFilesServer) Send(m *FileName) error {
	return x.ServerStream.SendMsg(m)
}

func _ShareFileService_ServerUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShareFileServiceServer).ServerUpload(m, &shareFileServiceServerUploadServer{stream})
}

type ShareFileService_ServerUploadServer interface {
	Send(*ChunkPackage) error
	grpc.ServerStream
}

type shareFileServiceServerUploadServer struct {
	grpc.ServerStream
}

func (x *shareFileServiceServerUploadServer) Send(m *ChunkPackage) error {
	return x.ServerStream.SendMsg(m)
}

var _ShareFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileStreaming.ShareFileService",
	HandlerType: (*ShareFileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShowFiles",
			Handler:       _ShareFileService_ShowFiles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ServerUpload",
			Handler:       _ShareFileService_ServerUpload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fileShare.proto",
}
