// Copyright 2018 Google Inc. All Rights Reserved.
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
package postgres

import (
	"database/sql"
	"flag"
	"os"
	"testing"

	"github.com/golang/glog"
	"github.com/google/trillian/storage/postgres/testdb"
)

var DB *sql.DB

func TestMain(m *testing.M) {
	flag.Parse()
	if !testdb.PGAvailable() {
		glog.Errorf("PG not available, skipping all PG storage tests")
		return
	}
	DB = testdb.OpenTestDBOrDie()
	defer DB.Close()
	ec := m.Run()
	os.Exit(ec)
}
