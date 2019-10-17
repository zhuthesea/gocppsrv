/*
 * @Author:          _
 *                 / '_)  LONG!
 *  *        _.----._/  /
 *  *      /          /
 *  *    _/  (   | ( |
 *  *   /__.-|_|--|_l
 * @since: 2019-10-17 17:21:08
 * @lastTime: 2019-10-17 17:30:13
 * @LastAuthor: LONG
 * @Description:
 */
package gocpp

/*
#include "util.h"
#include "cwrap.h"
*/
import "C"

func GoSum(a, b int) int {
	s := C.sum(C.int(a), C.int(b))

	return int(s)
}

func GoCallCpp() {
	C.call()
}
