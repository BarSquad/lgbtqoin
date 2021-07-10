package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"math/big"
	"strings"
)

func Public(private *big.Int) (publicKey string) {
	var e ecdsa.PrivateKey
	e.D = private
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return fmt.Sprintf("%x", elliptic.Marshal(secp256k1.S256(), e.X, e.Y))
}
func main() {
	privateKey, _ := rand.Int(rand.Reader, secp256k1.S256().N)
	log.Println(privateKey)
	log.Println(strings.ToUpper(Public(privateKey)))
}
