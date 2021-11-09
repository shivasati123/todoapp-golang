package handler

import (
	"fmt"
	"net/http"
	"todoapp-golang/database/helper"
	"todoapp-golang/utils"
)

func AddUser(writer http.ResponseWriter, request *http.Request){
	body := struct{
		Name       string    ` json:"name"`
		Email      string    ` json:"email_id"`
		City  string  `json:"city"`
		MobileNo string ` json:"mobile_no"`
		IsAdmin  bool  ` json:"is_admin"`
	}{}
	err := utils.ParseBody(request.Body,&body)
	if err!=nil{
		fmt.Println("hiii")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	res,err := helper.CreateUser(body.Name,body.Email,body.City,body.MobileNo,body.IsAdmin)
	if err !=nil{

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(writer,http.StatusOK,res)

}
func AddTask(writer http.ResponseWriter, request *http.Request){
	body := struct{

		Email      string    ` json:"email_id"`
		TaskTitle  string  `json:"task_title"`
		TaskDescription string `json:"task_description"`
		TaskTime string `json:"task_time"`

	}{}
	err := utils.ParseBody(request.Body,&body)
	if err!=nil{

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	email := body.Email
	userId := helper.GetUserId(email)
	res,err:=helper.AddUserTask(userId,body.TaskDescription,body.TaskTitle,body.TaskTime)
	if err !=nil{

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(writer,http.StatusOK,res)

}
