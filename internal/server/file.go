package server

import (
	"github.com/disism/saikan/ent"
	"strconv"
)

type File struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
	Name string `json:"name"`
	Size string `json:"size"`
}

func FMTFile(f *ent.Files) *File {
	r := &File{
		ID:   strconv.FormatUint(f.ID, 10),
		Hash: f.Hash,
		Name: f.Name,
		Size: strconv.FormatUint(f.Size, 10),
	}
	return r
}
