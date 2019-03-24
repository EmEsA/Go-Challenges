// Calculating Critical Path Method for an input given in /data/input.csv

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Task struct
type Task struct {
	id     string // task id
	es     int    // early start
	ef     int    // early finish
	ls     int    // late start
	lf     int    // late finish
	length int    // duration of the task
}

// State struct
type State struct {
	post []*Task // preceeding tasks
	pre  []*Task // sucessing tasks
}

// Forward pass for a given state
func (state *State) forwardPass() {
	if len(state.pre) == 0 { // start
		for _, task := range state.post {
			task.es = 0
			task.ls = 0
			task.ef = task.length
		}
	} else {
		for _, task := range state.post {
			//es equals maximum of ef of preceeding tasks
			task.es = chooseEs(state.pre)
			task.ef = task.es + task.length
		}
	}
}

// Backward pass for a given state
func (state *State) backwardPass() {
	if len(state.post) == 0 { // end
		for _, task := range state.pre {
			task.lf = task.ef
			task.ls = task.es
		}
	} else {
		for _, task := range state.pre {
			task.lf = chooseLf(state.post)
			task.ls = task.lf - task.length
		}
	}
}

// Choose max from early finish times
func chooseEs(pre []*Task) int {
	maxEf := 0
	for _, task := range pre {
		if task.ef > maxEf {
			maxEf = task.ef
		}
	}
	return maxEf
}

// Choose min from late start times
func chooseLf(post []*Task) int {
	minLs := -1
	for _, task := range post {
		if minLs == -1 {
			minLs = task.ls
		} else if task.ls < minLs {
			minLs = task.ls
		}
	}
	return minLs
}

// getData imports data from csv file and returns a slice of Tasks
func getData(path string) []State {
	csvFile, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	lines, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}
	stateCount, error := strconv.Atoi(lines[len(lines)-1][3])
	if error != nil {
		log.Fatal(error)
	}
	data := make([]State, stateCount)

	for i, line := range lines {
		if i == 0 {
			continue // skip header
		}
		// state from
		state1, error := strconv.Atoi(line[2])
		if error != nil {
			log.Fatal(error)
		}
		state1--
		// state to
		state2, error := strconv.Atoi(line[3])
		if error != nil {
			log.Fatal(error)
		}
		state2--
		// task length
		length, error := strconv.Atoi(line[1])
		if error != nil {
			log.Fatal(error)
		}
		// create task
		var task Task
		if state1 == 0 {
			task = Task{line[0], 0, -1, -1, -1, length}
		} else {
			task = Task{line[0], -1, -1, -1, 0, length}
		}

		data[state1].post = append(data[state1].post, &task)
		data[state2].pre = append(data[state2].pre, &task)
	}

	return data
}

func findCriticalPaths(data []State) map[string][][]string {
	paths := make(map[string][][]string)
	// loop through states
	for _, state := range data {
		// check preceeding tasks
		if len(state.pre) == 0 {
			// initial state, just add post tasks if they do not have any time buffer
			for _, task := range state.post {
				if task.es == task.ls {
					paths[task.id] = [][]string{[]string{task.id}}
				}
			}
		} else if len(state.post) != 0 {
			// regular state
			// loop through pre tasks
			for _, preTask := range state.pre {
				// pop preTask key
				temp, exists := paths[preTask.id]
				if exists {
					delete(paths, preTask.id)
					// loop through post tasks
					for _, postTask := range state.post {
						if postTask.es == postTask.ls {
							// add post task to every path ending with pre task
							// loop through paths
							for _, path := range temp {
								paths[postTask.id] = append(paths[postTask.id], append(path, postTask.id))
							}
						}
					}
				}
			}
		}
	}
	return paths
}

func main() {
	data := getData("data/input.csv")
	// forward pass
	for _, state := range data {
		state.forwardPass()
	}
	// backward pass
	for i := len(data) - 1; i >= 0; i-- {
		data[i].backwardPass()
	}

	paths := findCriticalPaths(data)
	fmt.Print("Critical Paths:")
	i := 1
	for _, pathArr := range paths {
		for _, path := range pathArr {
			fmt.Printf("\n%v: %s", i, strings.Join(path, ", "))
			i++
		}
	}
}
