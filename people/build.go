package people

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/utils"
	"log"
	"math/rand"
	"time"
)

type buildFunc func(*bufio.Writer, int)

func buildPeople(fileName string, buildFuncs []buildFunc, peopleId ...int) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, id := range peopleId {
		if id == 0 {
			continue
		}

		for _, buildFunc := range buildFuncs {
			buildFunc(writer, id)
		}
	}

	writer.Flush()
}

func buildPeopleRandom(fileName string, buildFuncs [][]buildFunc, peopleId ...int) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, id := range peopleId {
		if id == 0 {
			continue
		}

		group := getRandomBuildFunc(buildFuncs)
		for _, buildFunc := range group {
			buildFunc(writer, id)
		}
	}

	writer.Flush()
}

func getRandomBuildFunc(buildFuncs [][]buildFunc) []buildFunc {
	rand.Seed(time.Now().UnixNano())
	return buildFuncs[rand.Intn(len(buildFuncs))]
}
