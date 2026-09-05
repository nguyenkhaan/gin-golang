package main

import (
	"cloudian/cloudian-restful/internal/config"
	"cloudian/cloudian-restful/internal/db"
	"cloudian/cloudian-restful/internal/model"
	"flag"
	"log"

	"gorm.io/gorm"
)

//This function will auto run when in the package main file
func init() {
	config.LoadEnvVariables() 
	db.ConnectDB() 
} 
func MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.UserModel{}, 
		&model.TodoModel{}, 
	)
}
func MigrateDown(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.UserModel{},  
		&model.TodoModel{}, 
	)
}	

func main() {
	//We will support flag: --up or --down to decide whether run MigrateUp() or MigrateDown()
	//Syntax (flag package): up := flag.Bool(flagName, defaultValue, description) ----> up is a pointer 
	//Provide a flagNamer into run command. If we don't set the flagName in the command, *up will become default value 
	up := flag.Bool("up", false, "Run migration up") //If we don't pass --up to commands, *up = false 
	down := flag.Bool("down", false, "Migration down") 
	flag.Parse() 
	if *up {
		if err := MigrateUp(db.DB); err != nil {
			log.Fatal("Error while running migration: ", err)
		}
	} 
	if *down {
		if err := MigrateDown(db.DB); err != nil {
			log.Fatal("Error while downing migration: ", err) 
		}
	}
}