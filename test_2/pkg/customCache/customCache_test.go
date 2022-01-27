package customCache

import (
	"testing"

	"github.com/houqp/gtest"
	"github.com/stretchr/testify/assert"
)

type CacheTests struct {
	TestedCache Cache
}

func (s *CacheTests) Setup(t *testing.T)      {}
func (s *CacheTests) Teardown(t *testing.T)   {}
func (s *CacheTests) BeforeEach(t *testing.T) {}
func (s *CacheTests) AfterEach(t *testing.T) {
	s.TestedCache.Clear()
}

// ======================================================= Cache Test Suite

// should get the correct value when previously set
func (s *CacheTests) SubTest1SimpleSetGet(t *testing.T) {
	expectedValue := "expected value"

	s.TestedCache.Set("test", expectedValue)

	assert.Equal(t, expectedValue, s.TestedCache.Get("test"))
}

// should get "nil" when the key is unknown
func (s *CacheTests) SubTest2GetUnknown(t *testing.T) {
	assert.Equal(t, nil, s.TestedCache.Get("test"))
}

// should get the correct value when previously set multiple times (rewrite)
func (s *CacheTests) SubTest3MultipleSetGet(t *testing.T) {
	expectedValue := "expected value"

	s.TestedCache.Set("test", "random value 1")
	s.TestedCache.Set("test", expectedValue)

	assert.Equal(t, expectedValue, s.TestedCache.Get("test"))
}

// should clear the cache correctly
func (s *CacheTests) SubTest4Clear(t *testing.T) {
	s.TestedCache.Set("test", "random value 1")

	s.TestedCache.Clear()

	assert.Equal(t, nil, s.TestedCache.Get("test"))
}

// Start test suite for each cache type
func TestGtest(t *testing.T) {
	fileCache, _ := NewFileCache("/tmp/customCacheTest")
	fileOnlyCache, _ := NewFileOnlyCache("/tmp/customCacheTest")

	gtest.RunSubTests(t, &CacheTests{TestedCache: NewMemCache()})
	gtest.RunSubTests(t, &CacheTests{TestedCache: fileCache})
	gtest.RunSubTests(t, &CacheTests{TestedCache: fileOnlyCache})
}
