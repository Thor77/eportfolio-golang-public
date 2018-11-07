package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const input = "golang"

func main() {
	var password string
	for i := 0; len(password) < 4; i++ {
		merged := input + strconv.Itoa(i)
		hashed := md5.Sum([]byte(merged))
		hexed := hex.EncodeToString(hashed[:])
		if strings.HasPrefix(hexed, "000") {
			password += string(hexed[3])
		}
	}
	fmt.Printf("Password: %s\n", password)
}
