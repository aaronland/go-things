package things

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"	
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"github.com/tidwall/gjson"
	"sync"
	"time"
)

type Thing struct {
	DepictsId int64 `json:"depicts_id"`
	Note string `json:"note"`
}

type Geometry interface{}

type Properties map[string]interface{}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

func NewThing(ctx context.Context, depicts geojson.Feature, thing_text string) (geojson.Feature, error) {

	thing_id := int64(999)

	thing_name := fmt.Sprintf("Things I like about %s #%d", depicts.Name(), thing_id)

	now := time.Now()
	ymd := now.Format("2006-01-02")
	ts := now.Unix()
	
	depicts_id := whosonfirst.Id(depicts)
	hierarchies := whosonfirst.Hierarchies(depicts)

	country := whosonfirst.Country(depicts)
	placetype := whosonfirst.Placetype(depicts)
	// repo := whosonfirst.Repo(depicts)	
	source_geom := whosonfirst.Source(depicts)

	props := make(map[string]interface{})

	depicts_ids := make([]int64, 0)

	depicts_map := new(sync.Map)
	depicts_map.Store(depicts_id, true)

	pt, err := placetypes.GetPlacetypeByName(placetype)

	if err != nil {
		return nil, err
	}

	roles := []string{
		"common",
		"optional",
		"common_optional",
	}

	for _, d := range placetypes.AncestorsForRoles(pt, roles) {

		k := fmt.Sprintf("%s_id", d.Name)

		for _, hier := range hierarchies {

			id, ok := hier[k]

			if ok {
				depicts_map.LoadOrStore(id, true)
			}
		}
	}

	depicts_map.Range(func(k interface{}, v interface{}) bool {
		id := k.(int64)
		depicts_ids = append(depicts_ids, id)
		return true
	})

	props["wof:id"] = thing_id
	props["wof:name"] = thing_name

	props["things:text"] = thing_text
	props["things:created"] = ts
	props["things:lastmodified"] = ts	
	props["things:placetype"] = placetype
	
	props["edtf:inception"] = ymd
	props["edtf:cessation"] = ymd

	props["wof:placetype"] = "custom"
	props["wof:parent_id"] = depicts_id
	props["wof:country"] = country
	props["wof:depicts"] = depicts_ids
	props["wof:hierarchy"] = hierarchies

	// how to deal w/ wof:repo
	// props["wof:parent_repo"] = repo
	// props["wof:repo"] = "custom"

	props["iso:country"] = country
	props["src:geom"] = source_geom

	geom_rsp := gjson.GetBytes(depicts.Bytes(), "geometry")

	if !geom_rsp.Exists(){
		return nil, errors.New("Missing geometry")
	}

	geom := geom_rsp.Value()
	
	f := &Feature{
		Type:       "Feature",
		Geometry:   geom,
		Properties: props,
	}

	enc_f, err := json.Marshal(f)

	if err != nil {
		return nil, err
	}

	// see the way we're not trying to return a WOF specific
	// feature? that is on purpose since we have no idea what
	// sort of properties (notably  wof:placetype) have been
	// set above (20191216/thisisaaronland)

	br := bytes.NewReader(enc_f)
	return feature.LoadGeoJSONFeatureFromReader(br)
}
