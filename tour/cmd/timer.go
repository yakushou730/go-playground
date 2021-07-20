package cmd

import (
	"log"
	"playground/tour/internal/timer"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間處理格式",
	Long:  "時間處理格式",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得目前時間",
	Long:  "取得目前時間",
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := time.LoadLocation("Asia/Taipei")
		nowTime := timer.GetNowTime().In(location)
		log.Printf("輸出結果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "計算所需時間",
	Long:  "計算所需時間",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("輸出結果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要計算的時間，有效單位為時間戳記或已格式化後的時間`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持續時間，有效單位為"ns", "us", "ms", "s", "m", "h"`)
}
