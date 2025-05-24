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

// Memo 结构体用于保存所需的字段
type Memo struct {
	Name        string   `json:"name"`
	DisplayTime string   `json:"displayTime"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
}

// APIResponse 解析API返回的JSON数据
type APIResponse struct {
	Memos []struct {
		Name        string   `json:"name"`
		DisplayTime string   `json:"displayTime"`
		Content     string   `json:"content"`
		Tags        []string `json:"tags"`
	} `json:"memos"`
	NextPageToken string `json:"nextPageToken"`
}

func main() {
	// 1. 获取API数据
	// 只会获取最新的 10 条 memo，也就是说只有最新的 10 条 memo 会在每次脚本触发时更新内容
	memosHost := os.Getenv("MEMOS_HOST")
	if memosHost == "" {
		fmt.Println("警告: 未设置MEMOS_HOST环境变量")
		return
	}
	apiURL := memosHost + "/api/v1/memos?parent=users/1&filter=tag%20in%20%5B%222025%22%5D"

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("获取API数据失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应数据失败: %v\n", err)
		return
	}

	// 2. 解析API返回的JSON数据
	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("解析API数据失败: %v\n", err)
		return
	}

	// 3. 将API数据转换为Memo结构体
	var newMemos []Memo
	for _, m := range apiResponse.Memos {
		newMemos = append(newMemos, Memo{
			Name:        m.Name,
			DisplayTime: m.DisplayTime,
			Content:     m.Content,
			Tags:        m.Tags,
		})
	}

	// 4. 读取本地data.json文件（如果存在）
	var existingMemos []Memo
	if _, err := os.Stat("data.json"); err == nil {
		fileData, err := os.ReadFile("data.json")
		if err != nil {
			fmt.Printf("读取data.json失败: %v\n", err)
			return
		}

		if err := json.Unmarshal(fileData, &existingMemos); err != nil {
			fmt.Printf("解析data.json失败: %v\n", err)
			return
		}
	}

	// 5. 合并数据并根据name去重
	memoMap := make(map[string]Memo)

	// 先添加已有数据
	for _, memo := range existingMemos {
		memoMap[memo.Name] = memo
	}

	// 再添加新数据（如有重复，会覆盖旧数据）
	for _, memo := range newMemos {
		memoMap[memo.Name] = memo
	}

	// 将map转回切片
	var allMemos []Memo
	for _, memo := range memoMap {
		allMemos = append(allMemos, memo)
	}

	// 6. 按DisplayTime降序排序
	sort.Slice(allMemos, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, allMemos[i].DisplayTime)
		timeJ, _ := time.Parse(time.RFC3339, allMemos[j].DisplayTime)
		return timeI.After(timeJ)
	})

	// 7. 保存到data.json
	jsonData, err := json.MarshalIndent(allMemos, "", "  ")
	if err != nil {
		fmt.Printf("序列化数据失败: %v\n", err)
		return
	}

	if err := os.WriteFile("data.json", jsonData, 0644); err != nil {
		fmt.Printf("保存到data.json失败: %v\n", err)
		return
	}

	// 8. 生成README.md内容
	readmeContent := `---
date: '2025-01-18'
draft: false
title: '2025'
ShowBreadCrumbs: false
---
`
	for _, memo := range allMemos {
		// 解析时间并格式化为YYYY-MM-DD格式
		t, _ := time.Parse(time.RFC3339, memo.DisplayTime)
		formattedDate := t.Format("2006-01-02")

		// 添加标题和内容
		readmeContent += fmt.Sprintf("### %s\n\n", formattedDate)
		memoContent := strings.ReplaceAll(memo.Content, "\n", "\n\n")
		memoContent = strings.ReplaceAll(memoContent, "\n>", ">")
		memoContent = strings.Replace(memoContent, ">", "\n>", 1)
		readmeContent += fmt.Sprintf("%s\n\n", memoContent)
	}

	// 9. 保存到README.md
	if err := os.WriteFile("./content/posts/2025-01-18-2025-memo.md", []byte(readmeContent), 0644); err != nil {
		fmt.Printf("保存到README.md失败: %v\n", err)
		return
	}

	fmt.Println("处理完成！数据已保存到data.json和README.md")
}
