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

	route = append(route, nextSpaceObject)

	return determineOrbits(nextSpaceObject, spaceObjects, route)
}

func main() {
	var input = readInput("input-test-2.txt")
	var spaceObjects = extractSpaceObjects(input)
	orbits := 0

	for _, spaceObject := range spaceObjects {
		var r = determineOrbits(spaceObject, spaceObjects, []SpaceObject{spaceObject})
		orbits += len(r)
	}

	fmt.Println("solution part 1 :", orbits)
	os.Exit(0)
}
