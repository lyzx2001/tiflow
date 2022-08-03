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
// limitations under the License

package factory

import (
	"context"
	"net/url"
	"strings"

	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/cdc/sinkv2/eventsink"
	"github.com/pingcap/tiflow/cdc/sinkv2/eventsink/mq"
	"github.com/pingcap/tiflow/cdc/sinkv2/eventsink/mq/dmlproducer"
	"github.com/pingcap/tiflow/cdc/sinkv2/tablesink"
	"github.com/pingcap/tiflow/pkg/config"
	cerror "github.com/pingcap/tiflow/pkg/errors"
	"github.com/pingcap/tiflow/pkg/kafka"
)

// sinkType is the type of sink.
type sinkType int

const (
	mqSink sinkType = iota + 1
	txnSink
)

// SinkFactory is the factory of sink.
// It is responsible for creating sink and closing it.
// Because there is no way to convert the eventsink.EventSink[*model.RowChangedEvent]
// to eventsink.EventSink[eventsink.TableEvent].
// So we have to use this factory to create and store the sink.
type SinkFactory struct {
	sinkType sinkType
	mqSink   eventsink.EventSink[*model.RowChangedEvent]
	txnSink  eventsink.EventSink[*model.SingleTableTxn]
}

// New creates a new SinkFactory by schema.
func New(ctx context.Context,
	sinkURIStr string,
	config *config.ReplicaConfig,
	errCh chan error,
) (*SinkFactory, error) {
	sinkURI, err := getSinkURIAndAdjustConfigWithSinkURI(sinkURIStr, config)
	if err != nil {
		return nil, err
	}

	s := &SinkFactory{}
	schema := strings.ToLower(sinkURI.Scheme)
	// TODO: add more sink factory here.
	switch schema {
	case "kafka", "kafka+ssl":
		mqs, err := mq.NewKafkaDMLSink(ctx, sinkURI, config, errCh,
			kafka.NewSaramaAdminClient, dmlproducer.NewKafkaDMLProducer)
		if err != nil {
			return nil, err
		}
		s.mqSink = mqs
		s.sinkType = mqSink
	default:
		return nil,
			cerror.ErrSinkURIInvalid.GenWithStack("the sink scheme (%s) is not supported", schema)
	}
	return s, nil
}

// CreateTableSink creates a TableSink by schema.
func (s *SinkFactory) CreateTableSink(tableID model.TableID) tablesink.TableSink {
	switch s.sinkType {
	case mqSink:
		// We have to indicate the type here, otherwise it can not be compiled.
		return tablesink.New[*model.RowChangedEvent](tableID,
			s.mqSink, &eventsink.RowChangeEventAppender{})
	case txnSink:
		return tablesink.New[*model.SingleTableTxn](tableID,
			s.txnSink, &eventsink.TxnEventAppender{})
	default:
		panic("unknown sink type")
	}
}

// Close closes the sink.
func (s *SinkFactory) Close() error {
	switch s.sinkType {
	case mqSink:
		return s.mqSink.Close()
	case txnSink:
		return s.txnSink.Close()
	default:
		panic("unknown sink type")
	}
}

func getSinkURIAndAdjustConfigWithSinkURI(sinkURIStr string,
	config *config.ReplicaConfig,
) (*url.URL, error) {
	// parse sinkURI as a URI
	sinkURI, err := url.Parse(sinkURIStr)
	if err != nil {
		return nil, cerror.WrapError(cerror.ErrSinkURIInvalid, err)
	}
	if err := config.ValidateAndAdjust(sinkURI); err != nil {
		return nil, err
	}

	return sinkURI, nil
}