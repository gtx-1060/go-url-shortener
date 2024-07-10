package daos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"time"
)

var (
	ErrMaxRetriesExceed = errors.New("cant connect to db in non-concurrent mode because it's busy")
)

func (dao *Dao) execTx(ctx context.Context, options *sql.TxOptions, fn func(query RWQuery) error) error {
	tx, err := dao.nonConcurrentDb.BeginTx(ctx, options)
	if err != nil {
		return err
	}
	err = fn(RWQuery{Query{tx}})
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (dao *Dao) RetryTillFree(fn func(retryDao *Dao) error) error {
	var i uint
	for i = 0; i < dao.maxRetries; i++ {
		err := fn(dao)
		if err == nil {
			return nil
		}
		sqlErr, ok := err.(sqlite3.Error)
		if !ok || sqlErr.Code != sqlite3.ErrBusy || sqlErr.Code != sqlite3.ErrLocked {
			return err
		}
		time.Sleep(time.Duration(dao.retryTimeoutMs) * time.Millisecond)
	}
	return ErrMaxRetriesExceed
}

func (dao *Dao) StartTx(ctx context.Context, options *sql.TxOptions, op func(query RWQuery) error) error {
	return dao.RetryTillFree(func(retryDao *Dao) error {
		return retryDao.execTx(ctx, options, op)
	})
}
