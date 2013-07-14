/**
 * Created with IntelliJ IDEA.
 * User: sunlzx
 * Date: 13-7-14
 * Time: 下午4:37
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

import (
	"mylib"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	mylib.Hello()
	fmt.Println(os.Getpid())
	fmt.Println(os.Getuid())

}
