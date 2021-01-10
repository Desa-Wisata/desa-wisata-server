package handler

import (
	"net/http"

	"github.com/desa-wisata/desa-wisata-server/usecase"
)

type ctrl struct {
	uc usecase.Base
}

//Base ...
type Base interface {
	GetDestination(w http.ResponseWriter, r *http.Request)
}

//InitController ...
func InitController(c usecase.Base) Base {
	return &ctrl{c}
}
