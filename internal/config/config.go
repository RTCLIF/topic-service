package config

import (
	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path) // чтобы можно было из файла загрузить переменные как переменные окружения

	if err != nil {
		return err
	}
	return nil
}
