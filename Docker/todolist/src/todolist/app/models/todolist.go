package models

import (
	"github.com/revel/revel"
	"time"
)

type TodoListItem struct {
	ItemId       int
	UserId       int

	Scheduled     		time.Time
	Completion		time.Time
	Timestamp    		time.Time
	ScheduledFormatted	string
	CompletionFormatted	string
	TimestampFormatted	string
	Value        		string
	Done			bool
}

// TODO: Make an interface for Validate() and then validation can pass in the
// key prefix ("booking.")
func (todolistitem TodoListItem) Validate(v *revel.Validation) {
	v.Required(todolistitem.Value).Message("What you want to do?")
	v.Required(todolistitem.ScheduledFormatted).Message("When do you want to do it?")
}
func (t *TodoListItem) MarkAsDone() error {
	Completion := time.Now()
	CompletionFormatted := Completion.Format("2006-01-02 15:04:05")

	t.Completion = Completion;
	t.CompletionFormatted  = CompletionFormatted
	t.Done = true
	return nil
}

//func (b TodoListItem) String() string {
//	return fmt.Sprintf("Adding(%s,%s)", b.User, TodoListItem.Value)
//}

//
//func (b *TodoListItem) PostGet() error {
//	var (
//		obj interface{}
//		err error
//	)
//
//	err = todolist.Get(User{}, b.UserId)
//	if err != nil {
//		return fmt.Errorf("Error loading a booking's user (%d): %s", b.UserId, err)
//	}
//	b.User = obj.(*User)
//
//	obj, err = exe.Get(TodoListItem{}, TodoListItem.ItemId)
//	if err != nil {
//		return fmt.Errorf("Error loading a booking's hotel (%d): %s", TodoListItem.ItemId, err)
//	}
//	return nil
//}