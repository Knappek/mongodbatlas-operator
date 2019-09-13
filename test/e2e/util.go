package e2e

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
)

func isInDesiredState(t *testing.T, err error, name string, kind string, actualState interface{}, desiredState interface{}) (bool, error) {
	if err != nil {
		if errors.IsNotFound(err) {
			t.Logf("Waiting for availability of %s %s\n", name, kind)
			return false, nil
		}
		return false, err
	}

	if desiredState == actualState {
		t.Logf("%s %s available\n", name, kind)
		return true, nil
	}
	t.Logf("Waiting for full availability of %s %s\n", name, kind)
	return false, nil
}
