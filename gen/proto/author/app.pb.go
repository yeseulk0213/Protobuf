// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: proto/author/app.proto

package grpc_author

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AppRes_Status int32

const (
	AppRes_OK    AppRes_Status = 0
	AppRes_ERROR AppRes_Status = 1
)

// Enum value maps for AppRes_Status.
var (
	AppRes_Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	AppRes_Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x AppRes_Status) Enum() *AppRes_Status {
	p := new(AppRes_Status)
	*p = x
	return p
}

func (x AppRes_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppRes_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_author_app_proto_enumTypes[0].Descriptor()
}

func (AppRes_Status) Type() protoreflect.EnumType {
	return &file_proto_author_app_proto_enumTypes[0]
}

func (x AppRes_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppRes_Status.Descriptor instead.
func (AppRes_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_author_app_proto_rawDescGZIP(), []int{1, 0}
}

type AppReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      uint32               `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	NameSpace  string               `protobuf:"bytes,1,opt,name=name_space,json=nameSpace,proto3" json:"name_space,omitempty"`
	Traffics   []*AppReq_AppTraffic `protobuf:"bytes,3,rep,name=traffics,proto3" json:"traffics,omitempty"`
	Operations []*AppReq_Operation  `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *AppReq) Reset() {
	*x = AppReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReq) ProtoMessage() {}

func (x *AppReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReq.ProtoReflect.Descriptor instead.
func (*AppReq) Descriptor() ([]byte, []int) {
	return file_proto_author_app_proto_rawDescGZIP(), []int{0}
}

func (x *AppReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AppReq) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

func (x *AppReq) GetTraffics() []*AppReq_AppTraffic {
	if x != nil {
		return x.Traffics
	}
	return nil
}

func (x *AppReq) GetOperations() []*AppReq_Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type AppRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppRes) Reset() {
	*x = AppRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRes) ProtoMessage() {}

func (x *AppRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRes.ProtoReflect.Descriptor instead.
func (*AppRes) Descriptor() ([]byte, []int) {
	return file_proto_author_app_proto_rawDescGZIP(), []int{1}
}

type AppReq_AppTraffic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit  string `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Value uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Seq   uint32 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *AppReq_AppTraffic) Reset() {
	*x = AppReq_AppTraffic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppReq_AppTraffic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReq_AppTraffic) ProtoMessage() {}

func (x *AppReq_AppTraffic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReq_AppTraffic.ProtoReflect.Descriptor instead.
func (*AppReq_AppTraffic) Descriptor() ([]byte, []int) {
	return file_proto_author_app_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AppReq_AppTraffic) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *AppReq_AppTraffic) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AppReq_AppTraffic) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type AppReq_Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndPoint    string `protobuf:"bytes,1,opt,name=end_point,json=endPoint,proto3" json:"end_point,omitempty"`
	OperationId uint32 `protobuf:"varint,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *AppReq_Operation) Reset() {
	*x = AppReq_Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppReq_Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReq_Operation) ProtoMessage() {}

func (x *AppReq_Operation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReq_Operation.ProtoReflect.Descriptor instead.
func (*AppReq_Operation) Descriptor() ([]byte, []int) {
	return file_proto_author_app_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AppReq_Operation) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *AppReq_Operation) GetOperationId() uint32 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

var File_proto_author_app_proto protoreflect.FileDescriptor

var file_proto_author_app_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x2e, 0x41, 0x70,
	0x70, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x08, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x48, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x1a, 0x4b, 0x0a, 0x09, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x22, 0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32,
	0xa9, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_author_app_proto_rawDescOnce sync.Once
	file_proto_author_app_proto_rawDescData = file_proto_author_app_proto_rawDesc
)

func file_proto_author_app_proto_rawDescGZIP() []byte {
	file_proto_author_app_proto_rawDescOnce.Do(func() {
		file_proto_author_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_author_app_proto_rawDescData)
	})
	return file_proto_author_app_proto_rawDescData
}

var file_proto_author_app_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_author_app_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_author_app_proto_goTypes = []interface{}{
	(AppRes_Status)(0),        // 0: grpc_author.AppRes.Status
	(*AppReq)(nil),            // 1: grpc_author.AppReq
	(*AppRes)(nil),            // 2: grpc_author.AppRes
	(*AppReq_AppTraffic)(nil), // 3: grpc_author.AppReq.AppTraffic
	(*AppReq_Operation)(nil),  // 4: grpc_author.AppReq.Operation
}
var file_proto_author_app_proto_depIdxs = []int32{
	3, // 0: grpc_author.AppReq.traffics:type_name -> grpc_author.AppReq.AppTraffic
	4, // 1: grpc_author.AppReq.operations:type_name -> grpc_author.AppReq.Operation
	1, // 2: grpc_author.AppManager.Create:input_type -> grpc_author.AppReq
	1, // 3: grpc_author.AppManager.Update:input_type -> grpc_author.AppReq
	1, // 4: grpc_author.AppManager.Destroy:input_type -> grpc_author.AppReq
	2, // 5: grpc_author.AppManager.Create:output_type -> grpc_author.AppRes
	2, // 6: grpc_author.AppManager.Update:output_type -> grpc_author.AppRes
	2, // 7: grpc_author.AppManager.Destroy:output_type -> grpc_author.AppRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_author_app_proto_init() }
func file_proto_author_app_proto_init() {
	if File_proto_author_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_author_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_author_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_author_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppReq_AppTraffic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_author_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppReq_Operation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_author_app_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_author_app_proto_goTypes,
		DependencyIndexes: file_proto_author_app_proto_depIdxs,
		EnumInfos:         file_proto_author_app_proto_enumTypes,
		MessageInfos:      file_proto_author_app_proto_msgTypes,
	}.Build()
	File_proto_author_app_proto = out.File
	file_proto_author_app_proto_rawDesc = nil
	file_proto_author_app_proto_goTypes = nil
	file_proto_author_app_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AppManagerClient is the client API for AppManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppManagerClient interface {
	Create(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error)
	Update(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error)
	Destroy(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error)
}

type appManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAppManagerClient(cc grpc.ClientConnInterface) AppManagerClient {
	return &appManagerClient{cc}
}

func (c *appManagerClient) Create(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error) {
	out := new(AppRes)
	err := c.cc.Invoke(ctx, "/grpc_author.AppManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) Update(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error) {
	out := new(AppRes)
	err := c.cc.Invoke(ctx, "/grpc_author.AppManager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) Destroy(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppRes, error) {
	out := new(AppRes)
	err := c.cc.Invoke(ctx, "/grpc_author.AppManager/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppManagerServer is the server API for AppManager service.
type AppManagerServer interface {
	Create(context.Context, *AppReq) (*AppRes, error)
	Update(context.Context, *AppReq) (*AppRes, error)
	Destroy(context.Context, *AppReq) (*AppRes, error)
}

// UnimplementedAppManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAppManagerServer struct {
}

func (*UnimplementedAppManagerServer) Create(context.Context, *AppReq) (*AppRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAppManagerServer) Update(context.Context, *AppReq) (*AppRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAppManagerServer) Destroy(context.Context, *AppReq) (*AppRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}

func RegisterAppManagerServer(s *grpc.Server, srv AppManagerServer) {
	s.RegisterService(&_AppManager_serviceDesc, srv)
}

func _AppManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_author.AppManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).Create(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_author.AppManager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).Update(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_author.AppManager/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).Destroy(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_author.AppManager",
	HandlerType: (*AppManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AppManager_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AppManager_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _AppManager_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/author/app.proto",
}