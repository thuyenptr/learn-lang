package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func input() (int, error) {
	var n int
	fmt.Println("nhap vao so nguyen khac 0")
	_, err := fmt.Scanf("%d", &n)
	if err == nil && n == 0 {
		return  n, fmt.Errorf("loi so phai la so nguyen khac %d", n)
	}

	return n,err
}

func checkServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("error connect to %s with error %s", url, err)
		//time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server khong phan hoi sau %d", timeout)
}


func main() {
	if len(os.Args) < 2 {
		log.Println("thieu tham so")
		os.Exit(-1)
	}

	if err := checkServer(os.Args[1]); err != nil {
		log.Println("server dead")
		os.Exit(404)
	}

	log.Println("server is running")
}
