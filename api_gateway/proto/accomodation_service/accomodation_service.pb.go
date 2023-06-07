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

	ID            string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Wifi          bool   `protobuf:"varint,4,opt,name=wifi,proto3" json:"wifi,omitempty"`
	Kitchen       bool   `protobuf:"varint,5,opt,name=kitchen,proto3" json:"kitchen,omitempty"`
	AirCondition  bool   `protobuf:"varint,6,opt,name=airCondition,proto3" json:"airCondition,omitempty"`
	FreeParking   bool   `protobuf:"varint,7,opt,name=freeParking,proto3" json:"freeParking,omitempty"`
	AutoApproval  bool   `protobuf:"varint,8,opt,name=autoApproval,proto3" json:"autoApproval,omitempty"`
	Photos        []byte `protobuf:"bytes,9,opt,name=photos,proto3" json:"photos,omitempty"`
	MinGuests     string `protobuf:"bytes,10,opt,name=minGuests,proto3" json:"minGuests,omitempty"`
	MaxGuests     string `protobuf:"bytes,11,opt,name=maxGuests,proto3" json:"maxGuests,omitempty"`
	IDHost        string `protobuf:"bytes,12,opt,name=iDHost,proto3" json:"iDHost,omitempty"`
	PricePerGuest bool   `protobuf:"varint,13,opt,name=pricePerGuest,proto3" json:"pricePerGuest,omitempty"`
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

func (x *CreateRequest) GetMinGuests() string {
	if x != nil {
		return x.MinGuests
	}
	return ""
}

func (x *CreateRequest) GetMaxGuests() string {
	if x != nil {
		return x.MaxGuests
	}
	return ""
}

func (x *CreateRequest) GetIDHost() string {
	if x != nil {
		return x.IDHost
	}
	return ""
}

func (x *CreateRequest) GetPricePerGuest() bool {
	if x != nil {
		return x.PricePerGuest
	}
	return false
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

type UpdateAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	StartDate      string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate        string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	AccomodationId string `protobuf:"bytes,4,opt,name=accomodationId,proto3" json:"accomodationId,omitempty"`
	Price          string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateAvailabilityRequest) Reset() {
	*x = UpdateAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvailabilityRequest) ProtoMessage() {}

func (x *UpdateAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*UpdateAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAvailabilityRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateAvailabilityRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *UpdateAvailabilityRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *UpdateAvailabilityRequest) GetAccomodationId() string {
	if x != nil {
		return x.AccomodationId
	}
	return ""
}

func (x *UpdateAvailabilityRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type UpdateAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateAvailabilityResponse) Reset() {
	*x = UpdateAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvailabilityResponse) ProtoMessage() {}

func (x *UpdateAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*UpdateAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAvailabilityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllAccomodationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId string `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *GetAllAccomodationsRequest) Reset() {
	*x = GetAllAccomodationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAccomodationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAccomodationsRequest) ProtoMessage() {}

func (x *GetAllAccomodationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAccomodationsRequest.ProtoReflect.Descriptor instead.
func (*GetAllAccomodationsRequest) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllAccomodationsRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

type Accomodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Wifi          bool   `protobuf:"varint,4,opt,name=wifi,proto3" json:"wifi,omitempty"`
	Kitchen       bool   `protobuf:"varint,5,opt,name=kitchen,proto3" json:"kitchen,omitempty"`
	AirCondition  bool   `protobuf:"varint,6,opt,name=airCondition,proto3" json:"airCondition,omitempty"`
	FreeParking   bool   `protobuf:"varint,7,opt,name=freeParking,proto3" json:"freeParking,omitempty"`
	AutoApproval  bool   `protobuf:"varint,8,opt,name=autoApproval,proto3" json:"autoApproval,omitempty"`
	Photos        []byte `protobuf:"bytes,9,opt,name=photos,proto3" json:"photos,omitempty"`
	MinGuests     string `protobuf:"bytes,10,opt,name=minGuests,proto3" json:"minGuests,omitempty"`
	MaxGuests     string `protobuf:"bytes,11,opt,name=maxGuests,proto3" json:"maxGuests,omitempty"`
	IdHost        string `protobuf:"bytes,12,opt,name=idHost,proto3" json:"idHost,omitempty"`
	PricePerGuest bool   `protobuf:"varint,13,opt,name=pricePerGuest,proto3" json:"pricePerGuest,omitempty"`
}

func (x *Accomodation) Reset() {
	*x = Accomodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accomodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accomodation) ProtoMessage() {}

func (x *Accomodation) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accomodation.ProtoReflect.Descriptor instead.
func (*Accomodation) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{5}
}

func (x *Accomodation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Accomodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Accomodation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Accomodation) GetWifi() bool {
	if x != nil {
		return x.Wifi
	}
	return false
}

func (x *Accomodation) GetKitchen() bool {
	if x != nil {
		return x.Kitchen
	}
	return false
}

func (x *Accomodation) GetAirCondition() bool {
	if x != nil {
		return x.AirCondition
	}
	return false
}

func (x *Accomodation) GetFreeParking() bool {
	if x != nil {
		return x.FreeParking
	}
	return false
}

func (x *Accomodation) GetAutoApproval() bool {
	if x != nil {
		return x.AutoApproval
	}
	return false
}

func (x *Accomodation) GetPhotos() []byte {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Accomodation) GetMinGuests() string {
	if x != nil {
		return x.MinGuests
	}
	return ""
}

func (x *Accomodation) GetMaxGuests() string {
	if x != nil {
		return x.MaxGuests
	}
	return ""
}

func (x *Accomodation) GetIdHost() string {
	if x != nil {
		return x.IdHost
	}
	return ""
}

func (x *Accomodation) GetPricePerGuest() bool {
	if x != nil {
		return x.PricePerGuest
	}
	return false
}

type GetAllAccomodationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accomodations []*Accomodation `protobuf:"bytes,1,rep,name=accomodations,proto3" json:"accomodations,omitempty"`
}

func (x *GetAllAccomodationsResponse) Reset() {
	*x = GetAllAccomodationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAccomodationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAccomodationsResponse) ProtoMessage() {}

func (x *GetAllAccomodationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAccomodationsResponse.ProtoReflect.Descriptor instead.
func (*GetAllAccomodationsResponse) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllAccomodationsResponse) GetAccomodations() []*Accomodation {
	if x != nil {
		return x.Accomodations
	}
	return nil
}

type GetAllAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccomodationId string `protobuf:"bytes,1,opt,name=accomodationId,proto3" json:"accomodationId,omitempty"`
}

func (x *GetAllAvailabilityRequest) Reset() {
	*x = GetAllAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAvailabilityRequest) ProtoMessage() {}

func (x *GetAllAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*GetAllAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllAvailabilityRequest) GetAccomodationId() string {
	if x != nil {
		return x.AccomodationId
	}
	return ""
}

type Availability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate      string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate        string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	IdAccomodation string `protobuf:"bytes,4,opt,name=idAccomodation,proto3" json:"idAccomodation,omitempty"`
	Price          string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Availability) Reset() {
	*x = Availability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Availability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Availability) ProtoMessage() {}

func (x *Availability) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Availability.ProtoReflect.Descriptor instead.
func (*Availability) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{8}
}

func (x *Availability) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Availability) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Availability) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Availability) GetIdAccomodation() string {
	if x != nil {
		return x.IdAccomodation
	}
	return ""
}

func (x *Availability) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type GetAllAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Availabilities []*Availability `protobuf:"bytes,1,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
}

