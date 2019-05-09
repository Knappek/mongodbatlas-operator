// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// An ad.
type Ad struct {
	// The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of final app URLs that will be used on mobile if the user has the
	// specific app installed.
	FinalAppUrls []*common.FinalAppUrl `protobuf:"bytes,35,rep,name=final_app_urls,json=finalAppUrls,proto3" json:"final_app_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v1.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// Additional URLs for the ad that are tagged with a unique identifier that
	// can be referenced from other fields in the ad.
	UrlCollections []*common.UrlCollection `protobuf:"bytes,26,rep,name=url_collections,json=urlCollections,proto3" json:"url_collections,omitempty"`
	// The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// If this ad is system managed, then this field will indicate the source.
	// This field is read-only.
	SystemManagedResourceSource enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource `protobuf:"varint,27,opt,name=system_managed_resource_source,json=systemManagedResourceSource,proto3,enum=google.ads.googleads.v1.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource" json:"system_managed_resource_source,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	//	*Ad_VideoAd
	//	*Ad_ResponsiveSearchAd
	//	*Ad_LegacyResponsiveDisplayAd
	//	*Ad_AppAd
	//	*Ad_LegacyAppInstallAd
	//	*Ad_ResponsiveDisplayAd
	//	*Ad_DisplayUploadAd
	//	*Ad_AppEngagementAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_55cd72465d865885, []int{0}
}
func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (dst *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(dst, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalAppUrls() []*common.FinalAppUrl {
	if m != nil {
		return m.FinalAppUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetUrlCollections() []*common.UrlCollection {
	if m != nil {
		return m.UrlCollections
	}
	return nil
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Ad) GetSystemManagedResourceSource() enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource {
	if m != nil {
		return m.SystemManagedResourceSource
	}
	return enums.SystemManagedResourceSourceEnum_UNSPECIFIED
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

type Ad_VideoAd struct {
	VideoAd *common.VideoAdInfo `protobuf:"bytes,24,opt,name=video_ad,json=videoAd,proto3,oneof"`
}

type Ad_ResponsiveSearchAd struct {
	ResponsiveSearchAd *common.ResponsiveSearchAdInfo `protobuf:"bytes,25,opt,name=responsive_search_ad,json=responsiveSearchAd,proto3,oneof"`
}

type Ad_LegacyResponsiveDisplayAd struct {
	LegacyResponsiveDisplayAd *common.LegacyResponsiveDisplayAdInfo `protobuf:"bytes,28,opt,name=legacy_responsive_display_ad,json=legacyResponsiveDisplayAd,proto3,oneof"`
}

type Ad_AppAd struct {
	AppAd *common.AppAdInfo `protobuf:"bytes,29,opt,name=app_ad,json=appAd,proto3,oneof"`
}

type Ad_LegacyAppInstallAd struct {
	LegacyAppInstallAd *common.LegacyAppInstallAdInfo `protobuf:"bytes,30,opt,name=legacy_app_install_ad,json=legacyAppInstallAd,proto3,oneof"`
}

type Ad_ResponsiveDisplayAd struct {
	ResponsiveDisplayAd *common.ResponsiveDisplayAdInfo `protobuf:"bytes,31,opt,name=responsive_display_ad,json=responsiveDisplayAd,proto3,oneof"`
}

type Ad_DisplayUploadAd struct {
	DisplayUploadAd *common.DisplayUploadAdInfo `protobuf:"bytes,33,opt,name=display_upload_ad,json=displayUploadAd,proto3,oneof"`
}

type Ad_AppEngagementAd struct {
	AppEngagementAd *common.AppEngagementAdInfo `protobuf:"bytes,34,opt,name=app_engagement_ad,json=appEngagementAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (*Ad_VideoAd) isAd_AdData() {}

func (*Ad_ResponsiveSearchAd) isAd_AdData() {}

func (*Ad_LegacyResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_AppAd) isAd_AdData() {}

func (*Ad_LegacyAppInstallAd) isAd_AdData() {}

func (*Ad_ResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_DisplayUploadAd) isAd_AdData() {}

func (*Ad_AppEngagementAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

func (m *Ad) GetVideoAd() *common.VideoAdInfo {
	if x, ok := m.GetAdData().(*Ad_VideoAd); ok {
		return x.VideoAd
	}
	return nil
}

func (m *Ad) GetResponsiveSearchAd() *common.ResponsiveSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveSearchAd); ok {
		return x.ResponsiveSearchAd
	}
	return nil
}

func (m *Ad) GetLegacyResponsiveDisplayAd() *common.LegacyResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyResponsiveDisplayAd); ok {
		return x.LegacyResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetAppAd() *common.AppAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppAd); ok {
		return x.AppAd
	}
	return nil
}

func (m *Ad) GetLegacyAppInstallAd() *common.LegacyAppInstallAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyAppInstallAd); ok {
		return x.LegacyAppInstallAd
	}
	return nil
}

func (m *Ad) GetResponsiveDisplayAd() *common.ResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveDisplayAd); ok {
		return x.ResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetDisplayUploadAd() *common.DisplayUploadAdInfo {
	if x, ok := m.GetAdData().(*Ad_DisplayUploadAd); ok {
		return x.DisplayUploadAd
	}
	return nil
}

func (m *Ad) GetAppEngagementAd() *common.AppEngagementAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppEngagementAd); ok {
		return x.AppEngagementAd
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Ad) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Ad_OneofMarshaler, _Ad_OneofUnmarshaler, _Ad_OneofSizer, []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
		(*Ad_VideoAd)(nil),
		(*Ad_ResponsiveSearchAd)(nil),
		(*Ad_LegacyResponsiveDisplayAd)(nil),
		(*Ad_AppAd)(nil),
		(*Ad_LegacyAppInstallAd)(nil),
		(*Ad_ResponsiveDisplayAd)(nil),
		(*Ad_DisplayUploadAd)(nil),
		(*Ad_AppEngagementAd)(nil),
	}
}

func _Ad_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Ad)
	// ad_data
	switch x := m.AdData.(type) {
	case *Ad_TextAd:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextAd); err != nil {
			return err
		}
	case *Ad_ExpandedTextAd:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExpandedTextAd); err != nil {
			return err
		}
	case *Ad_CallOnlyAd:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CallOnlyAd); err != nil {
			return err
		}
	case *Ad_ExpandedDynamicSearchAd:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExpandedDynamicSearchAd); err != nil {
			return err
		}
	case *Ad_HotelAd:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelAd); err != nil {
			return err
		}
	case *Ad_ShoppingSmartAd:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShoppingSmartAd); err != nil {
			return err
		}
	case *Ad_ShoppingProductAd:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShoppingProductAd); err != nil {
			return err
		}
	case *Ad_GmailAd:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GmailAd); err != nil {
			return err
		}
	case *Ad_ImageAd:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageAd); err != nil {
			return err
		}
	case *Ad_VideoAd:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VideoAd); err != nil {
			return err
		}
	case *Ad_ResponsiveSearchAd:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponsiveSearchAd); err != nil {
			return err
		}
	case *Ad_LegacyResponsiveDisplayAd:
		b.EncodeVarint(28<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LegacyResponsiveDisplayAd); err != nil {
			return err
		}
	case *Ad_AppAd:
		b.EncodeVarint(29<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppAd); err != nil {
			return err
		}
	case *Ad_LegacyAppInstallAd:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LegacyAppInstallAd); err != nil {
			return err
		}
	case *Ad_ResponsiveDisplayAd:
		b.EncodeVarint(31<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponsiveDisplayAd); err != nil {
			return err
		}
	case *Ad_DisplayUploadAd:
		b.EncodeVarint(33<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DisplayUploadAd); err != nil {
			return err
		}
	case *Ad_AppEngagementAd:
		b.EncodeVarint(34<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppEngagementAd); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Ad.AdData has unexpected type %T", x)
	}
	return nil
}

func _Ad_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Ad)
	switch tag {
	case 6: // ad_data.text_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TextAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_TextAd{msg}
		return true, err
	case 7: // ad_data.expanded_text_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ExpandedTextAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ExpandedTextAd{msg}
		return true, err
	case 13: // ad_data.call_only_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CallOnlyAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_CallOnlyAd{msg}
		return true, err
	case 14: // ad_data.expanded_dynamic_search_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ExpandedDynamicSearchAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ExpandedDynamicSearchAd{msg}
		return true, err
	case 15: // ad_data.hotel_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_HotelAd{msg}
		return true, err
	case 17: // ad_data.shopping_smart_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ShoppingSmartAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ShoppingSmartAd{msg}
		return true, err
	case 18: // ad_data.shopping_product_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ShoppingProductAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ShoppingProductAd{msg}
		return true, err
	case 21: // ad_data.gmail_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.GmailAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_GmailAd{msg}
		return true, err
	case 22: // ad_data.image_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ImageAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ImageAd{msg}
		return true, err
	case 24: // ad_data.video_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.VideoAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_VideoAd{msg}
		return true, err
	case 25: // ad_data.responsive_search_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ResponsiveSearchAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ResponsiveSearchAd{msg}
		return true, err
	case 28: // ad_data.legacy_responsive_display_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.LegacyResponsiveDisplayAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_LegacyResponsiveDisplayAd{msg}
		return true, err
	case 29: // ad_data.app_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.AppAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_AppAd{msg}
		return true, err
	case 30: // ad_data.legacy_app_install_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.LegacyAppInstallAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_LegacyAppInstallAd{msg}
		return true, err
	case 31: // ad_data.responsive_display_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ResponsiveDisplayAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ResponsiveDisplayAd{msg}
		return true, err
	case 33: // ad_data.display_upload_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.DisplayUploadAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_DisplayUploadAd{msg}
		return true, err
	case 34: // ad_data.app_engagement_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.AppEngagementAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_AppEngagementAd{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Ad_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Ad)
	// ad_data
	switch x := m.AdData.(type) {
	case *Ad_TextAd:
		s := proto.Size(x.TextAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ExpandedTextAd:
		s := proto.Size(x.ExpandedTextAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_CallOnlyAd:
		s := proto.Size(x.CallOnlyAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ExpandedDynamicSearchAd:
		s := proto.Size(x.ExpandedDynamicSearchAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_HotelAd:
		s := proto.Size(x.HotelAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ShoppingSmartAd:
		s := proto.Size(x.ShoppingSmartAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ShoppingProductAd:
		s := proto.Size(x.ShoppingProductAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_GmailAd:
		s := proto.Size(x.GmailAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ImageAd:
		s := proto.Size(x.ImageAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_VideoAd:
		s := proto.Size(x.VideoAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ResponsiveSearchAd:
		s := proto.Size(x.ResponsiveSearchAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_LegacyResponsiveDisplayAd:
		s := proto.Size(x.LegacyResponsiveDisplayAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_AppAd:
		s := proto.Size(x.AppAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_LegacyAppInstallAd:
		s := proto.Size(x.LegacyAppInstallAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ResponsiveDisplayAd:
		s := proto.Size(x.ResponsiveDisplayAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_DisplayUploadAd:
		s := proto.Size(x.DisplayUploadAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_AppEngagementAd:
		s := proto.Size(x.AppEngagementAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v1.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad.proto", fileDescriptor_ad_55cd72465d865885)
}

var fileDescriptor_ad_55cd72465d865885 = []byte{
	// 1176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0xdf, 0x6e, 0xdb, 0x36,
	0x1b, 0xc6, 0x3f, 0xbb, 0x6d, 0xf2, 0x95, 0xed, 0xf2, 0x87, 0x6e, 0x56, 0x35, 0xcd, 0xba, 0xb6,
	0x43, 0x81, 0xae, 0xc5, 0xe4, 0x26, 0x5d, 0x3b, 0xc0, 0x45, 0x81, 0xc9, 0x4d, 0xd6, 0x64, 0x58,
	0x31, 0xcf, 0xae, 0x7d, 0x50, 0x64, 0x13, 0x18, 0x91, 0x56, 0xb4, 0x52, 0x24, 0x41, 0x4a, 0x5e,
	0xbd, 0xa3, 0xdd, 0xc9, 0x80, 0x1d, 0xee, 0x52, 0x76, 0x25, 0xc3, 0x8e, 0x77, 0x01, 0x03, 0x49,
	0x89, 0x8e, 0x93, 0x38, 0xd2, 0x89, 0x21, 0x92, 0xef, 0xf3, 0x7b, 0x5e, 0xbe, 0x24, 0x45, 0x19,
	0x3c, 0x8a, 0x39, 0x8f, 0x29, 0x69, 0x23, 0xac, 0xda, 0xf6, 0x51, 0x3f, 0x4d, 0xb6, 0xdb, 0x92,
	0x28, 0x9e, 0xcb, 0x88, 0xa8, 0x36, 0xc2, 0xbe, 0x90, 0x3c, 0xe3, 0xf0, 0x9e, 0x0d, 0xf0, 0x11,
	0x56, 0xbe, 0x8b, 0xf5, 0x27, 0xdb, 0xbe, 0x8b, 0xdd, 0xdc, 0x59, 0x84, 0x8b, 0x78, 0x9a, 0x72,
	0xd6, 0x46, 0x38, 0xcc, 0xa6, 0x82, 0x84, 0x09, 0x1b, 0x73, 0x65, 0xb1, 0x9b, 0xcf, 0x2a, 0x34,
	0x51, 0xae, 0x32, 0x9e, 0x86, 0x02, 0x49, 0x94, 0x92, 0x8c, 0xc8, 0x42, 0x56, 0x65, 0x35, 0x4e,
	0x18, 0xa2, 0x21, 0x12, 0x22, 0xcc, 0x25, 0x2d, 0x34, 0x4f, 0x2b, 0x34, 0xb9, 0xa4, 0x61, 0xc4,
	0x29, 0x25, 0x51, 0x96, 0x70, 0x56, 0x88, 0x1e, 0x2f, 0x12, 0x11, 0x96, 0xa7, 0xaa, 0x9c, 0x52,
	0x11, 0xfc, 0xe8, 0xe2, 0x60, 0x4c, 0x26, 0x49, 0x54, 0xc6, 0x7e, 0x7d, 0x71, 0xac, 0x9a, 0xaa,
	0x8c, 0xa4, 0x61, 0x8a, 0x18, 0x8a, 0x09, 0x0e, 0x09, 0xcb, 0x92, 0x6c, 0x1a, 0xda, 0x4a, 0x17,
	0x84, 0x3b, 0x05, 0xc1, 0xb4, 0x8e, 0xf2, 0x71, 0xfb, 0x17, 0x89, 0x84, 0x20, 0xb2, 0x2c, 0xed,
	0x56, 0xe9, 0x20, 0x92, 0x36, 0x62, 0x8c, 0x67, 0x48, 0xcf, 0xab, 0x18, 0xbd, 0xff, 0xf7, 0x0d,
	0xd0, 0x0c, 0x30, 0x7c, 0x0c, 0x9a, 0x09, 0xf6, 0x1a, 0x77, 0x1b, 0x0f, 0xaf, 0xed, 0xdc, 0x2e,
	0x16, 0xd6, 0x2f, 0x89, 0xfe, 0x01, 0xcb, 0x9e, 0x7f, 0x39, 0x42, 0x34, 0x27, 0xfd, 0x66, 0x82,
	0xe1, 0x0b, 0x00, 0x6c, 0x61, 0x73, 0x49, 0x95, 0xd7, 0xbc, 0x7b, 0xe9, 0xe1, 0xb5, 0x9d, 0xad,
	0x33, 0xa2, 0x41, 0x26, 0x13, 0x16, 0x5b, 0xd5, 0x55, 0x13, 0x3f, 0x94, 0x54, 0xc1, 0x1f, 0xc0,
	0xca, 0xdc, 0xaa, 0x28, 0xef, 0x33, 0x03, 0x78, 0xec, 0x2f, 0xda, 0x59, 0x76, 0x5d, 0xfc, 0x6f,
	0xb4, 0x2a, 0x10, 0x62, 0x28, 0x69, 0xff, 0xfa, 0x78, 0xd6, 0x50, 0x70, 0x1f, 0xac, 0x5b, 0x64,
	0xca, 0x8f, 0x12, 0x4a, 0x2c, 0x75, 0xad, 0x46, 0x5a, 0xab, 0x46, 0xf6, 0xc6, 0xa8, 0x0c, 0xa9,
	0x07, 0x36, 0x32, 0x89, 0xa2, 0xf7, 0x09, 0x8b, 0x35, 0x25, 0xcc, 0x48, 0x2a, 0x28, 0xca, 0x88,
	0x77, 0xdd, 0x54, 0xe6, 0x62, 0x5a, 0xab, 0x94, 0x0e, 0x25, 0x7d, 0x5b, 0x08, 0x61, 0x04, 0x36,
	0xcc, 0x86, 0x3a, 0xb5, 0x7f, 0x95, 0x07, 0x4c, 0x7e, 0xed, 0xaa, 0x59, 0xbf, 0x32, 0xc2, 0x5e,
	0xa9, 0xeb, 0xb7, 0x72, 0x49, 0x4f, 0xf5, 0x29, 0xf8, 0x12, 0x5c, 0xc3, 0x89, 0x12, 0x14, 0x4d,
	0x75, 0xd6, 0xde, 0xe5, 0x1a, 0xc9, 0x82, 0x42, 0x30, 0x94, 0x14, 0xee, 0x82, 0xcb, 0x7a, 0xf7,
	0x7a, 0x57, 0xee, 0x36, 0x1e, 0xae, 0xec, 0x3c, 0x59, 0x98, 0x92, 0xd9, 0x92, 0x7e, 0x80, 0xdf,
	0x4e, 0x05, 0xd9, 0x63, 0x79, 0x5a, 0x3c, 0xf6, 0x8d, 0x1a, 0x1e, 0x80, 0x16, 0xc2, 0x98, 0xe0,
	0xf0, 0x68, 0x1a, 0x5a, 0x59, 0x88, 0xb0, 0xf2, 0x5a, 0x26, 0x99, 0xcd, 0x33, 0xc9, 0x74, 0x39,
	0xa7, 0x36, 0x95, 0x35, 0x23, 0xeb, 0x4e, 0x5f, 0x9b, 0x88, 0x00, 0x2b, 0xf8, 0x23, 0x58, 0xb7,
	0x87, 0x24, 0x14, 0x92, 0x8c, 0x89, 0x24, 0x2c, 0x22, 0xde, 0x8d, 0x5a, 0xd9, 0xed, 0x1a, 0x9d,
	0xc9, 0xce, 0x3e, 0xf6, 0xd7, 0x2c, 0xaa, 0xe7, 0x48, 0x70, 0x04, 0x56, 0xe7, 0x0f, 0xb9, 0xf2,
	0x36, 0xcd, 0x6a, 0x7c, 0x51, 0xb5, 0x1a, 0x43, 0x49, 0x5f, 0x39, 0x55, 0x7f, 0x25, 0x3f, 0xd9,
	0x54, 0xf0, 0x09, 0xb8, 0xcc, 0x50, 0x4a, 0xbc, 0x9b, 0x35, 0xea, 0x6f, 0x22, 0xe1, 0xef, 0x0d,
	0x70, 0xe7, 0xd4, 0x11, 0x2f, 0xdf, 0xa3, 0xc5, 0x21, 0xf7, 0x6e, 0x9b, 0x69, 0xbf, 0xab, 0x98,
	0xf6, 0xc0, 0x40, 0xde, 0x58, 0x46, 0xbf, 0x40, 0x0c, 0xcc, 0xaf, 0xa9, 0xc5, 0x05, 0xe3, 0xfd,
	0xdb, 0x6a, 0xf1, 0x20, 0xdc, 0x03, 0xcb, 0x19, 0xf9, 0x90, 0x85, 0x08, 0x7b, 0x4b, 0x66, 0x5a,
	0x8f, 0xaa, 0x6a, 0xf4, 0x96, 0x7c, 0xc8, 0x02, 0x7c, 0xc0, 0xc6, 0x7c, 0xff, 0x7f, 0xfd, 0xa5,
	0xcc, 0xb4, 0xe0, 0x4f, 0x60, 0x8d, 0x7c, 0x10, 0x88, 0xe9, 0xfd, 0x51, 0xf2, 0x96, 0x0d, 0x6f,
	0xa7, 0x8a, 0xb7, 0x57, 0xe8, 0xe6, 0xb8, 0x2b, 0x64, 0xae, 0x17, 0xf6, 0xc1, 0xf5, 0x08, 0x51,
	0x1a, 0x72, 0x46, 0xa7, 0x9a, 0xfd, 0x91, 0x61, 0xfb, 0x95, 0xa7, 0x0b, 0x51, 0xfa, 0x3d, 0xa3,
	0x53, 0xc7, 0x05, 0x91, 0xeb, 0x81, 0xbf, 0x82, 0x4d, 0x97, 0x33, 0x9e, 0x32, 0x94, 0x26, 0x51,
	0xa8, 0x08, 0x92, 0xd1, 0xb1, 0x76, 0x58, 0x31, 0x0e, 0x2f, 0xea, 0x66, 0xbf, 0x6b, 0x01, 0x03,
	0xa3, 0x77, 0x76, 0x37, 0xc9, 0xf9, 0xc3, 0x70, 0x1f, 0xfc, 0xff, 0x98, 0x67, 0x84, 0x6a, 0xa7,
	0x55, 0xe3, 0x54, 0xf9, 0x7e, 0xdc, 0xd7, 0xf1, 0x8e, 0xbc, 0x7c, 0x6c, 0x9b, 0x10, 0x81, 0x75,
	0x75, 0xcc, 0x85, 0xd0, 0xaf, 0x34, 0x95, 0x22, 0x69, 0x4a, 0xbf, 0x6e, 0x90, 0x4f, 0xab, 0x90,
	0x83, 0x42, 0x38, 0xd0, 0x3a, 0x87, 0x5e, 0x55, 0xf3, 0xdd, 0x30, 0x06, 0x2d, 0x67, 0x21, 0x24,
	0xc7, 0x79, 0x64, 0x4c, 0xa0, 0x31, 0x79, 0x56, 0xd7, 0xa4, 0x67, 0x95, 0xce, 0xc6, 0xa5, 0xed,
	0x06, 0x74, 0x55, 0xe2, 0x14, 0x25, 0xa6, 0x2a, 0x1b, 0xf5, 0xaa, 0xf2, 0x5a, 0xc7, 0xcf, 0xaa,
	0x12, 0xdb, 0xa6, 0x26, 0x25, 0x29, 0x8a, 0xf5, 0x2b, 0xca, 0xfb, 0xb8, 0x1e, 0xe9, 0x40, 0xc7,
	0xcf, 0x48, 0x89, 0x6d, 0x6a, 0xd2, 0x24, 0xc1, 0x84, 0x6b, 0x92, 0x57, 0x8f, 0x34, 0xd2, 0xf1,
	0x33, 0xd2, 0xc4, 0x36, 0xe1, 0xcf, 0xe0, 0x86, 0x24, 0x4a, 0x70, 0xa6, 0x92, 0x09, 0x39, 0xb1,
	0xd3, 0x6e, 0x19, 0xea, 0xf3, 0x2a, 0x6a, 0xdf, 0x69, 0x4f, 0x6d, 0x32, 0x28, 0xcf, 0x8c, 0xc0,
	0xdf, 0x1a, 0x60, 0x8b, 0x92, 0x18, 0x45, 0xd3, 0xf0, 0x84, 0x67, 0x79, 0x89, 0x20, 0xec, 0x6d,
	0x19, 0xd3, 0x97, 0x55, 0xa6, 0xdf, 0x19, 0xc6, 0xcc, 0x7a, 0xd7, 0x12, 0x9c, 0xf7, 0x2d, 0xba,
	0x28, 0x00, 0x76, 0xc1, 0x92, 0xfe, 0x04, 0x40, 0xd8, 0xfb, 0xc4, 0x78, 0x7d, 0x5e, 0xe5, 0x15,
	0x08, 0xe1, 0xb8, 0x57, 0x90, 0x6e, 0xc0, 0xf7, 0x60, 0xa3, 0x98, 0x85, 0x46, 0x25, 0x4c, 0x65,
	0xfa, 0x25, 0x80, 0xb0, 0x77, 0xa7, 0x5e, 0xcd, 0x6c, 0xfa, 0x81, 0x10, 0x07, 0x56, 0x3a, 0xab,
	0x19, 0x3d, 0x33, 0x02, 0x53, 0xb0, 0x71, 0x7e, 0xad, 0x3e, 0x35, 0x66, 0x5f, 0xd5, 0x5f, 0xa0,
	0xd3, 0x55, 0x6a, 0xc9, 0x73, 0xea, 0x83, 0xc0, 0xba, 0xbb, 0xd4, 0x05, 0xe5, 0x08, 0x6b, 0xab,
	0x7b, 0xf5, 0x0e, 0x6e, 0x41, 0x19, 0x1a, 0xdd, 0xec, 0xe0, 0xe2, 0xf9, 0x6e, 0x6d, 0xa1, 0xeb,
	0x46, 0x58, 0x8c, 0x62, 0x92, 0x12, 0x66, 0x8e, 0xed, 0xfd, 0x7a, 0x16, 0x81, 0x10, 0x7b, 0x4e,
	0x37, 0xb3, 0x40, 0xf3, 0xdd, 0xdd, 0xab, 0x60, 0x19, 0xe1, 0x10, 0xa3, 0x0c, 0x75, 0xff, 0x6d,
	0x80, 0x07, 0x11, 0x4f, 0xfd, 0xca, 0x7f, 0x10, 0xdd, 0xe5, 0x00, 0xf7, 0xf4, 0xa5, 0xd9, 0x6b,
	0xbc, 0xfb, 0xb6, 0x88, 0x8e, 0x39, 0x45, 0x2c, 0xf6, 0xb9, 0x8c, 0xdb, 0x31, 0x61, 0xe6, 0x4a,
	0x2d, 0xbf, 0x97, 0x45, 0xa2, 0x2e, 0xf8, 0xeb, 0xf2, 0xc2, 0x3d, 0xfd, 0xd1, 0xbc, 0xf4, 0x3a,
	0x08, 0xfe, 0x6c, 0xde, 0xb3, 0x1f, 0x1a, 0x7e, 0x80, 0x95, 0xef, 0xbe, 0x39, 0xfc, 0xd1, 0xb6,
	0x5f, 0x5e, 0x7d, 0xea, 0xaf, 0x32, 0xe6, 0x30, 0xc0, 0xea, 0xd0, 0xc5, 0x1c, 0x8e, 0xb6, 0x0f,
	0x5d, 0xcc, 0x3f, 0xcd, 0x07, 0x76, 0xa0, 0xd3, 0x09, 0xb0, 0xea, 0x74, 0x5c, 0x54, 0xa7, 0x33,
	0xda, 0xee, 0x74, 0x5c, 0xdc, 0xd1, 0x92, 0x49, 0xf6, 0xe9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x9c, 0xa7, 0x91, 0x66, 0x0d, 0x00, 0x00,
}
