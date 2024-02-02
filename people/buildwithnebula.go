package people

import (
	"bufio"
	"github.com/thalesfu/CK2Commands"
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

	groups := GetFriendsGroup(space, player, coreFamily)

	peopleGeneratorMap := MakeFriendsWithNebula(space, groups)

	for _, group := range groups {
		for _, p := range group {
			if !p.IsBuilt {
				spg := BuildSinglePeople(space, p)
				if pg, ok := peopleGeneratorMap[spg.Me.ID]; ok {
					for _, g := range spg.ScriptGenerators {
						pg.AddScriptGenerator(g)
					}
				} else {
					peopleGeneratorMap[spg.Me.ID] = spg
				}
			}

			cureGenerator := buildCurePeopleScriptGenerator(space, p)

			if len(cureGenerator.ScriptGenerators) > 0 {
				if pg, ok := peopleGeneratorMap[cureGenerator.Me.ID]; ok {
					for _, g := range cureGenerator.ScriptGenerators {
						pg.AddScriptGenerator(g)
					}
				} else {
					peopleGeneratorMap[cureGenerator.Me.ID] = cureGenerator
				}
			}

			if NeedPollinate(space, p) {
				pollinateGenerator := buildPollinatePeopleScriptGenerator(space, p)

				if len(pollinateGenerator.ScriptGenerators) > 0 {
					if pg, ok := peopleGeneratorMap[pollinateGenerator.Me.ID]; ok {
						for _, g := range pollinateGenerator.ScriptGenerators {
							pg.AddScriptGenerator(g)
						}
					} else {
						peopleGeneratorMap[pollinateGenerator.Me.ID] = pollinateGenerator
					}
				}
			}
		}
	}

	vr := player.GetVassals(space)

	if vr.Ok {
		for _, v := range vr.Data {
			if v.IsReligionVassal(space) {
				if !v.IsBuilt {
					fsg := buildFatherPeopleScriptGenerator(space, v)
					if pg, ok := peopleGeneratorMap[fsg.Me.ID]; ok {
						for _, g := range fsg.ScriptGenerators {
							pg.AddScriptGenerator(g)
						}
					} else {
						peopleGeneratorMap[fsg.Me.ID] = fsg
					}
				}
			}
		}
	}

	peopleGenerators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range peopleGeneratorMap {
		peopleGenerators = append(peopleGenerators, p)
	}

	CK2Commands.BuildScript("auto", peopleGenerators...)
}
