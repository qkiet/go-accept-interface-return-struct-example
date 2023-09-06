package main

import (
	"example/accept-interface-return-struct/redis"
	"example/accept-interface-return-struct/service"
	"fmt"
)

func main() {
	db := redis.NewRedisClient()
	svc := service.NewService(db)
	v, err := svc.GetAge()
	if err != nil {
		fmt.Println("Alert!! cannot get age from redit db. Try to set..")
		err = svc.SetAge(5)
		if err != nil {
			fmt.Println("Alert!! cannot get age from redit db and cannot set age to redis db")
			return
		}
	}
	fmt.Println("Get age value is ", v)
}
