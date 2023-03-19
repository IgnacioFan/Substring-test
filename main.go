package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var digitOnlyReg = regexp.MustCompile("[0-9]+")

func main() {
	// find the path name for the current directory
	pwd, err := os.Getwd()

	if err = filepath.Walk(pwd+"/doc", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// check if the file is a regular file (not a directory)
		if !info.Mode().IsRegular() {
			return nil
		}
		// Print out sub-strings
		if b, err := os.ReadFile(path); err != nil {
			fmt.Printf("Read file err: %s\n", err)
		} else {
			PrintSubStrings(&b)
			fmt.Println("...")
		}
		return nil
	}); err != nil {
		fmt.Printf("file err: %s\n", err)
	}
}

func PrintSubStrings(b *[]byte) {
	content, ranges := splitContentAndRanges(b)
	for i := 0; i < len(ranges); i++ {
		if len(ranges[i]) != 0 {
			start := ranges[i][0] - 1
			end := ranges[i][1]
			fmt.Println(strings.Join(content[start:end], ""))
		}
	}
}

func splitContentAndRanges(chars *[]byte) (content []string, ranges [][]int) {
	strs := strings.Split(string(*chars), "---")
	return strings.Split(strs[0], ""), collectRanges(strs[1])
}

func collectRanges(ranges string) [][]int {
	rangeStrs := digitOnlyReg.FindAllString(ranges, -1)

	var rangeInts [][]int
	for i := 0; i < len(rangeStrs); i += 2 {
		s1, _ := strconv.Atoi(rangeStrs[i])
		s2, _ := strconv.Atoi(rangeStrs[i+1])
		if s2 > s1 {
			rangeInts = append(rangeInts, []int{s1, s2})
		} else {
			rangeInts = append(rangeInts, []int{s2, s1})
		}
	}
	// sort the rules by ASC order
	sort.Slice(rangeInts, func(i, j int) bool {
		return rangeInts[i][0] < rangeInts[j][0]
	})
	// merge overlap numbers
	for i := 1; i < len(rangeInts); i++ {
		if rangeInts[i-1][1] >= rangeInts[i][0] {
			rangeInts[i][0] = rangeInts[i-1][0]
			rangeInts[i-1] = []int{}
		}
	}
	return rangeInts
}
