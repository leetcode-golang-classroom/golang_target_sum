package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	for idx := 0; idx < b.N; idx++ {
		findTargetSumWays(nums, target)
	}
}
func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [1,1,1,1,1], target = 3",
			args: args{nums: []int{1, 1, 1, 1, 1}, target: 3},
			want: 5,
		},
		{
			name: "nums = [1], target = 1",
			args: args{nums: []int{1}, target: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
