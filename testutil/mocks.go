package testutil

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MockEthClient is a mock implementation of ethclient.Client for testing
type MockEthClient struct {
	mu               sync.RWMutex
	balances         map[common.Address]*big.Int
	nonces           map[common.Address]uint64
	receipts         map[common.Hash]*types.Receipt
	callResults      map[string][]byte
	callErrors       map[string]error
	gasPrice         *big.Int
	gasTipCap        *big.Int
	baseFee          *big.Int
	estimatedGas     uint64
	estimateGasError error
	sendTxError      error
	blockNumber      uint64
	chainID          *big.Int
}

// NewMockEthClient creates a new mock Ethereum client
func NewMockEthClient(chainID *big.Int) *MockEthClient {
	return &MockEthClient{
		balances:     make(map[common.Address]*big.Int),
		nonces:       make(map[common.Address]uint64),
		receipts:     make(map[common.Hash]*types.Receipt),
		callResults:  make(map[string][]byte),
		callErrors:   make(map[string]error),
		gasPrice:     big.NewInt(20000000000), // 20 gwei
		gasTipCap:    big.NewInt(2000000000),  // 2 gwei
		baseFee:      big.NewInt(10000000000), // 10 gwei
		estimatedGas: 100000,
		chainID:      chainID,
	}
}

// SetBalance sets the balance for an address
func (m *MockEthClient) SetBalance(addr common.Address, balance *big.Int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.balances[addr] = balance
}

// SetNonce sets the nonce for an address
func (m *MockEthClient) SetNonce(addr common.Address, nonce uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nonces[addr] = nonce
}

// SetCallResult sets the result for a contract call
func (m *MockEthClient) SetCallResult(key string, result []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callResults[key] = result
}

// SetCallError sets an error for a contract call
func (m *MockEthClient) SetCallError(key string, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callErrors[key] = err
}

// SetReceipt sets a transaction receipt
func (m *MockEthClient) SetReceipt(txHash common.Hash, receipt *types.Receipt) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.receipts[txHash] = receipt
}

// SetEstimatedGas sets the estimated gas
func (m *MockEthClient) SetEstimatedGas(gas uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.estimatedGas = gas
}

// SetEstimateGasError sets an error for gas estimation
func (m *MockEthClient) SetEstimateGasError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.estimateGasError = err
}

// SetSendTxError sets an error for sending transactions
func (m *MockEthClient) SetSendTxError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.sendTxError = err
}

// BalanceAt returns the balance of an address
func (m *MockEthClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if balance, ok := m.balances[account]; ok {
		return new(big.Int).Set(balance), nil
	}
	return big.NewInt(0), nil
}

// PendingNonceAt returns the nonce of an address
func (m *MockEthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if nonce, ok := m.nonces[account]; ok {
		return nonce, nil
	}
	return 0, nil
}

// EstimateGas estimates gas for a transaction
func (m *MockEthClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.estimateGasError != nil {
		return 0, m.estimateGasError
	}
	return m.estimatedGas, nil
}

// CallContract calls a contract
func (m *MockEthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Create a key from the contract address and data
	key := fmt.Sprintf("%s:%x", msg.To.Hex(), msg.Data)

	if err, ok := m.callErrors[key]; ok {
		return nil, err
	}

	if result, ok := m.callResults[key]; ok {
		return result, nil
	}

	// Return empty result by default
	return []byte{}, nil
}

// SendTransaction sends a transaction
func (m *MockEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.sendTxError != nil {
		return m.sendTxError
	}

	// Auto-create a successful receipt
	receipt := &types.Receipt{
		Status:            types.ReceiptStatusSuccessful,
		TxHash:            tx.Hash(),
		GasUsed:           21000,
		CumulativeGasUsed: 21000,
		Logs:              []*types.Log{},
	}

	// Store receipt (need to unlock and lock for write)
	m.mu.RUnlock()
	m.mu.Lock()
	m.receipts[tx.Hash()] = receipt
	m.mu.Unlock()
	m.mu.RLock()

	return nil
}

// TransactionReceipt gets a transaction receipt
func (m *MockEthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if receipt, ok := m.receipts[txHash]; ok {
		return receipt, nil
	}
	return nil, fmt.Errorf("not found")
}

// HeaderByNumber gets a block header
func (m *MockEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	header := &types.Header{
		BaseFee: new(big.Int).Set(m.baseFee),
		Number:  big.NewInt(int64(m.blockNumber)),
	}
	return header, nil
}

// Close closes the client
func (m *MockEthClient) Close() {
	// No-op for mock
}

// ChainID returns the chain ID
func (m *MockEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return new(big.Int).Set(m.chainID), nil
}

// BlockNumber returns the current block number
func (m *MockEthClient) BlockNumber(ctx context.Context) (uint64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.blockNumber, nil
}

// SetBlockNumber sets the block number
func (m *MockEthClient) SetBlockNumber(number uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.blockNumber = number
}

// SetGasPrice sets the gas price
func (m *MockEthClient) SetGasPrice(price *big.Int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gasPrice = new(big.Int).Set(price)
}

// SetGasTipCap sets the gas tip cap
func (m *MockEthClient) SetGasTipCap(tip *big.Int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gasTipCap = new(big.Int).Set(tip)
}

// SetBaseFee sets the base fee
func (m *MockEthClient) SetBaseFee(fee *big.Int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.baseFee = new(big.Int).Set(fee)
}
