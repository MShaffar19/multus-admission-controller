// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	keyspb "github.com/google/trillian/crypto/keyspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LogBackend struct {
	// name defines the name of the log backend for use in LogConfig messages and must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// backend_spec defines the RPC endpoint that clients should use to send requests
	// to this log backend. These should be in the same format as rpcBackendFlag in the
	// CTFE main and must not be an empty string.
	BackendSpec          string   `protobuf:"bytes,2,opt,name=backend_spec,json=backendSpec,proto3" json:"backend_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogBackend) Reset()         { *m = LogBackend{} }
func (m *LogBackend) String() string { return proto.CompactTextString(m) }
func (*LogBackend) ProtoMessage()    {}
func (*LogBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *LogBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogBackend.Unmarshal(m, b)
}
func (m *LogBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogBackend.Marshal(b, m, deterministic)
}
func (m *LogBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogBackend.Merge(m, src)
}
func (m *LogBackend) XXX_Size() int {
	return xxx_messageInfo_LogBackend.Size(m)
}
func (m *LogBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_LogBackend.DiscardUnknown(m)
}

var xxx_messageInfo_LogBackend proto.InternalMessageInfo

func (m *LogBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogBackend) GetBackendSpec() string {
	if m != nil {
		return m.BackendSpec
	}
	return ""
}

// LogBackendSet supports a configuration where a single set of frontends handle
// requests for multiple backends. For example this could be used to run different
// backends in different geographic regions.
type LogBackendSet struct {
	Backend              []*LogBackend `protobuf:"bytes,1,rep,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogBackendSet) Reset()         { *m = LogBackendSet{} }
func (m *LogBackendSet) String() string { return proto.CompactTextString(m) }
func (*LogBackendSet) ProtoMessage()    {}
func (*LogBackendSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *LogBackendSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogBackendSet.Unmarshal(m, b)
}
func (m *LogBackendSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogBackendSet.Marshal(b, m, deterministic)
}
func (m *LogBackendSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogBackendSet.Merge(m, src)
}
func (m *LogBackendSet) XXX_Size() int {
	return xxx_messageInfo_LogBackendSet.Size(m)
}
func (m *LogBackendSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogBackendSet.DiscardUnknown(m)
}

var xxx_messageInfo_LogBackendSet proto.InternalMessageInfo

func (m *LogBackendSet) GetBackend() []*LogBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

// LogConfigSet is a set of LogConfig messages.
type LogConfigSet struct {
	Config               []*LogConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogConfigSet) Reset()         { *m = LogConfigSet{} }
