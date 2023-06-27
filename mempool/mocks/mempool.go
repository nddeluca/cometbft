// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	abcicli "github.com/cometbft/cometbft/abci/client"
	abcitypes "github.com/cometbft/cometbft/abci/types"

	log "github.com/cometbft/cometbft/libs/log"

	mempool "github.com/cometbft/cometbft/mempool"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cometbft/cometbft/types"
)

// Mempool is an autogenerated mock type for the Mempool type
type Mempool struct {
	mock.Mock
}

// CheckTx provides a mock function with given fields: tx
func (_m *Mempool) CheckTx(tx types.Tx) (*abcicli.ReqRes, error) {
	ret := _m.Called(tx)

	var r0 *abcicli.ReqRes
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Tx) (*abcicli.ReqRes, error)); ok {
		return rf(tx)
	}
	if rf, ok := ret.Get(0).(func(types.Tx) *abcicli.ReqRes); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableTxsAvailable provides a mock function with given fields:
func (_m *Mempool) EnableTxsAvailable() {
	_m.Called()
}

// EnableTxsRemoved provides a mock function with given fields:
func (_m *Mempool) EnableTxsRemoved() {
	_m.Called()
}

// Flush provides a mock function with given fields:
func (_m *Mempool) Flush() {
	_m.Called()
}

// FlushAppConn provides a mock function with given fields:
func (_m *Mempool) FlushAppConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Lock provides a mock function with given fields:
func (_m *Mempool) Lock() {
	_m.Called()
}

// NewIterator provides a mock function with given fields:
func (_m *Mempool) NewIterator() mempool.MempoolIterator {
	ret := _m.Called()

	var r0 mempool.MempoolIterator
	if rf, ok := ret.Get(0).(func() mempool.MempoolIterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mempool.MempoolIterator)
		}
	}

	return r0
}

// ReapMaxBytesMaxGas provides a mock function with given fields: maxBytes, maxGas
func (_m *Mempool) ReapMaxBytesMaxGas(maxBytes int64, maxGas int64) types.Txs {
	ret := _m.Called(maxBytes, maxGas)

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int64, int64) types.Txs); ok {
		r0 = rf(maxBytes, maxGas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// ReapMaxTxs provides a mock function with given fields: max
func (_m *Mempool) ReapMaxTxs(max int) types.Txs {
	ret := _m.Called(max)

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int) types.Txs); ok {
		r0 = rf(max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// RemoveTxByKey provides a mock function with given fields: txKey
func (_m *Mempool) RemoveTxByKey(txKey types.TxKey) error {
	ret := _m.Called(txKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.TxKey) error); ok {
		r0 = rf(txKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogger provides a mock function with given fields: l
func (_m *Mempool) SetLogger(l log.Logger) {
	_m.Called(l)
}

// Size provides a mock function with given fields:
func (_m *Mempool) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SizeBytes provides a mock function with given fields:
func (_m *Mempool) SizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Mempool) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxsAvailable provides a mock function with given fields:
func (_m *Mempool) TxsAvailable() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// TxsRemoved provides a mock function with given fields:
func (_m *Mempool) TxsRemoved() <-chan types.TxKey {
	ret := _m.Called()

	var r0 <-chan types.TxKey
	if rf, ok := ret.Get(0).(func() <-chan types.TxKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan types.TxKey)
		}
	}

	return r0
}

// Unlock provides a mock function with given fields:
func (_m *Mempool) Unlock() {
	_m.Called()
}

// Update provides a mock function with given fields: blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn
func (_m *Mempool) Update(blockHeight int64, blockTxs types.Txs, deliverTxResponses []*abcitypes.ExecTxResult, newPreFn mempool.PreCheckFunc, newPostFn mempool.PostCheckFunc) error {
	ret := _m.Called(blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, types.Txs, []*abcitypes.ExecTxResult, mempool.PreCheckFunc, mempool.PostCheckFunc) error); ok {
		r0 = rf(blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMempool interface {
	mock.TestingT
	Cleanup(func())
}

// NewMempool creates a new instance of Mempool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMempool(t mockConstructorTestingTNewMempool) *Mempool {
	mock := &Mempool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
