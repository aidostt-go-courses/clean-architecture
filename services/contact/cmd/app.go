package main

import (
	"fmt"

	"aidostt.clean-architecture/pkg/store/postgres"
	"aidostt.clean-architecture/services/contact/internal/domain"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "!Example123",
		DbName:   "clean-architecture",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	artem := domain.NewContact("Grabelnikov", "Artem", "Sergeevich")
	aidana := domain.NewContact("Satybayeva", "Aidana", "Serikovna")
	group1 := domain.NewGroup("Group 1")

	fmt.Println(artem, aidana, group1)
}
