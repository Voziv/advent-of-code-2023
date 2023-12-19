package day1

import "testing"

func Test_findFirstDigit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args{"1"}, "1", false},
		{args{"one1"}, "1", false},
		{args{"21threeabc44nine"}, "2", false},
		{args{"one"}, "", true},
		{args{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := findFirstDigit(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFirstDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findFirstDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastDigit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args{"1"}, "1", false},
		{args{"one1"}, "1", false},
		{args{"21threeabc44nine"}, "4", false},
		{args{"one"}, "", true},
		{args{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := findLastDigit(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findLastDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findLastDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCalibrationDigits(t *testing.T) {
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
		{args{"21threeabc44nine"}, 24, false},
		{args{"one"}, 0, true},
		{args{""}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			got, err := parseCalibrationDigits(tt.args.input)
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
