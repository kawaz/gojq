package gojq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		input    interface{}
		expected interface{}
		err      string
	}{
		{
			name:     "number",
			query:    `.`,
			input:    128,
			expected: 128,
		},
		{
			name:     "string",
			query:    `.`,
			input:    "foo",
			expected: "foo",
		},
		{
			name:     "object",
			query:    `.`,
			input:    map[string]interface{}{"foo": 128},
			expected: map[string]interface{}{"foo": 128},
		},
		{
			name:     "array",
			query:    `.`,
			input:    []interface{}{"foo", 128},
			expected: []interface{}{"foo", 128},
		},
		{
			name:     "object index",
			query:    `.foo`,
			input:    map[string]interface{}{"foo": 128},
			expected: 128,
		},
		{
			name:     "object member",
			query:    `.["foo"]`,
			input:    map[string]interface{}{"foo": 128},
			expected: 128,
		},
		{
			name:  "expected object",
			query: `.foo|.bar`,
			input: map[string]interface{}{"foo": 128},
			err:   "expected an object but got: int",
		},
		{
			name:     "object optional",
			query:    `.foo|.bar?`,
			input:    map[string]interface{}{"foo": 128},
			expected: struct{}{},
		},
		{
			name:     "array index",
			query:    `.[2]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: 48,
		},
		{
			name:     "array index out of bound",
			query:    `. [ 4 ]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: nil,
		},
		{
			name:     "array slice start",
			query:    `.[2:]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: []interface{}{48, 64},
		},
		{
			name:     "array slice end",
			query:    `.[:2]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: []interface{}{16, 32},
		},
		{
			name:     "array slice start end",
			query:    `.[1:3]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: []interface{}{32, 48},
		},
		{
			name:     "array slice all",
			query:    `.[:]`,
			input:    []interface{}{16, 32, 48, 64},
			expected: []interface{}{16, 32, 48, 64},
		},
		{
			name:  "expected array",
			query: `.[0]`,
			input: map[string]interface{}{"foo": 128},
			err:   "expected an array but got: map",
		},
		{
			name:  "pipe",
			query: `.foo | . | .baz | .[1]`,
			input: map[string]interface{}{
				"foo": map[string]interface{}{
					"baz": []interface{}{"Hello", "world"},
				},
			},
			expected: "world",
		},
	}

	for _, tc := range testCases {
		parser := NewParser()
		t.Run(tc.name, func(t *testing.T) {
			query, err := parser.Parse(tc.query)
			assert.NoError(t, err)
			require.NoError(t, err)
			got, err := Run(query, tc.input)
			if err == nil {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, got)
			} else {
				assert.NotEqual(t, tc.err, "")
				require.Contains(t, err.Error(), tc.err)
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}
