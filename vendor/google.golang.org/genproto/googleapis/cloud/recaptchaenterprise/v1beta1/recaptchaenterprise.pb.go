// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto

package recaptchaenterprise // import "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum that reprensents the types of annotations.
type AnnotateAssessmentRequest_Annotation int32

const (
	// Default unspecified type.
	AnnotateAssessmentRequest_ANNOTATION_UNSPECIFIED AnnotateAssessmentRequest_Annotation = 0
	// Provides information that the event turned out to be legitimate.
	AnnotateAssessmentRequest_LEGITIMATE AnnotateAssessmentRequest_Annotation = 1
	// Provides information that the event turned out to be fraudulent.
	AnnotateAssessmentRequest_FRAUDULENT AnnotateAssessmentRequest_Annotation = 2
)

var AnnotateAssessmentRequest_Annotation_name = map[int32]string{
	0: "ANNOTATION_UNSPECIFIED",
	1: "LEGITIMATE",
	2: "FRAUDULENT",
}
var AnnotateAssessmentRequest_Annotation_value = map[string]int32{
	"ANNOTATION_UNSPECIFIED": 0,
	"LEGITIMATE":             1,
	"FRAUDULENT":             2,
}

func (x AnnotateAssessmentRequest_Annotation) String() string {
	return proto.EnumName(AnnotateAssessmentRequest_Annotation_name, int32(x))
}
func (AnnotateAssessmentRequest_Annotation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{1, 0}
}

// LINT.IfChange(classification_reason)
// Reasons contributing to the risk analysis verdict.
type Assessment_ClassificationReason int32

const (
	// Default unspecified type.
	Assessment_CLASSIFICATION_REASON_UNSPECIFIED Assessment_ClassificationReason = 0
	// The event appeared to be automated.
	Assessment_AUTOMATION Assessment_ClassificationReason = 1
	// The event was not made from the proper context on the real site.
	Assessment_UNEXPECTED_ENVIRONMENT Assessment_ClassificationReason = 2
	// Browsing behavior leading up to the event was generated was out of the
	// ordinary.
	Assessment_UNEXPECTED_USAGE_PATTERNS Assessment_ClassificationReason = 4
	// Too little traffic has been received from this site thus far to generate
	// quality risk analysis.
	Assessment_PROVISIONAL_RISK_ANALYSIS Assessment_ClassificationReason = 5
)

var Assessment_ClassificationReason_name = map[int32]string{
	0: "CLASSIFICATION_REASON_UNSPECIFIED",
	1: "AUTOMATION",
	2: "UNEXPECTED_ENVIRONMENT",
	4: "UNEXPECTED_USAGE_PATTERNS",
	5: "PROVISIONAL_RISK_ANALYSIS",
}
var Assessment_ClassificationReason_value = map[string]int32{
	"CLASSIFICATION_REASON_UNSPECIFIED": 0,
	"AUTOMATION":                        1,
	"UNEXPECTED_ENVIRONMENT":            2,
	"UNEXPECTED_USAGE_PATTERNS":         4,
	"PROVISIONAL_RISK_ANALYSIS":         5,
}

func (x Assessment_ClassificationReason) String() string {
	return proto.EnumName(Assessment_ClassificationReason_name, int32(x))
}
func (Assessment_ClassificationReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{3, 0}
}

// Enum that represents the types of invalid token reasons.
type TokenProperties_InvalidReason int32

const (
	// Default unspecified type.
	TokenProperties_INVALID_REASON_UNSPECIFIED TokenProperties_InvalidReason = 0
	// If the failure reason was not accounted for.
	TokenProperties_UNKNOWN_INVALID_REASON TokenProperties_InvalidReason = 1
	// The provided user verification token was malformed.
	TokenProperties_MALFORMED TokenProperties_InvalidReason = 2
	// The user verification token had expired.
	TokenProperties_EXPIRED TokenProperties_InvalidReason = 3
	// The user verification had already been seen.
	TokenProperties_DUPE TokenProperties_InvalidReason = 4
	// The user verification token did not match the provided site secret.
	// This may be a configuration error (e.g. development keys used in
	// production) or end users trying to use verification tokens from other
	// sites.
	TokenProperties_SITE_MISMATCH TokenProperties_InvalidReason = 5
	// The user verification token was not present.  It is a required input.
	TokenProperties_MISSING TokenProperties_InvalidReason = 6
)

