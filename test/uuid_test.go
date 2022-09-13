package test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUUID(t *testing.T){
	t.Log(uuid.NewV4())
}
