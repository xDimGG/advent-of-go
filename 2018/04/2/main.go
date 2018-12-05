package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	actionBeginShift = "BEGIN_SHIFT"
	actionFallAsleep = "FALL_ASLEEP"
	actionWakeUp     = "WAKE_UP"
)

type log struct {
	id     int
	date   string
	time   int
	action string
}

func parseLog(raw string) (l *log) {
	m, err := strconv.Atoi(raw[15:17])
	if err != nil {
		return
	}

	l = &log{date: raw[1:17], time: m}

	switch raw[19] {
	case 'G':
		fmt.Sscanf(raw[19:], "Guard #%d", &l.id)
		l.action = actionBeginShift
	case 'f':
		l.action = actionFallAsleep
	case 'w':
		l.action = actionWakeUp
	}

	return l
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	allLogs := make([]*log, 0)

	for scanner.Scan() {
		allLogs = append(allLogs, parseLog(scanner.Text()))
	}

	sort.Slice(allLogs, func(i, j int) bool {
		return allLogs[i].date < allLogs[j].date
	})

	sleepTimes := make(map[int]int)

	{
		var id, sleepTime int

		for _, log := range allLogs {
			switch log.action {
			case actionBeginShift:
				id = log.id
			case actionFallAsleep:
				sleepTime = log.time
			case actionWakeUp:
				for i := sleepTime; i < log.time; i++ {
					sleepTimes[(id<<6)|i]++
				}
			}
		}
	}

	highestCountID := 0
	highestCount := 0

	for id, count := range sleepTimes {
		if count > highestCount {
			highestCountID = id
			highestCount = count
		}
	}

	fmt.Println("Result:", (highestCountID>>6)*(highestCountID&((1<<6)-1)))
}
