// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1beta1/organization_settings.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1beta1"

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

// The mode of inclusion when running Asset Discovery.
// Asset discovery can be limited by explicitly identifying projects to be
// included or excluded. If INCLUDE_ONLY is set, then only those projects
// within the organization and their children are discovered during asset
// discovery. If EXCLUDE is set, then projects that don't match those
// projects are discovered during asset discovery. If neither are set, then
// all projects within the organization are discovered during asset
// discovery.
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode int32

const (
	// Unspecified. Setting the mode with this value will disable
	// inclusion/exclusion filtering for Asset Discovery.
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 0
	// Asset Discovery will capture only the resources within the projects
	// specified. All other resources will be ignored.
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 1
	// Asset Discovery will ignore all resources under the projects specified.
	// All other resources will be retrieved.
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 2
)

var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name = map[int32]string{
	0: "INCLUSION_MODE_UNSPECIFIED",
	1: "INCLUDE_ONLY",
	2: "EXCLUDE",
}
var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value = map[string]int32{
	"INCLUSION_MODE_UNSPECIFIED": 0,
	"INCLUDE_ONLY":               1,
	"EXCLUDE":                    2,
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) String() string {
	return proto.EnumName(OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, int32(x))
}
func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_8916391f4b80a5cd, []int{0, 0, 0}
}

// User specified settings that are attached to the Cloud Security Command
// Center (Cloud SCC) organization.
type OrganizationSettings struct {
	// The relative resource name of the settings. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/organizationSettings".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	// set to `true`, then discovery of assets will occur. If it is set to `false,
	// all historical assets will remain, but discovery of future assets will not
	// occur.
	EnableAssetDiscovery bool `protobuf:"varint,2,opt,name=enable_asset_discovery,json=enableAssetDiscovery,proto3" json:"enable_asset_discovery,omitempty"`
	// The configuration used for Asset Discovery runs.
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `protobuf:"bytes,3,opt,name=asset_discovery_config,json=assetDiscoveryConfig,proto3" json:"asset_discovery_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *OrganizationSettings) Reset()         { *m = OrganizationSettings{} }
func (m *OrganizationSettings) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings) ProtoMessage()    {}
func (*OrganizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_8916391f4b80a5cd, []int{0}
}
func (m *OrganizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings.Unmarshal(m, b)
}
func (m *OrganizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings.Merge(dst, src)
}
func (m *OrganizationSettings) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings.Size(m)
}
func (m *OrganizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings proto.InternalMessageInfo

func (m *OrganizationSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationSettings) GetEnableAssetDiscovery() bool {
	if m != nil {
		return m.EnableAssetDiscovery
	}
	return false
}

func (m *OrganizationSettings) GetAssetDiscoveryConfig() *OrganizationSettings_AssetDiscoveryConfig {
	if m != nil {
		return m.AssetDiscoveryConfig
	}
	return nil
}

// The configuration used for Asset Discovery runs.
type OrganizationSettings_AssetDiscoveryConfig struct {
	// The project ids to use for filtering asset discovery.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	// The mode to use for filtering asset discovery.
	InclusionMode        OrganizationSettings_AssetDiscoveryConfig_InclusionMode `protobuf:"varint,2,opt,name=inclusion_mode,json=inclusionMode,proto3,enum=google.cloud.securitycenter.v1beta1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode" json:"inclusion_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *OrganizationSettings_AssetDiscoveryConfig) Reset() {
	*m = OrganizationSettings_AssetDiscoveryConfig{}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings_AssetDiscoveryConfig) ProtoMessage()    {}
func (*OrganizationSettings_AssetDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_8916391f4b80a5cd, []int{0, 0}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Unmarshal(m, b)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings_AssetDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Merge(dst, src)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Size(m)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig proto.InternalMessageInfo

func (m *OrganizationSettings_AssetDiscoveryConfig) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *OrganizationSettings_AssetDiscoveryConfig) GetInclusionMode() OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	if m != nil {
		return m.InclusionMode
	}
	return OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OrganizationSettings)(nil), "google.cloud.securitycenter.v1beta1.OrganizationSettings")
	proto.RegisterType((*OrganizationSettings_AssetDiscoveryConfig)(nil), "google.cloud.securitycenter.v1beta1.OrganizationSettings.AssetDiscoveryConfig")
	proto.RegisterEnum("google.cloud.securitycenter.v1beta1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode", OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value)
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1beta1/organization_settings.proto", fileDescriptor_organization_settings_8916391f4b80a5cd)
}

