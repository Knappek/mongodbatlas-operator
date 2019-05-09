// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/geo_target_constant.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A geo target constant.
type GeoTargetConstant struct {
	// The resource name of the geo target constant.
	// Geo target constant resource names have the form:
	//
	// `geoTargetConstants/{geo_target_constant_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the geo target constant.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Geo target constant English name.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The ISO-3166-1 alpha-2 country code that is associated with the target.
	CountryCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Geo target constant target type.
	TargetType *wrappers.StringValue `protobuf:"bytes,6,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	// Geo target constant status.
	Status enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus `protobuf:"varint,7,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus" json:"status,omitempty"`
	// The fully qualified English name, consisting of the target's name and that
	// of its parent and country.
	CanonicalName        *wrappers.StringValue `protobuf:"bytes,8,opt,name=canonical_name,json=canonicalName,proto3" json:"canonical_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GeoTargetConstant) Reset()         { *m = GeoTargetConstant{} }
func (m *GeoTargetConstant) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstant) ProtoMessage()    {}
func (*GeoTargetConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_geo_target_constant_45d3d633f73a7113, []int{0}
}
func (m *GeoTargetConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstant.Unmarshal(m, b)
}
func (m *GeoTargetConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstant.Marshal(b, m, deterministic)
}
func (dst *GeoTargetConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstant.Merge(dst, src)
}
func (m *GeoTargetConstant) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstant.Size(m)
}
func (m *GeoTargetConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstant.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstant proto.InternalMessageInfo

func (m *GeoTargetConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GeoTargetConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GeoTargetConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GeoTargetConstant) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *GeoTargetConstant) GetTargetType() *wrappers.StringValue {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *GeoTargetConstant) GetStatus() enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus {
	if m != nil {
		return m.Status
	}
	return enums.GeoTargetConstantStatusEnum_UNSPECIFIED
}

func (m *GeoTargetConstant) GetCanonicalName() *wrappers.StringValue {
	if m != nil {
		return m.CanonicalName
	}
	return nil
}

func init() {
	proto.RegisterType((*GeoTargetConstant)(nil), "google.ads.googleads.v1.resources.GeoTargetConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/geo_target_constant.proto", fileDescriptor_geo_target_constant_45d3d633f73a7113)
}

