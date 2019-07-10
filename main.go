package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SESSION ensure global mongodb connection
var SESSION *mgo.Session

func init() {
	// mongodb manual connection using host ip. Replace your host IP address there
	session, err := mgo.Dial("172.17.0.2")
	// session, err := mgo.Dial("<HostIP>")
	Must(err)
	fmt.Println(err)
	SESSION = session
}

func main() {

	port := os.Getenv("PORT")
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(mapMongo)

	if port == "" {
		port = "8000"
	}
	r.POST("/api/v1/task", CreateTask)

	http.ListenAndServe(":"+port, r)

}

// close connection
func mapMongo(c *gin.Context) {
	s := SESSION.Clone()
	defer s.Close()
	c.Set("mongo", s.DB("mongotask"))
	c.Next()
}

// Must to catch the mongo panic issues
func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// NewTask Struct/model
type NewTask struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Task string
}

// Mongo bson generate New unique Id each request
func (self *NewTask) Init() {
	self.Id = bson.NewObjectId()
}

const (
	// CollectionTask is the collection name
	CollectionTask = "taskCollection"
)

// CreateTask to create new Task message
func CreateTask(c *gin.Context) {
	var newTask NewTask
	err := c.BindJSON(&newTask)
	if err != nil {
		c.Error(err)
		return
	}
	mongodb := c.MustGet("mongo").(*mgo.Database)
	con := mongodb.C(CollectionTask)
	// fmt.Println(newTask)
	con.Insert(newTask)
	if err != nil {
		c.Error(err)
	}

}
