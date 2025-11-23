// https://go.dev/tour/concurrency/10
package concurrency

import (
	"fmt"
	"sync"
)

// 練習：並行網頁爬蟲
//
// 任務要求：
// 1. 修改 Crawl 函數以並行獲取 URL
// 2. 確保不會重複獲取同一個 URL
// 3. 使用 map 來緩存已獲取的 URL，但需要確保並發安全（使用 mutex）

type Fetcher interface {
	// Fetch 返回 URL 的內容和該頁面上找到的 URL 切片
	Fetch(url string) (body string, urls []string, err error)
}

// urlCache 是一個並發安全的 URL 緩存
// 用於記錄已經獲取過的 URL，避免重複獲取
type urlCache struct {
	mu   sync.Mutex
	urls map[string]bool
}

// newURLCache 創建一個新的 URL 緩存
func newURLCache() *urlCache {
	return &urlCache{
		urls: make(map[string]bool),
	}
}

// tryMarkVisited 嘗試標記 URL 為已訪問
// 如果 URL 已經被訪問過，返回 false；否則標記並返回 true
// 這是一個原子操作，確保並發安全
func (c *urlCache) tryMarkVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.urls[url] {
		return false // 已經訪問過
	}
	c.urls[url] = true // 標記為已訪問
	return true        // 成功標記
}

// Crawl 使用 fetcher 遞歸爬取從 url 開始的頁面，最大深度為 depth
// 修改為並行版本，使用 goroutine 並行獲取 URL，並使用緩存避免重複獲取
func Crawl(url string, depth int, fetcher Fetcher, cache *urlCache) {
	// 如果深度小於等於 0，停止遞歸
	if depth <= 0 {
		return
	}

	// 原子地嘗試標記 URL 為已訪問
	// 如果 URL 已經被訪問過，直接返回，避免重複獲取
	// 這確保了在並發環境下不會重複獲取同一個 URL
	if !cache.tryMarkVisited(url) {
		return // URL 已經被訪問過，直接返回
	}

	// 獲取 URL 的內容
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印找到的內容
	fmt.Printf("found: %s %q\n", url, body)

	// 使用 WaitGroup 等待所有子 goroutine 完成
	var wg sync.WaitGroup

	// 並行處理所有找到的 URL
	for _, u := range urls {
		wg.Add(1) // 增加等待計數
		go func(u string) {
			defer wg.Done() // goroutine 完成時減少計數
			// 遞歸爬取子 URL，深度減 1
			Crawl(u, depth-1, fetcher, cache)
		}(u) // 注意：必須傳遞 u 作為參數，避免閉包問題
	}

	// 等待所有子 goroutine 完成
	wg.Wait()
}

func RunConcurrency10() {
	fmt.Println("=== 並行網頁爬蟲示例 ===")
	fmt.Println("使用 goroutine 並行獲取 URL，並使用 mutex 保護 URL 緩存")
	fmt.Println()

	// 創建 URL 緩存
	cache := newURLCache()

	// 開始爬取，從 "https://golang.org/" 開始，最大深度為 4
	Crawl("https://golang.org/", 4, fetcher, cache)

	fmt.Println()
	fmt.Println("=== 實現說明 ===")
	fmt.Println("1. 使用 sync.Mutex 保護 URL 緩存 map，確保並發安全")
	fmt.Println("2. 在獲取 URL 前檢查緩存，避免重複獲取")
	fmt.Println("3. 使用 goroutine 並行處理所有子 URL")
	fmt.Println("4. 使用 sync.WaitGroup 等待所有 goroutine 完成")
	fmt.Println("5. 遞歸深度控制確保不會無限遞歸")
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
