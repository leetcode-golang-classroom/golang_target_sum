package sol

import "testing"

func BenchmarkTestDFS(b *testing.B) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	for idx := 0; idx < b.N; idx++ {
		findTargetSumWaysDFS(nums, target)
	}
}
func Test_findTargetSumWaysDFS(t *testing.T) {
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
			if got := findTargetSumWaysDFS(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWaysDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
