package repositories

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/alex-bezverkhniy/gonotes/model"
)

func TestNoteRepository_Create(t *testing.T) {
	type fields struct {
		DataFileName string
		Notes        map[string]model.Note
	}
	type args struct {
		note model.Note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Note
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{}},
			args:    args{note: model.Note{ID: "1", Title: "Sample note"}},
			want:    model.Note{ID: "1", Title: "Sample note"},
			wantErr: false,
		},
		{
			name:    "error - ID should not be empty",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{Title: "Sample note"}}},
			args:    args{note: model.Note{Title: "Sample note"}},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - note should not be empty",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{Title: "Sample note"}}},
			args:    args{},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - Duplicate record",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{note: model.Note{ID: "1", Title: "Sample note"}},
			want:    model.Note{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nr := &NoteRepository{
				DataFileName: tt.fields.DataFileName,
				Notes:        tt.fields.Notes,
			}
			got, err := nr.Create(tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_Update(t *testing.T) {
	type fields struct {
		DataFileName string
		Notes        map[string]model.Note
	}
	type args struct {
		ID   string
		note model.Note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Note
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "1", note: model.Note{ID: "1", Title: "Updated Sample note"}},
			want:    model.Note{ID: "1", Title: "Updated Sample note"},
			wantErr: false,
		},
		{
			name:    "error - not found",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "10", note: model.Note{ID: "10", Title: "Updated Sample note"}},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - bad request - ID is empty",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "", note: model.Note{ID: "10", Title: "Updated Sample note"}},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - bad request - note is empty",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "1"},
			want:    model.Note{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nr := &NoteRepository{
				DataFileName: tt.fields.DataFileName,
				Notes:        tt.fields.Notes,
			}
			got, err := nr.Update(tt.args.ID, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteRepository(t *testing.T) {
	type args struct {
		dataFileName string
	}
	tests := []struct {
		name string
		args args
		want *NoteRepository
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{dataFileName: "../data.json"},
			want: &NoteRepository{DataFileName: "../data.json", Notes: map[string]model.Note{
				"1": model.Note{
					ID:        "1",
					Title:     "Note 1",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
					Desc:      "Test note #1",
					Content:   "Sample content of note #1",
				},
				"2": model.Note{
					ID:        "2",
					Title:     "Note 2",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914791-05:00"),
					Desc:      "Test note #2",
					Content:   "Sample content of note #2",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteRepository(tt.args.dataFileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_FindByID(t *testing.T) {
	type fields struct {
		DataFileName string
		Notes        map[string]model.Note
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Note
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "1"},
			want:    model.Note{ID: "1", Title: "Sample note"},
			wantErr: false,
		},
		{
			name:    "error - not found",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "11"},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - not found in the store",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"10": model.Note{ID: "10", Title: "Sample note"}}},
			args:    args{ID: "1"},
			want:    model.Note{},
			wantErr: true,
		},
		{
			name:    "error - ID should not by empty found",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: ""},
			want:    model.Note{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nr := &NoteRepository{
				DataFileName: tt.fields.DataFileName,
				Notes:        tt.fields.Notes,
			}
			got, err := nr.FindByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_FindAll(t *testing.T) {
	type fields struct {
		DataFileName string
		Notes        map[string]model.Note
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]model.Note
	}{
		// TODO: Add test cases.
		{
			name:   "success",
			fields: fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			want:   map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}},
		},
		{
			name:   "success - should load initial data",
			fields: fields{DataFileName: "../data.json"},
			want: map[string]model.Note{
				"1": model.Note{
					ID:        "1",
					Title:     "Note 1",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
					Desc:      "Test note #1",
					Content:   "Sample content of note #1",
				},
				"2": model.Note{
					ID:        "2",
					Title:     "Note 2",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914791-05:00"),
					Desc:      "Test note #2",
					Content:   "Sample content of note #2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nr := &NoteRepository{
				DataFileName: tt.fields.DataFileName,
				Notes:        tt.fields.Notes,
			}
			if got := nr.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_Delete(t *testing.T) {
	type fields struct {
		DataFileName string
		Notes        map[string]model.Note
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "1"},
			wantErr: false,
		},
		{
			name:    "error - not found",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: "11"},
			wantErr: true,
		},
		{
			name:    "error - not found in the store",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"10": model.Note{ID: "10", Title: "Sample note"}}},
			args:    args{ID: "1"},
			wantErr: true,
		},
		{
			name:    "error - ID should not by empty found",
			fields:  fields{DataFileName: "test-data.json", Notes: map[string]model.Note{"1": model.Note{ID: "1", Title: "Sample note"}}},
			args:    args{ID: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nr := &NoteRepository{
				DataFileName: tt.fields.DataFileName,
				Notes:        tt.fields.Notes,
			}
			if err := nr.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func Test_loadNotes(t *testing.T) {
	type args struct {
		dataFileName string
	}
	tests := []struct {
		name string
		args args
		want map[string]model.Note
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{dataFileName: "../data.json"},
			want: map[string]model.Note{
				"1": model.Note{
					ID:        "1",
					Title:     "Note 1",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
					Desc:      "Test note #1",
					Content:   "Sample content of note #1",
				},
				"2": model.Note{
					ID:        "2",
					Title:     "Note 2",
					CreatedAt: parseTime("2019-07-12T14:32:21.003914791-05:00"),
					Desc:      "Test note #2",
					Content:   "Sample content of note #2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadNotes(tt.args.dataFileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeNotes(t *testing.T) {
	type args struct {
		dataFileName string
		data         map[string]model.Note
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				dataFileName: "tmp.json",
				data: map[string]model.Note{
					"1": model.Note{
						ID:        "1",
						Title:     "Note 1",
						CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
						Desc:      "Test note #1",
						Content:   "Sample content of note #1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error - no data to save",
			args: args{
				dataFileName: "tmp.json",
				data:         map[string]model.Note{},
			},
			wantErr: true,
		},
		{
			name: "error - nil data to save",
			args: args{
				dataFileName: "tmp.json",
				data:         nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := storeNotes(tt.args.dataFileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("storeNotes() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(tt.args.dataFileName); os.IsNotExist(err) {
				t.Errorf("storeNotes() data file does not exists: %v", err)
			}
		})
	}
}

func Test_storeNotesCheckFileWasUpdated(t *testing.T) {
	want := map[string]model.Note{
		"1": model.Note{
			ID:        "1",
			Title:     "Note 1",
			CreatedAt: parseTime("2019-07-12T14:32:21.003914596-05:00"),
			Desc:      "Test note #1",
			Content:   "Sample content of note #1",
		},
	}

	err := storeNotes("tmp.json", want)

	if err != nil {
		t.Errorf("storeNotes() did not save data: %v", err)
	}

	want["2"] = model.Note{
		ID:        "2",
		Title:     "Note 2",
		CreatedAt: parseTime("2019-07-12T14:32:21.003914598-05:00"),
		Desc:      "Test note #2",
		Content:   "Sample content of note #2",
	}

	err = storeNotes("tmp.json", want)

	if err != nil {
		t.Errorf("storeNotes() did not save data: %v", err)
	}

	got := loadNotes("tmp.json")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NoteRepository.FindByID() = %v, want %v", got, want)
	}

}

func TestMain(m *testing.M) {
	// Before

	ret := m.Run()

	// After
	log.Println("Cleanup!")
	err := os.Remove("tmp.json")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(ret)
}
