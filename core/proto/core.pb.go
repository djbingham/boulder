// Code generated by protoc-gen-go.
// source: core/proto/core.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	core/proto/core.proto

It has these top-level messages:
	Challenge
	ValidationRecord
	ProblemDetails
	Precertificate
	SCTFetchingConfig
	SCTFetchingLogSet
	Certificate
	Registration
	Authorization
	Order
	Empty
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Challenge struct {
	Id                *int64              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type              *string             `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Status            *string             `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Uri               *string             `protobuf:"bytes,9,opt,name=uri" json:"uri,omitempty"`
	Token             *string             `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	KeyAuthorization  *string             `protobuf:"bytes,5,opt,name=keyAuthorization" json:"keyAuthorization,omitempty"`
	Validationrecords []*ValidationRecord `protobuf:"bytes,10,rep,name=validationrecords" json:"validationrecords,omitempty"`
	Error             *ProblemDetails     `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *Challenge) Reset()                    { *m = Challenge{} }
func (m *Challenge) String() string            { return proto1.CompactTextString(m) }
func (*Challenge) ProtoMessage()               {}
func (*Challenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Challenge) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Challenge) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Challenge) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Challenge) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Challenge) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Challenge) GetKeyAuthorization() string {
	if m != nil && m.KeyAuthorization != nil {
		return *m.KeyAuthorization
	}
	return ""
}

func (m *Challenge) GetValidationrecords() []*ValidationRecord {
	if m != nil {
		return m.Validationrecords
	}
	return nil
}

func (m *Challenge) GetError() *ProblemDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

type ValidationRecord struct {
	Hostname          *string  `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port              *string  `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	AddressesResolved [][]byte `protobuf:"bytes,3,rep,name=addressesResolved" json:"addressesResolved,omitempty"`
	AddressUsed       []byte   `protobuf:"bytes,4,opt,name=addressUsed" json:"addressUsed,omitempty"`
	Authorities       []string `protobuf:"bytes,5,rep,name=authorities" json:"authorities,omitempty"`
	Url               *string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	// A list of addresses tried before the address used (see
	// core/objects.go and the comment on the ValidationRecord structure
	// definition for more information.
	AddressesTried   [][]byte `protobuf:"bytes,7,rep,name=addressesTried" json:"addressesTried,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ValidationRecord) Reset()                    { *m = ValidationRecord{} }
func (m *ValidationRecord) String() string            { return proto1.CompactTextString(m) }
func (*ValidationRecord) ProtoMessage()               {}
func (*ValidationRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidationRecord) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ValidationRecord) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *ValidationRecord) GetAddressesResolved() [][]byte {
	if m != nil {
		return m.AddressesResolved
	}
	return nil
}

func (m *ValidationRecord) GetAddressUsed() []byte {
	if m != nil {
		return m.AddressUsed
	}
	return nil
}

func (m *ValidationRecord) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *ValidationRecord) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *ValidationRecord) GetAddressesTried() [][]byte {
	if m != nil {
		return m.AddressesTried
	}
	return nil
}

type ProblemDetails struct {
	ProblemType      *string `protobuf:"bytes,1,opt,name=problemType" json:"problemType,omitempty"`
	Detail           *string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	HttpStatus       *int32  `protobuf:"varint,3,opt,name=httpStatus" json:"httpStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProblemDetails) Reset()                    { *m = ProblemDetails{} }
func (m *ProblemDetails) String() string            { return proto1.CompactTextString(m) }
func (*ProblemDetails) ProtoMessage()               {}
func (*ProblemDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProblemDetails) GetProblemType() string {
	if m != nil && m.ProblemType != nil {
		return *m.ProblemType
	}
	return ""
}

func (m *ProblemDetails) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *ProblemDetails) GetHttpStatus() int32 {
	if m != nil && m.HttpStatus != nil {
		return *m.HttpStatus
	}
	return 0
}