var TokenProperties_InvalidReason_name = map[int32]string{
	0: "INVALID_REASON_UNSPECIFIED",
	1: "UNKNOWN_INVALID_REASON",
	2: "MALFORMED",
	3: "EXPIRED",
	4: "DUPE",
	5: "SITE_MISMATCH",
	6: "MISSING",
}
var TokenProperties_InvalidReason_value = map[string]int32{
	"INVALID_REASON_UNSPECIFIED": 0,
	"UNKNOWN_INVALID_REASON":     1,
	"MALFORMED":                  2,
	"EXPIRED":                    3,
	"DUPE":                       4,
	"SITE_MISMATCH":              5,
	"MISSING":                    6,
}

func (x TokenProperties_InvalidReason) String() string {
	return proto.EnumName(TokenProperties_InvalidReason_name, int32(x))
}
func (TokenProperties_InvalidReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{5, 0}
}

// The create assessment request message.
type CreateAssessmentRequest struct {
	// Required. The name of the project in which the assessment will be created,
	// in the format "projects/{project_number}".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The asessment details.
	Assessment           *Assessment `protobuf:"bytes,2,opt,name=assessment,proto3" json:"assessment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateAssessmentRequest) Reset()         { *m = CreateAssessmentRequest{} }
func (m *CreateAssessmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAssessmentRequest) ProtoMessage()    {}
func (*CreateAssessmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{0}
}
func (m *CreateAssessmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAssessmentRequest.Unmarshal(m, b)
}
func (m *CreateAssessmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAssessmentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAssessmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAssessmentRequest.Merge(dst, src)
}
func (m *CreateAssessmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAssessmentRequest.Size(m)
}
func (m *CreateAssessmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAssessmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAssessmentRequest proto.InternalMessageInfo

func (m *CreateAssessmentRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateAssessmentRequest) GetAssessment() *Assessment {
	if m != nil {
		return m.Assessment
	}
	return nil
}

// The request message to annotate an Assessment.
type AnnotateAssessmentRequest struct {
	// Required. The resource name of the Assessment, in the format
	// "projects/{project_number}/assessments/{assessment_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The annotation that will be assigned to the Event.
	Annotation           AnnotateAssessmentRequest_Annotation `protobuf:"varint,2,opt,name=annotation,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest_Annotation" json:"annotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AnnotateAssessmentRequest) Reset()         { *m = AnnotateAssessmentRequest{} }
func (m *AnnotateAssessmentRequest) String() string { return proto.CompactTextString(m) }
func (*AnnotateAssessmentRequest) ProtoMessage()    {}
func (*AnnotateAssessmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{1}
}
func (m *AnnotateAssessmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotateAssessmentRequest.Unmarshal(m, b)
}
func (m *AnnotateAssessmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotateAssessmentRequest.Marshal(b, m, deterministic)
}
func (dst *AnnotateAssessmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotateAssessmentRequest.Merge(dst, src)
}
func (m *AnnotateAssessmentRequest) XXX_Size() int {
	return xxx_messageInfo_AnnotateAssessmentRequest.Size(m)
}
func (m *AnnotateAssessmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotateAssessmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotateAssessmentRequest proto.InternalMessageInfo

func (m *AnnotateAssessmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnnotateAssessmentRequest) GetAnnotation() AnnotateAssessmentRequest_Annotation {
	if m != nil {
		return m.Annotation
	}
	return AnnotateAssessmentRequest_ANNOTATION_UNSPECIFIED
}

// Empty response for AnnotateAssessment.
type AnnotateAssessmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotateAssessmentResponse) Reset()         { *m = AnnotateAssessmentResponse{} }
func (m *AnnotateAssessmentResponse) String() string { return proto.CompactTextString(m) }
func (*AnnotateAssessmentResponse) ProtoMessage()    {}
func (*AnnotateAssessmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{2}
}
func (m *AnnotateAssessmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotateAssessmentResponse.Unmarshal(m, b)
}
func (m *AnnotateAssessmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotateAssessmentResponse.Marshal(b, m, deterministic)
}
func (dst *AnnotateAssessmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotateAssessmentResponse.Merge(dst, src)
}
func (m *AnnotateAssessmentResponse) XXX_Size() int {
	return xxx_messageInfo_AnnotateAssessmentResponse.Size(m)
}
func (m *AnnotateAssessmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotateAssessmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotateAssessmentResponse proto.InternalMessageInfo

// A recaptcha assessment resource.
type Assessment struct {
	// Output only. The resource name for the Assessment in the format
	// "projects/{project_number}/assessments/{assessment_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The event being assessed.
	Event *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	// Output only. Legitimate event score from 0.0 to 1.0.
	// (1.0 means very likely legitimate traffic while 0.0 means very likely
	// non-legitimate traffic).
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	// Output only. Properties of the provided event token.
	TokenProperties *TokenProperties `protobuf:"bytes,4,opt,name=token_properties,json=tokenProperties,proto3" json:"token_properties,omitempty"`
	// Output only. Reasons contributing to the risk analysis verdict.
	Reasons              []Assessment_ClassificationReason `protobuf:"varint,5,rep,packed,name=reasons,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.Assessment_ClassificationReason" json:"reasons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Assessment) Reset()         { *m = Assessment{} }
func (m *Assessment) String() string { return proto.CompactTextString(m) }
func (*Assessment) ProtoMessage()    {}
func (*Assessment) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{3}
}
func (m *Assessment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assessment.Unmarshal(m, b)
}
func (m *Assessment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assessment.Marshal(b, m, deterministic)
}
func (dst *Assessment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assessment.Merge(dst, src)
}
func (m *Assessment) XXX_Size() int {
	return xxx_messageInfo_Assessment.Size(m)
}
func (m *Assessment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assessment.DiscardUnknown(m)
}

var xxx_messageInfo_Assessment proto.InternalMessageInfo

func (m *Assessment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assessment) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Assessment) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Assessment) GetTokenProperties() *TokenProperties {
	if m != nil {
		return m.TokenProperties
	}
	return nil
}

func (m *Assessment) GetReasons() []Assessment_ClassificationReason {
	if m != nil {
		return m.Reasons
	}
	return nil
}

type Event struct {
	// The user response token provided by the reCAPTCHA client-side integration
	// on your site.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The site key that was used to invoke reCAPTCHA on your site and generate
	// the token.
	SiteKey              string   `protobuf:"bytes,2,opt,name=site_key,json=siteKey,proto3" json:"site_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{4}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Event) GetSiteKey() string {
	if m != nil {
		return m.SiteKey
	}
	return ""
}

type TokenProperties struct {
	// Output only. Whether the provided user response token is valid.
	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	// Output only. Reason associated with the response when valid = false.
	InvalidReason TokenProperties_InvalidReason `protobuf:"varint,2,opt,name=invalid_reason,json=invalidReason,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.TokenProperties_InvalidReason" json:"invalid_reason,omitempty"`
	// Output only. The timestamp corresponding to the generation of the token.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The hostname of the page on which the token was generated.
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Output only. Action name provided at token generation.
	Action               string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenProperties) Reset()         { *m = TokenProperties{} }
func (m *TokenProperties) String() string { return proto.CompactTextString(m) }
func (*TokenProperties) ProtoMessage()    {}
func (*TokenProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_recaptchaenterprise_7acd4e5481c3d535, []int{5}
}
func (m *TokenProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenProperties.Unmarshal(m, b)
}
func (m *TokenProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenProperties.Marshal(b, m, deterministic)
}
func (dst *TokenProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenProperties.Merge(dst, src)
}
func (m *TokenProperties) XXX_Size() int {
	return xxx_messageInfo_TokenProperties.Size(m)
}
func (m *TokenProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenProperties.DiscardUnknown(m)
}

var xxx_messageInfo_TokenProperties proto.InternalMessageInfo

func (m *TokenProperties) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *TokenProperties) GetInvalidReason() TokenProperties_InvalidReason {
	if m != nil {
		return m.InvalidReason
	}
	return TokenProperties_INVALID_REASON_UNSPECIFIED
}

func (m *TokenProperties) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *TokenProperties) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *TokenProperties) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAssessmentRequest)(nil), "google.cloud.recaptchaenterprise.v1beta1.CreateAssessmentRequest")
	proto.RegisterType((*AnnotateAssessmentRequest)(nil), "google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest")
	proto.RegisterType((*AnnotateAssessmentResponse)(nil), "google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentResponse")
	proto.RegisterType((*Assessment)(nil), "google.cloud.recaptchaenterprise.v1beta1.Assessment")
	proto.RegisterType((*Event)(nil), "google.cloud.recaptchaenterprise.v1beta1.Event")
	proto.RegisterType((*TokenProperties)(nil), "google.cloud.recaptchaenterprise.v1beta1.TokenProperties")
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest_Annotation", AnnotateAssessmentRequest_Annotation_name, AnnotateAssessmentRequest_Annotation_value)
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.Assessment_ClassificationReason", Assessment_ClassificationReason_name, Assessment_ClassificationReason_value)
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.TokenProperties_InvalidReason", TokenProperties_InvalidReason_name, TokenProperties_InvalidReason_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecaptchaEnterpriseServiceV1Beta1Client is the client API for RecaptchaEnterpriseServiceV1Beta1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecaptchaEnterpriseServiceV1Beta1Client interface {
	// Creates an Assessment of the likelihood an event is legitimate.
	CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*Assessment, error)
	// Annotates a previously created Assessment to provide additional information
	// on whether the event turned out to be authentic or fradulent.
	AnnotateAssessment(ctx context.Context, in *AnnotateAssessmentRequest, opts ...grpc.CallOption) (*AnnotateAssessmentResponse, error)
}

