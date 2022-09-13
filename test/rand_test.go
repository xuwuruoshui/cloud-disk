package test

import (
	"cloud-disk/core/helper"
	"testing"
)

func TestRand(t *testing.T) {
	t.Log(helper.RandCode())
}