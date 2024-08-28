package todocrud

import "time"

var tasks = []Task{
	{
		ID:       1,
		Title:    "New Task",
		TaskText: []string{"Clean da house", "Wash dishes"},
		Date:     time.Now().Format("02-01-2006 15:04:05"),
	},
	{
		ID:       2,
		Title:    "Pedro Day",
		TaskText: []string{"Wash a Pedro", "!Make a Pedro's Day dinner!"},
		Date:     time.Now().Format("02-01-2006 15:04:05"),
	},
}
