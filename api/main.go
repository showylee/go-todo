package main

import (
  "database/sql"
  //"gopkg.in/group.v1"
  "log"
  //"strconv"

  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

//var db = initDb()
//
//func initDb() *DB {
//  db, err := sql.Open("mysql", "root/test")
//  checkErr(err, "sql.Open failed")
//  //defer db.Close()
//
//  return db
//}
//
//func checkErr(err error, msg string){
//  if err != nil{
//    log.Fatalln(msg, err)
//  }
//}

func Cors() gin.HandlerFunc {
  return func(c *gin.Context){
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
    c.Next()
  }
}

type (PostFormat struct {
  Value string `json:"key"`
})

type Todo struct{
  id int
  user_id int
  item string
  delflg bool
}

func main() {
  r := gin.Default()

  r.Use(Cors())

  v1 := r.Group("api/v1")
  {
    v1.GET("/todo", GetItems)
    v1.GET("/todo/:id", GetItem)
    v1.POST("/todo", AddItem)
    v1.PUT("/todo/:id", UpdateItem)
    v1.DELETE("/todo/:id", DeleteItem)
    v1.OPTIONS("/todo", OptionItem)
    v1.OPTIONS("/todo/:id", OptionItem)
  }
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H {
      "message": "pong",
    })
  })
  r.Run(":8888")
}

func GetItems(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  db, err := sql.Open("mysql", "root:example@tcp(db:3306)/todo")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  rows, err := db.Query("select * from items")
  if err != nil {
    panic(err.Error())
  }

  var todoList []Todo
  for rows.Next(){
    todo := Todo{}
    if err := rows.Scan(&todo.id, &todo.user_id, &todo.item, &todo.delflg); err != nil {
      log.Fatal(err)
    }
    log.Printf("item: %s", todo.item);
    todoList = append(todoList, todo)
  }
  log.Printf("items item: %s", todoList[0].item)

  todo := Todo{}
  todo.id = 0
  todo.user_id = 0
  todo.item = "test test"
  todo.delflg = false

  c.JSON(200, todo)
}

func GetItem(c *gin.Context) {
  c.JSON(200, gin.H{"message": "item"})
}

func AddItem(c *gin.Context) { c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  var val PostFormat
  c.Bind(&val)

  log.Printf("post: %s", val.Value)

  //c.Request.ParseForm()
  //item := c.Request.Form["item"]

  c.JSON(200, val.Value)

  //db, err := sql.Open("mysql", "root:example@tcp(db:3306)/todo")
  //if err != nil {
  //  panic(err.Error())
  //}
  //defer db.Close()

  //ins, err := db.Prepare("insert into items(user_id, item, delflg) values(?, ?, ?)")
  //if err != nil {
  //  panic(err.Error())
  //}
  //ins.Exec(0, item, false)

  //rows, err := db.Query("select item from items")
  //if err != nil {
  //  panic(err.Error())
  //}

  //c.JSON(200, rows)
}

func UpdateItem(c *gin.Context) {
  c.JSON(200, gin.H{"message": "item"})
}

func DeleteItem(c *gin.Context) {
  c.JSON(200, gin.H{"message": "item"})
}

func OptionItem(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  c.Next()
}
