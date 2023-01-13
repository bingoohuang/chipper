package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/bingoohuang/chipper/core"
	"io"
)

type aesGcmTest struct {
	baseStepTest
	key   []byte
	nonce []byte
}

// NewAesGcmTest creates new factorial test
func NewAesGcmTest(n uint) core.Test {
	return &aesGcmTest{
		baseStepTest: baseStepTest{n: n},
		key:          randBytes(32),
		nonce:        randBytes(12),
	}
}

func (f *aesGcmTest) Name() string {
	return f.nameBase(AesGcm)
}

func (f *aesGcmTest) Start() {
	plaintext := []byte("黑龙江省鸡西市躻鋁路3356号秘紤小区8单元2381室")
	for f.p = 1; f.p <= f.n; f.p++ {
		f.aesGcm(plaintext)
	}

	f.isDone = true
}

func (f *aesGcmTest) aesGcm(plaintext []byte) (ciphertext []byte) {
	block, err := aes.NewCipher(f.key)
	if err != nil {
		panic(err.Error())
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return c.Seal(nil, f.nonce, plaintext, nil)
}

func randBytes(n int) []byte {
	key := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}
	return key
}
