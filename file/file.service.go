package file

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repo *repo
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) UploadPost(postId int, fh *multipart.FileHeader) error {
	name := s.genUniqueName(fh.Filename)
	if err := s.upload(name, fh); err != nil {
		return err
	}
	file := PostFile{
		File:   File{Name: name},
		PostId: postId,
	}
	return s.repo.savePost(&file)
}

func (s *Service) genUniqueName(name string) string {
	li := strings.LastIndex(name, ".")
	ext := name[li:]
	return uuid.NewString() + ext
}

func (s *Service) upload(name string, fh *multipart.FileHeader) error {
	dstf, err := os.Create("./assets/" + name)
	if err != nil {
		return err
	}
	defer dstf.Close()

	srcf, err := fh.Open()
	if err != nil {
		return err
	}
	defer srcf.Close()

	_, err = io.Copy(dstf, srcf)
	return err
}

func (s *Service) getById(id int) (*os.File, error) {
	f, err := s.repo.getById(id)
	if err != nil {
		return nil, err
	}
	return os.Open("./assets/" + f.Name)
}
