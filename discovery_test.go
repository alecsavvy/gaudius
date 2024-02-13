package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscoverySelection(t *testing.T) {
	nodes, err := SelectHealthyDiscoveryNode([]string{})
	require.NotNil(t, err)
	require.Empty(t, nodes)

}
