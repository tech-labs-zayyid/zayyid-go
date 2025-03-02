package repository

import (
	"context"
	"fmt"
	"strings"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/domain/third_party/model"
	"zayyid-go/infrastructure/logger"
)

func (t thirdPartyRepository) AddSalesPaymentRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (err error) {

	args := []interface{}{
		request.TransactionId,
		request.TransactionStatus,
		request.StatusMessage,
		request.StatusCode,
		request.PaymentType,
		request.OrderId,
		request.GrossAmount,
		request.FraudStatus,
		request.Bank,
	}

	query := `
		INSERT INTO product_marketing.sales_payment (
			transaction_id,
			transaction_status,
			status_message,
			status_code,
			payment_type,
			order_id,
			gross_amount,
			fraud_status,
			bank)
		VALUES
			($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	stmt, err := t.database.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	return
}

func (t thirdPartyRepository) UpdateTestimoniRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (err error) {

	args := []interface{}{}
	buildQuery := []string{}

	if request.TransactionStatus != "" {
		args = append(args, request.TransactionStatus)
		buildQuery = append(buildQuery, " fullname = $1")
	}
	if request.StatusMessage != "" {
		args = append(args, request.StatusMessage)
		buildQuery = append(buildQuery, " status_message = $2")
	}
	if request.StatusCode != 0 {
		args = append(args, request.StatusCode)
		buildQuery = append(buildQuery, " status_code = $3")
	}
	if request.PaymentType != "" {
		args = append(args, request.PaymentType)
		buildQuery = append(buildQuery, " payment_type = $4")
	}
	if request.GrossAmount != 0 {
		args = append(args, request.GrossAmount)
		buildQuery = append(buildQuery, " gross_amount = $5")
	}
	if request.FraudStatus != "" {
		args = append(args, request.FraudStatus)
		buildQuery = append(buildQuery, " fraud_atatus = $6")
	}
	if request.Bank != "" {
		args = append(args, request.Bank)
		buildQuery = append(buildQuery, " bank = $7")
	}

	buildQuery = append(buildQuery, " updated_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, request.TransactionId)
	args = append(args, request.OrderId)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_payment SET %s  WHERE transaction_id = ? AND order_id = ? `, updateQuery)

	logger.LogInfo(constant.QUERY, query)
	stmt, err := t.database.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	return
}

func (t thirdPartyRepository) GetSalesPaymentRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (response model.SalesPaymentResp, err error) {

	query := `
		SELECT
			transaction_status,
			transaction_id,
			status_message,
			status_code,
			payment_type,
			order_id,
			gross_amount,
			fraud_status,
			bank,
			created_at,
			updated_at
		FROM
			product_marketing.sales_payment
		WHERE
			transaction_id = $1
			AND order_id = $2`
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response, request.TransactionId, request.OrderId)
	if err != nil {
		return
	}

	return
}
