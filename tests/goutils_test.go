package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xq-120/goUtils/pkg"
)

// TestFormatBytes 测试字节格式化函数
func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{1023, "1023 B"},
		{1024, "1.00 KB"},
		{1048576, "1.00 MB"},
		{1073741824, "1.00 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := pkg.FormatBytes(tt.bytes)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestAdd(t *testing.T) {
	stack := pkg.New[string]()
	stack.Push("hello")
	stack.Push("world")
	assert.Equal(t, 2, stack.Count(), "Stack should contain 2 elements")
	assert.Equal(t, "world", stack.Pop(), "Top element should be 'world'")
	assert.Equal(t, 1, stack.Count(), "Stack should contain 1 element after pop")
	stack.Push("foo")
	stack.Push("bar")
	assert.Equal(t, 3, stack.Count(), "Stack should contain 3 elements after pushing 'foo' and 'bar'")
	assert.Equal(t, "bar", stack.Pop(), "Top element should be 'bar'")
	assert.Equal(t, 2, stack.Count(), "Stack should contain 2 elements after pop")
	assert.Equal(t, "foo", stack.Pop(), "Top element should be 'foo'")
	assert.Equal(t, 1, stack.Count(), "Stack should contain 1 element after pop")
	assert.Equal(t, "hello", stack.Pop(), "Top element should be 'hello'")
	assert.Equal(t, 0, stack.Count(), "Stack should be empty after popping all elements")
	assert.Equal(t, "", stack.Pop(), "Popping from an empty stack should return an empty string")
	assert.Equal(t, true, stack.IsEmpty(), "Stack should be empty")
	stack.Push("test")
	assert.Equal(t, 1, stack.Count(), "Stack should contain 1 element after pushing 'test'")
	assert.Equal(t, "test", stack.Pop(), "Top element should be 'test'")
	assert.Equal(t, 0, stack.Count(), "Stack should be empty after popping 'test'")
	assert.Equal(t, true, stack.IsEmpty(), "Stack should be empty")
	stack.PushItems([]string{"a", "b", "c"})
	assert.Equal(t, 3, stack.Count(), "Stack should contain 3 elements after pushing items")
	assert.Equal(t, "c", stack.Pop(), "Top element should be 'c'")
	assert.Equal(t, 2, stack.Count(), "Stack should contain 2 elements after pop")
	assert.Equal(t, "b", stack.Pop(), "Top element should be 'b'")
	assert.Equal(t, 1, stack.Count(), "Stack should contain 1 element after pop")
	assert.Equal(t, "a", stack.Pop(), "Top element should be 'a'")
	assert.Equal(t, 0, stack.Count(), "Stack should be empty after popping all elements")
	assert.Equal(t, "", stack.Pop(), "Popping from an empty stack should return an empty string")
	assert.Equal(t, true, stack.IsEmpty(), "Stack should be empty")
}
