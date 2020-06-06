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
			if got := eggDrop(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("eggDrop() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := eggDrop1(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("eggDrop1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := eggDrop2(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("eggDrop2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := eggDrop3(tt.args.eggsCount, tt.args.floorCount); got != tt.want {
				t.Errorf("eggDrop3() = %v, want %v", got, tt.want)
			}
		})
	}
}
