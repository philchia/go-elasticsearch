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

// ClusterOperatingSystemName type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/cluster/stats/types.ts#L253-L256
type ClusterOperatingSystemName struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

func (s *ClusterOperatingSystemName) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterOperatingSystemName returns a ClusterOperatingSystemName.
func NewClusterOperatingSystemName() *ClusterOperatingSystemName {
	r := &ClusterOperatingSystemName{}

	return r
}
