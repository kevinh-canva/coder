package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coder/coder/cli/clitest"
	"github.com/coder/coder/coderd/coderdtest"
	"github.com/coder/coder/pty/ptytest"
)

func TestUserList(t *testing.T) {
	t.Parallel()
	client := coderdtest.New(t, nil)
	coderdtest.CreateFirstUser(t, client)
	cmd, root := clitest.New(t, "users", "list")
	clitest.SetupConfig(t, client, root)
	pty := ptytest.New(t)
	cmd.SetIn(pty.Input())
	cmd.SetOut(pty.Output())
	errC := make(chan error)
	go func() {
		errC <- cmd.Execute()
	}()
	require.NoError(t, <-errC)
	pty.ExpectMatch("coder.com")
}
