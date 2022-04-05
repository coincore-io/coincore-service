package controllers

import (
	"coinwallet/form_validate"
	"coinwallet/global"
	"coinwallet/global/response"
	"coinwallet/services"
	"github.com/gookit/validate"
	"log"
	"strconv"
)

type NewsController struct {
	baseController
}

func (Self *NewsController) Index() {
	var srv services.NewsService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "news/index.html"
}

func (Self *NewsController) Add() {
	Self.Layout = "public/base.html"
	Self.TplName = "news/add.html"
}


func (Self *NewsController) Create() {
	var (
		vForm form_validate.NewsForm
		srv services.NewsService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(vForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "image")
	if err != nil {
		log.Println("upload--err",err)
	}
	vForm.Image = imgPath

	insertId := srv.Create(&vForm)
	url := global.URL_BACK

	if vForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *NewsController) Edit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.NewsService

	ver := srv.GetById(id)
	if ver == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["data"] = ver
	Self.Layout = "public/base.html"
	Self.TplName = "news/edit.html"
}

func (Self *NewsController) Update(){
	var (
		vForm form_validate.NewsForm
		srv services.NewsService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	if vForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(vForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "image")
	if err != nil {
		log.Println("upload--err",err)
	}
	if len(imgPath) > 0 {
		vForm.Image = imgPath
	}
	if srv.Update(&vForm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *NewsController) Del() {
	idStr := Self.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		Self.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}
	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", Self.Ctx)
	}
	var srv services.NewsService
	if srv.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}