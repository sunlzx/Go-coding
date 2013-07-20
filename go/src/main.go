/**
 * Created with IntelliJ IDEA.
 * User: sunlzx
 * Date: 13-7-14
 * Time: 下午4:37
 * To change this template use File | Settings | File Templates.
 */
package main


import (
	"fmt"
//	"time"
	"io"
	"net/http"
	"log"
)

var c = make(chan int)
var a string

func f() {
	a = "hello"

//	c <- 0
	<-c
	fmt.Print("f")

	c <- 0

}

func helloHandler(w http.ResponseWriter, r * http.Request) {
	io.WriteString(w, "hello world ! \n")
}

func main_bak() {
	http.HandleFunc("/", helloHandler)
	err  := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
