package logger

import (
	"database/sql"
	"encoding/json"

	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	logrus "github.com/sirupsen/logrus"
	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
)

const (
	LogCategoryDefault        = "default"
	LogCategorySystem         = "application.system"
	LogCategoryDatabase       = "application.database"
	LogCategoryGrpcConnection = "application.grpc.connection"
	LogCategoryGrpcTinkoff    = "application.grpc.tinkoff"
	LogCategoryLogic          = "application.logic"
)

func NewLogger(db database.DatabaseInterface) LoggerInterface {
	initLogrus(db.GetConnection())
	return &Logger{}
}

func initLogrus(db *gorm.DB) {
	sqlDB, _ := db.DB()
	hook := pglogrus.NewHook(sqlDB, map[string]interface{}{"category": LogCategoryDefault})
	hook.InsertFunc = func(db *sql.DB, entry *logrus.Entry) error {
		jsonData, err := json.Marshal(entry.Data)
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO logs(category, level, message, context, created_at) VALUES ($1,$2,$3,$4,$5);", entry.Data["category"], entry.Level, entry.Message, jsonData, entry.Time)
		return err
	}
	logrus.AddHook(hook)
}

type LoggerInterface interface {
	Info(category string, message string, context map[string]interface{})

	Warning(category string, message string, context map[string]interface{})

	Error(category string, message string, context map[string]interface{})

	Fatal(category string, message string, context map[string]interface{})

	Panic(category string, message string, context map[string]interface{})
}

type Logger struct{}

func (l Logger) Info(category string, message string, context map[string]interface{}) {
	context["category"] = category
	logrus.WithFields(context).Info(message)
}

func (l Logger) Warning(category string, message string, context map[string]interface{}) {
	context["category"] = category
	logrus.WithFields(context).Warning(message)
}

func (l Logger) Error(category string, message string, context map[string]interface{}) {
	context["category"] = category
	logrus.WithFields(context).Error(message)
}

func (l Logger) Fatal(category string, message string, context map[string]interface{}) {
	context["category"] = category
	logrus.WithFields(context).Fatal(message)
}

func (l Logger) Panic(category string, message string, context map[string]interface{}) {
	context["category"] = category
	logrus.WithFields(context).Panic(message)
}
