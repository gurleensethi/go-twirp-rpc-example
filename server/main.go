package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gurleensethi/go-twirp-rpc-example/rpc/notes"
	"github.com/twitchtv/twirp"
)

type notesService struct {
	Notes     []notes.Note
	CurrentId int32
}

func (s *notesService) GetAllNotes(ctx context.Context, params *notes.GetAllNotesParams) (*notes.GetAllNotesResult, error) {
	allNotes := make([]*notes.Note, 0)

	for _, note := range s.Notes {
		n := note
		allNotes = append(allNotes, &n)
	}

	return &notes.GetAllNotesResult{
		Notes: allNotes,
	}, nil
}

func (s *notesService) CreateNote(ctx context.Context, params *notes.CreateNoteParams) (*notes.Note, error) {
	if len(params.Text) < 4 {
		return nil, twirp.InvalidArgument.Error("Text should be min 4 characters.")
	}

	note := notes.Note{
		Id:        s.CurrentId,
		Text:      params.Text,
		CreatedAt: time.Now().UnixMilli(),
	}

	s.Notes = append(s.Notes, note)

	s.CurrentId++

	return &note, nil
}

func main() {
	notesServer := notes.NewNotesServiceServer(&notesService{})

	mux := http.NewServeMux()
	mux.Handle(notesServer.PathPrefix(), notesServer)

	err := http.ListenAndServe(":8000", notesServer)
	if err != nil {
		panic(err)
	}
}
