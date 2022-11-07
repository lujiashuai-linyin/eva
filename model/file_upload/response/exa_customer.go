package response

import "eva/model/file_upload"

type ExaCustomerResponse struct {
	Customer file_upload.ExaCustomer `json:"customer"`
}
