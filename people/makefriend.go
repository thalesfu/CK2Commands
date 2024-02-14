package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"log"
	"math"
	"sort"
)

func MakeFriends(group ...[]int) {
	file, err := utils.OpenFile("makefriends")
	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for gIndex, peoples := range group {
		peoples = utils.Filter[int](peoples, func(people int) bool {
			return people != 0
		})
		for i := 0; i < len(peoples); i++ {
			for j := i + 1; j < len(peoples); j++ {
				writer.WriteString(fmt.Sprintf("add_friend %d %d", peoples[i], peoples[j]))

				if !(gIndex == len(group)-1 && i == len(peoples)-2 && j == len(peoples)-1) {
					writer.WriteString("\n")
				}
			}
		}
	}

	writer.Flush()
}

func MakeFriendsWithNebula(space *nebulagolang.Space, group map[int][]*ck2nebula.People) map[int]*PeopleScriptGenerator {
	result := make(map[int]*PeopleScriptGenerator)
	for _, peoples := range group {
		for i := 0; i < len(peoples); i++ {
			p, ok := result[peoples[i].ID]
			if !ok {
				p = NewPeopleScriptGenerator(peoples[i])
				result[peoples[i].ID] = p
			}
			r := peoples[i].GetFriends(space)

			for j := i + 1; j < len(peoples); j++ {
				_, exists := r.Data[peoples[j].ID]
				if !r.Ok || !exists {
					fmt.Printf("%s.%s(%d) and %s.%s(%d) are friends\n", peoples[i].DynastyName, peoples[i].Name, peoples[i].ID, peoples[j].DynastyName, peoples[j].Name, peoples[j].ID)
					p.AddScriptGenerator(NewMakeFriendScriptGenerator(peoples[j]))
				}
			}
		}
	}

	return result
}

func GetFriendsGroup(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) map[int][]*ck2nebula.People {
	result := make(map[int][]*ck2nebula.People)

	coreFamilyFriends := GetCoreDynastyFriends(space, player, coreFamily)

	for _, p := range coreFamilyFriends {
		familyFriends := GetFamilyFriends(space, p)
		result[p.Dynasty] = familyFriends
	}

	result[0] = coreFamilyFriends

	return result
}

func GetCoreDynastyFriends(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) []*ck2nebula.People {
	fr := player.GetFriends(space)

	friendsFamilies := make(map[int][]*ck2nebula.People)

	if fr.Ok {
		for _, p := range fr.Data {
			if _, ok := coreFamily[p.Dynasty]; ok {
				if _, ok := friendsFamilies[p.Dynasty]; !ok {
					friendsFamilies[p.Dynasty] = make([]*ck2nebula.People, 0)
				}

				friendsFamilies[p.Dynasty] = append(friendsFamilies[p.Dynasty], p)
			}
		}
	}

	families := make(map[int]string)

	for v, k := range coreFamily {
		families[v] = k
	}

	familyVassals := make(map[int][]*ck2nebula.People)

	er := player.GetEmpirePeople(space)

	if er.Ok {
		for _, p := range er.Data {
			if _, ok := coreFamily[p.Dynasty]; ok {
				if _, ok := familyVassals[p.Dynasty]; !ok {
					familyVassals[p.Dynasty] = make([]*ck2nebula.People, 0)
				}

				familyVassals[p.Dynasty] = append(familyVassals[p.Dynasty], p)
			}
		}
	}

	vr := player.GetVassals(space)

	if vr.Ok {
		for _, p := range vr.Data {
			if _, ok := coreFamily[p.Dynasty]; ok {
				if _, ok := familyVassals[p.Dynasty]; !ok {
					familyVassals[p.Dynasty] = make([]*ck2nebula.People, 0)
				}

				familyVassals[p.Dynasty] = append(familyVassals[p.Dynasty], p)
			}
		}
	}

	result := make([]*ck2nebula.People, 0)
	result = append(result, player)

	for f, _ := range families {
		if fcs, ok := friendsFamilies[f]; ok {
			if len(fcs) == 1 {
				result = append(result, fcs[0])
			} else if len(fcs) > 1 {
				result = append(result, getTopClosedPeople(fcs, player))
			}

			delete(families, f)
			continue
		}

		if fcs, ok := familyVassals[f]; ok {
			if len(fcs) == 1 {
				result = append(result, fcs[0])
			} else if len(fcs) > 1 {
				people := getTopClosedPeople(fcs, player)
				result = append(result, people)
			}

			delete(families, f)
			continue
		}
	}

	for f, _ := range families {
		dynasty := ck2nebula.NewDynasty(player.StoryID, f)

		dpr := dynasty.GetAlivePeoples(space)
		if dpr.Ok {
			candidates := make([]*ck2nebula.People, 0)

			for _, p := range dpr.Data {
				candidates = append(candidates, p)
			}

			if len(candidates) == 1 {
				result = append(result, candidates[0])
			} else if len(candidates) > 1 {
				people := getTopClosedPeople(candidates, player)
				result = append(result, people)
			}

			delete(families, f)
		}
	}

	myDistanceFamilyResult := player.GetAliveDistanceFamilies(space)
	if myDistanceFamilyResult.Ok {
		candidates := make([]*ck2nebula.People, 0)

		for _, p := range myDistanceFamilyResult.Data {
			if _, ok := fr.Data[p.ID]; !ok {
				candidates = append(candidates, p)
			} else {
				return result
			}
		}

		if len(candidates) == 1 {
			result = append(result, candidates[0])
		} else if len(candidates) > 1 {
			people := getTopClosedPeople(candidates, player)
			result = append(result, people)
		}
	}

	return result
}

func GetFamilyFriends(space *nebulagolang.Space, player *ck2nebula.People) []*ck2nebula.People {
	result := make([]*ck2nebula.People, 0)

	fr := player.GetAliveFamilies(space)

	if fr.Ok {
		for _, f := range fr.Data {
			result = append(result, f)
		}
	}

	br := player.GetAliveBrothersAndSisters(space)

	if br.Ok {
		for _, b := range br.Data {
			result = append(result, b)
		}
	}

	result = append(result, player)

	return result
}

func getTopClosedPeople(people []*ck2nebula.People, base *ck2nebula.People) *ck2nebula.People {
	if len(people) == 0 {
		return nil
	}

	sort.Slice(people, func(i, j int) bool {
		iage := math.Abs(float64(people[i].Age - base.Age))
		jage := math.Abs(float64(people[j].Age - base.Age))

		if iage < 5 && jage >= 5 {
			return true
		}

		if iage >= 5 && jage < 5 {
			return false
		}

		if !people[i].IsFemale && people[j].IsFemale {
			return true
		}

		if people[i].IsFemale && !people[j].IsFemale {
			return false
		}

		ipoints := people[i].ModifiedDiplomacy + people[i].ModifiedMartial + people[i].Stewardship + people[i].Intrigue + people[i].Learning

		if people[i].Age < 16 {
			ipoints = ipoints * 2
		}

		jpoints := people[j].ModifiedDiplomacy + people[j].ModifiedMartial + people[j].Stewardship + people[j].Intrigue + people[j].Learning

		if people[j].Age < 16 {
			jpoints = jpoints * 2
		}

		return ipoints > jpoints
	})

	return people[0]
}
