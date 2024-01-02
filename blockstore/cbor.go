package blockstore

import (
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"
)

type CborStore struct {
	cst *cbor.BasicIpldStore
}

func NewCborStore(bs *Store) *CborStore {
	cst := cbor.NewCborStore(bs.bs)
	cst.DefaultMultihash = mh.SHA2_256

	return &CborStore{cst: cst}
}

func (s *CborStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return s.cst.Put(ctx, v)
}

func (s *CborStore) Get(ctx context.Context, id cid.Cid, out interface{}) error {
	return s.cst.Get(ctx, id, out)
}
