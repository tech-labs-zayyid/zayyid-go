package feature

import (
	"context"
	sharedHelperErr "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
	sharedModel "zayyid-go/domain/shared/model"
	sharedHelperRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) CreateAgentFeature(ctx context.Context, payload model.RegisterRequest, userId string) (resp model.UserRes, err error) {

	// if role not agent return error
	if payload.Role != "agent" {
		err = sharedHelperErr.New(403, "Role should be an agent", nil)
		return
	}

	// Get data by userId 
	userLogIn, err := f.repo.GetDataUserByUserId(ctx, userId)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return 
	}

	// if agent was register generate referal code
	referalCode, errGenerateReferal := sharedHelper.GenerateRandomString(10)
	if errGenerateReferal != nil {
		err = sharedHelperErr.HandleError(errGenerateReferal)
		return
	}

	payload.ReferalCode = referalCode

	// generate random string for password
	password, err := sharedHelper.GenerateRandomString(5)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// encrypt password
	encryptedPassword, err := sharedHelper.HashPassword(password)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// set up password
	payload.Password = encryptedPassword

	// generate agent id
	agentId := sharedHelperRepo.GenerateUuidAsIdTable()

	// Start transaction 
	trx := f.repo.OpenTransaction(ctx)

	// call repo to register user 
	err = f.repo.RegisterRepositoryTransaction(ctx, payload, agentId.String(), trx)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// mapping sales agent 
	err = f.repo.MappingSalesAgent(ctx, userId, agentId.String(), userLogIn.Email, trx)
	if err != nil {
		// rollback transaction 
		trx.Rollback()
		err = sharedHelperErr.HandleError(err)
		return 
	}

	// commit transaction 
	trx.Commit()

	// get one user by userid
	user, err := f.repo.GetUserById(ctx, agentId.String())
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// Generate token
	token, err := sharedHelper.GenerateToken(userId, payload.Role)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// generate refresh token
	refreshToken, err := sharedHelper.GenerateRefreshToken(user.Id, user.Role)
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return
	}

	// TODO: send password to agent email
	resp = model.UserRes{
		Id:             agentId.String(),
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		Role:           user.Role,
		WhatsAppNumber: user.WhatsAppNumber,
		CreatedAt:      user.CreatedAt,
		CreatedBy:      user.CreatedBy,
		TokenData: &model.TokenRes{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}

	return

}

func (f UserFeature) GetAgentFeature(ctx context.Context, query model.QueryAgentList, userId string) (resp model.AgentListPagination, err error) {
	
	agents, err := f.repo.GetAgentRepository(ctx, query, userId) 
	if err != nil {
		err = sharedHelperErr.HandleError(err)
		return 
	}

	totalPages := (len(agents) + query.Limit - 1) / query.Limit

	resp = model.AgentListPagination{
		Data: agents,
		Pagination: sharedModel.Pagination{
			Limit:     query.Limit,
			Page:      query.Page,
			TotalPage: totalPages,
			TotalRows:     len(agents),
		},
	}

	return 
	
}