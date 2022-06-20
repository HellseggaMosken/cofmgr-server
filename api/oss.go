package api

import (
	"cofmgr/service/ossservice"

	"github.com/gin-gonic/gin"
)

// @Tags OSS
// @Summary OSS Upload URL
// @Description Get a signed oss url for upload
// @Router /oss/upload [post]
// @Param token header string                         true "User or Admin Token"
// @Param body  body ossservice.GetUploadURLService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.OSSSignedURL "Sucess"
// @Failure 410 {string} string                  "Unsuported file type. Only suport doc/docx/pdf"
func OSSGetUploadURL(c *gin.Context) {
	var s ossservice.GetUploadURLService
	if !bind(c, &s) {
		return
	}
	signed, code := s.GetUploadURL()
	c.JSON(code, signed)
}

// @Tags OSS
// @Summary OSS Download URL
// @Description Get a signed oss url for download
// @Router /oss/download [post]
// @Param token header string                           true "User or Admin Token"
// @Param body  body ossservice.GetDownloadURLService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.OSSSignedURL "Sucess"
// @Failure 404 {string} string                  "Target file not exists"
func OSSGetUDownloaddURL(c *gin.Context) {
	var s ossservice.GetDownloadURLService
	if !bind(c, &s) {
		return
	}
	signed, code := s.GetDownloadURL()
	c.JSON(code, signed)
}