func (x *GetAllAvailabilityResponse) Reset() {
	*x = GetAllAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accomodation_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAvailabilityResponse) ProtoMessage() {}

func (x *GetAllAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accomodation_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*GetAllAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_accomodation_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllAvailabilityResponse) GetAvailabilities() []*Availability {
	if x != nil {
		return x.Availabilities
	}
	return nil
}

var File_accomodation_service_proto protoreflect.FileDescriptor

var file_accomodation_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69,
	0x66, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x77, 0x69, 0x66, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x69, 0x72, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x61, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x72, 0x65, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69,
	0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x78,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x44, 0x48, 0x6f, 0x73, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x44, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xa1, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x22, 0xf8, 0x02, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x66, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x77, 0x69, 0x66, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x65, 0x50, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x65,
	0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61,
	0x75, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x43,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xcd, 0x04, 0x0a,
	0x14, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01, 0x2a,
	0x22, 0x2e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x64, 0x64, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f,
	0x7b, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d,
	0x12, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x12, 0x9c, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12,
	0x2b, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x7b, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x4d, 0x4c, 0x2d, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_accomodation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_accomodation_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),               // 0: accomodation.CreateRequest
	(*CreateResponse)(nil),              // 1: accomodation.CreateResponse
	(*UpdateAvailabilityRequest)(nil),   // 2: accomodation.UpdateAvailabilityRequest
	(*UpdateAvailabilityResponse)(nil),  // 3: accomodation.UpdateAvailabilityResponse
	(*GetAllAccomodationsRequest)(nil),  // 4: accomodation.GetAllAccomodationsRequest
	(*Accomodation)(nil),                // 5: accomodation.Accomodation
	(*GetAllAccomodationsResponse)(nil), // 6: accomodation.GetAllAccomodationsResponse
	(*GetAllAvailabilityRequest)(nil),   // 7: accomodation.GetAllAvailabilityRequest
	(*Availability)(nil),                // 8: accomodation.Availability
	(*GetAllAvailabilityResponse)(nil),  // 9: accomodation.GetAllAvailabilityResponse
}
var file_accomodation_service_proto_depIdxs = []int32{
	5, // 0: accomodation.GetAllAccomodationsResponse.accomodations:type_name -> accomodation.Accomodation
	8, // 1: accomodation.GetAllAvailabilityResponse.availabilities:type_name -> accomodation.Availability
	0, // 2: accomodation.AccommodationService.Create:input_type -> accomodation.CreateRequest
	2, // 3: accomodation.AccommodationService.UpdateAvailability:input_type -> accomodation.UpdateAvailabilityRequest
	4, // 4: accomodation.AccommodationService.GetAllAccomodations:input_type -> accomodation.GetAllAccomodationsRequest
	7, // 5: accomodation.AccommodationService.GetAllAvailability:input_type -> accomodation.GetAllAvailabilityRequest
	1, // 6: accomodation.AccommodationService.Create:output_type -> accomodation.CreateResponse
	3, // 7: accomodation.AccommodationService.UpdateAvailability:output_type -> accomodation.UpdateAvailabilityResponse
	6, // 8: accomodation.AccommodationService.GetAllAccomodations:output_type -> accomodation.GetAllAccomodationsResponse
	9, // 9: accomodation.AccommodationService.GetAllAvailability:output_type -> accomodation.GetAllAvailabilityResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_accomodation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvailabilityRequest); i {
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
		file_accomodation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvailabilityResponse); i {
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
		file_accomodation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAccomodationsRequest); i {
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
		file_accomodation_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accomodation); i {
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
		file_accomodation_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAccomodationsResponse); i {
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
		file_accomodation_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAvailabilityRequest); i {
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
		file_accomodation_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Availability); i {
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
		file_accomodation_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAvailabilityResponse); i {
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
			NumMessages:   10,
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