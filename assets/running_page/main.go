package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// RunningActivity 表示单次跑步活动数据
type RunningActivity struct {
	Id                 int     `json:"id"`
	StartDate          string  `json:"start_date"`
	Distance           float64 `json:"distance"`
	MovingTime         int     `json:"moving_time"`
	ElapsedTime        int     `json:"elapsed_time"`
	TotalElevationGain float64 `json:"total_elevation_gain"`
	SportType          string  `json:"sport_type"`
	AverageSpeed       float64 `json:"average_speed"`
	MaxSpeed           float64 `json:"max_speed"`
	HasHeartrate       bool    `json:"has_heartrate"`
	AverageHeartrate   float64 `json:"average_heartrate"`
	MaxHeartrate       float64 `json:"max_heartrate"`
	ElevHigh           float64 `json:"elev_high"`
	ElevLow            float64 `json:"elev_low"`
}

type AuthToken struct {
	AccessToken string `json:"access_token"`
}

const (
	fileName = "./assets/running_page/2025_running_page_data.json"
)

func main() {
	// 更新 access_token
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	apiURL := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&refresh_token=%s&grant_type=refresh_token", clientID, clientSecret, refreshToken)
	req, _ := http.NewRequest("POST", apiURL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("获取 API: %s 失败: %v\n", apiURL, err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应数据失败: %v\n", err)
		return
	}
	var authApiResponse AuthToken
	if err := json.Unmarshal(body, &authApiResponse); err != nil {
		fmt.Printf("解析 API: %s 数据失败: %v\n", apiURL, err)
		return
	}
	if authApiResponse.AccessToken == "" {
		fmt.Printf("获取 access_token 失败: %v\n", authApiResponse)
		return
	}

	// 读取本地 data.json 文件（如果存在）
	var existingActivityList []RunningActivity
	if _, err := os.Stat(fileName); err == nil {
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("读取 %s 失败: %v\n", fileName, err)
			return
		}

		if err := json.Unmarshal(fileData, &existingActivityList); err != nil {
			fmt.Printf("解析 %s 失败: %v\n", fileName, err)
			return
		}
	}

	// 请求 api 数据
	latestStartDate, _ := time.Parse(time.RFC3339, existingActivityList[0].StartDate)
	apiURL = fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?after=%d", latestStartDate.Unix())
	// apiURL = fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?per_page=200")
	req, _ = http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authApiResponse.AccessToken))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("获取 API: %s 数据失败: %v\n", apiURL, err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取 API: %s 数据失败: %v\n", apiURL, err)
		return
	}
	var apiResponse []RunningActivity
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("解析 API: %s 数据失败: %v\n", apiURL, err)
		return
	}
	if len(apiResponse) == 0 {
		fmt.Printf("没有新的跑步数据需要更新")
		return
	}

	// 合并数据并根据 id 去重
	activityMap := make(map[int]RunningActivity)
	// 先添加已有数据
	for _, activity := range existingActivityList {
		activityMap[activity.Id] = activity
	}
	// 再添加新数据（如有重复，会覆盖旧数据）
	for _, activity := range apiResponse {
		if activity.SportType != "Run" {
			continue // 只处理跑步活动
		}
		activityMap[activity.Id] = activity
	}
	// 将 map 转回切片
	var allActivityList []RunningActivity
	for _, activity := range activityMap {
		allActivityList = append(allActivityList, activity)
	}

	// 按 start_date 降序排序
	sort.Slice(allActivityList, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, allActivityList[i].StartDate)
		timeJ, _ := time.Parse(time.RFC3339, allActivityList[j].StartDate)
		return timeI.After(timeJ)
	})
	// 保存到 data.json
	jsonData, err := json.MarshalIndent(allActivityList, "", "  ")
	if err != nil {
		fmt.Printf("序列化数据失败: %v\n", err)
		return
	}
	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		fmt.Printf("保存到 %s 失败: %v\n", fileName, err)
		return
	}

	// 生成README.md内容
	readmeContent := `+++
title = "Running Page"
menu = "main"
weight = 200
+++

`
	readmeContent += fmt.Sprintf(`# 我的跑步日记

**最后更新**: %s

![2025 Running Summary](https://raw.githubusercontent.com/A11Might/GitHubPoster/refs/heads/main/OUT_FOLDER/strava.svg)

`, time.Now().Format("2006年1月2日"))

	readmeContent += generateMarkdown(allActivityList)
	if err := os.WriteFile("./content/running_page.md", []byte(readmeContent), 0644); err != nil {
		fmt.Printf("保存到running_page.md失败: %v\n", err)
		return
	}

	fmt.Println("处理完成！数据已保存到data.json和README.md")
}

