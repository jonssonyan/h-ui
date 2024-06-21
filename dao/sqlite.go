package dao

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
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

var sqlInitStr = "create table account\n(\n    id             INTEGER                  not null\n        primary key autoincrement,\n    username       TEXT      default ''     not null\n        unique,\n    `pass`         TEXT      default ''     not null,\n    con_pass       TEXT      default ''     not null,\n    quota          INTEGER   default 0      not null,\n    download       INTEGER   default 0      not null,\n    upload         INTEGER   default 0      not null,\n    expire_time    INTEGER   default 0      not null,\n    kick_util_time INTEGER   default 0      not null,\n    device_no      INTEGER   default 3      not null,\n    `role`         TEXT      default 'user' not null,\n    deleted        INTEGER   default 0      not null,\n    create_time    TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime')),\n    update_time    TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime'))\n);\ncreate index account_deleted_index\n    on account (deleted);\ncreate index account_con_pass_index\n    on account (con_pass);\ncreate index account_pass_index\n    on account (`pass`);\ncreate index account_username_index\n    on account (username);\nINSERT INTO account (username, `pass`, con_pass, quota, download, upload, expire_time, device_no, role)\nVALUES ('sysadmin', '02f382b76ca1ab7aa06ab03345c7712fd5b971fb0c0f2aef98bac9cd',\n        'ee85cbdbb387a61de3d1d52e6773cc31e28c7913b32f4b7aa44ec61b', -1, 0, 0, 253370736000000, 6, 'admin');\ncreate table config\n(\n    id          INTEGER              not null\n        primary key autoincrement,\n    `key`       TEXT      default '' not null\n        unique,\n    `value`     TEXT      default '' not null,\n    remark      TEXT      default '' not null,\n    create_time TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime')),\n    update_time TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime'))\n);\ncreate index account_key_index\n    on config (`key`);\nINSERT INTO config (key, value, remark)\nVALUES ('H_UI_WEB_PORT', '9090', 'H UI Web 端口');\nINSERT INTO config (key, value, remark)\nVALUES ('JWT_SECRET', hex(randomblob(10)), 'JWT 密钥');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_ENABLE', '0', 'Hysteria2 开关');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_CONFIG', '', 'Hysteria2 配置');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_TRAFFIC_TIME', '1', 'Hysteria2 流量倍数');"

func InitSqliteDB(port string) {
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
		panic(fmt.Sprintf("sqlite open err: %v", err))
	}

	if _, err = os.Stat(constant.SqliteDBPath); os.IsNotExist(err) {
		panic("sqlite database file not found")
	}

	var count uint
	if err = sqliteDB.Raw("SELECT count(1) FROM sqlite_master WHERE type='table' AND (name = 'account' or name = 'config')").Scan(&count).Error; err != nil {
		panic(fmt.Sprintf("sqlite query table err: %v", err))
	}
	if count == 0 {
		if err = sqliteInit(sqlInitStr); err != nil {
			panic(fmt.Sprintf("sqlite table init err: %v", err))
		}
	}
	if port != "" {
		if tx := sqliteDB.Exec("UPDATE config set `value` = ? where `key` = 'H_UI_WEB_PORT'", port); tx.Error != nil {
			panic(fmt.Sprintf("sqlite exec err: %v", tx.Error.Error()))
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
					panic(fmt.Sprintf("sqlite exec err: %v", tx.Error.Error()))
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
			return
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
