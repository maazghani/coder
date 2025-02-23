package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coder/coder/v2/cli/clibase"
	"github.com/coder/coder/v2/cli/clitest"
	"github.com/coder/coder/v2/cli/config"
	"github.com/coder/coder/v2/enterprise/cli"
)

func newCLI(t *testing.T, args ...string) (*clibase.Invocation, config.Root) {
	var root cli.RootCmd
	cmd, err := root.Command(root.EnterpriseSubcommands())
	require.NoError(t, err)
	return clitest.NewWithCommand(t, cmd, args...)
}

func TestEnterpriseHandlersOK(t *testing.T) {
	t.Parallel()

	var root cli.RootCmd
	cmd, err := root.Command(root.EnterpriseSubcommands())
	require.NoError(t, err)

	clitest.HandlersOK(t, cmd)
}
