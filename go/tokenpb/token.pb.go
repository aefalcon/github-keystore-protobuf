// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token.proto

package tokenpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Links struct {
	AppTokens            string   `protobuf:"bytes,1,opt,name=app_tokens,json=appTokens,proto3" json:"app_tokens,omitempty"`
	InstallTokens        string   `protobuf:"bytes,2,opt,name=install_tokens,json=installTokens,proto3" json:"install_tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Links) Reset()         { *m = Links{} }
func (m *Links) String() string { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()    {}
func (*Links) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{0}
}
func (m *Links) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Links.Unmarshal(m, b)
}
func (m *Links) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Links.Marshal(b, m, deterministic)
}
func (dst *Links) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Links.Merge(dst, src)
}
func (m *Links) XXX_Size() int {
	return xxx_messageInfo_Links.Size(m)
}
func (m *Links) XXX_DiscardUnknown() {
	xxx_messageInfo_Links.DiscardUnknown(m)
}

var xxx_messageInfo_Links proto.InternalMessageInfo

func (m *Links) GetAppTokens() string {
	if m != nil {
		return m.AppTokens
	}
	return ""
}

func (m *Links) GetInstallTokens() string {
	if m != nil {
		return m.InstallTokens
	}
	return ""
}

type AppToken struct {
	App                  uint64               `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Token                []byte               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AppToken) Reset()         { *m = AppToken{} }
func (m *AppToken) String() string { return proto.CompactTextString(m) }
func (*AppToken) ProtoMessage()    {}
func (*AppToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{1}
}
func (m *AppToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppToken.Unmarshal(m, b)
}
func (m *AppToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppToken.Marshal(b, m, deterministic)
}
func (dst *AppToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppToken.Merge(dst, src)
}
func (m *AppToken) XXX_Size() int {
	return xxx_messageInfo_AppToken.Size(m)
}
func (m *AppToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AppToken.DiscardUnknown(m)
}

var xxx_messageInfo_AppToken proto.InternalMessageInfo

