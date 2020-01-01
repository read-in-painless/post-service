// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post.proto

package post

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_PENDING    Status = 0
	Status_ACCEPTED   Status = 1
	Status_MALFORMED  Status = 2
	Status_INPROGRESS Status = 3
	Status_REJECTED   Status = 4
	Status_INVALID    Status = 5
	Status_PUPLISHING Status = 6
)

var Status_name = map[int32]string{
	0: "PENDING",
	1: "ACCEPTED",
	2: "MALFORMED",
	3: "INPROGRESS",
	4: "REJECTED",
	5: "INVALID",
	6: "PUPLISHING",
}
var Status_value = map[string]int32{
	"PENDING":    0,
	"ACCEPTED":   1,
	"MALFORMED":  2,
	"INPROGRESS": 3,
	"REJECTED":   4,
	"INVALID":    5,
	"PUPLISHING": 6,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_post_79400db183fabfa1, []int{0}
}

type PostRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Date                 uint32   `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	PublishDate          uint32   `protobuf:"varint,5,opt,name=publishDate,proto3" json:"publishDate,omitempty"`
	Tags                 []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_post_79400db183fabfa1, []int{0}
}
func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (dst *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(dst, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PostRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *PostRequest) GetDate() uint32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *PostRequest) GetPublishDate() uint32 {
	if m != nil {
		return m.PublishDate
	}
	return 0
}

func (m *PostRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PostReponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReponse) Reset()         { *m = PostReponse{} }
func (m *PostReponse) String() string { return proto.CompactTextString(m) }
func (*PostReponse) ProtoMessage()    {}
func (*PostReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_post_79400db183fabfa1, []int{1}
}
func (m *PostReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReponse.Unmarshal(m, b)
}
func (m *PostReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReponse.Marshal(b, m, deterministic)
}
func (dst *PostReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReponse.Merge(dst, src)
}
func (m *PostReponse) XXX_Size() int {
	return xxx_messageInfo_PostReponse.Size(m)
}
func (m *PostReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostReponse proto.InternalMessageInfo

func (m *PostReponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PENDING
}

func (m *PostReponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PostRequest)(nil), "PostRequest")
	proto.RegisterType((*PostReponse)(nil), "PostReponse")
	proto.RegisterEnum("Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("proto/post.proto", fileDescriptor_post_79400db183fabfa1) }

var fileDescriptor_post_79400db183fabfa1 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0x35, 0x1a, 0xd7, 0x3a, 0x7e, 0x10, 0x86, 0x52, 0x42, 0x2f, 0x0d, 0x1e, 0x8a, 0xf4, 0x90,
	0x82, 0xfd, 0x05, 0x92, 0x6c, 0x35, 0x45, 0xe3, 0xb2, 0x69, 0x7b, 0x8f, 0xb8, 0x58, 0x41, 0xbb,
	0xa9, 0xbb, 0x39, 0xf4, 0xa7, 0xf4, 0xdf, 0x96, 0x59, 0x15, 0xbc, 0xbd, 0xf7, 0xe6, 0xed, 0xe3,
	0xcd, 0x0e, 0x04, 0xd5, 0x51, 0x5b, 0xfd, 0x5c, 0x69, 0x63, 0x63, 0x07, 0x47, 0x7f, 0x1e, 0xf4,
	0x84, 0x36, 0x56, 0xaa, 0x9f, 0x5a, 0x19, 0x8b, 0xb7, 0xd0, 0xb6, 0x3b, 0xbb, 0x57, 0xa1, 0x17,
	0x79, 0xe3, 0xae, 0x3c, 0x11, 0xbc, 0x03, 0x56, 0x1b, 0x75, 0xcc, 0x36, 0x61, 0xd3, 0xc9, 0x67,
	0x86, 0x08, 0xfe, 0x5a, 0x6f, 0x7e, 0xc3, 0x96, 0x53, 0x1d, 0x26, 0x6d, 0x53, 0x5a, 0x15, 0xfa,
	0x91, 0x37, 0x1e, 0x48, 0x87, 0x31, 0x82, 0x5e, 0x55, 0xaf, 0xf7, 0x3b, 0xf3, 0x95, 0xd2, 0xa8,
	0xed, 0x46, 0xd7, 0x12, 0xbd, 0xb2, 0xe5, 0xd6, 0x84, 0x2c, 0x6a, 0x51, 0x12, 0xe1, 0xd1, 0xfc,
	0x52, 0xad, 0xd2, 0xdf, 0x46, 0xe1, 0x03, 0x30, 0x63, 0x4b, 0x5b, 0x1b, 0xd7, 0x6d, 0x38, 0xe9,
	0xc4, 0x85, 0xa3, 0xf2, 0x2c, 0x63, 0x08, 0x9d, 0x83, 0x32, 0xa6, 0xdc, 0xaa, 0x73, 0xcd, 0x0b,
	0x7d, 0x3a, 0x00, 0x3b, 0x79, 0xb1, 0x07, 0x1d, 0xc1, 0xf3, 0x34, 0xcb, 0x67, 0x41, 0x03, 0xfb,
	0x70, 0x33, 0x4d, 0x12, 0x2e, 0xde, 0x79, 0x1a, 0x78, 0x38, 0x80, 0xee, 0x72, 0xba, 0x78, 0x5d,
	0xc9, 0x25, 0x4f, 0x83, 0x26, 0x0e, 0x01, 0xb2, 0x5c, 0xc8, 0xd5, 0x4c, 0xf2, 0xa2, 0x08, 0x5a,
	0x64, 0x96, 0xfc, 0x8d, 0x27, 0x64, 0xf6, 0x29, 0x27, 0xcb, 0x3f, 0xa7, 0x8b, 0x2c, 0x0d, 0xda,
	0x64, 0x15, 0x1f, 0x62, 0x91, 0x15, 0x73, 0xca, 0x65, 0x93, 0x18, 0x7c, 0x2a, 0x8e, 0x8f, 0xc0,
	0x92, 0xa3, 0xa2, 0xf5, 0xfa, 0xf1, 0xd5, 0x27, 0xdf, 0x5f, 0x98, 0xdb, 0x6b, 0xd4, 0x58, 0x33,
	0x77, 0x8b, 0x97, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x5c, 0x8a, 0xf7, 0x9f, 0x01, 0x00,
	0x00,
}
