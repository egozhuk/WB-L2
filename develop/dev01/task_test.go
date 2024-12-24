package main

import (
	"github.com/beevik/ntp"
	"testing"
	"time"
)

func TestGetNTPTime(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		wantErr bool
	}{
		{"ValidHost", "0.beevik-ntp.pool.ntp.org", false},
		{"InvalidHost", "invalid-ntp-host", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ntpTime, err := ntp.Time(tt.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNTPTime() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if ntpTime.IsZero() {
					t.Error("Получено нулевое время NTP")
				}

				// Проверяем, что время NTP не сильно отличается от системного времени
				// Допускаем разницу в несколько секунд из-за задержек сети и обработки
				systemTime := time.Now()
				if diff := systemTime.Sub(ntpTime); diff > time.Minute || diff < -time.Minute {
					t.Errorf("Время NTP сильно отличается от системного времени: разница %v", diff)
				}
			}
		})
	}
}
