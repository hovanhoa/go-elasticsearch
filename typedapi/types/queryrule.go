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
// https://github.com/elastic/elasticsearch-specification/tree/cdb84fa39f1401846dab6e1c76781fb3090527ed

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/queryruletype"
)

// QueryRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cdb84fa39f1401846dab6e1c76781fb3090527ed/specification/query_rules/_types/QueryRuleset.ts#L37-L42
type QueryRule struct {
	Actions  QueryRuleActions            `json:"actions"`
	Criteria []QueryRuleCriteria         `json:"criteria"`
	RuleId   string                      `json:"rule_id"`
	Type     queryruletype.QueryRuleType `json:"type"`
}

func (s *QueryRule) UnmarshalJSON(data []byte) error {

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

		case "actions":
			if err := dec.Decode(&s.Actions); err != nil {
				return fmt.Errorf("%s | %w", "Actions", err)
			}

		case "criteria":
			if err := dec.Decode(&s.Criteria); err != nil {
				return fmt.Errorf("%s | %w", "Criteria", err)
			}

		case "rule_id":
			if err := dec.Decode(&s.RuleId); err != nil {
				return fmt.Errorf("%s | %w", "RuleId", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewQueryRule returns a QueryRule.
func NewQueryRule() *QueryRule {
	r := &QueryRule{}

	return r
}
