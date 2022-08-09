package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func cmdRoot(w io.Writer) *cobra.Command {
	r := &cobra.Command{
		Use:   "enc",
		Short: "Enc is PGP for humans",
		Long: `
			A user-friendly CLI tool to work with PGP keys:
			create, add, list, encrypt, decrypt, sign, verify signatures.
		`,
	}
	r.AddCommand(cmdVersion(w))
	return r
}

func Main(w io.Writer) error {
	return cmdRoot(w).Execute()
}

func Entrypoint() {
	err := Main(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}