package logging

import "testing"

func TestLog(t *testing.T) {
	log := NewConsoleLogger("test")
	log.Info("???")
}
