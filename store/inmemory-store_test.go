package store

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/alex-bezverkhniy/gonotes/model"
)

const dataFileName = "data.json"

var sampleJSON = []byte(`
{
  "1": {
    "id": 1,
    "title": "simple title",
		"createdAt": "2019-07-12T14:32:21.003914596-05:00",
		"updatedAt": "2019-07-12T14:32:21.003914596-05:00",
    "content": "Simple content"
	}
}`)

type SpyDataLoader struct {
	Calls int
}

func (s *SpyDataLoader) Load() (map[int]model.Note, error) {
	s.Calls++
	return getMockData(), nil
}

func getMockData() map[int]model.Note {
	return map[int]model.Note{
		1: model.Note{},
	}
}

type SpyReader struct {
	r     io.Reader
	Calls int
}

func (s SpyReader) Read(p []byte) (n int, err error) {
	s.Calls++
	return s.r.Read(p)
}

func TestNewFileDataLoader(t *testing.T) {
	fdl := createDataLoader(createSpyReader())

	assertNotNull(t, fdl)
}

func TestLoad(t *testing.T) {
	sr := createSpyReader()
	fdl := createDataLoader(sr)

	got, err := fdl.Load()

	if err != nil {
		t.Errorf("returns an error '%q'", err)
	}

	assertStoreLen(t, got, 1)

}

func TestNewInMemoryNoteStore(t *testing.T) {
	spyDataLoader := &SpyDataLoader{}
	s := NewInMemoryNoteStore(spyDataLoader)
	t.Run("should create new instance", func(t *testing.T) {

		assertNotNull(t, s)
		assertNotNull(t, s.store)
		assertStoreLen(t, s.store, 1)
	})

	t.Run("should call Load method", func(t *testing.T) {
		if spyDataLoader.Calls != 1 {
			t.Errorf("not enough calls to data loader, want 1 got %d", spyDataLoader.Calls)
		}
	})
}

func TestGet(t *testing.T) {
	s := NewInMemoryNoteStore(createDataLoader(createSpyReader()))

	t.Run("returns note by id", func(t *testing.T) {
		want := model.Note{
			ID:        1,
			Title:     "simple title",
			Content:   "Simple content",
			CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
			UpdatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
		}
		got := s.Get(1)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func createSpyReader() *SpyReader {
	return &SpyReader{r: bytes.NewReader(sampleJSON)}
}

func createDataLoader(reader io.Reader) DataLoader {
	return NewFileDataLoader(reader)
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func assertNotNull(t *testing.T, i interface{}) {
	t.Helper()
	if i == nil {
		t.Errorf("instance should not be nil")
	}
}

func assertStoreLen(t *testing.T, m map[int]model.Note, want int) {
	t.Helper()
	if len(m) != want {
		t.Errorf("wrong size of slice, got %d want %d", len(m), want)
	}
}
