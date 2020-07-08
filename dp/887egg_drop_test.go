package main

import "testing"

func Test_eggDrop(t *testing.T) {
	type args struct {
		eggsCount  int
		floorCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				eggsCount:  2,
				floorCount: 100,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop1(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("superEggDrop1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop2(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("superEggDrop2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop3(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("superEggDrop3() = %v, want %v", got, tt.want)
			}
		})
	}
}