func (m *LogConfigSet) String() string { return proto.CompactTextString(m) }
func (*LogConfigSet) ProtoMessage()    {}
func (*LogConfigSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *LogConfigSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfigSet.Unmarshal(m, b)
}
func (m *LogConfigSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfigSet.Marshal(b, m, deterministic)
}
func (m *LogConfigSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfigSet.Merge(m, src)
}
func (m *LogConfigSet) XXX_Size() int {
	return xxx_messageInfo_LogConfigSet.Size(m)
}
func (m *LogConfigSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfigSet.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfigSet proto.InternalMessageInfo

func (m *LogConfigSet) GetConfig() []*LogConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// LogConfig describes the configuration options for a log instance.
//
// NEXT_ID: 18
type LogConfig struct {
	// The ID of a Trillian tree that stores the log data. The tree type must be
	// LOG for regular CT logs. For mirror logs it must be either PREORDERED_LOG
	// or LOG, and can change at runtime. CTFE in mirror mode uses only read API
	// which is common for both types.
	LogId int64 `protobuf:"varint,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// prefix is the name of the log. It will come after the global or
	// override handler prefix. For example if the handler prefix is "/logs"
	// and prefix is "vogon" the get-sth handler for this log will be
	// available at "/logs/vogon/ct/v1/get-sth". The prefix cannot be empty
	// and must not include "/" path separator characters.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// override_handler_prefix if set to a non empty value overrides the global
	// handler prefix for an individual log. For example this field is set to
	// "/otherlogs" then a log with prefix "vogon" will make it's get-sth handler
	// available at "/otherlogs/vogon/ct/v1/get-sth" regardless of what the
	// global prefix is. Can be set to '/' to make the get-sth handler register
	// at "/vogon/ct/v1/get-sth".
	OverrideHandlerPrefix string `protobuf:"bytes,13,opt,name=override_handler_prefix,json=overrideHandlerPrefix,proto3" json:"override_handler_prefix,omitempty"`
	// Paths to the files containing root certificates that are acceptable to the
	// log. The certs are served through get-roots endpoint. Optional in mirrors.
	RootsPemFile []string `protobuf:"bytes,3,rep,name=roots_pem_file,json=rootsPemFile,proto3" json:"roots_pem_file,omitempty"`
	// The private key used for signing STHs etc. Not required for mirrors.
	PrivateKey *any.Any `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The public key matching the above private key (if both are present). It is
	// used only by mirror logs for verifying the source log's signatures, but can
	// be specified for regular logs as well for the convenience of test tools.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// If reject_expired is true then the certificate validity period will be
	// checked against the current time during the validation of submissions.
	// This will cause expired and not-yet-valid certificates to be rejected.
	RejectExpired bool `protobuf:"varint,6,opt,name=reject_expired,json=rejectExpired,proto3" json:"reject_expired,omitempty"`
	// If reject_unexpired is true then CTFE rejects certificates that are within
	// their validity period with respect to the current time.
	RejectUnexpired bool `protobuf:"varint,17,opt,name=reject_unexpired,json=rejectUnexpired,proto3" json:"reject_unexpired,omitempty"`
	// If set, ext_key_usages will restrict the set of such usages that the
	// server will accept. By default all are accepted. The values specified
	// must be ones known to the x509 package.
	ExtKeyUsages []string `protobuf:"bytes,7,rep,name=ext_key_usages,json=extKeyUsages,proto3" json:"ext_key_usages,omitempty"`
	// not_after_start defines the start of the range of acceptable NotAfter
	// values, inclusive.
	// Leaving this unset implies no lower bound to the range.
	NotAfterStart *timestamp.Timestamp `protobuf:"bytes,8,opt,name=not_after_start,json=notAfterStart,proto3" json:"not_after_start,omitempty"`
	// not_after_limit defines the end of the range of acceptable NotAfter values,
	// exclusive.
	// Leaving this unset implies no upper bound to the range.
	NotAfterLimit *timestamp.Timestamp `protobuf:"bytes,9,opt,name=not_after_limit,json=notAfterLimit,proto3" json:"not_after_limit,omitempty"`
	// accept_only_ca controls whether or not *only* certificates with the CA bit
	// set will be accepted.
	AcceptOnlyCa bool `protobuf:"varint,10,opt,name=accept_only_ca,json=acceptOnlyCa,proto3" json:"accept_only_ca,omitempty"`
	// backend_name if set indicates which backend serves this log. The name must be
	// one of those defined in the LogBackendSet.
	LogBackendName string `protobuf:"bytes,11,opt,name=log_backend_name,json=logBackendName,proto3" json:"log_backend_name,omitempty"`
	// If set, the log is a mirror, i.e. it serves the data of another (source)
	// log. It doesn't handle write requests (add-chain, etc.), so it's not a
	// fully fledged RFC-6962 log, but the tree read requests like get-entries and
	// get-consistency-proof are compatible. A mirror doesn't have the source
	// log's key and can't sign STHs. Consequently, the log operator must ensure
	// to channel source log's STHs into CTFE.
	IsMirror bool `protobuf:"varint,12,opt,name=is_mirror,json=isMirror,proto3" json:"is_mirror,omitempty"`
	// The Maximum Merge Delay (MMD) of this log in seconds. See RFC6962 section 3
	// for definition of MMD. If zero, the log does not provide an MMD guarantee
	// (for example, it is a frozen log).
	MaxMergeDelaySec int32 `protobuf:"varint,14,opt,name=max_merge_delay_sec,json=maxMergeDelaySec,proto3" json:"max_merge_delay_sec,omitempty"`
	// The merge delay that the underlying log implementation is able/targeting to
	// provide. This option is exposed in CTFE metrics, and can be particularly
	// useful to catch when the log is behind but has not yet violated the strict
	// MMD limit.
	// Log operator should decide what exactly EMD means for them. For example, it
	// can be a 99-th percentile of merge delays that they observe, and they can
	// alert on the actual merge delay going above a certain multiple of this EMD.
	ExpectedMergeDelaySec int32 `protobuf:"varint,15,opt,name=expected_merge_delay_sec,json=expectedMergeDelaySec,proto3" json:"expected_merge_delay_sec,omitempty"`
	// The STH that this log will serve permanently (if present). Frozen STH must
	// be signed by this log's private key, and will be verified using the public
	// key specified in this config.
	FrozenSth            *SignedTreeHead `protobuf:"bytes,16,opt,name=frozen_sth,json=frozenSth,proto3" json:"frozen_sth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LogConfig) Reset()         { *m = LogConfig{} }
func (m *LogConfig) String() string { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()    {}
func (*LogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *LogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfig.Unmarshal(m, b)
}
func (m *LogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfig.Marshal(b, m, deterministic)
}
func (m *LogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfig.Merge(m, src)
}
func (m *LogConfig) XXX_Size() int {
	return xxx_messageInfo_LogConfig.Size(m)
}
func (m *LogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfig proto.InternalMessageInfo

func (m *LogConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *LogConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *LogConfig) GetOverrideHandlerPrefix() string {
	if m != nil {
		return m.OverrideHandlerPrefix
	}
	return ""
}

func (m *LogConfig) GetRootsPemFile() []string {
	if m != nil {
		return m.RootsPemFile
	}
	return nil
}

func (m *LogConfig) GetPrivateKey() *any.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *LogConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *LogConfig) GetRejectExpired() bool {
	if m != nil {
		return m.RejectExpired
	}
	return false
}

func (m *LogConfig) GetRejectUnexpired() bool {
	if m != nil {
		return m.RejectUnexpired
	}
	return false
}

func (m *LogConfig) GetExtKeyUsages() []string {
	if m != nil {
		return m.ExtKeyUsages
	}
	return nil
}

func (m *LogConfig) GetNotAfterStart() *timestamp.Timestamp {
	if m != nil {
		return m.NotAfterStart
	}
	return nil
}

func (m *LogConfig) GetNotAfterLimit() *timestamp.Timestamp {
	if m != nil {
		return m.NotAfterLimit
	}
	return nil
}

func (m *LogConfig) GetAcceptOnlyCa() bool {
	if m != nil {
		return m.AcceptOnlyCa
	}
	return false
}

func (m *LogConfig) GetLogBackendName() string {
	if m != nil {
		return m.LogBackendName
	}
	return ""
}

func (m *LogConfig) GetIsMirror() bool {
	if m != nil {
		return m.IsMirror
	}
	return false
}

func (m *LogConfig) GetMaxMergeDelaySec() int32 {
	if m != nil {
		return m.MaxMergeDelaySec
	}
	return 0
}

func (m *LogConfig) GetExpectedMergeDelaySec() int32 {
	if m != nil {
		return m.ExpectedMergeDelaySec
	}
	return 0
}

func (m *LogConfig) GetFrozenSth() *SignedTreeHead {
	if m != nil {
		return m.FrozenSth
	}
	return nil
}

// LogMultiConfig wraps up a LogBackendSet and corresponding LogConfigSet so
// that they can easily be parsed as a single proto.
type LogMultiConfig struct {
	// The set of backends that this configuration will use to send requests to.
	// The names of the backends in the LogBackendSet must all be distinct.
	Backends *LogBackendSet `protobuf:"bytes,1,opt,name=backends,proto3" json:"backends,omitempty"`
	// The set of logs that will use the above backends. All the protos in this
	// LogConfigSet must set a valid log_backend_name for the config to be usable.
	LogConfigs           *LogConfigSet `protobuf:"bytes,2,opt,name=log_configs,json=logConfigs,proto3" json:"log_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogMultiConfig) Reset()         { *m = LogMultiConfig{} }
func (m *LogMultiConfig) String() string { return proto.CompactTextString(m) }
func (*LogMultiConfig) ProtoMessage()    {}
func (*LogMultiConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}

func (m *LogMultiConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMultiConfig.Unmarshal(m, b)
}
func (m *LogMultiConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMultiConfig.Marshal(b, m, deterministic)
}
func (m *LogMultiConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMultiConfig.Merge(m, src)
}
func (m *LogMultiConfig) XXX_Size() int {
	return xxx_messageInfo_LogMultiConfig.Size(m)
}
func (m *LogMultiConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMultiConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogMultiConfig proto.InternalMessageInfo

func (m *LogMultiConfig) GetBackends() *LogBackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

func (m *LogMultiConfig) GetLogConfigs() *LogConfigSet {
	if m != nil {
		return m.LogConfigs
	}
	return nil
}

// SignedTreeHead represents the structure returned by the get-sth CT method.
// See RFC6962 sections 3.5 and 4.3 for reference.
// TODO(pavelkalinnikov): Find a better place for this type.
type SignedTreeHead struct {
	TreeSize             int64    `protobuf:"varint,1,opt,name=tree_size,json=treeSize,proto3" json:"tree_size,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sha256RootHash       []byte   `protobuf:"bytes,3,opt,name=sha256_root_hash,json=sha256RootHash,proto3" json:"sha256_root_hash,omitempty"`
	TreeHeadSignature    []byte   `protobuf:"bytes,4,opt,name=tree_head_signature,json=treeHeadSignature,proto3" json:"tree_head_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedTreeHead) Reset()         { *m = SignedTreeHead{} }
func (m *SignedTreeHead) String() string { return proto.CompactTextString(m) }
func (*SignedTreeHead) ProtoMessage()    {}
func (*SignedTreeHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{5}
}

func (m *SignedTreeHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTreeHead.Unmarshal(m, b)
}
func (m *SignedTreeHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTreeHead.Marshal(b, m, deterministic)
}
func (m *SignedTreeHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTreeHead.Merge(m, src)
}
func (m *SignedTreeHead) XXX_Size() int {
	return xxx_messageInfo_SignedTreeHead.Size(m)
}
func (m *SignedTreeHead) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTreeHead.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTreeHead proto.InternalMessageInfo

func (m *SignedTreeHead) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *SignedTreeHead) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SignedTreeHead) GetSha256RootHash() []byte {
	if m != nil {
		return m.Sha256RootHash
	}
	return nil
}

func (m *SignedTreeHead) GetTreeHeadSignature() []byte {
	if m != nil {
		return m.TreeHeadSignature
	}
	return nil
}

func init() {
	proto.RegisterType((*LogBackend)(nil), "configpb.LogBackend")
	proto.RegisterType((*LogBackendSet)(nil), "configpb.LogBackendSet")
	proto.RegisterType((*LogConfigSet)(nil), "configpb.LogConfigSet")
	proto.RegisterType((*LogConfig)(nil), "configpb.LogConfig")
	proto.RegisterType((*LogMultiConfig)(nil), "configpb.LogMultiConfig")
	proto.RegisterType((*SignedTreeHead)(nil), "configpb.SignedTreeHead")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x6f, 0xdb, 0x36,
	0x10, 0x87, 0xe7, 0xc6, 0xb5, 0xcf, 0x8e, 0x93, 0x30, 0x4b, 0xab, 0x65, 0x03, 0xe6, 0x19, 0x1b,
	0xe0, 0x61, 0x98, 0x3c, 0xa4, 0x48, 0xfb, 0xb0, 0x87, 0xa1, 0xcd, 0x36, 0x64, 0x48, 0xb2, 0x05,
	0x52, 0xfb, 0x4c, 0xd0, 0xd2, 0x59, 0xe2, 0x42, 0x89, 0x02, 0x49, 0x17, 0x56, 0x1e, 0xf6, 0x15,
	0xf6, 0x29, 0xf6, 0x3d, 0x07, 0xfe, 0x51, 0x82, 0xb4, 0x79, 0xe8, 0x93, 0xc5, 0xdf, 0x9f, 0xe3,
	0x1d, 0xef, 0xce, 0x30, 0xc9, 0x64, 0xbd, 0xe6, 0x45, 0xdc, 0x28, 0x69, 0x24, 0x19, 0xfa, 0x53,
	0xb3, 0x3a, 0x3e, 0x2d, 0xb8, 0x29, 0x37, 0xab, 0x38, 0x93, 0xd5, 0xb2, 0x90, 0xb2, 0x10, 0xb8,
	0x34, 0x8a, 0x0b, 0xc1, 0x59, 0xbd, 0xcc, 0x54, 0xdb, 0x18, 0xb9, 0xbc, 0xc1, 0x56, 0x37, 0xab,
	0xf0, 0xe3, 0x03, 0x1c, 0x7f, 0x11, 0xb4, 0xee, 0xb4, 0xda, 0xac, 0x97, 0xac, 0x6e, 0x03, 0xf5,
	0xf5, 0x87, 0x94, 0xe1, 0x15, 0x6a, 0xc3, 0xaa, 0xc6, 0x0b, 0xe6, 0x67, 0x00, 0x97, 0xb2, 0x78,
	0xc3, 0xb2, 0x1b, 0xac, 0x73, 0x42, 0xe0, 0x49, 0xcd, 0x2a, 0x8c, 0x7a, 0xb3, 0xde, 0x62, 0x94,
	0xb8, 0x6f, 0xf2, 0x0d, 0x4c, 0x56, 0x9e, 0xa6, 0xba, 0xc1, 0x2c, 0xfa, 0xcc, 0x71, 0xe3, 0x80,
	0xa5, 0x0d, 0x66, 0xf3, 0x5f, 0x60, 0xf7, 0x3e, 0x48, 0x8a, 0x86, 0xc4, 0xf0, 0x34, 0xf0, 0x51,
	0x6f, 0xd6, 0x5f, 0x8c, 0x4f, 0x3e, 0x8f, 0xbb, 0x22, 0xe3, 0x7b, 0x65, 0xd2, 0x89, 0xe6, 0x3f,
	0xc3, 0xe4, 0x52, 0x16, 0x67, 0x4e, 0x62, 0xfd, 0x3f, 0xc0, 0xc0, 0xeb, 0x83, 0xfd, 0xf0, 0x81,
	0xdd, 0xeb, 0x92, 0x20, 0x99, 0xff, 0x3b, 0x80, 0xd1, 0x1d, 0x4a, 0x8e, 0x60, 0x20, 0x64, 0x41,
	0x79, 0xee, 0x8a, 0xe8, 0x27, 0x3b, 0x42, 0x16, 0x7f, 0xe4, 0xe4, 0x19, 0x0c, 0x1a, 0x85, 0x6b,
	0xbe, 0x0d, 0xf9, 0x87, 0x13, 0x79, 0x09, 0xcf, 0xe5, 0x7b, 0x54, 0x8a, 0xe7, 0x48, 0x4b, 0x56,
	0xe7, 0x02, 0x15, 0x0d, 0xc2, 0x5d, 0x27, 0x3c, 0xea, 0xe8, 0x73, 0xcf, 0x5e, 0x7b, 0xdf, 0xb7,
	0x30, 0x55, 0x52, 0x1a, 0x4d, 0x1b, 0xac, 0xe8, 0x9a, 0x0b, 0x8c, 0xfa, 0xb3, 0xfe, 0x62, 0x94,
	0x4c, 0x1c, 0x7a, 0x8d, 0xd5, 0xef, 0x5c, 0x20, 0x39, 0x85, 0x71, 0xa3, 0xf8, 0x7b, 0x66, 0x90,
	0xde, 0x60, 0x1b, 0x3d, 0x99, 0xf5, 0xdc, 0x5b, 0xf8, 0xa6, 0xc4, 0x5d, 0x53, 0xe2, 0xd7, 0x75,
	0x9b, 0x40, 0x10, 0x5e, 0x60, 0x4b, 0x7e, 0x02, 0x68, 0x36, 0x2b, 0xc1, 0x33, 0xe7, 0xda, 0x71,
	0xae, 0x83, 0x38, 0xf4, 0xfc, 0xda, 0x31, 0x17, 0xd8, 0x26, 0xa3, 0xa6, 0xfb, 0x24, 0xdf, 0xc1,
	0x54, 0xe1, 0xdf, 0x98, 0x19, 0x8a, 0xdb, 0x86, 0x2b, 0xcc, 0xa3, 0xc1, 0xac, 0xb7, 0x18, 0x26,
	0xbb, 0x1e, 0xfd, 0xcd, 0x83, 0xe4, 0x7b, 0xd8, 0x0f, 0xb2, 0x4d, 0xdd, 0x09, 0x0f, 0x9c, 0x70,
	0xcf, 0xe3, 0xef, 0x3a, 0xd8, 0x16, 0x88, 0x5b, 0x63, 0x13, 0xa0, 0x1b, 0xcd, 0x0a, 0xd4, 0xd1,
	0x53, 0x5f, 0x20, 0x6e, 0xcd, 0x05, 0xb6, 0xef, 0x1c, 0x46, 0xde, 0xc0, 0x5e, 0x2d, 0x0d, 0x65,
	0x6b, 0x83, 0x8a, 0x6a, 0xc3, 0x94, 0x89, 0x86, 0x2e, 0xdd, 0xe3, 0x8f, 0x8a, 0x7c, 0xdb, 0x4d,
	0x5e, 0xb2, 0x5b, 0x4b, 0xf3, 0xda, 0x3a, 0x52, 0x6b, 0x78, 0x18, 0x43, 0xf0, 0x8a, 0x9b, 0x68,
	0xf4, 0xe9, 0x31, 0x2e, 0xad, 0xc1, 0x66, 0xcb, 0xb2, 0x0c, 0x1b, 0x43, 0x65, 0x2d, 0x5a, 0x9a,
	0xb1, 0x08, 0x5c, 0x59, 0x13, 0x8f, 0xfe, 0x55, 0x8b, 0xf6, 0x8c, 0x91, 0x05, 0xec, 0xdb, 0xd9,
	0xe8, 0xc6, 0xd9, 0x8d, 0xfa, 0xd8, 0x75, 0x79, 0x2a, 0xee, 0xa6, 0xf2, 0x4f, 0x3b, 0xf4, 0x5f,
	0xc2, 0x88, 0x6b, 0x5a, 0x71, 0xa5, 0xa4, 0x8a, 0x26, 0x2e, 0xd4, 0x90, 0xeb, 0x2b, 0x77, 0x26,
	0x3f, 0xc2, 0x61, 0xc5, 0xb6, 0xb4, 0x42, 0x55, 0x20, 0xcd, 0x51, 0xb0, 0x96, 0x6a, 0xcc, 0xa2,
	0xe9, 0xac, 0xb7, 0xd8, 0x49, 0xf6, 0x2b, 0xb6, 0xbd, 0xb2, 0xcc, 0xaf, 0x96, 0x48, 0x31, 0x23,
	0xaf, 0x20, 0xc2, 0x6d, 0x83, 0x99, 0xc1, 0xfc, 0x23, 0xcf, 0x9e, 0xf3, 0x1c, 0x75, 0xfc, 0x87,
	0x46, 0x58, 0x2b, 0x79, 0x8b, 0x35, 0xd5, 0xa6, 0x8c, 0xf6, 0xdd, 0x9b, 0x44, 0xf7, 0x9b, 0x90,
	0xf2, 0xa2, 0xc6, 0xfc, 0xad, 0x42, 0x3c, 0x47, 0x96, 0x27, 0x23, 0xaf, 0x4d, 0x4d, 0x39, 0xff,
	0x07, 0xa6, 0x97, 0xb2, 0xb8, 0xda, 0x08, 0xc3, 0xc3, 0x56, 0xbc, 0x80, 0x61, 0xa8, 0x5a, 0xbb,
	0xbd, 0x18, 0x9f, 0x3c, 0x7f, 0x6c, 0x23, 0x53, 0x34, 0xc9, 0x9d, 0x90, 0xbc, 0x82, 0xb1, 0x7d,
	0x2e, 0xaf, 0xd3, 0x6e, 0x71, 0xc6, 0x27, 0xcf, 0x1e, 0x59, 0x45, 0x6b, 0x03, 0xd1, 0x9d, 0xf4,
	0xfc, 0xbf, 0x1e, 0x4c, 0x1f, 0x66, 0x67, 0x1f, 0xd4, 0x28, 0x44, 0xaa, 0xf9, 0x2d, 0x86, 0xcd,
	0x1c, 0x5a, 0x20, 0xe5, 0xb7, 0x48, 0xbe, 0x82, 0xd1, 0xdd, 0xff, 0x92, 0xbb, 0xa6, 0x9f, 0xdc,
	0x03, 0xb6, 0x6b, 0xba, 0x64, 0x27, 0xa7, 0x2f, 0xa9, 0xdd, 0x2d, 0x5a, 0x32, 0x5d, 0x46, 0xfd,
	0x59, 0x6f, 0x31, 0x49, 0xa6, 0x1e, 0x4f, 0xa4, 0x34, 0xe7, 0x4c, 0x97, 0x24, 0x86, 0x43, 0x77,
	0x49, 0x89, 0x2c, 0xa7, 0x9a, 0x17, 0x35, 0x33, 0x1b, 0x85, 0x6e, 0xed, 0x26, 0xc9, 0x81, 0x09,
	0xb9, 0xa4, 0x1d, 0xb1, 0x1a, 0xb8, 0xc1, 0x7a, 0xf1, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x6f, 0xb2, 0x4a, 0x90, 0x05, 0x00, 0x00,
}
