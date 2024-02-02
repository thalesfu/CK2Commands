package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/flags"
)

type SetBuildPeopleFlagScriptGenerator struct {
	Type string
}

func NewSetBuildPeopleFlagScriptGenerator(t string) *SetBuildPeopleFlagScriptGenerator {
	return &SetBuildPeopleFlagScriptGenerator{
		Type: t,
	}
}

func (g *SetBuildPeopleFlagScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	scripts = append(scripts, fmt.Sprintf("set_flag = %s", flags.BuiltPeopleFlag))

	if g.Type == "" {
		return scripts
	}

	builtSpecialPeopleFlag := fmt.Sprintf(flags.BuiltSpecialTypePeopleFlagPattern, g.Type)

	scripts = append(scripts, fmt.Sprintf("set_flag = %s", builtSpecialPeopleFlag))

	return scripts
}

func buildSetBuiltPeopleScriptGenerator(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	peopleType := "common"
	if people.Age < 16 {
		peopleType = "child"
	}

	scriptGenerator.AddScriptGenerator(NewSetBuildPeopleFlagScriptGenerator(peopleType))

	return scriptGenerator
}

func BatchSetBuildPeopleFlag(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, len(people))

	for i, p := range people {
		generators[i] = buildSetBuiltPeopleScriptGenerator(space, p)
	}

	CK2Commands.BuildScript("builtflag", generators...)
}

func SetFriendsAndFriendsFriendsBuildPeopleFlag(space *nebulagolang.Space, people *ck2nebula.People) {
	fr := people.GetFriends(space)

	if fr.Ok {

		friendsMap := fr.Data

		for _, friend := range fr.Data {
			subFr := friend.GetFriends(space)

			if subFr.Ok {
				for _, subFriend := range subFr.Data {
					if _, exist := friendsMap[subFriend.ID]; !exist {
						friendsMap[subFriend.ID] = subFriend
					}
				}
			}
		}

		friends := make([]*ck2nebula.People, len(friendsMap))

		i := 0
		for _, friend := range friendsMap {
			friends[i] = friend
			i++
		}

		BatchSetBuildPeopleFlag(space, friends...)
	}
}
