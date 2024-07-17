package checker

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

func isSocksLive(addr string, target string) bool {
	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		return false
	}

	httpTransport := &http.Transport{
		Dial: dialer.Dial,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
		Timeout:   5 * time.Second,
	}

	response, err := httpClient.Get(target)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	return response.StatusCode == http.StatusOK
}

func checkProxy(proxy string) bool {
	if strings.HasPrefix(proxy, "socks") {
		return isSocksLive(proxy, "http://example.com")
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(proxy)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	return res.StatusCode == 200
}

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for proxy := range jobs {
		if checkProxy(proxy) {
			results <- fmt.Sprintf("✅ LIVE: %s", proxy)
		} else {
			results <- fmt.Sprintf("❌ DEAD: %s", proxy)
		}
	}
}

func CheckProxies(inputFile string, outputFile string, threads int) error {
	fmt.Printf("Checking proxies using %d threads...\n", threads)

	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outFile.Close()

	jobs := make(chan string, threads)
	results := make(chan string, threads)

	var wg sync.WaitGroup

	// start worker goroutines
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// start a goroutine to close the jobs channel when all jobs are sent
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- strings.TrimSpace(scanner.Text())
		}
		close(jobs)
	}()

	// start a goroutine to collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// process results
	for result := range results {
		fmt.Println(result)
		_, err := outFile.WriteString(result + "\n")
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	return nil
}
