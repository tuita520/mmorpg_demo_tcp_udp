// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Rqst_HeartBeating.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rqst_HeartBeating struct {
	Status           *int32 `protobuf:"varint,1,req,name=Status,def=0" json:"Status,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Rqst_HeartBeating) Reset()                    { *m = Rqst_HeartBeating{} }
func (m *Rqst_HeartBeating) String() string            { return proto.CompactTextString(m) }
func (*Rqst_HeartBeating) ProtoMessage()               {}
func (*Rqst_HeartBeating) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

const Default_Rqst_HeartBeating_Status int32 = 0

func (m *Rqst_HeartBeating) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_Rqst_HeartBeating_Status
}

func init() {
	proto.RegisterType((*Rqst_HeartBeating)(nil), "msg.Rqst_HeartBeating")
}

func init() { proto.RegisterFile("Rqst_HeartBeating.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0f, 0x2a, 0x2c, 0x2e,
	0x89, 0xf7, 0x48, 0x4d, 0x2c, 0x2a, 0x71, 0x4a, 0x4d, 0x2c, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x52, 0xe3, 0x12, 0xc4, 0x90, 0x17, 0x12,
	0xe4, 0x62, 0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd2, 0x60, 0xb5, 0x62,
	0x34, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x00, 0x7e, 0xf7, 0xcb, 0x46, 0x00, 0x00, 0x00,
}
