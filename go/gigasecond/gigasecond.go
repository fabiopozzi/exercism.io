package gigasecond

// Write a function AddGigasecond that works with time.Time.
// Also define a variable Birthday set to your (or someone else's) birthday.
// Run go test -v to see your gigasecond anniversary.

import "time"

const TestVersion = 2

const shortForm = "2006-01-02"

var Birthday, _ = time.Parse(shortForm, "1984-09-10")

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
