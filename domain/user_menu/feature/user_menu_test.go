package feature

import (
	"context"
	"encoding/json"
	"testing"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/domain/user_menu/helper"
	"zayyid-go/domain/user_menu/model"
	mockRepo "zayyid-go/domain/user_menu/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
)

func Test_GetList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserMenu := mockRepo.NewMockUserMenuRepository(mockCtrl)

	t.Run("Data not found or is empty", func(t *testing.T) {
		expected := model.GetList{
			Data: []model.User{},
			Pagination: sharedModel.Pagination{
				Limit:     5,
				Page:      1,
				TotalPage: 1,
				TotalRows: 0,
			},
			StatusCode: 200,
		}
		mockUserMenu.EXPECT().GetList(gomock.Any()).Return(expected, nil)

		param := model.User{
			Id:               "",
			Name:             "",
			Role:             "",
			IsActive:         false,
			TempIsActive:     0,
			CompanyPermision: []byte{},
			MenuId:           []int{},
		}
		dataList, err := mockUserMenu.GetList(param)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if dataList.StatusCode != expected.StatusCode {
			t.Errorf("status code : %d", dataList.StatusCode)
		}
	})

	t.Run("Data search found", func(t *testing.T) {
		dataExpected := []model.User{}
		dataExpected = append(dataExpected, model.User{
			Id:               "pamungkas@mail.com",
			Name:             "pamungkas",
			Role:             "admin",
			IsActive:         true,
			CompanyPermision: []byte{},
			MenuId:           []int{1, 2, 3},
			StatusCode:       200,
		})

		expected := model.GetList{
			Data: dataExpected,
			Pagination: sharedModel.Pagination{
				Limit:     5,
				Page:      1,
				TotalPage: 1,
				TotalRows: 1,
			},
			StatusCode: 200,
		}

		param := model.User{
			Search: "pamungkas",
		}

		mockUserMenu.EXPECT().GetList(param).Return(expected, nil)
		dataList, err := mockUserMenu.GetList(param)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if dataList.StatusCode != expected.StatusCode {
			t.Errorf("status code : %d", dataList.StatusCode)
		}
	})
}

func Test_GetDataById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserMenu := mockRepo.NewMockUserMenuRepository(mockCtrl)
	id := "pamungkas@mail.com"

	t.Run("Data not found", func(t *testing.T) {
		expected := model.User{
			Id:               "",
			Name:             "",
			Role:             "",
			IsActive:         false,
			TempIsActive:     0,
			CompanyPermision: []byte{},
			MenuId:           []int{},
			StatusCode:       200,
		}
		mockUserMenu.EXPECT().GetDataById(id).Return(expected, nil)
		dataUser, err := mockUserMenu.GetDataById(id)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expectedMenu := []model.UserAuth{}
		mockUserMenu.EXPECT().GetDataUserMenuById(id).Return(expectedMenu, nil)
		_, errMenu := mockUserMenu.GetDataUserMenuById(id)
		if errMenu != nil {
			t.Errorf("unexpected error: %v", errMenu)
		}

		expectedResponse := &model.UserResponse{
			Id:               dataUser.Id,
			Name:             dataUser.Name,
			Role:             dataUser.Role,
			CreatedAt:        dataUser.CreatedAt,
			IsActive:         dataUser.IsActive,
			CompanyPermision: []string{},
			MenuId:           []int{},
			StatusCode:       200,
		}

		if expectedResponse.StatusCode != expected.StatusCode {
			t.Errorf("status code : %d", dataUser.StatusCode)
		}
	})

	t.Run("Data search found", func(t *testing.T) {
		expected := model.User{
			Id:               id,
			Name:             "pamungkas",
			Role:             "admin",
			IsActive:         true,
			TempIsActive:     1,
			CompanyPermision: []byte{},
			MenuId:           []int{1, 2, 3},
			StatusCode:       200,
		}
		mockUserMenu.EXPECT().GetDataById(id).Return(expected, nil)
		dataUser, err := mockUserMenu.GetDataById(id)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expectedMenu := []model.UserAuth{}
		expectedMenu = append(expectedMenu, model.UserAuth{
			Id:        int64(1),
			UserId:    "pamungkas@mail.com",
			MenuId:    1,
			Permision: []byte{},
		})
		mockUserMenu.EXPECT().GetDataUserMenuById(id).Return(expectedMenu, nil)
		dataMenu, errMenu := mockUserMenu.GetDataUserMenuById(id)
		if errMenu != nil {
			t.Errorf("unexpected error: %v", errMenu)
		}

		var menuId []int
		for _, v := range dataMenu {
			menuId = append(menuId, v.MenuId)
		}

		var companyPermission []string
		if len(dataUser.CompanyPermision) > 0 {
			errPermission := json.Unmarshal(dataUser.CompanyPermision, &companyPermission)
			if errPermission != nil {
				t.Errorf("unexpected error: %v", errPermission)
			}
		}

		expectedResponse := &model.UserResponse{
			Id:               dataUser.Id,
			Name:             dataUser.Name,
			Role:             dataUser.Role,
			CreatedAt:        dataUser.CreatedAt,
			IsActive:         dataUser.IsActive,
			CompanyPermision: companyPermission,
			MenuId:           menuId,
			StatusCode:       200,
		}

		if expectedResponse.Id == "" {
			t.Errorf("status code : %d", dataUser.StatusCode)
		}
	})
}

