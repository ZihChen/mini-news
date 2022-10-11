package tag_handler

import "mini-news/app/business/tag_business"

type Handler struct {
	tagBusiness tag_business.Interface
}

func NewHandler() *Handler {
	return &Handler{
		tagBusiness: tag_business.NewBusiness(),
	}
}
