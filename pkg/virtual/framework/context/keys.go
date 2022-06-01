/*
Copyright 2021 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"context"
	"errors"
)

type virtualWorkspaceNameKeyType string

// virtualWorkspaceNameKey is a context key that contains the name of the
// virtual workspace that should serve a given request according to its URL path.
const virtualWorkspaceNameKey virtualWorkspaceNameKeyType = "VirtualWorkspaceName"

// WithVirtualWorkspaceName adds the VirtualWorkspace name to the context.
func WithVirtualWorkspaceName(ctx context.Context, virtualWorkspaceName string) context.Context {
	return context.WithValue(ctx, virtualWorkspaceNameKey, virtualWorkspaceName)
}

// VirtualWorkspaceNameFrom retrieves the VirtualWorkspace name from the context, if any.
func VirtualWorkspaceNameFrom(ctx context.Context) (string, error) {
	wcn, hasVirtualWorkspaceName := ctx.Value(virtualWorkspaceNameKey).(string)
	if !hasVirtualWorkspaceName {
		return "", errors.New("context must contain a valid non-empty virtual workspace name")
	}
	return wcn, nil
}
