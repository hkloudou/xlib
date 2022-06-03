package xcert

import (
	"log"
	"reflect"
	"testing"
)

func Test_x(t *testing.T) {
	// GenerateEcdsaCert(Template())
	// tmp := NewEcdsaP256()

	tmp := NewRsa256()
	// tmp.
	log.Println("type", reflect.TypeOf(tmp.Public()))
}