var fileDescriptor_geo_target_constant_45d3d633f73a7113 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x49, 0xb6, 0xae, 0x3a, 0xdb, 0x16, 0xcc, 0x41, 0x42, 0x2d, 0xb2, 0x55, 0x0a, 0x0b,
	0xc2, 0xc4, 0x54, 0xf1, 0x90, 0xa2, 0x92, 0x2e, 0xb2, 0xe8, 0x41, 0x4a, 0x5a, 0xf6, 0x20, 0x0b,
	0x61, 0x9a, 0x79, 0x0e, 0x91, 0xcd, 0xbc, 0x30, 0x33, 0xa9, 0xec, 0xd9, 0x6f, 0xe2, 0xd1, 0x4f,
	0x22, 0x7e, 0x14, 0x3f, 0x85, 0xec, 0x4c, 0x92, 0xcb, 0xba, 0xba, 0xb7, 0x97, 0xbc, 0xff, 0xef,
	0xff, 0xfe, 0x79, 0x33, 0x21, 0xe7, 0x02, 0x51, 0x2c, 0x21, 0x62, 0x5c, 0x47, 0xae, 0x5c, 0x57,
	0xb7, 0x71, 0xa4, 0x40, 0x63, 0xa3, 0x0a, 0xd0, 0x91, 0x00, 0xcc, 0x0d, 0x53, 0x02, 0x4c, 0x5e,
	0xa0, 0xd4, 0x86, 0x49, 0x43, 0x6b, 0x85, 0x06, 0x83, 0x13, 0x47, 0x50, 0xc6, 0x35, 0xed, 0x61,
	0x7a, 0x1b, 0xd3, 0x1e, 0x3e, 0x7a, 0xb3, 0xcd, 0x1f, 0x64, 0x53, 0xfd, 0xd5, 0x3b, 0xd7, 0x86,
	0x99, 0x46, 0xbb, 0x11, 0x47, 0x8f, 0x5b, 0xde, 0x3e, 0xdd, 0x34, 0x9f, 0xa3, 0xaf, 0x8a, 0xd5,
	0x35, 0xa8, 0xae, 0x7f, 0xdc, 0xf9, 0xd7, 0x65, 0xc4, 0xa4, 0x44, 0xc3, 0x4c, 0x89, 0xb2, 0xed,
	0x3e, 0xf9, 0x39, 0x20, 0x0f, 0x66, 0x80, 0xd7, 0x76, 0xc2, 0xb4, 0x1d, 0x10, 0x3c, 0x25, 0x07,
	0x5d, 0xc0, 0x5c, 0xb2, 0x0a, 0x42, 0x6f, 0xec, 0x4d, 0xee, 0x67, 0xfb, 0xdd, 0xcb, 0x8f, 0xac,
	0x82, 0xe0, 0x19, 0xf1, 0x4b, 0x1e, 0x0e, 0xc6, 0xde, 0x64, 0x74, 0xf6, 0xa8, 0xfd, 0x3a, 0xda,
	0xa5, 0xa0, 0xef, 0xa5, 0x79, 0xf5, 0x72, 0xce, 0x96, 0x0d, 0x64, 0x7e, 0xc9, 0x83, 0xe7, 0x64,
	0xcf, 0x1a, 0xed, 0x59, 0xf9, 0xf1, 0x86, 0xfc, 0xca, 0xa8, 0x52, 0x0a, 0xa7, 0xb7, 0xca, 0xe0,
	0x2d, 0xd9, 0x2f, 0xb0, 0x91, 0x46, 0xad, 0xf2, 0x02, 0x39, 0x84, 0x77, 0x76, 0x20, 0x47, 0x2d,
	0x31, 0x45, 0x0e, 0xc1, 0x6b, 0x32, 0x6a, 0x17, 0x67, 0x56, 0x35, 0x84, 0xc3, 0x1d, 0x78, 0xe2,
	0x80, 0xeb, 0x55, 0x0d, 0xc1, 0x17, 0x32, 0x74, 0x7b, 0x0e, 0xef, 0x8e, 0xbd, 0xc9, 0xe1, 0x59,
	0x46, 0xb7, 0x9d, 0xa5, 0x3d, 0x28, 0xba, 0xb1, 0xc5, 0x2b, 0x4b, 0xbf, 0x93, 0x4d, 0xb5, 0xad,
	0x97, 0xb5, 0x13, 0x82, 0x29, 0x39, 0x2c, 0x98, 0x44, 0x59, 0x16, 0x6c, 0xe9, 0x16, 0x7e, 0x6f,
	0x87, 0xb4, 0x07, 0x3d, 0xb3, 0x3e, 0x8f, 0x8b, 0x6f, 0x3e, 0x39, 0x2d, 0xb0, 0xa2, 0xff, 0xbd,
	0x72, 0x17, 0x0f, 0x37, 0xf2, 0x5c, 0xae, 0xfd, 0x2f, 0xbd, 0x4f, 0x1f, 0x5a, 0x58, 0xe0, 0x92,
	0x49, 0x41, 0x51, 0x89, 0x48, 0x80, 0xb4, 0xd3, 0xbb, 0xcb, 0x59, 0x97, 0xfa, 0x1f, 0xff, 0xc2,
	0x79, 0x5f, 0x7d, 0xf7, 0x07, 0xb3, 0x34, 0xfd, 0xe1, 0x9f, 0xcc, 0x9c, 0x65, 0xca, 0x35, 0x75,
	0xe5, 0xba, 0x9a, 0xc7, 0x34, 0xeb, 0x94, 0xbf, 0x3a, 0xcd, 0x22, 0xe5, 0x7a, 0xd1, 0x6b, 0x16,
	0xf3, 0x78, 0xd1, 0x6b, 0x7e, 0xfb, 0xa7, 0xae, 0x91, 0x24, 0x29, 0xd7, 0x49, 0xd2, 0xab, 0x92,
	0x64, 0x1e, 0x27, 0x49, 0xaf, 0xbb, 0x19, 0xda, 0xb0, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x7f, 0x9c, 0x83, 0xb7, 0x03, 0x00, 0x00,
}
