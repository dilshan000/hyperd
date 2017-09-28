// Code generated by protoc-gen-gogo.
// source: persist.proto
// DO NOT EDIT!

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import api "github.com/hyperhq/runv/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PersistPodLayout struct {
	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GlobalSpec string   `protobuf:"bytes,11,opt,name=globalSpec,proto3" json:"globalSpec,omitempty"`
	Containers []string `protobuf:"bytes,21,rep,name=containers" json:"containers,omitempty"`
	Volumes    []string `protobuf:"bytes,22,rep,name=volumes" json:"volumes,omitempty"`
	Interfaces []string `protobuf:"bytes,23,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *PersistPodLayout) Reset()                    { *m = PersistPodLayout{} }
func (m *PersistPodLayout) String() string            { return proto.CompactTextString(m) }
func (*PersistPodLayout) ProtoMessage()               {}
func (*PersistPodLayout) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{0} }

func (m *PersistPodLayout) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PersistPodLayout) GetGlobalSpec() string {
	if m != nil {
		return m.GlobalSpec
	}
	return ""
}

func (m *PersistPodLayout) GetContainers() []string {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *PersistPodLayout) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *PersistPodLayout) GetInterfaces() []string {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type PersistPodMeta struct {
	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Services  []*UserService    `protobuf:"bytes,11,rep,name=services" json:"services,omitempty"`
	Labels    map[string]string `protobuf:"bytes,12,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt int64             `protobuf:"varint,21,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *PersistPodMeta) Reset()                    { *m = PersistPodMeta{} }
func (m *PersistPodMeta) String() string            { return proto.CompactTextString(m) }
func (*PersistPodMeta) ProtoMessage()               {}
func (*PersistPodMeta) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{1} }

func (m *PersistPodMeta) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PersistPodMeta) GetServices() []*UserService {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *PersistPodMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *PersistPodMeta) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type SandboxPersistInfo struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PersistInfo []byte `protobuf:"bytes,2,opt,name=PersistInfo,proto3" json:"PersistInfo,omitempty"`
}

func (m *SandboxPersistInfo) Reset()                    { *m = SandboxPersistInfo{} }
func (m *SandboxPersistInfo) String() string            { return proto.CompactTextString(m) }
func (*SandboxPersistInfo) ProtoMessage()               {}
func (*SandboxPersistInfo) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{2} }

func (m *SandboxPersistInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SandboxPersistInfo) GetPersistInfo() []byte {
	if m != nil {
		return m.PersistInfo
	}
	return nil
}

type PersistContainer struct {
	Id       string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pod      string                    `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	Spec     *UserContainer            `protobuf:"bytes,11,opt,name=spec" json:"spec,omitempty"`
	Descript *api.ContainerDescription `protobuf:"bytes,12,opt,name=descript" json:"descript,omitempty"`
}

func (m *PersistContainer) Reset()                    { *m = PersistContainer{} }
func (m *PersistContainer) String() string            { return proto.CompactTextString(m) }
func (*PersistContainer) ProtoMessage()               {}
func (*PersistContainer) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{3} }

func (m *PersistContainer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PersistContainer) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *PersistContainer) GetSpec() *UserContainer {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PersistContainer) GetDescript() *api.ContainerDescription {
	if m != nil {
		return m.Descript
	}
	return nil
}

type PersistVolume struct {
	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pod      string                 `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	Spec     *UserVolume            `protobuf:"bytes,11,opt,name=spec" json:"spec,omitempty"`
	Descript *api.VolumeDescription `protobuf:"bytes,12,opt,name=descript" json:"descript,omitempty"`
}

func (m *PersistVolume) Reset()                    { *m = PersistVolume{} }
func (m *PersistVolume) String() string            { return proto.CompactTextString(m) }
func (*PersistVolume) ProtoMessage()               {}
func (*PersistVolume) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{4} }

func (m *PersistVolume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersistVolume) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *PersistVolume) GetSpec() *UserVolume {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PersistVolume) GetDescript() *api.VolumeDescription {
	if m != nil {
		return m.Descript
	}
	return nil
}

