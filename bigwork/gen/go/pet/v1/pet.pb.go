// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pet.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

// PetType represents the different types of pets in the pet store.
type PetType int32

const (
	PetType_PET_TYPE_UNSPECIFIED PetType = 0
	PetType_PET_TYPE_CAT         PetType = 1
	PetType_PET_TYPE_DOG         PetType = 2
	PetType_PET_TYPE_SNAKE       PetType = 3
	PetType_PET_TYPE_HAMSTER     PetType = 4
)

var PetType_name = map[int32]string{
	0: "PET_TYPE_UNSPECIFIED",
	1: "PET_TYPE_CAT",
	2: "PET_TYPE_DOG",
	3: "PET_TYPE_SNAKE",
	4: "PET_TYPE_HAMSTER",
}

var PetType_value = map[string]int32{
	"PET_TYPE_UNSPECIFIED": 0,
	"PET_TYPE_CAT":         1,
	"PET_TYPE_DOG":         2,
	"PET_TYPE_SNAKE":       3,
	"PET_TYPE_HAMSTER":     4,
}

func (x PetType) String() string {
	return proto.EnumName(PetType_name, int32(x))
}

func (PetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{0}
}

// Pet represents a pet in the pet store.
type Pet struct {
	PetType              PetType            `protobuf:"varint,1,opt,name=pet_type,json=petType,proto3,enum=pet.v1.PetType" json:"pet_type,omitempty"`
	PetId                string             `protobuf:"bytes,2,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Name                 string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *datetime.DateTime `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{0}
}

