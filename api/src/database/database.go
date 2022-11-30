package database

import (
	"disperinaker-api/config"
	"disperinaker-api/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		fmt.Println("error opening connection : ", err)
	}

	if DB.Migrator().HasTable(&model.CalonSertificate{}) == false {
		if DB.Migrator().HasTable("sertificates") {
			err = DB.Migrator().RenameTable("sertificates", &model.CalonSertificate{})
		}
	}

	err = DB.AutoMigrate(
		&model.CalonSertificate{},
		&model.PesertaSertificate{},

		&model.CalonTraining{},
		&model.PesertaTraining{},
	)
	if err != nil {
		fmt.Print("error migrating table : ", err)
	} else {
		fmt.Println("BERHASILLLL")
	}

	return DB
}
