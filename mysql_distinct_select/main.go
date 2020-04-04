package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type demoUser struct {
	gorm.Model
	Name    string
	Version string
}

func main() {
	db := initDB()
	defer db.Close()
	err := cleanUP(db)
	if err != nil {
		log.Fatalln(err)
	}
	err = seed(db)
	vs, err := distinctSelect(db, "version")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("versions:", vs, "length:", len(vs))
}

func distinctSelect(db *gorm.DB, column string) ([]string, error) {
	cs := []string{}
	// err := db.Table("demo_users").
	// 	Select("distinct " + column).Scan(&cs).Error

	// us := []demoUser{}
	// err := db.
	// 	Select("distinct " + column).Find(&us).Error

	// log.Println("users:", us)
	// for _, u := range us {
	// 	cs = append(cs, u.Version)
	// }

	a := []struct {
		//E string `gorm:"cloumn:`+ column +`"`
		Abc     string `gorm:"cloumn:version"`
		Version string
	}{}

	err := db.Table("demo_users").
		Select("distinct " + column).Scan(&a).Error

	log.Println(a)
	for _, v := range a {
		cs = append(cs, v.Abc)

	}
	return cs, err
}

func cleanUP(db *gorm.DB) error {
	return db.Delete(&demoUser{}).Error
}

func seed(db *gorm.DB) error {
	for i := 0; i < 20; i++ {
		n := i % 5
		v := i % 9
		u := demoUser{
			Name:    "hag" + strconv.Itoa(n),
			Version: strconv.Itoa(v),
		}
		err := db.Create(&u).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456789@/mydb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&demoUser{}).Error
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(true)
	return db
}
