package storage

import (
	"testing"
)

func TestNewAppend(t *testing.T) {
	ap, err := NewAppend("/tmp/test-binlog")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if err := Close(); err != nil {
		t.Error(err)
		t.Fail()
	}
}
