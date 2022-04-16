package clock
import "fmt"
// Define the Clock type here.
type Clock struct {
    hours int
    minutes int
}
func Round(c Clock) Clock {
    m := c.minutes
    h := c.hours
    if m >= 60 || m < 0 {
        addHours := m / 60
        m %= 60
        if m < 0 {
            m += 60
            addHours--
        }
        h += addHours
    }
    if h >= 24 || h < 0 {
        h %= 24
        if h < 0 {
            h += 24
        }
    }
	return Clock{h,m}
}
func New(h, m int) Clock {
	return Round(Clock{h,m})
}
func (c Clock) Add(m int) Clock {
	c.minutes += m
	return Round(c)
}
func (c Clock) Subtract(m int) Clock {
	c.minutes -= m
    return Round(c)
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}