package foo

import (
	"github.com/learning-go-book/internal_example/foo/internal"
	"github.com/learning-go-book/internal_example/foo/sibling"
)

func UseDoubler(i int) int {
	sibling.AlsoUseDoubler(i)
	return internal.Doubler(i)
}
