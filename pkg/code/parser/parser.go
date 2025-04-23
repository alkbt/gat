package parser

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
)

func getModulePath(reader io.Reader) (string, error) {
	const op = "get module path"
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("%s: read module file: %w", op, err)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", fmt.Errorf("%s: parse module file: %w", op, err)
	}

	return f.Module.Mod.Path, nil
}
