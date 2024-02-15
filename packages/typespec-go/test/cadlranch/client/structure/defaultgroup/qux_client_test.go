//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultgroup_test

import (
	"context"
	"defaultgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuxClient_Eight(t *testing.T) {
	client, err := defaultgroup.NewServiceClient(defaultgroup.ClientTypeDefault, nil)
	require.NoError(t, err)
	resp, err := client.NewQuxClient().Eight(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