type PersistInterface struct {
	Id       string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pod      string                    `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	Spec     *UserInterface            `protobuf:"bytes,11,opt,name=spec" json:"spec,omitempty"`
	Descript *api.InterfaceDescription `protobuf:"bytes,12,opt,name=descript" json:"descript,omitempty"`
}

func (m *PersistInterface) Reset()                    { *m = PersistInterface{} }
func (m *PersistInterface) String() string            { return proto.CompactTextString(m) }
func (*PersistInterface) ProtoMessage()               {}
func (*PersistInterface) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{5} }

func (m *PersistInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PersistInterface) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *PersistInterface) GetSpec() *UserInterface {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PersistInterface) GetDescript() *api.InterfaceDescription {
	if m != nil {
		return m.Descript
	}
	return nil
}

type PersistPortmappings struct {
	Pod          string         `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	ContainerIP  string         `protobuf:"bytes,2,opt,name=containerIP,proto3" json:"containerIP,omitempty"`
	PortMappings []*PortMapping `protobuf:"bytes,11,rep,name=portMappings" json:"portMappings,omitempty"`
}

func (m *PersistPortmappings) Reset()                    { *m = PersistPortmappings{} }
func (m *PersistPortmappings) String() string            { return proto.CompactTextString(m) }
func (*PersistPortmappings) ProtoMessage()               {}
func (*PersistPortmappings) Descriptor() ([]byte, []int) { return fileDescriptorPersist, []int{6} }

func (m *PersistPortmappings) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *PersistPortmappings) GetContainerIP() string {
	if m != nil {
		return m.ContainerIP
	}
	return ""
}

func (m *PersistPortmappings) GetPortMappings() []*PortMapping {
	if m != nil {
		return m.PortMappings
	}
	return nil
}

func init() {
	proto.RegisterType((*PersistPodLayout)(nil), "types.PersistPodLayout")
	proto.RegisterType((*PersistPodMeta)(nil), "types.PersistPodMeta")
	proto.RegisterType((*SandboxPersistInfo)(nil), "types.SandboxPersistInfo")
	proto.RegisterType((*PersistContainer)(nil), "types.PersistContainer")
	proto.RegisterType((*PersistVolume)(nil), "types.PersistVolume")
	proto.RegisterType((*PersistInterface)(nil), "types.PersistInterface")
	proto.RegisterType((*PersistPortmappings)(nil), "types.PersistPortmappings")
}

func init() { proto.RegisterFile("persist.proto", fileDescriptorPersist) }

