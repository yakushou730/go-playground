package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	globalID int
	idLocker sync.Mutex

	// 新用戶到來，進行登記
	enteringChannel = make(chan *User)
	// 用戶離開，進行登記
	leavingChannel = make(chan *User)
	// 廣播專用的用戶普通消息，緩衝是盡可能避免出現異常情況堵塞
	messageChannel = make(chan Message, 8)
)

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (u *User) String() string {
	return u.Addr + ", UID:" + strconv.Itoa(u.ID) + ", Enter at:" + u.EnterAt.Format("2006-01-02 15:04:05+8000")
}

type Message struct {
	OwnerID int
	Content string
}

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster 用於記錄聊天室使用者，並進行訊息廣播
// 1. 新使用者進來
// 2. 使用者普通訊息
// 3. 使用者離開
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users, user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			for user := range users {
				if user.ID == msg.OwnerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1. 新使用者進來，建制該使用者的實例
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 2. 用於寫入操作的 goroutine
	go sendMessage(conn, user.MessageChannel)

	// 3. 給目前使用者發送歡迎訊息，向所有使用者告知新使用者到來
	user.MessageChannel <- "Welcome, " + user.String()
	msg := Message{
		OwnerID: user.ID,
		Content: "user:`" + strconv.Itoa(user.ID) + "` has enter",
	}
	messageChannel <- msg

	// 4. 記錄到全域使用者清單中，避免用鎖
	enteringChannel <- user

	// 控制超時用戶彈出
	var userActive = make(chan struct{})
	go func() {
		d := 1 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	// 5. 循環讀取使用者輸入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg.Content = strconv.Itoa(user.ID) + ":" + input.Text()
		messageChannel <- msg

		// 用戶活躍
		userActive <- struct{}{}
	}
	if err := input.Err(); err != nil {
		log.Println("讀取錯誤:", err)
	}

	// 6. 使用者離開
	leavingChannel <- user
	msg.Content = "user:`" + strconv.Itoa(user.ID) + "`has left"
	messageChannel <- msg
}

func GenUserID() int {
	idLocker.Lock()
	defer idLocker.Unlock()

	globalID++
	return globalID
}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
