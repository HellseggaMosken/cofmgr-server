package ossservice

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"cofmgr/logger"
	"cofmgr/service"
	"cofmgr/service/serializer"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type GetUploadURLService struct {
	Filename string `json:"filename"  binding:"required"`
}

func (s *GetUploadURLService) GetUploadURL() (*serializer.OSSSignedURL, service.Status) {
	ext := strings.ToLower(filepath.Ext(s.Filename))
	if ext != ".doc" && ext != ".docx" && ext != ".pdf" {
		return nil, service.StatusOSSUnsupportedFileType
	}

	signed := &serializer.OSSSignedURL{
		Filename:     fmt.Sprintf("[%d]%s", time.Now().Unix(), s.Filename),
		ContentType:  mime.TypeByExtension(ext),
		Method:       "put",
		ExpiredInSec: 600,
	}

	options := []oss.Option{
		oss.ContentType(signed.ContentType),
	}

	url, err := bucket.SignURL(signed.Filename,
		oss.HTTPPut, signed.ExpiredInSec, options...)
	if err != nil {
		logger.Panic("when sign oss upload url:", signed, err)
	}

	signed.URL = url

	return signed, service.StatusOK
}
