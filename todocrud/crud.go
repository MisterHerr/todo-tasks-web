package todocrud

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func printTask(t Task, w http.ResponseWriter) {
	fmt.Fprintf(w, "\t~{{-------------------------------}}~\n\nID\t%d\n\n", t.ID)
	fmt.Fprintf(w, "~~Title~~ || %v\n", t.Title)
	fmt.Fprintf(w, "~~Text of task~~\n")

	for i, val := range t.TaskText {
		fmt.Fprintf(w, "\t%d | **%v** |\n", i+1, val)
	}

	fmt.Fprintf(w, "\nDate of Add || %v\n\n", t.Date)
	fmt.Fprintf(w, "\t~{{-------------------------------}}~\n\n\n\n")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	for _, task := range tasks {
		printTask(task, w)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queryID, _ := strconv.Atoi(params["id"])

	if queryID == 0 {
		fmt.Fprintf(w, "You are just put non-digit ID\nPlease put the digits of Task-ID what you want")

		log.Println("User just tryied GET string ID")
		return
	}
	for _, task := range tasks {
		if task.ID == queryID {
			printTask(task, w)
			return
		}
	}


	w.Write([]byte("Task with this ID is not exist"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deleteID, _ := strconv.Atoi(params["id"])

	if deleteID == 0 {
		fmt.Fprintf(w, "You are just put non-digit ID\nPlease put the digits of Task-ID what you want")

		log.Println("User just tryied DELETE string ID")
		return
	}

	for indexT, task := range tasks {
		if task.ID == deleteID {
			tasks = append(tasks[:indexT], tasks[indexT+1:]...)

			log.Printf("Task with ID {%d} has been deleted", deleteID)
			fmt.Fprintf(w, "Task has been deleted! ;)")
			return
		}
	}
	fmt.Fprintf(w, "Task with this ID is not exist")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t Task

	json.NewDecoder(r.Body).Decode(&t)

	t.ID = rand.Intn(100000)
	t.Date = time.Now().Format("02-01-2006 15:04:05")
	tasks = append(tasks, t)

	printTask(t, w)
}
