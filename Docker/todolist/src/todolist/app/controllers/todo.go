package controllers

import (
	"github.com/revel/revel"
	"todolist/app/models"
	"todolist/app/routes"
	"gopkg.in/mgo.v2/bson"
	"time"
);
type TodoList struct {
	App
}

func (c TodoList) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}
	return nil
}

func (c TodoList) Index() revel.Result {
	uncompletedTasks := make([]models.TodoListItem, 0, 10)
	completedTasks := make([]models.TodoListItem, 0, 10)
	userId := c.connected().UserId
	err := todolist.Find(bson.M{"userid": userId, "done": false}).Sort("-scheduled").All(&uncompletedTasks)
	if err != nil {
		panic(err)
	}
	err = todolist.Find(bson.M{"userid": userId, "done": true}).Sort("-scheduled").All(&completedTasks)
	if err != nil {
		panic(err)
	}

	return c.Render(uncompletedTasks, completedTasks)
}

func (c TodoList) Settings() revel.Result {
	return c.Render()
}
func (c TodoList) SaveSettings(password, verifyPassword string) revel.Result {
	models.ValidatePassword(c.Validation, password)
	c.Validation.Required(verifyPassword).
		Message("Please verify your password")
	c.Validation.Required(verifyPassword == password).
		Message("Your password doesn't match")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.Redirect(App.Index)
	}

	//bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//_, err := c.Txn.Exec("update User set HashedPassword = ? where UserId = ?",
	//	bcryptPassword, c.connected().UserId)
	//if err != nil {
	//	panic(err)
	//}
	c.Flash.Success("Password updated")
	return c.Redirect(routes.TodoList.Settings())
}

func (c TodoList) Add(todolistitem models.TodoListItem) revel.Result {
	c.Validation.Required(todolistitem.Value).Message("What you want to do?")
	c.Validation.Required(todolistitem.ScheduledFormatted).Message("When do you want to do it?")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.TodoList.Index())

	}
	lastItem := models.TodoListItem{}
	userId := c.connected().UserId
	layout := "2006-01-02 15:04:05"
	scheduled, _ := time.Parse(layout, todolistitem.ScheduledFormatted + ":00")
	timestamp := time.Now()
	err := todolist.Find(bson.M{"userid": userId}).Sort("-itemid").One(&lastItem)
	if err != nil {
		todolistitem.ItemId = 1
	}else{
		todolistitem.ItemId = lastItem.ItemId + 1
	}
	todolistitem.UserId = userId
	todolistitem.Scheduled = scheduled
	todolistitem.ScheduledFormatted = scheduled.Format("2006-01-02 15:04:05")
	todolistitem.Timestamp = timestamp
	todolistitem.TimestampFormatted = timestamp.Format("2006-01-02 15:04:05")
	todolist.Insert(todolistitem)

	c.Flash.Success("<strong>@TODO:</strong> 	" + todolistitem.Value + " scheduled for " + todolistitem.ScheduledFormatted)
	return c.Redirect(routes.TodoList.Index())

}
func (c TodoList) Done(ItemId int) revel.Result {
	userId := c.connected().UserId
	todolistitem := models.TodoListItem{}
	err := todolist.Find(bson.M{"itemid": ItemId, "userid": userId}).One(&todolistitem)
	if err != nil { panic(err) }

	todolistitem.MarkAsDone()
	err = todolist.Remove(bson.M{"itemid": ItemId, "userid": userId})
	if err != nil { panic(err) }
	c.Validation.Required(todolistitem.Done).Message("")

	err = todolist.Insert(todolistitem);
	if err != nil { panic(err) }

	c.Flash.Success("<strong>@DONE at " + todolistitem.CompletionFormatted  + " </strong> "+ todolistitem.Value)

	return c.Redirect(routes.TodoList.Index())

}
func (c TodoList) Delete(ItemId int) revel.Result {
	userId := c.connected().UserId
	err := todolist.Remove(bson.M{"itemid": ItemId, "userid": userId})
	if err != nil { panic(err) }

	c.Flash.Success("<strong>Item successfully deleted.</strong>")

	return c.Redirect(routes.TodoList.Index())
}