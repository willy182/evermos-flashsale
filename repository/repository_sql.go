package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	// RepoSQL abstraction
	RepoSQL interface {
		WithTransaction(ctx context.Context, txFunc func(ctx context.Context, repo RepoSQL) error) (err error)
		Free()
	}

	repoSQLImpl struct {
		db *gorm.DB

		// register all repository from modules
	}
)

var (
	globalRepoSQL RepoSQL
)

var (
	once sync.Once
)

// SetSharedRepository set the global singleton "RepoSQL" implementation
func SetSharedRepository(db *sql.DB) {
	once.Do(func() {
		setSharedRepoSQL(db)
	})
}

// setSharedRepoSQL set the global singleton "RepoSQL" implementation
func setSharedRepoSQL(db *sql.DB) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic(err)
	}

	globalRepoSQL = NewRepositorySQL(gormDB)
}

// GetSharedRepoSQL returns the global singleton "RepoSQL" implementation
func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

// NewRepositorySQL constructor
func NewRepositorySQL(db *gorm.DB) RepoSQL {

	return &repoSQLImpl{
		db: db,
	}
}

// WithTransaction run transaction for each repository with context, include handle canceled or timeout context
func (r *repoSQLImpl) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, repo RepoSQL) error) (err error) {
	tx := r.db.Begin()
	err = tx.Error
	if err != nil {
		return err
	}

	// reinit new repository in different memory address with tx value
	manager := NewRepositorySQL(tx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		manager.Free()
	}()

	errChan := make(chan error)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := txFunc(ctx, manager); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("Canceled or timeout: %v", ctx.Err())
	case e := <-errChan:
		return e
	}
}

func (r *repoSQLImpl) Free() {
	// make nil all repository
}
