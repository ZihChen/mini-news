package tag_business

import (
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/structs/request"
)

func (b *business) CreateTag(option request.CreateTagOption) (goErr errorcode.Error) {
	insertMap, err := helper.StructToMap(option)
	if err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorBusiness, errorcode.StructToMapError, err.Error())
		return
	}

	if goErr = b.tagRepo.CreateTagByMap(insertMap); goErr != nil {
		return
	}

	return
}
