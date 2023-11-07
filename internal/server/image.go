package server

import (
	"fmt"
	"github.com/disism/saikan/ent"
	"github.com/disism/saikan/ent/files"
	"github.com/disism/saikan/ent/images"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Images interface {
	AddImage() error
}

const (
	ImageNotFound = "image not found"
)

type Image struct {
	ID     string `json:"id"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	File   *File  `json:"file"`
}

func FMTImage(i *ent.Images) *Image {
	r := &Image{
		ID:     strconv.FormatUint(i.ID, 10),
		Width:  i.Width,
		Height: i.Height,
		File:   FMTFile(i.Edges.File),
	}
	return r
}

type AddImageForm struct {
	Hash   string `form:"hash" json:"hash" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	Size   string `form:"size" json:"size" binding:"required"`
	Height int32  `form:"height" json:"height" binding:"required"`
	Width  int32  `form:"width" json:"width" binding:"required"`
}

func (s *Server) AddImage() error {

	var form AddImageForm
	if err := s.ctx.ShouldBind(&form); err != nil {
		ErrorBadRequest(s.ctx, err.Error())
		return nil
	}

	tx, err := s.client.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("add image tx error: %w", err)
	}
	defer tx.Rollback()

	f, err := func() (*ent.Files, error) {
		query, err := tx.Files.Query().
			Where(
				files.HashEQ(form.Hash),
			).
			Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				size, err := strconv.ParseUint(form.Size, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("parse size error: %w", err)
				}

				create, err := tx.Files.
					Create().
					SetName(form.Name).
					SetHash(form.Hash).
					SetSize(size).
					Save(s.ctx)
				if err != nil {
					return nil, fmt.Errorf("add image create file error: %w", err)
				}
				return create, nil
			}
			return nil, fmt.Errorf("add image query file error: %w", err)
		}
		return query, nil
	}()
	if err != nil {
		return err
	}

	image, err := func() (*ent.Images, error) {
		query, err := tx.Images.
			Query().
			Where(
				images.HasFileWith(
					files.IDEQ(f.ID),
				),
			).Only(s.ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				create, err := tx.Images.Create().
					SetFile(f).
					SetHeight(form.Height).
					SetWidth(form.Width).
					Save(s.ctx)
				if err != nil {
					return nil, fmt.Errorf("add image create image error: %w", err)
				}
				return create, nil
			}
			return nil, fmt.Errorf("add image query image error: %w", err)
		}
		return query, nil
	}()

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("add image commit error: %w", err)
	}
	Success(s.ctx, gin.H{
		"code": http.StatusOK,
		"id":   strconv.FormatUint(image.ID, 10),
		"hash": form.Hash,
	})

	return nil
}
