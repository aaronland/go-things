package main

import (
	
)

import (
	"context"
	_ "github.com/aaronland/go-things"
	"github.com/aaronland/go-things/http"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-crumb"	
	"github.com/aaronland/go-things/assets/templates"			
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"
	gohttp "net/http"
	"html/template"
	"flag"
	"fmt"
	"log"
)

func main() {

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")
	
	reader_source := flag.String("reader-source", "", "...")
	writer_source := flag.String("writer-source", "", "...")

	crumb_dsn := flag.String("crumb-dsn", "secret=P3sbIgCiRHTKt8RqaudRfU8MdrkEnh7L ttl=300 separator=: extra=x key=/add/", "...")
	
	path_templates := flag.String("templates", "", "An optional string for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")
	
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

	cr_config, err := crumb.NewCrumbConfigFromDSN(*crumb_dsn)

	if err != nil {
		log.Fatal(err)
	}
	
	// start of sudo put me in a package

	t := template.New("things").Funcs(template.FuncMap{
		"Add": func(i int, offset int) int {
			return i + offset
		},
	})

	if *path_templates != "" {

		t, err = t.ParseGlob(*path_templates)

		if err != nil {
			log.Fatal(err)
		}

	} else {

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				log.Fatal(err)				
			}

			t, err = t.Parse(string(body))

			if err != nil {
				log.Fatal(err)				
			}
		}
	}

	// end of sudo put me in a package

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	
	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	//

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}
	
	err = http.AppendStaticAssetHandlers(mux)
	
	if err != nil {
		log.Fatal(err)
	}

	//
	
	things_opts := http.ThingsHandlerOptions{
		Templates: t,
		Reader: r,
		Writer: wr,
	}

	things_handler, err := http.ThingsHandler(things_opts)

	if err != nil {
		log.Fatal(err)
	}

	things_handler = bootstrap.AppendResourcesHandler(things_handler, bootstrap_opts)
	things_handler = crumb.EnsureCrumbHandler(cr_config, things_handler)
	
	mux.Handle("/", things_handler)
	
	add_opts := http.AddThingHandlerOptions{
		Reader: r,
		Writer: wr,
	}

	add_handler, err := http.AddThingHandler(add_opts)

	if err != nil {
		log.Fatal(err)
	}

	add_handler = crumb.EnsureCrumbHandler(cr_config, add_handler)
	
	mux.Handle("/add/", add_handler)
	
	//

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s\n", addr)
	
	err = gohttp.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatal(err)
	}
	
}
