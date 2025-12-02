package main

import (
	"awsc/logger"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type Function struct {
	Name string `json:"name"`
}

func invokeFuncName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	funcName := r.PathValue("funcName")
	logger.L().Info("received request", zap.String("funcName", funcName))
}

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	logger.InitLogger(env)
	defer logger.Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /invoke/{funcName}", invokeFuncName)

	logger.L().Info("Сервер запущен", zap.Int("port", 8080))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), mux); err != nil {
		logger.L().Fatal("Ошибка запуска сервера", zap.Error(err))
	}
}
