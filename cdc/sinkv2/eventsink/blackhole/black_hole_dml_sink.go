// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package blackhole

import (
	"github.com/pingcap/log"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/cdc/sinkv2/eventsink"
	"go.uber.org/zap"
)

// Assert EventSink[E event.TableEvent] implementation
var _ eventsink.EventSink[*model.RowChangedEvent] = (*sink)(nil)

type sink struct{}

// New create a black hole DML sink.
func New() *sink {
	return &sink{}
}

// WriteEvents log the events.
func (s *sink) WriteEvents(rows ...*eventsink.CallbackableEvent[*model.RowChangedEvent]) error {
	for _, row := range rows {
		log.Debug("BlockHoleSink: WriteEvents", zap.Any("row", row))
	}

	return nil
}

// Close do nothing.
func (s *sink) Close() error {
	return nil
}
