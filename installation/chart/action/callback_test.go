package action

import (
	"errors"
	"testing"

	"github.com/kyma-project/manager-toolkit/installation/base/resource"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestPreApplyWithPredicate(t *testing.T) {
	t.Run("update only matching resources", func(t *testing.T) {
		u := &unstructured.Unstructured{}
		u.SetKind("TestKind")
		err := PreApplyWithPredicate(
			func(u *unstructured.Unstructured) error {
				u.SetNamespace("test-ns")
				return nil
			},
			resource.HasKind("TestKind"),
		)(u)
		require.NoError(t, err)
		require.Equal(t, "test-ns", u.GetNamespace())
	})
	t.Run("handle error", func(t *testing.T) {
		u := &unstructured.Unstructured{}
		u.SetKind("TestKind")
		testErr := errors.New("test error")
		err := PreApplyWithPredicate(
			func(u *unstructured.Unstructured) error {
				return testErr
			},
			resource.HasKind("TestKind"),
		)(u)
		require.ErrorIs(t, err, testErr)
	})
	t.Run("skip non-matching resources", func(t *testing.T) {
		err := PreApplyWithPredicate(
			func(u *unstructured.Unstructured) error {
				// should not be called
				t.Fail()
				return nil
			},
			resource.HasKind("TestKind"),
		)(&unstructured.Unstructured{})
		require.NoError(t, err)
	})
}

func TestPostUninstallWithPredicate(t *testing.T) {
	t.Run("run only for matching resources", func(t *testing.T) {
		u := unstructured.Unstructured{}
		u.SetKind("TestKind")
		done, err := PostUninstallWithPredicate(
			func(u unstructured.Unstructured) (bool, error) {
				return true, nil
			},
			resource.HasKind("TestKind"),
		)(u)
		require.NoError(t, err)
		require.True(t, done)
	})
	t.Run("handle error from post action", func(t *testing.T) {
		u := unstructured.Unstructured{}
		u.SetKind("TestKind")
		testErr := errors.New("test error")
		done, err := PostUninstallWithPredicate(
			func(u unstructured.Unstructured) (bool, error) {
				return false, testErr
			},
			resource.HasKind("TestKind"),
		)(u)
		require.ErrorIs(t, err, testErr)
		require.False(t, done)
	})
	t.Run("skip non-matching resources", func(t *testing.T) {
		u := unstructured.Unstructured{}
		u.SetKind("OtherKind")
		done, err := PostUninstallWithPredicate(
			func(u unstructured.Unstructured) (bool, error) {
				// should not be called
				t.Fail()
				return false, nil
			},
			resource.HasKind("TestKind"),
		)(u)
		require.NoError(t, err)
		require.True(t, done)
	})
}
