package main

import (
	"fmt"
	faktory "github.com/contribsys/faktory/client"
)

func main() {
	client, err := faktory.Open()
	if err != nil {
		panic(err)
	}

	job := faktory.NewJob("report", "test@mail.ru")
	job.Queue = "critical"
	err = client.Push(job)
	if err != nil {
		fmt.Println("Ошибка постановки задачи в очередь")
	}
}
