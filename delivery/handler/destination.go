package handler

import "net/http"

func (ctrl *ctrl) GetDestination(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctrl.uc.GetDestination(ctx).Default(w)
}
