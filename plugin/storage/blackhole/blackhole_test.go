// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blackhole

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
)

func withBlackhole(f func(store *Store)) {
	f(NewStore())
}

func TestStoreGetDependencies(t *testing.T) {
	withBlackhole(func(store *Store) {
		links, err := store.GetDependencies(context.Background(), time.Now(), time.Hour)
		assert.NoError(t, err)
		assert.Empty(t, links)
	})
}

func TestStoreWriteSpan(t *testing.T) {
	withBlackhole(func(store *Store) {
		err := store.WriteSpan(context.Background(), nil)
		assert.NoError(t, err)
	})
}

func TestStoreGetTrace(t *testing.T) {
	withBlackhole(func(store *Store) {
		trace, err := store.GetTrace(context.Background(), model.NewTraceID(1, 2))
		assert.Error(t, err)
		assert.Nil(t, trace)
	})
}

func TestStoreGetServices(t *testing.T) {
	withBlackhole(func(store *Store) {
		serviceNames, err := store.GetServices(context.Background())
		assert.NoError(t, err)
		assert.Len(t, serviceNames, 0)
	})
}

func TestStoreGetAllOperations(t *testing.T) {
	withBlackhole(func(store *Store) {
		operations, err := store.GetOperations(
			context.Background(),
			spanstore.OperationQueryParameters{},
		)
		assert.NoError(t, err)
		assert.Len(t, operations, 0)
	})
}

func TestStoreFindTraces(t *testing.T) {
	withBlackhole(func(store *Store) {
		traces, err := store.FindTraces(context.Background(), nil)
		assert.NoError(t, err)
		assert.Len(t, traces, 0)
	})
}

func TestStoreFindTraceIDs(t *testing.T) {
	withBlackhole(func(store *Store) {
		traceIDs, err := store.FindTraceIDs(context.Background(), nil)
		assert.NoError(t, err)
		assert.Len(t, traceIDs, 0)
	})
}
