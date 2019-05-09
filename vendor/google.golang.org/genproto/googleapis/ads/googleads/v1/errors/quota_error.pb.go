// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/quota_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible quota errors.
type QuotaErrorEnum_QuotaError int32

const (
	// Enum unspecified.
	QuotaErrorEnum_UNSPECIFIED QuotaErrorEnum_QuotaError = 0
	// The received error code is not known in this version.
	QuotaErrorEnum_UNKNOWN QuotaErrorEnum_QuotaError = 1
	// Too many requests.
	QuotaErrorEnum_RESOURCE_EXHAUSTED QuotaErrorEnum_QuotaError = 2
	// Access is prohibited.
	QuotaErrorEnum_ACCESS_PROHIBITED QuotaErrorEnum_QuotaError = 3
	// Too many requests in a short amount of time.
	QuotaErrorEnum_RESOURCE_TEMPORARILY_EXHAUSTED QuotaErrorEnum_QuotaError = 4
)

var QuotaErrorEnum_QuotaError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "RESOURCE_EXHAUSTED",
	3: "ACCESS_PROHIBITED",
	4: "RESOURCE_TEMPORARILY_EXHAUSTED",
}
var QuotaErrorEnum_QuotaError_value = map[string]int32{
	"UNSPECIFIED":                    0,
	"UNKNOWN":                        1,
	"RESOURCE_EXHAUSTED":             2,
	"ACCESS_PROHIBITED":              3,
	"RESOURCE_TEMPORARILY_EXHAUSTED": 4,
}

func (x QuotaErrorEnum_QuotaError) String() string {
	return proto.EnumName(QuotaErrorEnum_QuotaError_name, int32(x))
}
func (QuotaErrorEnum_QuotaError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_quota_error_731d878185294ca1, []int{0, 0}
}

// Container for enum describing possible quota errors.
type QuotaErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaErrorEnum) Reset()         { *m = QuotaErrorEnum{} }
func (m *QuotaErrorEnum) String() string { return proto.CompactTextString(m) }
func (*QuotaErrorEnum) ProtoMessage()    {}
func (*QuotaErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_quota_error_731d878185294ca1, []int{0}
}
func (m *QuotaErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaErrorEnum.Unmarshal(m, b)
}
func (m *QuotaErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaErrorEnum.Marshal(b, m, deterministic)
}
func (dst *QuotaErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaErrorEnum.Merge(dst, src)
}
func (m *QuotaErrorEnum) XXX_Size() int {
	return xxx_messageInfo_QuotaErrorEnum.Size(m)
}
func (m *QuotaErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QuotaErrorEnum)(nil), "google.ads.googleads.v1.errors.QuotaErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.QuotaErrorEnum_QuotaError", QuotaErrorEnum_QuotaError_name, QuotaErrorEnum_QuotaError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/quota_error.proto", fileDescriptor_quota_error_731d878185294ca1)
}

var fileDescriptor_quota_error_731d878185294ca1 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4e, 0xb3, 0x40,
	0x14, 0xc5, 0x3f, 0xe8, 0x17, 0x4d, 0xa6, 0x89, 0xad, 0x93, 0xe8, 0xc2, 0x98, 0x2e, 0x78, 0x80,
	0x41, 0xe2, 0x6e, 0x5c, 0x4d, 0xe9, 0xd8, 0x12, 0x15, 0x10, 0x4a, 0xfd, 0x13, 0x92, 0x06, 0xa5,
	0x99, 0x34, 0x69, 0x67, 0x2a, 0x43, 0xbb, 0xf3, 0x19, 0x7c, 0x07, 0x97, 0x3e, 0x8a, 0x8f, 0xd2,
	0xa7, 0x30, 0xc3, 0x08, 0x75, 0xa3, 0x2b, 0x0e, 0x37, 0xbf, 0x73, 0xe6, 0x9e, 0x0b, 0xce, 0x98,
	0x10, 0x6c, 0x31, 0xb3, 0xb3, 0x5c, 0xda, 0x5a, 0x2a, 0xb5, 0x71, 0xec, 0x59, 0x51, 0x88, 0x42,
	0xda, 0x2f, 0x6b, 0x51, 0x66, 0xd3, 0xea, 0x07, 0xad, 0x0a, 0x51, 0x0a, 0xd8, 0xd3, 0x18, 0xca,
	0x72, 0x89, 0x1a, 0x07, 0xda, 0x38, 0x48, 0x3b, 0x4e, 0x4e, 0xeb, 0xc4, 0xd5, 0xdc, 0xce, 0x38,
	0x17, 0x65, 0x56, 0xce, 0x05, 0x97, 0xda, 0x6d, 0xbd, 0x19, 0xe0, 0xe0, 0x56, 0x65, 0x52, 0x45,
	0x53, 0xbe, 0x5e, 0x5a, 0xaf, 0x00, 0xec, 0x26, 0xb0, 0x03, 0xda, 0x89, 0x1f, 0x87, 0xd4, 0xf5,
	0x2e, 0x3d, 0x3a, 0xe8, 0xfe, 0x83, 0x6d, 0xb0, 0x9f, 0xf8, 0x57, 0x7e, 0x70, 0xe7, 0x77, 0x0d,
	0x78, 0x0c, 0x60, 0x44, 0xe3, 0x20, 0x89, 0x5c, 0x3a, 0xa5, 0xf7, 0x23, 0x92, 0xc4, 0x63, 0x3a,
	0xe8, 0x9a, 0xf0, 0x08, 0x1c, 0x12, 0xd7, 0xa5, 0x71, 0x3c, 0x0d, 0xa3, 0x60, 0xe4, 0xf5, 0x3d,
	0x35, 0x6e, 0x41, 0x0b, 0xf4, 0x1a, 0x7c, 0x4c, 0x6f, 0xc2, 0x20, 0x22, 0x91, 0x77, 0xfd, 0xf0,
	0xc3, 0xfa, 0xbf, 0xbf, 0x35, 0x80, 0xf5, 0x2c, 0x96, 0xe8, 0xef, 0x5a, 0xfd, 0xce, 0x6e, 0xc7,
	0x50, 0x35, 0x09, 0x8d, 0xc7, 0xc1, 0xb7, 0x85, 0x89, 0x45, 0xc6, 0x19, 0x12, 0x05, 0xb3, 0xd9,
	0x8c, 0x57, 0x3d, 0xeb, 0x5b, 0xae, 0xe6, 0xf2, 0xb7, 0xd3, 0x5e, 0xe8, 0xcf, 0xbb, 0xd9, 0x1a,
	0x12, 0xf2, 0x61, 0xf6, 0x86, 0x3a, 0x8c, 0xe4, 0x12, 0x69, 0xa9, 0xd4, 0xc4, 0x41, 0xd5, 0x93,
	0xf2, 0xb3, 0x06, 0x52, 0x92, 0xcb, 0xb4, 0x01, 0xd2, 0x89, 0x93, 0x6a, 0x60, 0x6b, 0x5a, 0x7a,
	0x8a, 0x31, 0xc9, 0x25, 0xc6, 0x0d, 0x82, 0xf1, 0xc4, 0xc1, 0x58, 0x43, 0x4f, 0x7b, 0xd5, 0x76,
	0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x72, 0x66, 0x2a, 0xf7, 0x01, 0x00, 0x00,
}
