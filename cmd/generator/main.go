package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"telefun/config"
	"telefun/entity"
	"telefun/repository"

	"github.com/jmoiron/sqlx"
)

type provider struct {
	name   string
	prefix string
}

var providers = []provider{
	{name: "telkomsel", prefix: "+62852"},
	{name: "telkomsel", prefix: "+62811"},
	{name: "xl", prefix: "+62878"},
	{name: "tri", prefix: "+62896"},
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect(cfg.DB.Driver, cfg.DB.DSN())
	if err != nil {
		panic(err)
	}

	repo := repository.NewPhoneNumberRepository(db)

	var i, count uint64
	var numbers []entity.PhoneNumber

	i, _ = strconv.ParseUint(os.Getenv("START"), 10, 64)
	log.Println("Will start generating from ", i)

	start := time.Now()
	for ; i < 10_000_000; i++ {
		for _, provider := range providers {
			num := fmt.Sprintf("%s%07d", provider.prefix, i)
			pn, _ := entity.ParsePhoneNumber(num)

			numbers = append(numbers, pn)
			count++
		}

		if count >= 1000 {
			log.Println("Inserting to DB ...", i, numbers[0].Get(), numbers[len(numbers)-1].Get())

			err := repo.BulkCreate(context.Background(), numbers)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(59 * time.Millisecond)

			// reset
			count = 0
			numbers = []entity.PhoneNumber{}
		}
	}

	if len(numbers) > 0 {
		log.Println("Insert left over data to DB ...", numbers[0].Get(), numbers[len(numbers)].Get())

		err := repo.BulkCreate(context.Background(), numbers)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println(time.Since(start))
}
