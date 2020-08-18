// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	keychain "github.com/FleekHQ/space-daemon/core/keychain"
	crypto "github.com/libp2p/go-libp2p-core/crypto"

	mock "github.com/stretchr/testify/mock"
)

// Keychain is an autogenerated mock type for the Keychain type
type Keychain struct {
	mock.Mock
}

// GenerateKeyFromMnemonic provides a mock function with given fields: _a0
func (_m *Keychain) GenerateKeyFromMnemonic(_a0 ...keychain.GenerateKeyFromMnemonicOpts) (string, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...keychain.GenerateKeyFromMnemonicOpts) string); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...keychain.GenerateKeyFromMnemonicOpts) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateKeyPair provides a mock function with given fields:
func (_m *Keychain) GenerateKeyPair() ([]byte, []byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func() []byte); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GenerateKeyPairWithForce provides a mock function with given fields:
func (_m *Keychain) GenerateKeyPairWithForce() ([]byte, []byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func() []byte); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStoredKeyPairInLibP2PFormat provides a mock function with given fields:
func (_m *Keychain) GetStoredKeyPairInLibP2PFormat() (crypto.PrivKey, crypto.PubKey, error) {
	ret := _m.Called()

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	var r1 crypto.PubKey
	if rf, ok := ret.Get(1).(func() crypto.PubKey); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(crypto.PubKey)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStoredPublicKey provides a mock function with given fields:
func (_m *Keychain) GetStoredPublicKey() (crypto.PubKey, error) {
	ret := _m.Called()

	var r0 crypto.PubKey
	if rf, ok := ret.Get(0).(func() crypto.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PubKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportExistingKeyPair provides a mock function with given fields: priv
func (_m *Keychain) ImportExistingKeyPair(priv crypto.PrivKey) error {
	ret := _m.Called(priv)

	var r0 error
	if rf, ok := ret.Get(0).(func(crypto.PrivKey) error); ok {
		r0 = rf(priv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sign provides a mock function with given fields: _a0
func (_m *Keychain) Sign(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
