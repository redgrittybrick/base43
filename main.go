package main

import (
	"encoding/hex"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Command line options
	var hexFlag = flag.Bool("hex", false, 
		"Decode to / encode from HEX (otherwise BINARY)")
	var decodeFlag = flag.Bool("decode", false, 
		"DECODE from base43 (otherwise ENCODEs to base43)")
	var verboseFlag = flag.Bool("v", false, "Be verbose")
	flag.Parse()

	if *verboseFlag {
		l := log.New(os.Stderr, "", 0)
		other := "binary"
		if *hexFlag {
			other = "hexadecimal"
		}
		if *decodeFlag {
			l.Printf("Decoding from Base43 to %s\n", other)
		} else {
			l.Printf("Encoding from %s to Base43\n", other)
		}
	}

	// Read data from STDIN
	data := readAllStdIn()

	b43tool := base43{}
	// Decide if data is to be encoded or to be decoded
	if *decodeFlag {  // DECODE FROM BASE43 to binary or hex
		decoded, err := b43tool.Decode(data)
		check(err)
		if *hexFlag {
			fmt.Println(hex.EncodeToString(decoded))
		} else {
			fmt.Print(string(decoded))
		}
	} else {  // ENCODE from binary or hex TO BASE43
		if *hexFlag {
			tmp := make([]byte, hex.DecodedLen(len(data)))
			_, err := hex.Decode(tmp, data)
			check(err)
			fmt.Print(string(b43tool.Encode(tmp)))
		} else {
			fmt.Print(string(b43tool.Encode(data)))
		}
	}

}

// Write error message and exit
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Reads from STDIN into a byte slice until EOF detected.
func readAllStdIn() []byte {
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		buf = buf[:n]
		nChunks++
		nBytes += int64(len(buf))
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	return buf
}
