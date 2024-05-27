package v1

import (
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/utils/pointer"
	"testing"
)

func TestPluginDefaults(t *testing.T) {
	tests := []struct {
		name     string
		config   runtime.Object
		expected runtime.Object
	}{
		{
			name:   "empty ProxyArgs",
			config: &ProxyArgs{},
			expected: &ProxyArgs{
				FilterWaitDurationSeconds: pointer.Int32(30),
				LabelKeysToSkipPrefixing:  nil,
			},
		},
		{
			name: "non-default ProxyArgs",
			config: &ProxyArgs{
				FilterWaitDurationSeconds: pointer.Int32(60),
				LabelKeysToSkipPrefixing:  []string{"foo.com/bar", "foo.com/baz"},
			},
			expected: &ProxyArgs{
				FilterWaitDurationSeconds: pointer.Int32(60),
				LabelKeysToSkipPrefixing:  []string{"foo.com/bar", "foo.com/baz"},
			},
		},
		{
			name:   "empty CandidateArgs",
			config: &CandidateArgs{},
			expected: &CandidateArgs{
				PreBindWaitDurationSeconds: pointer.Int32(30),
			},
		},
		{
			name: "non-default CandidateArgs",
			config: &CandidateArgs{
				PreBindWaitDurationSeconds: pointer.Int32(120),
			},
			expected: &CandidateArgs{
				PreBindWaitDurationSeconds: pointer.Int32(120),
			},
		},
	}

	for _, tc := range tests {
		scheme := runtime.NewScheme()
		utilruntime.Must(AddToScheme(scheme))
		t.Run(tc.name, func(t *testing.T) {
			scheme.Default(tc.config)
			if diff := cmp.Diff(tc.config, tc.expected); diff != "" {
				t.Errorf("unexpected plugin default (-want, +got):\n%s", diff)
			}
		})
	}
}
