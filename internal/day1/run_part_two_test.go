package day1

import "testing"

func Test_findFirstNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args{"219"}, "2", false},
		{args{"abc"}, "", true},
		{args{"two1nine"}, "2", false},
		{args{"1nine"}, "1", false},
		{args{"abcnine"}, "9", false},
		{args{"eightwo"}, "8", false},
		{args{"3h"}, "3", false},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := findFirstNumber(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFirstNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findFirstNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args{"219"}, "9", false},
		{args{"abc"}, "", true},
		{args{"two1nine"}, "9", false},
		{args{"1nine"}, "9", false},
		{args{"abcnine"}, "9", false},
		{args{"eightwo"}, "2", false},
		{args{"3h"}, "3", false},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := findLastNumber(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findLastNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findLastNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCalibrationNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    int
		wantErr bool
	}{
		{args{"1"}, 11, false},
		{args{"one1"}, 11, false},
		{args{"21threeabc44nine"}, 29, false},
		{args{"eightwo"}, 82, false},
		{args{"one"}, 11, false},
		{args{""}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := parseCalibrationNumber(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCalibrationNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseCalibrationNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
