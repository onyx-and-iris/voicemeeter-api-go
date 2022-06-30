package voicemeeter_test

import (
	"os"
	"testing"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

var (
	vmRem = voicemeeter.GetRemote("potato")
)

func TestMain(m *testing.M) {
	vmRem.Login()
	code := m.Run()
	vmRem.Logout()
	os.Exit(code)
}

func sync() {
	time.Sleep(30 * time.Millisecond)
	for vmRem.Pdirty() || vmRem.Mdirty() {
	}
}
