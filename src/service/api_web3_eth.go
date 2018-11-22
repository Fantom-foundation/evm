package service

import (
	//"compress/gzip"
	"context"
	//"errors"
	"fmt"
	//"io"
	//"math/big"
	//"os"
	//"runtime"
	//"strings"
	//"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	//"github.com/ethereum/go-ethereum/core/rawdb"
	//"github.com/andrecronje/evm/src/service/internal/ethapi"
)

var ErrNotImplemented = fmt.Errorf("Not implemented yet")

// GetAPIsEth return the collection of RPC services the ethereum package offers.
func GetEthAPIs(s *Service) []rpc.API {
	apis := []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(s),
			Public:    true,
		}, /*{
			Namespace: "eth",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.APIBackend, false),
			Public:    true,
		},*/{
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, /*{
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s.chainConfig, s),
		},*/ /*{
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},*/
	}

	return apis
}

// PublicEthereumAPI provides an API to access Service full node-related
// information.
type PublicEthereumAPI struct {
	e *Service
}

// NewPublicEthereumAPI creates a new Ethereum protocol API for full nodes.
func NewPublicEthereumAPI(e *Service) *PublicEthereumAPI {
	return &PublicEthereumAPI{e}
}

// Etherbase is the address that mining rewards will be send to
func (api *PublicEthereumAPI) Etherbase() (common.Address, error) {
	/*
		return api.e.Etherbase()
	*/
	return common.Address{}, ErrNotImplemented
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase)
func (api *PublicEthereumAPI) Coinbase() (common.Address, error) {
	/*
		return api.Etherbase()
	*/
	return common.Address{}, ErrNotImplemented
}

// Hashrate returns the POW hashrate
func (api *PublicEthereumAPI) Hashrate() hexutil.Uint64 {
	/*
		return hexutil.Uint64(api.e.Miner().HashRate())
	*/
	return hexutil.Uint64(0)
}

// ChainId is the EIP-155 replay-protection chain id for the current ethereum chain config.
func (api *PublicEthereumAPI) ChainId() hexutil.Uint64 {
	/*
		chainID := new(big.Int)
		if config := api.e.chainConfig; config.IsEIP155(api.e.blockchain.CurrentBlock().Number()) {
			chainID = config.ChainID
		}
		return (hexutil.Uint64)(chainID.Uint64())
	*/
	return hexutil.Uint64(0)
}

// PrivateAdminAPI is the collection of Ethereum full node-related APIs
// exposed over the private admin endpoint.
type PrivateAdminAPI struct {
	eth *Service
}

// NewPrivateAdminAPI creates a new API definition for the full node private
// admin methods of the Ethereum service.
func NewPrivateAdminAPI(eth *Service) *PrivateAdminAPI {
	return &PrivateAdminAPI{eth: eth}
}

// ExportChain exports the current blockchain into a local file.
func (api *PrivateAdminAPI) ExportChain(file string) (bool, error) {
	/*
		// Make sure we can create the file to export into
		out, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return false, err
		}
		defer out.Close()

		var writer io.Writer = out
		if strings.HasSuffix(file, ".gz") {
			writer = gzip.NewWriter(writer)
			defer writer.(*gzip.Writer).Close()
		}

		// Export the blockchain
		if err := api.eth.BlockChain().Export(writer); err != nil {
			return false, err
		}
		return true, nil
	*/
	return false, ErrNotImplemented
}

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	/*
		for _, b := range bs {
			if !chain.HasBlock(b.Hash(), b.NumberU64()) {
				return false
			}
		}

		return true
	*/
	return false
}

// ImportChain imports a blockchain from a local file.
func (api *PrivateAdminAPI) ImportChain(file string) (bool, error) {
	/*
		// Make sure the can access the file to import
		in, err := os.Open(file)
		if err != nil {
			return false, err
		}
		defer in.Close()

		var reader io.Reader = in
		if strings.HasSuffix(file, ".gz") {
			if reader, err = gzip.NewReader(reader); err != nil {
				return false, err
			}
		}

		// Run actual the import in pre-configured batches
		stream := rlp.NewStream(reader, 0)

		blocks, index := make([]*types.Block, 0, 2500), 0
		for batch := 0; ; batch++ {
			// Load a batch of blocks from the input file
			for len(blocks) < cap(blocks) {
				block := new(types.Block)
				if err := stream.Decode(block); err == io.EOF {
					break
				} else if err != nil {
					return false, fmt.Errorf("block %d: failed to parse: %v", index, err)
				}
				blocks = append(blocks, block)
				index++
			}
			if len(blocks) == 0 {
				break
			}

			if hasAllBlocks(api.eth.BlockChain(), blocks) {
				blocks = blocks[:0]
				continue
			}
			// Import the batch and reset the buffer
			if _, err := api.eth.BlockChain().InsertChain(blocks); err != nil {
				return false, fmt.Errorf("batch %d: failed to insert: %v", batch, err)
			}
			blocks = blocks[:0]
		}
		return true, nil
	*/
	return false, ErrNotImplemented
}

