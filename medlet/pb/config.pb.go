// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package medletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	GlobalConfig
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	PprofConfig
	MiscConfig
	StatsConfig
	InfluxdbConfig
	SyncConfig
*/
package medletpb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reporting modules.
type StatsConfig_ReportingModule int32

const (
	StatsConfig_Influxdb StatsConfig_ReportingModule = 0
)

var StatsConfig_ReportingModule_name = map[int32]string{
	0: "Influxdb",
}
var StatsConfig_ReportingModule_value = map[string]int32{
	"Influxdb": 0,
}

func (x StatsConfig_ReportingModule) String() string {
	return proto.EnumName(StatsConfig_ReportingModule_name, int32(x))
}
func (StatsConfig_ReportingModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{8, 0}
}

// Med global configurations.
type Config struct {
	// Global config
	Global *GlobalConfig `protobuf:"bytes,1,opt,name=global" json:"global,omitempty"`
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,2,opt,name=network" json:"network,omitempty"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,3,opt,name=chain" json:"chain,omitempty"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,4,opt,name=rpc" json:"rpc,omitempty"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats,omitempty"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc,omitempty"`
	// App Config.
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app,omitempty"`
	// Sync Service Config.
	Sync *SyncConfig `protobuf:"bytes,200,opt,name=sync" json:"sync,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetGlobal() *GlobalConfig {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetStats() *StatsConfig {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Config) GetSync() *SyncConfig {
	if m != nil {
		return m.Sync
	}
	return nil
}

type GlobalConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir,omitempty"`
}

