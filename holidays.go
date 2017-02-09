// Package holidays provides functions for calculating bank holidays in England & Wales for a given year.
package holidays

import (
	"sort"
	"time"
)

type byTimeAscending []time.Time

func (t byTimeAscending) Len() int           { return len(t) }
func (t byTimeAscending) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byTimeAscending) Less(i, j int) bool { return t[i].Before(t[j]) }

var exceptionsRemove = map[int][]time.Time{
	2012: {
		time.Date(2012, time.May, 28, 0, 0, 0, 0, time.Local), // Spring Bank Holiday moved due to Jubilee
	},
}

var exceptionsAdd = map[int][]time.Time{
	2012: {
		time.Date(2012, time.June, 4, 0, 0, 0, 0, time.Local), // Spring Bank Holiday moved due to Jubilee
		time.Date(2012, time.June, 5, 0, 0, 0, 0, time.Local), // Extra bank holiday for Jubilee
	},
}

// ForYear returns a time.Time slice representing the bank holidays for the provided year.
func ForYear(year int) []time.Time {
	holsMap := getHols(year)

	result := make([]time.Time, len(holsMap))
	i := 0
	for hol := range holsMap {
		result[i] = hol
		i++
	}

	sort.Sort(byTimeAscending(result))
	return result
}

// Check returns whether or not the provided time.Time is bank holiday. It ignores the time portion, checking only the date.
func Check(date time.Time) bool {
	holsMap := getHols(date.Year())
	// strip time
	d := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	_, ok := holsMap[d]
	return ok
}

func getHols(year int) map[time.Time]struct{} {
	yep := struct{}{}
	holsMap := make(map[time.Time]struct{})

	// new years day
	nyd := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	if nyd.Weekday() == time.Saturday {
		nyd = nyd.AddDate(0, 0, 2)
	} else if nyd.Weekday() == time.Sunday {
		nyd = nyd.AddDate(0, 0, 1)
	}
	holsMap[nyd] = yep

	// easter
	es := easterSunday(year)
	gf := es.AddDate(0, 0, -2)
	em := es.AddDate(0, 0, 1)
	holsMap[gf] = yep
	holsMap[em] = yep

	// early may
	holsMap[firstOfMonth(year, time.May, time.Monday)] = yep

	// spring
	holsMap[lastOfMonth(year, time.May, time.Monday)] = yep

	// summer
	holsMap[lastOfMonth(year, time.August, time.Monday)] = yep

	// xmas
	xmas := time.Date(year, time.December, 25, 0, 0, 0, 0, time.Local)
	if xmas.Weekday() == time.Friday {
		// boxing day is the following monday
		holsMap[xmas] = yep
		holsMap[xmas.AddDate(0, 0, 3)] = yep
	} else if xmas.Weekday() == time.Saturday {
		// xmas is the following monday, boxing day tuesday
		holsMap[xmas.AddDate(0, 0, 2)] = yep
		holsMap[xmas.AddDate(0, 0, 3)] = yep
	} else if xmas.Weekday() == time.Sunday {
		// xmas is the following monday, boxing day tuesday
		holsMap[xmas.AddDate(0, 0, 1)] = yep
		holsMap[xmas.AddDate(0, 0, 2)] = yep
	} else {
		holsMap[xmas] = yep
		holsMap[xmas.AddDate(0, 0, 1)] = yep
	}

	// exceptions
	for _, e := range exceptionsAdd[year] {
		holsMap[e] = yep
	}

	for _, e := range exceptionsRemove[year] {
		delete(holsMap, e)
	}

	return holsMap
}

// Calculate the date of Easter Sunday (https://en.wikipedia.org/wiki/Computus#Software - ported from Python function IanTaylorEasterJscr)
func easterSunday(year int) time.Time {
	a := year % 19
	b := year >> 2
	c := b/25 + 1
	d := (c * 3) >> 2
	e := ((a * 19) - ((c*8 + 5) / 25) + d + 15) % 30
	e += (29578 - a - e*32) >> 10
	e -= ((year % 7) + b - d + e + 2) % 7
	d = e >> 5
	day := e - d*31
	month := d + 3
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// Calculate the first occurrence of the given weekday in the given month/year
func firstOfMonth(year int, month time.Month, weekday time.Weekday) time.Time {
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	diff := int(weekday - d.Weekday())
	if diff < 0 {
		diff += 7
	}
	return d.AddDate(0, 0, diff)
}

// Calculate the last occurrence of the given weekday in the given month/year
func lastOfMonth(year int, month time.Month, weekday time.Weekday) time.Time {
	d := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	diff := int(weekday - d.Weekday())
	if diff >= 0 {
		diff -= 7
	}
	return d.AddDate(0, 0, diff)
}
