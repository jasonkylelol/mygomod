package mygomod

import "testing"

func Test_generics(t *testing.T) {
	t.Run("", func(t *testing.T) {
		classicSum()
		genericSum()
	})
}
