package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const iterations = 10000

func TestNewSet(t *testing.T) {
	NewSet()
}

func TestSetInsert(t *testing.T) {
	s := NewSet()

	for i := 0; i < iterations; i++ {
		e := EntryInt(rand.Int31())
		s = s.Insert(e)
		assert.Equal(t, e, s.Find(e))
	}
}

func TestSetOperations(t *testing.T) {
	s := NewSet()

	for i := 0; i < iterations; i++ {
		e := EntryInt(rand.Int31() % 256)

		if rand.Int()%2 == 0 {
			s = s.Insert(e)
			assert.Equal(t, e, s.Find(e))
		} else {
			s = s.Delete(e)
			assert.Equal(t, nil, s.Find(e))
		}
	}
}
