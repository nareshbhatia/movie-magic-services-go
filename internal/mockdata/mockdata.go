package mockdata

import (
	"embed"
	"io"

	movie "github.com/nareshbhatia/movie-magicservices-go/gen/go/models/movie/v1"
	seed "github.com/nareshbhatia/movie-magicservices-go/gen/go/seed/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

//go:embed */*.json
var seedFS embed.FS

// Loader parses seed data from JSON files and acts as the data store for the application.
//
// The JSON files are parsed into their respective protobuf messages and stored in the wrapper.
type Loader struct {
	wrapper *seed.Wrapper
}

func NewLoader() *Loader {
	return &Loader{
		wrapper: &seed.Wrapper{},
	}
}

func (l *Loader) Load() error {
	if w, err := loadSeedData("seed/movies.json"); err == nil {
		l.wrapper.Movies = w.GetMovies()
	} else {
		return err
	}

	return nil
}

func (l *Loader) GetMovies() []*movie.Movie {
	return l.wrapper.GetMovies()
}

// Close is for demo purposes only and does nothing.
func (l *Loader) Close() error {
	return nil
}

func loadSeedData(path string) (*seed.Wrapper, error) {
	f, err := seedFS.Open(path)
	if err != nil {
		return nil, err
	}

	bts, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var wrapper seed.Wrapper
	if err := protojson.Unmarshal(bts, &wrapper); err != nil {
		return nil, err
	}

	return &wrapper, nil
}
