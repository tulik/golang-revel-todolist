package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"todolist/app/models"
	"gopkg.in/mgo.v2"
	r "github.com/revel/revel"
	"time"
)
var(
	err error
	session *mgo.Session
	database *mgo.Database
	users *mgo.Collection
	todolist *mgo.Collection

)
type GorpController struct {
	*r.Controller
}
func init(){
	session, err = mgo.Dial("127.0.0.1")
	database = session.DB("test");
	users = database.C("Users");
	todolist = database.C("TotoList")

}

func InitDB() {
	database.DropDatabase();
	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("demo"), bcrypt.DefaultCost)
	demoUser := &models.User{0, "Demo User", "demo", "demo", bcryptPassword}
	err := users.Insert(demoUser)
	if err != nil {
		print(err)
	}
	Completion := time.Now();
	CompletionFormatted := Completion.Format("2006-01-02 15:04:05")
	Timestamp := time.Now();
	TimestampFormatted := Timestamp.Format("2006-01-02 15:04:05")

	ScheduledOne := time.Date(2017, time.February, 10, 18, 0, 0, 0, time.UTC)
	ScheduledOneFormatted := ScheduledOne.Format("2006-01-02 15:04:05")
	ValueOne := "Go application live!"
	taskOne := models.TodoListItem{1, 0,
		ScheduledOne,
		Completion,
		Timestamp,
		ScheduledOneFormatted,// TODO: Make an interface for Validate() and then validation can pass in the
// key prefix ("booking.")

		CompletionFormatted,
		TimestampFormatted,
		ValueOne,
		false,
	}

	ScheduledTwo := time.Date(2017, time.February, 10, 28, 0, 0, 0, time.UTC)
	ScheduledTwoFormatted := ScheduledTwo.Format("2006-01-02 15:04:05")
	ValueTwo := "Make it responsive!"
	taskTwo := models.TodoListItem{2, 0,
		ScheduledTwo,
		Completion,
		Timestamp,
		ScheduledTwoFormatted,
		CompletionFormatted,
		TimestampFormatted,
		ValueTwo,
		false,
	}
	ScheduledThree := time.Date(2017, time.February, 10, 19, 10, 23, 0, time.UTC)
	ScheduledThreeFormatted := ScheduledThree.Format("2006-01-02 15:04:05")
	ValueThree := "Have fun with Node.js"
	taskThree := models.TodoListItem{3, 0,
		ScheduledTwo,
		Completion,
		Timestamp,
		ScheduledThreeFormatted,
		CompletionFormatted,
		TimestampFormatted,
		ValueThree,
		false,
	}
	err = todolist.Insert(taskOne, taskTwo, taskThree)
	if err != nil {
		print(err)
	}

}
