package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type SpaceObject struct {
	object          string
	orbittingObject string
}

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func extractSpaceObjects(input string) []SpaceObject {
	var spaceObjects []SpaceObject
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		spaceObjects = append(spaceObjects, SpaceObject{s[0], s[1]})
	}

	return spaceObjects
}

func determineOrbits(spaceObject SpaceObject, spaceObjects []SpaceObject, route []SpaceObject) []SpaceObject {
	var nextSpaceObject = spaceObject

	for _, s := range spaceObjects {
		if s.orbittingObject == spaceObject.object {
			nextSpaceObject = s
			break
		}
	}

	if nextSpaceObject == spaceObject {
		return route
	}

	route = prependSpaceObject(route, nextSpaceObject)

	return determineOrbits(nextSpaceObject, spaceObjects, route)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func prependSpaceObject(x []SpaceObject, y SpaceObject) []SpaceObject {
	x = append(x, SpaceObject{})
	copy(x[1:], x)
	x[0] = y

	return x
}

func main() {
	var input = readInput("input.txt")
	var spaceObjects = extractSpaceObjects(input)
	orbits := 0

	var me, santa []SpaceObject

	for _, spaceObject := range spaceObjects {
		var r = determineOrbits(spaceObject, spaceObjects, []SpaceObject{spaceObject})
		orbits += len(r)

		if spaceObject.orbittingObject == "YOU" {
			me = r
		}

		if spaceObject.orbittingObject == "SAN" {
			santa = r
		}
	}

	var l = Min(len(me), len(santa))
	var pointOfDifference = 0

	for i := 0; i < l; i++ {
		if me[i] != santa[i] {
			pointOfDifference = i
			break
		}
	}

	var orbitalTransfers = (len(me) - 1) + (len(santa) - 1) - (pointOfDifference * 2)

	fmt.Println("solution part 1 :", orbits)

	fmt.Println("solution part 2 :", orbitalTransfers)
	os.Exit(0)
}
