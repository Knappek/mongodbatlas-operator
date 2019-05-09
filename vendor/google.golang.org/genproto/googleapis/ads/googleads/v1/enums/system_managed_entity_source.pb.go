// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/system_managed_entity_source.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

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

// Enum listing the possible system managed entity sources.
type SystemManagedResourceSourceEnum_SystemManagedResourceSource int32

const (
	// Not specified.
	SystemManagedResourceSourceEnum_UNSPECIFIED SystemManagedResourceSourceEnum_SystemManagedResourceSource = 0
	// Used for return value only. Represents value unknown in this version.
	SystemManagedResourceSourceEnum_UNKNOWN SystemManagedResourceSourceEnum_SystemManagedResourceSource = 1
	// Generated ad variations experiment ad.
	SystemManagedResourceSourceEnum_AD_VARIATIONS SystemManagedResourceSourceEnum_SystemManagedResourceSource = 2
)

var SystemManagedResourceSourceEnum_SystemManagedResourceSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_VARIATIONS",
}
var SystemManagedResourceSourceEnum_SystemManagedResourceSource_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"AD_VARIATIONS": 2,
}

func (x SystemManagedResourceSourceEnum_SystemManagedResourceSource) String() string {
	return proto.EnumName(SystemManagedResourceSourceEnum_SystemManagedResourceSource_name, int32(x))
}
func (SystemManagedResourceSourceEnum_SystemManagedResourceSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_system_managed_entity_source_2cb6e967d17e4e0c, []int{0, 0}
}

// Container for enum describing possible system managed entity sources.
type SystemManagedResourceSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemManagedResourceSourceEnum) Reset()         { *m = SystemManagedResourceSourceEnum{} }
func (m *SystemManagedResourceSourceEnum) String() string { return proto.CompactTextString(m) }
func (*SystemManagedResourceSourceEnum) ProtoMessage()    {}
func (*SystemManagedResourceSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_managed_entity_source_2cb6e967d17e4e0c, []int{0}
}
func (m *SystemManagedResourceSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Unmarshal(m, b)
}
func (m *SystemManagedResourceSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Marshal(b, m, deterministic)
}
func (dst *SystemManagedResourceSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemManagedResourceSourceEnum.Merge(dst, src)
}
func (m *SystemManagedResourceSourceEnum) XXX_Size() int {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Size(m)
}
func (m *SystemManagedResourceSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemManagedResourceSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SystemManagedResourceSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SystemManagedResourceSourceEnum)(nil), "google.ads.googleads.v1.enums.SystemManagedResourceSourceEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource", SystemManagedResourceSourceEnum_SystemManagedResourceSource_name, SystemManagedResourceSourceEnum_SystemManagedResourceSource_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/system_managed_entity_source.proto", fileDescriptor_system_managed_entity_source_2cb6e967d17e4e0c)
}

var fileDescriptor_system_managed_entity_source_2cb6e967d17e4e0c = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0xfd, 0xd7, 0x1f, 0x14, 0x32, 0xc4, 0xd9, 0x4b, 0x75, 0xca, 0xf6, 0x00, 0x29, 0xc5, 0xbb,
	0x78, 0x63, 0xe6, 0xea, 0x28, 0x62, 0x37, 0x56, 0x57, 0x41, 0x0a, 0x25, 0x2e, 0x21, 0x14, 0xd6,
	0x64, 0xee, 0xcb, 0x06, 0x7b, 0x1d, 0x2f, 0x7d, 0x14, 0x1f, 0xc5, 0x5b, 0x5f, 0x40, 0x9a, 0xb8,
	0x82, 0x17, 0xee, 0x26, 0x1c, 0xf2, 0x9d, 0xef, 0x9c, 0xf3, 0x1d, 0x74, 0x23, 0xb5, 0x96, 0x0b,
	0x11, 0x30, 0x0e, 0x81, 0x83, 0x35, 0xda, 0x84, 0x81, 0x50, 0xeb, 0x0a, 0x02, 0xd8, 0x82, 0x11,
	0x55, 0x51, 0x31, 0xc5, 0xa4, 0xe0, 0x85, 0x50, 0xa6, 0x34, 0xdb, 0x02, 0xf4, 0x7a, 0x35, 0x17,
	0x78, 0xb9, 0xd2, 0x46, 0xfb, 0x5d, 0xb7, 0x86, 0x19, 0x07, 0xdc, 0x28, 0xe0, 0x4d, 0x88, 0xad,
	0xc2, 0xe9, 0xf9, 0xce, 0x60, 0x59, 0x06, 0x4c, 0x29, 0x6d, 0x98, 0x29, 0xb5, 0x02, 0xb7, 0xdc,
	0x7f, 0x45, 0x97, 0xa9, 0xb5, 0x78, 0x70, 0x0e, 0x53, 0xe1, 0xb4, 0x53, 0xfb, 0x46, 0x6a, 0x5d,
	0xf5, 0x13, 0x74, 0xb6, 0x87, 0xe2, 0x1f, 0xa3, 0xf6, 0x2c, 0x49, 0x27, 0xd1, 0x6d, 0x7c, 0x17,
	0x47, 0xc3, 0xce, 0x3f, 0xbf, 0x8d, 0x0e, 0x67, 0xc9, 0x7d, 0x32, 0x7e, 0x4a, 0x3a, 0x2d, 0xff,
	0x04, 0x1d, 0xd1, 0x61, 0x91, 0xd1, 0x69, 0x4c, 0x1f, 0xe3, 0x71, 0x92, 0x76, 0xbc, 0xc1, 0x57,
	0x0b, 0xf5, 0xe6, 0xba, 0xc2, 0x7b, 0x63, 0x0f, 0x2e, 0x7e, 0x79, 0x46, 0xf6, 0x6e, 0xe7, 0x38,
	0xa9, 0x83, 0x4f, 0x5a, 0xcf, 0x83, 0x1f, 0x01, 0xa9, 0x17, 0x4c, 0x49, 0xac, 0x57, 0x32, 0x90,
	0x42, 0xd9, 0xb3, 0x76, 0x4d, 0x2e, 0x4b, 0xf8, 0xa3, 0xd8, 0x6b, 0xfb, 0xbe, 0x79, 0xff, 0x47,
	0x94, 0xbe, 0x7b, 0xdd, 0x91, 0x93, 0xa2, 0x1c, 0xb0, 0x83, 0x35, 0xca, 0x42, 0x5c, 0x17, 0x00,
	0x1f, 0xbb, 0x79, 0x4e, 0x39, 0xe4, 0xcd, 0x3c, 0xcf, 0xc2, 0xdc, 0xce, 0x3f, 0xbd, 0x9e, 0xfb,
	0x24, 0x84, 0x72, 0x20, 0xa4, 0x61, 0x10, 0x92, 0x85, 0x84, 0x58, 0xce, 0xcb, 0x81, 0x0d, 0x76,
	0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x52, 0x59, 0xad, 0xca, 0xf0, 0x01, 0x00, 0x00,
}
