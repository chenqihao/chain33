// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	account.proto
	blockchain.proto
	common.proto
	config.proto
	db.proto
	evmcontract.proto
	executor.proto
	executorTrade.proto
	game.proto
	p2p.proto
	pbft.proto
	relay.proto
	rpc.proto
	statistic.proto
	transaction.proto
	wallet.proto

It has these top-level messages:
	Account
	ReceiptExecAccountTransfer
	ReceiptAccountTransfer
	ReqBalance
	Accounts
	ReqTokenBalance
	ReqAccountTokenAssets
	TokenAsset
	ReplyAccountTokenAssets
	ExecAccount
	AllExecBalance
	Header
	Block
	Blocks
	BlockPid
	BlockDetails
	Headers
	HeadersPid
	BlockOverview
	BlockDetail
	Receipts
	PrivacyKV
	PrivacyKVToken
	ReceiptsAndPrivacyKV
	ReceiptCheckTxList
	ChainStatus
	ReqBlocks
	MempoolSize
	ReplyBlockHeight
	BlockBody
	IsCaughtUp
	IsNtpClockSync
	BlockChainQuery
	BlockSequence
	BlockSequences
	ParaChainBlockDetail
	Reply
	ReqString
	ReplyString
	ReplyStrings
	ReqInt
	Int64
	ReqHash
	ReplyHash
	ReqNil
	ReqHashes
	ReplyHashes
	KeyValue
	TxHash
	TimeStatus
	ReqKey
	Config
	Log
	MemPool
	Consensus
	Wallet
	Store
	BlockChain
	P2P
	Rpc
	Exec
	Authority
	Pprof
	LeafNode
	InnerNode
	MAVLProof
	StoreNode
	LocalDBSet
	LocalDBList
	LocalDBGet
	LocalReplyValue
	StoreSet
	StoreSetWithSync
	StoreGet
	StoreReplyValue
	EVMContractObject
	EVMContractData
	EVMContractState
	EVMContractAction
	ReceiptEVMContract
	EVMStateChangeItem
	EVMContractDataCmd
	EVMContractStateCmd
	ReceiptEVMContractCmd
	CheckEVMAddrReq
	CheckEVMAddrResp
	EstimateEVMGasReq
	EstimateEVMGasResp
	EvmDebugReq
	EvmDebugResp
	Genesis
	CoinsAction
	CoinsGenesis
	CoinsTransfer
	CoinsTransferToExec
	CoinsWithdraw
	Hashlock
	HashlockAction
	HashlockLock
	HashlockUnlock
	HashlockSend
	HashRecv
	Hashlockquery
	Ticket
	TicketAction
	TicketMiner
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ExecTxList
	Query
	Norm
	NormAction
	NormPut
	Cert
	CertAction
	CertPut
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PreRetrieve
	PerformRetrieve
	CancelRetrieve
	ReqRetrieveInfo
	RetrieveQuery
	TokenAction
	TokenPreCreate
	TokenFinishCreate
	TokenRevokeCreate
	Token
	ReqTokens
	ReplyTokens
	ReceiptToken
	TokenRecv
	ReplyAddrRecvForTokens
	ArrayConfig
	StringConfig
	Int32Config
	ConfigItem
	ModifyConfig
	ManageAction
	ReceiptConfig
	ReplyConfig
	HistoryCertStore
	PrivacyAction
	Public2Privacy
	Privacy2Privacy
	Privacy2Public
	UTXOGlobalIndex
	KeyInput
	PrivacyInput
	KeyOutput
	PrivacyOutput
	GroupUTXOGlobalIndex
	LocalUTXOItem
	ReqUTXOPubKeys
	PublicKeyData
	GroupUTXOPubKey
	ResUTXOPubKeys
	ReqPrivacyToken
	AmountDetail
	ReplyPrivacyAmounts
	ReplyUTXOsOfAmount
	ReceiptPrivacyOutput
	AmountsOfUTXO
	TokenNamesOfUTXO
	UTXOGlobalIndex4Print
	KeyInput4Print
	KeyOutput4Print
	PrivacyInput4Print
	PrivacyOutput4Print
	Public2Privacy4Print
	Privacy2Privacy4Print
	Privacy2Public4Print
	PrivacyAction4Print
	Trade
	TradeForSell
	TradeForBuy
	TradeForRevokeSell
	TradeForBuyLimit
	TradeForSellMarket
	TradeForRevokeBuy
	SellOrder
	BuyLimitOrder
	ReceiptBuyBase
	ReceiptSellBase
	ReceiptTradeBuyMarket
	ReceiptTradeBuyLimit
	ReceiptTradeBuyRevoke
	ReceiptTradeSell
	ReceiptSellMarket
	ReceiptTradeRevoke
	ReqAddrTokens
	ReqTokenSellOrder
	ReqTokenBuyOrder
	ReplyBuyOrder
	ReplySellOrder
	ReplySellOrders
	ReplyBuyOrders
	ReplyTradeOrder
	ReplyTradeOrders
	Game
	GameAction
	GameMatch
	GameCancel
	GameClose
	GameOpen
	GameList
	GameInfos
	ReplyGameList
	ReceiptGame
	P2PGetPeerInfo
	P2PPeerInfo
	P2PVersion
	P2PVerAck
	P2PPing
	P2PPong
	P2PGetAddr
	P2PAddr
	P2PAddrList
	P2PExternalInfo
	P2PGetBlocks
	P2PGetMempool
	P2PInv
	Inventory
	P2PGetData
	P2PTx
	P2PBlock
	Versions
	BroadCastData
	P2PGetHeaders
	P2PHeaders
	InvData
	InvDatas
	Peer
	PeerList
	NodeNetInfo
	PeersReply
	PeersInfo
	Operation
	Checkpoint
	Entry
	ViewChange
	Summary
	Result
	Request
	RequestClient
	RequestPrePrepare
	RequestPrepare
	RequestCommit
	RequestCheckpoint
	RequestViewChange
	RequestAck
	RequestNewView
	ClientReply
	RelayAction
	RelayCreate
	RelayOrder
	RelayAccept
	RelayRevoke
	RelayConfirmTx
	RelayVerify
	RelayVerifyCli
	BtcHeader
	BtcHeaders
	BtcTransaction
	Vin
	Vout
	BtcSpv
	RelayLastRcvBtcHeader
	ReceiptRelayRcvBTCHeaders
	ReceiptRelayLog
	ReqRelayAddrCoins
	ReplyRelayOrders
	QueryRelayOrderParam
	QueryRelayOrderResult
	ReqRelayBtcHeaderHeightList
	ReplyRelayBtcHeadHeightList
	ReqRelayQryBTCHeadHeight
	ReplayRelayQryBTCHeadHeight
	TotalFee
	ReqGetTotalCoins
	ReplyGetTotalCoins
	IterateRangeByStateHash
	TicketStatistic
	TicketMinerInfo
	CreateTx
	UnsignTx
	NoBalanceTx
	SignedTx
	Transaction
	Transactions
	RingSignature
	RingSignatureItem
	Signature
	AddrOverview
	ReqAddr
	ReqPrivacy
	HexTx
	ReplyTxInfo
	ReqTxList
	ReplyTxList
	TxHashList
	ReplyTxInfos
	ReceiptLog
	Receipt
	ReceiptData
	TxResult
	TransactionDetail
	TransactionDetails
	ReqAddrs
	ReqDecodeRawTransaction
	WalletTxDetail
	WalletTxDetails
	WalletAccountStore
	WalletAccountPrivacy
	WalletPwHash
	WalletStatus
	WalletAccounts
	WalletAccount
	WalletUnLock
	GenSeedLang
	GetSeedByPw
	SaveSeedByPw
	ReplySeed
	ReqWalletSetPasswd
	ReqNewAccount
	MinerFlag
	ReqWalletTransactionList
	ReqWalletImportPrivKey
	ReqWalletSendToAddress
	ReqWalletSetFee
	ReqWalletSetLabel
	ReqWalletMergeBalance
	ReqStr
	ReplyStr
	ReqTokenPreCreate
	ReqTokenFinishCreate
	ReqTokenRevokeCreate
	ReqSellToken
	ReqBuyToken
	ReqRevokeSell
	ReqModifyConfig
	ReqSignRawTx
	ReplySignRawTx
	ReportErrEvent
	Int32
	ReqPub2Pri
	ReqPri2Pri
	ReqPri2Pub
	ReqCreateUTXOs
	ReplyPrivacyPkPair
	ReqPrivBal4AddrToken
	ReplyPrivacyBalance
	PrivacyDBStore
	UTXO
	UTXOHaveTxHash
	UTXOs
	UTXOHaveTxHashs
	ReqUTXOGlobalIndex
	UTXOBasic
	UTXOIndex4Amount
	ResUTXOGlobalIndex
	FTXOsSTXOsInOneTx
	ReqCreateTransaction
	RealKeyInput
	UTXOBasics
	CreateTransactionCache
	ReqCacheTxList
	ReplyCacheTxList
	ReqPrivacyAccount
	ReqPPrivacyAccount
	ReplyPrivacyAccount
	ReqCreateCacheTxKey
	ReqBindMiner
	ReplyBindMiner
	ReqPrivacyTransactionList
	ReqRescanUtxos
	RepRescanResult
	RepRescanUtxos
	PrivacySignatureParam
