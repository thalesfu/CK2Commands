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

func MakeFriendsWithNebula(writer *bufio.Writer, space *nebulagolang.Space, group [][]*ck2nebula.People) {
	for _, peoples := range group {
		for i := 0; i < len(peoples); i++ {
			r := peoples[i].GetFriends(space)

			for j := i + 1; j < len(peoples); j++ {
				_, exists := r.Data[peoples[j].ID]
				if !r.Ok || !exists {
					writer.WriteString(fmt.Sprintf("add_friend %d %d\n", peoples[i].ID, peoples[j].ID))
					fmt.Printf("%s.%s and %s.%s family are friends\n", peoples[i].DynastyName, peoples[i].Name, peoples[j].DynastyName, peoples[j].Name)
				}
			}
		}
	}
}

func GetFriendsGroup(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) ([][]*ck2nebula.People, map[int]map[int]*ck2nebula.People) {
	result := make([][]*ck2nebula.People, 0)
	newFriendsByFamily := make(map[int]map[int]*ck2nebula.People)

	coreFamilyFriends, newFriends := GetCoreDynastyFriends(space, player, coreFamily)

	for _, p := range newFriends {
		if _, ok := newFriendsByFamily[p.Dynasty]; !ok {
			newFriendsByFamily[p.Dynasty] = make(map[int]*ck2nebula.People)
		}

		newFriendsByFamily[p.Dynasty][p.ID] = p
	}

	for _, p := range coreFamilyFriends {
		familyFriends, nf := GetFamilyFriends(space, p)
		result = append(result, familyFriends)
		for _, p := range nf {
			if _, ok := newFriendsByFamily[p.Dynasty]; !ok {
				newFriendsByFamily[p.Dynasty] = make(map[int]*ck2nebula.People)
			}

			newFriendsByFamily[p.Dynasty][p.ID] = p
		}
	}

	return result, newFriendsByFamily
}

func GetCoreDynastyFriends(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) ([]*ck2nebula.People, []*ck2nebula.People) {
	newFriends := make([]*ck2nebula.People, 0)
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

	r := player.GetVassals(space)

	if !r.Ok {
		return nil, nil
	}

	familyCourtiers := make(map[int][]*ck2nebula.People)

	for _, p := range r.Data {
		if _, ok := coreFamily[p.Dynasty]; ok {
			if _, ok := familyCourtiers[p.Dynasty]; !ok {
				familyCourtiers[p.Dynasty] = make([]*ck2nebula.People, 0)
			}

			familyCourtiers[p.Dynasty] = append(familyCourtiers[p.Dynasty], p)
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

		if fcs, ok := familyCourtiers[f]; ok {
			if len(fcs) == 1 {
				result = append(result, fcs[0])
				newFriends = append(newFriends, fcs[0])
			} else if len(fcs) > 1 {
				people := getTopClosedPeople(fcs, player)
				result = append(result, people)
				newFriends = append(newFriends, people)
			}

			delete(families, f)
			continue
		}
	}

	return result, newFriends
}

func GetFamilyFriends(space *nebulagolang.Space, player *ck2nebula.People) ([]*ck2nebula.People, []*ck2nebula.People) {
	newFriends := make([]*ck2nebula.People, 0)
	result := make([]*ck2nebula.People, 0)

	friendsResult := player.GetFriends(space)

	fr := player.GetFamilies(space)

	if fr.Ok {
		for _, f := range fr.Data {
			result = append(result, f)
		}
	}

	br := player.GetBrothersAndSisters(space)

	if br.Ok {
		for _, b := range br.Data {
			result = append(result, b)
		}
	}

	for _, p := range result {
		if friendsResult.Ok {
			if _, ok := friendsResult.Data[p.ID]; !ok {
				newFriends = append(newFriends, p)
			}
		}
	}

	result = append(result, player)

	return result, newFriends
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
