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

	for d, group := range groups {
		if d == 0 {
			continue
		}
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

			if NeedPollinate(space, p) && p.Dynasty == d {
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
			if v.IsReligionVassal(space) && player.Religion == v.Religion {
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

	BuildAllReligionVassal(space, player, player.Religion, peopleGeneratorMap, nil)
	BuildMarriage(space, player.StoryID, coreFamily, peopleGeneratorMap)

	peopleGenerators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range peopleGeneratorMap {
		peopleGenerators = append(peopleGenerators, p)
	}

	CK2Commands.BuildScript("auto", peopleGenerators...)
}

func BuildAllReligionVassal(space *nebulagolang.Space, people *ck2nebula.People, religion string, allPeople map[int]*PeopleScriptGenerator, scannedPeople map[int]*ck2nebula.People) {
	vr := people.GetVassals(space)

	if scannedPeople == nil {
		scannedPeople = make(map[int]*ck2nebula.People)
	}

	scannedPeople[people.ID] = people

	if vr.Ok {
		for _, v := range vr.Data {
			if v.IsReligionVassal(space) && religion == v.Religion {
				if !v.IsBuilt {
					fsg := buildFatherPeopleScriptGenerator(space, v)
					if pg, ok := allPeople[fsg.Me.ID]; ok {
						for _, g := range fsg.ScriptGenerators {
							pg.AddScriptGenerator(g)
						}
					} else {
						allPeople[fsg.Me.ID] = fsg
					}
				}
			}

			if _, ok := scannedPeople[v.ID]; !ok {
				BuildAllReligionVassal(space, v, religion, allPeople, scannedPeople)
			}
		}
	}
}

func BuildMarriage(space *nebulagolang.Space, storyId int, coreFamily map[int]string, allPeople map[int]*PeopleScriptGenerator) {
	singlePeople := ck2nebula.GetSinglePeople(space, storyId)
	vassalFamilyMen := make(map[int]*ck2nebula.People)
	vassalFamilyWomen := make(map[int]*ck2nebula.People)
	familyMen := make(map[int]*ck2nebula.People)
	familyWomen := make(map[int]*ck2nebula.People)
	vassalMen := make(map[int]*ck2nebula.People)
	vassalWomen := make(map[int]*ck2nebula.People)
	men := make(map[int]*ck2nebula.People)
	women := make(map[int]*ck2nebula.People)

	if singlePeople.Ok {
		for id, p := range singlePeople.Data {
			if _, ok := coreFamily[p.Dynasty]; ok {
				if p.Government == "" {
					if p.IsFemale {
						familyWomen[id] = p
					} else {
						familyMen[id] = p
					}
				} else {
					if p.IsFemale {
						vassalFamilyWomen[id] = p
					} else {
						vassalFamilyMen[id] = p
					}
				}
			} else {
				if p.Government == "" {
					if p.IsFemale {
						women[id] = p
					} else {
						men[id] = p
					}
				} else {
					if p.IsFemale {
						vassalWomen[id] = p
					} else {
						vassalMen[id] = p
					}
				}
			}
		}
	}

	for _, p := range vassalFamilyMen {
		wife := chooseMarriagePartner(p, familyWomen, nil, women)

		if wife != nil {
			hg, wg := buildMarriageScriptGenerator(p, wife, false)
			if pg, ok := allPeople[hg.Me.ID]; ok {
				for _, g := range hg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[hg.Me.ID] = hg
			}

			if pg, ok := allPeople[wg.Me.ID]; ok {
				for _, g := range wg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[wg.Me.ID] = wg
			}
		}
	}

	for _, p := range familyMen {
		wife := chooseMarriagePartner(p, familyWomen, nil, women)

		if wife != nil {
			hg, wg := buildMarriageScriptGenerator(p, wife, false)
			if pg, ok := allPeople[hg.Me.ID]; ok {
				for _, g := range hg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[hg.Me.ID] = hg
			}

			if pg, ok := allPeople[wg.Me.ID]; ok {
				for _, g := range wg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[wg.Me.ID] = wg
			}
		}
	}

	for _, p := range vassalFamilyWomen {
		husband := chooseMarriagePartner(p, familyMen, nil, men)

		if husband != nil {
			hg, wg := buildMarriageScriptGenerator(p, husband, true)
			if pg, ok := allPeople[hg.Me.ID]; ok {
				for _, g := range hg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[hg.Me.ID] = hg
			}

			if pg, ok := allPeople[wg.Me.ID]; ok {
				for _, g := range wg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				allPeople[wg.Me.ID] = wg
			}
		}
	}
}

func chooseMarriagePartner(people *ck2nebula.People, familyCandidates map[int]*ck2nebula.People, vassalCandidates map[int]*ck2nebula.People, normal map[int]*ck2nebula.People) *ck2nebula.People {
	partner := randomChooseAPeople(filterMarriageablePeople(people, familyCandidates))

	if partner != nil {
		delete(familyCandidates, partner.ID)
		return partner
	}

	partner = randomChooseAPeople(filterMarriageablePeople(people, vassalCandidates))

	if partner != nil {
		delete(vassalCandidates, partner.ID)
		return partner
	}

	partner = randomChooseAPeople(filterMarriageablePeople(people, normal))

	if partner != nil {
		delete(normal, partner.ID)
	}

	return partner
}

func filterMarriageablePeople(people *ck2nebula.People, candidates map[int]*ck2nebula.People) []*ck2nebula.People {
	result := make([]*ck2nebula.People, 0)

	if candidates == nil {
		return result
	}

	for _, c := range candidates {
		if people.GrandFather == c.GrandFather {
			continue
		}

		if people.GrandMother == c.GrandFather {
			continue
		}

		if people.MaternalGrandFather == c.MaternalGrandFather {
			continue
		}

		if people.MaternalGrandMother == c.MaternalGrandMother {
			continue
		}

		if people.Father == c.Father {
			continue
		}

		if people.Mother == c.Mother {
			continue
		}

		if !people.IsFemale {
			if c.Age > 45 {
				continue
			}

			if c.IsConsort {
				continue
			}
		} else {
			if people.Age > 45 {
				if people.Age < 50 {
					continue
				}
			}
		}

		if c.Age-people.Age > 10 {
			continue
		}

		result = append(result, c)
	}

	return result
}

func randomChooseAPeople(people []*ck2nebula.People) *ck2nebula.People {
	if len(people) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	return people[rand.Intn(len(people))]
}
