//   Copyright (C) 2018 ZVChain
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <https://www.gnu.org/licenses/>.

package base

import (
	"io"
	"math/big"

	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/common/ed25519"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	VRFPublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	VRFPrivateKeySize = 64
	// RandomValueSize is the size, in bytes, of VRF random value as used in this package.
	VRFRandomValueSize = 32
	// ProveSize is the size, in bytes, of VRF prove as used in this package.
	VRFProveSize = 81
)

// VRFPublicKey is the type of Ed25519 public keys.
type VRFPublicKey ed25519.PublicKey

// VRFPrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.
type VRFPrivateKey ed25519.PrivateKey

// VRFRandomValue is the output random value of VRF_Ed25519.
type VRFRandomValue ed25519.VRFRandomValue //RandomValueSize = 32 in bytes

// VRFProve is the output prove of VRF_Ed25519.
type VRFProve ed25519.VRFProve //ProveSize = 81 in bytes

func (vp VRFPublicKey) GetHexString() string {
	return common.ToHex(vp)
}

func Hex2VRFPublicKey(hex string) VRFPublicKey {
	return VRFPublicKey(common.FromHex(hex))
}

func Hex2VRFPrivateKey(hex string) VRFPrivateKey {
	return VRFPrivateKey(common.FromHex(hex))
}

func (vp VRFPrivateKey) GetHexString() string {
	return common.ToHex(vp)
}

// VRFGenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func VRFGenerateKey(rand io.Reader) (publicKey VRFPublicKey, privateKey VRFPrivateKey, err error) {
	pk, sk, err := ed25519.GenerateKey(rand)
	return VRFPublicKey(pk), VRFPrivateKey(sk), err
}

// VRFGenerateProve assume <pk, sk> were generated by ed25519.GenerateKey()
func VRFGenerateProve(pk VRFPublicKey, sk VRFPrivateKey, m []byte) (pi VRFProve, err error) {
	prove, err := ed25519.ECVRFProve(ed25519.PublicKey(pk), ed25519.PrivateKey(sk), m)
	return VRFProve(prove), err
}

func VRFProof2hash(pi VRFProve) VRFRandomValue {
	return VRFRandomValue(ed25519.ECVRFProof2hash(ed25519.VRFProve(pi)))
}

func VRFVerify(pk VRFPublicKey, pi VRFProve, m []byte) (bool, error) {
	return ed25519.ECVRFVerify(ed25519.PublicKey(pk), ed25519.VRFProve(pi), m)
}

func (vp VRFProve) Big() *big.Int {
	return new(big.Int).SetBytes([]byte(vp))
}

func (rv VRFRandomValue) Big() *big.Int {
	return new(big.Int).SetBytes([]byte(rv))
}
