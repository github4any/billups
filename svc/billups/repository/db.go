// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createScoreboardStmt, err = db.PrepareContext(ctx, createScoreboard); err != nil {
		return nil, fmt.Errorf("error preparing query CreateScoreboard: %w", err)
	}
	if q.deleteScoreboardStmt, err = db.PrepareContext(ctx, deleteScoreboard); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteScoreboard: %w", err)
	}
	if q.getScoreboardStmt, err = db.PrepareContext(ctx, getScoreboard); err != nil {
		return nil, fmt.Errorf("error preparing query GetScoreboard: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createScoreboardStmt != nil {
		if cerr := q.createScoreboardStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScoreboardStmt: %w", cerr)
		}
	}
	if q.deleteScoreboardStmt != nil {
		if cerr := q.deleteScoreboardStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScoreboardStmt: %w", cerr)
		}
	}
	if q.getScoreboardStmt != nil {
		if cerr := q.getScoreboardStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScoreboardStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createScoreboardStmt *sql.Stmt
	deleteScoreboardStmt *sql.Stmt
	getScoreboardStmt    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createScoreboardStmt: q.createScoreboardStmt,
		deleteScoreboardStmt: q.deleteScoreboardStmt,
		getScoreboardStmt:    q.getScoreboardStmt,
	}
}
