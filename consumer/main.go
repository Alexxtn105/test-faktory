package main

import (
	"context"
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
	"log"
	"time"
)

func main() {
	mgr := worker.NewManager()

	// регистрируем задачи
	mgr.Register("email", sendEmail)
	mgr.Register("report", prepareReport)

	// устанавливаем количество горутин для исполнения задач
	mgr.Concurrency = 5

	// извлекаем задачи в соответствии с заданным приоритетом
	mgr.ProcessStrictPriorityQueues("critical", "default", "low_priority")

	// как альтернатива - можно использовать вес приоритета
	//mgr.ProcessWeightedPriorityQueues(map[string]int{"critical": 3, "default": 2, "low_priority": 1})

	// Запускаем менеджер
	err := mgr.Run()
	if err != nil {
		panic(err)
	}
}

func sendEmail(ctx context.Context, args ...any) error {
	help := worker.HelperFor(ctx)
	log.Printf("Working on job with ID: %s\n", help.Jid())

	addr := args[0].(string)
	subject := args[1].(string)

	fmt.Printf("Sending mail to %s with subject %s", addr, subject)
	time.Sleep(5 * time.Second)

	return nil
}

func prepareReport(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)
	log.Printf("Working on job with ID: %s\n", help.Jid())
	addr := args[0].(string)

	fmt.Println("Preparing report for the user: " + addr)
	time.Sleep(10 * time.Second)

	// также направляем письмо пользователю как задачу - добавляем ее в очередь
	return help.With(func(cl *faktory.Client) error {
		job := faktory.NewJob("email", addr, "Report is ready!")
		return cl.Push(job)
	})
}
