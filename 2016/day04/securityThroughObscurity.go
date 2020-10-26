package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/fowbi/advent-of-code-2016/tools"
)

type Line struct {
	Original string
	Name     string
	SectorID int
	Checksum []string
}

func extractLine(line string) Line {
	re := regexp.MustCompile(`([\w-]+)-(\d+)\[(\w+)\]`)
	matches := re.FindSubmatch([]byte(line))

	letters := strings.Split(strings.Join(strings.Split(string(matches[1]), "-"), ``), "")
	sort.Strings(letters)

	return Line{
		line,
		strings.Join(letters, ``),
		tools.Atoi(string(matches[2])),
		strings.Split(string(matches[3]), ""),
	}
}

type SortByLength []string

func (s SortByLength) Len() int {
	return len(s)
}

func (s SortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func isValidRoom(line Line) (bool, int) {
	var ss []string
	js := line.Name
	for i := 0; i < len(js)-1; {
		if js[i+1] != js[i] {
			ss, js = append(ss, js[:i+1]), js[i+1:]
			i = 0
		} else {
			i++
		}

	}
	ss = append(ss, js[:])
	sort.Stable(SortByLength(ss))

	for i, v := range line.Checksum {
		fmt.Println(v, string(ss[i][0]))
		if v != string(ss[i][0]) {
			return false, 0
		}
	}

	return true, line.SectorID
}

func decrypt(line Line) string {
	runes := []rune(line.Name)

	for _, r := range runes {
		tools.Caesar(r, line.SectorID)
	}

	return ""
}

func main() {
	list := tools.ReadLines("input.txt")

	sectorIdSum := 0
	for _, row := range list {
		line := extractLine(row)
		if valid, num := isValidRoom(line); valid {
			sectorIdSum += num

			fmt.Println(decrypt(line))
		}
	}

	fmt.Println("Solution part 1 :", sectorIdSum)
}