type Precertificate struct {
	Der              []byte `protobuf:"bytes,1,opt,name=der" json:"der,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Precertificate) Reset()                    { *m = Precertificate{} }
func (m *Precertificate) String() string            { return proto1.CompactTextString(m) }
func (*Precertificate) ProtoMessage()               {}
func (*Precertificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Precertificate) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

type SCTFetchingConfig struct {
	LogSets          []*SCTFetchingLogSet `protobuf:"bytes,1,rep,name=logSets" json:"logSets,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SCTFetchingConfig) Reset()                    { *m = SCTFetchingConfig{} }
func (m *SCTFetchingConfig) String() string            { return proto1.CompactTextString(m) }
func (*SCTFetchingConfig) ProtoMessage()               {}
func (*SCTFetchingConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SCTFetchingConfig) GetLogSets() []*SCTFetchingLogSet {
	if m != nil {
		return m.LogSets
	}
	return nil
}

type SCTFetchingLogSet struct {
	LogURLs          []string `protobuf:"bytes,1,rep,name=logURLs" json:"logURLs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SCTFetchingLogSet) Reset()                    { *m = SCTFetchingLogSet{} }
func (m *SCTFetchingLogSet) String() string            { return proto1.CompactTextString(m) }
func (*SCTFetchingLogSet) ProtoMessage()               {}
func (*SCTFetchingLogSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SCTFetchingLogSet) GetLogURLs() []string {
	if m != nil {
		return m.LogURLs
	}
	return nil
}

type Certificate struct {
	RegistrationID   *int64  `protobuf:"varint,1,opt,name=registrationID" json:"registrationID,omitempty"`
	Serial           *string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	Digest           *string `protobuf:"bytes,3,opt,name=digest" json:"digest,omitempty"`
	Der              []byte  `protobuf:"bytes,4,opt,name=der" json:"der,omitempty"`
	Issued           *int64  `protobuf:"varint,5,opt,name=issued" json:"issued,omitempty"`
	Expires          *int64  `protobuf:"varint,6,opt,name=expires" json:"expires,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto1.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Certificate) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Certificate) GetSerial() string {
	if m != nil && m.Serial != nil {
		return *m.Serial
	}
	return ""
}

func (m *Certificate) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

func (m *Certificate) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

func (m *Certificate) GetIssued() int64 {
	if m != nil && m.Issued != nil {
		return *m.Issued
	}
	return 0
}

func (m *Certificate) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

type Registration struct {
	Id               *int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Key              []byte   `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Contact          []string `protobuf:"bytes,3,rep,name=contact" json:"contact,omitempty"`
	ContactsPresent  *bool    `protobuf:"varint,4,opt,name=contactsPresent" json:"contactsPresent,omitempty"`
	Agreement        *string  `protobuf:"bytes,5,opt,name=agreement" json:"agreement,omitempty"`
	InitialIP        []byte   `protobuf:"bytes,6,opt,name=initialIP" json:"initialIP,omitempty"`
	CreatedAt        *int64   `protobuf:"varint,7,opt,name=createdAt" json:"createdAt,omitempty"`
	Status           *string  `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto1.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Registration) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Registration) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Registration) GetContact() []string {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Registration) GetContactsPresent() bool {
	if m != nil && m.ContactsPresent != nil {
		return *m.ContactsPresent
	}
	return false
}

func (m *Registration) GetAgreement() string {
	if m != nil && m.Agreement != nil {
		return *m.Agreement
	}
	return ""
}

func (m *Registration) GetInitialIP() []byte {
	if m != nil {
		return m.InitialIP
	}
	return nil
}

func (m *Registration) GetCreatedAt() int64 {
	if m != nil && m.CreatedAt != nil {
		return *m.CreatedAt
	}
	return 0
}

