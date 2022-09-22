package user_handler

import "mini-news/app/business/user_business"

type Handler struct {
	userBusiness user_business.Interface
}

func NewsHandler() *Handler {
	return &Handler{
		userBusiness: user_business.NewBusiness(),
	}
}
