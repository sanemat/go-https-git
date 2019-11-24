package httpsgit

import (
	"context"
	"flag"
	"io"
)

type cmdPush struct{}

func (cp *cmdPush) name() string {
	return "push"
}

func (cp *cmdPush) description() string {
	return "https push"
}

func (cp *cmdPush) run(ctx context.Context, argv []string, outStream io.Writer, errStream io.Writer) error {
	fs := flag.NewFlagSet("httpsgit push", flag.ContinueOnError)
	fs.SetOutput(errStream)

	if err := fs.Parse(argv); err != nil {
		return err
	}

	return nil
}
