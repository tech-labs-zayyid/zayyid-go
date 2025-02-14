package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	userMenuModel "zayyid-go/domain/user_menu/model"
)

func (r AtomicOperation) UpdateDataUser(ctx context.Context, request userMenuModel.User) (err error) {

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
		buildQuery = append(buildQuery, " company_permission = ?")
	}

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, request.Id)
	query := fmt.Sprintf("UPDATE cms_user SET %s  WHERE id = ?", updateQuery)
	stmt, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		err = errors.New("🔥 [update] | " + err.Error())
		return
	}

	defer stmt.Close()

	return
}
