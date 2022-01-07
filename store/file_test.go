package store

import (
	"strconv"
	"testing"
)

func TestFile(t *testing.T) {
	path := "../test_data/data.log"
	f := OpenFile(path)
	for i := 0; i < 10; i++ {
		f.WriteString("hello" + strconv.FormatInt(int64(i), 10))
	}
}

func TestCreateFixedFile(t *testing.T) {
	path := "../test_data/fixedSize.log"
	CreateFixedSizeFile(path, 10)
}
