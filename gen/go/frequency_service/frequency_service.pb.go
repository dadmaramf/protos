// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: frequency_service/frequency_service.proto

package frequency_servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FrequencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Timenow   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timenow,proto3" json:"timenow,omitempty"`
	Frequency float64                `protobuf:"fixed64,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *FrequencyResponse) Reset() {
	*x = FrequencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frequency_service_frequency_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyResponse) ProtoMessage() {}

func (x *FrequencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_frequency_service_frequency_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyResponse.ProtoReflect.Descriptor instead.
func (*FrequencyResponse) Descriptor() ([]byte, []int) {
	return file_frequency_service_frequency_service_proto_rawDescGZIP(), []int{0}
}

func (x *FrequencyResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *FrequencyResponse) GetTimenow() *timestamppb.Timestamp {
	if x != nil {
		return x.Timenow
	}
	return nil
}

func (x *FrequencyResponse) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

type FrequencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FrequencyRequest) Reset() {
	*x = FrequencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frequency_service_frequency_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyRequest) ProtoMessage() {}

func (x *FrequencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frequency_service_frequency_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyRequest.ProtoReflect.Descriptor instead.
func (*FrequencyRequest) Descriptor() ([]byte, []int) {
	return file_frequency_service_frequency_service_proto_rawDescGZIP(), []int{1}
}

func (x *FrequencyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_frequency_service_frequency_service_proto protoreflect.FileDescriptor

var file_frequency_service_frequency_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e,
	0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6e,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x6f, 0x77, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x46,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x55, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x12, 0x16, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x64, 0x61, 0x64, 0x6d, 0x61, 0x72,
	0x61, 0x6d, 0x66, 0x2e, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frequency_service_frequency_service_proto_rawDescOnce sync.Once
	file_frequency_service_frequency_service_proto_rawDescData = file_frequency_service_frequency_service_proto_rawDesc
)

func file_frequency_service_frequency_service_proto_rawDescGZIP() []byte {
	file_frequency_service_frequency_service_proto_rawDescOnce.Do(func() {
		file_frequency_service_frequency_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_frequency_service_frequency_service_proto_rawDescData)
	})
	return file_frequency_service_frequency_service_proto_rawDescData
}

var file_frequency_service_frequency_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frequency_service_frequency_service_proto_goTypes = []any{
	(*FrequencyResponse)(nil),     // 0: rand.FrequencyResponse
	(*FrequencyRequest)(nil),      // 1: rand.FrequencyRequest
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_frequency_service_frequency_service_proto_depIdxs = []int32{
	2, // 0: rand.FrequencyResponse.timenow:type_name -> google.protobuf.Timestamp
	1, // 1: rand.StreamingService.NewConnect:input_type -> rand.FrequencyRequest
	0, // 2: rand.StreamingService.NewConnect:output_type -> rand.FrequencyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_frequency_service_frequency_service_proto_init() }
func file_frequency_service_frequency_service_proto_init() {
	if File_frequency_service_frequency_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frequency_service_frequency_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FrequencyResponse); i {
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
		file_frequency_service_frequency_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FrequencyRequest); i {
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
			RawDescriptor: file_frequency_service_frequency_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frequency_service_frequency_service_proto_goTypes,
		DependencyIndexes: file_frequency_service_frequency_service_proto_depIdxs,
		MessageInfos:      file_frequency_service_frequency_service_proto_msgTypes,
	}.Build()
	File_frequency_service_frequency_service_proto = out.File
	file_frequency_service_frequency_service_proto_rawDesc = nil
	file_frequency_service_frequency_service_proto_goTypes = nil
	file_frequency_service_frequency_service_proto_depIdxs = nil
}
