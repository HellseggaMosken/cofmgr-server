package service

type Status = int

const (
	StatusOK                     Status = 200
	StatusInternalError                 = 500
	StatusBadParam                      = 400
	StatusWrongAccount                  = 401 // wrong password or wrong account
	StatusNoAuth                        = 403 // no token or token invalid
	StatusNotFoundItem                  = 404
	StatusEmailExists                   = 406
	StatusOSSUnsupportedFileType        = 410
)
