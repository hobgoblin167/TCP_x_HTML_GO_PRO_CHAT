// сервер на Go
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var strCh chan string
var messages string

func main() {
	fmt.Println("Сервер запущен")
	//go TcpConnection()
	HTMLConnection()

}
func TcpConnection() {
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		fmt.Println("Ошибка создания сервера:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка принятия подключения:", err)
			continue
		}

		go HandleClient(conn) //горутина для обработки разных клиентов tcp

	}
}
func HandleClient(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Новый клиент: %s\n", clientAddr)
	for {
		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		fmt.Println("Получено сообщение от", clientAddr, " : ", string(buffer[:n]))
		messages += fmt.Sprintf("<div class='message'>%s > %s</div>", clientAddr, string(buffer[:n]))
	}

}
func HTMLConnection() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, _ := os.ReadFile("index.html")
		htmlStr := string(html)

		if messages == "" {
			htmlStr = fmt.Sprintf(htmlStr, "Сообщений пока нет")
		} else {
			htmlStr = fmt.Sprintf(htmlStr, messages) //ищет %s в нашем стринге и подставляет туда переменную messages
		}

		w.Write([]byte(htmlStr))
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
