// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: calculator/calcpb/sum.proto

package calcpb

import (
	context "context"
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

type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{0}
}

func (x *Sum) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Sum) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum *Sum `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{1}
}

func (x *SumRequest) GetSum() *Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{2}
}

func (x *SumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumber int32 `protobuf:"varint,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeRequest) GetPrimeNumber() int32 {
	if x != nil {
		return x.PrimeNumber
	}
	return 0
}

type PrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{4}
}

func (x *PrimeResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type NumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumberRequest) Reset() {
	*x = NumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberRequest) ProtoMessage() {}

func (x *NumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberRequest.ProtoReflect.Descriptor instead.
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{5}
}

func (x *NumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ComputedResponse) Reset() {
	*x = ComputedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_sum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedResponse) ProtoMessage() {}

func (x *ComputedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_sum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedResponse.ProtoReflect.Descriptor instead.
func (*ComputedResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_sum_proto_rawDescGZIP(), []int{6}
}

func (x *ComputedResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calcpb_sum_proto protoreflect.FileDescriptor

var file_calculator_calcpb_sum_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2f, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63,
	0x61, 0x6c, 0x63, 0x70, 0x62, 0x22, 0x4d, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x03, 0x73, 0x75,
	0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x31, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0d, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2a, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd3, 0x01, 0x0a, 0x0a, 0x53, 0x75,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x19, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x13, 0x5a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calcpb_sum_proto_rawDescOnce sync.Once
	file_calculator_calcpb_sum_proto_rawDescData = file_calculator_calcpb_sum_proto_rawDesc
)

func file_calculator_calcpb_sum_proto_rawDescGZIP() []byte {
	file_calculator_calcpb_sum_proto_rawDescOnce.Do(func() {
		file_calculator_calcpb_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calcpb_sum_proto_rawDescData)
	})
	return file_calculator_calcpb_sum_proto_rawDescData
}

var file_calculator_calcpb_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_calculator_calcpb_sum_proto_goTypes = []interface{}{
	(*Sum)(nil),              // 0: calcpb.Sum
	(*SumRequest)(nil),       // 1: calcpb.SumRequest
	(*SumResponse)(nil),      // 2: calcpb.SumResponse
	(*PrimeRequest)(nil),     // 3: calcpb.PrimeRequest
	(*PrimeResponse)(nil),    // 4: calcpb.PrimeResponse
	(*NumberRequest)(nil),    // 5: calcpb.NumberRequest
	(*ComputedResponse)(nil), // 6: calcpb.ComputedResponse
}
var file_calculator_calcpb_sum_proto_depIdxs = []int32{
	0, // 0: calcpb.SumRequest.sum:type_name -> calcpb.Sum
	1, // 1: calcpb.SumService.Add:input_type -> calcpb.SumRequest
	3, // 2: calcpb.SumService.PrimeNumberDescomposition:input_type -> calcpb.PrimeRequest
	5, // 3: calcpb.SumService.ComputeAverage:input_type -> calcpb.NumberRequest
	2, // 4: calcpb.SumService.Add:output_type -> calcpb.SumResponse
	4, // 5: calcpb.SumService.PrimeNumberDescomposition:output_type -> calcpb.PrimeResponse
	6, // 6: calcpb.SumService.ComputeAverage:output_type -> calcpb.ComputedResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_calcpb_sum_proto_init() }
func file_calculator_calcpb_sum_proto_init() {
	if File_calculator_calcpb_sum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calcpb_sum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeResponse); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberRequest); i {
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
		file_calculator_calcpb_sum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedResponse); i {
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
			RawDescriptor: file_calculator_calcpb_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calcpb_sum_proto_goTypes,
		DependencyIndexes: file_calculator_calcpb_sum_proto_depIdxs,
		MessageInfos:      file_calculator_calcpb_sum_proto_msgTypes,
	}.Build()
	File_calculator_calcpb_sum_proto = out.File
	file_calculator_calcpb_sum_proto_rawDesc = nil
	file_calculator_calcpb_sum_proto_goTypes = nil
	file_calculator_calcpb_sum_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	// unary api
	Add(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server streaming api
	PrimeNumberDescomposition(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDescompositionClient, error)
	// Client streaming api
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Add(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calcpb.SumService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) PrimeNumberDescomposition(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDescompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/calcpb.SumService/PrimeNumberDescomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimeNumberDescompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimeNumberDescompositionClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type sumServicePrimeNumberDescompositionClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimeNumberDescompositionClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[1], "/calcpb.SumService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceComputeAverageClient{stream}
	return x, nil
}

type SumService_ComputeAverageClient interface {
	Send(*NumberRequest) error
	CloseAndRecv() (*ComputedResponse, error)
	grpc.ClientStream
}

type sumServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *sumServiceComputeAverageClient) Send(m *NumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceComputeAverageClient) CloseAndRecv() (*ComputedResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	// unary api
	Add(context.Context, *SumRequest) (*SumResponse, error)
	// Server streaming api
	PrimeNumberDescomposition(*PrimeRequest, SumService_PrimeNumberDescompositionServer) error
	// Client streaming api
	ComputeAverage(SumService_ComputeAverageServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Add(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedSumServiceServer) PrimeNumberDescomposition(*PrimeRequest, SumService_PrimeNumberDescompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDescomposition not implemented")
}
func (*UnimplementedSumServiceServer) ComputeAverage(SumService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calcpb.SumService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Add(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_PrimeNumberDescomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).PrimeNumberDescomposition(m, &sumServicePrimeNumberDescompositionServer{stream})
}

type SumService_PrimeNumberDescompositionServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type sumServicePrimeNumberDescompositionServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimeNumberDescompositionServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SumService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).ComputeAverage(&sumServiceComputeAverageServer{stream})
}

type SumService_ComputeAverageServer interface {
	SendAndClose(*ComputedResponse) error
	Recv() (*NumberRequest, error)
	grpc.ServerStream
}

type sumServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *sumServiceComputeAverageServer) SendAndClose(m *ComputedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceComputeAverageServer) Recv() (*NumberRequest, error) {
	m := new(NumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calcpb.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _SumService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDescomposition",
			Handler:       _SumService_PrimeNumberDescomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _SumService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/sum.proto",
}
