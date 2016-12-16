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

func (todolistitem TodoListItem) Validate(v *revel.Validation) {
	v.Required(todolistitem.Value).Message("<strong>What</strong> you want to do?")
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