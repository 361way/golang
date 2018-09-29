/*
code by www.361way.com,from golang.org
itybku@139.com
AES CFB加解密
*/
package main

import (
        "crypto/aes"
        "crypto/cipher"
        "crypto/rand"
        "encoding/hex"
        "fmt"
        "io"
)

func ExampleNewCFBDecrypter() {
        key, _ := hex.DecodeString("6368616e676520746869732070617373")
        ciphertext, _ := hex.DecodeString("e38932f30048f4cf2ecff113b29c4aed3dc0fb65c8d16ae0171aee54d207")

        block, err := aes.NewCipher(key)
        if err != nil {
                panic(err)
        }

        if len(ciphertext) < aes.BlockSize {
                panic("ciphertext too short")
        }
        iv := ciphertext[:aes.BlockSize]
        ciphertext = ciphertext[aes.BlockSize:]

        stream := cipher.NewCFBDecrypter(block, iv)

        stream.XORKeyStream(ciphertext, ciphertext)
        fmt.Printf("%s\n", ciphertext)
}

func ExampleNewCFBEncrypter() {
        key, _ := hex.DecodeString("6368616e676520746869732070617373")
        plaintext := []byte("www.361way.com")

        block, err := aes.NewCipher(key)
        if err != nil {
                panic(err)
        }

        ciphertext := make([]byte, aes.BlockSize+len(plaintext))
        iv := ciphertext[:aes.BlockSize]
        if _, err := io.ReadFull(rand.Reader, iv); err != nil {
                panic(err)
        }

        stream := cipher.NewCFBEncrypter(block, iv)
        stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

        fmt.Printf("%x\n", ciphertext)
}

func main() {
        ExampleNewCFBDecrypter()
        ExampleNewCFBEncrypter()
}
