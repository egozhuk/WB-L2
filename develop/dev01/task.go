package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получение точного времени
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// Обработка ошибки
		logToStderr(err)
		os.Exit(1)
	}

	// Получение локального времени
	localTime := time.Now()

	fmt.Printf("Local Time: %v\n", localTime)
	fmt.Printf("NTP Time:   %v\n", ntpTime)
}

// logToStderr выводит ошибку в STDERR
func logToStderr(err error) {
	log.SetOutput(os.Stderr)
	log.Printf("Error fetching NTP time: %v\n", err)
}
