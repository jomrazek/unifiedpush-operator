// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/job_placeholder_field.proto

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

// Possible values for Job placeholder fields.
type JobPlaceholderFieldEnum_JobPlaceholderField int32

const (
	// Not specified.
	JobPlaceholderFieldEnum_UNSPECIFIED JobPlaceholderFieldEnum_JobPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	JobPlaceholderFieldEnum_UNKNOWN JobPlaceholderFieldEnum_JobPlaceholderField = 1
	// Data Type: STRING. Required. If only JOB_ID is specified, then it must be
	// unique. If both JOB_ID and LOCATION_ID are specified, then the
	// pair must be unique.
	// ID) pair must be unique.
	JobPlaceholderFieldEnum_JOB_ID JobPlaceholderFieldEnum_JobPlaceholderField = 2
	// Data Type: STRING. Combination of JOB_ID and LOCATION_ID must be unique
	// per offer.
	JobPlaceholderFieldEnum_LOCATION_ID JobPlaceholderFieldEnum_JobPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with job title to be shown in
	// dynamic ad.
	JobPlaceholderFieldEnum_TITLE JobPlaceholderFieldEnum_JobPlaceholderField = 4
	// Data Type: STRING. Job subtitle to be shown in dynamic ad.
	JobPlaceholderFieldEnum_SUBTITLE JobPlaceholderFieldEnum_JobPlaceholderField = 5
	// Data Type: STRING. Description of job to be shown in dynamic ad.
	JobPlaceholderFieldEnum_DESCRIPTION JobPlaceholderFieldEnum_JobPlaceholderField = 6
	// Data Type: URL. Image to be displayed in the ad. Highly recommended for
	// image ads.
	JobPlaceholderFieldEnum_IMAGE_URL JobPlaceholderFieldEnum_JobPlaceholderField = 7
	// Data Type: STRING. Category of property used to group like items together
	// for recommendation engine.
	JobPlaceholderFieldEnum_CATEGORY JobPlaceholderFieldEnum_JobPlaceholderField = 8
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	JobPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS JobPlaceholderFieldEnum_JobPlaceholderField = 9
	// Data Type: STRING. Complete property address, including postal code.
	JobPlaceholderFieldEnum_ADDRESS JobPlaceholderFieldEnum_JobPlaceholderField = 10
	// Data Type: STRING. Salary or salary range of job to be shown in dynamic
	// ad.
	JobPlaceholderFieldEnum_SALARY JobPlaceholderFieldEnum_JobPlaceholderField = 11
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific job and its location).
	JobPlaceholderFieldEnum_FINAL_URLS JobPlaceholderFieldEnum_JobPlaceholderField = 12
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	JobPlaceholderFieldEnum_FINAL_MOBILE_URLS JobPlaceholderFieldEnum_JobPlaceholderField = 14
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	JobPlaceholderFieldEnum_TRACKING_URL JobPlaceholderFieldEnum_JobPlaceholderField = 15
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	JobPlaceholderFieldEnum_ANDROID_APP_LINK JobPlaceholderFieldEnum_JobPlaceholderField = 16
	// Data Type: STRING_LIST. List of recommended job IDs to show together with
	// this item.
	JobPlaceholderFieldEnum_SIMILAR_JOB_IDS JobPlaceholderFieldEnum_JobPlaceholderField = 17
	// Data Type: STRING. iOS app link.
	JobPlaceholderFieldEnum_IOS_APP_LINK JobPlaceholderFieldEnum_JobPlaceholderField = 18
	// Data Type: INT64. iOS app store ID.
	JobPlaceholderFieldEnum_IOS_APP_STORE_ID JobPlaceholderFieldEnum_JobPlaceholderField = 19
)

