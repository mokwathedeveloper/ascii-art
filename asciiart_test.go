package main

import (
	"ascii-art/component"
	"reflect"
	"testing"
)

func TestMakeAsciiGraphicsMap(t *testing.T) {
	type args struct {
		fileContents []string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := component.MakeAsciiGraphicsMap(tt.args.fileContents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeAsciiGraphicsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTheCorrectFile(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := component.GetTheCorrectFile(tt.args.args); got != tt.want {
				t.Errorf("GetTheCorrectFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		str          string
		fileContents []string
		asciiMap     map[rune]int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component.Output(tt.args.str, tt.args.fileContents, tt.args.asciiMap)
		})
	}
}

func TestHandleError(t *testing.T) {
	type args struct {
		er error
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component.HandleError(tt.args.er)
		})
	}
}
