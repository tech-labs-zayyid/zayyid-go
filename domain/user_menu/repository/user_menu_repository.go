package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"middleware-cms-api/domain/user_menu/model"

	"github.com/jmoiron/sqlx"
)

func (r userMenuRepository) OpenTransaction() (tx *sqlx.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}
	ctx := context.Background()
	tx, _ = r.database.DB.BeginTxx(ctx, &sqlTxOptions)
	return
}

func (r userMenuRepository) RollbackTransaction(tx *sqlx.Tx) (rollBack error) {

	rollBack = tx.Rollback()

	return
}

func (r userMenuRepository) CommitTransaction(tx *sqlx.Tx) (commit error) {

	commit = tx.Rollback()

	return
}

func (r userMenuRepository) GetList(request model.User) (response model.GetList, err error) {

	var args []interface{}

	query := `
	SELECT
		a.id,
		a.name,
		a.role,
		a.created_at,
		a.is_active,
		a.company_permission
	FROM
		cms_user a
	WHERE
		1 = 1
	`

	queryCount := `
		SELECT
			COUNT(a.id)
		FROM
			cms_user a
		WHERE
			1 = 1
	`

	if request.Search != "" {
		appTypeCondition := ` AND (a.id LIKE ? OR a.name LIKE ?) `
		query += appTypeCondition
		queryCount += appTypeCondition
		args = append(args, "%"+request.Search+"%")
		args = append(args, "%"+request.Search+"%")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var count int64

		err := r.database.Get(&count, queryCount, args...)
		if err != nil {
			return
		}

		response.Pagination.TotalRows = count
		response.Pagination.Limit = request.Limit
		response.Pagination.Page = request.Page
		response.Pagination.TotalPage = int64(math.Ceil(float64(count) / float64(request.Limit)))
	}()

	go func() {
		defer wg.Done()
		var offset int64
		offset = 0
		if request.Page > 1 {
			offset = (request.Limit * (request.Page - 1))
		}
		query += fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d`, " a.id ", " DESC ", request.Limit, offset)
		// logger.LogInfo(constant.QUERY, baseQuery)
		err = r.database.Select(&response.Data, query, args...)
		if err != nil {
			err = errors.New("🔥 [get list] | " + err.Error())
			return
		}
	}()

	wg.Wait()

	return
}

func (r userMenuRepository) GetDataById(id string) (response model.User, err error) {

	err = r.database.Get(&response, `
		SELECT
			a.id,
			a.name,
			a.role,
			a.created_at,
			a.is_active,
			a.company_permission
		FROM
			cms_user a
		WHERE
			a.id = ?
		LIMIT 1;`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}

		return response, errors.New("🔥 [get data by id] | " + err.Error())
	}

	return response, nil
}

func (r userMenuRepository) GetDataUserMenuById(id string) (response []model.UserAuth, err error) {

	err = r.database.Select(&response, `
		SELECT
			IFNULL(b.menu_id, 0) AS menu_id,
			b.permission
		FROM
			cms_user a
			LEFT JOIN cms_user_auth b ON b.user_id = a.id
		WHERE
			a.id = ?`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}

		return response, errors.New("🔥 [get data by id] | " + err.Error())
	}

	return response, nil
}

func (r userMenuRepository) GetDataMenuById(id int) (response model.Menu, err error) {

	err = r.database.Get(&response, `
		SELECT
			a.id,
			a.name,
			a.description,
			a.permission
		FROM
			cms_menu a
		WHERE
			a.id = ?
		LIMIT 1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}

		return response, errors.New("🔥 [get data by id] | " + err.Error())
	}

	return response, nil
}

func (r userMenuRepository) GetListDataMenu() (response []model.Menu, err error) {

	err = r.database.Select(&response, `
		SELECT
			a.id,
			a.name,
			a.description,
			a.permission
		FROM
			cms_menu a`)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}

		return response, errors.New("🔥 [get data by id] | " + err.Error())
	}

	return response, nil
}

func (r userMenuRepository) CreateData(ctx context.Context, request model.User, tx *sqlx.Tx) (err error) {

	query := `
		INSERT INTO cms_user(
			id,
			name,
			role,
			created_at,
			is_active,
			company_permission
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		request.Id,
		request.Name,
		request.Role,
		sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		request.TempIsActive,
		request.CompanyPermision,
	)

	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}

	return
}

func (r userMenuRepository) CreateDataUserAuth(ctx context.Context, request model.UserAuth, tx *sqlx.Tx) (err error) {

	query := `
		INSERT INTO cms_user_auth(
			user_id,
			menu_id,
			permission
		) VALUES (?, ?, ?)
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		request.UserId,
		request.MenuId,
		request.Permision,
	)

	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}

	return
}

func (r userMenuRepository) UpdateData(ctx context.Context, request model.User, tx *sqlx.Tx) (err error) {

	var args []interface{}

	buildQuery := []string{}
	if request.Name != "" {
		args = append(args, request.Name)
		buildQuery = append(buildQuery, " name = ?")
	}

	if request.Role != "" {
		args = append(args, request.Role)
		buildQuery = append(buildQuery, " role = ?")
	}

	args = append(args, request.TempIsActive)
	buildQuery = append(buildQuery, " is_active = ?")

	if string(request.CompanyPermision) != "" {
		args = append(args, request.CompanyPermision)
		buildQuery = append(buildQuery, " config = ?")
	}

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, request.Id)
	query := fmt.Sprintf("UPDATE cms_user SET %s  WHERE id = ?", updateQuery)
	stmt, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		err = errors.New("🔥 [update] | " + err.Error())
		return
	}

	defer stmt.Close()

	return
}

func (r userMenuRepository) DeleteData(ctx context.Context, id string, tx *sqlx.Tx) (err error) {

	query := `
		DELETE FROM
			cms_user
		WHERE
			id = ?
	`

	stmt, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		err = errors.New("🔥 [delete] | " + err.Error())
		return
	}

	defer stmt.Close()

	return
}

func (r userMenuRepository) DeleteDataUserAuth(ctx context.Context, userId string, tx *sqlx.Tx) (err error) {

	query := `
		DELETE FROM
			cms_user_auth
		WHERE
			user_id = ?
	`

	stmt, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		err = errors.New("🔥 [delete] | " + err.Error())
		return
	}

	defer stmt.Close()

	return
}
