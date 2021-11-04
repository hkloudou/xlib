package xnet

import (
	"log"
	"testing"
)

func Test_GetPublicMainIP(t *testing.T) {
	log.Println(1)
	t.Log(GetLocalIP())
	log.Println(2)
}
