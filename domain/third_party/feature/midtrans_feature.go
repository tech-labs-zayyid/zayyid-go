package feature

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/third_party/model"
)

func (t ThirdPartyFeature) MidtransNotificationFeature(ctx context.Context, request model.MidtransNotificationBodyReq) (err error) {

	combination := fmt.Sprintf("%s%s%s%s", request.OrderId, request.StatusCode, request.GrossAmount, t.config.Midtrans.ServerKey)
	hash := sha512.Sum512([]byte(combination))
	signature := hex.EncodeToString(hash[:])

	if signature != request.SignatureKey {
		return sharedError.HandleError(sharedError.New(http.StatusUnauthorized, sharedConstant.ErrInvalidSignatureKey, errors.New(sharedConstant.ErrInvalidSignatureKey)))
	}

	statusCode, err := strconv.Atoi(request.StatusCode)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	grossAmount, err := strconv.ParseFloat(request.GrossAmount, 64)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	payload := model.FrontendNotificationBodyReq{
		TransactionStatus: request.TransactionStatus,
		TransactionId:     request.TransactionId,
		StatusMessage:     request.StatusMessage,
		StatusCode:        statusCode,
		PaymentType:       request.PaymentType,
		OrderId:           request.OrderId,
		GrossAmount:       grossAmount,
		FraudStatus:       request.FraudStatus,
		Bank:              request.Bank,
	}

	_, err = t.repo.GetSalesPaymentRepository(ctx, payload)
	switch err {
	case sql.ErrNoRows:
		err = t.repo.AddSalesPaymentRepository(ctx, payload)

	case nil:
		err = t.repo.UpdateSalesPaymentRepository(ctx, payload)
	}

	return
}

func (t ThirdPartyFeature) FrontendPaymentNotificationFeature(ctx context.Context, request model.FrontendNotificationBodyReq) (err error) {

	if request.TransactionStatus != sharedConstant.SETTELMENT && request.FraudStatus != sharedConstant.ACCEPT {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrStatusNotSettelment, errors.New(sharedConstant.ErrStatusNotSettelment))
	} else {
		_, err = t.repo.GetSalesPaymentRepository(ctx, request)
		switch err {
		case sql.ErrNoRows:
			err = t.repo.AddSalesPaymentRepository(ctx, request)

		case nil:
			err = t.repo.UpdateSalesPaymentRepository(ctx, request)
		}
	}

	return
}
