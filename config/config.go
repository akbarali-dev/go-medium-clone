package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Port     string
	Postgres postgres
}

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(_ string) config {
	// godotenv.Load(path + "/.env")

	// conf := viper.New()
	// conf.AutomaticEnv()

	rootDir := findProjectRoot()
	envPath := filepath.Join(rootDir, ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("❌ .env file not found at %s: %v", envPath, err)
	}

	// 2. Env o‘zgaruvchilarni yuklash
	v := viper.New()
	v.AutomaticEnv()

	cfg := config{
		Postgres: postgres{
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetString("POSTGRES_PORT"),
			User:     v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			Database: v.GetString("POSTGRES_DATABASE"),
		},
		Port: v.GetString("PORT"),
	}

	log.Printf("📦 Loaded config: %+v", cfg.Postgres)
	return cfg
}

// 🔎 Bu funksiya hozirgi fayldan root katalogga ko‘tariladi
func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("❌ Could not get working directory")
	}
	for {
		if fileExists(filepath.Join(dir, ".env")) {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	log.Fatal("❌ .env file not found in any parent directory")
	return ""
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// func tryLoadEnv(path string) {
// 	err := godotenv.Load(path)
// 	if err == nil {
// 		log.Printf("✅ Loaded environment from %s", path)
// 		return
// 	}
// }
