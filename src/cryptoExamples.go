package main

import (
	"bytes"
	"crypto"
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

var HashMapArray []string

func examCrypto() {
	examDSA()
	// examCryptoRand()
}
func examDSA() {
	var privKey dsa.PrivateKey
	dsa.GenerateParameters(&privKey.PublicKey.Parameters,
		rand.Reader,
		dsa.L1024N160)
	dsa.GenerateKey(&privKey, rand.Reader)
	fmt.Println(privKey)
	hashS := make([]byte, 1024, 1024)
	r, s, err := dsa.Sign(rand.Reader, &privKey,
		hashS)
	fmt.Printf("%X, %X, %e\n", r, s, err)
	if dsa.Verify(&privKey.PublicKey, hashS, r, s) {
		println("验证通过")
	}
}
func examAES() {

}

func examCryptoRand() {
	c := 24
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	// fmt.Println(bytes.Equal(b, make([]byte, c)))
	// var reader bytes.Reader = bytes.NewReader(b)
	for index := range b {
		if index%4 == 0 {
			switch index {
			case 0:

			default:
				print("-")

			}
		}
		fmt.Printf("%X", b[index])
	}
	println("\n--------------------------------")
	max := big.NewInt(100)
	// ia, errI := strconv.Atoi(string(b))
	// errI = errI
	fmt.Println("Buffer:", fmt.Sprintf("%X", b))
	n, err := rand.Int(bytes.NewBuffer(b), max)
	n, err = rand.Int(bytes.NewReader(b), max)
	if err == nil {
		fmt.Println("Max:", max)
		fmt.Println("RandInt:", n)
		fmt.Println("Buffer:", string(b))
	}
	n, err = rand.Prime(bytes.NewBuffer(b), 31)
	if err == nil {
		fmt.Println("Prime:", n)
		fmt.Println("Buffer:", string(b))
	}
	retInt1 := 0
	retInt1, err = rand.Read(b)
	if err == nil && retInt1 > 0 {
		fmt.Println("Rand.Read:", fmt.Sprintf("%x", string(b)))
	}

}
func examCryptoMD5() {
	HashMapArray = []string{"MD4", // import golang.org/x/crypto/md4
		"MD5",         // import crypto/md5
		"SHA1",        // import crypto/sha1
		"SHA224",      // import crypto/sha256
		"SHA256",      // import crypto/sha256
		"SHA384",      // import crypto/sha512
		"SHA512",      // import crypto/sha512
		"MD5SHA1",     // no implementation; MD5+SHA1 used for TLS RSA
		"RIPEMD160",   // import golang.org/x/crypto/ripemd160
		"SHA3_224",    // import golang.org/x/crypto/sha3
		"SHA3_256",    // import golang.org/x/crypto/sha3
		"SHA3_384",    // import golang.org/x/crypto/sha3
		"SHA3_512",    // import golang.org/x/crypto/sha3
		"SHA512_224",  // import crypto/sha512
		"SHA512_256",  // import crypto/sha512
		"BLAKE2s_256", // import golang.org/x/crypto/blake2s
		"BLAKE2b_256", // import golang.org/x/crypto/blake2b
		"BLAKE2b_384", // import golang.org/x/crypto/blake2b
		"BLAKE2b_512"} // import golang.org/x/crypto/blake2b}

	var myHash crypto.Hash = crypto.MD4
	myHash = myHash

	data := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	has := md5.Sum(data)
	fmt.Printf("%X\n", has) //将[]byte转成16进制

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%X\n", h.Sum(nil))

	for index := crypto.MD4; index <= crypto.BLAKE2b_512; index++ {
		myHash = crypto.Hash(index)
		// print(HashMapArray[index-1])
		// if myHash.Available() {
		// 	fmt.Println(" *****available*****")
		// } else {
		// 	fmt.Println(" unavailable")
		// }

	}

}