func Test_UpsertDataManual(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserMenu := mockRepo.NewMockUserMenuRepository(mockCtrl)
	id := "pamungkas@mail.com"
	ctx := context.Background()
	var tx *sqlx.Tx

	expectedUser := model.User{
		Id:               id,
		Name:             "pamungkas",
		Role:             "admin",
		IsActive:         true,
		TempIsActive:     1,
		CompanyPermision: []byte{},
		MenuId:           []int{1, 2, 3},
		StatusCode:       200,
	}
	mockUserMenu.EXPECT().GetDataById(id).Return(expectedUser, nil)
	dataUser, err := mockUserMenu.GetDataById(id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	request := model.User{
		Id:               id,
		Name:             "pamungkas",
		Role:             "admin",
		IsActive:         true,
		TempIsActive:     1,
		CompanyPermision: []byte{},
		MenuId:           []int{1, 2, 3},
		StatusCode:       200,
	}

	param, errCheck := helper.CheckRequest(request)
	if errCheck != nil {
		t.Errorf("unexpected error: %v", errCheck)
	}

	mockUserMenu.EXPECT().OpenTransaction().Return(tx)
	startTransaction := mockUserMenu.OpenTransaction()

	if dataUser.Id != "" {
		mockUserMenu.EXPECT().DeleteDataUserAuth(ctx, dataUser.Id, startTransaction)
		err = mockUserMenu.DeleteDataUserAuth(ctx, dataUser.Id, startTransaction)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	for _, v := range param.MenuId {
		expectedMenu := model.Menu{
			Id:          int64(v),
			Name:        "",
			Description: "",
			Permision:   []byte{},
		}
		mockUserMenu.EXPECT().GetDataMenuById(v).Return(expectedMenu, nil)
		dataMenu, errDataMenu := mockUserMenu.GetDataMenuById(v)
		if errDataMenu != nil {
			t.Errorf("unexpected error: %v", errDataMenu)
		}

		var permission []string
		var paramPermission []byte
		if len(dataMenu.Permision) > 0 {
			err = json.Unmarshal(dataMenu.Permision, &permission)
			if err != nil {
				mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
				mockUserMenu.RollbackTransaction(startTransaction)
				t.Errorf("unexpected error: %v", err)
			}

			paramPermission, err = json.Marshal(permission)
			if err != nil {
				mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
				mockUserMenu.RollbackTransaction(startTransaction)
				t.Errorf("unexpected error: %v", err)
			}
		}

		param := model.UserAuth{
			UserId:    param.Id,
			MenuId:    v,
			Permision: paramPermission,
		}

		mockUserMenu.EXPECT().CreateDataUserAuth(ctx, param, startTransaction).Return(nil)
		err = mockUserMenu.CreateDataUserAuth(ctx, param, startTransaction)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	t.Run("Create data success", func(t *testing.T) {
		mockUserMenu.EXPECT().CreateData(ctx, param, startTransaction).Return(nil)
		err := mockUserMenu.CreateData(ctx, param, startTransaction)
		if err != nil {
			mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
			mockUserMenu.RollbackTransaction(startTransaction)
			t.Errorf("unexpected error: %v", err)
		}

		if param.StatusCode != 200 {
			mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
			mockUserMenu.RollbackTransaction(startTransaction)
			t.Errorf("can not create data for  %s", param.Id)
		}
	})

	t.Run("Update data success", func(t *testing.T) {
		mockUserMenu.EXPECT().UpdateData(ctx, param, startTransaction).Return(nil)
		err := mockUserMenu.UpdateData(ctx, param, startTransaction)
		if err != nil {
			mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
			mockUserMenu.RollbackTransaction(startTransaction)
			t.Errorf("unexpected error: %v", err)
		}

		if param.StatusCode != 200 {
			mockUserMenu.EXPECT().RollbackTransaction(startTransaction).Return(nil)
			mockUserMenu.RollbackTransaction(startTransaction)
			t.Errorf("can not update data for  %s", param.Id)
		}
	})

	mockUserMenu.EXPECT().CommitTransaction(startTransaction).Return(nil)
	mockUserMenu.CommitTransaction(startTransaction)
}

func Test_DeleteData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserMenu := mockRepo.NewMockUserMenuRepository(mockCtrl)
	id := "pamungkas@mail.com"
	ctx := context.Background()
	var tx *sqlx.Tx

	mockUserMenu.EXPECT().OpenTransaction().Return(tx)
	startTransaction := mockUserMenu.OpenTransaction()

	t.Run("Delete data success", func(t *testing.T) {
		mockUserMenu.EXPECT().DeleteData(ctx, id, startTransaction).Return(nil)
		err := mockUserMenu.DeleteData(ctx, id, startTransaction)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_GetListMenu(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockuserMenu := mockRepo.NewMockUserMenuRepository(mockCtrl)

	t.Run("Data not found or is empty", func(t *testing.T) {
		expected := []model.Menu{}
		expected = append(expected, model.Menu{
			Id:          int64(0),
			Name:        "",
			Description: "",
			Permision:   []byte{},
		})
		mockuserMenu.EXPECT().GetListDataMenu().Return(expected, nil)
		_, err := mockuserMenu.GetListDataMenu()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

	})

	t.Run("Data found", func(t *testing.T) {
		expected := []model.Menu{}
		expected = append(expected, model.Menu{
			Id:          int64(1),
			Name:        "User Menu",
			Description: "User Menu",
			Permision:   []byte{},
		})
		mockuserMenu.EXPECT().GetListDataMenu().Return(expected, nil)
		listDataMenu, err := mockuserMenu.GetListDataMenu()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(listDataMenu) < len(expected) {
			t.Errorf("data corrupt")
		}
	})
}
