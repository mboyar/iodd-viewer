package main

import (
	"mime/multipart"
	"reflect"
	"testing"
)

func Test_parseIODDbyFile(t *testing.T) {
	type args struct {
		file multipart.File
	}
	tests := []struct {
		name string
		args args
		want IODevice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseIODDbyFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIODDbyFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
