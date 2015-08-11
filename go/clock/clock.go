package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hour   int
	minute int
}

func Time(hour, minute int) Clock {
	c := Clock{hour, minute}

	switch {
	case c.minute >= 60:
		for c.minute >= 60 {
			c.hour++
			c.minute -= 60
		}
	case c.minute < 0:
		for c.minute < 0 {
			c.hour--
			c.minute += 60
		}
	}

	if c.hour < 0 {
		c.hour = 24 + (c.hour % 24)
	} else {
		c.hour = c.hour % 24
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%.02d:%.02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return Time(c.hour, c.minute+minutes)
}
