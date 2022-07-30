package sol

import "sort"

type ByStart []*Interval

func (a ByStart) Len() int {
	return len(a)
}
func (a ByStart) Less(i, j int) bool {
	return a[i].start < a[j].start
}

func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */
/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */
func CanAttendMeetings(intervals []*Interval) bool {
	sort.Sort(ByStart(intervals))
	nIntervals := len(intervals)
	if nIntervals <= 1 {
		return true
	}
	overlapEnd := intervals[0].end
	for pos := 1; pos < nIntervals; pos++ {
		if overlapEnd > intervals[pos].start {
			return false
		} else {
			overlapEnd = intervals[pos].end
		}
	}
	return true
}
