package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	fmt.Println("Запуск HTTP сервера!")

	err := http_server.StartHTTPServer()
	if err != nil {
		fmt.Println("Во время работы сервера произошла ошибка:", err)
	} else {
		fmt.Println("Сервер завершился успешно!")
	}
}
