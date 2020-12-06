package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetInput(year, day int) string {
	return getInput(year, day)
}

func GetInputRows(year, day int) []string {
	input := strings.TrimSpace(GetInput(year, day))
	return strings.Split(input, "\n")
}

func GetInputRowsAsInt(year, day int) []int {
	rows := GetInputRows(year, day)
	intRows := make([]int, len(rows))
	for i, s := range rows {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intRows[i] = x
	}
	return intRows
}
