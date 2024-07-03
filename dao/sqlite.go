package dao

import (
	"errors"
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

var sqlInitStr = "CREATE TABLE account\n(\n    id             INTEGER PRIMARY KEY AUTOINCREMENT,\n    username       TEXT    NOT NULL UNIQUE DEFAULT '',\n    pass           TEXT    NOT NULL        DEFAULT '',\n    con_pass       TEXT    NOT NULL        DEFAULT '',\n    quota          INTEGER NOT NULL        DEFAULT 0,\n    download       INTEGER NOT NULL        DEFAULT 0,\n    upload         INTEGER NOT NULL        DEFAULT 0,\n    expire_time    INTEGER NOT NULL        DEFAULT 0,\n    kick_util_time INTEGER NOT NULL        DEFAULT 0,\n    device_no      INTEGER NOT NULL        DEFAULT 3,\n    role           TEXT    NOT NULL        DEFAULT 'user',\n    deleted        INTEGER NOT NULL        DEFAULT 0,\n    create_time    TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,\n    update_time    TIMESTAMP               DEFAULT CURRENT_TIMESTAMP\n);\nCREATE INDEX account_deleted_index ON account (deleted);\nCREATE INDEX account_username_index ON account (username);\nCREATE INDEX account_con_pass_index ON account (con_pass);\nCREATE INDEX account_pass_index ON account (pass);\nINSERT INTO account (username, pass, con_pass, quota, download, upload, expire_time, device_no, role)\nVALUES ('sysadmin', '02f382b76ca1ab7aa06ab03345c7712fd5b971fb0c0f2aef98bac9cd',\n        'sysadmin.sysadmin', -1, 0, 0, 253370736000000, 6, 'admin');\nCREATE TABLE config\n(\n    id          INTEGER PRIMARY KEY AUTOINCREMENT,\n    key         TEXT NOT NULL UNIQUE DEFAULT '',\n    value       TEXT NOT NULL        DEFAULT '',\n    remark      TEXT NOT NULL        DEFAULT '',\n    create_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,\n    update_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP\n);\nCREATE INDEX config_key_index ON config (key);\nINSERT INTO config (key, value, remark)\nVALUES ('H_UI_WEB_PORT', '8081', 'H UI Web Port');\nINSERT INTO config (key, value, remark)\nVALUES ('H_UI_CRT_PATH', '', 'H UI Crt File Path');\nINSERT INTO config (key, value, remark)\nVALUES ('H_UI_KEY_PATH', '', 'H UI Key File Path');\nINSERT INTO config (key, value, remark)\nVALUES ('JWT_SECRET', hex(randomblob(10)), 'JWT Secret');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_ENABLE', '0', 'Hysteria2 Switch');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_CONFIG', '', 'Hysteria2 Config');\nINSERT INTO config (key, value, remark)\nVALUES ('HYSTERIA2_TRAFFIC_TIME', '1', 'Hysteria2 Traffic Time');"

var sqliteDB *gorm.DB

func InitSqliteDB(port string) error {
	var err error
	sqliteDB, err = gorm.Open(sqlite.Open(constant.SqliteDBPath), &gorm.Config{
		TranslateError: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logrus.Errorf("sqlite open err: %v", err)
		return errors.New("sqlite open err")
	}

	var count uint
	if err = sqliteDB.Raw("SELECT count(1) FROM sqlite_master WHERE type='table' AND (name = 'account' or name = 'config')").Scan(&count).Error; err != nil {
		logrus.Errorf("sqlite query table err: %v", err)
		return errors.New("sqlite query table err")
	}
	if count == 0 {
		if err = sqliteInit(sqlInitStr); err != nil {
			return err
		}
	}
	if port != "" {
		if tx := sqliteDB.Exec("UPDATE config set `value` = ? where `key` = 'H_UI_WEB_PORT'", port); tx.Error != nil {
			logrus.Errorf("sqlite exec err: %v", tx.Error)
			return errors.New("sqlite exec err")
		}
	}
	return nil
}

func sqliteInit(sqlStr string) error {
	if sqliteDB != nil {
		sqls := strings.Split(strings.Replace(sqlStr, "\r\n", "\n", -1), ";\n")
		for _, s := range sqls {
			s = strings.TrimSpace(s)
			if s != "" {
				tx := sqliteDB.Exec(s)
				if tx.Error != nil {
					logrus.Errorf("sqlite exec err: %v", tx.Error)
					return errors.New("sqlite exec err")
				}
			}
		}
	}
	return nil
}

func CloseSqliteDB() error {
	if sqliteDB != nil {
		db, err := sqliteDB.DB()
		if err != nil {
			logrus.Errorf("sqlite err: %v", err)
			return errors.New("sqlite err")
		}
		if err = db.Close(); err != nil {
			logrus.Errorf("sqlite close err: %v", err)
			return errors.New("sqlite close err")
		}
	}
	return nil
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