func (m *Registration) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type Authorization struct {
	Id               *string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Identifier       *string      `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
	RegistrationID   *int64       `protobuf:"varint,3,opt,name=registrationID" json:"registrationID,omitempty"`
	Status           *string      `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Expires          *int64       `protobuf:"varint,5,opt,name=expires" json:"expires,omitempty"`
	Challenges       []*Challenge `protobuf:"bytes,6,rep,name=challenges" json:"challenges,omitempty"`
	Combinations     []byte       `protobuf:"bytes,7,opt,name=combinations" json:"combinations,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Authorization) Reset()                    { *m = Authorization{} }
func (m *Authorization) String() string            { return proto1.CompactTextString(m) }
func (*Authorization) ProtoMessage()               {}
func (*Authorization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Authorization) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Authorization) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *Authorization) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Authorization) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Authorization) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

func (m *Authorization) GetChallenges() []*Challenge {
	if m != nil {
		return m.Challenges
	}
	return nil
}

func (m *Authorization) GetCombinations() []byte {
	if m != nil {
		return m.Combinations
	}
	return nil
}

type Order struct {
	Id                *int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	RegistrationID    *int64   `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	Expires           *int64   `protobuf:"varint,3,opt,name=expires" json:"expires,omitempty"`
	Csr               []byte   `protobuf:"bytes,4,opt,name=csr" json:"csr,omitempty"`
	Error             []byte   `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	CertificateSerial *string  `protobuf:"bytes,6,opt,name=certificateSerial" json:"certificateSerial,omitempty"`
	Authorizations    []string `protobuf:"bytes,7,rep,name=authorizations" json:"authorizations,omitempty"`
	Status            *string  `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto1.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Order) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Order) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Order) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

func (m *Order) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *Order) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Order) GetCertificateSerial() string {
	if m != nil && m.CertificateSerial != nil {
		return *m.CertificateSerial
	}
	return ""
}

func (m *Order) GetAuthorizations() []string {
	if m != nil {
		return m.Authorizations
	}
	return nil
}

func (m *Order) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto1.RegisterType((*Challenge)(nil), "core.Challenge")
	proto1.RegisterType((*ValidationRecord)(nil), "core.ValidationRecord")
	proto1.RegisterType((*ProblemDetails)(nil), "core.ProblemDetails")
	proto1.RegisterType((*Precertificate)(nil), "core.Precertificate")
	proto1.RegisterType((*SCTFetchingConfig)(nil), "core.SCTFetchingConfig")
	proto1.RegisterType((*SCTFetchingLogSet)(nil), "core.SCTFetchingLogSet")
	proto1.RegisterType((*Certificate)(nil), "core.Certificate")
	proto1.RegisterType((*Registration)(nil), "core.Registration")
	proto1.RegisterType((*Authorization)(nil), "core.Authorization")
	proto1.RegisterType((*Order)(nil), "core.Order")
	proto1.RegisterType((*Empty)(nil), "core.Empty")
}

