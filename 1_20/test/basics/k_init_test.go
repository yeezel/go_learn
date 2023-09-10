package test

import (
	"golearn/basics"
	"testing"
)

func TestInitDemo(t *testing.T) {
	basics.InitDemo()
	t.Log(" 10 must be even!")
}
