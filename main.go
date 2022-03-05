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
	cryptPasswd, err := crypt.Generate([]byte(passwd), []byte("$5$h34d3mpty"))
	if err != nil {
		panic(err)
	}

	// Verify password
	err = crypt.Verify(cryptPasswd, []byte(passwd))
	if err != nil {
		panic(err)
	}
	// Output to terminal
	fmt.Println(cryptPasswd)

	// Require env var to be set before editing /etc/shadow
	if os.Getenv("REPLACE_IN_SHADOW") != "" {
		// Open file
		f, err := os.Open("/etc/shadow")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// Set up variables
		var lines []string
		var lineOfInterest int = -1
		var currentLine int = 0

		// Create new scanner from the file
		scanner := bufio.NewScanner(f)
		// Read line-by-line as long as there are more lines
		for scanner.Scan() {
			// Get line and append it
			line := scanner.Text()
			lines = append(lines, line)
			// Check if it includes the username
			if strings.Contains(line, usr) {
				// Set line of interest if we find it
				lineOfInterest = currentLine
				break
			}
			currentLine++
		}
		// The scanner has exited, check if it was due to an error
		err = scanner.Err()
		if err != nil {
			panic(err)
		}

		// Check if we did find the correct line with the username
		if lineOfInterest > 0 {
			fmt.Println("line to look at:", lines[lineOfInterest])
		} else {
			panic("Line with username " + usr + " not found")
		}

		// Split line by colon separator
		// This results in 9 fields but there really are only 8
		fields := strings.Split(lines[lineOfInterest], ":")
		// Set field 2 to the encrypted password
		fields[1] = cryptPasswd
		// Create the edited line by reassembling fields
		newLine := strings.Join(fields, ":")
		// Replace the old line with the new one
		lines[lineOfInterest] = newLine

		// Write lines to file
		writer := bufio.NewWriter(f)
		for _, line := range lines {
			_, err := writer.WriteString(line)
			if err != nil {
				panic("the write was short:" + err.Error())
			}
		}
		writer.Flush()
	}
}
