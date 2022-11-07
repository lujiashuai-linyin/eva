package response

import "eva/model/file_upload"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File file_upload.ExaFile `json:"file"`
}
