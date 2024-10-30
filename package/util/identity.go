package util

import (
	"Mou1ght-Server/internal/controller"
	"Mou1ght-Server/internal/model"
	"math/rand"
	"time"
)

func GenerateIdentity() uint {
	rand.NewSource(time.Now().UnixNano())
	name := rand.Intn(100000)

	_, exisit := controller.CheckExistName(&model.User{}, string(rune(name)))
	for exisit {
		name = rand.Intn(100000)
		_, exisit = controller.CheckExistName(&model.User{}, string(rune(name)))
	}

	return uint(name)
}
