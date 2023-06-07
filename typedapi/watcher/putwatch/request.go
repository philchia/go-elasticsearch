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
// https://github.com/elastic/elasticsearch-specification/tree/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96

package putwatch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putwatch
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/watcher/put_watch/WatcherPutWatchRequest.ts#L30-L54
type Request struct {
	Actions        map[string]types.WatcherAction `json:"actions,omitempty"`
	Condition      *types.WatcherCondition        `json:"condition,omitempty"`
	Input          *types.WatcherInput            `json:"input,omitempty"`
	Metadata       types.Metadata                 `json:"metadata,omitempty"`
	ThrottlePeriod *string                        `json:"throttle_period,omitempty"`
	Transform      *types.TransformContainer      `json:"transform,omitempty"`
	Trigger        *types.TriggerContainer        `json:"trigger,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Actions: make(map[string]types.WatcherAction, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putwatch request: %w", err)
	}

	return &req, nil
}
