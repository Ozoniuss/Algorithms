package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Mineral int

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
)

type Blueprint struct {
	ID           int
	oreCost      int
	clayCost     int
	obsidianCost [2]int
	geodeCost    [2]int
}

func mustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

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

type State struct {
	minute         int
	oreRobots      int
	ore            int
	clayRobots     int
	clay           int
	obsidianRobots int
	obsidian       int
	geodeRobots    int
	geode          int
}

// Action is the interface that each buying action must implement.
type Action interface {
	doAction(s State, bp Blueprint) State
}

// This round, no robot is bought.
type VoidAction struct{}

type BuyOre struct {
}
type BuyClay struct {
}
type BuyObsidian struct {
}
type BuyGeode struct {
}

func (v VoidAction) doAction(s State, bp Blueprint) State {
	return s
}

func (b BuyOre) doAction(s State, bp Blueprint) State {
	s.ore -= bp.oreCost
	if s.ore < 0 {
		panic("something bad happened")
	}
	s.oreRobots += 1
	return s
}

func (b BuyClay) doAction(s State, bp Blueprint) State {
	s.ore -= bp.clayCost
	if s.ore < 0 {
		panic("something bad happened")
	}
	s.clayRobots += 1
	return s
}

func (b BuyObsidian) doAction(s State, bp Blueprint) State {
	s.ore -= bp.obsidianCost[0]
	s.clay -= bp.obsidianCost[1]
	if s.ore < 0 || s.clay < 0 {
		panic("something bad happened")
	}
	s.obsidianRobots += 1
	return s
}
func (b BuyGeode) doAction(s State, bp Blueprint) State {
	s.ore -= bp.geodeCost[0]
	s.obsidian -= bp.geodeCost[1]
	if s.ore < 0 || s.obsidian < 0 {
		panic("something bad happened")
	}
	s.geodeRobots += 1
	return s
}

func collect(s State) State {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
	return s
}

func canBuyOre(s State, bp Blueprint) bool {
	return s.ore >= bp.oreCost
}

func canBuyClay(s State, bp Blueprint) bool {
	return s.ore >= bp.clayCost
}

func canBuyObsidian(s State, bp Blueprint) bool {
	return s.ore >= bp.obsidianCost[0] && s.clay >= bp.obsidianCost[1]
}

func canBuyGeode(s State, bp Blueprint) bool {
	return s.ore >= bp.geodeCost[0] && s.obsidian >= bp.geodeCost[1]
}

type Stats struct {
	state  State
	action Action
}

// maxTheoreticalGeode returns the maximum theoretical geodes one can collect
// from this state, including the geodes they already have.
func maxTheoreticalGeode(s State, bp Blueprint, rounds int) int {

	//fmt.Println(bp)

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

	//fmt.Println("clay robot at minute", s.minute)

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

	//fmt.Println("obsidian robot at minute", s.minute)

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
	//fmt.Println("geode robot at minute", s.minute)

	// Keep collecting geode
	for s.minute < rounds+1 {
		s.geode += s.geodeRobots
		s.geodeRobots += 1
		s.minute += 1
	}

	return s.geode
}

// determineQualityLevel determines the quality level of a single blueprint.
func determineQualityLevel(bp Blueprint) int {
	initialState := State{
		minute:         0,
		oreRobots:      1,
		ore:            0,
		clayRobots:     0,
		clay:           0,
		obsidianRobots: 0,
		obsidian:       0,
		geodeRobots:    0,
		geode:          0,
	}

	initialStats := Stats{
		state:  initialState,
		action: VoidAction{},
	}

	q := make([]Stats, 0)
	hashes := make(map[Stats]struct{})

	q = append(q, initialStats)
	hashes[initialStats] = struct{}{}

	maxGeode := 0

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		delete(hashes, current)

		// Increase the minute before the collection.
		current.state.minute++
		fmt.Println(current.state)

		// Explored all 23-minute states
		if current.state.minute == 24 {
			return maxGeode
		}

		// Collect the resources at the beginning of each minute.
		current.state = collect(current.state)

		if current.state.minute == 23 {
			switch current.action.(type) {
			// If this round a geode robot is build, we will have one
			// additional geode at the end of minute 24.
			case BuyGeode:
				if current.state.geode+current.state.geodeRobots+1 > maxGeode {
					maxGeode = current.state.geode + current.state.geodeRobots + 1
				}
				// Otherwise, we
			default:
				if current.state.geode+current.state.geodeRobots > maxGeode {
					maxGeode = current.state.geode + current.state.geodeRobots
				}
			}
		}

		// Execute the action.

		//fmt.Printf("initial: %+v\n", current.state)
		current.state = current.action.doAction(current.state, bp)
		//fmt.Printf("after: %+v\n", current.state)

		dontBuyOre := false
		dontBuyClay := false
		dontBuyObsidian := false
		doSomething := false

		// Have the ore for a clay, obsidian and geode at every round. In this
		// case you should never buy ore robots anymore.
		if current.state.oreRobots >= max([]int{bp.clayCost, bp.obsidianCost[0], bp.geodeCost[0]}) {
			dontBuyOre = true
		}

		// Have the clay for an obsidian robot at any round, there's no need to
		// buy clay anymore.
		if current.state.clayRobots >= bp.obsidianCost[1] {
			dontBuyClay = true
		}

		// Same as clay for obsidian. In this case, only buy geode robots from
		// now on.
		if current.state.obsidianRobots >= bp.geodeCost[1] {
			act := BuyGeode{}
			q = append(q, Stats{
				state:  current.state,
				action: act,
			})
			continue
		}

		// If I can buy any of the robots, it's completely pointless to wait
		// further.
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
			//fmt.Println("adding ore action")
		}
		if !dontBuyClay && canBuyClay(current.state, bp) {
			acts = append(acts, BuyClay{})
			//fmt.Println("adding clay action")

		}
		if !dontBuyObsidian && canBuyObsidian(current.state, bp) {
			acts = append(acts, BuyObsidian{})
			//fmt.Println("adding obs action")

		}
		if canBuyGeode(current.state, bp) {
			acts = append(acts, BuyGeode{})
			//fmt.Println("adding geode action")

		}
		if !doSomething {
			acts = append(acts, VoidAction{})
		}
		//fmt.Println("acts", len(acts))
		for _, act := range acts {
			stat := Stats{
				state:  current.state,
				action: act,
			}
			if _, ok := hashes[stat]; !ok {
				q = append(q, stat)
				hashes[stat] = struct{}{}
			}
		}

	}
	return maxGeode
}

