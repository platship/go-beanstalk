/*
 * @Author: coller
 * @Date: 2023-01-29 13:00:23
 * @LastEditTime: 2023-01-29 13:00:32
 * @Desc:
 */
package beanstalk

import (
	"strconv"
	"time"
)

type dur time.Duration

func (d dur) String() string {
	return strconv.FormatInt(int64(time.Duration(d)/time.Second), 10)
}