func (m *AppToken) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *AppToken) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AppToken) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type InstallToken struct {
	App                  uint64               `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Install              uint64               `protobuf:"varint,2,opt,name=install,proto3" json:"install,omitempty"`
	Token                []byte               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InstallToken) Reset()         { *m = InstallToken{} }
func (m *InstallToken) String() string { return proto.CompactTextString(m) }
func (*InstallToken) ProtoMessage()    {}
func (*InstallToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{2}
}
func (m *InstallToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallToken.Unmarshal(m, b)
}
func (m *InstallToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallToken.Marshal(b, m, deterministic)
}
func (dst *InstallToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallToken.Merge(dst, src)
}
func (m *InstallToken) XXX_Size() int {
	return xxx_messageInfo_InstallToken.Size(m)
}
func (m *InstallToken) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallToken.DiscardUnknown(m)
}

var xxx_messageInfo_InstallToken proto.InternalMessageInfo

func (m *InstallToken) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *InstallToken) GetInstall() uint64 {
	if m != nil {
		return m.Install
	}
	return 0
}

func (m *InstallToken) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *InstallToken) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type GetAppTokenRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppTokenRequest) Reset()         { *m = GetAppTokenRequest{} }
func (m *GetAppTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppTokenRequest) ProtoMessage()    {}
func (*GetAppTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{3}
}
func (m *GetAppTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppTokenRequest.Unmarshal(m, b)
}
func (m *GetAppTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GetAppTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppTokenRequest.Merge(dst, src)
}
func (m *GetAppTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppTokenRequest.Size(m)
}
func (m *GetAppTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppTokenRequest proto.InternalMessageInfo

func (m *GetAppTokenRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type GetAppTokenResponse struct {
	Token                *AppToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetAppTokenResponse) Reset()         { *m = GetAppTokenResponse{} }
func (m *GetAppTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppTokenResponse) ProtoMessage()    {}
func (*GetAppTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{4}
}
func (m *GetAppTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppTokenResponse.Unmarshal(m, b)
}
func (m *GetAppTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GetAppTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppTokenResponse.Merge(dst, src)
}
func (m *GetAppTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppTokenResponse.Size(m)
}
func (m *GetAppTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppTokenResponse proto.InternalMessageInfo

func (m *GetAppTokenResponse) GetToken() *AppToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetInstallTokenRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Install              uint64   `protobuf:"varint,2,opt,name=install,proto3" json:"install,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstallTokenRequest) Reset()         { *m = GetInstallTokenRequest{} }
func (m *GetInstallTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstallTokenRequest) ProtoMessage()    {}
func (*GetInstallTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{5}
}
func (m *GetInstallTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstallTokenRequest.Unmarshal(m, b)
}
func (m *GetInstallTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstallTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GetInstallTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstallTokenRequest.Merge(dst, src)
}
func (m *GetInstallTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstallTokenRequest.Size(m)
}
func (m *GetInstallTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstallTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstallTokenRequest proto.InternalMessageInfo

func (m *GetInstallTokenRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *GetInstallTokenRequest) GetInstall() uint64 {
	if m != nil {
		return m.Install
	}
	return 0
}

type GetInstallTokenResponse struct {
	Token                *InstallToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetInstallTokenResponse) Reset()         { *m = GetInstallTokenResponse{} }
func (m *GetInstallTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstallTokenResponse) ProtoMessage()    {}
func (*GetInstallTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{6}
}
func (m *GetInstallTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstallTokenResponse.Unmarshal(m, b)
}
func (m *GetInstallTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstallTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GetInstallTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstallTokenResponse.Merge(dst, src)
}
func (m *GetInstallTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetInstallTokenResponse.Size(m)
}
func (m *GetInstallTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstallTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstallTokenResponse proto.InternalMessageInfo

func (m *GetInstallTokenResponse) GetToken() *InstallToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type PutAppTokenRequest struct {
	Token                *AppToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PutAppTokenRequest) Reset()         { *m = PutAppTokenRequest{} }
func (m *PutAppTokenRequest) String() string { return proto.CompactTextString(m) }
func (*PutAppTokenRequest) ProtoMessage()    {}
func (*PutAppTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{7}
}
func (m *PutAppTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutAppTokenRequest.Unmarshal(m, b)
}
func (m *PutAppTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutAppTokenRequest.Marshal(b, m, deterministic)
}
func (dst *PutAppTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutAppTokenRequest.Merge(dst, src)
}
func (m *PutAppTokenRequest) XXX_Size() int {
	return xxx_messageInfo_PutAppTokenRequest.Size(m)
}
func (m *PutAppTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutAppTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutAppTokenRequest proto.InternalMessageInfo

func (m *PutAppTokenRequest) GetToken() *AppToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type PutAppTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutAppTokenResponse) Reset()         { *m = PutAppTokenResponse{} }
func (m *PutAppTokenResponse) String() string { return proto.CompactTextString(m) }
func (*PutAppTokenResponse) ProtoMessage()    {}
func (*PutAppTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{8}
}
func (m *PutAppTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutAppTokenResponse.Unmarshal(m, b)
}
func (m *PutAppTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutAppTokenResponse.Marshal(b, m, deterministic)
}
func (dst *PutAppTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutAppTokenResponse.Merge(dst, src)
}
func (m *PutAppTokenResponse) XXX_Size() int {
	return xxx_messageInfo_PutAppTokenResponse.Size(m)
}
func (m *PutAppTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutAppTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutAppTokenResponse proto.InternalMessageInfo

type PutInstallTokenRequest struct {
	Token                *InstallToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PutInstallTokenRequest) Reset()         { *m = PutInstallTokenRequest{} }
func (m *PutInstallTokenRequest) String() string { return proto.CompactTextString(m) }
func (*PutInstallTokenRequest) ProtoMessage()    {}
func (*PutInstallTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{9}
}
func (m *PutInstallTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutInstallTokenRequest.Unmarshal(m, b)
}
func (m *PutInstallTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutInstallTokenRequest.Marshal(b, m, deterministic)
}
func (dst *PutInstallTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutInstallTokenRequest.Merge(dst, src)
}
func (m *PutInstallTokenRequest) XXX_Size() int {
	return xxx_messageInfo_PutInstallTokenRequest.Size(m)
}
func (m *PutInstallTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutInstallTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutInstallTokenRequest proto.InternalMessageInfo

func (m *PutInstallTokenRequest) GetToken() *InstallToken {
	if m != nil {
		return m.Token
	}
	return nil
}

type PutInstallTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutInstallTokenResponse) Reset()         { *m = PutInstallTokenResponse{} }
func (m *PutInstallTokenResponse) String() string { return proto.CompactTextString(m) }
func (*PutInstallTokenResponse) ProtoMessage()    {}
func (*PutInstallTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{10}
}
func (m *PutInstallTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutInstallTokenResponse.Unmarshal(m, b)
}
func (m *PutInstallTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutInstallTokenResponse.Marshal(b, m, deterministic)
}
func (dst *PutInstallTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutInstallTokenResponse.Merge(dst, src)
}
func (m *PutInstallTokenResponse) XXX_Size() int {
	return xxx_messageInfo_PutInstallTokenResponse.Size(m)
}
func (m *PutInstallTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutInstallTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutInstallTokenResponse proto.InternalMessageInfo

type DeleteAppTokenRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAppTokenRequest) Reset()         { *m = DeleteAppTokenRequest{} }
func (m *DeleteAppTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAppTokenRequest) ProtoMessage()    {}
func (*DeleteAppTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{11}
}
func (m *DeleteAppTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAppTokenRequest.Unmarshal(m, b)
}
func (m *DeleteAppTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAppTokenRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteAppTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAppTokenRequest.Merge(dst, src)
}
func (m *DeleteAppTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAppTokenRequest.Size(m)
}
func (m *DeleteAppTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAppTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAppTokenRequest proto.InternalMessageInfo

func (m *DeleteAppTokenRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type DeleteAppTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAppTokenResponse) Reset()         { *m = DeleteAppTokenResponse{} }
func (m *DeleteAppTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAppTokenResponse) ProtoMessage()    {}
func (*DeleteAppTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{12}
}
func (m *DeleteAppTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAppTokenResponse.Unmarshal(m, b)
}
func (m *DeleteAppTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAppTokenResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteAppTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAppTokenResponse.Merge(dst, src)
}
func (m *DeleteAppTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAppTokenResponse.Size(m)
}
func (m *DeleteAppTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAppTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAppTokenResponse proto.InternalMessageInfo

type DeleteInstallTokenRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Install              uint64   `protobuf:"varint,2,opt,name=install,proto3" json:"install,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstallTokenRequest) Reset()         { *m = DeleteInstallTokenRequest{} }
func (m *DeleteInstallTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstallTokenRequest) ProtoMessage()    {}
func (*DeleteInstallTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{13}
}
func (m *DeleteInstallTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstallTokenRequest.Unmarshal(m, b)
}
func (m *DeleteInstallTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstallTokenRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteInstallTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstallTokenRequest.Merge(dst, src)
}
func (m *DeleteInstallTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInstallTokenRequest.Size(m)
}
func (m *DeleteInstallTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstallTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstallTokenRequest proto.InternalMessageInfo

func (m *DeleteInstallTokenRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *DeleteInstallTokenRequest) GetInstall() uint64 {
	if m != nil {
		return m.Install
	}
	return 0
}

type DeleteInstallTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstallTokenResponse) Reset()         { *m = DeleteInstallTokenResponse{} }
func (m *DeleteInstallTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteInstallTokenResponse) ProtoMessage()    {}
func (*DeleteInstallTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_b9d5b4aaffb4e8fe, []int{14}
}
func (m *DeleteInstallTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstallTokenResponse.Unmarshal(m, b)
}
func (m *DeleteInstallTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstallTokenResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteInstallTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstallTokenResponse.Merge(dst, src)
}
func (m *DeleteInstallTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteInstallTokenResponse.Size(m)
}
func (m *DeleteInstallTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstallTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstallTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Links)(nil), "token.Links")
	proto.RegisterType((*AppToken)(nil), "token.AppToken")
	proto.RegisterType((*InstallToken)(nil), "token.InstallToken")
	proto.RegisterType((*GetAppTokenRequest)(nil), "token.GetAppTokenRequest")
	proto.RegisterType((*GetAppTokenResponse)(nil), "token.GetAppTokenResponse")
	proto.RegisterType((*GetInstallTokenRequest)(nil), "token.GetInstallTokenRequest")
	proto.RegisterType((*GetInstallTokenResponse)(nil), "token.GetInstallTokenResponse")
	proto.RegisterType((*PutAppTokenRequest)(nil), "token.PutAppTokenRequest")
	proto.RegisterType((*PutAppTokenResponse)(nil), "token.PutAppTokenResponse")
	proto.RegisterType((*PutInstallTokenRequest)(nil), "token.PutInstallTokenRequest")
	proto.RegisterType((*PutInstallTokenResponse)(nil), "token.PutInstallTokenResponse")
	proto.RegisterType((*DeleteAppTokenRequest)(nil), "token.DeleteAppTokenRequest")
	proto.RegisterType((*DeleteAppTokenResponse)(nil), "token.DeleteAppTokenResponse")
	proto.RegisterType((*DeleteInstallTokenRequest)(nil), "token.DeleteInstallTokenRequest")
	proto.RegisterType((*DeleteInstallTokenResponse)(nil), "token.DeleteInstallTokenResponse")
}

func init() { proto.RegisterFile("token.proto", fileDescriptor_token_b9d5b4aaffb4e8fe) }

var fileDescriptor_token_b9d5b4aaffb4e8fe = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0xce, 0xd2, 0x40,
	0x10, 0xc7, 0x53, 0x0b, 0x22, 0x03, 0x82, 0xd9, 0x0a, 0x94, 0x06, 0x14, 0x9b, 0x68, 0xe4, 0x52,
	0x12, 0xbc, 0xa9, 0x17, 0x95, 0x84, 0x98, 0x48, 0x42, 0x0a, 0x17, 0xbd, 0x18, 0x30, 0x2b, 0x69,
	0x28, 0xdd, 0xda, 0x2e, 0xc6, 0x87, 0xf0, 0x39, 0x7d, 0x0e, 0x97, 0xdd, 0x6d, 0x69, 0xe9, 0x56,
	0x3f, 0xf2, 0xe5, 0xbb, 0xd1, 0x99, 0xff, 0xcc, 0xfc, 0x66, 0xff, 0x13, 0xa0, 0x41, 0xc9, 0x1e,
	0x07, 0x4e, 0x18, 0x11, 0x4a, 0x50, 0x95, 0x7f, 0x58, 0x4f, 0x77, 0x84, 0xec, 0x7c, 0x3c, 0xe1,
	0xc1, 0xed, 0xf1, 0xfb, 0x84, 0x7a, 0x07, 0x1c, 0xd3, 0xcd, 0x21, 0x14, 0x3a, 0x7b, 0x01, 0xd5,
	0x4f, 0x5e, 0xb0, 0x8f, 0xd1, 0x10, 0x60, 0x13, 0x86, 0x5f, 0x79, 0x59, 0x6c, 0x6a, 0x23, 0xed,
	0x65, 0xdd, 0xad, 0xb3, 0xc8, 0x9a, 0x07, 0xd0, 0x73, 0x68, 0x79, 0x01, 0x2b, 0xf4, 0xfd, 0x44,
	0x72, 0x8f, 0x4b, 0x1e, 0xca, 0xa8, 0x90, 0xd9, 0x01, 0x3c, 0x78, 0x27, 0x6b, 0xd0, 0x23, 0xd0,
	0x59, 0x3d, 0x6f, 0x55, 0x71, 0x4f, 0x3f, 0xd1, 0x63, 0x10, 0x58, 0xbc, 0xb6, 0xe9, 0x8a, 0x0f,
	0xf4, 0x1a, 0x00, 0xff, 0x0a, 0xbd, 0x68, 0x43, 0x3d, 0x12, 0x98, 0x3a, 0x4b, 0x35, 0xa6, 0x96,
	0x23, 0xc0, 0x9d, 0x04, 0xdc, 0x59, 0x27, 0xe0, 0x6e, 0x46, 0x6d, 0xff, 0xd6, 0xa0, 0xf9, 0x31,
	0x43, 0xa0, 0x18, 0x6a, 0x42, 0x4d, 0x32, 0xf2, 0xb1, 0x15, 0x37, 0xf9, 0x3c, 0xe3, 0xe8, 0xe5,
	0x38, 0x95, 0xab, 0x70, 0x5e, 0x00, 0x9a, 0x63, 0x9a, 0xbc, 0x80, 0x8b, 0x7f, 0x1c, 0x99, 0xa8,
	0xc8, 0x64, 0xbf, 0x05, 0x23, 0xa7, 0x8b, 0x43, 0x12, 0xc4, 0x98, 0x3d, 0xb2, 0x04, 0xd2, 0xf8,
	0xd4, 0xb6, 0x23, 0x1c, 0x4d, 0x75, 0x22, 0x6b, 0xcf, 0xa0, 0xcb, 0xaa, 0xb3, 0x6b, 0x97, 0x4e,
	0x2a, 0xdf, 0x9e, 0x75, 0xe9, 0x15, 0xba, 0x48, 0x8e, 0x71, 0x9e, 0xc3, 0x90, 0x1c, 0x39, 0xad,
	0x64, 0x79, 0x03, 0x68, 0x79, 0x2c, 0x6c, 0x7c, 0xc3, 0x45, 0x3a, 0x60, 0xe4, 0x8a, 0xc5, 0x78,
	0xfb, 0x03, 0x74, 0x59, 0x58, 0xb5, 0xdf, 0x15, 0x60, 0x7d, 0xe8, 0x15, 0x9a, 0xc8, 0xfe, 0x63,
	0xe8, 0xcc, 0xb0, 0x8f, 0x29, 0xfe, 0xbf, 0x51, 0x26, 0x74, 0x2f, 0xa5, 0xb2, 0xc9, 0x1c, 0xfa,
	0x22, 0x73, 0x5b, 0x1f, 0x06, 0x60, 0xa9, 0x1a, 0x89, 0x31, 0xd3, 0x3f, 0x3a, 0x00, 0x8f, 0xac,
	0x28, 0x89, 0x30, 0x9a, 0x41, 0x23, 0x73, 0x38, 0xa8, 0x2f, 0x1f, 0xa0, 0x78, 0x74, 0x96, 0xa5,
	0x4a, 0x49, 0x7f, 0x97, 0xd0, 0xbe, 0xb0, 0x1e, 0x0d, 0xcf, 0x72, 0xc5, 0x42, 0xd6, 0x93, 0xb2,
	0xb4, 0xec, 0xc8, 0xb8, 0x32, 0x4e, 0xa6, 0x5c, 0xc5, 0xd3, 0x48, 0xb9, 0x14, 0xc6, 0x9f, 0xb8,
	0x2e, 0x3c, 0x4b, 0xb9, 0xd4, 0x07, 0x91, 0x72, 0x95, 0x58, 0x8d, 0x16, 0xd0, 0xca, 0xfb, 0x87,
	0x06, 0xb2, 0x42, 0x79, 0x01, 0xd6, 0xb0, 0x24, 0x2b, 0xdb, 0x7d, 0x06, 0x54, 0xf4, 0x0a, 0x8d,
	0x72, 0x45, 0x2a, 0xcc, 0x67, 0xff, 0x50, 0x48, 0xa3, 0x77, 0x60, 0x64, 0xe3, 0x2b, 0x1c, 0xfd,
	0xf4, 0xbe, 0xdd, 0x81, 0x55, 0xef, 0xeb, 0x5f, 0x6a, 0x5c, 0x10, 0x6e, 0xb7, 0xf7, 0xf9, 0xdf,
	0xd9, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x77, 0xe0, 0xcb, 0x3a, 0x06, 0x00, 0x00,
}
