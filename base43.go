/*
Based on https://github.com/btcsuite/btcutil/blob/v1.0.2/base58/base58.go
*/

package main

import (
	"fmt"
	"log"
	"math/big"
)

type base43 struct { }

const base43CharSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ$*+-./:"

var bigRadix = big.NewInt(43)
var bigZero = big.NewInt(0)

// Convert Base43 to binary data
func (b43 base43) Decode(b43Data []byte) ([]byte, error) {

	// Check the data looks like Base43
	loop:
	for n,b := range b43Data {
		for _,c := range base43CharSet {
			if b == byte(c) { // valid base43 character
				continue loop
			}
		}
		log.Fatalf("Data contains non-base43 character '%c' at position %d\n",
			rune(b), n+1)
	}

	// Create an array of rune values
	var b43Value [256]byte
	for i := 0; i<256; i++ {
		b43Value[i] = 255
	}
	for n, r := range base43CharSet {
		b43Value[byte(r)] = byte(n)
	}

	// Start decoding

	answer := new(big.Int)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(b43Data) -1; i>= 0; i-- {
		tmp := b43Value[b43Data[i]]
		if tmp == 255 {
			return []byte(""), 
				fmt.Errorf("Undetected illegal base43 character '%c'\n", tmp)
		}
		scratch.SetInt64(int64(tmp))
		scratch.Mul(j, scratch)
		answer.Add(answer, scratch)
		j.Mul(j, bigRadix)
	}

	return answer.Bytes(), nil
}

// Encodes binary bytes in Base43 
func (b43 base43) Encode(binData []byte) ([]byte) {
	x := new(big.Int)
	x.SetBytes(binData)

	answer := make([]byte, 0, len(binData)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, base43CharSet[mod.Int64()])
	}
	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return answer
}
