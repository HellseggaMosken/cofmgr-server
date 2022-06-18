package ossservice

import (
	"mime"
	"path/filepath"

	"cofmgr/logger"
	"cofmgr/service"
	"cofmgr/service/serializer"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type GetDownloadURLService struct {
	Filename string `json:"filename"  binding:"required"`
}

func (s *GetDownloadURLService) GetDownloadURL() (*serializer.OSSSignedURL, service.Status) {
	if exists, err := bucket.IsObjectExist(s.Filename); err != nil {
		logger.Panic("when check oss file exist:", s.Filename, err)
	} else if !exists {
		return nil, service.StatusNotFoundItem
	}

	signed := &serializer.OSSSignedURL{
		ContentType:  mime.TypeByExtension(filepath.Ext(s.Filename)),
		Filename:     s.Filename,
		Method:       "get",
		ExpiredInSec: 600,
	}

	url, err := bucket.SignURL(s.Filename, oss.HTTPGet, signed.ExpiredInSec)
	if err != nil {
		logger.Panic("when sign oss download url:", signed, err)
	}

	signed.URL = url

	return signed, service.StatusOK
}
