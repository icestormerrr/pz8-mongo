package notes_test

import (
	"context"
	"testing"

	"github.com/icestormerrr/pz8-mongo/internal/db"
	"github.com/icestormerrr/pz8-mongo/internal/notes"
)

func TestCreateAndGet(t *testing.T) {
	ctx := context.Background()
	deps, err := db.ConnectMongo(ctx, "mongodb://root:secret@localhost:27017/?authSource=admin", "pz8_test")

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		deps.Client.Disconnect(ctx)
		deps.Client.Database("pz8_test").Drop(ctx)
	})
	r, err := notes.NewRepo(deps.Database)
	if err != nil {
		t.Fatal(err)
	}

	created, err := r.Create(ctx, "T1", "C1")
	if err != nil {
		t.Fatal(err)
	}

	got, err := r.ByID(ctx, created.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "T1" {
		t.Fatalf("want T1 got %s", got.Title)
	}
}
