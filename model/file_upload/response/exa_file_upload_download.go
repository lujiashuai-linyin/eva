package response

import "eva/model/file_upload"

type ExaFileResponse struct {
	File file_upload.ExaFileUploadAndDownload `json:"file"`
}
