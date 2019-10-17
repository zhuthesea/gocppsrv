/*
 * @Author:          _
 *                 / '_)  LONG!
 *  *        _.----._/  /
 *  *      /          /
 *  *    _/  (   | ( |
 *  *   /__.-|_|--|_l
 * @since: 2019-10-17 17:21:08
 * @lastTime: 2019-10-17 17:30:40
 * @LastAuthor: LONG
 * @Description:
 */
package main

import (
	"fmt"
	"net/http"
	"strconv"

	gocppfunctions "github.com/zhuthesea/gocppsrv/gocpp"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	var rs int = gocppfunctions.GoSum(4, 1)
	w.Write([]byte("helloWorld -- from cpp return " + strconv.Itoa(rs)))
}

func main() {
	gocppfunctions.GoCallCpp()

	var rs int = gocppfunctions.GoSum(4, 2)
	fmt.Println(rs)

	http.HandleFunc("/helloWorld", helloWorld)
	http.ListenAndServe(":8001", nil)

}
