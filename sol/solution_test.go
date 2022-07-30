package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	intervals := []*Interval{{0, 30}, {5, 10}, {15, 20}}
	for idx := 0; idx < b.N; idx++ {
		CanAttendMeetings(intervals)
	}
}
func TestCanAttendMeetings(t *testing.T) {
	type args struct {
		intervals []*Interval
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "intervals = [(0,30),(5,10),(15,20)]",
			args: args{intervals: []*Interval{{0, 30}, {5, 10}, {15, 20}}},
			want: false,
		},
		{
			name: "intervals = [(5,8),(9,15)]",
			args: args{intervals: []*Interval{{5, 8}, {9, 15}}},
			want: true,
		},
		{
			name: "intervals = []",
			args: args{intervals: []*Interval{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("CanAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
