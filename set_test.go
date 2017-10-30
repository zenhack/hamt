package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	NewSet()
}

func TestSetInsert(t *testing.T) {
	s := NewSet()

	for i := 0; i < iterations; i++ {
		e := EntryInt(rand.Int31())
		s = s.Insert(e)
		assert.True(t, s.Include(e))
	}
}

func TestSetOperations(t *testing.T) {
	s := NewSet()

	for i := 0; i < iterations; i++ {
		assert.Equal(t, s.hamt.Size(), s.Size())

		e := EntryInt(rand.Int31() % 256)
		var ss Set

		if rand.Int()%2 == 0 {
			ss = s.Insert(e)

			assert.True(t, ss.Include(e))

			if s.Include(e) {
				assert.Equal(t, s.Size(), ss.Size())
			} else {
				assert.Equal(t, s.Size()+1, ss.Size())
			}
		} else {
			ss = s.Delete(e)

			assert.False(t, ss.Include(e))

			if s.Include(e) {
				assert.Equal(t, s.Size()-1, ss.Size())
			} else {
				assert.Equal(t, s.Size(), ss.Size())
			}
		}

		s = ss
	}
}

func TestSetFirstRest(t *testing.T) {
	s := NewSet()
	e, ss := s.FirstRest()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, ss.Size())

	s = s.Insert(EntryInt(42))
	e, ss = s.FirstRest()

	assert.Equal(t, EntryInt(42), e)
	assert.Equal(t, 0, ss.Size())

	s = s.Insert(EntryInt(2049))
	size := s.Size()

	for i := 0; i < size; i++ {
		e, s = s.FirstRest()

		assert.NotEqual(t, nil, e)
		assert.Equal(t, 1-i, s.Size())
	}
}

func TestSetSize(t *testing.T) {
	assert.Equal(t, 0, NewSet().Size())
}

func BenchmarkSetInsert(b *testing.B) {
	s := NewSet()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			s = s.Insert(EntryInt(i))
		}
	}
}

func BenchmarkSetSize(b *testing.B) {
	s := NewSet()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			s = s.Insert(EntryInt(i))
			b.Log(s.Size())
		}
	}
}
