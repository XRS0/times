package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Particpant struct {
  Host bool
  Conn *websocket.Conn
}

type RoomMap struct {
  Mutex sync.RWMutex
  Map map[string][]Particpant
}

func (r *RoomMap) Init() {
  r.Map = make(map[string][]Particpant)
}

func (r *RoomMap) Get(roomID string) []Particpant {
  r.Mutex.RLock()
  defer r.Mutex.RUnlock()


  return r.Map[roomID]
}

func (r *RoomMap) CreateRoom() string {
  r.Mutex.Lock()
  defer r.Mutex.Unlock()

  rand.Seed(time.Now().UnixNano())
  var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
  b := make([]rune, 8)

  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }

  roomID := string(b)

  r.Map[roomID] = []Particpant{}

  return roomID
}

func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
  r.Mutex.Lock()
  defer r.Mutex.Unlock()

  p := Particpant{host, conn}

  log.Println("Inserting into Room with RoomID: ", roomID)
  r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
  r.Mutex.Lock()
  defer r.Mutex.Unlock()

  delete(r.Map, roomID)
}







