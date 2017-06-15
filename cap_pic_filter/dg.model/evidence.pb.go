// Code generated by protoc-gen-go.
// source: evidence.proto
// DO NOT EDIT!

package dg_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MetaEvidence struct {
	DeviceNo      string         `protobuf:"bytes,1,opt,name=DeviceNo,json=deviceNo" json:"DeviceNo"`
	Timestamp     int64          `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp"`
	LicencePlate  string         `protobuf:"bytes,3,opt,name=LicencePlate,json=licencePlate" json:"LicencePlate"`
	PlateType     int32          `protobuf:"varint,4,opt,name=PlateType,json=plateType" json:"PlateType"`
	OriginalImage *EvidenceImage `protobuf:"bytes,8,opt,name=OriginalImage,json=originalImage" json:"OriginalImage"`
	DriverImage   *EvidenceImage `protobuf:"bytes,9,opt,name=DriverImage,json=driverImage" json:"DriverImage"`
	CoDriverImage *EvidenceImage `protobuf:"bytes,10,opt,name=CoDriverImage,json=coDriverImage" json:"CoDriverImage"`
}

func (m *MetaEvidence) Reset()                    { *m = MetaEvidence{} }
func (m *MetaEvidence) String() string            { return proto.CompactTextString(m) }
func (*MetaEvidence) ProtoMessage()               {}
func (*MetaEvidence) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *MetaEvidence) GetOriginalImage() *EvidenceImage {
	if m != nil {
		return m.OriginalImage
	}
	return nil
}

func (m *MetaEvidence) GetDriverImage() *EvidenceImage {
	if m != nil {
		return m.DriverImage
	}
	return nil
}

func (m *MetaEvidence) GetCoDriverImage() *EvidenceImage {
	if m != nil {
		return m.CoDriverImage
	}
	return nil
}

type EvidenceImage struct {
	URI     string `protobuf:"bytes,1,opt,name=URI,json=uRI" json:"URI"`
	BinData []byte `protobuf:"bytes,2,opt,name=BinData,json=binData,proto3" json:"BinData"`
}

func (m *EvidenceImage) Reset()                    { *m = EvidenceImage{} }
func (m *EvidenceImage) String() string            { return proto.CompactTextString(m) }
func (*EvidenceImage) ProtoMessage()               {}
func (*EvidenceImage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func init() {
	proto.RegisterType((*MetaEvidence)(nil), "dg.model.MetaEvidence")
	proto.RegisterType((*EvidenceImage)(nil), "dg.model.EvidenceImage")
}

func init() { proto.RegisterFile("evidence.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xd9, 0xad, 0xba, 0xcd, 0xb4, 0x15, 0x99, 0x8b, 0x41, 0x3c, 0x48, 0x4f, 0x9e, 0x7a,
	0xd0, 0x93, 0x88, 0x17, 0xad, 0x87, 0x82, 0xff, 0x08, 0xeb, 0x03, 0x64, 0x9b, 0xa1, 0x04, 0xda,
	0xa6, 0xd4, 0xb8, 0xe0, 0xf3, 0xf9, 0x62, 0x66, 0xd3, 0x16, 0xeb, 0xc5, 0xbd, 0x65, 0xbe, 0xef,
	0xf7, 0xcd, 0x0c, 0x13, 0x38, 0xa6, 0xad, 0x56, 0xd4, 0x96, 0x94, 0x75, 0xbd, 0xb1, 0x06, 0x43,
	0x55, 0x65, 0x8d, 0x51, 0x54, 0xa7, 0xdf, 0x4b, 0x88, 0x9f, 0xc9, 0xca, 0xc7, 0x11, 0xc0, 0x33,
	0x08, 0x73, 0x47, 0x97, 0xf4, 0x62, 0xf8, 0xe2, 0x62, 0x71, 0xc9, 0x44, 0xa8, 0xc6, 0x1a, 0xcf,
	0x81, 0xad, 0x75, 0x43, 0x1f, 0x56, 0x36, 0x1d, 0x5f, 0x3a, 0x33, 0x10, 0xcc, 0x4e, 0x02, 0xa6,
	0x10, 0x3f, 0x39, 0xce, 0x35, 0x79, 0xab, 0xa5, 0x25, 0x1e, 0xf8, 0x74, 0x5c, 0xcf, 0xb4, 0x5d,
	0x07, 0xff, 0x58, 0x7f, 0x75, 0xc4, 0x0f, 0x1c, 0x70, 0x28, 0x58, 0x37, 0x09, 0x78, 0x07, 0xc9,
	0x6b, 0xaf, 0x2b, 0xdd, 0xca, 0xba, 0x68, 0x64, 0x45, 0x3c, 0x74, 0x44, 0x74, 0x75, 0x9a, 0x4d,
	0xeb, 0x66, 0xd3, 0x9a, 0xde, 0x16, 0x89, 0x99, 0xd3, 0x78, 0x03, 0x51, 0xde, 0xeb, 0x2d, 0xf5,
	0x43, 0x98, 0xfd, 0x1f, 0x8e, 0xd4, 0x2f, 0xbb, 0x9b, 0xfc, 0x60, 0xe6, 0x61, 0xd8, 0x33, 0xb9,
	0x9c, 0xd3, 0xe9, 0x2d, 0x24, 0x7f, 0x7c, 0x3c, 0x81, 0xe0, 0x5d, 0x14, 0xe3, 0x01, 0x83, 0x4f,
	0x51, 0x20, 0x87, 0xd5, 0xbd, 0x6e, 0x73, 0x69, 0xa5, 0xbf, 0x5c, 0x2c, 0x56, 0x9b, 0xa1, 0xdc,
	0x1c, 0xf9, 0x3f, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x24, 0x11, 0x52, 0x10, 0xa5, 0x01,
	0x00, 0x00,
}
