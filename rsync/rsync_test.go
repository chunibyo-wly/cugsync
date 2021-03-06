package rsync

import (
	"os"
	"testing"
)

func TestRsyncCheck(t *testing.T) {
	hasRsync, info := CheckRsync()
	if !hasRsync {
		t.Errorf("rsync is not installed on this machine, stop.")
	}
	t.Logf("check rsync ok. version info: \n%s", info)
}

func TestExecRsync(t *testing.T) {
	defer os.RemoveAll("test_rsync")
	defer os.Remove("test_rsync.log")
	conf := InitConfig("./test_rsync", "/eclipse/swtchart/releases/0.12.0/release/", "mirrors.tuna.tsinghua.edu.cn", []string{"-avz", "--delete"})
	if err := ExecCommand(conf); err != nil {
		t.Error(err)
	}
}