var fileDescriptorPersist = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0x76, 0xdb, 0x7f, 0x3b, 0xe9, 0xa6, 0xfe, 0xcd, 0x36, 0x4c, 0x85, 0x50, 0xa8,
	0x84, 0xd4, 0xab, 0x54, 0x2a, 0x02, 0x31, 0xee, 0x10, 0x2f, 0x52, 0xa5, 0x4d, 0xaa, 0x52, 0xc1,
	0xbd, 0x93, 0x78, 0xad, 0x45, 0x6a, 0x1b, 0xdb, 0xa9, 0xc8, 0x25, 0xdf, 0x80, 0x1b, 0xee, 0xf8,
	0x78, 0x7c, 0x10, 0x14, 0xc7, 0x79, 0x29, 0xad, 0xc4, 0x05, 0x77, 0xf6, 0x73, 0x9e, 0xf3, 0xf4,
	0x97, 0x63, 0xbb, 0x70, 0x2e, 0xa9, 0xd2, 0x4c, 0x9b, 0x50, 0x2a, 0x61, 0x04, 0x3a, 0x36, 0x85,
	0xa4, 0x7a, 0x14, 0xae, 0x98, 0x59, 0xe7, 0x71, 0x98, 0x88, 0xcd, 0x74, 0x5d, 0x48, 0xaa, 0xd6,
	0x5f, 0xa6, 0x2a, 0xe7, 0xdb, 0x29, 0x91, 0x6c, 0x9a, 0x52, 0x9d, 0x28, 0x26, 0x0d, 0x13, 0x5c,
	0x57, 0x6d, 0x23, 0xdf, 0xb6, 0x55, 0x9b, 0xf1, 0x4f, 0x0f, 0x86, 0x8b, 0x2a, 0x75, 0x21, 0xd2,
	0x5b, 0x52, 0x88, 0xdc, 0xa0, 0x0b, 0xe8, 0xb1, 0x14, 0x7b, 0x81, 0x37, 0x39, 0x8b, 0x7a, 0x2c,
	0x45, 0x4f, 0x00, 0x56, 0x99, 0x88, 0x49, 0xb6, 0x94, 0x34, 0xc1, 0xbe, 0xd5, 0x3b, 0x4a, 0x59,
	0x4f, 0x04, 0x37, 0x84, 0x71, 0xaa, 0x34, 0xbe, 0x0a, 0xfa, 0x65, 0xbd, 0x55, 0x10, 0x86, 0xff,
	0xb6, 0x22, 0xcb, 0x37, 0x54, 0xe3, 0x6b, 0x5b, 0xac, 0xb7, 0x65, 0x27, 0xe3, 0x86, 0xaa, 0x7b,
	0x92, 0x50, 0x8d, 0x1f, 0x56, 0x9d, 0xad, 0x32, 0xfe, 0xe5, 0xc1, 0x45, 0x8b, 0x77, 0x47, 0x0d,
	0xd9, 0x83, 0x0b, 0xe1, 0x54, 0x53, 0xb5, 0x65, 0x65, 0x80, 0x1f, 0xf4, 0x27, 0xfe, 0x0c, 0x85,
	0xd5, 0x17, 0x7e, 0xd4, 0x54, 0x2d, 0xab, 0x52, 0xd4, 0x78, 0xd0, 0x0d, 0x9c, 0x64, 0x24, 0xa6,
	0x99, 0xc6, 0x03, 0xeb, 0x7e, 0xea, 0xdc, 0xbb, 0x3f, 0x13, 0xde, 0x5a, 0xcf, 0x7b, 0x6e, 0x54,
	0x11, 0xb9, 0x06, 0xf4, 0x18, 0xce, 0x12, 0x45, 0x89, 0xa1, 0xe9, 0x1b, 0x83, 0xaf, 0x02, 0x6f,
	0xd2, 0x8f, 0x5a, 0x61, 0x74, 0x03, 0x7e, 0xa7, 0x09, 0x0d, 0xa1, 0xff, 0x99, 0x16, 0x0e, 0xb4,
	0x5c, 0xa2, 0x4b, 0x38, 0xde, 0x92, 0x2c, 0xa7, 0xb8, 0x67, 0xb5, 0x6a, 0xf3, 0xba, 0xf7, 0xca,
	0x1b, 0x7f, 0x00, 0xb4, 0x24, 0x3c, 0x8d, 0xc5, 0x57, 0x47, 0x31, 0xe7, 0xf7, 0x62, 0xef, 0x4b,
	0x03, 0xf0, 0x3b, 0x65, 0x9b, 0x32, 0x88, 0xba, 0xd2, 0xf8, 0x47, 0x7b, 0x9a, 0x6f, 0xeb, 0xf1,
	0xef, 0xc5, 0x0c, 0xa1, 0x2f, 0x45, 0xea, 0x20, 0xca, 0x25, 0x9a, 0xc0, 0x91, 0xae, 0x4f, 0xd6,
	0x9f, 0x5d, 0x76, 0xc6, 0xd7, 0xa4, 0x44, 0xd6, 0x81, 0x5e, 0xc0, 0x69, 0x7d, 0xa3, 0xf0, 0xc0,
	0xba, 0x1f, 0x85, 0x44, 0xb2, 0xb0, 0xf1, 0xbd, 0x6b, 0xef, 0x5b, 0xd4, 0x58, 0xc7, 0xdf, 0x3d,
	0x38, 0x77, 0x5c, 0x9f, 0xec, 0xc9, 0x23, 0x04, 0x47, 0x9c, 0x6c, 0xa8, 0xc3, 0xb2, 0xeb, 0x03,
	0x60, 0xcf, 0x76, 0xc0, 0xfe, 0xef, 0x80, 0x55, 0x31, 0x8e, 0x6a, 0xb6, 0x47, 0x75, 0x6d, 0xa9,
	0x2a, 0xd3, 0x61, 0xa4, 0xce, 0xa8, 0xe6, 0xf5, 0x7d, 0xfb, 0xa7, 0x51, 0x35, 0x29, 0x7f, 0x19,
	0x55, 0xe3, 0x3b, 0xcc, 0xf5, 0xcd, 0x83, 0x07, 0xcd, 0x55, 0x54, 0x66, 0x43, 0xa4, 0x64, 0x7c,
	0xa5, 0x6b, 0x14, 0xaf, 0x45, 0x09, 0xc0, 0x6f, 0xde, 0xd8, 0x7c, 0xe1, 0x20, 0xbb, 0x12, 0x7a,
	0x09, 0x03, 0x29, 0x94, 0xb9, 0x73, 0x19, 0x7f, 0x3c, 0x8f, 0x45, 0x5b, 0x8a, 0x76, 0x7c, 0xf1,
	0x89, 0xfd, 0x6f, 0x78, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x65, 0x89, 0x38, 0x67, 0x70, 0x04,
	0x00, 0x00,
}
