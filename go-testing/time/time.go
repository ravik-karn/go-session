package time

import (
	"time"
)

var now = time.Now
func AfterOneMinute() time.Time {
	currTime := now()
	return currTime.Add(time.Second * 60)
}
