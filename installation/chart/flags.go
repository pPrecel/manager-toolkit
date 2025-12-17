package chart

import (
	"fmt"

	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/strvals"
)

type FlagsBuilder interface {
	Build() (map[string]interface{}, error)
	With(string, interface{}) *flagsBuilder
}

// flagsBuilder is used to build Helm chart flags in a structured way.
type flagsBuilder struct {
	flags map[string]interface{}
}

func NewFlagsBuilder() FlagsBuilder {
	return &flagsBuilder{
		flags: map[string]interface{}{},
	}
}

// Build constructs the final map of Helm chart flags.
func (fb *flagsBuilder) Build() (map[string]interface{}, error) {
	flags := map[string]interface{}{}
	for key, value := range fb.flags {
		flag := fmt.Sprintf("%s=%v", key, value)
		err := strvals.ParseInto(flag, flags)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %s flag", flag)
		}
	}
	return flags, nil
}

// With adds a new flag to the builder in form of key-value pair.
// example: With("global.commonLabels.managedBy", "my-manager")
func (fb *flagsBuilder) With(key string, value interface{}) *flagsBuilder {
	fb.flags[key] = value
	return fb
}