// PublicDebugAPI is the collection of Ethereum full node APIs exposed
// over the public debugging endpoint.
type PublicDebugAPI struct {
	eth *Service
}

// NewPublicDebugAPI creates a new API definition for the full node-
// related public debug methods of the Ethereum service.
func NewPublicDebugAPI(eth *Service) *PublicDebugAPI {
	return &PublicDebugAPI{eth: eth}
}

// DumpBlock retrieves the entire state of the database at a given block.
func (api *PublicDebugAPI) DumpBlock(blockNr rpc.BlockNumber) (state.Dump, error) {
	/*
		if blockNr == rpc.PendingBlockNumber {
			// If we're dumping the pending state, we need to request
			// both the pending block as well as the pending state from
			// the miner and operate on those
			_, stateDb := api.eth.miner.Pending()
			return stateDb.RawDump(), nil
		}
		var block *types.Block
		if blockNr == rpc.LatestBlockNumber {
			block = api.eth.blockchain.CurrentBlock()
		} else {
			block = api.eth.blockchain.GetBlockByNumber(uint64(blockNr))
		}
		if block == nil {
			return state.Dump{}, fmt.Errorf("block #%d not found", blockNr)
		}
		stateDb, err := api.eth.BlockChain().StateAt(block.Root())
		if err != nil {
			return state.Dump{}, err
		}
		return stateDb.RawDump(), nil
	*/
	return state.Dump{}, ErrNotImplemented
}

// PrivateDebugAPI is the collection of Ethereum full node APIs exposed over
// the private debugging endpoint.
type PrivateDebugAPI struct {
	config *params.ChainConfig
	eth    *Service
}

// NewPrivateDebugAPI creates a new API definition for the full node-related
// private debug methods of the Ethereum service.
func NewPrivateDebugAPI(config *params.ChainConfig, eth *Service) *PrivateDebugAPI {
	return &PrivateDebugAPI{config: config, eth: eth}
}

// Preimage is a debug API function that returns the preimage for a sha3 hash, if known.
func (api *PrivateDebugAPI) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	/*
		if preimage := rawdb.ReadPreimage(api.eth.ChainDb(), hash); preimage != nil {
			return preimage, nil
		}
		return nil, errors.New("unknown preimage")
	*/
	return nil, ErrNotImplemented
}

// BadBlockArgs represents the entries in the list returned when bad blocks are queried.
type BadBlockArgs struct {
	Hash  common.Hash            `json:"hash"`
	Block map[string]interface{} `json:"block"`
	RLP   string                 `json:"rlp"`
}

// GetBadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
// and returns them as a JSON list of block-hashes
func (api *PrivateDebugAPI) GetBadBlocks(ctx context.Context) ([]*BadBlockArgs, error) {
	/*
		blocks := api.eth.BlockChain().BadBlocks()
		results := make([]*BadBlockArgs, len(blocks))

		var err error
		for i, block := range blocks {
			results[i] = &BadBlockArgs{
				Hash: block.Hash(),
			}
			if rlpBytes, err := rlp.EncodeToBytes(block); err != nil {
				results[i].RLP = err.Error() // Hacky, but hey, it works
			} else {
				results[i].RLP = fmt.Sprintf("0x%x", rlpBytes)
			}
			if results[i].Block, err = ethapi.RPCMarshalBlock(block, true, true); err != nil {
				results[i].Block = map[string]interface{}{"error": err.Error()}
			}
		}
		return results, nil
	*/
	return nil, ErrNotImplemented
}

// StorageRangeResult is the result of a debug_storageRangeAt API call.
type StorageRangeResult struct {
	Storage storageMap   `json:"storage"`
	NextKey *common.Hash `json:"nextKey"` // nil if Storage includes the last key in the trie.
}

type storageMap map[common.Hash]storageEntry

type storageEntry struct {
	Key   *common.Hash `json:"key"`
	Value common.Hash  `json:"value"`
}

