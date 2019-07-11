package external

import (
	"log"

	"github.com/ShotaKitazawa/tabemap-api/utils"
	"github.com/comail/colog"
)

type Logger struct{}

func (logger Logger) Debug(args ...interface{}) {
	log.Printf("debug: %v", args...)
}
func (logger Logger) Info(args ...interface{}) {
	log.Printf("info: %v", args...)
}
func (logger Logger) Warn(args ...interface{}) {
	log.Printf("warn: %v", args...)
}
func (logger Logger) Error(args ...interface{}) {
	log.Printf("error: %v", args...)
}

func init() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	switch utils.GetEnvOrDefault("LOG_MIN_LEVEL", "info") {
	case "debug":
		log.Println("info: Set Minimum LogLevel: Debug")
		colog.SetMinLevel(colog.LDebug)
	case "info":
		log.Println("info: Set Minimum LogLevel: Info")
		colog.SetMinLevel(colog.LInfo)
	case "warn":
		log.Println("info: Set Minimum LogLevel: Warn")
		colog.SetMinLevel(colog.LWarning)
	case "error":
		log.Println("info: Set Minimum LogLevel: Error")
		colog.SetMinLevel(colog.LError)
	default:
		log.Println("info: Set Minimum LogLevel: Info")
		colog.SetMinLevel(colog.LInfo)
	}
	colog.Register()
}