// formatDuration 将秒数转换为 h:mm:ss 格式
func formatDuration(seconds int) string {
	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("0:%02d:%02d", m, s)
}

// speedToPace 将 m/s 转换为 分:秒/公里 格式
func speedToPace(speed float64) string {
	if speed <= 0 {
		return "0'00\""
	}
	// 米/秒 -> 公里/小时: speed * 3.6
	// 配速(分/公里) = 60 / (speed * 3.6)
	minutesPerKm := 60.0 / (speed * 3.6)
	minutes := int(minutesPerKm)
	seconds := int((minutesPerKm - float64(minutes)) * 60)
	return fmt.Sprintf("%d'%02d\"", minutes, seconds)
}

// generateMarkdown 根据跑步数据生成 Markdown 内容
func generateMarkdown(activities []RunningActivity) string {
	var md strings.Builder

	// 1. 最新活动详情（第一个活动）
	if len(activities) > 0 {
		latest := activities[0]
		startDate, _ := time.Parse(time.RFC3339, latest.StartDate)
		dateStr := startDate.Format("2006年1月2日")
		distanceKm := latest.Distance / 1000
		pace := speedToPace(latest.AverageSpeed)
		movingTime := formatDuration(latest.MovingTime)
		elapsedTime := formatDuration(latest.ElapsedTime)

		md.WriteString(fmt.Sprintf("### %s\n\n", dateStr))
		md.WriteString(fmt.Sprintf("| 📏 **距离**: %.2f公里 | ⚡ **平均配速**: %s/公里 |\n", distanceKm, pace))
		md.WriteString("|---|---|\n")
		md.WriteString(fmt.Sprintf("| ⏱️ **移动时间**: %s | ⏲️ **全程耗时**: %s |\n", movingTime, elapsedTime))
		md.WriteString(fmt.Sprintf("| 👟 **爬升海拔**: %.1f米 | ⛰️ **海拔范围**: %.1f-%.1f米 |\n", latest.TotalElevationGain, latest.ElevLow, latest.ElevHigh))

		if latest.HasHeartrate {
			md.WriteString(fmt.Sprintf("| 💓 **平均心率**: %.1f次/分 | 🔝 **最大心率**: %.0f次/分 |\n", latest.AverageHeartrate, latest.MaxHeartrate))
		} else {
			md.WriteString("| 💓 **平均心率**: - | 🔝 **最大心率**: - |\n")
		}
		md.WriteString("\n")
	}

	// 2. 当月每日数据
	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()
	var currentMonthActivities []RunningActivity

	for _, activity := range activities {
		startDate, _ := time.Parse(time.RFC3339, activity.StartDate)
		if startDate.Month() == currentMonth && startDate.Year() == currentYear {
			currentMonthActivities = append(currentMonthActivities, activity)
		}
	}

	md.WriteString(fmt.Sprintf("### %d年%d月每日数据\n\n", currentYear, currentMonth))
	md.WriteString("| 日期 | 距离 | 平均配速 | 总时间 |\n")
	md.WriteString("|---|---|---|---|\n")

	for _, activity := range currentMonthActivities {
		startDate, _ := time.Parse(time.RFC3339, activity.StartDate)
		dateStr := startDate.Format("2006-01-02")
		distanceKm := activity.Distance / 1000
		pace := speedToPace(activity.AverageSpeed)
		timeStr := formatDuration(activity.ElapsedTime)

		md.WriteString(fmt.Sprintf("| %s | %.2f | %s | %s |\n", dateStr, distanceKm, pace, timeStr))
	}

	md.WriteString("\n")

	// 3. 各月跑步汇总
	monthlySummary := make(map[string]struct {
		TotalDistance float64
		TotalTime     int
		ActivityCount int
		AvgPace       string
	})

	for _, activity := range activities {
		startDate, _ := time.Parse(time.RFC3339, activity.StartDate)
		key := fmt.Sprintf("%d-%02d", startDate.Year(), startDate.Month())
		monthData := monthlySummary[key]
		monthData.TotalDistance += activity.Distance / 1000
		monthData.TotalTime += activity.ElapsedTime
		monthData.ActivityCount++
		monthlySummary[key] = monthData
	}

	// 计算平均配速
	for key, monthData := range monthlySummary {
		if monthData.TotalDistance > 0 && monthData.TotalTime > 0 {
			// 总时间(秒) / 总距离(公里) = 秒/公里
			secPerKm := float64(monthData.TotalTime) / monthData.TotalDistance
			monthData.AvgPace = speedToPace(1000 / secPerKm)
			monthlySummary[key] = monthData
		}
	}

	md.WriteString("### 2025年各月跑步汇总\n\n")
	md.WriteString("| 月份 | 总距离 | 平均配速 | 总时间 | 活动次数 |\n")
	md.WriteString("|---|---|---|---|---|\n")

	// 按月份排序
	var monthKeys []string
	for k := range monthlySummary {
		monthKeys = append(monthKeys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(monthKeys)))

	for _, key := range monthKeys {
		year, month := 0, 0
		fmt.Sscanf(key, "%d-%02d", &year, &month)
		if year != 2025 {
			continue
		}
		monthData := monthlySummary[key]
		md.WriteString(fmt.Sprintf("| **%d月** | %.2f | %s | %s | %d |\n",
			month,
			monthData.TotalDistance,
			monthData.AvgPace,
			formatDuration(monthData.TotalTime),
			monthData.ActivityCount))
	}

	md.WriteString("\n")

	// 4. 历年数据对比
	yearlySummary := make(map[int]struct {
		TotalDistance float64
		TotalTime     float64
		ActivityCount int
	})

	for _, activity := range activities {
		startDate, _ := time.Parse(time.RFC3339, activity.StartDate)
		year := startDate.Year()
		data := yearlySummary[year]
		data.TotalDistance += activity.Distance / 1000
		data.TotalTime += float64(activity.ElapsedTime) / 3600 // 转换为小时
		data.ActivityCount++
		yearlySummary[year] = data
	}

	md.WriteString("### 历年数据对比\n\n")
	md.WriteString("| 年份 | 总距离 | 平均配速 | 总时间(小时) | 活动次数 |\n")
	md.WriteString("|---|---|---|---|---|\n")

	var yearKeys []int
	for y := range yearlySummary {
		yearKeys = append(yearKeys, y)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(yearKeys)))

	for _, year := range yearKeys {
		data := yearlySummary[year]
		avgPace := "-"
		if data.TotalDistance > 0 && data.TotalTime > 0 {
			// 总时间(秒) / 总距离(公里) = 秒/公里
			secPerKm := (data.TotalTime * 3600) / data.TotalDistance
			// 转换为 m/s
			speedInMS := 1000 / secPerKm
			avgPace = speedToPace(speedInMS)
		}
		md.WriteString(fmt.Sprintf("| **%d** | %.2f | %s | %.2f | %d |\n",
			year,
			data.TotalDistance,
			avgPace,
			data.TotalTime,
			data.ActivityCount))
	}

	return md.String()
}
