package search

import (
	"fmt"
	"log"
)

//保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的
// 新类型搜索的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数， 为每个数据源单独启动goroutine来执行这个函数
// 并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searhResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searhResults {
		results <- result
	}
}

// Display 从每个单独的goroutine接收到结果后
// 在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直被阻塞， 直到有结果写入
	// 一旦通道关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
