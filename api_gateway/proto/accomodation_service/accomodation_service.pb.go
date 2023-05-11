// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: accomodation_service.proto

package accomodation

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Location     string `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	Wifi         bool   `protobuf:"varint,4,opt,name=Wifi,proto3" json:"Wifi,omitempty"`
	Kitchen      bool   `protobuf:"varint,5,opt,name=Kitchen,proto3" json:"Kitchen,omitempty"`
	AirCondition bool   `protobuf:"varint,6,opt,name=AirCondition,proto3" json:"AirCondition,omitempty"`
	FreeParking  bool   `protobuf:"varint,7,opt,name=FreeParking,proto3" json:"FreeParking,omitempty"`
	AutoApproval bool   `protobuf:"varint,8,opt,name=AutoApproval,proto3" json:"AutoApproval,omitempty"`
	Photos       []byte `protobuf:"bytes,9,opt,name=Photos,proto3" json:"Photos,omitempty"`
	MinGuests    int64  `protobuf:"varint,10,opt,name=MinGuests,proto3" json:"MinGuests,omitempty"`
	MaxGuests    int64  `protobuf:"varint,11,opt,name=MaxGuests,proto3" json:"MaxGuests,omitempty"`
	IDHost       string `protobuf:"bytes,12,opt,name=IDHost,proto3" json:"IDHost,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateRequest) GetWifi() bool {
	if x != nil {
		return x.Wifi
	}
	return false
}

func (x *CreateRequest) GetKitchen() bool {
	if x != nil {
		return x.Kitchen
	}
	return false
}

func (x *CreateRequest) GetAirCondition() bool {
	if x != nil {
		return x.AirCondition
	}
	return false
}

func (x *CreateRequest) GetFreeParking() bool {
	if x != nil {
		return x.FreeParking
	}
	return false
}

func (x *CreateRequest) GetAutoApproval() bool {
	if x != nil {
		return x.AutoApproval
	}
	return false
}

func (x *CreateRequest) GetPhotos() []byte {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *CreateRequest) GetMinGuests() int64 {
	if x != nil {
		return x.MinGuests
	}
	return 0
}

func (x *CreateRequest) GetMaxGuests() int64 {
	if x != nil {
		return x.MaxGuests
	}
	return 0
}

func (x *CreateRequest) GetIDHost() string {
	if x != nil {
		return x.IDHost
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_accomodation_service_proto protoreflect.FileDescriptor

var file_accomodation_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x69,
	0x66, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x57, 0x69, 0x66, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x69, 0x72, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x41, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x72, 0x65, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x46, 0x72, 0x65, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69,
	0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d,
	0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d, 0x61, 0x78,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x44, 0x48, 0x6f, 0x73, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x44, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x2a,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x7c, 0x0a, 0x14, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x64, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x4d, 0x4c, 0x2d, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accomodation_service_proto_rawDescOnce sync.Once
	file_accomodation_service_proto_rawDescData = file_accomodation_service_proto_rawDesc
)

func file_accomodation_service_proto_rawDescGZIP() []byte {
	file_accomodation_service_proto_rawDescOnce.Do(func() {
		file_accomodation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accomodation_service_proto_rawDescData)
	})
	return file_accomodation_service_proto_rawDescData
}

var file_accomodation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_accomodation_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: accomodation.CreateRequest
	(*CreateResponse)(nil), // 1: accomodation.CreateResponse
}
var file_accomodation_service_proto_depIdxs = []int32{
	0, // 0: accomodation.AccommodationService.Create:input_type -> accomodation.CreateRequest
	1, // 1: accomodation.AccommodationService.Create:output_type -> accomodation.CreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accomodation_service_proto_init() }
func file_accomodation_service_proto_init() {
	if File_accomodation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accomodation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_accomodation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_accomodation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accomodation_service_proto_goTypes,
		DependencyIndexes: file_accomodation_service_proto_depIdxs,
		MessageInfos:      file_accomodation_service_proto_msgTypes,
	}.Build()
	File_accomodation_service_proto = out.File
	file_accomodation_service_proto_rawDesc = nil
	file_accomodation_service_proto_goTypes = nil
	file_accomodation_service_proto_depIdxs = nil
}
