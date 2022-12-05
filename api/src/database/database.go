package database

import (
	"disperinaker-api/config"
	"disperinaker-api/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("error opening connection : ", err)
	}
	/*
		if DB.Migrator().HasTable(&model.CalonSertificate{}) == false {
			if DB.Migrator().HasTable("sertificates") {
				err = DB.Migrator().RenameTable("sertificates", &model.CalonSertificate{})
			}
		}
	*/
	err = DB.AutoMigrate(
		&model.CalonPesertaSertifikasi{},
		&model.PesertaSertifikasiSetelahPelatihan{},

		&model.PesertaTestMinatBakat{},
		&model.PesertaPelatihan{},
	)
	if err != nil {
		fmt.Print("error migrating table : ", err)
	} else {
		fmt.Println("BERHASILLLL")
	}

	return DB
}
