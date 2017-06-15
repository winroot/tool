// Code generated by protoc-gen-go.
// source: matrix.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecognizeType int32

const (
	RecognizeType_REC_TYPE_DEFAULT RecognizeType = 0
	RecognizeType_REC_TYPE_VEHICLE RecognizeType = 1
	RecognizeType_REC_TYPE_FACE    RecognizeType = 2
	RecognizeType_REC_TYPE_ALL     RecognizeType = 3
)

var RecognizeType_name = map[int32]string{
	0: "REC_TYPE_DEFAULT",
	1: "REC_TYPE_VEHICLE",
	2: "REC_TYPE_FACE",
	3: "REC_TYPE_ALL",
}
var RecognizeType_value = map[string]int32{
	"REC_TYPE_DEFAULT": 0,
	"REC_TYPE_VEHICLE": 1,
	"REC_TYPE_FACE":    2,
	"REC_TYPE_ALL":     3,
}

func (x RecognizeType) String() string {
	return proto.EnumName(RecognizeType_name, int32(x))
}
func (RecognizeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type RankType int32

const (
	RankType_RANK_TYPE_DEFAULT RankType = 0
	RankType_RANK_TYPE_VEHICLE RankType = 1
	RankType_RANK_TYPE_FACE    RankType = 2
)

var RankType_name = map[int32]string{
	0: "RANK_TYPE_DEFAULT",
	1: "RANK_TYPE_VEHICLE",
	2: "RANK_TYPE_FACE",
}
var RankType_value = map[string]int32{
	"RANK_TYPE_DEFAULT": 0,
	"RANK_TYPE_VEHICLE": 1,
	"RANK_TYPE_FACE":    2,
}

func (x RankType) String() string {
	return proto.EnumName(RankType_name, int32(x))
}
func (RankType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

type RecognizeFunctions int32

const (
	RecognizeFunctions_RECFUNC_NONE                    RecognizeFunctions = 0
	RecognizeFunctions_RECFUNC_VEHICLE                 RecognizeFunctions = 1
	RecognizeFunctions_RECFUNC_VEHICLE_DETECT          RecognizeFunctions = 10
	RecognizeFunctions_RECFUNC_VEHICLE_TRACK           RecognizeFunctions = 11
	RecognizeFunctions_RECFUNC_VEHICLE_STYLE           RecognizeFunctions = 12
	RecognizeFunctions_RECFUNC_VEHICLE_COLOR           RecognizeFunctions = 13
	RecognizeFunctions_RECFUNC_VEHICLE_MARKER          RecognizeFunctions = 14
	RecognizeFunctions_RECFUNC_VEHICLE_PLATE           RecognizeFunctions = 15
	RecognizeFunctions_RECFUNC_VEHICLE_FEATURE_VECTOR  RecognizeFunctions = 16
	RecognizeFunctions_RECFUNC_VEHICLE_DRIVER_NOBELT   RecognizeFunctions = 170
	RecognizeFunctions_RECFUNC_VEHICLE_DRIVER_PHONE    RecognizeFunctions = 171
	RecognizeFunctions_RECFUNC_VEHICLE_CODRIVER_NOBELT RecognizeFunctions = 172
	RecognizeFunctions_RECFUNC_FACE                    RecognizeFunctions = 2
	RecognizeFunctions_RECFUNC_FACE_DETECTOR           RecognizeFunctions = 20
	RecognizeFunctions_RECFUNC_FACE_FEATURE_VECTOR     RecognizeFunctions = 21
	RecognizeFunctions_RECFUNC_FACE_ALIGNMENT          RecognizeFunctions = 22
	RecognizeFunctions_RECFUNC_FACE_QUALITY            RecognizeFunctions = 23
	RecognizeFunctions_RECFUNC_PEDESTRIAN_ATTR         RecognizeFunctions = 3
	RecognizeFunctions_RECFUNC_NON_VEHICLE_ATTR        RecognizeFunctions = 4
)

var RecognizeFunctions_name = map[int32]string{
	0:   "RECFUNC_NONE",
	1:   "RECFUNC_VEHICLE",
	10:  "RECFUNC_VEHICLE_DETECT",
	11:  "RECFUNC_VEHICLE_TRACK",
	12:  "RECFUNC_VEHICLE_STYLE",
	13:  "RECFUNC_VEHICLE_COLOR",
	14:  "RECFUNC_VEHICLE_MARKER",
	15:  "RECFUNC_VEHICLE_PLATE",
	16:  "RECFUNC_VEHICLE_FEATURE_VECTOR",
	170: "RECFUNC_VEHICLE_DRIVER_NOBELT",
	171: "RECFUNC_VEHICLE_DRIVER_PHONE",
	172: "RECFUNC_VEHICLE_CODRIVER_NOBELT",
	2:   "RECFUNC_FACE",
	20:  "RECFUNC_FACE_DETECTOR",
	21:  "RECFUNC_FACE_FEATURE_VECTOR",
	22:  "RECFUNC_FACE_ALIGNMENT",
	23:  "RECFUNC_FACE_QUALITY",
	3:   "RECFUNC_PEDESTRIAN_ATTR",
	4:   "RECFUNC_NON_VEHICLE_ATTR",
}
var RecognizeFunctions_value = map[string]int32{
	"RECFUNC_NONE":                    0,
	"RECFUNC_VEHICLE":                 1,
	"RECFUNC_VEHICLE_DETECT":          10,
	"RECFUNC_VEHICLE_TRACK":           11,
	"RECFUNC_VEHICLE_STYLE":           12,
	"RECFUNC_VEHICLE_COLOR":           13,
	"RECFUNC_VEHICLE_MARKER":          14,
	"RECFUNC_VEHICLE_PLATE":           15,
	"RECFUNC_VEHICLE_FEATURE_VECTOR":  16,
	"RECFUNC_VEHICLE_DRIVER_NOBELT":   170,
	"RECFUNC_VEHICLE_DRIVER_PHONE":    171,
	"RECFUNC_VEHICLE_CODRIVER_NOBELT": 172,
	"RECFUNC_FACE":                    2,
	"RECFUNC_FACE_DETECTOR":           20,
	"RECFUNC_FACE_FEATURE_VECTOR":     21,
	"RECFUNC_FACE_ALIGNMENT":          22,
	"RECFUNC_FACE_QUALITY":            23,
	"RECFUNC_PEDESTRIAN_ATTR":         3,
	"RECFUNC_NON_VEHICLE_ATTR":        4,
}

func (x RecognizeFunctions) String() string {
	return proto.EnumName(RecognizeFunctions_name, int32(x))
}
func (RecognizeFunctions) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

type Time struct {
	Seconds  int64 `protobuf:"varint,1,opt,name=Seconds,json=seconds" json:"Seconds"`
	NanoSecs int64 `protobuf:"varint,2,opt,name=NanoSecs,json=nanoSecs" json:"NanoSecs"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type MatrixError struct {
	Code    int32  `protobuf:"varint,1,opt,name=Code,json=code" json:"Code"`
	Message string `protobuf:"bytes,2,opt,name=Message,json=message" json:"Message"`
}

func (m *MatrixError) Reset()                    { *m = MatrixError{} }
func (m *MatrixError) String() string            { return proto.CompactTextString(m) }
func (*MatrixError) ProtoMessage()               {}
func (*MatrixError) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func init() {
	proto.RegisterType((*Time)(nil), "dg.model.Time")
	proto.RegisterType((*MatrixError)(nil), "dg.model.MatrixError")
	proto.RegisterEnum("dg.model.RecognizeType", RecognizeType_name, RecognizeType_value)
	proto.RegisterEnum("dg.model.RankType", RankType_name, RankType_value)
	proto.RegisterEnum("dg.model.RecognizeFunctions", RecognizeFunctions_name, RecognizeFunctions_value)
}

func init() { proto.RegisterFile("matrix.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0x5f, 0x73, 0xd2, 0x40,
	0x14, 0xc5, 0x4d, 0x89, 0x82, 0xb7, 0xb4, 0xdd, 0x5e, 0xa1, 0xc5, 0xb6, 0x5a, 0x65, 0x7c, 0x70,
	0x78, 0xe8, 0x8b, 0x8f, 0xfa, 0xb2, 0x6e, 0x6f, 0x2c, 0x43, 0x08, 0xb8, 0x5c, 0x3a, 0xc3, 0x13,
	0x22, 0x64, 0x18, 0x46, 0x49, 0x3a, 0x09, 0xce, 0xa8, 0x1f, 0x49, 0x9d, 0xf1, 0x2b, 0xba, 0x84,
	0xbf, 0x49, 0xf1, 0x8d, 0x3d, 0x67, 0xf7, 0xb7, 0xe7, 0x9e, 0x25, 0x50, 0x9c, 0x0e, 0x66, 0xd1,
	0xe4, 0xfb, 0xd5, 0x5d, 0x14, 0xce, 0x42, 0x2c, 0x8c, 0xc6, 0x57, 0xd3, 0x70, 0xe4, 0x7f, 0xad,
	0xbe, 0x03, 0x9b, 0x27, 0x53, 0x1f, 0x2b, 0x90, 0xef, 0xf8, 0xc3, 0x30, 0x18, 0xc5, 0x15, 0xeb,
	0x85, 0xf5, 0x3a, 0xa7, 0xf3, 0xf1, 0x62, 0x89, 0x67, 0x50, 0xf0, 0x06, 0x41, 0x68, 0xdc, 0xb8,
	0xb2, 0x97, 0x58, 0x85, 0x60, 0xb9, 0xae, 0xbe, 0x85, 0xfd, 0x66, 0xc2, 0xa5, 0x28, 0x0a, 0x23,
	0x44, 0xb0, 0x95, 0xa1, 0x26, 0x84, 0x87, 0xda, 0x1e, 0x9a, 0xdf, 0x73, 0x70, 0xd3, 0x8f, 0xe3,
	0xc1, 0xd8, 0x4f, 0x4e, 0x3f, 0xd6, 0xf9, 0xe9, 0x62, 0x59, 0xfb, 0x04, 0x07, 0xda, 0xdc, 0x31,
	0x0e, 0x26, 0x3f, 0x7d, 0xfe, 0x71, 0xe7, 0x63, 0x09, 0x84, 0x26, 0xd5, 0xe7, 0x5e, 0x9b, 0xfa,
	0xd7, 0xe4, 0xc8, 0xae, 0xcb, 0xe2, 0x41, 0x4a, 0xbd, 0xa5, 0x9b, 0xba, 0x72, 0x49, 0x58, 0x78,
	0x6c, 0x0e, 0xaf, 0x54, 0x47, 0x2a, 0x12, 0x7b, 0x28, 0xa0, 0xb8, 0x96, 0xa4, 0xeb, 0x8a, 0x5c,
	0xcd, 0x85, 0x82, 0x1e, 0x04, 0x5f, 0x12, 0x78, 0x19, 0x8e, 0xb5, 0xf4, 0x1a, 0x59, 0x7a, 0x4a,
	0xde, 0xe0, 0x11, 0x0e, 0x37, 0xf2, 0x82, 0x5f, 0xfb, 0x6b, 0x03, 0xae, 0x03, 0x3b, 0xdf, 0x82,
	0xe1, 0x6c, 0x12, 0x06, 0xf1, 0xf2, 0x5a, 0xa7, 0xeb, 0xa9, 0xbe, 0xd7, 0xf2, 0xc8, 0x30, 0x9f,
	0xc0, 0xd1, 0x4a, 0xd9, 0x10, 0xcf, 0xe0, 0x24, 0x23, 0x9a, 0x14, 0x4c, 0x8a, 0x05, 0xe0, 0x53,
	0x28, 0x67, 0x3d, 0xd6, 0x52, 0x35, 0xc4, 0xfe, 0x2e, 0xab, 0xc3, 0x3d, 0x43, 0x2c, 0xee, 0xb2,
	0x54, 0xcb, 0x6d, 0x69, 0x71, 0xb0, 0xeb, 0xb2, 0xa6, 0xd4, 0x0d, 0xd2, 0xe2, 0x70, 0xd7, 0xb1,
	0xb6, 0x2b, 0x99, 0xc4, 0x11, 0x56, 0xe1, 0x79, 0xd6, 0x72, 0x48, 0x72, 0x57, 0xcf, 0xab, 0x51,
	0x6c, 0xd0, 0xc2, 0xec, 0x79, 0x76, 0x6f, 0x0e, 0x5d, 0xbf, 0x25, 0x6d, 0xa6, 0x7f, 0x4f, 0xa6,
	0xd3, 0x5f, 0x16, 0xbe, 0x84, 0x8b, 0xff, 0xec, 0x69, 0xdf, 0xcc, 0x2b, 0xfa, 0x6d, 0xe1, 0x2b,
	0xb8, 0xbc, 0x1f, 0x3e, 0x0d, 0xfa, 0x63, 0x6d, 0x77, 0xbb, 0x7c, 0xe4, 0xad, 0xf4, 0x73, 0x65,
	0xd9, 0xa1, 0x49, 0x56, 0xc2, 0x4b, 0x38, 0x4f, 0x59, 0x99, 0xe8, 0xe5, 0xed, 0x56, 0x92, 0x0d,
	0xd2, 0xad, 0x7f, 0xf0, 0x9a, 0xe4, 0xb1, 0x38, 0x31, 0x7f, 0xd3, 0x52, 0xca, 0xfb, 0xd8, 0x35,
	0x2e, 0xf7, 0xc4, 0x29, 0x9e, 0xc3, 0xe9, 0xca, 0x69, 0xd3, 0x35, 0x75, 0x58, 0xd7, 0xa5, 0xd7,
	0x97, 0xcc, 0x5a, 0xe4, 0xf0, 0x02, 0x2a, 0x5b, 0x8f, 0xbf, 0x1e, 0x25, 0x71, 0xed, 0xcf, 0x8f,
	0x92, 0xaf, 0xed, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x7f, 0xbe, 0x6d, 0x7d, 0x03,
	0x00, 0x00,
}