// StorageRangeAt returns the storage at the given block height and transaction index.
func (api *PrivateDebugAPI) StorageRangeAt(ctx context.Context, blockHash common.Hash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (StorageRangeResult, error) {
	/*
		_, _, statedb, err := api.computeTxEnv(blockHash, txIndex, 0)
		if err != nil {
			return StorageRangeResult{}, err
		}
		st := statedb.StorageTrie(contractAddress)
		if st == nil {
			return StorageRangeResult{}, fmt.Errorf("account %x doesn't exist", contractAddress)
		}
		return storageRangeAt(st, keyStart, maxResult)
	*/
	return StorageRangeResult{}, ErrNotImplemented
}

func storageRangeAt(st state.Trie, start []byte, maxResult int) (StorageRangeResult, error) {
	it := trie.NewIterator(st.NodeIterator(start))
	result := StorageRangeResult{Storage: storageMap{}}
	for i := 0; i < maxResult && it.Next(); i++ {
		_, content, _, err := rlp.Split(it.Value)
		if err != nil {
			return StorageRangeResult{}, err
		}
		e := storageEntry{Value: common.BytesToHash(content)}
		if preimage := st.GetKey(it.Key); preimage != nil {
			preimage := common.BytesToHash(preimage)
			e.Key = &preimage
		}
		result.Storage[common.BytesToHash(it.Key)] = e
	}
	// Add the 'next key' so clients can continue downloading.
	if it.Next() {
		next := common.BytesToHash(it.Key)
		result.NextKey = &next
	}
	return result, nil
}

// GetModifiedAccountsByNumber returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *PrivateDebugAPI) GetModifiedAccountsByNumber(startNum uint64, endNum *uint64) ([]common.Address, error) {
	/*
		var startBlock, endBlock *types.Block

		startBlock = api.eth.blockchain.GetBlockByNumber(startNum)
		if startBlock == nil {
			return nil, fmt.Errorf("start block %x not found", startNum)
		}

		if endNum == nil {
			endBlock = startBlock
			startBlock = api.eth.blockchain.GetBlockByHash(startBlock.ParentHash())
			if startBlock == nil {
				return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
			}
		} else {
			endBlock = api.eth.blockchain.GetBlockByNumber(*endNum)
			if endBlock == nil {
				return nil, fmt.Errorf("end block %d not found", *endNum)
			}
		}
		return api.getModifiedAccounts(startBlock, endBlock)
	*/
	return nil, ErrNotImplemented
}

// GetModifiedAccountsByHash returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *PrivateDebugAPI) GetModifiedAccountsByHash(startHash common.Hash, endHash *common.Hash) ([]common.Address, error) {
	/*
		var startBlock, endBlock *types.Block
		startBlock = api.eth.blockchain.GetBlockByHash(startHash)
		if startBlock == nil {
			return nil, fmt.Errorf("start block %x not found", startHash)
		}

		if endHash == nil {
			endBlock = startBlock
			startBlock = api.eth.blockchain.GetBlockByHash(startBlock.ParentHash())
			if startBlock == nil {
				return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
			}
		} else {
			endBlock = api.eth.blockchain.GetBlockByHash(*endHash)
			if endBlock == nil {
				return nil, fmt.Errorf("end block %x not found", *endHash)
			}
		}
		return api.getModifiedAccounts(startBlock, endBlock)
	*/
	return nil, ErrNotImplemented
}

func (api *PrivateDebugAPI) getModifiedAccounts(startBlock, endBlock *types.Block) ([]common.Address, error) {
	/*
		if startBlock.Number().Uint64() >= endBlock.Number().Uint64() {
			return nil, fmt.Errorf("start block height (%d) must be less than end block height (%d)", startBlock.Number().Uint64(), endBlock.Number().Uint64())
		}

		oldTrie, err := trie.NewSecure(startBlock.Root(), trie.NewDatabase(api.eth.chainDb), 0)
		if err != nil {
			return nil, err
		}
		newTrie, err := trie.NewSecure(endBlock.Root(), trie.NewDatabase(api.eth.chainDb), 0)
		if err != nil {
			return nil, err
		}

		diff, _ := trie.NewDifferenceIterator(oldTrie.NodeIterator([]byte{}), newTrie.NodeIterator([]byte{}))
		iter := trie.NewIterator(diff)

		var dirty []common.Address
		for iter.Next() {
			key := newTrie.GetKey(iter.Key)
			if key == nil {
				return nil, fmt.Errorf("no preimage found for hash %x", iter.Key)
			}
			dirty = append(dirty, common.BytesToAddress(key))
		}
		return dirty, nil
	*/
	return nil, ErrNotImplemented
}
