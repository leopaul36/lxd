// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lxd/migration/migrate.proto

/*
Package migration is a generated protocol buffer package.

It is generated from these files:
	lxd/migration/migrate.proto

It has these top-level messages:
	IDMapType
	Config
	Device
	Snapshot
	MigrationHeader
	MigrationControl
*/
package migration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MigrationFSType int32

const (
	MigrationFSType_RSYNC MigrationFSType = 0
	MigrationFSType_BTRFS MigrationFSType = 1
	MigrationFSType_ZFS   MigrationFSType = 2
)

var MigrationFSType_name = map[int32]string{
	0: "RSYNC",
	1: "BTRFS",
	2: "ZFS",
}
var MigrationFSType_value = map[string]int32{
	"RSYNC": 0,
	"BTRFS": 1,
	"ZFS":   2,
}

func (x MigrationFSType) Enum() *MigrationFSType {
	p := new(MigrationFSType)
	*p = x
	return p
}
func (x MigrationFSType) String() string {
	return proto.EnumName(MigrationFSType_name, int32(x))
}
func (x *MigrationFSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MigrationFSType_value, data, "MigrationFSType")
	if err != nil {
		return err
	}
	*x = MigrationFSType(value)
	return nil
}
func (MigrationFSType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CRIUType int32

const (
	CRIUType_CRIU_RSYNC CRIUType = 0
	CRIUType_PHAUL      CRIUType = 1
	CRIUType_NONE       CRIUType = 2
)

var CRIUType_name = map[int32]string{
	0: "CRIU_RSYNC",
	1: "PHAUL",
	2: "NONE",
}
var CRIUType_value = map[string]int32{
	"CRIU_RSYNC": 0,
	"PHAUL":      1,
	"NONE":       2,
}

func (x CRIUType) Enum() *CRIUType {
	p := new(CRIUType)
	*p = x
	return p
}
func (x CRIUType) String() string {
	return proto.EnumName(CRIUType_name, int32(x))
}
func (x *CRIUType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CRIUType_value, data, "CRIUType")
	if err != nil {
		return err
	}
	*x = CRIUType(value)
	return nil
}
func (CRIUType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IDMapType struct {
	Isuid            *bool  `protobuf:"varint,1,req,name=isuid" json:"isuid,omitempty"`
	Isgid            *bool  `protobuf:"varint,2,req,name=isgid" json:"isgid,omitempty"`
	Hostid           *int32 `protobuf:"varint,3,req,name=hostid" json:"hostid,omitempty"`
	Nsid             *int32 `protobuf:"varint,4,req,name=nsid" json:"nsid,omitempty"`
	Maprange         *int32 `protobuf:"varint,5,req,name=maprange" json:"maprange,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IDMapType) Reset()                    { *m = IDMapType{} }
func (m *IDMapType) String() string            { return proto.CompactTextString(m) }
func (*IDMapType) ProtoMessage()               {}
func (*IDMapType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IDMapType) GetIsuid() bool {
	if m != nil && m.Isuid != nil {
		return *m.Isuid
	}
	return false
}

func (m *IDMapType) GetIsgid() bool {
	if m != nil && m.Isgid != nil {
		return *m.Isgid
	}
	return false
}

func (m *IDMapType) GetHostid() int32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *IDMapType) GetNsid() int32 {
	if m != nil && m.Nsid != nil {
		return *m.Nsid
	}
	return 0
}

func (m *IDMapType) GetMaprange() int32 {
	if m != nil && m.Maprange != nil {
		return *m.Maprange
	}
	return 0
}

type Config struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Device struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config           []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Snapshot struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	LocalConfig      []*Config `protobuf:"bytes,2,rep,name=localConfig" json:"localConfig,omitempty"`
	Profiles         []string  `protobuf:"bytes,3,rep,name=profiles" json:"profiles,omitempty"`
	Ephemeral        *bool     `protobuf:"varint,4,req,name=ephemeral" json:"ephemeral,omitempty"`
	LocalDevices     []*Device `protobuf:"bytes,5,rep,name=localDevices" json:"localDevices,omitempty"`
	Architecture     *int32    `protobuf:"varint,6,req,name=architecture" json:"architecture,omitempty"`
	Stateful         *bool     `protobuf:"varint,7,req,name=stateful" json:"stateful,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Snapshot) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Snapshot) GetLocalConfig() []*Config {
	if m != nil {
		return m.LocalConfig
	}
	return nil
}

func (m *Snapshot) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Snapshot) GetEphemeral() bool {
	if m != nil && m.Ephemeral != nil {
		return *m.Ephemeral
	}
	return false
}

func (m *Snapshot) GetLocalDevices() []*Device {
	if m != nil {
		return m.LocalDevices
	}
	return nil
}

func (m *Snapshot) GetArchitecture() int32 {
	if m != nil && m.Architecture != nil {
		return *m.Architecture
	}
	return 0
}