type recaptchaEnterpriseServiceV1Beta1Client struct {
	cc *grpc.ClientConn
}

func NewRecaptchaEnterpriseServiceV1Beta1Client(cc *grpc.ClientConn) RecaptchaEnterpriseServiceV1Beta1Client {
	return &recaptchaEnterpriseServiceV1Beta1Client{cc}
}

func (c *recaptchaEnterpriseServiceV1Beta1Client) CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*Assessment, error) {
	out := new(Assessment)
	err := c.cc.Invoke(ctx, "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recaptchaEnterpriseServiceV1Beta1Client) AnnotateAssessment(ctx context.Context, in *AnnotateAssessmentRequest, opts ...grpc.CallOption) (*AnnotateAssessmentResponse, error) {
	out := new(AnnotateAssessmentResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecaptchaEnterpriseServiceV1Beta1Server is the server API for RecaptchaEnterpriseServiceV1Beta1 service.
type RecaptchaEnterpriseServiceV1Beta1Server interface {
	// Creates an Assessment of the likelihood an event is legitimate.
	CreateAssessment(context.Context, *CreateAssessmentRequest) (*Assessment, error)
	// Annotates a previously created Assessment to provide additional information
	// on whether the event turned out to be authentic or fradulent.
	AnnotateAssessment(context.Context, *AnnotateAssessmentRequest) (*AnnotateAssessmentResponse, error)
}

func RegisterRecaptchaEnterpriseServiceV1Beta1Server(s *grpc.Server, srv RecaptchaEnterpriseServiceV1Beta1Server) {
	s.RegisterService(&_RecaptchaEnterpriseServiceV1Beta1_serviceDesc, srv)
}

func _RecaptchaEnterpriseServiceV1Beta1_CreateAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).CreateAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).CreateAssessment(ctx, req.(*CreateAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecaptchaEnterpriseServiceV1Beta1_AnnotateAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnotateAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).AnnotateAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).AnnotateAssessment(ctx, req.(*AnnotateAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecaptchaEnterpriseServiceV1Beta1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1",
	HandlerType: (*RecaptchaEnterpriseServiceV1Beta1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAssessment",
			Handler:    _RecaptchaEnterpriseServiceV1Beta1_CreateAssessment_Handler,
		},
		{
			MethodName: "AnnotateAssessment",
			Handler:    _RecaptchaEnterpriseServiceV1Beta1_AnnotateAssessment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto",
}

func init() {
	proto.RegisterFile("google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto", fileDescriptor_recaptchaenterprise_7acd4e5481c3d535)
}

var fileDescriptor_recaptchaenterprise_7acd4e5481c3d535 = []byte{
	// 945 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0xf9, 0xe8, 0xc7, 0xab, 0xda, 0x35, 0xa3, 0xd5, 0x92, 0x46, 0x0b, 0x74, 0x2d, 0x81,
	0xa2, 0x1e, 0x6c, 0xb5, 0x20, 0x04, 0x5d, 0x38, 0xb8, 0xce, 0x34, 0x3b, 0x6a, 0xe2, 0x44, 0x63,
	0xa7, 0x2c, 0x50, 0xc9, 0x72, 0xdd, 0xd9, 0xac, 0xd9, 0xc4, 0x36, 0x1e, 0xb7, 0xd2, 0x0a, 0xed,
	0x85, 0x0b, 0x67, 0xc4, 0x95, 0x13, 0x17, 0x24, 0xfe, 0x13, 0xb8, 0x70, 0xe0, 0xcc, 0x8d, 0x0b,
	0x57, 0xfe, 0x00, 0x84, 0x3c, 0x63, 0x27, 0xe9, 0xe2, 0x85, 0x52, 0xb8, 0xe5, 0xcd, 0x7b, 0xf3,
	0x7b, 0x1f, 0xbf, 0x37, 0x3f, 0x07, 0x0e, 0x27, 0x71, 0x3c, 0x99, 0x32, 0x23, 0x98, 0xc6, 0x17,
	0xe7, 0x46, 0xca, 0x02, 0x3f, 0xc9, 0x82, 0xc7, 0x3e, 0x8b, 0x32, 0x96, 0x26, 0x69, 0xc8, 0x99,
	0x71, 0xb9, 0x77, 0xc6, 0x32, 0x7f, 0xaf, 0xca, 0xa7, 0x27, 0x69, 0x9c, 0xc5, 0xa8, 0x23, 0x31,
	0x74, 0x81, 0xa1, 0x57, 0xc5, 0x15, 0x18, 0xed, 0xbb, 0x45, 0x36, 0x3f, 0x09, 0x0d, 0x3f, 0x8a,
	0xe2, 0xcc, 0xcf, 0xc2, 0x38, 0xe2, 0x12, 0xa7, 0xfd, 0x7a, 0xe1, 0x15, 0xd6, 0xd9, 0xc5, 0x23,
	0x23, 0x0b, 0x67, 0x8c, 0x67, 0xfe, 0x2c, 0x91, 0x01, 0xda, 0x97, 0x0a, 0xbc, 0x62, 0xa5, 0xcc,
	0xcf, 0x98, 0xc9, 0x39, 0xe3, 0x7c, 0xc6, 0xa2, 0x8c, 0xb2, 0xcf, 0x2e, 0x18, 0xcf, 0xd0, 0x1d,
	0x58, 0x49, 0xfc, 0x94, 0x45, 0x59, 0x4b, 0xd9, 0x51, 0x3a, 0xeb, 0xb4, 0xb0, 0x90, 0x0b, 0xe0,
	0xcf, 0x83, 0x5b, 0xb5, 0x1d, 0xa5, 0xb3, 0xb1, 0xff, 0xb6, 0x7e, 0xdd, 0x8a, 0xf5, 0xa5, 0x44,
	0x4b, 0x38, 0xda, 0x6f, 0x0a, 0x6c, 0x9b, 0xb2, 0x81, 0x8a, 0x5a, 0x10, 0x34, 0x22, 0x7f, 0xc6,
	0x8a, 0x4a, 0xc4, 0x6f, 0x14, 0x01, 0x2c, 0x3a, 0x16, 0x75, 0x6c, 0xed, 0xdb, 0xff, 0xa2, 0x8e,
	0x17, 0x25, 0x2b, 0x3d, 0x61, 0x1c, 0xd1, 0xa5, 0x0c, 0xda, 0x03, 0x80, 0x85, 0x07, 0xb5, 0xe1,
	0x8e, 0x69, 0xdb, 0x43, 0xd7, 0x74, 0xc9, 0xd0, 0xf6, 0xc6, 0xb6, 0x33, 0xc2, 0x16, 0x39, 0x22,
	0xb8, 0xab, 0xbe, 0x84, 0xb6, 0x00, 0xfa, 0xb8, 0x47, 0x5c, 0x32, 0x30, 0x5d, 0xac, 0x2a, 0xb9,
	0x7d, 0x44, 0xcd, 0x71, 0x77, 0xdc, 0xc7, 0xb6, 0xab, 0xd6, 0xb4, 0xbb, 0xd0, 0xae, 0xca, 0xce,
	0x93, 0x38, 0xe2, 0x4c, 0xfb, 0xbd, 0x0e, 0xb0, 0x38, 0xae, 0x6c, 0x1d, 0x43, 0x93, 0x5d, 0x2e,
	0xa6, 0x6f, 0x5c, 0xbf, 0x6b, 0x9c, 0x5f, 0xa3, 0xf2, 0x36, 0xba, 0x0d, 0x4d, 0x1e, 0xc4, 0x29,
	0x6b, 0xd5, 0x77, 0x94, 0x4e, 0x8d, 0x4a, 0x03, 0x9d, 0x83, 0x9a, 0xc5, 0x4f, 0x58, 0xe4, 0x25,
	0x69, 0x9c, 0xb0, 0x34, 0x0b, 0x19, 0x6f, 0x35, 0x44, 0x9e, 0xf7, 0xae, 0x9f, 0xc7, 0xcd, 0x11,
	0x46, 0x73, 0x00, 0x7a, 0x2b, 0xbb, 0x7a, 0x80, 0x02, 0x58, 0x4d, 0x99, 0xcf, 0xe3, 0x88, 0xb7,
	0x9a, 0x3b, 0xf5, 0xce, 0xd6, 0x3e, 0xb9, 0xc9, 0x0a, 0xe9, 0xd6, 0xd4, 0xe7, 0x3c, 0x7c, 0x14,
	0x06, 0x92, 0x2f, 0x81, 0x48, 0x4b, 0x64, 0xed, 0x3b, 0x05, 0x6e, 0x57, 0x45, 0xa0, 0x37, 0xe0,
	0x9e, 0xd5, 0x37, 0x1d, 0x87, 0x1c, 0x11, 0x4b, 0x32, 0x48, 0xb1, 0xe9, 0x54, 0x11, 0x69, 0x8e,
	0xdd, 0xe1, 0x40, 0x84, 0xa8, 0x4a, 0x4e, 0xfa, 0xd8, 0xc6, 0x0f, 0x47, 0xd8, 0x72, 0x71, 0xd7,
	0xc3, 0xf6, 0x09, 0xa1, 0x43, 0x7b, 0x20, 0x48, 0x45, 0xaf, 0xc2, 0xf6, 0x92, 0x6f, 0xec, 0x98,
	0x3d, 0xec, 0x8d, 0x4c, 0xd7, 0xc5, 0xd4, 0x76, 0xd4, 0x46, 0xee, 0x1e, 0xd1, 0xe1, 0x09, 0x71,
	0xc8, 0xd0, 0x36, 0xfb, 0x1e, 0x25, 0xce, 0xb1, 0x67, 0xda, 0x66, 0xff, 0x23, 0x87, 0x38, 0x6a,
	0x53, 0x7b, 0x17, 0x9a, 0xb8, 0xe4, 0x44, 0x8c, 0xaa, 0xe0, 0x5b, 0x1a, 0x68, 0x1b, 0xd6, 0x78,
	0x98, 0x31, 0xef, 0x09, 0x7b, 0x2a, 0x38, 0x5f, 0xa7, 0xab, 0xb9, 0x7d, 0xcc, 0x9e, 0x6a, 0xdf,
	0xd4, 0xe1, 0xd6, 0x73, 0xd3, 0xce, 0x41, 0x2e, 0xfd, 0x69, 0x78, 0x2e, 0x40, 0xd6, 0xa8, 0x34,
	0x50, 0x04, 0x5b, 0x61, 0x24, 0x7e, 0x7a, 0x72, 0x40, 0xc5, 0xa3, 0xe9, 0xdd, 0x98, 0x56, 0x9d,
	0x48, 0xbc, 0x62, 0xee, 0x9b, 0xe1, 0xb2, 0x89, 0xee, 0xc3, 0x46, 0x20, 0xb4, 0xc5, 0xcb, 0x65,
	0x47, 0x2c, 0xd9, 0xc6, 0x7e, 0xbb, 0x4c, 0x56, 0x6a, 0x92, 0xee, 0x96, 0x9a, 0x44, 0x41, 0x86,
	0xe7, 0x07, 0xa8, 0x0d, 0x6b, 0x8f, 0x63, 0x9e, 0x89, 0xd5, 0x6f, 0x88, 0x8e, 0xe7, 0x76, 0xae,
	0x4c, 0x7e, 0x20, 0x5e, 0x7d, 0x53, 0x2a, 0x93, 0xb4, 0xb4, 0xaf, 0x14, 0xd8, 0xbc, 0x52, 0x11,
	0x7a, 0x0d, 0xda, 0xc4, 0x3e, 0x31, 0xfb, 0xa4, 0x5b, 0x4d, 0xb0, 0x20, 0xf4, 0xd8, 0x1e, 0x7e,
	0x68, 0x7b, 0x57, 0xe3, 0x54, 0x05, 0x6d, 0xc2, 0xfa, 0xc0, 0xec, 0x1f, 0x0d, 0xe9, 0x00, 0x77,
	0xd5, 0x1a, 0xda, 0x80, 0x55, 0xfc, 0x70, 0x44, 0x28, 0xee, 0xaa, 0x75, 0xb4, 0x06, 0x8d, 0xee,
	0x78, 0x84, 0xd5, 0x06, 0x7a, 0x19, 0x36, 0x1d, 0xe2, 0x62, 0x6f, 0x40, 0x9c, 0x81, 0xe9, 0x5a,
	0x0f, 0xd4, 0x66, 0x1e, 0x39, 0x20, 0x8e, 0x43, 0xec, 0x9e, 0xba, 0xb2, 0xff, 0x53, 0x1d, 0xee,
	0xd1, 0x72, 0xa2, 0x78, 0x3e, 0x51, 0x87, 0xa5, 0x97, 0x61, 0xc0, 0x4e, 0xf6, 0x0e, 0xf3, 0xb9,
	0xa2, 0x1f, 0x14, 0x50, 0x9f, 0xd7, 0x61, 0x64, 0x5e, 0x9f, 0x97, 0x17, 0x68, 0x78, 0xfb, 0x46,
	0xba, 0xac, 0xbd, 0xff, 0xc5, 0xcf, 0xbf, 0x7e, 0x5d, 0x7b, 0x47, 0xeb, 0xcc, 0x3f, 0x55, 0x9f,
	0x4b, 0xed, 0xff, 0x20, 0x49, 0xe3, 0x4f, 0x59, 0x90, 0x71, 0x63, 0xf7, 0x99, 0xb1, 0x10, 0x6f,
	0x7e, 0xb0, 0xa4, 0xe4, 0xe8, 0x17, 0x05, 0xd0, 0x5f, 0xe5, 0x0d, 0x59, 0xff, 0x83, 0x34, 0xb7,
	0xbb, 0xff, 0x0d, 0xa4, 0x50, 0xd8, 0xb2, 0xbf, 0xbd, 0x45, 0x7f, 0xf9, 0x5e, 0x2d, 0x75, 0xb7,
	0xdc, 0x9c, 0xb1, 0xfb, 0xec, 0xa0, 0xf8, 0x0c, 0xb0, 0x03, 0x65, 0xf7, 0xf0, 0x0f, 0x05, 0xde,
	0x0c, 0xe2, 0x59, 0x59, 0xc9, 0xdf, 0xd4, 0x70, 0xd8, 0xaa, 0x60, 0x7e, 0x94, 0xef, 0xfd, 0x48,
	0xf9, 0xf8, 0x93, 0xe2, 0xfe, 0x24, 0x9e, 0xfa, 0xd1, 0x44, 0x8f, 0xd3, 0x89, 0x31, 0x61, 0x91,
	0x78, 0x15, 0x86, 0x74, 0xf9, 0x49, 0xc8, 0xff, 0xf9, 0x6f, 0xc4, 0xfd, 0x0a, 0xdf, 0xb7, 0xb5,
	0x46, 0xcf, 0xa2, 0xf8, 0xfb, 0x5a, 0xa7, 0x27, 0x93, 0x58, 0x62, 0x5c, 0x15, 0xd5, 0xe8, 0xc5,
	0x06, 0xfe, 0x58, 0x86, 0x9e, 0x8a, 0xd0, 0xd3, 0x8a, 0xd0, 0xd3, 0x13, 0x99, 0xf0, 0x6c, 0x45,
	0x14, 0xf9, 0xd6, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x13, 0x3b, 0x55, 0xea, 0x08, 0x00,
	0x00,
}
