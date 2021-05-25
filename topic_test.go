package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateSet(t *testing.T) {
	tree := NewTopicTree()
	tree.Add("a", "a")
	assert.Equal(t, "a", tree.Get("a"), "suitable save")
	tree.Add("a", "b")
	assert.Equal(t, "a", tree.Get("a"), "suitable save")
}
func TestRemoveSet(t *testing.T) {
	tree := NewTopicTree()
	tree.Add("a", "a")
	assert.Equal(t, "a", tree.Get("a"), "suitable save")
	tree.Add("a", "b")
	assert.Equal(t, "a", tree.Get("a"), "suitable save")
	tree.Remove("a")
	assert.Equal(t, nil, tree.Get("a"), "suitable save")
	tree.Add("a", "b")
	assert.Equal(t, "b", tree.Get("a"), "suitable save")
}
