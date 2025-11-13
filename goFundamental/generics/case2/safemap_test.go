package safemap_test

import (
	"golang/generics/case2/safemap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	sm := safemap.NewSafeMap[string, int]()
	sm.Set("key", 1)

	ans, err := sm.Get("key")
	assert.Equal(t, 1, ans)
	assert.NoError(t, err)
}
