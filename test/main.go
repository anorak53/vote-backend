package main

import (
	"fmt"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	// สร้าง registry สำหรับเก็บ metrics ทั้งหมด
	r := metrics.NewRegistry()

	// สร้าง gauge metric สำหรับเก็บเวลาปัจจุบัน
	gauge := metrics.NewGauge()
	r.Register("current_time", gauge)

	// อัพเดท gauge metric ในลูป
	go func() {
		for {
			// อัพเดท gauge ด้วยเวลาปัจจุบันในรูปแบบ Unix timestamp
			gauge.Update(time.Now().Unix())

			// รอ 1 วินาที
			time.Sleep(1 * time.Second)
		}
	}()

	// แสดงผล metrics บน console
	for {
		// ล้างหน้าจอ
		fmt.Print("\033[H\033[2J")

		// ดึงค่าเวลาปัจจุบันจาก gauge metric และแปลงเป็นรูปแบบเวลาที่อ่านได้
		currentTime := time.Unix(gauge.Value(), 0).Format("2006-01-02 15:04:05")
		fmt.Println("Current Time:", currentTime)

		// รอ 1 วินาทีเพื่อแสดงผลใหม่
		time.Sleep(1 * time.Second)
	}
}
