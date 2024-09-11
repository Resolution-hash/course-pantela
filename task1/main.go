package main

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

/*TODO:
1. Если еще нет, создать аккаунт на [github](https://github.com/). Создать репозиторий и загрузить туда проект
2. Добавить глобальную переменную `message`
3. Добавить `POST` handler, который будет принимать json с полем `message` и записывать его содержимое в нашу переменную.
4. Обновить `GET` handler, чтобы он возвращал “hello, `message` ”
5. Сделать коммит и загрузить изменения в созданный заранее репозиторий
6. Прислать мне ссылку на ваш репозиторий
7. Также прислать мне скриншоты кода и постмана :)
*/

type Message struct {
	Text string `json:"text"`
}

var message Message

// type Message struct{
// 	Text string `json:"text"`
// }

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, ", message.Text)

}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = json.Unmarshal(data, &message)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Println("message:", message.Text)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", MessageHandler).Methods("POST")

	fmt.Println("Server is starting on port:8080")
	http.ListenAndServe(":8080", router)
}
