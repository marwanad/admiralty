package delegatepod

import (
	"fmt"
	"testing"

	"admiralty.io/multicluster-scheduler/pkg/common"
	"github.com/google/go-cmp/cmp"
)

func TestChangeLabels(t *testing.T) {
	tests := []struct {
		name            string
		inputLabels     map[string]string
		labelKeysToSkip []string
		expected        map[string]string
	}{
		{
			name: "with no labels to skip",
			inputLabels: map[string]string{
				"a":           "b",
				"c":           "d",
				"foo.com/bar": "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
			labelKeysToSkip: nil,
			expected: map[string]string{
				fmt.Sprintf("%s%s", common.KeyPrefix, "a"):   "b",
				fmt.Sprintf("%s%s", common.KeyPrefix, "c"):   "d",
				fmt.Sprintf("%s%s", common.KeyPrefix, "bar"): "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
		},
		{
			name: "with no labels to skip not in input",
			inputLabels: map[string]string{
				"a": "b",
				"c": "d",
			},
			labelKeysToSkip: []string{"foo.com/bar"},
			expected: map[string]string{
				fmt.Sprintf("%s%s", common.KeyPrefix, "a"): "b",
				fmt.Sprintf("%s%s", common.KeyPrefix, "c"): "d",
			},
		},
		{
			name: "with labels to skip",
			inputLabels: map[string]string{
				"a":           "b",
				"c":           "d",
				"foo.com/bar": "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
			labelKeysToSkip: []string{"a", "foo.com/bar", common.KeyPrefix},
			expected: map[string]string{
				"a": "b",
				fmt.Sprintf("%s%s", common.KeyPrefix, "c"): "d",
				"foo.com/bar": "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
		},
		{
			name: "with labels to skip",
			inputLabels: map[string]string{
				"a":           "b",
				"c":           "d",
				"foo.com/bar": "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
			labelKeysToSkip: []string{"a", "foo.com/bar", common.KeyPrefix},
			expected: map[string]string{
				"a": "b",
				fmt.Sprintf("%s%s", common.KeyPrefix, "c"): "d",
				"foo.com/bar": "baz",
				fmt.Sprintf("%s%s", common.KeyPrefix, "qux"): "quux",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			labels, _ := ChangeLabels(tc.inputLabels, tc.labelKeysToSkip)
			if diff := cmp.Diff(labels, tc.expected); diff != "" {
				t.Errorf("unexpected plugin default (-want, +got):\n%s", diff)
			}
		})
	}
}
