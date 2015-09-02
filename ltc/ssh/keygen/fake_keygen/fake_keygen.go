// This file was generated by counterfeiter
package fake_keygen

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/ssh/keygen"
)

type FakeKeyGenerator struct {
	GenerateRSAPrivateKeyStub        func(bits int) (pemEncodedPrivateKey string, err error)
	generateRSAPrivateKeyMutex       sync.RWMutex
	generateRSAPrivateKeyArgsForCall []struct {
		bits int
	}
	generateRSAPrivateKeyReturns struct {
		result1 string
		result2 error
	}
	GenerateRSAKeyPairStub        func(bits int) (pemEncodedPrivateKey string, authorizedKey string, err error)
	generateRSAKeyPairMutex       sync.RWMutex
	generateRSAKeyPairArgsForCall []struct {
		bits int
	}
	generateRSAKeyPairReturns struct {
		result1 string
		result2 string
		result3 error
	}
}

func (fake *FakeKeyGenerator) GenerateRSAPrivateKey(bits int) (pemEncodedPrivateKey string, err error) {
	fake.generateRSAPrivateKeyMutex.Lock()
	fake.generateRSAPrivateKeyArgsForCall = append(fake.generateRSAPrivateKeyArgsForCall, struct {
		bits int
	}{bits})
	fake.generateRSAPrivateKeyMutex.Unlock()
	if fake.GenerateRSAPrivateKeyStub != nil {
		return fake.GenerateRSAPrivateKeyStub(bits)
	} else {
		return fake.generateRSAPrivateKeyReturns.result1, fake.generateRSAPrivateKeyReturns.result2
	}
}

func (fake *FakeKeyGenerator) GenerateRSAPrivateKeyCallCount() int {
	fake.generateRSAPrivateKeyMutex.RLock()
	defer fake.generateRSAPrivateKeyMutex.RUnlock()
	return len(fake.generateRSAPrivateKeyArgsForCall)
}

func (fake *FakeKeyGenerator) GenerateRSAPrivateKeyArgsForCall(i int) int {
	fake.generateRSAPrivateKeyMutex.RLock()
	defer fake.generateRSAPrivateKeyMutex.RUnlock()
	return fake.generateRSAPrivateKeyArgsForCall[i].bits
}

func (fake *FakeKeyGenerator) GenerateRSAPrivateKeyReturns(result1 string, result2 error) {
	fake.GenerateRSAPrivateKeyStub = nil
	fake.generateRSAPrivateKeyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyGenerator) GenerateRSAKeyPair(bits int) (pemEncodedPrivateKey string, authorizedKey string, err error) {
	fake.generateRSAKeyPairMutex.Lock()
	fake.generateRSAKeyPairArgsForCall = append(fake.generateRSAKeyPairArgsForCall, struct {
		bits int
	}{bits})
	fake.generateRSAKeyPairMutex.Unlock()
	if fake.GenerateRSAKeyPairStub != nil {
		return fake.GenerateRSAKeyPairStub(bits)
	} else {
		return fake.generateRSAKeyPairReturns.result1, fake.generateRSAKeyPairReturns.result2, fake.generateRSAKeyPairReturns.result3
	}
}

func (fake *FakeKeyGenerator) GenerateRSAKeyPairCallCount() int {
	fake.generateRSAKeyPairMutex.RLock()
	defer fake.generateRSAKeyPairMutex.RUnlock()
	return len(fake.generateRSAKeyPairArgsForCall)
}

func (fake *FakeKeyGenerator) GenerateRSAKeyPairArgsForCall(i int) int {
	fake.generateRSAKeyPairMutex.RLock()
	defer fake.generateRSAKeyPairMutex.RUnlock()
	return fake.generateRSAKeyPairArgsForCall[i].bits
}

func (fake *FakeKeyGenerator) GenerateRSAKeyPairReturns(result1 string, result2 string, result3 error) {
	fake.GenerateRSAKeyPairStub = nil
	fake.generateRSAKeyPairReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

var _ keygen.KeyGenerator = new(FakeKeyGenerator)
