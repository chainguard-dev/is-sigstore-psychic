package main

import (
	"encoding/base64"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"math/big"
	"os"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

func ParseECDSASignature(sigString string) (*big.Int, *big.Int, error) {
	var (
		r, s  = &big.Int{}, &big.Int{}
		inner cryptobyte.String
	)
	sig, err := base64.StdEncoding.DecodeString(sigString)
	if err != nil {
		return nil, nil, errors.New("invalid base64") // probably CSV header line
	}
	input := cryptobyte.String(sig)
	if !input.ReadASN1(&inner, asn1.SEQUENCE) ||
		!input.Empty() ||
		!inner.ReadASN1Integer(r) ||
		!inner.ReadASN1Integer(s) ||
		!inner.Empty() {
		return nil, nil, errors.New("not ECDSA signature")
	}
	return r, s, nil
}

func main() {
	r := csv.NewReader(os.Stdin)
	for {
		record, err := r.Read() // record: logIndex,uuid,signature
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else if r.FieldsPerRecord != 3 {
			log.Fatal("Unexpected input data; expected 3 fields.")
		}
		logIndex := record[0]
		uuid := record[1]
		sig := record[2]

		r, s, err := ParseECDSASignature(sig)
		if err != nil {
			continue // invalid ECDSA signature; it's probably in another format
		}
		zero := big.NewInt(0)
		if r.Cmp(zero) == 0 && s.Cmp(zero) == 0 {
			log.Printf("Found psychic signature with UUID %s (idx: %s)", uuid, logIndex)
		}
	}
}
