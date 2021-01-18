package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var client *redis.Client
var wordCount, charCount, wpm int
var msg string

type Summary struct {
	Words       int
	Chars       int
	WordsPerMin int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

//listening for incoming connections1
/*func reader(conn *websocket.Conn) {
for {

	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(messageType)
	log.Println(string(p))
	msg = string(p)
	charCount = len(msg)
	wordCount = WordCount(string(p))
	if charCount >= 5 {
		wpm = (charCount) / 5
	}
	saveData()
	//summ()
	//summary := Summary{Words: wordCount, Chars: charCount, WordsPerMin: wpm}
	//msg := string(summary)
	summary := Summary{Words: wordCount, Chars: charCount, WordsPerMin: wpm}
	if err := websocket.WriteJSON(conn, summary); err != nil {
		log.Println(err)
		return
	}
	 calc()


//echo back msg to client
}
/*if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		calc()
	}
}*/

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))
		log.Println(messageType)
		log.Println(string(p))
		msg = string(p)
		charCount = len(msg)
		wordCount = WordCount(string(p))
		if charCount >= 5 {
			wpm = (charCount) / 5
		}
		saveData()
		summary := Summary{Words: wordCount, Chars: charCount, WordsPerMin: wpm}
		if err := websocket.WriteJSON(conn, summary); err != nil {
			log.Println(err)
			return

			calc()
		}
	}
}
func calc() {
	fmt.Println("msg: ", msg)
	fmt.Println("charcount:", charCount)
	fmt.Println("wordcount:", wordCount)
	fmt.Println("wpmCount:", wpm)
	//summ()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Websocket Endpoint")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	//upgrading this connection to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully Connected..")
	Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

}

//WordCount 1
func WordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}

func saveData() {
	client = redis.NewClient(&redis.Options{
		//telling client where redis server is
		Addr: "localhost:6379",
	})
	//client.LPush("rollno")
	client.LPush("Words", wordCount)
	client.LPush("Characters", charCount)
	client.LPush("Wpm", wpm)

}

func main() {
	fmt.Println("Go Websockets")
	setupRoutes()
	calc()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
