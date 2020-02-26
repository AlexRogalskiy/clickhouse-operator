// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clickhouse

import (
	"sync"
)

var (
	dbConnectionPool               = sync.Map{}
	dbConnectionPoolEntryInitMutex = sync.Mutex{}
)

func GetPooledDBConnection(params *CHConnectionParams) *CHConnection {
	if connection, existed := dbConnectionPool.Load(params); existed {
		return connection.(*CHConnection)
	}

	dbConnectionPoolEntryInitMutex.Lock()
	defer dbConnectionPoolEntryInitMutex.Unlock()

	// Double check
	if connection, existed := dbConnectionPool.Load(params); existed {
		return connection.(*CHConnection)
	}

	connection := NewConnection(params)

	dbConnectionPool.Store(params, connection)

	return connection
}

// TODO we need to be able to remove entries from the pool
func DropHost(host string) {

}
