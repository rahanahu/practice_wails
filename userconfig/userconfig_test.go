package userconfig

import (
	"bytes"
	"testing"
)

func TestUserConfig_Saveconfig(t *testing.T) {
	tests := []struct {
		name    string
		a       *UserConfig
		wantW   string
		wantErr bool
	}{
		{
			name: "normal",
			a: &UserConfig{
				Dir: "testdir",
			},
			wantW:   `{"directory":"testdir"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.a.saveconfig(w); (err != nil) != tt.wantErr {
				t.Errorf("UserConfig.Saveconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("UserConfig.Saveconfig() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
