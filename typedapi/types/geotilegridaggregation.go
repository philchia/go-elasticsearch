// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// GeoTileGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/_types/aggregations/bucket.ts#L192-L198
type GeoTileGridAggregation struct {
	Bounds    GeoBounds `json:"bounds,omitempty"`
	Field     *string   `json:"field,omitempty"`
	Meta      Metadata  `json:"meta,omitempty"`
	Name      *string   `json:"name,omitempty"`
	Precision *int      `json:"precision,omitempty"`
	ShardSize *int      `json:"shard_size,omitempty"`
	Size      *int      `json:"size,omitempty"`
}

func (s *GeoTileGridAggregation) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "bounds":
			if err := dec.Decode(&s.Bounds); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "precision":
			if err := dec.Decode(&s.Precision); err != nil {
				return err
			}

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewGeoTileGridAggregation returns a GeoTileGridAggregation.
func NewGeoTileGridAggregation() *GeoTileGridAggregation {
	r := &GeoTileGridAggregation{}

	return r
}
