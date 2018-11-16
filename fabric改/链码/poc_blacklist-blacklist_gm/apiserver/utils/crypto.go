package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/peersafe/gm-crypto/sm2"
	"github.com/peersafe/gm-crypto/sm4"
)


func IsFileExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("ReadFile open file [%s] err : %s\n", path, err.Error())
		return nil, err
	}
	info, err := os.Stat(path)
	fileSize := info.Size()
	buff := make([]byte, fileSize)
	_, err = f.Read(buff)
	if err != nil {
		fmt.Printf("ReadFile read file [%s] err : %s\n", path, err.Error())
		return nil, err
	}

	f.Close()

	return buff, nil
}

func GenerateKey(bitsize int) []byte {
	key := ""

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < bitsize/2; i++ {
		key += fmt.Sprintf("%x", r.Int63())
	}

	return []byte(key)
}

func SM4Encrypt(originData, key []byte) ([]byte, error) {
	c, err := sm4.NewCipher(key, nil)
	if err != nil {
		return nil, err
	}

	cipherData := make([]byte, (len(originData)/16+1)*16+16)
	c.Encrypt(cipherData, originData)

	return cipherData, nil
}

func SM4Decrypt(cipherData, key []byte) ([]byte, error) {
	c, err := sm4.NewCipher(key, nil)
	if err != nil {
		return nil, err
	}

	originData := make([]byte, (len(cipherData)/16+1)*16+16)
	c.Decrypt(originData, cipherData)

	return originData, nil
}

func SM2Encrypt(originData, key []byte) ([]byte, error) {
	cert, err := sm2.ReadCertificateFromMem(key)
	if err != nil {
		return nil, err
	}

	pub, ok := cert.PublicKey.(*sm2.PublicKey)
	if !ok {
		return nil, fmt.Errorf("read sm2 public key from cert error !")
	}

	return pub.Encrypt(originData)
}

func SM2Decrypt(originData, key []byte) ([]byte, error) {
	priv, err := sm2.ReadPrivateKeyFromMem(key, nil)
	if err != nil {
		return nil, err
	}

	return priv.Decrypt(originData)
}
