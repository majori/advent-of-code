package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func fetchInput(year, day int) string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)

	token := os.Getenv("AOC_SESSION")
	if token == "" {
		panic("Environment variable \"AOC_SESSION\" is missing")
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		if err != nil {
			panic(err)
		}

		panic(fmt.Sprintf("Fetching input failed: %s", res.Status))
	}

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(s)
}

func getInput(year, day int) string {
	key := getCacheKey(year, day)

	if input := readCacheValue(key); input != "" {
		return input
	}

	input := fetchInput(year, day)
	if err := writeCacheValue(key, input); err != nil {
		panic(err)
	}

	return input
}

func submit(year, day, level int, answer interface{}) {
	token := os.Getenv("AOC_SESSION")
	if token == "" {
		panic("Environment variable \"AOC_SESSION\" is missing")
	}

	data := url.Values{}
	data.Set("level", fmt.Sprint(level))
	data.Set("answer", fmt.Sprint(answer))

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day), strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	req.AddCookie(&http.Cookie{Name: "session", Value: token})

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
			log.Fatal(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			log.Fatal(err)
	}
	log.Println(string(body))
}