func init() { proto1.RegisterFile("core/proto/core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x5f, 0x4e, 0xf3, 0x46,
	0x10, 0x97, 0xe3, 0x98, 0xe0, 0x89, 0x1b, 0xc2, 0x8a, 0x52, 0xab, 0xaa, 0x90, 0xe5, 0x07, 0x64,
	0xa1, 0x16, 0x54, 0x6e, 0x40, 0x43, 0x91, 0x90, 0x90, 0x8a, 0x36, 0xd0, 0x87, 0xbe, 0x19, 0x7b,
	0x48, 0xb6, 0x38, 0xde, 0x68, 0x77, 0x83, 0x9a, 0xde, 0xa1, 0x07, 0xe9, 0xc5, 0xfa, 0xd8, 0xef,
	0x0a, 0x9f, 0x76, 0xd6, 0x49, 0x9c, 0x04, 0xde, 0x66, 0x7e, 0x33, 0xde, 0xf9, 0xf3, 0x9b, 0x19,
	0xc3, 0xb7, 0x85, 0x54, 0x78, 0x35, 0x57, 0xd2, 0xc8, 0x2b, 0x2b, 0x5e, 0x92, 0xc8, 0xba, 0x56,
	0x4e, 0xff, 0xe9, 0x40, 0x38, 0x9a, 0xe6, 0x55, 0x85, 0xf5, 0x04, 0xd9, 0x00, 0x3a, 0xa2, 0x8c,
	0xbd, 0xc4, 0xcb, 0x7c, 0xde, 0x11, 0x25, 0x63, 0xd0, 0x35, 0xcb, 0x39, 0xc6, 0x9d, 0xc4, 0xcb,
	0x42, 0x4e, 0x32, 0x3b, 0x85, 0x03, 0x6d, 0x72, 0xb3, 0xd0, 0xf1, 0x01, 0xa1, 0x8d, 0xc6, 0x86,
	0xe0, 0x2f, 0x94, 0x88, 0x43, 0x02, 0xad, 0xc8, 0x4e, 0x20, 0x30, 0xf2, 0x0d, 0xeb, 0xd8, 0x27,
	0xcc, 0x29, 0xec, 0x02, 0x86, 0x6f, 0xb8, 0xbc, 0x59, 0x98, 0xa9, 0x54, 0xe2, 0xef, 0xdc, 0x08,
	0x59, 0xc7, 0x01, 0x39, 0xec, 0xe1, 0xec, 0x16, 0x8e, 0xdf, 0xf3, 0x4a, 0x94, 0xa4, 0x29, 0x2c,
	0xa4, 0x2a, 0x75, 0x0c, 0x89, 0x9f, 0xf5, 0xaf, 0x4f, 0x2f, 0xa9, 0x96, 0xdf, 0xd7, 0x66, 0x4e,
	0x66, 0xbe, 0xff, 0x01, 0xbb, 0x80, 0x00, 0x95, 0x92, 0x2a, 0xee, 0x25, 0x5e, 0xd6, 0xbf, 0x3e,
	0x71, 0x5f, 0x3e, 0x2a, 0xf9, 0x52, 0xe1, 0xec, 0x16, 0x4d, 0x2e, 0x2a, 0xcd, 0x9d, 0x4b, 0xfa,
	0xbf, 0x07, 0xc3, 0xdd, 0x37, 0xd9, 0xf7, 0x70, 0x38, 0x95, 0xda, 0xd4, 0xf9, 0x0c, 0xa9, 0x39,
	0x21, 0x5f, 0xeb, 0xb6, 0x45, 0x73, 0xa9, 0xcc, 0xaa, 0x45, 0x56, 0x66, 0x3f, 0xc2, 0x71, 0x5e,
	0x96, 0x0a, 0xb5, 0x46, 0xcd, 0x51, 0xcb, 0xea, 0x1d, 0xcb, 0xd8, 0x4f, 0xfc, 0x2c, 0xe2, 0xfb,
	0x06, 0x96, 0x40, 0xbf, 0x01, 0x9f, 0x35, 0x96, 0x71, 0x37, 0xf1, 0xb2, 0x88, 0xb7, 0x21, 0xf2,
	0x70, 0x7d, 0x31, 0x02, 0x75, 0x1c, 0x24, 0x7e, 0x16, 0xf2, 0x36, 0xe4, 0x9a, 0x5f, 0x35, 0x8c,
	0x58, 0x91, 0x9d, 0xc3, 0x60, 0x1d, 0xea, 0x49, 0x09, 0x2c, 0xe3, 0x1e, 0x25, 0xb0, 0x83, 0xa6,
	0x7f, 0xc2, 0x60, 0xbb, 0x13, 0x36, 0xda, 0xdc, 0x21, 0x4f, 0x96, 0x7b, 0x57, 0x70, 0x1b, 0xb2,
	0x23, 0x50, 0x92, 0x73, 0x53, 0x75, 0xa3, 0xb1, 0x33, 0x80, 0xa9, 0x31, 0xf3, 0xb1, 0x1b, 0x0f,
	0xcb, 0x7a, 0xc0, 0x5b, 0x48, 0x9a, 0xda, 0x58, 0x58, 0xa0, 0x32, 0xe2, 0x55, 0x14, 0xb9, 0x41,
	0x9b, 0x77, 0x89, 0x8a, 0x62, 0x44, 0xdc, 0x8a, 0xe9, 0x1d, 0x1c, 0x8f, 0x47, 0x4f, 0x77, 0x68,
	0x8a, 0xa9, 0xa8, 0x27, 0x23, 0x59, 0xbf, 0x8a, 0x09, 0xfb, 0x19, 0x7a, 0x95, 0x9c, 0x8c, 0xd1,
	0xe8, 0xd8, 0x23, 0xf6, 0xbf, 0x73, 0x1c, 0xb6, 0x3c, 0x1f, 0xc8, 0xce, 0x57, 0x7e, 0xe9, 0x4f,
	0x5b, 0xef, 0x38, 0x2b, 0x8b, 0xe9, 0x9d, 0x67, 0xfe, 0xe0, 0xde, 0x09, 0xf9, 0x4a, 0x4d, 0xff,
	0xf5, 0xa0, 0x3f, 0x6a, 0x25, 0x76, 0x0e, 0x03, 0x85, 0x13, 0xa1, 0x8d, 0xa2, 0x41, 0xb8, 0xbf,
	0x6d, 0xb6, 0x62, 0x07, 0xa5, 0x6d, 0x40, 0x25, 0xf2, 0x75, 0x2b, 0x9c, 0x46, 0x2d, 0x12, 0x13,
	0xd4, 0xa6, 0x19, 0xfe, 0x46, 0x5b, 0x15, 0xdc, 0x5d, 0x17, 0x6c, 0x3d, 0x85, 0xd6, 0x0b, 0x2c,
	0x69, 0x0b, 0x7c, 0xde, 0x68, 0x36, 0x57, 0xfc, 0x6b, 0x2e, 0x14, 0xba, 0x45, 0xf3, 0xf9, 0x4a,
	0x4d, 0xff, 0xf3, 0x20, 0xe2, 0xad, 0x34, 0xf6, 0xd6, 0x76, 0x08, 0xfe, 0x1b, 0x2e, 0x29, 0xa3,
	0x88, 0x5b, 0xd1, 0x3e, 0x56, 0xc8, 0xda, 0xe4, 0x85, 0xa1, 0x39, 0x0c, 0xf9, 0x4a, 0x65, 0x19,
	0x1c, 0x35, 0xa2, 0x7e, 0x54, 0xa8, 0xb1, 0x36, 0x94, 0xdc, 0x21, 0xdf, 0x85, 0xd9, 0x0f, 0x10,
	0xe6, 0x13, 0x85, 0x38, 0xb3, 0x3e, 0x6e, 0x63, 0x37, 0x80, 0xb5, 0x8a, 0x5a, 0x18, 0x91, 0x57,
	0xf7, 0x8f, 0x94, 0x70, 0xc4, 0x37, 0x80, 0xb5, 0x16, 0x0a, 0x73, 0x83, 0xe5, 0x8d, 0xa1, 0x35,
	0xf4, 0xf9, 0x06, 0x68, 0x9d, 0x94, 0xc3, 0xf6, 0x49, 0xb1, 0xcb, 0xf8, 0xcd, 0xf6, 0x41, 0xd8,
	0x54, 0x1a, 0x52, 0xa5, 0x67, 0x00, 0xa2, 0xc4, 0xda, 0xd2, 0x86, 0xaa, 0xa1, 0xa0, 0x85, 0x7c,
	0x40, 0xa3, 0xff, 0x29, 0x8d, 0x2e, 0x83, 0xee, 0xd6, 0x51, 0x6b, 0x91, 0x10, 0x6c, 0x91, 0xc0,
	0xae, 0x00, 0x8a, 0xd5, 0xdd, 0xb4, 0x0c, 0xd9, 0xa9, 0x3c, 0x72, 0x53, 0xb9, 0xbe, 0xa7, 0xbc,
	0xe5, 0xc2, 0x52, 0x88, 0x0a, 0x39, 0x7b, 0x11, 0x35, 0xc5, 0xd4, 0xd4, 0x85, 0x88, 0x6f, 0x61,
	0xe9, 0x17, 0x0f, 0x82, 0xdf, 0x94, 0x9d, 0x8a, 0x5d, 0x4a, 0xf7, 0x0b, 0xe9, 0x7c, 0x58, 0x48,
	0x2b, 0x61, 0x7f, 0x3b, 0xe1, 0x21, 0xf8, 0x85, 0x5e, 0x4f, 0x5e, 0xa1, 0x95, 0xbd, 0xcf, 0xee,
	0x2e, 0x06, 0x84, 0x39, 0xc5, 0x1e, 0xaf, 0xd6, 0x86, 0x8e, 0xdd, 0x70, 0xbb, 0xc3, 0xb2, 0x6f,
	0xa0, 0x33, 0xd3, 0x66, 0x48, 0xd3, 0x99, 0x09, 0xf9, 0x0e, 0xfa, 0x29, 0xc5, 0x3d, 0x08, 0x7e,
	0x9d, 0xcd, 0xcd, 0xf2, 0x97, 0xde, 0x1f, 0x01, 0xfd, 0x97, 0xbe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x8b, 0xbb, 0x88, 0xaf, 0x06, 0x00, 0x00,
}
