package main

import (
	"bufio"
	"fmt"
	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
	"os"
	"strings"
)

func main() {
	// Define desired password
	passwd := os.Getenv("PASSWD")
	if passwd == "" {
		passwd = "fail"
	}
	// Define user who's password to change
	usr := os.Getenv("USR")
	if usr == "" {
		usr = "angle"
	}

	// Create new SHA256 crypt.Crypter
	crypt := crypt.SHA256.New()
	// Generate password with salt
	ret, err := crypt.Generate([]byte(passwd), []byte("$5$h34d3mpty"))
	if err != nil {
		panic(err)
	}

	// Verify password
	err = crypt.Verify(ret, []byte(passwd))
	if err != nil {
		panic(err)
	}
	// Output to terminal
	fmt.Println(ret)

	// Require env var to be set before editing /etc/shadow
	if os.Getenv("REPLACE_IN_SHADOW") != "" {
		// Open file
		f, err := os.Open("/etc/shadow")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		// Read line-by-line as long as there are more lines
		for scanner.Scan() {
			// Get line
			line := scanner.Text()
			// Check if it includes the username
			if strings.Contains(line, usr) {
				// Split line by colon separator
				// This results in 9 fields but there really are only 8
				fields := strings.Split(line, ":")
				// TODO: just prints for now but make it alter the value of second field
				for fieldIndex, fieldValue := range fields {
					println(fieldIndex, fieldValue)
				}
			}
		}
		// The scanner has exited, check if it was due to an error
		err = scanner.Err()
		if err != nil {
			panic(err)
		}
	}
}
