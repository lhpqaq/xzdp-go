package mysql

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"xzdp/conf"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				TablePrefix:   "tb_",
			},
		},
	)
	//全局注册创建时间
	DB.Callback().Create().Before("gorm:create").Register("update_timestamps_before_create", func(tx *gorm.DB) {
		if tx.Error == nil {
			now := time.Now().Format("2006-01-02 15:04:05")
			if tx.Statement.Schema != nil {
				if createTimeField := tx.Statement.Schema.LookUpField("create_time"); createTimeField != nil {
					createTimeField.Set(context.Background(), tx.Statement.ReflectValue, now)
				}
				if createTimeField := tx.Statement.Schema.LookUpField("update_time"); createTimeField != nil {
					createTimeField.Set(context.Background(), tx.Statement.ReflectValue, now)
				}
			}
		}
	})
	//全局注册更新时间
	DB.Callback().Update().Before("gorm:update").Register("update_timestamps_before_update", func(tx *gorm.DB) {
		if tx.Error == nil {
			now := time.Now().Format("2006-01-02 15:04:05")
			if tx.Statement.Schema != nil {
				if updateTimeField := tx.Statement.Schema.LookUpField("update_time"); updateTimeField != nil {
					_ = updateTimeField.Set(context.Background(), tx.Statement.ReflectValue, now)
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}
	//autoMigrateTable()
}

//func autoMigrateTable() {
//	DB.AutoMigrate(
//		&model.User{},
//	)
//}
