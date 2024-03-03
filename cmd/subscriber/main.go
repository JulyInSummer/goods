package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"goods_project/internal/config"
	"goods_project/internal/storage/clickhouse-storage"
	"goods_project/internal/types"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	ch := clickhouse_storage.NewCh(cfg)

	messages := make(chan *nats.Msg, 1000)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	batch, batchErr := ch.PrepareBatch(ctx, "INSERT INTO logs(Id, ProjectId, Name, Description, Priority, Removed) VALUES ($1, $2, $3, $4, $5, $6)")
	if batchErr != nil {
		log.Printf("PrepareBatch: %s", batchErr.Error())
		return
	}

	sub, err := nc.ChanQueueSubscribe("logs", "add", messages)
	if err != nil {
		log.Println("failed to subscribe")
		return
	}

	defer func() {
		sub.Unsubscribe()
		close(messages)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-messages:
			var chlog types.GoodResponse
			if err = json.Unmarshal(msg.Data, &chlog); err != nil {
				log.Println("failed to unmarshal")
				return
			}
			fmt.Println(string(msg.Data))
			if err = batch.AppendStruct(&chlog); err != nil {
				fmt.Printf("AppendStruct: %s", err.Error())
				return
			}

			log.Printf("Received a log from goods service: Id - %d. Batch size = %d", chlog.Id, batch.Rows())

			if batch.Rows() >= 5 {
				if err = batch.Send(); err != nil {
					log.Printf("Send: %s", err.Error())
					return
				}

				batch, batchErr = ch.PrepareBatch(ctx, "INSERT INTO logs(Id, ProjectId, Name, Description, Priority, Removed) VALUES ($1, $2, $3, $4, $5, $6)")
				if batchErr != nil {
					log.Printf("PrepareBatch: %s", batchErr.Error())
					return
				}
			}
		}
	}
}
