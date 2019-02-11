// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package repository

import (
	"database/sql"
	"github.com/goodrain/go-demo/dbinfo"
	"github.com/sirupsen/logrus"
)

type mysqlDBInfoRepo struct {
	DB *sql.DB
}

// NewMysqlDBInfoRepository will create an implementation of author.Repositorier
func NewMysqlDBInfoRepository(db *sql.DB) dbinfo.Repositorier {
	return &mysqlDBInfoRepo{
		DB: db,
	}
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (m *mysqlDBInfoRepo) Ping() bool {
	if  m.DB == nil {
		return false
	}

	err := m.DB.Ping()
	if err != nil {
		logrus.Debugf("error pinging sql.DB: %v", err)
		return false
	}
	return true
}
