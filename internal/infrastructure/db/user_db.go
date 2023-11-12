package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/thumperq/golib/database"
)

type UserDB interface {
	DeleteUsers(ctx context.Context) error
}
type UserDb struct {
	pgDb database.PgDB
}

func NewUserDb(pgDb database.PgDB) UserDB {
	return UserDb{
		pgDb: pgDb,
	}
}

func (db UserDb) DeleteUsers(ctx context.Context) error {
	err := db.pgDb.WithTransaction(ctx, func(tx pgx.Tx) error {
		sql, args, err := sb.Delete("users").ToSql()
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