var fileDescriptor_organization_settings_8916391f4b80a5cd = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6b, 0x14, 0x31,
	0x14, 0xc6, 0xcd, 0xae, 0xa8, 0xcd, 0xda, 0xb2, 0x84, 0xa1, 0x0c, 0x8b, 0xe8, 0x50, 0x0f, 0xce,
	0x29, 0x43, 0xab, 0x37, 0x0f, 0xa2, 0x3b, 0x23, 0x0c, 0xb4, 0x33, 0x75, 0x96, 0x82, 0x8a, 0x10,
	0xb2, 0x99, 0x18, 0x22, 0xb3, 0x79, 0xc3, 0x24, 0x5b, 0xa8, 0x07, 0x2f, 0x7a, 0xf5, 0xef, 0xf5,
	0x2a, 0xcd, 0x4e, 0xa1, 0x23, 0x8b, 0xec, 0xa1, 0xb7, 0xe4, 0xfb, 0x5e, 0x7e, 0xdf, 0x7b, 0xe4,
	0xe1, 0x37, 0x0a, 0x40, 0x35, 0x32, 0x11, 0x0d, 0xac, 0xeb, 0xc4, 0x4a, 0xb1, 0xee, 0xb4, 0xbb,
	0x12, 0xd2, 0x38, 0xd9, 0x25, 0x97, 0xc7, 0x4b, 0xe9, 0xf8, 0x71, 0x02, 0x9d, 0xe2, 0x46, 0x7f,
	0xe7, 0x4e, 0x83, 0x61, 0x56, 0x3a, 0xa7, 0x8d, 0xb2, 0xb4, 0xed, 0xc0, 0x01, 0x79, 0xbe, 0x01,
	0x50, 0x0f, 0xa0, 0x43, 0x00, 0xed, 0x01, 0xb3, 0x27, 0x7d, 0x0a, 0x6f, 0x75, 0xc2, 0x8d, 0x01,
	0xe7, 0x51, 0x3d, 0xe2, 0xe8, 0xcf, 0x18, 0x07, 0xe5, 0xad, 0x88, 0x45, 0x9f, 0x40, 0x08, 0xbe,
	0x6f, 0xf8, 0x4a, 0x86, 0x28, 0x42, 0xf1, 0x5e, 0xe5, 0xcf, 0xe4, 0x15, 0x3e, 0x94, 0x86, 0x2f,
	0x1b, 0xc9, 0xb8, 0xb5, 0xd2, 0xb1, 0x5a, 0x5b, 0x01, 0x97, 0xb2, 0xbb, 0x0a, 0x47, 0x11, 0x8a,
	0x1f, 0x55, 0xc1, 0xc6, 0x7d, 0x7b, 0x6d, 0xa6, 0x37, 0x1e, 0xf9, 0x85, 0xf0, 0xe1, 0x3f, 0xf5,
	0x4c, 0x80, 0xf9, 0xaa, 0x55, 0x38, 0x8e, 0x50, 0x3c, 0x39, 0x29, 0xe8, 0x0e, 0x73, 0xd0, 0x6d,
	0x5d, 0xd2, 0x61, 0xd4, 0xdc, 0x53, 0xab, 0x80, 0x6f, 0x51, 0x67, 0xbf, 0x47, 0x38, 0xd8, 0x56,
	0x4e, 0x9e, 0xe1, 0x49, 0xdb, 0xc1, 0x37, 0x29, 0x1c, 0xd3, 0xb5, 0x0d, 0x51, 0x34, 0x8e, 0xf7,
	0x2a, 0xdc, 0x4b, 0x79, 0x6d, 0xc9, 0x4f, 0x84, 0x0f, 0xb4, 0x11, 0xcd, 0xda, 0x5e, 0xff, 0xc1,
	0x0a, 0x6a, 0xe9, 0xe7, 0x3d, 0x38, 0xf9, 0x72, 0xb7, 0x8d, 0xd3, 0xfc, 0x26, 0xe4, 0x0c, 0x6a,
	0x59, 0xed, 0xeb, 0xdb, 0xd7, 0xa3, 0x02, 0xef, 0x0f, 0x7c, 0xf2, 0x14, 0xcf, 0xf2, 0x62, 0x7e,
	0x7a, 0xb1, 0xc8, 0xcb, 0x82, 0x9d, 0x95, 0x69, 0xc6, 0x2e, 0x8a, 0xc5, 0x79, 0x36, 0xcf, 0xdf,
	0xe7, 0x59, 0x3a, 0xbd, 0x47, 0xa6, 0xf8, 0xb1, 0xf7, 0xd3, 0x8c, 0x95, 0xc5, 0xe9, 0xa7, 0x29,
	0x22, 0x13, 0xfc, 0x30, 0xfb, 0xe8, 0x95, 0xe9, 0xe8, 0xdd, 0x0f, 0xfc, 0x42, 0xc0, 0x6a, 0x97,
	0x09, 0xce, 0xd1, 0xe7, 0x0f, 0x7d, 0x99, 0x82, 0x86, 0x1b, 0x45, 0xa1, 0x53, 0x89, 0x92, 0xc6,
	0xaf, 0x50, 0xb2, 0xb1, 0x78, 0xab, 0xed, 0x7f, 0x37, 0xf9, 0xf5, 0x50, 0x5e, 0x3e, 0xf0, 0xaf,
	0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x07, 0x96, 0xfd, 0x8f, 0x06, 0x03, 0x00, 0x00,
}
