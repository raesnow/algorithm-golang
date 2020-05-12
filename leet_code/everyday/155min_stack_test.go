package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	assert.Equal(t, -3, obj.GetMin(), "failed")

	obj.Pop()
	obj.Pop()

	assert.Equal(t, -2, obj.GetMin(), "failed")

	obj = Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)

	assert.Equal(t, 2147483647, obj.Top(), "failed")

	obj.Pop()

	assert.Equal(t, 2147483646, obj.GetMin(), "failed")

	obj.Pop()

	assert.Equal(t, 2147483646, obj.GetMin(), "failed")

	obj.Pop()
	obj.Push(2147483647)

	assert.Equal(t, 2147483647, obj.Top(), "failed")
	assert.Equal(t, 2147483647, obj.GetMin(), "failed")

	obj.Push(-2147483648)

	assert.Equal(t, -2147483648, obj.Top(), "failed")
	assert.Equal(t, -2147483648, obj.GetMin(), "failed")

	obj.Pop()

	assert.Equal(t, 2147483647, obj.GetMin(), "failed")
}
