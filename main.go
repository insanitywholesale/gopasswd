package main

import (
	"fmt"
	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
	"os"
)

func main() {
	passwd := os.Getenv("PASSWD")
	if passwd == "" {
		passwd = "fail"
	}
	crypt := crypt.SHA256.New()
	ret, _ := crypt.Generate([]byte(passwd), []byte("$5$h34d3mpty"))

	err := crypt.Verify(ret, []byte(passwd))
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}
