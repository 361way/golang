package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

//func oneTimePassword(key []byte, value []byte) uint32 {
func oneTimePassword(key []byte, value []byte) string {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)
	pwd := number % 1000000

	//return pwd
	return strconv.Itoa(int(pwd))
}

// all []byte in this program are treated as Big Endian

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your key:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		os.Exit(1)
	}
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpaces = strings.Replace(inputNoSpaces, "\n", "", -1)
	fmt.Println(inputNoSpaces)
	fmt.Println(len(inputNoSpaces))
	fmt.Println(reflect.TypeOf(inputNoSpaces))

	key, err := base32.StdEncoding.DecodeString("N77D5BE6LHGVABEA")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(key, toBytes(epochSeconds/30))
	fmt.Println(pwd)
	fmt.Println(len(pwd))
	fmt.Println(reflect.TypeOf(pwd))

	secondsRemaining := 30 - (epochSeconds % 30)
	fmt.Printf("%06s (%d second(s) remaining)\n", pwd, secondsRemaining)
	if inputNoSpaces == pwd {
		fmt.Println("ok")
	} else {
		fmt.Println("false")
	}

}

