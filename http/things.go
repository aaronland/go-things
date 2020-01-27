package http

import (
	"errors"
	"html/template"
	_ "log"
	gohttp "net/http"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"
	
)

type ThingsHandlerOptions struct {
	Templates *template.Template
	// Endpoints *Endpoints
	Reader reader.Reader
	Writer writer.Writer
}

type ThingsVars struct {

}

func ThingsHandler(opts ThingsHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("things")

	if t == nil {
		return nil, errors.New("Missing index template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := ThingsVars{

		}

		RenderTemplate(rsp, t, vars)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
