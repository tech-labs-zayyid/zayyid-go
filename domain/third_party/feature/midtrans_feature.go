package feature

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/third_party/model"
)

func (t ThirdPartyFeature) NotificationFeature(ctx context.Context, request model.MidtransBodyReq) (err error) {

	log.Println("request.TransactionStatus : ", request.TransactionStatus)
	log.Println("request.TransactionId : ", request.TransactionId)
	log.Println("request.StatusMessage : ", request.StatusMessage)
	log.Println("request.StatusCode : ", request.StatusCode)
	log.Println("request.SignatureKey : ", request.SignatureKey)
	log.Println("request.PaymentType : ", request.PaymentType)
	log.Println("request.OrderId : ", request.OrderId)
	log.Println("request.GrossAmount : ", request.GrossAmount)
	log.Println("request.FraudStatus : ", request.FraudStatus)
	log.Println("request.Bank : ", request.Bank)

	input := fmt.Sprintf("%s%s%s%s", request.OrderId, request.StatusCode, request.GrossAmount, t.config.Midtrans.ServerKey)

	// Menghitung hash SHA-512
	hash := sha512.Sum512([]byte(input))
	signature := hex.EncodeToString(hash[:])

	log.Println("signature : ", signature)

	if signature != request.SignatureKey {
		err = sharedError.New(http.StatusUnauthorized, sharedConstant.ErrInvalidSignatureKey, errors.New(sharedConstant.ErrInvalidSignatureKey))
		return sharedError.ResponseErrorWithContext(ctx, err, t.SlackConf)
	}

	return
}
