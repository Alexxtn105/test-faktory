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
	for i := 1; i <= 10; i++ {

		job := faktory.NewJob("report", fmt.Sprintf("test%d@mail.ru", i))
		job.Queue = "critical"
		err = client.Push(job)
		if err != nil {
			fmt.Println("Ошибка постановки задачи в очередь")
		}
	}

}
