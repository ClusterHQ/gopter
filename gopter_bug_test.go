package gopter_bug

import (
	"fmt"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

type Foo string

func TestPtrOfFoo(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("PtrOf", prop.ForAll(
		func(foo *Foo,
		) bool {
			return false
		},
		gen.PtrOf(GenFoo()),
	))
	properties.TestingRun(t)
}

func GenFoo() gopter.Gen {
	return gen.SliceOfN(16, gen.Rune()).Map(func(v interface{}) interface{} {
		asRunes, ok := v.([]rune)
		if !ok {
			panic("Oh craps")
		}
		fmt.Println(asRunes, Foo(asRunes))
		return Foo(asRunes)
	})
}
