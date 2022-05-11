package http

import (
	"reflect"
	"testing"
)

func TestHttpGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				url: "https://127.0.0.1:8004/api/v1/node",
			},
			want:   []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HttpGet(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
