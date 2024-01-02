package main

import (
	"context"
	"fmt"
	"os"

	"github.com/olehmushka/ipfs-storing-example/blockstore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bs := blockstore.New()
	// Put value to blockstore
	cid, err := bs.Put(ctx, []byte("hello world 2"))
	if err != nil {
		panic(err)
	}
	b, err := bs.Get(ctx, cid)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cid.String()))
	fmt.Println(string(b))

	// Put value to cbor store
	cst := blockstore.NewCborStore(bs)
	cid2, err := cst.Put(ctx, "hello world 3")
	if err != nil {
		panic(err)
	}
	var v string
	if err := cst.Get(ctx, cid2, &v); err != nil {
		panic(err)
	}
	fmt.Println(string(cid2.String()))
	fmt.Println(string(v))

	// Put image to blockstore
	f, err := os.Open("public/pexels-christian-heitz-842711.jpg")
	if err != nil {
		panic(err)
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	defer f.Close()
	body, err := os.ReadFile("public/pexels-christian-heitz-842711.jpg")
	if err != nil {
		panic(err)
	}
	cid3, err := bs.Put(ctx, body)
	if err != nil {
		panic(err)
	}
	b3, err := bs.Get(ctx, cid3)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cid3.String()))
	fmt.Println(string(b3) == string(body))
	err = os.WriteFile("public/out.jpg", b3, info.Mode())
	if err != nil {
		panic(err)
	}
}