*/
package types

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

// Account 的信息
type Account struct {
	// coins标识，目前只有0 一个值
	Currency int32 `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	// 账户可用余额
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	// 账户冻结余额
	Frozen int64 `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	// 账户的地址
	Addr string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	// 合约地址
	ExecAddr string `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	// 转移前
	Prev *Account `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptExecAccountTransfer) Reset()                    { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()               {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	// 转移前
	Prev *Account `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptAccountTransfer) Reset()                    { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()               {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	// 地址列表
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	// 执行器名称
	Execer    string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
	StateHash string `protobuf:"bytes,3,opt,name=stateHash" json:"stateHash,omitempty"`
}

func (m *ReqBalance) Reset()                    { *m = ReqBalance{} }
func (m *ReqBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()               {}
func (*ReqBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ReqBalance) GetStateHash() string {
	if m != nil {
		return m.StateHash
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc []*Account `protobuf:"bytes,1,rep,name=acc" json:"acc,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ReqTokenBalance struct {
	Addresses   []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	TokenSymbol string   `protobuf:"bytes,2,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	Execer      string   `protobuf:"bytes,3,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqTokenBalance) Reset()                    { *m = ReqTokenBalance{} }
func (m *ReqTokenBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenBalance) ProtoMessage()               {}
func (*ReqTokenBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqTokenBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqTokenBalance) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *ReqTokenBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type ReqAccountTokenAssets struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Execer  string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqAccountTokenAssets) Reset()                    { *m = ReqAccountTokenAssets{} }
func (m *ReqAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReqAccountTokenAssets) ProtoMessage()               {}
func (*ReqAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqAccountTokenAssets) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReqAccountTokenAssets) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type TokenAsset struct {
	Symbol  string   `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *TokenAsset) Reset()                    { *m = TokenAsset{} }
func (m *TokenAsset) String() string            { return proto.CompactTextString(m) }
func (*TokenAsset) ProtoMessage()               {}
func (*TokenAsset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TokenAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenAsset) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ReplyAccountTokenAssets struct {
	TokenAssets []*TokenAsset `protobuf:"bytes,1,rep,name=tokenAssets" json:"tokenAssets,omitempty"`
}

func (m *ReplyAccountTokenAssets) Reset()                    { *m = ReplyAccountTokenAssets{} }
func (m *ReplyAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReplyAccountTokenAssets) ProtoMessage()               {}
func (*ReplyAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReplyAccountTokenAssets) GetTokenAssets() []*TokenAsset {
	if m != nil {
		return m.TokenAssets
	}
	return nil
}

type ExecAccount struct {
	Execer  string   `protobuf:"bytes,1,opt,name=execer" json:"execer,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *ExecAccount) Reset()                    { *m = ExecAccount{} }
func (m *ExecAccount) String() string            { return proto.CompactTextString(m) }
func (*ExecAccount) ProtoMessage()               {}
func (*ExecAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ExecAccount) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ExecAccount) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AllExecBalance struct {
	Addr        string         `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	ExecAccount []*ExecAccount `protobuf:"bytes,2,rep,name=ExecAccount" json:"ExecAccount,omitempty"`
}

func (m *AllExecBalance) Reset()                    { *m = AllExecBalance{} }
func (m *AllExecBalance) String() string            { return proto.CompactTextString(m) }
func (*AllExecBalance) ProtoMessage()               {}
func (*AllExecBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AllExecBalance) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *AllExecBalance) GetExecAccount() []*ExecAccount {
	if m != nil {
		return m.ExecAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ReqTokenBalance)(nil), "types.ReqTokenBalance")
	proto.RegisterType((*ReqAccountTokenAssets)(nil), "types.ReqAccountTokenAssets")
	proto.RegisterType((*TokenAsset)(nil), "types.TokenAsset")
	proto.RegisterType((*ReplyAccountTokenAssets)(nil), "types.ReplyAccountTokenAssets")
	proto.RegisterType((*ExecAccount)(nil), "types.ExecAccount")
	proto.RegisterType((*AllExecBalance)(nil), "types.AllExecBalance")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xe5, 0xa6, 0xbb, 0xdd, 0x4c, 0xc5, 0x22, 0x2c, 0xb1, 0x44, 0x2b, 0x0e, 0x91, 0x4f,
	0x39, 0xa0, 0x1e, 0x58, 0x5e, 0xa0, 0x48, 0x48, 0x70, 0x59, 0x24, 0xd3, 0x13, 0x27, 0x5c, 0x77,
	0x2a, 0xaa, 0x86, 0x24, 0xb5, 0x5d, 0x44, 0x78, 0x00, 0x9e, 0x1b, 0xd9, 0x99, 0x34, 0x2e, 0xa5,
	0xab, 0xde, 0x32, 0x5f, 0xff, 0xf9, 0xcd, 0x4c, 0x0c, 0xcf, 0x94, 0xd6, 0xf5, 0xbe, 0x72, 0xb3,
	0xc6, 0xd4, 0xae, 0xe6, 0x57, 0xae, 0x6d, 0xd0, 0x8a, 0x2d, 0x4c, 0xe6, 0x9d, 0x9f, 0xdf, 0xc3,
	0x8d, 0xde, 0x1b, 0x83, 0x95, 0x6e, 0x33, 0x96, 0xb3, 0xe2, 0x4a, 0x1e, 0x6c, 0x9e, 0xc1, 0x64,
	0xa9, 0x4a, 0x55, 0x69, 0xcc, 0x46, 0x39, 0x2b, 0x12, 0xd9, 0x9b, 0xfc, 0x0e, 0xae, 0xd7, 0xa6,
	0xfe, 0x8d, 0x55, 0x96, 0x84, 0x00, 0x59, 0x9c, 0xc3, 0x58, 0xad, 0x56, 0x26, 0x1b, 0xe7, 0xac,
	0x48, 0x65, 0xf8, 0x16, 0x7f, 0x18, 0xdc, 0x4b, 0xd4, 0xb8, 0x69, 0xdc, 0x87, 0x5f, 0xa8, 0xa9,
	0xf1, 0xc2, 0xa8, 0xca, 0xae, 0xd1, 0x78, 0x00, 0xf4, 0x6e, 0x5f, 0xc6, 0x42, 0xd9, 0xc1, 0xe6,
	0x02, 0xc6, 0x8d, 0xc1, 0x9f, 0xa1, 0xfb, 0xf4, 0xed, 0xed, 0x2c, 0xd0, 0xcf, 0x48, 0x41, 0x86,
	0x18, 0x2f, 0x60, 0xd2, 0x01, 0xbb, 0xc0, 0x72, 0x9a, 0xd6, 0x87, 0xc5, 0x1a, 0xee, 0x88, 0xe3,
	0x5f, 0x86, 0xbe, 0x0f, 0xbb, 0xac, 0xcf, 0xe8, 0xe9, 0x3e, 0xdf, 0x00, 0x24, 0xee, 0xde, 0xd3,
	0xaa, 0x5e, 0x43, 0xea, 0xd7, 0x80, 0xd6, 0xa2, 0xcd, 0x58, 0x9e, 0x14, 0xa9, 0x1c, 0x1c, 0x7e,
	0x91, 0x7e, 0x5a, 0x34, 0x41, 0x34, 0x95, 0x64, 0xf9, 0x2a, 0xeb, 0x94, 0xc3, 0x8f, 0xca, 0x7e,
	0x0f, 0x73, 0xa5, 0x72, 0x70, 0x88, 0x37, 0x70, 0x43, 0x5d, 0x2d, 0xcf, 0x21, 0x51, 0x5a, 0x07,
	0xe5, 0x53, 0x26, 0x1f, 0x12, 0x1b, 0x78, 0x2e, 0x71, 0xb7, 0xa8, 0xb7, 0x58, 0x5d, 0x06, 0x95,
	0xc3, 0xd4, 0xf9, 0xec, 0x2f, 0xed, 0x8f, 0x65, 0x5d, 0x12, 0x59, 0xec, 0x8a, 0xb0, 0x93, 0x18,
	0x5b, 0x7c, 0x82, 0x97, 0x12, 0x77, 0xfd, 0x7a, 0x7d, 0xc1, 0xdc, 0x5a, 0x74, 0xd6, 0xff, 0x4a,
	0xa4, 0x4f, 0x47, 0xee, 0xcd, 0x73, 0x1b, 0x10, 0x8f, 0x00, 0x83, 0x80, 0xcf, 0xb2, 0x1d, 0x4d,
	0x57, 0x4e, 0x96, 0xbf, 0x0a, 0xfd, 0xe1, 0xe7, 0xae, 0x42, 0x61, 0xf1, 0x08, 0xaf, 0x24, 0x36,
	0x65, 0xfb, 0x1f, 0xb8, 0x07, 0x9a, 0xb7, 0x33, 0x69, 0x95, 0x2f, 0x48, 0x68, 0x48, 0x94, 0x71,
	0x96, 0xf8, 0x0c, 0xd3, 0xe8, 0x77, 0x8e, 0xc6, 0x60, 0x47, 0x87, 0xbc, 0x1c, 0xf0, 0x2b, 0xdc,
	0xce, 0xcb, 0xd2, 0x6b, 0xf6, 0x57, 0xea, 0x5f, 0x13, 0x1b, 0x5e, 0x13, 0x7f, 0x77, 0xd4, 0x36,
	0x1b, 0x05, 0x56, 0x4e, 0x9a, 0x51, 0x44, 0xc6, 0x69, 0xcb, 0xeb, 0xf0, 0xfc, 0x1f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x28, 0x3a, 0x97, 0x0f, 0x04, 0x00, 0x00,
}
