package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения: ", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Начинайте вводить сообщения")
	for {
		message, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Ошибка подключения: ", err)
			return
		}
		fmt.Println("Отправлено :", message)
	}

}
