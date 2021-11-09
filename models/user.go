package models
type User struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email_id" json:"email_id"`
	City  string `db:"city" json:"city"`
	MobileNo string `db:"mobile_no" json:"mobile_no"`
	IsAdmin  bool  `db:"is_admin" json:"is_admin"`
}
type ToDoTask struct{
	Email  string    `db:"email_id" json:"email_id"`
	TaskTitle  string    `db:"task_title" json:"task_title"`
	TaskDescription  string    `db:"task_description" json:"task_description"`
	TaskTime string  `db:"task_time" json:"task_time"`

}
