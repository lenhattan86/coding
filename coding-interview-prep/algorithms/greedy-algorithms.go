package main

import (
	"fmt"
	"sort"
)

// Activity represents an activity with a start and finish time.
type Activity struct {
	start  int
	finish int
}

// ByFinish implements sort.Interface for []Activity based on the finish field.
type ByFinish []Activity

func (a ByFinish) Len() int           { return len(a) }
func (a ByFinish) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFinish) Less(i, j int) bool { return a[i].finish < a[j].finish }

// ActivitySelection solves the activity selection problem using a greedy approach.
// It returns the maximum number of non-overlapping activities that can be performed.
func ActivitySelection(activities []Activity) []Activity {
	if len(activities) == 0 {
		return []Activity{}
	}

	// 1. Sort activities by their finish times.
	sort.Sort(ByFinish(activities))

	// 2. Select the first activity.
	selectedActivities := []Activity{activities[0]}
	lastFinishTime := activities[0].finish

	// 3. Iterate through the remaining activities.
	//    If an activity's start time is greater than or equal to the last selected activity's finish time,
	//    select it.
	for i := 1; i < len(activities); i++ {
		if activities[i].start >= lastFinishTime {
			selectedActivities = append(selectedActivities, activities[i])
			lastFinishTime = activities[i].finish
		}
	}

	return selectedActivities
}

func main() {
	fmt.Println("--- Activity Selection Problem ---")

	// Example 1
	activities1 := []Activity{
		{start: 1, finish: 4},
		{start: 3, finish: 5},
		{start: 0, finish: 6},
		{start: 5, finish: 7},
		{start: 3, finish: 8},
		{start: 5, finish: 9},
		{start: 6, finish: 10},
		{start: 8, finish: 11},
		{start: 8, finish: 12},
		{start: 2, finish: 13},
		{start: 12, finish: 14},
	}
	fmt.Println("Activities 1:", activities1)
	selected1 := ActivitySelection(activities1)
	fmt.Println("Selected Activities 1:", selected1) // Expected: [{1 4} {5 7} {8 11} {12 14}]
	fmt.Printf("Max activities: %d\n", len(selected1))
	fmt.Println("--------------------")

	// Example 2
	activities2 := []Activity{
		{start: 10, finish: 20},
		{start: 12, finish: 25},
		{start: 20, finish: 30},
	}
	fmt.Println("Activities 2:", activities2)
	selected2 := ActivitySelection(activities2)
	fmt.Println("Selected Activities 2:", selected2) // Expected: [{10 20} {20 30}]
	fmt.Printf("Max activities: %d\n", len(selected2))
	fmt.Println("--------------------")

	// Example 3 (empty list)
	activities3 := []Activity{}
	fmt.Println("Activities 3:", activities3)
	selected3 := ActivitySelection(activities3)
	fmt.Println("Selected Activities 3:", selected3) // Expected: []
	fmt.Printf("Max activities: %d\n", len(selected3))
	fmt.Println("--------------------")
}
