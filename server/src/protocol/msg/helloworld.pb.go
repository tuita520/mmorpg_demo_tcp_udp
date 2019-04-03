// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Helloworld struct {
	Id               *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Str              *string `protobuf:"bytes,2,req,name=str" json:"str,omitempty"`
	Opt              *int32  `protobuf:"varint,3,opt,name=opt" json:"opt,omitempty"`
	Ok               *OK     `protobuf:"bytes,4,req,name=ok" json:"ok,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Helloworld) Reset()                    { *m = Helloworld{} }
func (m *Helloworld) String() string            { return proto.CompactTextString(m) }
func (*Helloworld) ProtoMessage()               {}
func (*Helloworld) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *Helloworld) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Helloworld) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func (m *Helloworld) GetOpt() int32 {
	if m != nil && m.Opt != nil {
		return *m.Opt
	}
	return 0
}

func (m *Helloworld) GetOk() *OK {
	if m != nil {
		return m.Ok
	}
	return nil
}

func init() {
	proto.RegisterType((*Helloworld)(nil), "msg.helloworld")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d,
	0x4e, 0x97, 0xe2, 0xf0, 0xf7, 0x86, 0x70, 0x95, 0xdc, 0xb9, 0xb8, 0x10, 0x4a, 0x84, 0xb8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x85, 0xb8, 0xb9, 0x98, 0x8b, 0x4b, 0x8a,
	0x24, 0x98, 0x14, 0x98, 0x34, 0x38, 0x41, 0x9c, 0xfc, 0x82, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x56, 0x21, 0x61, 0x2e, 0xa6, 0xfc, 0x6c, 0x09, 0x16, 0x05, 0x26, 0x0d, 0x6e, 0x23, 0x76, 0xbd,
	0xdc, 0xe2, 0x74, 0x3d, 0x7f, 0x6f, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf5, 0xe5, 0x90,
	0x6a, 0x00, 0x00, 0x00,
}
