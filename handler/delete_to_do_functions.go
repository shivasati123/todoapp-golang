package handler

import (
	"net/http"
	"todoapp-golang/database/helper"
	"todoapp-golang/utils"
)

func DeleteTask(writer http.ResponseWriter, request *http.Request){
	body := struct{
		Tasktitle       string    ` json:"task_title"`
		Email      string    ` json:"email_id"`

	}{}
	err := utils.ParseBody(request.Body,&body)
	if err!=nil{

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	email := body.Email
	userId := helper.GetUserId(email)
	res,err:=helper.DeleteUserTask(userId,body.Tasktitle)
	if err !=nil{

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(writer,http.StatusOK,res)

}
