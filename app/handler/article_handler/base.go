package article_handler

import "mini-news/app/business/article_tag_business"

type Handler struct {
	articleTagBusiness article_tag_business.Interface
}

func NewHandler() *Handler {
	return &Handler{
		articleTagBusiness: article_tag_business.NewBusiness(),
	}
}
