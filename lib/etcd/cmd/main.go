package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://localhost:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go func() {
		ticker := time.NewTicker(time.Second)
		for ; true; <-ticker.C {
			if _, err := cli.Put(context.Background(), "final_reminder_count", ""); err != nil {
				log.Println(err)
			} else {
				log.Println("ok")
			}
		}
	}()

	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		_, err := cli.Put(context.Background(), "reminders", fmt.Sprintf("[{\"name\":\"%d\"}]", i))
	//		if err != nil {
	//			log.Println("cannot put value")
	//		}
	//		time.Sleep(2 * time.Second)
	//	}
	//}()
	//
	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		_, err := cli.Put(context.Background(), "remind_interval", fmt.Sprint(i*100))
	//		if err != nil {
	//			log.Println("cannot put value")
	//		}
	//		time.Sleep(2 * time.Second)
	//	}
	//}()
	//
	//go func() {
	//	rch := cli.Watch(context.Background(), "remind_interval")
	//	for wresp := range rch {
	//		for _, ev := range wresp.Events {
	//			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//		}
	//	}
	//}()

	var ch chan int
	<-ch
}