func dfs(current Stats, bp Blueprint, maxGeode *int, hashes map[Stats]struct{}, rounds int) {

	// Remove the current state from the hash map.
	//delete(hashes, current)

	// Increase the minute before the collection.
	current.state.minute++
	//fmt.Println(current.state.minute)

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

	dontBuyOre := false
	dontBuyClay := false
	dontBuyObsidian := false
	doSomething := false

	// Have the ore for a clay, obsidian and geode at every round. In this
	// case you should never buy ore robots anymore.
	if current.state.oreRobots >= max([]int{bp.clayCost, bp.obsidianCost[0], bp.geodeCost[0]}) {
		dontBuyOre = true
	}

	// Have the clay for an obsidian robot at any round, there's no need to
	// buy clay anymore.
	if current.state.clayRobots >= bp.obsidianCost[1] {
		dontBuyClay = true
	}

	if current.state.obsidianRobots >= bp.geodeCost[1] {
		dontBuyObsidian = true
	}

	// Same as clay for obsidian. In this case, only buy geode robots from
	// now on.
	if current.state.obsidianRobots >= bp.geodeCost[1] && current.state.oreRobots >= bp.geodeCost[0] {
		act := BuyGeode{}
		stat := Stats{
			state:  current.state,
			action: act,
		}
		hashes[stat] = struct{}{}
		dfs(stat, bp, maxGeode, hashes, rounds)
		return
	}

	// If I can buy any of the robots, it's completely pointless to wait
	// further.
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
		//fmt.Println("adding ore action")
	}
	if !dontBuyClay && canBuyClay(current.state, bp) {
		acts = append(acts, BuyClay{})
		//fmt.Println("adding clay action")

	}
	if !dontBuyObsidian && canBuyObsidian(current.state, bp) {
		acts = append(acts, BuyObsidian{})
		//fmt.Println("adding obs action")

	}
	if canBuyGeode(current.state, bp) {
		acts = append(acts, BuyGeode{})
		//fmt.Println("adding geode action")

	}
	if !doSomething {
		acts = append(acts, VoidAction{})
	}
	//fmt.Println("acts", len(acts))
	for _, act := range acts {
		stat := Stats{
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

	initialState := State{
		minute:         0,
		oreRobots:      1,
		ore:            0,
		clayRobots:     0,
		clay:           0,
		obsidianRobots: 0,
		obsidian:       0,
		geodeRobots:    0,
		geode:          0,
	}

	initialStats := Stats{
		state:  initialState,
		action: VoidAction{},
	}

	part := 2
	if part == 1 {
		quality := 0
		for _, bp := range blueprints {
			maxGeode := 0
			hashStates := make(map[Stats]struct{})
			hashStates[initialStats] = struct{}{}
			dfs(initialStats, bp, &maxGeode, hashStates, 24)
			quality += maxGeode * bp.ID
		}

		fmt.Printf("Quality level: %d", quality)
	} else {
		qualities := [3]int{}
		for i := 0; i < 3; i++ {
			maxGeode := 0
			hashStates := make(map[Stats]struct{})
			hashStates[initialStats] = struct{}{}
			dfs(initialStats, blueprints[i], &maxGeode, hashStates, 32)
			qualities[i] = maxGeode
		}
		fmt.Println(qualities)
	}

}
