package mockdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLoader(t *testing.T) {
	t.Parallel()

	t.Run("Load", func(t *testing.T) {
		t.Parallel()

		loader := NewLoader()
		require.NoError(t, loader.Load())

		assert.NotEmpty(t, loader.GetMovies())
	})
}
