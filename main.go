package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
	"flag"
	"time"
)

func comparePublicKey(key1, key2 *ecdsa.PublicKey) bool {
	x := key1.X.Cmp(key2.X)
	y := key2.Y.Cmp(key2.Y)
	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}


func benchmarkVerify(publicKey, hash, sig []byte, times int) {
	for i := 0; i< times; i++ {
		crypto.VerifySignature(publicKey,hash,sig)
		//result := crypto.VerifySignature(publicKey,hash,sig)
		//fmt.Printf("verify result: %v\n",result)
	}
}

func benchmarkRecovery(publicKey,hash,sig []byte, times int) {
	for i := 0; i< times; i++ {
		crypto.Ecrecover(hash,sig)
		//rPublicKey, err := crypto.Ecrecover(hash,sig)
		//if err != nil {
		//	fmt.Printf("err: %v\n",err)
		//	return
		//}
		//
		//if bytes.Compare(publicKey,rPublicKey) != 0 {
		//	fmt.Println("recovery fail!")
		//} else {
		//	fmt.Println("recovery success!")
		//}
	}
}

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	if err != nil {
		fmt.Printf("err: %v\n",err)
		return
	}

	data := []byte("hello, world")
	hash := sha256.Sum256(data)

	sig, err := crypto.Sign(hash[:],privateKey)
	if err != nil {
		fmt.Printf("err: %v\n",err)
		return
	}

	//fmt.Printf("sig: %X\n",sig)

	bytePublicKey := crypto.FromECDSAPub(&(privateKey.PublicKey))
	//
	//result := crypto.VerifySignature(bytePublicKey,hash[:],sig[:64])
	//
	//fmt.Printf("verify result: %v\n",result)
	//
	//
	//rPublicKey, err := crypto.Ecrecover(hash[:],sig)
	//if err != nil {
	//	fmt.Printf("err: %v\n",err)
	//	return
	//}
	//
	//if bytes.Compare(bytePublicKey,rPublicKey) != 0 {
	//	fmt.Println("recovery fail!")
	//} else {
	//	fmt.Println("recovery success!")
	//}

	times := flag.Int("times",1000,"benchmark times")
	flag.Parse()

	start := time.Now()
	benchmarkRecovery(bytePublicKey,hash[:],sig,*times)
	used := time.Since(start)
	fmt.Printf("run benchmarkRecovery %v times use %v mills.\n",*times,used.Nanoseconds()/1000000)
	start = time.Now()
	benchmarkVerify(bytePublicKey,hash[:],sig[:64],*times)
	used = time.Since(start)
	fmt.Printf("run benchmarkVerify %v times use %v mills.\n",*times,used.Nanoseconds()/1000000)
}
