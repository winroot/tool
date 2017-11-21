// Code generated by protoc-gen-go.
// source: deepdata.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CutboardType int32

const (
	CutboardType_CUTBOARD_TYPE_UNKNOWN CutboardType = 0
	CutboardType_CUTBOARD_TYPE_VEHICLE CutboardType = 1
	CutboardType_CUTBOARD_TYPE_SYMBOL  CutboardType = 2
	CutboardType_CUTBOARD_TYPE_PLATE   CutboardType = 3
)

var CutboardType_name = map[int32]string{
	0: "CUTBOARD_TYPE_UNKNOWN",
	1: "CUTBOARD_TYPE_VEHICLE",
	2: "CUTBOARD_TYPE_SYMBOL",
	3: "CUTBOARD_TYPE_PLATE",
}
var CutboardType_value = map[string]int32{
	"CUTBOARD_TYPE_UNKNOWN": 0,
	"CUTBOARD_TYPE_VEHICLE": 1,
	"CUTBOARD_TYPE_SYMBOL":  2,
	"CUTBOARD_TYPE_PLATE":   3,
}

func (x CutboardType) String() string {
	return proto.EnumName(CutboardType_name, int32(x))
}
func (CutboardType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GenericObj struct {
	ObjType ObjType     `protobuf:"varint,1,opt,name=ObjType,json=objType,enum=dg.model.ObjType" json:"ObjType,omitempty"`
	FmtType DataFmtType `protobuf:"varint,2,opt,name=FmtType,json=fmtType,enum=dg.model.DataFmtType" json:"FmtType,omitempty"`
	BinData []byte      `protobuf:"bytes,3,opt,name=BinData,json=binData,proto3" json:"BinData,omitempty"`
}

func (m *GenericObj) Reset()                    { *m = GenericObj{} }
func (m *GenericObj) String() string            { return proto.CompactTextString(m) }
func (*GenericObj) ProtoMessage()               {}
func (*GenericObj) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GenericObj) GetObjType() ObjType {
	if m != nil {
		return m.ObjType
	}
	return ObjType_OBJ_TYPE_UNKNOWN
}

func (m *GenericObj) GetFmtType() DataFmtType {
	if m != nil {
		return m.FmtType
	}
	return DataFmtType_UNKNOWNFMT
}

func (m *GenericObj) GetBinData() []byte {
	if m != nil {
		return m.BinData
	}
	return nil
}

func init() {
	proto.RegisterType((*GenericObj)(nil), "dg.model.GenericObj")
	proto.RegisterEnum("dg.model.CutboardType", CutboardType_name, CutboardType_value)
}

func init() { proto.RegisterFile("deepdata.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x4d, 0x2d,
	0x48, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x49, 0xd7, 0xcb,
	0xcd, 0x4f, 0x49, 0xcd, 0x91, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x2b, 0xb5,
	0x30, 0x72, 0x71, 0xb9, 0xa7, 0xe6, 0xa5, 0x16, 0x65, 0x26, 0xfb, 0x27, 0x65, 0x09, 0x69, 0x73,
	0xb1, 0xfb, 0x27, 0x65, 0x85, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x09,
	0xea, 0xc1, 0x34, 0xea, 0x41, 0x25, 0x82, 0xd8, 0xf3, 0x21, 0x0c, 0x21, 0x7d, 0x2e, 0x76, 0xb7,
	0xdc, 0x12, 0xb0, 0x62, 0x26, 0xb0, 0x62, 0x51, 0x84, 0x62, 0x97, 0xc4, 0x92, 0x44, 0xa8, 0x64,
	0x10, 0x7b, 0x1a, 0x84, 0x21, 0x24, 0xc1, 0xc5, 0xee, 0x94, 0x99, 0x07, 0x92, 0x92, 0x60, 0x56,
	0x60, 0xd4, 0xe0, 0x09, 0x62, 0x4f, 0x82, 0x70, 0xb5, 0xca, 0xb9, 0x78, 0x9c, 0x4b, 0x4b, 0x92,
	0xf2, 0x13, 0x8b, 0x52, 0xc0, 0x2a, 0x25, 0xb9, 0x44, 0x9d, 0x43, 0x43, 0x9c, 0xfc, 0x1d, 0x83,
	0x5c, 0xe2, 0x43, 0x22, 0x03, 0x5c, 0xe3, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18,
	0x30, 0xa5, 0xc2, 0x5c, 0x3d, 0x3c, 0x9d, 0x7d, 0x5c, 0x05, 0x18, 0x85, 0x24, 0xb8, 0x44, 0x50,
	0xa5, 0x82, 0x23, 0x7d, 0x9d, 0xfc, 0x7d, 0x04, 0x98, 0x84, 0xc4, 0xb9, 0x84, 0x51, 0x65, 0x02,
	0x7c, 0x1c, 0x43, 0x5c, 0x05, 0x98, 0x93, 0xd8, 0xc0, 0xc1, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0x42, 0xe1, 0x40, 0x30, 0x01, 0x00, 0x00,
}
