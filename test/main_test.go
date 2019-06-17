package test

import (
	"fmt"
	"github.com/dncmn/com"
	"os"
	"testing"
)

func TestOsRead(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(dir)
}

func TestConvert(t *testing.T) {
	str := "1234"
	v, err := com.StrTo(str).Int64()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v)
}

func TestURLEncode(t *testing.T) {
	str := "http://www.oschina.net/search?scope=bbs&q=C语言"
	rs := com.UrlEncode(str)
	fmt.Println(rs)
}
