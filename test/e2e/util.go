package e2e

import (
	"testing"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func waitForNonEmptyStatus(t *testing.T, err error, name string, kind string, actualStatus interface{}, emptyStatus interface{}) (bool, error) {
	if err != nil {
		if apierrors.IsNotFound(err) {
			t.Logf("Waiting for availability of %s %s\n", name, kind)
			return false, nil
		}
		return false, err
	}

	if (emptyStatus) == actualStatus {
		t.Logf("Waiting for full availability of %s %s\n", name, kind)
		return false, nil
	}
	t.Logf("%s %s available\n", name, kind)
	return true, nil
}
