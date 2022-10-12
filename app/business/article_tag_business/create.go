package article_tag_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) CreateArticleTag(option request.CreateArticleTagOption) (goErr errorcode.Error) {
	count, goErr := b.tagRepo.GetTagsCountByTagIDsAndUserID(option.TagIDs, option.UserID)
	if goErr != nil {
		return
	}

	if len(option.TagIDs) != int(count) {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.TagsCountNotMatchError, "")
		return
	}

	insertFields := []map[string]interface{}{}
	for _, tagID := range option.TagIDs {
		insertFields = append(insertFields, map[string]interface{}{
			"article_id": option.ArticleID,
			"tag_id":     tagID,
		})
	}

	if goErr = b.articleTagRepo.CreateArticleTagByMap(insertFields); goErr != nil {
		return
	}

	return
}
