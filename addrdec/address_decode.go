package addrdec

import (
	"fmt"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	TGCPublicKeyPrefix       = "PUB_"
	TGCPublicKeyK1Prefix     = "PUB_K1_"
	TGCPublicKeyR1Prefix     = "PUB_R1_"
	TGCPublicKeyPrefixCompat = "EVS"

	//TGC stuff
	TGC_mainnetPublic = addressEncoder.AddressType{"tgc", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(TGCPublicKeyPrefixCompat), nil}
	// TGC_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// TGC_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, TGCPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(TGCPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, TGCPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(TGCPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, TGCPublicKeyPrefixCompat) { // "TGC"
		pubKeyMaterial = pubKey[len(TGCPublicKeyPrefixCompat):] // strip "TGC"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", TGCPublicKeyK1Prefix, TGCPublicKeyR1Prefix, TGCPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(TGC_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, TGC_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte) string {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, TGC_mainnetPublic.ChecksumType))
	return string(TGC_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", TGC_mainnetPublic.Alphabet)
}
