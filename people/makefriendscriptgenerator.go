package people

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

type MakeFriendScriptGenerator struct {
	Friend *ck2nebula.People
}

func NewMakeFriendScriptGenerator(friend *ck2nebula.People) *MakeFriendScriptGenerator {
	return &MakeFriendScriptGenerator{
		Friend: friend,
	}

}

func (m *MakeFriendScriptGenerator) Generate() []string {
	scripts := make([]string, 0)

	gender := "男"

	if m.Friend.IsFemale {
		gender = "女"
	}

	var comment string

	if m.Friend.DynastyName != "" {
		comment = fmt.Sprintf("# 和 %s 家的 %s (%s) 交朋友", m.Friend.DynastyName, m.Friend.Name, gender)
	} else {
		comment = fmt.Sprintf("# 和 %s (%s) 交朋友", m.Friend.Name, gender)
	}

	scripts = append(scripts, fmt.Sprintf("add_friend = %d\t%s", m.Friend.ID, comment))

	return scripts
}
