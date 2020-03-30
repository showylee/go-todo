package main

import (
  "database/sql"
  //"gopkg.in/group.v1"
  "log"
  //"strconv"
  "os"

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

type PostFormat struct {
  Val string `json:"item"`
}

type Todo struct{
  Id int
  UserId int
  Item string
  Delflg bool
}
type TodoList []Todo

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

  logfile, err := os.Create("/go/src/api/log.txt")
  if err != nil {
    return
  }
  defer logfile.Close()

  var todoList TodoList
  for rows.Next(){
    todo := Todo{}
    if err := rows.Scan(&todo.Id, &todo.UserId, &todo.Item, &todo.Delflg); err != nil {
      log.Fatal(err)
    }
    logfile.Write(([]byte)(todo.Item))
    logfile.Write(([]byte)("\n"))
    todoList = append(todoList, todo)
  }

  c.JSON(200, todoList)
}

func GetItem(c *gin.Context) {
  c.JSON(200, gin.H{"message": "item"})
}

func AddItem(c *gin.Context) { c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  var val PostFormat
  c.BindJSON(&val)

  logfile, err := os.Create("/go/src/api/log.txt")
  if err != nil {
    return
  }
  defer logfile.Close()

  item := val.Val 

  db, err := sql.Open("mysql", "root:example@tcp(db:3306)/todo")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  ins, err := db.Prepare("insert into items(user_id, item, delflg) values(?, ?, ?)")
  if err != nil {
    panic(err.Error())
  }
  ins.Exec(0, item, false)

  rows, err := db.Query("select item from items")
  if err != nil {
    panic(err.Error())
  }

  var res []string
  for rows.Next() {
    if err := rows.Scan(&item); err != nil {
      log.Fatal(err)
    }
    logfile.Write(([]byte)(item))
    res = append(res, item)
  }

  c.JSON(200, res)
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
