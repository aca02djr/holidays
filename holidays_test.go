package holidays

import (
	"testing"
	"time"
)

func TestForYear(t *testing.T) {
	tests := []struct {
		Year     int
		Holidays []time.Time
	}{
		{2010, []time.Time{
			time.Date(2010, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.April, 2, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.April, 5, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.May, 3, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.May, 31, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.August, 30, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.December, 27, 0, 0, 0, 0, time.Local),
			time.Date(2010, time.December, 28, 0, 0, 0, 0, time.Local),
		}},
		{2011, []time.Time{
			time.Date(2011, time.January, 3, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.April, 22, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.April, 25, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.May, 2, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.May, 30, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.August, 29, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.December, 26, 0, 0, 0, 0, time.Local),
			time.Date(2011, time.December, 27, 0, 0, 0, 0, time.Local),
		}},
		{2012, []time.Time{
			time.Date(2012, time.January, 2, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.April, 6, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.April, 9, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.May, 7, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.June, 4, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.June, 5, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.August, 27, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2012, time.December, 26, 0, 0, 0, 0, time.Local),
		}},
		{2013, []time.Time{
			time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.March, 29, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.April, 1, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.May, 6, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.May, 27, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.August, 26, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2013, time.December, 26, 0, 0, 0, 0, time.Local),
		}},
		{2014, []time.Time{
			time.Date(2014, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.April, 18, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.April, 21, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.May, 5, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.May, 26, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.August, 25, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2014, time.December, 26, 0, 0, 0, 0, time.Local),
		}},
		{2015, []time.Time{
			time.Date(2015, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.April, 3, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.April, 6, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.May, 4, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.May, 25, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.August, 31, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2015, time.December, 28, 0, 0, 0, 0, time.Local),
		}},
		{2016, []time.Time{
			time.Date(2016, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.March, 25, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.March, 28, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.May, 2, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.May, 30, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.August, 29, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.December, 26, 0, 0, 0, 0, time.Local),
			time.Date(2016, time.December, 27, 0, 0, 0, 0, time.Local),
		}},
		{2017, []time.Time{
			time.Date(2017, time.January, 2, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.April, 14, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.April, 17, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.May, 1, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.May, 29, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.August, 28, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2017, time.December, 26, 0, 0, 0, 0, time.Local),
		}},
		{2018, []time.Time{
			time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.March, 30, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.April, 2, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.May, 7, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.May, 28, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.August, 27, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.December, 25, 0, 0, 0, 0, time.Local),
			time.Date(2018, time.December, 26, 0, 0, 0, 0, time.Local),
		}},
	}

	for _, test := range tests {
		actual := ForYear(test.Year)
		if len(actual) != len(test.Holidays) {
			t.Errorf("For year %d: expected result length %d, got %d", test.Year, len(test.Holidays), len(actual))
			continue
		}
		for i, hol := range actual {
			if !hol.Equal(test.Holidays[i]) {
				t.Errorf("For year %d: expected %s, got %s", test.Year, test.Holidays[i], hol)
			}
		}
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		Date     time.Time
		Expected bool
	}{
		{time.Date(2012, time.January, 1, 17, 0, 0, 0, time.Local), false},
		{time.Date(2012, time.January, 2, 14, 32, 11, 0, time.Local), true},
		{time.Date(2012, time.May, 28, 0, 0, 0, 0, time.Local), false},
		{time.Date(2012, time.June, 5, 0, 0, 0, 0, time.Local), true},
	}

	for _, test := range tests {
		actual := Check(test.Date)
		if actual != test.Expected {
			t.Errorf("Check %s: expected %t, got %t", test.Date, test.Expected, actual)
		}
	}
}
