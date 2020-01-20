package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"log"
)

type Cipher struct {
	passwd  string
	content []byte
	key     []byte
	iv      []byte
}

func NewCipher(passwd string) *Cipher {
	return &Cipher{
		passwd: passwd,
	}
}

func (c *Cipher) Ecrypt(content []byte) []byte {
	if content == nil {
		fmt.Println("nil")
	}
	c.setKeyVi()
	block, err := aes.NewCipher(c.key)
	if err != nil {
		log.Fatalln(err)
	}
	text := make([]byte, len(content))
	stream := cipher.NewCFBEncrypter(block, c.iv)
	stream.XORKeyStream(text, content)
	return text
}

func (c *Cipher) Decrypt(ciphertext []byte) []byte {
	if ciphertext == nil {
		fmt.Println("nil")
	}
	c.setKeyVi()

	block, err := aes.NewCipher(c.key)
	if err != nil {
		fmt.Println("ciphertext too short", err)
	}

	//if len(ciphertext) < aes.BlockSize {
	//	fmt.Println("ciphertext too short")
	//	//panic("ciphertext too short")
	//}

	aesStream := cipher.NewCFBDecrypter(block, c.iv)
	aesStream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}

func (c *Cipher) setKeyVi() {
	key, iv := c.byteToKey(c.passwd, 32)
	c.key = key
	c.iv = iv
}

func (c *Cipher) byteToKey(passwd string, keylen int) ([]byte, []byte) {
	pass := []byte(passwd)
	prev := []byte{}
	key := []byte{}
	iv := []byte{}

	remain := 0
	for len(key) < keylen {
		hash := md5.Sum(append(prev, pass...))
		remain = keylen - len(key)
		if remain < 16 {
			key = append(key, hash[:remain]...)
		} else {
			key = append(key, hash[:]...)
		}
		prev = hash[:]
	}

	hash := md5.Sum(append(prev, pass...))
	if remain < 16 {
		iv = append(prev[remain:], hash[:remain]...)
	} else {
		iv = hash[:]
	}

	return key, iv
}