func (m *GlobalConfig) Reset()                    { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string            { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()               {}
func (*GlobalConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *GlobalConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *GlobalConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

type NetworkConfig struct {
	// Med seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed,omitempty"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen,omitempty"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Network ID
	NetworkId uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Rount table syncing interval - Millisecond
	RouteTableSyncLoopInterval uint32 `protobuf:"varint,5,opt,name=route_table_sync_loop_interval,json=routeTableSyncLoopInterval,proto3" json:"route_table_sync_loop_interval,omitempty"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *NetworkConfig) GetRouteTableSyncLoopInterval() uint32 {
	if m != nil {
		return m.RouteTableSyncLoopInterval
	}
	return 0
}

type ChainConfig struct {
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir,omitempty"`
	// start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine,omitempty"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner,omitempty"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers []string `protobuf:"bytes,24,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers,omitempty"`
	// Block cache size
	BlockCacheSize uint32 `protobuf:"varint,25,opt,name=block_cache_size,json=blockCacheSize,proto3" json:"block_cache_size,omitempty"`
	// Tail cache size
	TailCacheSize uint32 `protobuf:"varint,26,opt,name=tail_cache_size,json=tailCacheSize,proto3" json:"tail_cache_size,omitempty"`
	// Blockpool size
	BlockPoolSize uint32 `protobuf:"varint,27,opt,name=block_pool_size,json=blockPoolSize,proto3" json:"block_pool_size,omitempty"`
	// Transaction pool size
	TransactionPoolSize uint32 `protobuf:"varint,28,opt,name=transaction_pool_size,json=transactionPoolSize,proto3" json:"transaction_pool_size,omitempty"`
	// TODO account manager
	// Miner private key.
	Privkey string `protobuf:"bytes,29,opt,name=privkey,proto3" json:"privkey,omitempty"`
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetStartMine() bool {
	if m != nil {
		return m.StartMine
	}
	return false
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

func (m *ChainConfig) GetBlockCacheSize() uint32 {
	if m != nil {
		return m.BlockCacheSize
	}
	return 0
}

func (m *ChainConfig) GetTailCacheSize() uint32 {
	if m != nil {
		return m.TailCacheSize
	}
	return 0
}

func (m *ChainConfig) GetBlockPoolSize() uint32 {
	if m != nil {
		return m.BlockPoolSize
	}
	return 0
}

func (m *ChainConfig) GetTransactionPoolSize() uint32 {
	if m != nil {
		return m.TransactionPoolSize
	}
	return 0
}

func (m *ChainConfig) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen,omitempty"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen,omitempty"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module,omitempty"`
	// Connection limit.
	ConnectionLimits int32 `protobuf:"varint,4,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

func (m *RPCConfig) GetConnectionLimits() int32 {
	if m != nil {
		return m.ConnectionLimits
	}
	return 0
}

type AppConfig struct {
	// log level
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// log file path
	LogFile string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	// log file age, unit is s.
	LogAge uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age,omitempty"`
	// pprof config
	Pprof *PprofConfig `protobuf:"bytes,4,opt,name=pprof" json:"pprof,omitempty"`
	// App version
	Version string `protobuf:"bytes,100,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func (m *AppConfig) GetPprof() *PprofConfig {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *AppConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen,omitempty"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile,omitempty"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile,omitempty"`
}

func (m *PprofConfig) Reset()                    { *m = PprofConfig{} }
func (m *PprofConfig) String() string            { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()               {}
func (*PprofConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *PprofConfig) GetHttpListen() string {
	if m != nil {
		return m.HttpListen
	}
	return ""
}

func (m *PprofConfig) GetCpuprofile() string {
	if m != nil {
		return m.Cpuprofile
	}
	return ""
}

func (m *PprofConfig) GetMemprofile() string {
	if m != nil {
		return m.Memprofile
	}
	return ""
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper,omitempty"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

type StatsConfig struct {
	// Enable metrics of not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=medletpb.StatsConfig_ReportingModule" json:"reporting_module,omitempty"`
	// Influxdb config.
	Influxdb    *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb,omitempty"`
	MetricsTags []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{8} }

func (m *StatsConfig) GetEnableMetrics() bool {
	if m != nil {
		return m.EnableMetrics
	}
	return false
}

func (m *StatsConfig) GetReportingModule() []StatsConfig_ReportingModule {
	if m != nil {
		return m.ReportingModule
	}
	return nil
}

func (m *StatsConfig) GetInfluxdb() *InfluxdbConfig {
	if m != nil {
		return m.Influxdb
	}
	return nil
}

func (m *StatsConfig) GetMetricsTags() []string {
	if m != nil {
		return m.MetricsTags
	}
	return nil
}

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Auth password.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *InfluxdbConfig) Reset()                    { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string            { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()               {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{9} }

func (m *InfluxdbConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InfluxdbConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InfluxdbConfig) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *InfluxdbConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InfluxdbConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SyncConfig struct {
	// Minimum ChunkSize on seeding
	SeedingMinChunkSize uint64 `protobuf:"varint,1,opt,name=seeding_min_chunk_size,json=seedingMinChunkSize,proto3" json:"seeding_min_chunk_size,omitempty"`
	// Maximum ChunkSize on seeding
	SeedingMaxChunkSize uint64 `protobuf:"varint,2,opt,name=seeding_max_chunk_size,json=seedingMaxChunkSize,proto3" json:"seeding_max_chunk_size,omitempty"`
	// Maximum Number of Concurrent Peers on seeding
	SeedingMaxConcurrentPeers uint32 `protobuf:"varint,3,opt,name=seeding_max_concurrent_peers,json=seedingMaxConcurrentPeers,proto3" json:"seeding_max_concurrent_peers,omitempty"`
	// Size of Chunk on Downloading
	DownloadChunkSize uint64 `protobuf:"varint,4,opt,name=download_chunk_size,json=downloadChunkSize,proto3" json:"download_chunk_size,omitempty"`
	// Maximum Number of Concurrent DownloadTasks
	DownloadMaxConcurrentTasks uint32 `protobuf:"varint,5,opt,name=download_max_concurrent_tasks,json=downloadMaxConcurrentTasks,proto3" json:"download_max_concurrent_tasks,omitempty"`
	// Chunk Cache Size
	DownloadChunkCacheSize uint64 `protobuf:"varint,6,opt,name=download_chunk_cache_size,json=downloadChunkCacheSize,proto3" json:"download_chunk_cache_size,omitempty"`
	// Minimum Number of Peers to pass majority check
	MinimumPeers uint32 `protobuf:"varint,7,opt,name=minimum_peers,json=minimumPeers,proto3" json:"minimum_peers,omitempty"`
	// Blockchunk request interval - second
	RequestInterval uint32 `protobuf:"varint,10,opt,name=request_interval,json=requestInterval,proto3" json:"request_interval,omitempty"`
	// Timeout for finishing download  - second
	FinisherTimeout uint32 `protobuf:"varint,11,opt,name=finisher_timeout,json=finisherTimeout,proto3" json:"finisher_timeout,omitempty"`
}

func (m *SyncConfig) Reset()                    { *m = SyncConfig{} }
func (m *SyncConfig) String() string            { return proto.CompactTextString(m) }
func (*SyncConfig) ProtoMessage()               {}
func (*SyncConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{10} }

func (m *SyncConfig) GetSeedingMinChunkSize() uint64 {
	if m != nil {
		return m.SeedingMinChunkSize
	}
	return 0
}

func (m *SyncConfig) GetSeedingMaxChunkSize() uint64 {
	if m != nil {
		return m.SeedingMaxChunkSize
	}
	return 0
}

func (m *SyncConfig) GetSeedingMaxConcurrentPeers() uint32 {
	if m != nil {
		return m.SeedingMaxConcurrentPeers
	}
	return 0
}

func (m *SyncConfig) GetDownloadChunkSize() uint64 {
	if m != nil {
		return m.DownloadChunkSize
	}
	return 0
}

func (m *SyncConfig) GetDownloadMaxConcurrentTasks() uint32 {
	if m != nil {
		return m.DownloadMaxConcurrentTasks
	}
	return 0
}

func (m *SyncConfig) GetDownloadChunkCacheSize() uint64 {
	if m != nil {
		return m.DownloadChunkCacheSize
	}
	return 0
}

func (m *SyncConfig) GetMinimumPeers() uint32 {
	if m != nil {
		return m.MinimumPeers
	}
	return 0
}

func (m *SyncConfig) GetRequestInterval() uint32 {
	if m != nil {
		return m.RequestInterval
	}
	return 0
}

func (m *SyncConfig) GetFinisherTimeout() uint32 {
	if m != nil {
		return m.FinisherTimeout
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "medletpb.Config")
	proto.RegisterType((*GlobalConfig)(nil), "medletpb.GlobalConfig")
	proto.RegisterType((*NetworkConfig)(nil), "medletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "medletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "medletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "medletpb.AppConfig")
	proto.RegisterType((*PprofConfig)(nil), "medletpb.PprofConfig")
	proto.RegisterType((*MiscConfig)(nil), "medletpb.MiscConfig")
	proto.RegisterType((*StatsConfig)(nil), "medletpb.StatsConfig")
	proto.RegisterType((*InfluxdbConfig)(nil), "medletpb.InfluxdbConfig")
	proto.RegisterType((*SyncConfig)(nil), "medletpb.SyncConfig")
	proto.RegisterEnum("medletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 1150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x96, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xc7, 0x2b, 0x59, 0xb6, 0xb5, 0x23, 0xc9, 0x76, 0xe8, 0xc4, 0x59, 0xe7, 0xbb, 0x2a, 0x52,
	0x38, 0x08, 0x60, 0xa0, 0x4e, 0x2f, 0x3d, 0x14, 0x45, 0x2a, 0xa0, 0x85, 0x11, 0xbb, 0x30, 0x36,
	0xbe, 0x2f, 0xa8, 0x5d, 0x6a, 0x45, 0x78, 0x45, 0xb2, 0x24, 0xe5, 0xd8, 0x39, 0xf5, 0x35, 0x7a,
	0x2e, 0xfa, 0x1c, 0xbd, 0xf6, 0x31, 0x7a, 0xec, 0x63, 0x14, 0x33, 0xcb, 0xd5, 0xae, 0x84, 0xde,
	0x96, 0xff, 0xf9, 0xcd, 0x0c, 0x39, 0x1c, 0x8e, 0x04, 0xc3, 0x4c, 0xab, 0x99, 0x2c, 0x4e, 0x8d,
	0xd5, 0x5e, 0xb3, 0xfe, 0x42, 0xe4, 0xa5, 0xf0, 0x66, 0x3a, 0xfe, 0xb7, 0x0b, 0x3b, 0x13, 0x32,
	0xb1, 0x53, 0xd8, 0x29, 0x4a, 0x3d, 0xe5, 0x65, 0xdc, 0x79, 0xd5, 0x39, 0x19, 0x9c, 0x1d, 0x9d,
	0xd6, 0xd4, 0xe9, 0xcf, 0xa4, 0x57, 0x5c, 0x12, 0x28, 0xf6, 0x0d, 0xec, 0x2a, 0xe1, 0x3f, 0x69,
	0x7b, 0x13, 0x77, 0xc9, 0xe1, 0x71, 0xe3, 0xf0, 0x4b, 0x65, 0x08, 0x1e, 0x35, 0xc7, 0xde, 0xc2,
	0x76, 0x36, 0xe7, 0x52, 0xc5, 0x5b, 0xe4, 0xf0, 0xa8, 0x71, 0x98, 0xa0, 0x1c, 0xf0, 0x8a, 0x61,
	0xaf, 0x61, 0xcb, 0x9a, 0x2c, 0xee, 0x11, 0x7a, 0xd8, 0xa0, 0xc9, 0xd5, 0x24, 0x80, 0x68, 0xc7,
	0x98, 0xce, 0x73, 0xef, 0xe2, 0x7c, 0x33, 0xe6, 0x47, 0x94, 0xeb, 0x98, 0xc4, 0xb0, 0x13, 0xe8,
	0x2d, 0xa4, 0xcb, 0x62, 0x41, 0xec, 0xc3, 0x86, 0xbd, 0x94, 0x2e, 0x0b, 0x28, 0x11, 0x98, 0x9d,
	0x1b, 0x13, 0xcf, 0x36, 0xb3, 0xbf, 0x37, 0xa6, 0xce, 0xce, 0x8d, 0x61, 0x6f, 0xa0, 0xe7, 0xee,
	0x55, 0x16, 0xff, 0xdd, 0xd9, 0x8c, 0xf8, 0xf1, 0x5e, 0xad, 0x22, 0x22, 0x32, 0x9e, 0xc0, 0xb0,
	0x5d, 0x47, 0x76, 0x0c, 0x7d, 0x3a, 0x68, 0x2a, 0x73, 0xaa, 0xf8, 0x28, 0xd9, 0xa5, 0xf5, 0x79,
	0xce, 0x62, 0xd8, 0xcd, 0xb9, 0xe7, 0xb9, 0xb4, 0xf1, 0xe0, 0x55, 0xe7, 0x24, 0x4a, 0xea, 0xe5,
	0xf8, 0xaf, 0x0e, 0x8c, 0xd6, 0x8a, 0xcb, 0x18, 0xf4, 0x9c, 0x10, 0x18, 0x62, 0xeb, 0x24, 0x4a,
	0xe8, 0x9b, 0x1d, 0xc1, 0x4e, 0x29, 0x9d, 0x17, 0x2a, 0xee, 0x92, 0x1a, 0x56, 0xec, 0x25, 0x0c,
	0x8c, 0x95, 0xb7, 0xdc, 0x8b, 0xf4, 0x46, 0xdc, 0xd3, 0x2d, 0x44, 0x09, 0x04, 0xe9, 0x83, 0xb8,
	0x67, 0xcf, 0x01, 0xc2, 0x5d, 0xe1, 0xae, 0x7a, 0xb4, 0xab, 0x28, 0x28, 0xe7, 0x39, 0xfb, 0x11,
	0x5e, 0x58, 0xbd, 0xf4, 0x22, 0xf5, 0x7c, 0x5a, 0x8a, 0x14, 0x8f, 0x95, 0x96, 0x5a, 0x9b, 0x54,
	0x2a, 0x2f, 0xec, 0x2d, 0x2f, 0xe3, 0x6d, 0x72, 0x79, 0x42, 0xd4, 0x35, 0x42, 0x58, 0x86, 0x0b,
	0xad, 0xcd, 0x79, 0x20, 0xc6, 0x7f, 0x6e, 0xc1, 0xa0, 0x75, 0xdb, 0x78, 0xd6, 0x42, 0x28, 0xe1,
	0xa4, 0xa3, 0x36, 0x8a, 0x92, 0x7a, 0x89, 0xa7, 0xb8, 0x11, 0xf7, 0x58, 0x84, 0x21, 0x19, 0xc2,
	0x0a, 0x37, 0xe9, 0x3c, 0xb7, 0x3e, 0x5d, 0x48, 0x25, 0xe2, 0x87, 0xaf, 0x3a, 0x27, 0xfd, 0x24,
	0x22, 0xe5, 0x52, 0x2a, 0xc1, 0x9e, 0x40, 0x3f, 0xd3, 0x52, 0x4d, 0xb9, 0x13, 0xf1, 0x23, 0x72,
	0x5c, 0xad, 0xd9, 0x43, 0xd8, 0x46, 0x27, 0x1b, 0x1f, 0x91, 0xa1, 0x5a, 0xb0, 0x17, 0x00, 0x86,
	0x3b, 0x67, 0xe6, 0x16, 0x7d, 0x1e, 0x87, 0xaa, 0xac, 0x14, 0xf6, 0x16, 0x1e, 0x38, 0x59, 0x28,
	0xee, 0x97, 0x56, 0xa4, 0x99, 0x34, 0x73, 0x61, 0x5d, 0x1c, 0x53, 0x65, 0x0f, 0x56, 0x86, 0x49,
	0xa5, 0xb3, 0x13, 0x38, 0x98, 0x96, 0x3a, 0xbb, 0x49, 0x33, 0x9e, 0xcd, 0x45, 0xea, 0xe4, 0x67,
	0x11, 0x1f, 0x53, 0x55, 0xf6, 0x48, 0x9f, 0xa0, 0xfc, 0x51, 0x7e, 0x16, 0xec, 0x6b, 0xd8, 0xf7,
	0x5c, 0x96, 0x6d, 0xf0, 0x09, 0x81, 0x23, 0x94, 0xd7, 0xb8, 0x2a, 0xa2, 0xd1, 0xba, 0xac, 0xb8,
	0xa7, 0x15, 0x47, 0xf2, 0x95, 0xd6, 0x25, 0x71, 0x67, 0xf0, 0xc8, 0x5b, 0xae, 0x1c, 0xcf, 0xbc,
	0xd4, 0xaa, 0x45, 0x3f, 0x23, 0xfa, 0xb0, 0x65, 0x5c, 0xf9, 0xc4, 0xb0, 0x8b, 0xd7, 0x8f, 0xdd,
	0xf0, 0xbc, 0xaa, 0x7e, 0x58, 0x8e, 0x7f, 0xef, 0x40, 0xb4, 0x7a, 0x6a, 0x58, 0x73, 0x6b, 0xb2,
	0x34, 0x74, 0x55, 0xd5, 0x6b, 0x91, 0x35, 0xd9, 0xc5, 0xaa, 0xb1, 0xe6, 0xde, 0x9b, 0x74, 0xad,
	0xeb, 0x00, 0xa5, 0x0d, 0x60, 0xa1, 0xf3, 0x65, 0x29, 0xe2, 0xad, 0x06, 0xb8, 0x24, 0x05, 0x6b,
	0x9c, 0x69, 0xa5, 0x44, 0xb5, 0xf7, 0x52, 0x2e, 0xa4, 0x77, 0xd4, 0x80, 0xdb, 0xc9, 0x41, 0x63,
	0xb8, 0x20, 0x7d, 0xfc, 0x47, 0x07, 0xa2, 0xd5, 0x43, 0x64, 0x4f, 0x21, 0x2a, 0x75, 0x91, 0x96,
	0xe2, 0x56, 0x54, 0xb3, 0x2b, 0x4a, 0xfa, 0xa5, 0x2e, 0x2e, 0x70, 0x8d, 0xaf, 0x0c, 0x8d, 0x33,
	0x59, 0x8a, 0xba, 0xbf, 0x4a, 0x5d, 0xfc, 0x24, 0x4b, 0xc1, 0x1e, 0x03, 0x7e, 0xa6, 0xbc, 0x10,
	0xf4, 0x12, 0x46, 0xc9, 0x4e, 0xa9, 0x8b, 0xf7, 0x05, 0xee, 0x65, 0xdb, 0x18, 0xab, 0x67, 0x61,
	0xf6, 0xb4, 0x46, 0xca, 0x15, 0xca, 0xf5, 0x48, 0x21, 0x06, 0x2b, 0x78, 0x2b, 0xac, 0x93, 0x5a,
	0xd1, 0x04, 0x8a, 0x92, 0x7a, 0x39, 0x56, 0x30, 0x68, 0xf1, 0x9b, 0x35, 0xaa, 0x36, 0xda, 0xae,
	0xd1, 0x0b, 0x80, 0xcc, 0x2c, 0xd1, 0xa3, 0xd9, 0x6c, 0x4b, 0x41, 0xfb, 0x42, 0x2c, 0x6a, 0x7b,
	0x78, 0xbc, 0x8d, 0x32, 0xfe, 0x00, 0xd0, 0x8c, 0x31, 0xf6, 0x3d, 0x3c, 0xcd, 0xc5, 0x8c, 0x2f,
	0x4b, 0x8f, 0x6f, 0xdd, 0x79, 0x6d, 0x05, 0x55, 0x01, 0x1b, 0x58, 0xd8, 0x90, 0x3e, 0x0e, 0xc8,
	0x87, 0x40, 0x60, 0x5d, 0x26, 0x68, 0x1f, 0xff, 0xd6, 0x85, 0x41, 0x6b, 0x80, 0xb2, 0xd7, 0xb0,
	0x27, 0x14, 0xbd, 0xfa, 0x85, 0xf0, 0x56, 0x66, 0x8e, 0x22, 0xf4, 0x93, 0x51, 0xa5, 0x5e, 0x56,
	0x22, 0xbb, 0x82, 0x03, 0x2b, 0x8c, 0xb6, 0x5e, 0xaa, 0xa2, 0xbe, 0x6c, 0xec, 0x86, 0xbd, 0xb3,
	0xd7, 0xff, 0x3b, 0x98, 0x4f, 0x93, 0x9a, 0xae, 0xfa, 0x20, 0xd9, 0xb7, 0xeb, 0x02, 0xfb, 0x16,
	0xfa, 0x52, 0xcd, 0xca, 0xe5, 0x5d, 0x3e, 0xa5, 0x61, 0x38, 0x38, 0x8b, 0x9b, 0x48, 0xe7, 0xc1,
	0x12, 0xae, 0x64, 0x45, 0xb2, 0x2f, 0x61, 0x18, 0xf6, 0x99, 0x7a, 0x5e, 0xb8, 0x78, 0x48, 0x0d,
	0x37, 0x08, 0xda, 0x35, 0x2f, 0xdc, 0xf8, 0x25, 0xec, 0x6f, 0x24, 0x67, 0x43, 0xe8, 0xd7, 0x11,
	0x0f, 0xbe, 0x18, 0xdf, 0xc1, 0xde, 0x7a, 0x7c, 0x9c, 0xb5, 0x73, 0xed, 0x7c, 0x28, 0x1e, 0x7d,
	0xa3, 0x86, 0x41, 0xe8, 0xbe, 0x46, 0x09, 0x7d, 0xb3, 0x3d, 0xe8, 0xe6, 0xd3, 0x70, 0x43, 0xdd,
	0x7c, 0x8a, 0xcc, 0xd2, 0x09, 0x4b, 0xfd, 0x14, 0x25, 0xf4, 0x8d, 0x63, 0x0a, 0x47, 0xcc, 0x27,
	0x6d, 0x73, 0x9a, 0x9a, 0x51, 0xb2, 0x5a, 0x8f, 0xff, 0xd9, 0x02, 0x68, 0x7e, 0x3f, 0xd8, 0x3b,
	0x38, 0xc2, 0xb1, 0x4e, 0x25, 0x95, 0x2a, 0xcd, 0xe6, 0x4b, 0x75, 0x53, 0xbd, 0x6c, 0xdc, 0x48,
	0x2f, 0x39, 0x0c, 0xd6, 0x4b, 0xa9, 0x26, 0x68, 0xa3, 0x97, 0xdd, 0x76, 0xe2, 0x77, 0x6d, 0xa7,
	0xee, 0xba, 0x13, 0xbf, 0x6b, 0x9c, 0x7e, 0x80, 0x67, 0x6b, 0x4e, 0x5a, 0x65, 0x4b, 0x6b, 0x85,
	0xf2, 0xa9, 0x11, 0x38, 0xf4, 0xaa, 0x77, 0x72, 0xdc, 0x72, 0x5d, 0x11, 0x57, 0x08, 0xb0, 0x53,
	0x38, 0xcc, 0xf5, 0x27, 0x55, 0x6a, 0x9e, 0xb7, 0x53, 0xf6, 0x28, 0xe5, 0x83, 0xda, 0xd4, 0x24,
	0x7c, 0x0f, 0xcf, 0x57, 0xfc, 0x46, 0x46, 0xcf, 0xdd, 0x8d, 0xab, 0x7f, 0x50, 0x6a, 0x68, 0x2d,
	0xe5, 0x35, 0x12, 0xec, 0x3b, 0x38, 0xde, 0x48, 0xd9, 0x1a, 0xa8, 0x3b, 0x94, 0xf8, 0x68, 0x2d,
	0x71, 0x33, 0x59, 0xbf, 0x82, 0xd1, 0x42, 0x2a, 0xb9, 0x58, 0x2e, 0xc2, 0xf9, 0x76, 0x29, 0xdb,
	0x30, 0x88, 0xd5, 0x91, 0xde, 0x60, 0x4b, 0xff, 0xba, 0x14, 0xce, 0x37, 0x3f, 0x73, 0x40, 0xdc,
	0x7e, 0xd0, 0xeb, 0xdf, 0x36, 0x44, 0x67, 0x52, 0x49, 0x37, 0x17, 0x36, 0xf5, 0x72, 0x21, 0xf4,
	0xd2, 0x53, 0xcf, 0x8e, 0x92, 0xfd, 0x5a, 0xbf, 0xae, 0xe4, 0xe9, 0x0e, 0xfd, 0x13, 0x7b, 0xf7,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x56, 0x96, 0xc7, 0x99, 0x09, 0x00, 0x00,
}
