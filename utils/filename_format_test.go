package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameFormat_GetName(t *testing.T) {
	as := assert.New(t)
	ns := NameFormat(`{{date|2006-01-02}}`)
	as.Equal(ns.GetName(), "2019-07-02")
}

func TestNameFormat_GetName1(t *testing.T) {
	as := assert.New(t)
	ns := NameFormat(`{{md5}}`)
	as.Equal(len(ns.GetName()), 32)
}
