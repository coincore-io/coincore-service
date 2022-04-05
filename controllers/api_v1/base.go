package api_v1


import "github.com/astaxie/beego"

const HttpAuthKey = "Authorization"


type baseController struct {
	beego.Controller
}

type RetJson struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func RetResource(status bool, code int, data interface{}, msg string) (apijson *RetJson) {
	apijson = &RetJson{Status: status, Code: code, Data: data, Msg: msg}
	return
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}