package slice

import (
	"testing"
)

func TestEqualsUnordered(t *testing.T) {
	type args struct {
		x interface{}
		y interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Slice of Strings",
			args: args{
				x: []string{"Hallo", "Test", "Hier"},
				y: []string{"Test", "Hallo", "Hier"},
			},
			want: true,
		},
		{
			name: "Slice of Unequal Types",
			args: args{
				x: []int{1, 2, 3},
				y: []string{"Test", "Hallo", "Hier"},
			},
			want: false,
		},
		{
			name: "Unequal Slice of Structs",
			args: args{
				x: []struct {
					ID   int
					Text string
				}{
					{
						ID:   4,
						Text: "Vier",
					}, {
						ID:   2,
						Text: "Zwei",
					}, {
						ID:   3,
						Text: "Drei",
					}},
				y: []struct {
					ID   int
					Text string
				}{
					{
						ID:   3,
						Text: "Drei",
					}, {
						ID:   2,
						Text: "Zwei",
					}, {
						ID:   1,
						Text: "Eins",
					}},
			},
			want: false,
		},
		{
			name: "Slice of Structs",
			args: args{
				x: []struct {
					ID   int
					Text string
				}{
					{
						ID:   1,
						Text: "Eins",
					}, {
						ID:   2,
						Text: "Zwei",
					}, {
						ID:   3,
						Text: "Drei",
					}},
				y: []struct {
					ID   int
					Text string
				}{
					{
						ID:   3,
						Text: "Drei",
					}, {
						ID:   2,
						Text: "Zwei",
					}, {
						ID:   1,
						Text: "Eins",
					}},
			},
			want: true,
		},
		{
			name: "Slice of Numbers",
			args: args{
				x: []int{1, 2, 3},
				y: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "Unequal Array of String",
			args: args{
				x: [...]string{"A", "B", "C"},
				y: [...]string{"J", "I", "H"},
			},
			want: false,
		},
		{
			name: "Array of String",
			args: args{
				x: [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
				y: [...]string{"J", "I", "H", "G", "F", "E", "D", "C", "B", "A"},
			},
			want: true,
		},
		{
			name: "Array of Numbers",
			args: args{
				x: [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				y: [...]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "Unequal array of numbers",
			args: args{
				x: [...]int{1, 2, 3},
				y: [...]int{4, 5, 6},
			},
			want: false,
		},
		{
			name: "Nil with something else",
			args: args{
				x: nil,
				y: [...]int{2, 1, 3, 0, 4, 5, 7, 6, 8, 9},
			},
			want: false,
		},
		{
			name: "String comparison",
			args: args{
				x: "X",
				y: "Y",
			},
			want: false,
		},
		{
			name: "Empty arrays, same type",
			args: args{
				x: []int{},
				y: []int{},
			},
			want: true,
		},
		{
			name: "Empty arrays, different type",
			args: args{
				x: []string{},
				y: []int{},
			},
			want: false,
		},
		{
			name: "Different arrays",
			args: args{
				x: [...]int{2, 1, 3, 0, 4, 5, 7, 6, 8},
				y: [...]int{2, 1, 3, 0, 4, 5, 7, 6, 8, 9},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualsUnordered(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("EqualsUnordered() = %v, want %v", got, tt.want)
			}
		})
	}
}