func (m *Snapshot) GetStateful() bool {
	if m != nil && m.Stateful != nil {
		return *m.Stateful
	}
	return false
}

type MigrationHeader struct {
	Fs               *MigrationFSType `protobuf:"varint,1,req,name=fs,enum=migration.MigrationFSType" json:"fs,omitempty"`
	Criu             *CRIUType        `protobuf:"varint,2,opt,name=criu,enum=migration.CRIUType" json:"criu,omitempty"`
	Idmap            []*IDMapType     `protobuf:"bytes,3,rep,name=idmap" json:"idmap,omitempty"`
	SnapshotNames    []string         `protobuf:"bytes,4,rep,name=snapshotNames" json:"snapshotNames,omitempty"`
	Snapshots        []*Snapshot      `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MigrationHeader) Reset()                    { *m = MigrationHeader{} }
func (m *MigrationHeader) String() string            { return proto.CompactTextString(m) }
func (*MigrationHeader) ProtoMessage()               {}
func (*MigrationHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MigrationHeader) GetFs() MigrationFSType {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return MigrationFSType_RSYNC
}

func (m *MigrationHeader) GetCriu() CRIUType {
	if m != nil && m.Criu != nil {
		return *m.Criu
	}
	return CRIUType_CRIU_RSYNC
}

func (m *MigrationHeader) GetIdmap() []*IDMapType {
	if m != nil {
		return m.Idmap
	}
	return nil
}

func (m *MigrationHeader) GetSnapshotNames() []string {
	if m != nil {
		return m.SnapshotNames
	}
	return nil
}

func (m *MigrationHeader) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type MigrationControl struct {
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// optional failure message if sending a failure
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MigrationControl) Reset()                    { *m = MigrationControl{} }
func (m *MigrationControl) String() string            { return proto.CompactTextString(m) }
func (*MigrationControl) ProtoMessage()               {}
func (*MigrationControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MigrationControl) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *MigrationControl) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IDMapType)(nil), "migration.IDMapType")
	proto.RegisterType((*Config)(nil), "migration.Config")
	proto.RegisterType((*Device)(nil), "migration.Device")
	proto.RegisterType((*Snapshot)(nil), "migration.Snapshot")
	proto.RegisterType((*MigrationHeader)(nil), "migration.MigrationHeader")
	proto.RegisterType((*MigrationControl)(nil), "migration.MigrationControl")
	proto.RegisterEnum("migration.MigrationFSType", MigrationFSType_name, MigrationFSType_value)
	proto.RegisterEnum("migration.CRIUType", CRIUType_name, CRIUType_value)
}

func init() { proto.RegisterFile("lxd/migration/migrate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x62, 0xc7, 0xa9, 0x7d, 0xdb, 0x2f, 0x98, 0x4b, 0x85, 0x46, 0x85, 0x45, 0x64, 0x21,
	0x11, 0xb2, 0x48, 0x4b, 0x10, 0x0f, 0x00, 0x29, 0xa1, 0x95, 0x68, 0x40, 0x93, 0x76, 0x01, 0x1b,
	0x34, 0xb2, 0x6f, 0x9c, 0x11, 0xfe, 0x93, 0xc7, 0xae, 0xe8, 0x8a, 0x0d, 0x6f, 0xcb, 0x4b, 0x20,
	0x8f, 0x7f, 0xea, 0x08, 0x24, 0x76, 0xe7, 0x9c, 0x39, 0xbe, 0xc7, 0xf7, 0xcc, 0xc0, 0x93, 0xe8,
	0x7b, 0x70, 0x1a, 0xcb, 0x30, 0x17, 0x85, 0x4c, 0x93, 0x06, 0xd1, 0x3c, 0xcb, 0xd3, 0x22, 0x45,
	0xa7, 0x3b, 0xf0, 0x7e, 0x80, 0x73, 0x79, 0x7e, 0x25, 0xb2, 0xeb, 0xbb, 0x8c, 0xf0, 0x18, 0x2c,
	0xa9, 0x4a, 0x19, 0xb0, 0xc1, 0xc4, 0x98, 0xda, 0xbc, 0x26, 0xb5, 0x1a, 0xca, 0x80, 0x19, 0xad,
	0x1a, 0xca, 0x00, 0x1f, 0xc3, 0x68, 0x97, 0xaa, 0x42, 0x06, 0xcc, 0x9c, 0x18, 0x53, 0x8b, 0x37,
	0x0c, 0x11, 0x86, 0x89, 0x92, 0x01, 0x1b, 0x6a, 0x55, 0x63, 0x3c, 0x01, 0x3b, 0x16, 0x59, 0x2e,
	0x92, 0x90, 0x98, 0xa5, 0xf5, 0x8e, 0x7b, 0x67, 0x30, 0x5a, 0xa6, 0xc9, 0x56, 0x86, 0xe8, 0x82,
	0xf9, 0x8d, 0xee, 0x74, 0xb6, 0xc3, 0x2b, 0x58, 0x25, 0xdf, 0x8a, 0xa8, 0x24, 0x9d, 0xec, 0xf0,
	0x9a, 0x78, 0xef, 0x61, 0x74, 0x4e, 0xb7, 0xd2, 0x27, 0x9d, 0x25, 0x62, 0x6a, 0x3e, 0xd1, 0x18,
	0x5f, 0xc0, 0xc8, 0xd7, 0xf3, 0x98, 0x31, 0x31, 0xa7, 0x87, 0x8b, 0x87, 0xf3, 0x6e, 0xd9, 0x79,
	0x1d, 0xc4, 0x1b, 0x83, 0xf7, 0xd3, 0x00, 0x7b, 0x93, 0x88, 0x4c, 0xed, 0xd2, 0xe2, 0xaf, 0xb3,
	0x5e, 0xc1, 0x61, 0x94, 0xfa, 0x22, 0x5a, 0xfe, 0x63, 0x60, 0xdf, 0x55, 0x2d, 0x9b, 0xe5, 0xe9,
	0x56, 0x46, 0xa4, 0x98, 0x39, 0x31, 0xa7, 0x0e, 0xef, 0x38, 0x3e, 0x05, 0x87, 0xb2, 0x1d, 0xc5,
	0x94, 0x8b, 0x48, 0x37, 0x64, 0xf3, 0x7b, 0x01, 0x5f, 0xc3, 0x91, 0x1e, 0x54, 0x6f, 0xa7, 0x98,
	0xf5, 0x47, 0x5e, 0x7d, 0xc2, 0xf7, 0x6c, 0xe8, 0xc1, 0x91, 0xc8, 0xfd, 0x9d, 0x2c, 0xc8, 0x2f,
	0xca, 0x9c, 0xd8, 0x48, 0x37, 0xbc, 0xa7, 0x55, 0x3f, 0xa5, 0x0a, 0x51, 0xd0, 0xb6, 0x8c, 0xd8,
	0x81, 0xce, 0xed, 0xb8, 0xf7, 0x6b, 0x00, 0x0f, 0xae, 0xda, 0x88, 0x0b, 0x12, 0x01, 0xe5, 0x38,
	0x03, 0x63, 0xab, 0x74, 0x17, 0xe3, 0xc5, 0x49, 0xef, 0x07, 0x3a, 0xdf, 0x6a, 0x53, 0xbd, 0x18,
	0x6e, 0x6c, 0x15, 0x3e, 0x87, 0xa1, 0x9f, 0xcb, 0x92, 0x19, 0x93, 0xc1, 0x74, 0xbc, 0x78, 0xd4,
	0xaf, 0x87, 0x5f, 0xde, 0x68, 0x9b, 0x36, 0xe0, 0x0c, 0x2c, 0x19, 0xc4, 0x22, 0xd3, 0xb5, 0x1c,
	0x2e, 0x8e, 0x7b, 0xce, 0xee, 0x0d, 0xf2, 0xda, 0x82, 0xcf, 0xe0, 0x7f, 0xd5, 0x5c, 0xcd, 0x5a,
	0xc4, 0xa4, 0xd8, 0x50, 0x57, 0xb9, 0x2f, 0xe2, 0x4b, 0x70, 0x5a, 0xa1, 0xad, 0xab, 0x9f, 0xdf,
	0x5e, 0x2e, 0xbf, 0x77, 0x79, 0x2b, 0x70, 0xbb, 0x25, 0x96, 0x69, 0x52, 0xe4, 0x69, 0x84, 0x0c,
	0x0e, 0x54, 0xe9, 0xfb, 0xa4, 0x54, 0xf3, 0xf2, 0x5b, 0x5a, 0x9d, 0xc4, 0xa4, 0x94, 0x08, 0x49,
	0xaf, 0xe7, 0xf0, 0x96, 0xce, 0xce, 0x7a, 0xa5, 0xd5, 0x65, 0xa0, 0x03, 0x16, 0xdf, 0x7c, 0x5e,
	0x2f, 0xdd, 0xff, 0x2a, 0xf8, 0xf6, 0x9a, 0xaf, 0x36, 0xee, 0x00, 0x0f, 0xc0, 0xfc, 0xb2, 0xda,
	0xb8, 0xc6, 0xec, 0x14, 0xec, 0xb6, 0x10, 0x1c, 0x03, 0x54, 0xf8, 0x6b, 0xcf, 0xff, 0xe9, 0xe2,
	0xcd, 0xcd, 0x07, 0x77, 0x80, 0x36, 0x0c, 0xd7, 0x1f, 0xd7, 0xef, 0x5c, 0xe3, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x27, 0xb6, 0xaa, 0xc4, 0x03, 0x00, 0x00,
}
