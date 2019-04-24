package main

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	// PM9SCREW is a flag
	PM9SCREW = "\tPM9SCREW\t"
)

var (
	pm9ScrewMyCryptKey []byte
)

func main() {
	key := flag.String("k", "", "pm9 screw crypt key")
	f := flag.String("f", "", "screwed file")
	screw := flag.String("screw", PM9SCREW, "screw")
	flag.Parse()
	fmt.Printf("%q\n", *screw)

	var err error
	k2 := *key
	if k2 == "" {
		k2, err = loadKey()
		if err != nil {
			log.Fatal(err)
		}
	}

	pm9ScrewMyCryptKey, err = parseKey(k2)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}
	data, err := decrypt(bs, *screw, pm9ScrewMyCryptKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func loadKey() (string, error) {
	bs, err := ioutil.ReadFile(os.ExpandEnv(`$HOME/.php-unscrew`))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func parseKey(s string) ([]byte, error) {
	bs, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	for i, x := range bs {
		if i&1 == 0 {
			buf.WriteByte(x)
		}
	}
	return buf.Bytes(), nil
}

func decrypt(data []byte, screw string, key []byte) ([]byte, error) {
	if len(data) < len(screw) {
		return nil, fmt.Errorf("invalid data")
	}

	data = data[len(screw):]
	var buf bytes.Buffer
	for i, x := range data {
		r := key[(len(data)-i)%len(key)] ^ (^x)
		if err := buf.WriteByte(r); err != nil {
			return nil, err
		}
	}
	r, err := zlib.NewReader(&buf)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
