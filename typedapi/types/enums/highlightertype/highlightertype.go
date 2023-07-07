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

// Package highlightertype
package highlightertype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/_global/search/_types/highlighting.ts#L80-L86
type HighlighterType struct {
	Name string
}

var (
	Plain = HighlighterType{"plain"}

	Fvh = HighlighterType{"fvh"}

	Unified = HighlighterType{"unified"}
)

func (h HighlighterType) MarshalText() (text []byte, err error) {
	return []byte(h.String()), nil
}

func (h *HighlighterType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "plain":
		*h = Plain
	case "fvh":
		*h = Fvh
	case "unified":
		*h = Unified
	default:
		*h = HighlighterType{string(text)}
	}

	return nil
}

func (h HighlighterType) String() string {
	return h.Name
}
