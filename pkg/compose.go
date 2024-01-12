// Composability in go
// This is the length and breadth of golaang

package pkg

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	//io.Reader
	bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: *bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func HashAndBroadCast(r HashReader) error {
	hash := r.hash()
	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	return err
	// }
	// hash := sha1.Sum(b)
	// fmt.Println(hash)
	fmt.Println(hash[:])

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes: ", string(b))

	return nil
}
