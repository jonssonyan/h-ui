package dao

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"h-ui/model/constant"
	"log"
	"os"
	"strings"
	"time"
)

var sqliteDB *gorm.DB

var sqlInitStr = "create table account\n(\n    id          INTEGER              not null\n        primary key autoincrement,\n    username    TEXT      default '' not null,\n    pass        TEXT      default '' not null,\n    quota       INTEGER   default 0  not null,\n    download    INTEGER   default 0  not null,\n    upload      INTEGER   default 0  not null,\n    expire_time INTEGER   default 0  not null,\n    deleted     INTEGER   default 0  not null,\n    create_time TIMESTAMP default CURRENT_TIMESTAMP,\n    update_time TIMESTAMP default CURRENT_TIMESTAMP\n);\ncreate index account_deleted_index on account (deleted);\ncreate index account_pass_index on account (pass);\ncreate index account_username_index on account (username);\ncreate table config\n(\n    id          INTEGER              not null\n        primary key autoincrement,\n    `key`       TEXT      default '' not null,\n    value       TEXT      default '' not null,\n    remark      TEXT      default '' not null,\n    create_time TIMESTAMP default CURRENT_TIMESTAMP,\n    update_time TIMESTAMP default CURRENT_TIMESTAMP\n);\ncreate index account_key_index on config (`key`);\nINSERT INTO config (id, `key`, value, remark, create_time, update_time)\nVALUES (1, 'WebServerPort', '8081', 'Web 服务端口', '2024-04-20 00:00:00', '2024-04-20 00:00:00');\nINSERT INTO config (id, `key`, value, remark, create_time, update_time)\nVALUES (2, 'Hysteria2Port', '443', 'Hysteria2 端口', '2024-04-20 00:00:00', '2024-04-20 00:00:00');\nINSERT INTO config (id, `key`, value, remark, create_time, update_time)\nVALUES (3, 'Hysteria2ApiPort', '3443', 'Hysteria2 Api 端口', '2024-04-20 00:00:00', '2024-04-20 00:00:00');\nINSERT INTO config (id, `key`, value, remark, create_time, update_time)\nVALUES (4, 'Hysteria2CertPath', '', 'Hysteria2 Cert 路径', '2024-04-20 00:00:00', '2024-04-20 00:00:00');\nINSERT INTO config (id, `key`, value, remark, create_time, update_time)\nVALUES (5, 'Hysteria2KeyPath', '', 'Hysteria2 Key 路径', '2024-04-20 00:00:00', '2024-04-20 00:00:00');"

func InitSqliteDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	var err error
	sqliteDB, err = gorm.Open(sqlite.Open(constant.SqliteDBPath), &gorm.Config{
		TranslateError: true,
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("sqlite connect err: %v", err))
	}

	var count uint
	if err := sqliteDB.Raw("SELECT count(1) FROM sqlite_master WHERE type='table' AND (name = 'account' or name = 'config')").Scan(&count).Error; err != nil {
		logrus.Errorf("sqlite query database err: %v", err)
		panic(err)
	}
	if count == 0 {
		if err = sqliteInit(sqlInitStr); err != nil {
			logrus.Errorf("sqlite database import err: %v", err)
			panic(err)
		}
	}
}

func sqliteInit(sqlStr string) error {
	if sqliteDB != nil {
		sqls := strings.Split(strings.Replace(sqlStr, "\r\n", "\n", -1), ";\n")
		for _, s := range sqls {
			s = strings.TrimSpace(s)
			if s != "" {
				tx := sqliteDB.Exec(s)
				if tx.Error != nil {
					logrus.Errorf("sqlite exec err: %v", tx.Error.Error())
					panic(tx.Error.Error())
				}
			}
		}
	}
	return nil
}

func CloseSqliteDB() {
	if sqliteDB != nil {
		db, err := sqliteDB.DB()
		if err != nil {
			logrus.Errorf("sqlite err: %v", err)
			return
		}
		if err = db.Close(); err != nil {
			logrus.Errorf("sqlite close err: %v", err)
		}
	}
}

func Paginate(pageNumParam *int64, pageSizeParam *int64) func(db *gorm.DB) *gorm.DB {
	var pageNum int64 = 1
	var pageSize int64 = 10
	if pageNumParam == nil || *pageNumParam == 0 {
		pageNum = *pageNumParam
	}
	if pageSizeParam == nil || *pageSizeParam == 0 {
		pageSize = *pageSizeParam
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
