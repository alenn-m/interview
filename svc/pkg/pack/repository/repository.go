package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/alenn-m/interview/svc/pkg/pack/entity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

type Repository interface {
	Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error)
	Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Pack, error)
	List(ctx context.Context) ([]*entity.Pack, error)
}

type repository struct {
	conn *sqlx.DB
}

// Options holds options for new instance
type Options struct {
	fx.In

	PostgresDB *sqlx.DB
}

func New(o Options) Repository {
	return &repository{conn: o.PostgresDB}
}

func (r *repository) builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		RunWith(r.conn)
}

func (r *repository) Create(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	now := time.Now()
	query := r.builder().
		Insert("packs").
		Columns("name", "amount", "created_at", "updated_at").
		Values(pack.Name, pack.Amount, now, now).
		Suffix("RETURNING id, name, amount, created_at, updated_at")

	err := query.QueryRowContext(ctx).Scan(
		&pack.ID,
		&pack.Name,
		&pack.Amount,
		&pack.CreatedAt,
		&pack.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return pack, nil
}

func (r *repository) Update(ctx context.Context, pack *entity.Pack) (*entity.Pack, error) {
	now := time.Now()
	query := r.builder().
		Update("packs").
		Set("name", pack.Name).
		Set("amount", pack.Amount).
		Set("updated_at", now).
		Where(sq.Eq{"id": pack.ID}).
		Suffix("RETURNING id, name, amount, created_at, updated_at")

	err := query.QueryRowContext(ctx).Scan(
		&pack.ID,
		&pack.Name,
		&pack.Amount,
		&pack.CreatedAt,
		&pack.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return pack, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := r.builder().
		Delete("packs").
		Where(sq.Eq{"id": id})

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) Get(ctx context.Context, id int) (*entity.Pack, error) {
	query := r.builder().
		Select("id", "name", "amount", "created_at", "updated_at").
		From("packs").
		Where(sq.Eq{"id": id})

	var pack entity.Pack
	err := query.QueryRowContext(ctx).Scan(
		&pack.ID,
		&pack.Name,
		&pack.Amount,
		&pack.CreatedAt,
		&pack.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &pack, nil
}

func (r *repository) List(ctx context.Context) ([]*entity.Pack, error) {
	query := r.builder().
		Select("id", "name", "amount", "created_at", "updated_at").
		From("packs").
		OrderBy("amount DESC")

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packs []*entity.Pack
	for rows.Next() {
		var pack entity.Pack
		err := rows.Scan(
			&pack.ID,
			&pack.Name,
			&pack.Amount,
			&pack.CreatedAt,
			&pack.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		packs = append(packs, &pack)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packs, nil
}
