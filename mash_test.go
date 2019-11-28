package uheprng

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMash(t *testing.T) {
	a := assert.New(t)
	mash := NewMash()

	a.Equal(0.4109745647292584, mash.Next("abcdefg"))
	mash.Init()
	a.Equal(0.8897808578331023, mash.Next("あいうえお"))
	mash.Init()
	a.Equal(0.4895239816978574, mash.Next("12345"))
}
