package blockstore

import (
	"context"

	"github.com/ipfs/boxo/blockstore"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
)

type Store struct {
	bs blockstore.Blockstore
}

func New() *Store {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds)

	return &Store{bs: bs}
}

func (s *Store) Put(ctx context.Context, data []byte) (cid.Cid, error) {
	b := blocks.NewBlock(data)
	return b.Cid(), s.bs.Put(ctx, b)
}

func (s *Store) Get(ctx context.Context, id cid.Cid) ([]byte, error) {
	b, err := s.bs.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return b.RawData(), nil
}
