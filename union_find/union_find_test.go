package union_find

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsvalidIndex(t *testing.T) {
	uFind, err := NewUnionFind(10)
	assert.Nil(t, err)

	for i := 0; i < 10; i++ {
		assert.True(t, uFind.isValidIndex(i))
	}
	assert.False(t, uFind.isValidIndex(-1))
	assert.False(t, uFind.isValidIndex(10))
	assert.False(t, uFind.isValidIndex(11))
}

func TestInitialization(t *testing.T) {
	uFind, err := NewUnionFind(-1)
	assert.NotNil(t, err)
	assert.Nil(t, uFind)

	n := rand.Intn(100) + 1
	uFind, err = NewUnionFind(n)
	assert.Nil(t, err)
	assert.NotNil(t, uFind)
	assert.Equal(t, n, len(uFind.parent))
	assert.Equal(t, n, len(uFind.size))
	assert.Equal(t, n, uFind.n)
	for i := 0; i < n; i++ {
		assert.Equal(t, 1, uFind.size[i])
		assert.Equal(t, i, uFind.parent[i])
	}
}

func TestUnion(t *testing.T) {
	uFind, _ := NewUnionFind(10)

	err := uFind.Union(-1, 2)
	assert.NotNil(t, err)

	err = uFind.Union(1, -2)
	assert.NotNil(t, err)

	err = uFind.Union(-1, -2)
	assert.NotNil(t, err)

	err = uFind.Union(5, 6)
	assert.Nil(t, err)
	uFind.Union(7, 8)
	uFind.Union(5, 8)
	uFind.Union(5, 5)
	uFind.Union(5, 9)

	group7, _ := uFind.Find(7)
	group9, _ := uFind.Find(9)
	assert.Equal(t, group7, group9)
}

func TestFind(t *testing.T) {
	uFind, _ := NewUnionFind(10)
	_, err := uFind.Find(-1)
	assert.NotNil(t, err)
	root5, err := uFind.Find(5)
	assert.Nil(t, err)
	assert.Equal(t, 5, root5)

	uFind.Union(7, 8)
	uFind.Union(5, 8)
	uFind.Union(5, 5)
	uFind.Union(5, 9)

	root5, err = uFind.Find(5)
	assert.Nil(t, err)
	assert.Equal(t, 7, root5)

	root9, err := uFind.Find(5)
	assert.Nil(t, err)
	assert.Equal(t, 7, root9)
}

func TestFindWithPathCompression(t *testing.T) {
	uFind, _ := NewUnionFind(10)
	uFind.Union(5, 6)
	uFind.Union(6, 7)
	uFind.Union(1, 2)
	uFind.Union(1, 5)
	assert.Equal(t, 5, uFind.findWithPathCompression(1))
	assert.Equal(t, 5, uFind.findWithPathCompression(2))
}
