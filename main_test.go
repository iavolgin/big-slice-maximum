package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "Positive size",
			size: 5,
			want: 5,
		},
		{
			name: "Zero size",
			size: 0,
			want: 0,
		},
		{
			name: "Negative size",
			size: -5,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.want {
				t.Errorf("generateRandomElements() length = %d, want %d", len(got), tt.want)
				return
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		want    int
		wantErr bool
	}{
		{name: "Positive numbers", data: []int{1, 5, 3, 9, 2}, want: 9, wantErr: false},
		{name: "Single element", data: []int{42}, want: 42, wantErr: false},
		{name: "Empty slice", data: []int{}, want: 0, wantErr: true},
		{name: "Nil slice", data: nil, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.data)

			if tt.wantErr && err == nil {
				t.Errorf("maximum() expected error, but didn't get")
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}
