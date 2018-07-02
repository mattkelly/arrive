package v1alpha1

import (
	"fmt"
)

func (o *Operand) Resolve() (interface{}, error) {
	// TODO handle o == nil case?
	if o.Value != nil {
		return *o.Value, nil
	}

	if o.ValueFrom != nil {
		return nil, fmt.Errorf("ValueFrom not yet supported")
	}

	return nil, fmt.Errorf("must specify either value or valueFrom")
}
