package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Databases struct {
	LogEn bool
	DB    *gorm.DB
	err   error
	Con   ConnectionDB
}
type ConnectionDB struct {
	User     string
	Pass     string
	Protocol string
	Host     string
	Port     string
}

type TblData struct {
	gorm.Model
	Name   string
	Value  float64
	Satuan string
}

func (d *Databases) Begin(dbname string) *Databases {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  false,         // Disable color
	// 	},
	// )
	// dsn := d.Con.User + ":" + d.Con.Pass + "@" + d.Con.Protocol + "(" + d.Con.Host + ":" + d.Con.Port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	// d.DB, d.err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })
	// if d.err != nil {
	// 	d.Log("error database not connected...")
	// } else {
	// 	d.Log("success connect db ", d.Con.User, " ", d.Con.Host)
	// }
	// return d
	// // db.AutoMigrate(&Status{}, &Command{})
	return nil
}

func (d *Databases) ConSqlLite() *Databases {
	// github.com/mattn/go-sqlite3
	d.DB, d.err = gorm.Open(sqlite.Open("car.db"), &gorm.Config{})
	return d
}

func (d *Databases) Log(msg ...interface{}) {
	if d.LogEn {
		log.Println(msg...)
	}
}

//=============

func (d *Databases) Init() {
	d.DB.AutoMigrate(&TblData{})

	// d.DB.Create(&TblData{
	// 	Name:   "Bensin",
	// 	Value:  12.5,
	// 	Satuan: "Liter",
	// })
}

func (d *Databases) Insert(data *TblData) {
	d.DB.Create(data)
}

func (d *Databases) List() (data []TblData) {
	d.DB.Find(&data)
	return data
}
