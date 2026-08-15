package infra

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func EnsureCleanOutputDir(dir string, force bool) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // fresh directory, nothing to do
		}
		return err
	}
	if len(entries) == 0 {
		return nil
	}
	if !force {
		return fmt.Errorf("%s already exists and is not empty — use --force to overwrite, or pick a different name", dir)
	}
	return os.RemoveAll(dir)
}

func CopyTemplateFS(fsys embed.FS, root, outDir string, data any) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		content, err := fsys.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		dest := filepath.Join(outDir, strings.TrimSuffix(rel, ".tmpl"))
		if err := EnsureDir(filepath.Dir(dest)); err != nil {
			return err
		}
		f, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		tmpl := template.Must(template.New(d.Name()).Parse(string(content)))
		return tmpl.Execute(f, data)
	})
}
