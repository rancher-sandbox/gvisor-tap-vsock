package virtualnetwork

import (
	"testing"

	"github.com/containers/gvisor-tap-vsock/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

func TestCreatePortsForwarderDisablesSourceIPPropagation(t *testing.T) {
	configuration := &types.Configuration{DisableSourceIPPropagation: true}
	fw, err := createPortsForwarder(configuration, stack.New(stack.Options{}))
	require.NoError(t, err)
	assert.True(t, fw.DisableSourceIPPropagation)
}
