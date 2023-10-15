package filesystem

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/evolidev/config"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var tmp embed.FS

func TestEmbedFS(t *testing.T) {
	t.Run("has dir should return true if directory exists", func(t *testing.T) {
		embedFs := NewFS(tmp)

		assert.True(t, embedFs.HasDir("resources/views"))
	})

	t.Run("has dir should return false if dir does not exists", func(t *testing.T) {
		embedFs := NewFS(tmp)

		assert.False(t, embedFs.HasDir("not_exists"))
	})

	t.Run("sub should return sub tree", func(t *testing.T) {
		embedFs := NewFS(tmp)
		sub, _ := fs.Sub(embedFs.FS(), "resources/views")
		f, _ := sub.Open("templates/layout.html")
		i, _ := f.Stat()

		assert.Equal(t, "layout.html", i.Name())
	})
}

func TestLocalFS(t *testing.T) {
	config.SetDirectory("./")
	cnf := config.NewConfig("storage")
	abs, _ := filepath.Abs("./")
	cnf.Set("local.base_path", abs)

	localStore := NewFS(os.DirFS(abs))

	assert.True(t, localStore.HasDir("resources/views"))
}
