package helper

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func SetupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, dbMock, err := sqlmock.New()
	require.NoError(t, err)

	err = sqlDB.PingContext(context.Background())
	require.NoError(t, err)

	return sqlDB, dbMock
}
