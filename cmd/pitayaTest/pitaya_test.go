package pitaya_test

import (
	pitaya "mygotraining/cmd/pitayaTest"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	port := pitaya.GetFreePort(t)
	t.Log(port)
}
