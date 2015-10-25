package http

import (
	j "../"
	"testing"
)

func Test_CanAssign(t *testing.T) {
	var client j.Client = NewClient()
	_ = client
}
