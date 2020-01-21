package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now = func() time.Time {
		return time.Date(2010, 10, 12, 12, 05, 0, 0, time.UTC)
	}
	expres :=  time.Date(2010, 10, 12, 12, 06, 0, 0, time.UTC)
	actres := AfterOneMinute()
	assert.Equal(t, expres, actres)
}
