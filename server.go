package main

import (
	"io"

	"github.com/chrollo-lucifer-12/file-storage/p2p"
)

type FileServerOpts struct {
	StorageRoot       string
	PathTransformFunc PathTransformFunc
	Transport         p2p.Transport
}

type FileServer struct {
	FileServerOpts
	store *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {

	storeOpts := StoreOpts{
		Root:              opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}

	return &FileServer{
		store:          NewStore(storeOpts),
		FileServerOpts: opts,
	}
}

func (s *FileServerOpts) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	return nil
}

func (s *FileServer) Store(key string, r io.Reader) error {
	return s.store.Write(key, r)
}
