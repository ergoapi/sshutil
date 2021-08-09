package sshutil

import (
	"testing"
)

func TestSSH_CmdAsync(t *testing.T) {
	type args struct {
		host string
		cmd  string
	}

	var (
		ssh = SSH{
			User:    "root",
			Pass:    "",
			PkFile:  "/Users/ysicing/.ssh/id_rsa",
			PkPass:  "",
			Timeout: nil,
			Debug:   true,
		}
	)

	tests := []struct {
		name    string
		fields  SSH
		args    args
		wantErr bool
	}{
		{
			"test ssh run w",
			ssh,
			args{
				host: "172.16.16.55",
				cmd:  "w",
			},
			false,
		},
		{
			"test ssh run w",
			ssh,
			args{
				host: "172.16.16.56",
				cmd:  "w",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SSH{
				User:    tt.fields.User,
				Pass:    tt.fields.Pass,
				PkFile:  tt.fields.PkFile,
				PkPass:  tt.fields.PkPass,
				Timeout: tt.fields.Timeout,
				Debug:   tt.fields.Debug,
			}
			if err := ss.CmdAsync(tt.args.host, tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("CmdAsync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
