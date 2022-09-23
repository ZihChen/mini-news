package helper

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"mini-news/app/global/errorcode"
)

func StructToMap(myStruct interface{}) (myMap map[string]interface{}, err error) {
	myMap = make(map[string]interface{})

	byteData, err := jsoniter.Marshal(myStruct)
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(byteData, &myMap); err != nil {
		return
	}

	return
}

func ErrorHandle(errComponent errorcode.ErrComponent, errType errorcode.ResponseErrType, errMsg string) (goError errorcode.Error) {
	goError = errorcode.NewGoError()

	goError.SetErrComponent(errComponent)
	goError.SetResponseErrType(errType)
	goError.SetErrMsg(errMsg)

	return
}

func PrintLog(param interface{}) {
	byteData, _ := jsoniter.Marshal(param)
	show := string(byteData)
	fmt.Println("PrintLog:" + show)
	panic(show)
}