var JobPlaceholderFieldEnum_JobPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "JOB_ID",
	3:  "LOCATION_ID",
	4:  "TITLE",
	5:  "SUBTITLE",
	6:  "DESCRIPTION",
	7:  "IMAGE_URL",
	8:  "CATEGORY",
	9:  "CONTEXTUAL_KEYWORDS",
	10: "ADDRESS",
	11: "SALARY",
	12: "FINAL_URLS",
	14: "FINAL_MOBILE_URLS",
	15: "TRACKING_URL",
	16: "ANDROID_APP_LINK",
	17: "SIMILAR_JOB_IDS",
	18: "IOS_APP_LINK",
	19: "IOS_APP_STORE_ID",
}
var JobPlaceholderFieldEnum_JobPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"JOB_ID":              2,
	"LOCATION_ID":         3,
	"TITLE":               4,
	"SUBTITLE":            5,
	"DESCRIPTION":         6,
	"IMAGE_URL":           7,
	"CATEGORY":            8,
	"CONTEXTUAL_KEYWORDS": 9,
	"ADDRESS":             10,
	"SALARY":              11,
	"FINAL_URLS":          12,
	"FINAL_MOBILE_URLS":   14,
	"TRACKING_URL":        15,
	"ANDROID_APP_LINK":    16,
	"SIMILAR_JOB_IDS":     17,
	"IOS_APP_LINK":        18,
	"IOS_APP_STORE_ID":    19,
}

func (x JobPlaceholderFieldEnum_JobPlaceholderField) String() string {
	return proto.EnumName(JobPlaceholderFieldEnum_JobPlaceholderField_name, int32(x))
}
func (JobPlaceholderFieldEnum_JobPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_job_placeholder_field_97a24396124ab4a9, []int{0, 0}
}

// Values for Job placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type JobPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobPlaceholderFieldEnum) Reset()         { *m = JobPlaceholderFieldEnum{} }
func (m *JobPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*JobPlaceholderFieldEnum) ProtoMessage()    {}
func (*JobPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_placeholder_field_97a24396124ab4a9, []int{0}
}
func (m *JobPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *JobPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *JobPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobPlaceholderFieldEnum.Merge(dst, src)
}
func (m *JobPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_JobPlaceholderFieldEnum.Size(m)
}
func (m *JobPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_JobPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_JobPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*JobPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.JobPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.JobPlaceholderFieldEnum_JobPlaceholderField", JobPlaceholderFieldEnum_JobPlaceholderField_name, JobPlaceholderFieldEnum_JobPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/job_placeholder_field.proto", fileDescriptor_job_placeholder_field_97a24396124ab4a9)
}

