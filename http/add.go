package http

import (
	"bufio"
	"bytes"	
	_ "errors"
	"log"
	gohttp "net/http"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"
	wof_reader "github.com/whosonfirst/go-whosonfirst-reader"
	wof_writer "github.com/whosonfirst/go-whosonfirst-writer"	
	"github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/options"
	"github.com/aaronland/go-things"		
)

type AddThingHandlerOptions struct {
	Reader reader.Reader
	Writer writer.Writer
}

func AddThingHandler(opts AddThingHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		ctx := req.Context()

		log.Println("HELLO")
		
		if req.Method != "POST" {
			gohttp.Error(rsp, "Method not allowed.", gohttp.StatusMethodNotAllowed)
			return
		}
		
		depicts_id := int64(-1)
		text := ""
		
		depicts, err := wof_reader.LoadFeatureFromID(ctx, opts.Reader, depicts_id)

		if err != nil {
 			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		thing, err := things.NewThing(ctx, depicts, text)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		bw := bufio.NewWriter(&buf)
		
		ex_opts, err := options.NewDefaultOptions()

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}
	
		err = export.Export(thing.Bytes(), ex_opts, bw)
		
		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		bw.Flush()
		body := buf.Bytes()
		
		err = wof_writer.WriteFeatureBytes(ctx, opts.Writer, body)
		
		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)			
			return
		}

		// headers
		
		rsp.Write(body)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
