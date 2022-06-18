package config

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetConfig(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *ServerConfig
		wantErr bool
	}{
		{
			name: "successful",
			args: args{
				reader: strings.NewReader(`{"port":"8080"}`),
			},
			want: &ServerConfig{
				Port: "8080",
			},
			wantErr: false,
		},
		{
			name: "fail case (wrong json syntax)",
			args: args{
				reader: strings.NewReader(`{port:8080}`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfig(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
