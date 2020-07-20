package core

import (
	"github.com/jexia/semaphore/pkg/codec"
	"github.com/jexia/semaphore/pkg/core/instance"
	"github.com/jexia/semaphore/pkg/flow"
	"github.com/jexia/semaphore/pkg/functions"
	"github.com/jexia/semaphore/pkg/metadata"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/labels"
	"github.com/jexia/semaphore/pkg/specs/template"
	"github.com/jexia/semaphore/pkg/specs/types"
)

// DefaultOnError sets the default values for not defined properties
func DefaultOnError(err *specs.OnError) {
	if err == nil {
		err = &specs.OnError{}
	}

	if err.Status == nil {
		err.Status = &specs.Property{
			Type:  types.Int64,
			Label: labels.Optional,
			Reference: &specs.PropertyReference{
				Resource: "error",
				Path:     "status",
			},
		}
	}

	if err.Message == nil {
		err.Message = &specs.Property{
			Type:  types.String,
			Label: labels.Optional,
			Reference: &specs.PropertyReference{
				Resource: "error",
				Path:     "message",
			},
		}
	}
}

// MergeOnError merges the right on error specs into the left on error
func MergeOnError(left *specs.OnError, right *specs.OnError) {
	if left == nil || right == nil {
		return
	}

	if left.Message == nil {
		left.Message = right.Message.Clone()
	}

	if left.Status == nil {
		left.Status = right.Status.Clone()
	}

	if len(left.Params) == 0 {
		left.Params = make(map[string]*specs.Property, len(right.Params))

		for key, param := range right.Params {
			left.Params[key] = param.Clone()
		}
	}

	if left.Response == nil {
		left.Response = right.Response.Clone()
	}
}

// ConstructErrorHandle clones any previously defined error objects or error handles
func ConstructErrorHandle(manifest *specs.FlowsManifest) {
	for _, flow := range manifest.Flows {
		DefaultOnError(flow.OnError)

		if flow.OnError.Response == nil {
			flow.OnError.Response = manifest.Error.Clone()
		}

		for _, node := range flow.Nodes {
			if node.OnError == nil {
				node.OnError = flow.OnError.Clone()
				continue
			}

			MergeOnError(node.OnError, flow.OnError)
		}
	}

	for _, proxy := range manifest.Proxy {
		DefaultOnError(proxy.OnError)

		if proxy.OnError.Response == nil {
			proxy.OnError.Response = manifest.Error.Clone()
		}

		for _, node := range proxy.Nodes {
			if node.OnError == nil {
				node.OnError = proxy.OnError.Clone()
				continue
			}

			MergeOnError(node.OnError, proxy.OnError)
		}
	}
}

// NewError constructs a new error object from the given parameter map and codec
func NewError(ctx instance.Context, node *specs.Node, mem functions.Collection, constructor codec.Constructor, err *specs.OnError) (*flow.OnError, error) {
	if err == nil {
		return nil, nil
	}

	var codec codec.Manager
	var meta *metadata.Manager
	var stack functions.Stack

	if err.Response != nil && constructor != nil {
		params := err.Response

		// TODO: check if I would like props to be defined like this
		manager, err := constructor.New(template.JoinPath(node.Name, template.ErrorResource), params)
		if err != nil {
			return nil, err
		}

		codec = manager
		stack = mem[params]
		meta = metadata.NewManager(ctx, node.Name, params.Header)
	}

	return flow.NewOnError(stack, codec, meta, err.Status, err.Message), nil
}
