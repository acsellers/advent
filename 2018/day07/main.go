package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var stepRegex = regexp.MustCompile(`Step (.) must be finished before step (.) can begin.`)

func main() {
	fmt.Println(Do(demo))
	fmt.Println(Do(input))
	fmt.Println(Do2(demo, 0, 2))
	fmt.Println(Do2(input, 60, 5))
}

func Do(input string) string {
	steps := map[string]*Step{}
	for _, line := range strings.Split(input, "\n") {
		parts := stepRegex.FindStringSubmatch(line)[1:]
		if _, ok := steps[parts[0]]; !ok {
			steps[parts[0]] = &Step{Name: parts[0]}
		}
		if _, ok := steps[parts[1]]; !ok {
			steps[parts[1]] = &Step{Name: parts[1]}
		}
		steps[parts[1]].Pre = append(steps[parts[1]].Pre, parts[0])
	}

	done := ""
	for len(done) < len(steps) {
		runnable := []string{}
	StepCheck:
		for _, step := range steps {
			if strings.Contains(done, step.Name) {
				continue StepCheck
			}
			for _, pre := range step.Pre {
				if !strings.Contains(done, pre) {
					continue StepCheck
				}
			}
			runnable = append(runnable, step.Name)
		}
		sort.Strings(runnable)
		done += runnable[0]
	}
	return done
}

func Do2(input string, base int, workerCount int) string {
	steps := map[string]*Step{}
	for _, line := range strings.Split(input, "\n") {
		parts := stepRegex.FindStringSubmatch(line)[1:]
		if _, ok := steps[parts[0]]; !ok {
			steps[parts[0]] = &Step{Name: parts[0], Done: -1, Worker: -1}
		}
		if _, ok := steps[parts[1]]; !ok {
			steps[parts[1]] = &Step{Name: parts[1], Done: -1, Worker: -1}
		}
		steps[parts[1]].Pre = append(steps[parts[1]].Pre, parts[0])
	}

	workers := make([]string, 5)

	done := ""
	time := 0
	fmt.Println("T\t1\t2\t3\t4\t5\tDone")
	for len(done) < len(steps) {
		runnable := []string{}
		for _, step := range steps {
			if step.Done == time {
				done += step.Name
				workers[step.Worker] = ""
			}
		}
	StepCheck:
		for _, step := range steps {
			if strings.Contains(done, step.Name) {
				continue StepCheck
			}
			for _, pre := range step.Pre {
				if !strings.Contains(done, pre) {
					continue StepCheck
				}
			}
			if step.Done > 0 {
				continue StepCheck
			}
			runnable = append(runnable, step.Name)
		}
		sort.Strings(runnable)
		assignable := strings.Join(runnable, "")
		for i := 0; i < workerCount; i++ {
			worker := workers[i]
			if worker == "" {
				if len(runnable) > 0 {
					task := runnable[0]
					workers[i] = steps[task].Name
					steps[task].Worker = i
					steps[task].Done = time + base + int([]rune(steps[task].Name)[0]-'A'+1)
					runnable = runnable[1:]
				}
			}
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", time, workers[0], workers[1], workers[2], workers[3], workers[4], done, assignable)
		// distribute steps
		time++
	}
	return fmt.Sprint(time - 1)
}

type Step struct {
	Name   string
	Pre    []string
	Done   int
	Worker int
}

var demo = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

var input = `Step Y must be finished before step L can begin.
Step N must be finished before step D can begin.
Step Z must be finished before step A can begin.
Step F must be finished before step L can begin.
Step H must be finished before step G can begin.
Step I must be finished before step S can begin.
Step M must be finished before step U can begin.
Step R must be finished before step J can begin.
Step T must be finished before step D can begin.
Step U must be finished before step D can begin.
Step O must be finished before step X can begin.
Step B must be finished before step D can begin.
Step X must be finished before step V can begin.
Step J must be finished before step V can begin.
Step D must be finished before step A can begin.
Step K must be finished before step P can begin.
Step Q must be finished before step C can begin.
Step S must be finished before step E can begin.
Step A must be finished before step V can begin.
Step G must be finished before step L can begin.
Step C must be finished before step W can begin.
Step P must be finished before step W can begin.
Step V must be finished before step W can begin.
Step E must be finished before step W can begin.
Step W must be finished before step L can begin.
Step P must be finished before step E can begin.
Step T must be finished before step K can begin.
Step A must be finished before step G can begin.
Step G must be finished before step P can begin.
Step N must be finished before step S can begin.
Step R must be finished before step D can begin.
Step M must be finished before step G can begin.
Step Z must be finished before step L can begin.
Step M must be finished before step T can begin.
Step S must be finished before step L can begin.
Step S must be finished before step W can begin.
Step O must be finished before step J can begin.
Step Z must be finished before step D can begin.
Step A must be finished before step C can begin.
Step P must be finished before step V can begin.
Step A must be finished before step P can begin.
Step B must be finished before step C can begin.
Step R must be finished before step S can begin.
Step X must be finished before step S can begin.
Step T must be finished before step P can begin.
Step Y must be finished before step E can begin.
Step G must be finished before step E can begin.
Step Y must be finished before step K can begin.
Step J must be finished before step P can begin.
Step I must be finished before step Q can begin.
Step E must be finished before step L can begin.
Step X must be finished before step J can begin.
Step T must be finished before step X can begin.
Step M must be finished before step O can begin.
Step K must be finished before step A can begin.
Step D must be finished before step W can begin.
Step H must be finished before step C can begin.
Step F must be finished before step R can begin.
Step B must be finished before step Q can begin.
Step M must be finished before step Q can begin.
Step D must be finished before step S can begin.
Step Y must be finished before step I can begin.
Step M must be finished before step K can begin.
Step S must be finished before step G can begin.
Step X must be finished before step L can begin.
Step D must be finished before step V can begin.
Step B must be finished before step X can begin.
Step C must be finished before step L can begin.
Step V must be finished before step L can begin.
Step Z must be finished before step Q can begin.
Step Z must be finished before step H can begin.
Step M must be finished before step S can begin.
Step O must be finished before step C can begin.
Step B must be finished before step A can begin.
Step U must be finished before step V can begin.
Step U must be finished before step A can begin.
Step X must be finished before step G can begin.
Step K must be finished before step C can begin.
Step T must be finished before step S can begin.
Step K must be finished before step G can begin.
Step U must be finished before step B can begin.
Step A must be finished before step E can begin.
Step F must be finished before step V can begin.
Step Q must be finished before step A can begin.
Step F must be finished before step Q can begin.
Step J must be finished before step L can begin.
Step O must be finished before step E can begin.
Step O must be finished before step Q can begin.
Step I must be finished before step K can begin.
Step I must be finished before step P can begin.
Step J must be finished before step D can begin.
Step Q must be finished before step P can begin.
Step S must be finished before step C can begin.
Step U must be finished before step P can begin.
Step S must be finished before step P can begin.
Step O must be finished before step B can begin.
Step Z must be finished before step F can begin.
Step R must be finished before step V can begin.
Step D must be finished before step L can begin.
Step Y must be finished before step T can begin.
Step G must be finished before step C can begin.`
