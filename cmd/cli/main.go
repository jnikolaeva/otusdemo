package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"

	"github.com/arahna/otusdemo/user/application"
	"github.com/arahna/otusdemo/user/infrastructure/postgres"
)

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: cli COMMAND [arg...]
Commands:
  load-users [-count C]
	  Loads test user data
	  Use -count option to specify the number of users.`)
	}

	flag.Parse()

	if flag.Arg(0) == "load-users" {
		args := flag.Args()[1:]
		loadUsersFlagSet := flag.NewFlagSet("load-users", flag.ExitOnError)
		usersCountPtr := loadUsersFlagSet.Int("count", 1000, "Users count")
		if err := loadUsersFlagSet.Parse(args); err != nil {
			log.Fatal(err)
		}
		if *usersCountPtr <= 0 {
			log.Fatal("error: -count flag must be greater than 0")
		}

		connConfig, err := postgres.ParseEnvConfig("")
		if err != nil {
			log.Fatal(err.Error())
		}
		connectionPool, err := postgres.NewConnectionPool(connConfig)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer connectionPool.Close()

		repository := postgres.New(connectionPool)
		service := application.NewService(repository)

		if err := loadTestUsers(service, *usersCountPtr); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		flag.Usage()
		os.Exit(2)
	}
}

func loadTestUsers(service application.Service, count int) error {
	for i := 0; i < count; i++ {
		if _, err := service.Create(faker.Username(), faker.FirstName(), faker.LastName(), faker.Email(), faker.Phonenumber()); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
