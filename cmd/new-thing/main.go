package main

import (
	
)

import (
	"bufio"
	"bytes"
	"context"
	"github.com/aaronland/go-things"	
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"	
	wof_reader "github.com/whosonfirst/go-whosonfirst-reader"
	wof_writer "github.com/whosonfirst/go-whosonfirst-writer"	
	"github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/options"
	_ "io/ioutil"
	"flag"
	_ "fmt"
	"log"
)

func main() {

	reader_source := flag.String("reader-source", "", "...")
	writer_source := flag.String("writer-source", "", "...")
	
	depicts_id := flag.Int64("depicts-id", -1, "...")
	text := flag.String("text", "", "...")
	
	flag.Parse()

	ctx := context.Background()
	
	r, err := reader.NewReader(ctx, *reader_source)

	if err != nil {
		log.Fatal(err)
	}

	wr, err := writer.NewWriter(ctx, *writer_source)

	if err != nil {
		log.Fatal(err)
	}
	
	depicts, err := wof_reader.LoadFeatureFromID(ctx, r, *depicts_id)

	if err != nil {
		log.Fatal(err)
	}

	//
	
	thing, err := things.NewThing(ctx, depicts, *text)

	if err != nil {
		log.Fatal(err)
	}

	//
	
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)

	ex_opts, err := options.NewDefaultOptions()

	if err != nil {
		log.Fatal(err)
	}
	
	err = export.Export(thing.Bytes(), ex_opts, bw)

	if err != nil {
		log.Fatal(err)
	}

	bw.Flush()
	body := buf.Bytes()

	//

	err = wof_writer.WriteFeatureBytes(ctx, wr, body)

	if err != nil {
		log.Fatal(err)
	}
}
