package sm4

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func getTime() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}
func TestSM4(t *testing.T) {
	var err error
	d0 := make([]byte, 16)
	d1 := make([]byte, 16)
	var start, end int64
	blockSize := 16
	cricleTimes := 1

	key := []byte("1234567890abcdef")
	c, err := NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	var msg []byte
	for i := 0; i < blockSize; i++ {
		msg = append(msg, 'a')
	}

	start = getTime()
	for i := 0; i < cricleTimes; i++ {
		c.Encrypt(d0, msg)
	}
	end = getTime()
	fmt.Printf("encrypt %d times use %d ms!\n", cricleTimes, end-start)

	start = getTime()
	for i := 0; i < cricleTimes; i++ {
		c.Decrypt(d1, d0)
	}
	end = getTime()
	fmt.Printf("decrypt %d times use %d ms!\n", cricleTimes, end-start)
}

func BenchmarkSM4(t *testing.B) {
	t.ReportAllocs()
	key := []byte("1234567890abcdef")
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	WriteKeyToPem("key.pem", key, nil)
	key, err := ReadKeyFromPem("key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
	c, err := NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < t.N; i++ {
		d0 := make([]byte, 16)
		c.Encrypt(d0, data)
		d1 := make([]byte, 16)
		c.Decrypt(d1, d0)
	}
}

// func TestErrKeyLen(t *testing.T) {
// 	fmt.Printf("\n--------------test key len------------------")
// 	key := []byte("1234567890abcdefg")
// 	_, err := NewCipher(key)
// 	if err != nil {
// 		fmt.Println("\nError key len !")
// 	}
// 	key = []byte("1234")
// 	_, err = NewCipher(key)
// 	if err != nil {
// 		fmt.Println("Error key len !")
// 	}
// 	fmt.Println("------------------end----------------------")
// }

func testCompare(key1, key2 []byte) bool {
	if len(key1) != len(key2) {
		return false
	}
	for i, v := range key1 {
		if i == 1 {
			fmt.Println("type of v", reflect.TypeOf(v))
		}
		a := key2[i]
		if a != v {
			return false
		}
	}
	return true
}
