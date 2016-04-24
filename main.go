package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/vintikzzz/proxy-check/check"
)

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println(" cat proxylist.txt | proxy-check [<args>]")
		fmt.Println("Proxy format:")
		fmt.Println(" schema://ip:port")
		return
	}
	workerFlag := flag.Int("workers", 20, "Number of workers")
	targetFlag := flag.String("target", "http://aliexpress.com", "Target test url")
	timeoutFlag := flag.Int("timeout", 5, "Timeout in seconds")
	checkFlag := flag.String("check", "Aliexpress", "Text that expected at target site")
	verboseFlag := flag.Bool("verbose", false, "Enables verbose mode")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	in := make(chan *url.URL)
	out := make(chan *check.Result)
	targetURL, _ := url.Parse(*targetFlag)
	duration := time.Duration(int64(time.Second) * int64(*timeoutFlag))
	check.Check(in, out, duration, targetURL, *checkFlag, *workerFlag)
	go func() {
		for scanner.Scan() {
			proxyURL, _ := url.Parse(scanner.Text())
			in <- proxyURL
		}
		close(in)
	}()
	for r := range out {
		if r.Err == nil {
			if *verboseFlag {
				fmt.Printf("URL:      %s\n", r.URL.String())
				fmt.Printf("Delay(s): %.2f\n", time.Duration(r.Delay).Seconds())
				fmt.Printf("Result:   OK\n")
				fmt.Println()
			} else {
				fmt.Println(r.URL.String())
			}
		} else {
			if *verboseFlag {
				fmt.Printf("URL:      %s\n", r.URL.String())
				fmt.Printf("Error:    %s\n", r.Err)
				fmt.Printf("Result:   FAIL\n")
				fmt.Println()
			}
		}
	}
}
