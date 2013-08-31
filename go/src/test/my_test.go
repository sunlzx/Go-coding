/**
 * Created with IntelliJ IDEA.
 * User: sunlzx
 * Date: 13-7-20
 * Time: 上午8:30
 * To change this template use File | Settings | File Templates.
 */
package test

import (
	"testing"
	"fmt"
//	"crypto/sha1"
	"crypto/md5"
	"io"
)

func test1() {
	fmt.Println("testiong")
}

func TestAdd(t *testing.T) {
	fmt.Println("hello go")
}


func TestCrypto(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")

	Bytes := h.Sum(nil)

	str := string(Bytes)

	fmt.Printf("%x", Bytes)

	fmt.Println(str)
}

