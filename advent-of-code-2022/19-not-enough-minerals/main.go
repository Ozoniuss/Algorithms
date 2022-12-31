package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Mineral is an enum defining all possible minerals.
type Mineral int

const (
	ORE Mineral = iota
	CLAY
	OBSIDIAN
	GEODE
)

// Blueprint stores all the data that is available from a blueprint.
type Blueprint struct {
	ID           int
	oreCost      int
	clayCost     int
	obsidianCost [2]int
	geodeCost    [2]int
}

// mustAtoi is a helper that makes strconv.Atoi annoy me less.
func mustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

// max returns the maximum element of an array of integers.
func max(nums []int) int {
	max := -1
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// processLine reads a line and converts it into a Blueprint.
func processLine(line string) Blueprint {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(line, -1)
	return Blueprint{
		ID:       mustAtoi(matches[0]),
		oreCost:  mustAtoi(matches[1]),
		clayCost: mustAtoi(matches[2]),
		obsidianCost: [2]int{
			mustAtoi(matches[3]), mustAtoi(matches[4]),
		},
		geodeCost: [2]int{
			mustAtoi(matches[5]), mustAtoi(matches[6]),
		},
	}
}

// ResourceCount stores the amount of resources available.
type ResourceCount struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

// RobotCount stores the number of robots.
type RobotCount struct {
	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int
}

// EconomyState stores all the relevant information that is relevant for each state of
// the exploration.
type EconomyState struct {
	// minute is relevant for the economy, because the economy grows with time.
	minute int
	ResourceCount
	RobotCount
}

// Action is the interface that every possible action during a round must
// implement.
type Action interface {
	doAction(s EconomyState, bp Blueprint) EconomyState
}

// No robot is bought.
type VoidAction struct{}

// An ore robot is bought.
type BuyOre struct {
}

// A clay robot is bought.
type BuyClay struct {
}

// An obsidian robot is bought.
type BuyObsidian struct {
}

// A geode robot is bought.
type BuyGeode struct {
}

/*
	The changes in state reflected by performing each action, as described in
	the problem statement.
*/

func (v VoidAction) doAction(s EconomyState, bp Blueprint) EconomyState {
	return s
}

func (b BuyOre) doAction(s EconomyState, bp Blueprint) EconomyState {
	s.ore -= bp.oreCost
	if s.ore < 0 {
		panic("something bad happened")
	}
	s.oreRobots += 1
	return s
}

func (b BuyClay) doAction(s EconomyState, bp Blueprint) EconomyState {
	s.ore -= bp.clayCost
	if s.ore < 0 {
		panic("something bad happened")
	}
	s.clayRobots += 1
	return s
}

func (b BuyObsidian) doAction(s EconomyState, bp Blueprint) EconomyState {
	s.ore -= bp.obsidianCost[0]
	s.clay -= bp.obsidianCost[1]
	if s.ore < 0 || s.clay < 0 {
		panic("something bad happened")
	}
	s.obsidianRobots += 1
	return s
}
func (b BuyGeode) doAction(s EconomyState, bp Blueprint) EconomyState {
	s.ore -= bp.geodeCost[0]
	s.obsidian -= bp.geodeCost[1]
	if s.ore < 0 || s.obsidian < 0 {
		panic("something bad happened")
	}
	s.geodeRobots += 1
	return s
}

//Helper functions to determine whether it's possible to perform the purchase.

func canBuyOre(s EconomyState, bp Blueprint) bool {
	return s.ore >= bp.oreCost
}

func canBuyClay(s EconomyState, bp Blueprint) bool {
	return s.ore >= bp.clayCost
}

func canBuyObsidian(s EconomyState, bp Blueprint) bool {
	return s.ore >= bp.obsidianCost[0] && s.clay >= bp.obsidianCost[1]
}

func canBuyGeode(s EconomyState, bp Blueprint) bool {
	return s.ore >= bp.geodeCost[0] && s.obsidian >= bp.geodeCost[1]
}

// collect is called during every state, to simulate robots collecting
// resources.
func collect(s EconomyState) EconomyState {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
	return s
}

// State stores all the relevant information that is required during an
// exploration state.
type State struct {
	state  EconomyState
	action Action
}

// maxTheoreticalGeode returns the maximum theoretical geodes one can collect
// from this state, including the geodes they already have.
func maxTheoreticalGeode(s EconomyState, bp Blueprint, rounds int) int {

	/* Note that state is passed by copy, so it's fine to modify it. */

	// No clay robots yet, try to buy a clay robot as fast as possible.
	if s.clayRobots == 0 {
		// At this state you have a fixed amount of ore and some ore robots.
		// Find out what is the fastest state to get a clay robot, assuming
		// you get an ore robot for free at each round.
		for s.ore < bp.clayCost {
			// Collect the ore and get one more ore robot.
			s.ore += s.oreRobots
			s.oreRobots += 1
			s.minute += 1
			if s.minute == rounds+1 {
				return s.geode
			}
		}
		// Once I have the ore to build a clay robot, the next state I build it.
		s.clayRobots += 1
		if s.minute == rounds+1 {
			return s.geode
		}
	}

	// Same approach as above for obsidian robots. Note that it's not possible
	// to have an obsidian robot before having a clay robot, so the first if
	// should be checked first.
	if s.obsidianRobots == 0 {
		// This is the same process as for clay. It however also assumes that
		// you have the necessary ore to build the obsidian robot.
		for s.clay < bp.obsidianCost[1] {
			// Collect the clay and get one more clay robot.
			s.clay += s.clayRobots
			s.clayRobots += 1
			s.minute += 1
			if s.minute == rounds+1 {
				return s.geode
			}
		}
		s.obsidianRobots += 1
		s.minute += 1
		if s.minute == rounds+1 {
			return s.geode
		}
	}

	// Exact same as for obsidian.
	if s.geodeRobots == 0 {
		for s.obsidian < bp.geodeCost[1] {
			s.obsidian += s.obsidianRobots
			s.obsidianRobots += 1
			s.minute += 1
			if s.minute == rounds+1 {
				return s.geode
			}
		}
		s.geodeRobots += 1
		s.minute += 1
		if s.minute == rounds+1 {
			return s.geode
		}
	}
	// Keep collecting geode
	for s.minute < rounds+1 {
		s.geode += s.geodeRobots
		s.geodeRobots += 1
		s.minute += 1
	}

	return s.geode
}

func dfs(current State, bp Blueprint, maxGeode *int, hashes map[State]struct{}, rounds int) {

	// Increase the minute before the collection.
	current.state.minute++

	if maxTheoreticalGeode(current.state, bp, rounds) <= *maxGeode {
		return
	}

	// Collect the resources at the beginning of each minute.
	current.state = collect(current.state)

	// At minute 23 (or 31) I can determine my geode count and will return.
	if current.state.minute == rounds-1 {
		switch current.action.(type) {
		// If this round a geode robot is build, we will have one
		// additional geode at the end of minute 24.
		case BuyGeode:
			if current.state.geode+current.state.geodeRobots+1 > *maxGeode {
				*maxGeode = current.state.geode + current.state.geodeRobots + 1
			}
			// Otherwise, we
		default:
			if current.state.geode+current.state.geodeRobots > *maxGeode {
				*maxGeode = current.state.geode + current.state.geodeRobots
			}
		}
		return
	}

	// Execute the action.
	current.state = current.action.doAction(current.state, bp)

	// The meaning of these variables is explained below.
	dontBuyOre := false
	dontBuyClay := false
	dontBuyObsidian := false
	doSomething := false

	// Have the ore for a clay, obsidian and geode at every round. In this
	// case you should never buy ore robots anymore.
	if current.state.oreRobots >= max([]int{bp.clayCost, bp.obsidianCost[0], bp.geodeCost[0]}) {
		dontBuyOre = true
	}

	// Have the clay for an obsidian robot at every round. In this case you
	// should never buy clay robots anymore.
	if current.state.clayRobots >= bp.obsidianCost[1] {
		dontBuyClay = true
	}

	// Have the obsidian for a geode robot at every round. In this case you
	// should never buy obsidian robots anymore.
	if current.state.obsidianRobots >= bp.geodeCost[1] {
		dontBuyObsidian = true
	}

	// If you can buy geode robots at every round, it doesn't make sense to buy
	// other robots.
	if current.state.obsidianRobots >= bp.geodeCost[1] && current.state.oreRobots >= bp.geodeCost[0] {
		act := BuyGeode{}
		stat := State{
			state:  current.state,
			action: act,
		}
		hashes[stat] = struct{}{}
		dfs(stat, bp, maxGeode, hashes, rounds)
		return
	}

	// If I can buy any of the robots, don't consider a wait action the next
	// round.
	if canBuyClay(current.state, bp) &&
		canBuyOre(current.state, bp) &&
		canBuyGeode(current.state, bp) &&
		canBuyObsidian(current.state, bp) {
		doSomething = true
	}

	acts := []Action{}

	// We're only interested in buying robots if we can buy them next round
	// and it actually makes sense to buy more.
	if !dontBuyOre && canBuyOre(current.state, bp) {
		acts = append(acts, BuyOre{})
	}
	if !dontBuyClay && canBuyClay(current.state, bp) {
		acts = append(acts, BuyClay{})

	}
	if !dontBuyObsidian && canBuyObsidian(current.state, bp) {
		acts = append(acts, BuyObsidian{})

	}
	if canBuyGeode(current.state, bp) {
		acts = append(acts, BuyGeode{})

	}
	if !doSomething {
		acts = append(acts, VoidAction{})
	}
	for _, act := range acts {
		stat := State{
			state:  current.state,
			action: act,
		}
		if _, ok := hashes[stat]; !ok {
			hashes[stat] = struct{}{}
			dfs(stat, bp, maxGeode, hashes, rounds)
		}
	}

}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	blueprints := make([]Blueprint, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		bp := processLine(line)
		blueprints = append(blueprints, bp)
	}

	initialState := EconomyState{
		minute: 0,
		ResourceCount: ResourceCount{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geode:    0,
		},
		RobotCount: RobotCount{
			oreRobots:      1,
			clayRobots:     0,
			obsidianRobots: 0,
			geodeRobots:    0,
		},
	}

	initialStats := State{
		state:  initialState,
		action: VoidAction{},
	}

	part := 2
	if part == 1 {
		t := time.Now()
		quality := 0
		for _, bp := range blueprints {
			maxGeode := 0
			hashStates := make(map[State]struct{})
			hashStates[initialStats] = struct{}{}
			dfs(initialStats, bp, &maxGeode, hashStates, 24)
			quality += maxGeode * bp.ID
		}
		fmt.Println(time.Since(t))
		fmt.Printf("Quality level: %d", quality)
	} else {
		t := time.Now()
		qualities := [3]int{}
		values := make(chan int, 3)
		for i := 0; i < 3; i++ {
			go func(i int) {
				maxGeode := 0
				hashStates := make(map[State]struct{})
				hashStates[initialStats] = struct{}{}
				dfs(initialStats, blueprints[i], &maxGeode, hashStates, 32)
				values <- maxGeode
			}(i)
		}
		for i := 0; i < 3; i++ {
			qualities[i] = <-values
		}
		fmt.Println(time.Since(t))
		fmt.Printf("Qualities multiply to %d", func(nums [3]int) int {
			return nums[0] * nums[1] * nums[2]
		}(qualities))
	}

}