func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (m *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(m, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetPetType() PetType {
	if m != nil {
		return m.PetType
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

func (m *Pet) GetPetId() string {
	if m != nil {
		return m.PetId
	}
	return ""
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pet) GetCreatedAt() *datetime.DateTime {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type GetPetRequest struct {
	PetId                string   `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPetRequest) Reset()         { *m = GetPetRequest{} }
func (m *GetPetRequest) String() string { return proto.CompactTextString(m) }
func (*GetPetRequest) ProtoMessage()    {}
func (*GetPetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{1}
}

func (m *GetPetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPetRequest.Unmarshal(m, b)
}
func (m *GetPetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPetRequest.Marshal(b, m, deterministic)
}
func (m *GetPetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPetRequest.Merge(m, src)
}
func (m *GetPetRequest) XXX_Size() int {
	return xxx_messageInfo_GetPetRequest.Size(m)
}
func (m *GetPetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPetRequest proto.InternalMessageInfo

func (m *GetPetRequest) GetPetId() string {
	if m != nil {
		return m.PetId
	}
	return ""
}

type GetPetResponse struct {
	Pet                  *Pet     `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPetResponse) Reset()         { *m = GetPetResponse{} }
func (m *GetPetResponse) String() string { return proto.CompactTextString(m) }
func (*GetPetResponse) ProtoMessage()    {}
func (*GetPetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{2}
}

func (m *GetPetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPetResponse.Unmarshal(m, b)
}
func (m *GetPetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPetResponse.Marshal(b, m, deterministic)
}
func (m *GetPetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPetResponse.Merge(m, src)
}
func (m *GetPetResponse) XXX_Size() int {
	return xxx_messageInfo_GetPetResponse.Size(m)
}
func (m *GetPetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPetResponse proto.InternalMessageInfo

func (m *GetPetResponse) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

type PutPetRequest struct {
	PetType              PetType  `protobuf:"varint,1,opt,name=pet_type,json=petType,proto3,enum=pet.v1.PetType" json:"pet_type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutPetRequest) Reset()         { *m = PutPetRequest{} }
func (m *PutPetRequest) String() string { return proto.CompactTextString(m) }
func (*PutPetRequest) ProtoMessage()    {}
func (*PutPetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{3}
}

func (m *PutPetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutPetRequest.Unmarshal(m, b)
}
func (m *PutPetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutPetRequest.Marshal(b, m, deterministic)
}
func (m *PutPetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutPetRequest.Merge(m, src)
}
func (m *PutPetRequest) XXX_Size() int {
	return xxx_messageInfo_PutPetRequest.Size(m)
}
func (m *PutPetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutPetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutPetRequest proto.InternalMessageInfo

func (m *PutPetRequest) GetPetType() PetType {
	if m != nil {
		return m.PetType
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

func (m *PutPetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PutPetResponse struct {
	Pet                  *Pet     `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutPetResponse) Reset()         { *m = PutPetResponse{} }
func (m *PutPetResponse) String() string { return proto.CompactTextString(m) }
func (*PutPetResponse) ProtoMessage()    {}
func (*PutPetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{4}
}

func (m *PutPetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutPetResponse.Unmarshal(m, b)
}
func (m *PutPetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutPetResponse.Marshal(b, m, deterministic)
}
func (m *PutPetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutPetResponse.Merge(m, src)
}
func (m *PutPetResponse) XXX_Size() int {
	return xxx_messageInfo_PutPetResponse.Size(m)
}
func (m *PutPetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutPetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutPetResponse proto.InternalMessageInfo

func (m *PutPetResponse) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

type DeletePetRequest struct {
	PetId                string   `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePetRequest) Reset()         { *m = DeletePetRequest{} }
func (m *DeletePetRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePetRequest) ProtoMessage()    {}
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{5}
}

func (m *DeletePetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePetRequest.Unmarshal(m, b)
}
func (m *DeletePetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePetRequest.Marshal(b, m, deterministic)
}
func (m *DeletePetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePetRequest.Merge(m, src)
}
func (m *DeletePetRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePetRequest.Size(m)
}
func (m *DeletePetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePetRequest proto.InternalMessageInfo

func (m *DeletePetRequest) GetPetId() string {
	if m != nil {
		return m.PetId
	}
	return ""
}

type DeletePetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePetResponse) Reset()         { *m = DeletePetResponse{} }
func (m *DeletePetResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePetResponse) ProtoMessage()    {}
func (*DeletePetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fa7e3fe0bcd61d, []int{6}
}

func (m *DeletePetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePetResponse.Unmarshal(m, b)
}
func (m *DeletePetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePetResponse.Marshal(b, m, deterministic)
}
func (m *DeletePetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePetResponse.Merge(m, src)
}
func (m *DeletePetResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePetResponse.Size(m)
}
func (m *DeletePetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("pet.v1.PetType", PetType_name, PetType_value)
	proto.RegisterType((*Pet)(nil), "pet.v1.Pet")
	proto.RegisterType((*GetPetRequest)(nil), "pet.v1.GetPetRequest")
	proto.RegisterType((*GetPetResponse)(nil), "pet.v1.GetPetResponse")
	proto.RegisterType((*PutPetRequest)(nil), "pet.v1.PutPetRequest")
	proto.RegisterType((*PutPetResponse)(nil), "pet.v1.PutPetResponse")
	proto.RegisterType((*DeletePetRequest)(nil), "pet.v1.DeletePetRequest")
	proto.RegisterType((*DeletePetResponse)(nil), "pet.v1.DeletePetResponse")
}

func init() { proto.RegisterFile("pet.proto", fileDescriptor_a0fa7e3fe0bcd61d) }

var fileDescriptor_a0fa7e3fe0bcd61d = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x26, 0x21, 0x6d, 0x26, 0x34, 0x5d, 0x86, 0x06, 0x99, 0x48, 0x48, 0x91, 0x0f, 0x28,
	0xf4, 0xe0, 0xa8, 0x81, 0x0b, 0xc7, 0xb4, 0x31, 0x25, 0x42, 0xb4, 0x2b, 0xdb, 0x1c, 0xe0, 0x62,
	0x99, 0x66, 0x54, 0x45, 0x6a, 0xea, 0x6d, 0x32, 0x8d, 0xd4, 0xdf, 0xe0, 0xb7, 0xf8, 0x29, 0xb4,
	0xde, 0xda, 0xb1, 0x23, 0x24, 0xe8, 0x6d, 0xf4, 0xfc, 0xe6, 0xcd, 0x7b, 0x6f, 0x65, 0x68, 0x6b,
	0x62, 0x4f, 0xaf, 0x52, 0x4e, 0xb1, 0x65, 0xc6, 0xcd, 0x49, 0xbf, 0x7f, 0x9d, 0xa6, 0xd7, 0x37,
	0x34, 0xe2, 0x07, 0x4d, 0xa3, 0x79, 0xc2, 0xc4, 0x8b, 0x25, 0x59, 0x8e, 0xfb, 0x4b, 0x40, 0x43,
	0x11, 0xe3, 0x31, 0xec, 0x6b, 0xe2, 0xd8, 0x50, 0x1c, 0x31, 0x10, 0xc3, 0xee, 0xf8, 0xd0, 0xb3,
	0xeb, 0x9e, 0x22, 0x8e, 0x1e, 0x34, 0x05, 0x7b, 0xda, 0x0e, 0xd8, 0x03, 0xa3, 0x1c, 0x2f, 0xe6,
	0x4e, 0x7d, 0x20, 0x86, 0xed, 0xe0, 0x99, 0x26, 0x9e, 0xcd, 0x11, 0xa1, 0x79, 0x9b, 0x2c, 0xc9,
	0x69, 0x64, 0x60, 0x36, 0xe3, 0x07, 0x80, 0xab, 0x15, 0x25, 0x4c, 0xf3, 0x38, 0x61, 0xa7, 0x39,
	0x10, 0xc3, 0xce, 0xb8, 0xe7, 0x59, 0x3f, 0x9e, 0x39, 0xe6, 0x4d, 0x13, 0xa6, 0x68, 0xb1, 0xa4,
	0xa0, 0xfd, 0x48, 0x9c, 0xb0, 0xfb, 0x16, 0x0e, 0xce, 0x89, 0x15, 0x71, 0x40, 0x77, 0xf7, 0xb4,
	0xe6, 0xd2, 0x45, 0x51, 0xba, 0xe8, 0x8e, 0xa0, 0x9b, 0xf3, 0xd6, 0x3a, 0xbd, 0x5d, 0x13, 0xbe,
	0x81, 0x86, 0x26, 0xce, 0x58, 0x9d, 0x71, 0xa7, 0x94, 0x20, 0x30, 0xb8, 0x7b, 0x09, 0x07, 0xea,
	0xbe, 0x2c, 0xfc, 0x94, 0xd8, 0x79, 0xbe, 0xfa, 0x36, 0x9f, 0x71, 0x90, 0x0b, 0xfe, 0x9f, 0x83,
	0x77, 0x20, 0xa7, 0x74, 0x43, 0x4c, 0xff, 0x4e, 0xf7, 0x12, 0x5e, 0x94, 0xa8, 0x56, 0xfe, 0xf8,
	0x0e, 0xf6, 0x1e, 0x8d, 0xa1, 0x03, 0x47, 0xca, 0x8f, 0xe2, 0xe8, 0xbb, 0xf2, 0xe3, 0x6f, 0x17,
	0xa1, 0xf2, 0xcf, 0x66, 0x9f, 0x66, 0xfe, 0x54, 0xd6, 0x50, 0xc2, 0xf3, 0xe2, 0xcb, 0xd9, 0x24,
	0x92, 0xa2, 0x82, 0x4c, 0x2f, 0xcf, 0x65, 0x1d, 0x11, 0xba, 0x05, 0x12, 0x5e, 0x4c, 0xbe, 0xf8,
	0xb2, 0x81, 0x47, 0x20, 0x0b, 0xec, 0xf3, 0xe4, 0x6b, 0x18, 0xf9, 0x81, 0x6c, 0x8e, 0x7f, 0x0b,
	0x38, 0x54, 0xc4, 0x21, 0xa7, 0x2b, 0x0a, 0x69, 0xb5, 0x59, 0x5c, 0x11, 0x7e, 0x84, 0x96, 0x6d,
	0x1e, 0x7b, 0x79, 0xc4, 0xca, 0x8b, 0xf5, 0x5f, 0xed, 0xc2, 0xd6, 0xbf, 0x5b, 0x33, 0xab, 0xb6,
	0xb2, 0xed, 0x6a, 0xe5, 0x4d, 0xb6, 0xab, 0xd5, 0x66, 0xdd, 0x1a, 0x9e, 0x42, 0xbb, 0x68, 0x04,
	0x9d, 0x9c, 0xb6, 0xdb, 0x67, 0xff, 0xf5, 0x5f, 0xbe, 0xe4, 0x1a, 0xa7, 0xfb, 0x3f, 0x4c, 0xd9,
	0xa3, 0xcd, 0xc9, 0xcf, 0x56, 0xf6, 0x07, 0xbc, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0x6a,
	0x96, 0xe4, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PetStoreServiceClient is the client API for PetStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetStoreServiceClient interface {
	GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error)
	PutPet(ctx context.Context, in *PutPetRequest, opts ...grpc.CallOption) (*PutPetResponse, error)
	DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetResponse, error)
}

type petStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetStoreServiceClient(cc grpc.ClientConnInterface) PetStoreServiceClient {
	return &petStoreServiceClient{cc}
}

func (c *petStoreServiceClient) GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error) {
	out := new(GetPetResponse)
	err := c.cc.Invoke(ctx, "/pet.v1.PetStoreService/GetPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) PutPet(ctx context.Context, in *PutPetRequest, opts ...grpc.CallOption) (*PutPetResponse, error) {
	out := new(PutPetResponse)
	err := c.cc.Invoke(ctx, "/pet.v1.PetStoreService/PutPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetResponse, error) {
	out := new(DeletePetResponse)
	err := c.cc.Invoke(ctx, "/pet.v1.PetStoreService/DeletePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServiceServer is the server API for PetStoreService service.
type PetStoreServiceServer interface {
	GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error)
	PutPet(context.Context, *PutPetRequest) (*PutPetResponse, error)
	DeletePet(context.Context, *DeletePetRequest) (*DeletePetResponse, error)
}

// UnimplementedPetStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPetStoreServiceServer struct {
}

func (*UnimplementedPetStoreServiceServer) GetPet(ctx context.Context, req *GetPetRequest) (*GetPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (*UnimplementedPetStoreServiceServer) PutPet(ctx context.Context, req *PutPetRequest) (*PutPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPet not implemented")
}
func (*UnimplementedPetStoreServiceServer) DeletePet(ctx context.Context, req *DeletePetRequest) (*DeletePetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}

func RegisterPetStoreServiceServer(s *grpc.Server, srv PetStoreServiceServer) {
	s.RegisterService(&_PetStoreService_serviceDesc, srv)
}

func _PetStoreService_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pet.v1.PetStoreService/GetPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).GetPet(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_PutPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).PutPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pet.v1.PetStoreService/PutPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).PutPet(ctx, req.(*PutPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pet.v1.PetStoreService/DeletePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).DeletePet(ctx, req.(*DeletePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pet.v1.PetStoreService",
	HandlerType: (*PetStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPet",
			Handler:    _PetStoreService_GetPet_Handler,
		},
		{
			MethodName: "PutPet",
			Handler:    _PetStoreService_PutPet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetStoreService_DeletePet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet.proto",
}
