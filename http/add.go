package http

import (
	"bufio"
	"bytes"	
	_ "errors"
	"encoding/json"
	"log"
	gohttp "net/http"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"
	wof_reader "github.com/whosonfirst/go-whosonfirst-reader"
	wof_writer "github.com/whosonfirst/go-whosonfirst-writer"	
	"github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/options"
	"github.com/aaronland/go-things"
	"github.com/whosonfirst/go-sanitize"			
)

type AddThingHandlerOptions struct {
	Reader reader.Reader
	Writer writer.Writer
}

func AddThingHandler(opts AddThingHandlerOptions) (gohttp.Handler, error) {

	s_opts := sanitize.DefaultOptions()

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		switch req.Method {
		case "PUT":
			// pass
		default:
			gohttp.Error(rsp, "Method not allowed.", gohttp.StatusMethodNotAllowed)
			return
		}

		var thing *things.Thing

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&thing)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		depicts_id := thing.DepictsId
		note := thing.Note

		log.Println("THING", depicts_id, note)
		
		text, err := sanitize.SanitizeString(note, s_opts)

		if err != nil {
 			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}
		
		ctx := req.Context()
		
		depicts, err := wof_reader.LoadFeatureFromID(ctx, opts.Reader, depicts_id)

		if err != nil {
			log.Println("SAD 2", err)			
 			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		thing_f, err := things.NewThing(ctx, depicts, text)

		if err != nil {
			log.Println("SAD 3", err)			
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
	
		err = export.Export(thing_f.Bytes(), ex_opts, bw)
		
		if err != nil {
			log.Println("SAD 4", err)
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		bw.Flush()
		body := buf.Bytes()
		
		err = wof_writer.WriteFeatureBytes(ctx, opts.Writer, body)
		
		if err != nil {
			log.Println("SAD 5", err)			
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
