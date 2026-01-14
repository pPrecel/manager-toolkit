package annotation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	annotationFormat = "%s.kyma-project.io/managed-by-%s-disclaimer"
	messageFormat    = "DO NOT EDIT - This resource is managed by %s.\nAny modifications are discarded and the resource is reverted to the original state."
)

// AddDoNotEditDisclaimer adds a "do not edit" disclaimer annotation to the given unstructured Kubernetes object
func AddDoNotEditDisclaimer(managerName string, obj unstructured.Unstructured) unstructured.Unstructured {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	annotations[fmt.Sprintf(annotationFormat, managerName, managerName)] = fmt.Sprintf(messageFormat, managerName)
	obj.SetAnnotations(annotations)

	return obj
}
