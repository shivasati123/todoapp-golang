package helper

import (
	"fmt"
	"todoapp-golang/database"
	"todoapp-golang/models"
)

func CreateUser(name string, email string , city string, mobNo string, isAdmin bool) (map[string]bool,error) {
	// language=SQL
	SQL := `INSERT INTO user_profile(name, email_id,mobile_no,city,is_admin ) VALUES ($1, $2,$3,$4,$5) RETURNING id;`
	var userID string
	err := database.ToDoApp.Get(&userID, SQL, name, email,mobNo,city, isAdmin)
	res :=make(map[string] bool)
	if err !=nil{
		fmt.Println(err)
		res["error"]=true
		return res,err
  }
	res["success"]=true
	return res,nil



}
func GetUserId(email string)string{
	// language=SQL
	SQL := `SELECT id from user_profile where email_id =$1;`
	var userID string
	err := database.ToDoApp.Get(&userID, SQL, email)
	if err != nil {

	}
	return userID
}
func AddUserTask(ID, taskDescription ,taskTitle, taskTime  string)(map[string]bool,error){
	// language=SQL
	fmt.Println(ID,taskDescription ,taskTitle, taskTime )
	SQL := `INSERT INTO to_do_task (user_id,task_title,task_description,task_time) VALUES ($1, $2,$3,$4) RETURNING user_id;`
	var userID string
	err := database.ToDoApp.Get(&userID, SQL,ID,taskTitle,taskDescription,taskTime)
	res :=make(map[string] bool)
	if err !=nil{
		fmt.Println(err)
		res["error"]=true
		return res,err
	}
	res["success"]=true
	return res,nil

}
func DeleteUserTask(Id,task string)(map[string]bool,error){
	// language=SQL
	SQL := `DELETE  from to_do_task where user_id =$1 and task_title = $2 RETURNING user_id;`

	_,err := database.ToDoApp.Exec( SQL,Id,task)
	res :=make(map[string] bool)
	if err !=nil{
		fmt.Println(err)
		res["error"]=true
		return res,err
	}
	res["success"]=true
	return res,nil


}
func GetAllTask()([]models.ToDoTask,error){
	SQL := `SELECT  task_title,task_description,task_time from to_do_task`
	users := make([]models.ToDoTask,0)
	err := database.ToDoApp.Select(&users,SQL)
	if err!= nil{
		return nil,err
	}
	return users,nil

}
