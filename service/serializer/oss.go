package serializer

type OSSSignedURL struct {
	URL          string `json:"url"`
	Filename     string `json:"filename"`
	ContentType  string `json:"contentType"`
	Method       string `json:"method"`
	ExpiredInSec int64  `json:"expiredInSec"`
}
