package handler

import (
	"fmt"
	"net/http"
	"todoapp-golang/database/helper"
	"todoapp-golang/utils"
)

func GetTask(writer http.ResponseWriter, request *http.Request){
	fmt.Println("HIIIIIIIIIIIIIIIIIIIIIII")
	res,err:= helper.GetAllTask()
	if err != nil{
		fmt.Println("hiiiiiiii")
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println("hiii")
		return

	}
	utils.RespondJSON(writer,http.StatusOK,res)

}