var fileDescriptor_job_placeholder_field_97a24396124ab4a9 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0xc6, 0xba, 0xd5, 0x2d, 0xab, 0xe7, 0x82, 0x26, 0x10, 0xbb, 0xd8, 0x1e, 0x20,
	0x51, 0xc5, 0x15, 0xe1, 0xca, 0x49, 0xdc, 0xca, 0x6d, 0x1a, 0x47, 0x71, 0xda, 0x51, 0x54, 0x29,
	0x4a, 0x97, 0x10, 0x8a, 0xd2, 0xb8, 0xaa, 0xbb, 0x3d, 0x10, 0x97, 0x3c, 0x0a, 0xf7, 0xbc, 0x04,
	0x17, 0x48, 0xbc, 0x01, 0x72, 0x4c, 0x3b, 0x09, 0xc1, 0x6e, 0xa2, 0xe3, 0x73, 0xbe, 0xff, 0xcf,
	0xb1, 0xcf, 0x01, 0x6f, 0x0b, 0x21, 0x8a, 0x32, 0xb7, 0xd2, 0x4c, 0x5a, 0x3a, 0x54, 0xd1, 0x7d,
	0xdf, 0xca, 0xab, 0xbb, 0xb5, 0xb4, 0x3e, 0x8b, 0x65, 0xb2, 0x29, 0xd3, 0xdb, 0xfc, 0x93, 0x28,
	0xb3, 0x7c, 0x9b, 0x7c, 0x5c, 0xe5, 0x65, 0x66, 0x6e, 0xb6, 0x62, 0x27, 0xd0, 0xa5, 0xe6, 0xcd,
	0x34, 0x93, 0xe6, 0x41, 0x6a, 0xde, 0xf7, 0xcd, 0x5a, 0xfa, 0xea, 0xf5, 0xde, 0x79, 0xb3, 0xb2,
	0xd2, 0xaa, 0x12, 0xbb, 0x74, 0xb7, 0x12, 0x95, 0xd4, 0xe2, 0xeb, 0x5f, 0x06, 0xb8, 0x18, 0x89,
	0x65, 0xf8, 0xe0, 0x3d, 0x50, 0xd6, 0xa4, 0xba, 0x5b, 0x5f, 0x7f, 0x37, 0x40, 0xef, 0x1f, 0x35,
	0xd4, 0x05, 0xed, 0x69, 0xc0, 0x43, 0xe2, 0xd2, 0x01, 0x25, 0x1e, 0x7c, 0x82, 0xda, 0xe0, 0x64,
	0x1a, 0x8c, 0x03, 0x76, 0x13, 0xc0, 0x06, 0x02, 0xa0, 0x39, 0x62, 0x4e, 0x42, 0x3d, 0x68, 0x28,
	0xd2, 0x67, 0x2e, 0x8e, 0x29, 0x0b, 0x54, 0xe2, 0x08, 0xb5, 0xc0, 0x71, 0x4c, 0x63, 0x9f, 0xc0,
	0xa7, 0xa8, 0x03, 0x4e, 0xf9, 0xd4, 0xd1, 0xa7, 0x63, 0x45, 0x7a, 0x84, 0xbb, 0x11, 0x0d, 0x15,
	0x0c, 0x9b, 0xe8, 0x19, 0x68, 0xd1, 0x09, 0x1e, 0x92, 0x64, 0x1a, 0xf9, 0xf0, 0x44, 0xd1, 0x2e,
	0x8e, 0xc9, 0x90, 0x45, 0x73, 0x78, 0x8a, 0x2e, 0x40, 0xcf, 0x65, 0x41, 0x4c, 0xde, 0xc7, 0x53,
	0xec, 0x27, 0x63, 0x32, 0xbf, 0x61, 0x91, 0xc7, 0x61, 0x4b, 0x75, 0x82, 0x3d, 0x2f, 0x22, 0x9c,
	0x43, 0xa0, 0x3a, 0xe1, 0xd8, 0xc7, 0xd1, 0x1c, 0xb6, 0xd1, 0x19, 0x00, 0x03, 0x1a, 0x60, 0x5f,
	0xd9, 0x71, 0xd8, 0x41, 0x2f, 0xc0, 0xb9, 0x3e, 0x4f, 0x98, 0x43, 0x7d, 0xa2, 0xd3, 0x67, 0x08,
	0x82, 0x4e, 0x1c, 0x61, 0x77, 0x4c, 0x83, 0x61, 0xfd, 0xe3, 0x2e, 0x7a, 0x0e, 0x20, 0x0e, 0xbc,
	0x88, 0x51, 0x2f, 0xc1, 0x61, 0x98, 0xf8, 0x34, 0x18, 0x43, 0x88, 0x7a, 0xa0, 0xcb, 0xe9, 0x84,
	0xfa, 0x38, 0x4a, 0xf4, 0x65, 0x39, 0x3c, 0x57, 0x62, 0xca, 0xf8, 0x03, 0x86, 0x94, 0x78, 0x9f,
	0xe1, 0x31, 0x8b, 0x88, 0x7a, 0x84, 0x9e, 0xf3, 0xb3, 0x01, 0xae, 0x6e, 0xc5, 0xda, 0x7c, 0x74,
	0x6e, 0xce, 0xcb, 0x91, 0x58, 0xca, 0xbf, 0xdf, 0x3e, 0x54, 0x43, 0x0b, 0x1b, 0x1f, 0x9c, 0x3f,
	0xda, 0x42, 0x94, 0x69, 0x55, 0x98, 0x62, 0x5b, 0x58, 0x45, 0x5e, 0xd5, 0x23, 0xdd, 0xaf, 0xcf,
	0x66, 0x25, 0xff, 0xb3, 0x4d, 0xef, 0xea, 0xef, 0x17, 0xe3, 0x68, 0x88, 0xf1, 0x57, 0xe3, 0x72,
	0xa8, 0xad, 0x70, 0x26, 0x4d, 0x1d, 0xaa, 0x68, 0xd6, 0x37, 0xd5, 0x0e, 0xc8, 0x6f, 0xfb, 0xfa,
	0x02, 0x67, 0x72, 0x71, 0xa8, 0x2f, 0x66, 0xfd, 0x45, 0x5d, 0xff, 0x61, 0x5c, 0xe9, 0xa4, 0x6d,
	0xe3, 0x4c, 0xda, 0xf6, 0x81, 0xb0, 0xed, 0x59, 0xdf, 0xb6, 0x6b, 0x66, 0xd9, 0xac, 0x1b, 0x7b,
	0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xc4, 0x8b, 0x2f, 0xe5, 0x02, 0x00, 0x00,
}
