package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"log"
	"math/rand"
	"time"
)

type buildFuncWithNebula func(*bufio.Writer, *nebulagolang.Space, *ck2nebula.People)

func buildPeopleWithNebula(fileName string, buildFuncs []buildFuncWithNebula, space *nebulagolang.Space, people ...*ck2nebula.People) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, p := range people {
		for _, bf := range buildFuncs {
			bf(writer, space, p)
		}
	}

	writer.Flush()
}

func buildPeopleRandomWithNebula(fileName string, buildFuncs [][]buildFuncWithNebula, space *nebulagolang.Space, people ...*ck2nebula.People) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, p := range people {
		group := getRandomBuildFuncWithNebula(buildFuncs)
		for _, bf := range group {
			bf(writer, space, p)
		}
	}

	writer.Flush()
}

func getRandomBuildFuncWithNebula(buildFuncs [][]buildFuncWithNebula) []buildFuncWithNebula {
	rand.Seed(time.Now().UnixNano())
	return buildFuncs[rand.Intn(len(buildFuncs))]
}

func AutoBuild(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) {
	file, err := utils.OpenFile("autobuild")
	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	defer func() {
		err := writer.Flush()
		if err != nil {
			log.Println(err)
		}
	}()

	group, newFriends := GetFriendsGroup(space, player, coreFamily)

	MakeFriendsWithNebula(writer, space, group)

	for _, d := range newFriends {
		for _, p := range d {
			BuildSinglePeople(writer, space, p)
		}
	}
}

func BuildPeople(space *nebulagolang.Space, people ...*ck2nebula.People) {
	var functions = []buildFuncWithNebula{
		BuildSinglePeople,
	}

	buildPeopleWithNebula("buildpeople", functions, space, people...)
}

func BuildSinglePeople(writer *bufio.Writer, space *nebulagolang.Space, people *ck2nebula.People) {
	modifiers := GetPropertyModifiers(space, people)

	BuildPeopleTrait(writer, space, people, modifiers)

	for _, modifier := range modifiers {
		if modifier.ModifiedValue > 0 {
			fmt.Printf("%s.%s add %s %d\n", people.DynastyName, people.Name, modifier.Property, modifier.ModifiedValue)
			writeAddPropertyValue(writer, string(modifier.Property), people.ID, modifier.ModifiedValue)
		}
	}
}
